package main

import (
	"errors"
	"github.com/docker/docker/client"
	"github.com/selflow/selflow/internal/sfenvironment"
)

var ContainerNotFound = errors.New("container not found")

type selflowClient struct {
	daemonName              string
	daemonPort              string
	daemonNetworkName       string
	daemonDockerImage       string
	daemonBaseDirectory     string
	daemonHostBaseDirectory string

	daemonIsDebug   bool
	daemonDebugPort string

	dockerClient client.APIClient
	dockerOpts   []client.Opt
}

func newSelflowClient() *selflowClient {
	daemonDebugPort := sfenvironment.GetDaemonDebugPort()
	return &selflowClient{
		daemonName:              sfenvironment.GetDaemonName(),
		daemonPort:              sfenvironment.GetDaemonPort(),
		daemonNetworkName:       sfenvironment.GetDaemonNetwork(),
		daemonDockerImage:       sfenvironment.GetDaemonImage(),
		daemonBaseDirectory:     sfenvironment.GetDaemonBaseDirectory(),
		daemonHostBaseDirectory: sfenvironment.GetDaemonHostBaseDirectory(),

		daemonIsDebug:   daemonDebugPort != "",
		daemonDebugPort: daemonDebugPort,

		dockerOpts: []client.Opt{client.FromEnv},
	}
}

func (sc *selflowClient) init() error {
	var err error
	sc.dockerClient, err = client.NewClientWithOpts(sc.dockerOpts...)
	if err != nil {
		return err
	}

	return err
}
