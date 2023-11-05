package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func NewRecreateDaemonCommand(selflowClient *selflowClient) *cobra.Command {
	return &cobra.Command{
		Use:   "recreate-daemon",
		Short: "Recreate the Selflow Daemon",
		Run: func(cmd *cobra.Command, _ []string) {
			if err := startRecreateDaemon(selflowClient); err != nil {
				if _, err = fmt.Fprintf(os.Stderr, "%v\n", err); err != nil {
					panic(err)
				}
				os.Exit(1)
			}
		},
	}
}

func startRecreateDaemon(selflowClient *selflowClient) error {
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

	return nil
}
