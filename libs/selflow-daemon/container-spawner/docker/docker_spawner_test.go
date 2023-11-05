package docker

import (
	"context"
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/selflow/selflow/libs/selflow-daemon/steps/container"
	"io"
	"reflect"
	"sort"
	"strings"
	"testing"
)

type mockedDockerClient struct {
	client.APIClient
	ImagePullReadCloser io.ReadCloser
	ImagePullError      error
}

func (m mockedDockerClient) ImagePull(_ context.Context, _ string, _ types.ImagePullOptions) (io.ReadCloser, error) {
	return m.ImagePullReadCloser, m.ImagePullError
}

func TestNewSpawner(t *testing.T) {
	type args struct {
		dockerClient client.APIClient
	}
	tests := []struct {
		name string
		args args
		want container.ContainerSpawner
	}{
		{
			name: "default",
			args: args{&mockedDockerClient{}},
			want: &spawner{&mockedDockerClient{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSpawner(tt.args.dockerClient); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSpawner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildEnvString(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "default",
			args: args{
				key:   "MY_KEY",
				value: "MY_VALUE",
			},
			want: "MY_KEY=MY_VALUE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildEnvString(tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("buildEnvString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildEnvMap(t *testing.T) {
	type args struct {
		environmentVariables map[string]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "default",
			args: args{
				environmentVariables: map[string]string{
					"MY_ENV":  "MY_VALUE",
					"MY_ENV2": "MY_VALUE2",
				},
			},
			want: []string{"MY_ENV=MY_VALUE", "MY_ENV2=MY_VALUE2"},
		},
		{
			name: "no env",
			args: args{
				environmentVariables: map[string]string{},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildEnvMap(tt.args.environmentVariables)
			sort.Strings(got)
			sort.Strings(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildEnvMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_spawner_pullDockerImage(t *testing.T) {
	type fields struct {
		dockerClient client.APIClient
	}
	type args struct {
		ctx       context.Context
		imageName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "can not pull docker image",
			args: args{
				ctx:       context.TODO(),
				imageName: "toto",
			},
			fields: fields{
				dockerClient: mockedDockerClient{
					ImagePullError: errors.New("some-error"),
				},
			},
			wantErr: true,
		},
		{
			name: "image pulled",
			args: args{
				ctx:       context.TODO(),
				imageName: "toto",
			},
			fields: fields{
				dockerClient: mockedDockerClient{
					ImagePullReadCloser: io.NopCloser(strings.NewReader("Hello World"))},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &spawner{
				dockerClient: tt.fields.dockerClient,
			}
			if err := d.pullDockerImage(tt.args.ctx, tt.args.imageName); (err != nil) != tt.wantErr {
				t.Errorf("pullDockerImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
