package main

import (
	"context"
	"fmt"
	"hash/fnv"
	"io"
	"log/slog"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/selflow/selflow/apps/selflow-cli/models"
	"github.com/selflow/selflow/apps/selflow-daemon/server/proto"
	"github.com/selflow/selflow/libs/selflow-daemon/sfenvironment"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewRunCommand(selflowClient *selflowClient) *cobra.Command {
	return &cobra.Command{
		Use:   "run ./filename",
		Short: "Start a workflow on the Selflow-Daemon and wait for the end of its execution.",
		Long: "Start a workflow on the Selflow-Daemon and wait for the end of its execution.\n\n" +
			"If the command is stopped, the workflow will not be stopped. " +
			"The commands also gives logs about what is happening.\n\n" +
			"The workflow file must follow the selflow workflow syntax (https://selflow.github.io/selflow/docs/workflow-syntax).\n\n" +
			"It can be written using __YAML__ or __JSON__.",
		Args: cobra.MinimumNArgs(1),
		Run: func(_ *cobra.Command, args []string) {
			if err := startRun(selflowClient, args[0]); err != nil {
				if _, err = fmt.Fprintf(os.Stderr, "%v\n", err); err != nil {
					panic(err)
				}
				os.Exit(1)
			}
		},
	}
}

func GetAnsiColor(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	hash := h.Sum32()

	return fmt.Sprintf("\033[3%dm", (hash%6)+1)
}

func CastTimeLog(logTimeAsString string) string {
	t, err := time.Parse(time.RFC3339Nano, logTimeAsString)
	if err != nil {
		return "invalid-date"
	}
	return t.Format(time.RFC3339)
}

func startRun(selflowClient *selflowClient, fileName string) error {
	ctx := context.Background()
	err := selflowClient.init()
	if err != nil {
		return fmt.Errorf("fail to initialize the client: %v", err)
	}

	_, err = selflowClient.createNetworkIfNotExists(ctx, selflowClient.daemonNetworkName)
	if err != nil {
		return fmt.Errorf("fail to create the daemon network: %v", err)
	}

	_, err = selflowClient.startDaemon(ctx)
	if err != nil {
		return fmt.Errorf("fail to start the daemon: %v", err)
	}

	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("fail to read the provided file: %v", err)
	}

	conn, err := grpc.Dial("localhost:10011", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("fail to connect to the daemon: %v", err)
	}
	defer func() {
		_ = conn.Close()
	}()

	c := proto.NewDaemonClient(conn)

	runId, err := c.StartRun(ctx, &proto.StartRun_Request{RunConfig: fileContent})

	if err != nil {
		return fmt.Errorf("fail to start run: %v", err)
	}

	fmt.Printf("Run started with Id: %v\n", runId)

	stream, err := c.GetLogStream(ctx, &proto.GetLogStream_Request{RunId: runId.GetRunId()})
	if err != nil {
		return err
	}

	workflowLogsReader, workflowLogsWriter := io.Pipe()
	workflowTerminated := make(chan bool, 1)
	stepStatusCh := make(chan models.StepStatus)

	// updateStatus is function that calls the daemon to get the status of each step of a run and applies it
	updateStatus := func() error {
		runStatus, err := c.GetRunStatus(ctx, &proto.GetRunStatus_Request{RunId: runId.GetRunId()})
		if err != nil {
			return err
		}

		for stepId, stepStatus := range runStatus.State {
			stepStatusCh <- models.StepStatus{
				StepId: stepId,
				Status: stepStatus.Name,
			}
		}
		return nil

	}

	// This goroutine updates status of the run
	go func() {
		ticker := time.NewTicker(50 * time.Millisecond)
		_ = updateStatus()

	EventLoop:
		for {
			select {
			case <-ticker.C:
				_ = updateStatus()

			case <-workflowTerminated:
				break EventLoop
			}
		}
	}()

	// This goroutine handles logs recieved over a GRPC stream
	go func(ctx context.Context) {
		for {
			l, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				slog.ErrorContext(ctx, "Fail to read from stream", "error", err)
				break
			}

			_, err = workflowLogsWriter.Write(l.GetMetadata())
			if err != nil {
				slog.ErrorContext(ctx, "Fail to write log")
			}
			_, _ = workflowLogsWriter.Write([]byte("\n"))

		}

		err := updateStatus()
		workflowTerminated <- err != nil
	}(ctx)

	//----------------------------//
	//--- Initialize Bubbletea ---//
	//----------------------------//

	model := models.NewRunModel(ctx, workflowTerminated, workflowLogsReader, stepStatusCh)

	var bubbleteaOptions []tea.ProgramOption
	// Handle sessions without tty
	if !term.IsTerminal(int(os.Stdout.Fd())) {
		slog.DebugContext(ctx, "Not tty detected")
		bubbleteaOptions = append(bubbleteaOptions, tea.WithInput(nil))
	}
	if sfenvironment.UseJsonLogs {
		bubbleteaOptions = append(bubbleteaOptions, tea.WithoutRenderer())
	}

	//---------------------//
	//--- Start process ---//
	//---------------------//

	p := tea.NewProgram(model, bubbleteaOptions...)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error starting Bubble Tea program:", err)
		os.Exit(1)
	}
	return nil

}
