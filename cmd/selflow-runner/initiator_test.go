package main

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/selflow/selflow/internal/config"
	selflow_runner_proto "github.com/selflow/selflow/internal/selflow-runner-proto"
	mock_selflow_runner_proto "github.com/selflow/selflow/internal/selflow-runner-proto/mock_selflow-runner-proto"
	"github.com/selflow/selflow/pkg/workflow"
	"github.com/selflow/selflow/pkg/workflow/mock_workflow"
	"testing"
)

type workflowBuilderMock struct {
	buildWorkflowError error
	workflow           workflow.Workflow
}

func (b workflowBuilderMock) BuildWorkflow(_ *config.Flow) (workflow.Workflow, error) {
	if b.buildWorkflowError != nil {
		return nil, b.buildWorkflowError
	}
	if b.workflow != nil {
		return b.workflow, nil
	}
	return workflow.NewWorkflow(0), nil
}

const validWorkflowFile = "./testdata/flow.json"
const invalidWorkflowFile = "./testdata/invalid_flow.json"
const notExistingWorkflowFile = "./testdata/not_existing_file.json"

func TestSelflowRunnerPlugin_InitRunner(t *testing.T) {
	type fields struct {
		configFileLocation   string
		workflowInitError    error
		workflowExecuteError error
		buildWorkflowError   error
	}
	type args struct {
		containerSpawner selflow_runner_proto.ContainerSpawner
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "default",
			fields: fields{
				configFileLocation: validWorkflowFile,
			},
			wantErr: false,
			args: args{
				containerSpawner: mock_selflow_runner_proto.NewMockContainerSpawner(&gomock.Controller{}),
			},
		},
		{
			name: "file does not exists",
			fields: fields{
				configFileLocation: notExistingWorkflowFile,
			},
			wantErr: true,
			args: args{
				containerSpawner: mock_selflow_runner_proto.NewMockContainerSpawner(&gomock.Controller{}),
			},
		},
		{
			name: "file is invalid",
			fields: fields{
				configFileLocation: invalidWorkflowFile,
			},
			wantErr: true,
			args: args{
				containerSpawner: mock_selflow_runner_proto.NewMockContainerSpawner(&gomock.Controller{}),
			},
		},
		{
			name: "fail to build workflow",
			fields: fields{
				configFileLocation: validWorkflowFile,
				buildWorkflowError: errors.New("build error"),
			},
			wantErr: true,
			args: args{
				containerSpawner: mock_selflow_runner_proto.NewMockContainerSpawner(&gomock.Controller{}),
			},
		},
		{
			name: "workflow init fails",
			fields: fields{
				configFileLocation: validWorkflowFile,
				workflowInitError:  errors.New("init error"),
			},
			wantErr: true,
			args: args{
				containerSpawner: mock_selflow_runner_proto.NewMockContainerSpawner(&gomock.Controller{}),
			},
		},
		{
			name: "workflow execution fails",
			fields: fields{
				configFileLocation:   validWorkflowFile,
				workflowExecuteError: errors.New("execution error"),
			},
			wantErr: true,
			args: args{
				containerSpawner: mock_selflow_runner_proto.NewMockContainerSpawner(&gomock.Controller{}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			wf := mock_workflow.NewMockWorkflow(ctrl)
			wf.EXPECT().Init().Return(tt.fields.workflowInitError).AnyTimes()
			wf.EXPECT().Execute(gomock.Any()).Return(map[string]map[string]string{}, tt.fields.workflowExecuteError).AnyTimes()

			s := &SelflowRunnerPlugin{
				workflowBuilder: workflowBuilderMock{
					workflow:           wf,
					buildWorkflowError: tt.fields.buildWorkflowError,
				},
				configFileLocation: tt.fields.configFileLocation,
			}
			if err := s.InitRunner(context.TODO(), tt.args.containerSpawner); (err != nil) != tt.wantErr {
				t.Errorf("InitRunner() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
