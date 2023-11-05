package main

import (
	"context"
	"fmt"
	proto2 "github.com/selflow/selflow/apps/selflow-daemon/server/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

func NewStatusCommand(selflowClient *selflowClient) *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Get the state of a workflow",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := startStatus(selflowClient, args[0]); err != nil {
				if _, err = fmt.Fprintf(os.Stderr, "%v\n", err); err != nil {
					panic(err)
				}
				os.Exit(1)
			}
		},
	}
}

func startStatus(selflowClient *selflowClient, runId string) error {
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

	conn, err := grpc.Dial("localhost:10011", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("fail to connect to the daemon: %v", err)
	}
	defer func() {
		_ = conn.Close()
	}()

	c := proto2.NewDaemonClient(conn)

	response, err := c.GetRunStatus(ctx, &proto2.GetRunStatus_Request{RunId: runId})

	log.Println(response)

	return nil

}
