package container_spawner

import (
	"fmt"
	"github.com/docker/docker/client"
	"sync"
)

var (
	dockerClient client.APIClient
	cliMutex     sync.Mutex
)

func GetClient() (client.APIClient, error) {
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
