package main

import (
	"context"
	"github.com/selflow/selflow/cmd/selflow-daemon/server/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

func NewRunCommand(selflowClient *selflowClient) *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "Start a workflow",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := startRun(selflowClient, args[0]); err != nil {
				panic(err)
			}
		},
	}
}

func startRun(selflowClient *selflowClient, fileName string) error {
	ctx := context.Background()
	err := selflowClient.init()
	if err != nil {
		return err
	}

	_, err = selflowClient.createNetworkIfNotExists(ctx, selflowClient.daemonNetworkName)
	if err != nil {
		return err
	}

	_, err = selflowClient.startDaemon(ctx)
	if err != nil {
		panic(err)
	}

	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	conn, err := grpc.Dial("localhost:10011", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() {
		_ = conn.Close()
	}()

	c := proto.NewDaemonClient(conn)

	r, err := c.StartRun(ctx, &proto.StartRun_Request{RunConfig: bytes})

	log.Println(r)
	log.Println(err)

	return nil

}
