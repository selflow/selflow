package container_spawner

import (
	"context"
	"github.com/docker/docker/client"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
	"testing"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
)

type mockClient struct {
	client.APIClient
}

func (m *mockClient) ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkConfig *network.NetworkingConfig, platform *specs.Platform, containerName string) (container.ContainerCreateCreatedBody, error) {
	return container.ContainerCreateCreatedBody{ID: "abcdefg"}, nil
}

func (m *mockClient) IsErrNotFound(err error) bool {
	return false
}

func TestCreateContainer(t *testing.T) {
	cli := &mockClient{}
	ctx := context.Background()
	config := &SpawnConfig{
		Environment:   map[string]string{"ENV1": "value1", "ENV2": "value2"},
		Image:         "myimage",
		Mounts:        []Mountable{},
		ContainerName: "mycontainer",
	}

	pc, err := createContainer(ctx, cli, config)
	if err != nil {
		t.Fatalf("StartContainerDetached failed: %v", err)
	}

	if pc.containerName != config.ContainerName {
		t.Errorf("containerName expected %s but got %s", config.ContainerName, pc.containerName)
	}

	if pc.containerId != "abcdefg" {
		t.Errorf("containerId expected abcdefg but got %s", pc.containerId)
	}
}

func TestCreateContainerImageNotFound(t *testing.T) {
	cli := &mockClient{}
	ctx := context.Background()
	config := &SpawnConfig{
		Environment:   map[string]string{"ENV1": "value1", "ENV2": "value2"},
		Image:         "myimage",
		Mounts:        []Mountable{},
		ContainerName: "mycontainer",
	}

	pc, err := createContainer(ctx, cli, config)
	if err != nil {
		t.Fatalf("StartContainerDetached failed: %v", err)
	}

	if pc.containerName != config.ContainerName {
		t.Errorf("containerName expected %s but got %s", config.ContainerName, pc.containerName)
	}

	if pc.containerId != "abcdefg" {
		t.Errorf("containerId expected abcdefg but got %s", pc.containerId)
	}
}
