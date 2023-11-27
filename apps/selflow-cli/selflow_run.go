package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/selflow/selflow/apps/selflow-daemon/server/proto"
	"github.com/selflow/selflow/libs/selflow-daemon/sfenvironment"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"hash/fnv"
	"io"
	"log/slog"
	"os"
	"time"
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
		Run: func(cmd *cobra.Command, args []string) {
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

	bytes, err := os.ReadFile(fileName)
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

	runId, err := c.StartRun(ctx, &proto.StartRun_Request{RunConfig: bytes})

	if err != nil {
		return fmt.Errorf("fail to start run: %v", err)
	}

	fmt.Printf("Run started with Id: %v\n", runId)

	stream, err := c.GetLogStream(ctx, &proto.GetLogStream_Request{RunId: runId.GetRunId()})
	if err != nil {
		return err
	}

	for {
		l, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("fail to read from stream: %v", err)
		}

		if sfenvironment.UseJsonLogs {
			metadata := map[string]interface{}{}
			if err = json.Unmarshal(l.GetMetadata(), &metadata); err != nil {
				metadata = nil
			}

			delete(metadata, "msg")
			delete(metadata, "level")
			delete(metadata, "source")
			delete(metadata, "time")

			slog.Log(ctx, slog.LevelInfo, l.GetMessage(), slog.String("stepId", l.GetName()), slog.Any("metadata", metadata))
		} else {
			fmt.Printf("%s%v\t[%4v]\t%v:\t%v\033[0m\n", GetAnsiColor(l.GetName()), CastTimeLog(l.GetDateTime()), l.GetLevel(), l.GetName(), l.GetMessage())
		}
	}

	return nil

}
