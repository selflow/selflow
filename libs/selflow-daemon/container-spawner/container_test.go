package container_spawner

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
	"io"
	"log"
	"reflect"
	"testing"
)

type mockDockerClient struct {
	client.APIClient
	createFunctionCallCount uint
	createFunction          func(count uint) (container.CreateResponse, error)
	pullFunction            func() (io.ReadCloser, error)
	startFunction           func() error
	logsFunction            func() (io.ReadCloser, error)
}

func (mdc *mockDockerClient) ContainerCreate(_ context.Context, _ *container.Config, _ *container.HostConfig, _ *network.NetworkingConfig, _ *specs.Platform, _ string) (container.CreateResponse, error) {
	defer func() {
		mdc.createFunctionCallCount++
	}()
	return mdc.createFunction(mdc.createFunctionCallCount)
}

func (mdc *mockDockerClient) ImagePull(ctx context.Context, ref string, options types.ImagePullOptions) (io.ReadCloser, error) {
	if mdc.pullFunction != nil {
		return mdc.pullFunction()
	}
	return io.NopCloser(bytes.NewReader([]byte("Coucou c'est moi moumou la reine des mouettes !"))), nil
}

func (mdc *mockDockerClient) ContainerStart(ctx context.Context, container string, options types.ContainerStartOptions) error {
	if mdc.startFunction != nil {
		return mdc.startFunction()
	}
	return nil
}

func makeFakeDockerLog(b string) []byte {
	size := make([]byte, 4)
	binary.BigEndian.PutUint32(size, uint32(len(b)))
	v := []byte{byte(1), 0x0, 0x0, 0x0}
	return append(append(v, size...), b...)
}

func (mdc *mockDockerClient) ContainerLogs(ctx context.Context, container string, options types.ContainerLogsOptions) (io.ReadCloser, error) {
	if mdc.logsFunction != nil {
		return mdc.logsFunction()
	}
	return io.NopCloser(bytes.NewReader(makeFakeDockerLog("Vim > Emacs"))), nil
}

type mockNotFoundErr struct {
	error
}

func (mockNotFoundErr) NotFound() bool {
	return true
}

func Test_createContainer(t *testing.T) {
	log.SetOutput(io.Discard)
	type args struct {
		ctx    context.Context
		cli    client.APIClient
		config *SpawnConfig
	}
	tests := []struct {
		name    string
		args    args
		want    *pluginContainer
		wantErr bool
	}{
		{
			name: "basic config",
			args: args{
				ctx: context.TODO(),
				cli: &mockDockerClient{
					createFunction: func(count uint) (container.CreateResponse, error) {
						return container.CreateResponse{
							ID:       "12345",
							Warnings: nil,
						}, nil
					},
				},
				config: &SpawnConfig{
					Image:               "brioche",
					ContainerName:       "painperdu",
					ContainerLogsWriter: io.Discard,
				},
			},
			want: &pluginContainer{
				containerName: "painperdu",
				containerId:   "12345",
			},
			wantErr: false,
		},
		{
			name: "image does not exists at all",
			args: args{
				ctx: context.TODO(),
				cli: &mockDockerClient{
					createFunction: func(count uint) (container.CreateResponse, error) {
						return container.CreateResponse{}, mockNotFoundErr{}
					},
					pullFunction: func() (io.ReadCloser, error) {
						return nil, mockNotFoundErr{}
					},
				},
				config: &SpawnConfig{
					Image:               "iwanttoexist",
					ContainerName:       "painperdu",
					ContainerLogsWriter: io.Discard,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "fail to initialize container",
			args: args{
				ctx: context.TODO(),
				cli: &mockDockerClient{
					createFunction: func(count uint) (container.CreateResponse, error) {
						return container.CreateResponse{}, errors.New("fail to create the container")
					},
				},
				config: &SpawnConfig{
					Image:               "iwillfail",
					ContainerName:       "painperdu",
					ContainerLogsWriter: io.Discard,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "image must be fetched",
			args: args{
				ctx: context.TODO(),
				cli: &mockDockerClient{
					createFunction: func(count uint) (container.CreateResponse, error) {
						if count == 0 {
							return container.CreateResponse{}, mockNotFoundErr{}
						}
						return container.CreateResponse{
							ID:       "12345",
							Warnings: nil,
						}, nil
					},
				},
				config: &SpawnConfig{
					Image:               "IToldYouIExist",
					ContainerName:       "painperdu",
					ContainerLogsWriter: io.Discard,
				},
			},
			want: &pluginContainer{
				containerName: "painperdu",
				containerId:   "12345",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createContainer(tt.args.ctx, tt.args.cli, tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("createContainer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createContainer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_startContainer(t *testing.T) {
	type args struct {
		ctx    context.Context
		cli    client.APIClient
		config *pluginContainer
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "default",
			args: args{
				ctx: context.TODO(),
				cli: &mockDockerClient{},
				config: &pluginContainer{
					containerName: "Merveilleux",
					containerId:   "merveilleux",
				},
			},
			wantErr: false,
		},
		{
			name: "container start error",
			args: args{
				ctx: context.TODO(),
				cli: &mockDockerClient{
					startFunction: func() error {
						return errors.New("start error")
					},
				},
				config: &pluginContainer{
					containerName: "Mille-feuilles",
					containerId:   "millefeuilles",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := startContainer(tt.args.ctx, tt.args.cli, tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("startContainer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_transferContainerLogs(t *testing.T) {
	type args struct {
		ctx    context.Context
		cli    client.APIClient
		config *pluginContainer
	}
	tests := []struct {
		name       string
		args       args
		wantWriter string
		wantErr    bool
	}{
		{
			name: "default",
			args: args{
				ctx: context.TODO(),
				cli: &mockDockerClient{},
				config: &pluginContainer{
					containerName: "fraisier",
					containerId:   "pomme-damour",
				},
			},
			wantWriter: "Vim > Emacs",
		},
		{
			name: "can not access docker logs",
			args: args{
				ctx: context.TODO(),
				cli: &mockDockerClient{
					logsFunction: func() (io.ReadCloser, error) {
						return nil, errors.New("sflog transfer failed")
					},
				},
				config: &pluginContainer{
					containerName: "fraisier",
					containerId:   "pomme-damour",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			err := transferContainerLogs(tt.args.ctx, tt.args.cli, tt.args.config, writer)
			if (err != nil) != tt.wantErr {
				t.Errorf("transferContainerLogs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("transferContainerLogs() gotWriter = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}
