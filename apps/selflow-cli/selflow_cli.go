package main

import (
	"context"
	"github.com/selflow/selflow/libs/core/sflog"
	"github.com/spf13/cobra"
	"log"
)

const (
	ForceDaemonRecreationFlag = "recreate-daemon"
)

var rootCmd = &cobra.Command{
	Use:   "selflow",
	Short: "Selflow is a workflow orchestration tool",
	Long: `Selflow is a workflow orchestration tool.
This is a command line interface to interact with the Selflow Daemon.

Built with Love by Anthony-Jhoiro (https://github.com/Anthony-Jhoiro)`,
}

func init() {
	sflog.Init()

	sc := newSelflowClient()

	rootCmd.PersistentFlags().Bool(ForceDaemonRecreationFlag, false, "Kill and recreate the daemon if it already exists")

	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Lookup(ForceDaemonRecreationFlag).Changed {
			err := sc.init()
			if err != nil {
				log.Fatalf("fail to initialize selflow: %v", err)
			}
			err = sc.clearDaemon(context.Background())
			if err != nil {
				log.Fatalf("fail to stop daemon : %v", err)
			}
		}
	}

	rootCmd.AddCommand(
		NewRunCommand(sc),
		NewRecreateDaemonCommand(sc),
		NewStatusCommand(sc),
	)
}

func main() {

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf(err.Error())
	}
}
