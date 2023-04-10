package main

import (
	"errors"
	"github.com/docker/docker/client"
	"testing"
)

func Test_newSelflowClient(t *testing.T) {
	tests := []struct {
		name string
		want *selflowClient
	}{
		{
			name: "default",
			want: &selflowClient{
				daemonName:              "selflow-daemon",
				daemonPort:              "10011",
				daemonNetworkName:       "selflow-daemon",
				daemonDockerImage:       "selflow-daemon:latest",
				daemonBaseDirectory:     "/etc/selflow",
				daemonHostBaseDirectory: "/etc/selflow",

				daemonIsDebug:   false,
				daemonDebugPort: "",

				dockerOpts: []client.Opt{client.FromEnv},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newSelflowClient()
			if got.daemonName != tt.want.daemonName {
				t.Errorf("newSelflowClient() daemonName = %v, want %v", got.daemonName, tt.want.daemonName)
			}
			if got.daemonPort != tt.want.daemonPort {
				t.Errorf("newSelflowClient() daemonPort = %v, want %v", got.daemonPort, tt.want.daemonPort)
			}
			if got.daemonNetworkName != tt.want.daemonNetworkName {
				t.Errorf("newSelflowClient() daemonNetworkName = %v, want %v", got.daemonNetworkName, tt.want.daemonNetworkName)
			}
			if got.daemonDockerImage != tt.want.daemonDockerImage {
				t.Errorf("newSelflowClient() daemonDockerImage = %v, want %v", got.daemonDockerImage, tt.want.daemonDockerImage)
			}
			if got.daemonBaseDirectory != tt.want.daemonBaseDirectory {
				t.Errorf("newSelflowClient() daemonBaseDirectory = %v, want %v", got.daemonBaseDirectory, tt.want.daemonBaseDirectory)
			}
			if got.daemonHostBaseDirectory != tt.want.daemonHostBaseDirectory {
				t.Errorf("newSelflowClient() daemonHostBaseDirectory = %v, want %v", got.daemonHostBaseDirectory, tt.want.daemonHostBaseDirectory)
			}
			if got.daemonIsDebug != tt.want.daemonIsDebug {
				t.Errorf("newSelflowClient() daemonIsDebug = %v, want %v", got.daemonIsDebug, tt.want.daemonIsDebug)
			}
			if got.daemonDebugPort != tt.want.daemonDebugPort {
				t.Errorf("newSelflowClient() daemonDebugPort = %v, want %v", got.daemonDebugPort, tt.want.daemonDebugPort)
			}
		})
	}
}

func Test_selflowClient_init(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "default",
			fields: fields{
				dockerOpts: nil,
			},
			wantErr: false,
		},
		{
			name: "default",
			fields: fields{
				dockerOpts: []client.Opt{func(c *client.Client) error {
					return errors.New("some error")
				}},
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
			if err := sc.init(); (err != nil) != tt.wantErr {
				t.Errorf("init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
