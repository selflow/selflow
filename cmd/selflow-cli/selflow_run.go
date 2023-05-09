package main

import (
	"context"
	"fmt"
	"github.com/selflow/selflow/cmd/selflow-daemon/server/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"os"
)

func NewRunCommand(selflowClient *selflowClient) *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "Start a workflow",
		Args:  cobra.MinimumNArgs(1),
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
		fmt.Printf("%v [%4v] %v: %v\n", l.GetDateTime(), l.GetLevel(), l.GetName(), l.GetMessage())
	}

	return nil

}
