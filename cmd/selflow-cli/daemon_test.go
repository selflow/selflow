package main

import (
	"context"
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
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
	dockerClient            client.APIClient
	dockerOpts              []client.Opt
}

type mockDockerClient struct {
	client.APIClient
	NetworkInspectNet types.NetworkResource
	NetworkInspectErr error

	NetworkCreateNet types.NetworkCreateResponse
	NetworkCreateErr error

	ContainerStopErr error

	ContainerRemoveErr error
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
			if err := sc.clearContainer(tt.args.ctx, tt.args.containerName); (err != nil) != tt.wantErr {
				t.Errorf("clearContainer() error = %v, wantErr %v", err, tt.wantErr)
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
		// TODO: Add test cases.
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
			got, err := sc.createDaemon(tt.args.ctx, tt.args.containerName)
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
		// TODO: Add test cases.
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
			got, err := sc.getRunningDaemon(tt.args.ctx, tt.args.containerName)
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
		// TODO: Add test cases.
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
			got, err := sc.startDaemon(tt.args.ctx, tt.args.selflowClient)
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
