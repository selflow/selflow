package main

import (
	"context"
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	cs "github.com/selflow/selflow/pkg/container-spawner"
	"testing"
	"time"
)

type fields struct {
	daemonName              string
	daemonPort              string
	daemonNetworkName       string
	daemonDockerImage       string
	daemonBaseDirectory     string
	daemonHostBaseDirectory string
	daemonIsDebug           bool
	daemonDebugPort         string
	dockerClient            ContainerSpawner
	dockerOpts              []client.Opt
}

type mockDockerClient struct {
	ContainerSpawner
	NetworkInspectNet types.NetworkResource
	NetworkInspectErr error

	NetworkCreateNet types.NetworkCreateResponse
	NetworkCreateErr error

	ContainerStopErr error

	ContainerRemoveErr error

	SpawnAsyncContainerId string
	SpawnAsyncError       error

	DaemonHostResponse string

	ContainerInspectContainer types.ContainerJSON
	ContainerInspectError     error
}

func (m mockDockerClient) NetworkInspect(_ context.Context, _ string, _ types.NetworkInspectOptions) (types.NetworkResource, error) {
	return m.NetworkInspectNet, m.NetworkInspectErr
}

func (m mockDockerClient) NetworkCreate(_ context.Context, _ string, _ types.NetworkCreate) (types.NetworkCreateResponse, error) {
	return m.NetworkCreateNet, m.NetworkCreateErr
}

func (m mockDockerClient) ContainerStop(_ context.Context, _ string, _ *time.Duration) error {
	return m.ContainerStopErr
}

func (m mockDockerClient) ContainerRemove(_ context.Context, _ string, _ types.ContainerRemoveOptions) error {
	return m.ContainerRemoveErr
}

func (m mockDockerClient) SpawnAsync(_ context.Context, _ *cs.SpawnConfig) (string, error) {
	return m.SpawnAsyncContainerId, m.SpawnAsyncError
}

func (m mockDockerClient) DaemonHost() string {
	return m.DaemonHostResponse
}

func (m mockDockerClient) ContainerInspect(_ context.Context, _ string) (types.ContainerJSON, error) {
	return m.ContainerInspectContainer, m.ContainerInspectError
}

type mockNotFoundErr struct {
}

func (m mockNotFoundErr) NotFound() bool {
	return true
}
func (m mockNotFoundErr) Error() string {
	return "not-found error"
}

func Test_selflowClient_clearContainer(t *testing.T) {

	type args struct {
		ctx           context.Context
		containerName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "container does not exists",
			fields: fields{
				dockerClient: mockDockerClient{
					ContainerStopErr: mockNotFoundErr{},
				},
			},
			wantErr: false,
		},
		{
			name: "fail to stop container",
			fields: fields{
				dockerClient: mockDockerClient{
					ContainerStopErr: errors.New("fail to stop container"),
				},
			},
			wantErr: true,
		},
		{
			name: "fail to remove the container",
			fields: fields{
				dockerClient: mockDockerClient{
					ContainerRemoveErr: errors.New("fail to remove the container"),
				},
			},
			wantErr: true,
		},
		{
			name: "container removed",
			fields: fields{
				dockerClient: mockDockerClient{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := &selflowClient{
				daemonName:              tt.fields.daemonName,
				daemonPort:              tt.fields.daemonPort,
				daemonNetworkName:       tt.fields.daemonNetworkName,
				daemonDockerImage:       tt.fields.daemonDockerImage,
				daemonBaseDirectory:     tt.fields.daemonBaseDirectory,
				daemonHostBaseDirectory: tt.fields.daemonHostBaseDirectory,
				daemonIsDebug:           tt.fields.daemonIsDebug,
				daemonDebugPort:         tt.fields.daemonDebugPort,
				dockerClient:            tt.fields.dockerClient,
				dockerOpts:              tt.fields.dockerOpts,
			}
			if err := sc.clearDaemon(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("clearDaemon() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_selflowClient_createDaemon(t *testing.T) {
	type args struct {
		ctx           context.Context
		containerName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "creation succeeded",
			fields: fields{
				dockerClient: mockDockerClient{
					DaemonHostResponse:    "unix:///var/run/docker.sock",
					SpawnAsyncContainerId: "toto",
				},
				daemonPort:      "1111",
				daemonIsDebug:   true,
				daemonDebugPort: "2222",

				daemonName:              "toto",
				daemonBaseDirectory:     "/I/need/more/tea",
				daemonNetworkName:       "some-daemon-network",
				daemonDockerImage:       "some-daemon-docker-image",
				daemonHostBaseDirectory: "/I/need/more/coffee",
			},
			wantErr: false,
			want:    "toto",
		},
		{
			name: "creation fails",
			fields: fields{
				dockerClient: mockDockerClient{
					DaemonHostResponse: "unix:///var/run/docker.sock",
					SpawnAsyncError:    errors.New("some-error"),
				},
				daemonPort:      "1111",
				daemonIsDebug:   true,
				daemonDebugPort: "2222",

				daemonName:              "toto",
				daemonBaseDirectory:     "/I/need/more/tea",
				daemonNetworkName:       "some-daemon-network",
				daemonDockerImage:       "some-daemon-docker-image",
				daemonHostBaseDirectory: "/I/need/more/coffee",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := &selflowClient{
				daemonName:              tt.fields.daemonName,
				daemonPort:              tt.fields.daemonPort,
				daemonNetworkName:       tt.fields.daemonNetworkName,
				daemonDockerImage:       tt.fields.daemonDockerImage,
				daemonBaseDirectory:     tt.fields.daemonBaseDirectory,
				daemonHostBaseDirectory: tt.fields.daemonHostBaseDirectory,
				daemonIsDebug:           tt.fields.daemonIsDebug,
				daemonDebugPort:         tt.fields.daemonDebugPort,
				dockerClient:            tt.fields.dockerClient,
				dockerOpts:              tt.fields.dockerOpts,
			}
			got, err := sc.createDaemon(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("createDaemon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("createDaemon() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_selflowClient_createNetworkIfNotExists(t *testing.T) {
	type args struct {
		ctx         context.Context
		networkName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "network exists",
			fields: fields{
				dockerClient: mockDockerClient{
					NetworkInspectNet: types.NetworkResource{ID: "toto"},
				},
			},
			want: "toto",
		},
		{
			name: "can not create network",
			fields: fields{
				dockerClient: mockDockerClient{
					NetworkInspectErr: errors.New("can not create network"),
				},
			},
			wantErr: true,
		},
		{
			name: "network not found and creation failed",
			fields: fields{
				dockerClient: mockDockerClient{
					NetworkInspectErr: mockNotFoundErr{},
					NetworkCreateErr:  errors.New("network create err"),
				},
			},
			wantErr: true,
		},
		{
			name: "network not found and creation succeeded",
			fields: fields{
				dockerClient: mockDockerClient{
					NetworkInspectErr: mockNotFoundErr{},
					NetworkCreateNet: types.NetworkCreateResponse{
						ID: "toto",
					},
				},
			},
			want: "toto",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := &selflowClient{
				daemonName:              tt.fields.daemonName,
				daemonPort:              tt.fields.daemonPort,
				daemonNetworkName:       tt.fields.daemonNetworkName,
				daemonDockerImage:       tt.fields.daemonDockerImage,
				daemonBaseDirectory:     tt.fields.daemonBaseDirectory,
				daemonHostBaseDirectory: tt.fields.daemonHostBaseDirectory,
				daemonIsDebug:           tt.fields.daemonIsDebug,
				daemonDebugPort:         tt.fields.daemonDebugPort,
				dockerClient:            tt.fields.dockerClient,
				dockerOpts:              tt.fields.dockerOpts,
			}
			got, err := sc.createNetworkIfNotExists(tt.args.ctx, tt.args.networkName)
			if (err != nil) != tt.wantErr {
				t.Errorf("createNetworkIfNotExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("createNetworkIfNotExists() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_selflowClient_getRunningDaemon(t *testing.T) {
	type args struct {
		ctx           context.Context
		containerName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "container found",
			fields: fields{
				dockerClient: mockDockerClient{
					ContainerInspectContainer: types.ContainerJSON{
						ContainerJSONBase: &types.ContainerJSONBase{ID: "toto"},
					},
				},
			},
			want: "toto",
		},
		{
			name: "container not found",
			fields: fields{
				dockerClient: mockDockerClient{
					ContainerInspectError: errors.New("some-error"),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := &selflowClient{
				daemonName:              tt.fields.daemonName,
				daemonPort:              tt.fields.daemonPort,
				daemonNetworkName:       tt.fields.daemonNetworkName,
				daemonDockerImage:       tt.fields.daemonDockerImage,
				daemonBaseDirectory:     tt.fields.daemonBaseDirectory,
				daemonHostBaseDirectory: tt.fields.daemonHostBaseDirectory,
				daemonIsDebug:           tt.fields.daemonIsDebug,
				daemonDebugPort:         tt.fields.daemonDebugPort,
				dockerClient:            tt.fields.dockerClient,
				dockerOpts:              tt.fields.dockerOpts,
			}
			got, err := sc.getRunningDaemon(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("getRunningDaemon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getRunningDaemon() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_selflowClient_startDaemon(t *testing.T) {
	type args struct {
		ctx           context.Context
		selflowClient *selflowClient
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "daemon already started",
			fields: fields{
				dockerClient: mockDockerClient{
					ContainerInspectContainer: types.ContainerJSON{
						ContainerJSONBase: &types.ContainerJSONBase{ID: "toto"},
					},
				},
			},
			want: "toto",
		},
		{
			name: "fail to clear daemon",
			fields: fields{
				dockerClient: mockDockerClient{
					ContainerInspectError: errors.New("some-error"),
					ContainerStopErr:      errors.New("another-error"),
				},
			},
			wantErr: true,
		},
		{
			name: "fail to create daemon",
			fields: fields{
				dockerClient: mockDockerClient{
					ContainerInspectError: errors.New("some-error"),
					SpawnAsyncError:       errors.New("another-error"),
				},
			},
			wantErr: true,
		},
		{
			name: "daemon created",
			fields: fields{
				dockerClient: mockDockerClient{
					ContainerInspectError: errors.New("some-error"),
					SpawnAsyncContainerId: "toto",
				},
			},
			want: "toto",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := &selflowClient{
				daemonName:              tt.fields.daemonName,
				daemonPort:              tt.fields.daemonPort,
				daemonNetworkName:       tt.fields.daemonNetworkName,
				daemonDockerImage:       tt.fields.daemonDockerImage,
				daemonBaseDirectory:     tt.fields.daemonBaseDirectory,
				daemonHostBaseDirectory: tt.fields.daemonHostBaseDirectory,
				daemonIsDebug:           tt.fields.daemonIsDebug,
				daemonDebugPort:         tt.fields.daemonDebugPort,
				dockerClient:            tt.fields.dockerClient,
				dockerOpts:              tt.fields.dockerOpts,
			}
			got, err := sc.startDaemon(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("startDaemon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("startDaemon() got = %v, want %v", got, tt.want)
			}
		})
	}
}
