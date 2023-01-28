package main

import (
	"fmt"
	"github.com/docker/docker/client"
	"sync"
)

var (
	dockerClient *client.Client
	cliMutex     sync.Mutex
)

func GetClient() (*client.Client, error) {
	if dockerClient == nil {
		var err error

		cliMutex.Lock()
		dockerClient, err = client.NewClientWithOpts(client.FromEnv)
		if err != nil {
			return nil, fmt.Errorf("fail to create docker client : %v", err)
		}

		cliMutex.Unlock()
	}

	return dockerClient, nil
}
