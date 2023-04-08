package main

import (
	"context"
	"github.com/spf13/cobra"
	"log"
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

	daemonId, err := selflowClient.startDaemon(ctx, selflowClient)
	if err != nil {
		panic(err)
	}

	log.Println(daemonId)

	return nil

}
