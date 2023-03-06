package workflow

import (
	"bytes"
	"context"
	"errors"
	"log"
	"os"
	"reflect"
	"testing"
)

func Test_areRequirementsFullFilled(t *testing.T) {
	type args struct {
		step         Step
		dependencies map[Step][]Step
	}
	stepA, _ := makeSimpleStep("step-a")
	stepSuccessB, _ := makeSimpleStep("step-b")
	stepSuccessB.Status = SUCCESS
	stepSuccessC, _ := makeSimpleStep("step-c")
	stepSuccessC.Status = SUCCESS
	stepPendingD, _ := makeSimpleStep("step-c")
	stepPendingD.Status = PENDING

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Step with 2 success requirements",
			args: args{
				step:         stepA,
				dependencies: map[Step][]Step{stepA: {stepSuccessB, stepSuccessC}},
			},
			want: true,
		},
		{
			name: "Step with 1 success and 1 pending requirements",
			args: args{
				step:         stepA,
				dependencies: map[Step][]Step{stepA: {stepSuccessB, stepPendingD}},
			},
			want: false,
		},
		{
			name: "Step with no requirements",
			args: args{
				step:         stepA,
				dependencies: map[Step][]Step{stepA: {}},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areRequirementsFullFilled(tt.args.step, tt.args.dependencies); got != tt.want {
				t.Errorf("areRequirementsFullFilled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeSimpleWorkflow(t *testing.T) {
	tests := []struct {
		name string
		want *SimpleWorkflow
	}{
		{
			name: "Create simple workflow",
			want: &SimpleWorkflow{
				steps:        []Step{},
				status:       CREATED,
				dependencies: map[Step][]Step{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeSimpleWorkflow(1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeSimpleWorkflow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleWorkflow_GetStatus(t *testing.T) {
	type fields struct {
		steps        []Step
		status       Status
		dependencies map[Step][]Step
	}
	tests := []struct {
		name   string
		fields fields
		want   Status
	}{
		{
			name: "Pending status",
			fields: fields{
				status: PENDING,
			},
			want: PENDING,
		},
		{
			name: "Created status",
			fields: fields{
				status: CREATED,
			},
			want: CREATED,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SimpleWorkflow{
				steps:        tt.fields.steps,
				status:       tt.fields.status,
				dependencies: tt.fields.dependencies,
			}
			if got := s.GetStatus(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_joinErrorList(t *testing.T) {
	type args struct {
		errorLst []error
	}
	tests := []struct {
		name           string
		args           args
		wantErr        bool
		wantErrMessage string
	}{
		{
			name:           "Empty errors",
			args:           args{},
			wantErr:        true,
			wantErrMessage: "",
		},
		{
			name:           "2 errors",
			args:           args{errorLst: []error{errors.New("aaa"), errors.New("bbb")}},
			wantErr:        true,
			wantErrMessage: "aaa ; bbb ; ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := joinErrorList(tt.args.errorLst)
			if (err != nil) != tt.wantErr {
				t.Errorf("joinErrorList() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err.Error() != tt.wantErrMessage {
				t.Errorf("joinErrorList() error = %v, wantErrMessage %v", err.Error(), tt.wantErrMessage)
			}
		})
	}
}

func TestSimpleWorkflow_Execute(t *testing.T) {
	type fields struct {
		steps        []Step
		status       Status
		dependencies map[Step][]Step
	}
	stepA, _ := makeSimpleStep("step-a")
	stepB, _ := makeSimpleStep("step-b")
	stepC, _ := makeSimpleStep("step-c")

	stepD, _ := makeSimpleStep("step-d")
	stepE, _ := makeSimpleStep("step-e")
	stepF, _ := makeSimpleStep("step-f")

	errorA, _ := makeSimpleStep("error-a")
	errorB, _ := makeSimpleStep("error-b")
	errorD, _ := makeErrorStep("error-d")

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       map[string]map[string]string
		wantStatus map[Step]Status
		wantErr    bool
	}{
		{
			name: "3 steps execution",
			fields: fields{
				steps:  []Step{stepD, stepE, stepF},
				status: CREATED,
				dependencies: map[Step][]Step{
					stepA: {stepF, stepE},
					stepB: {},
					stepC: {},
				},
			},
			args: args{
				ctx: context.TODO(),
			},
			want: map[string]map[string]string{
				"step-d": {},
				"step-e": {},
				"step-f": {},
			},
			wantErr: false,
			wantStatus: map[Step]Status{
				stepE: SUCCESS,
				stepF: SUCCESS,
				stepD: SUCCESS,
			},
		},
		{
			name: "with cancel",
			fields: fields{
				steps:  []Step{errorA, errorB, errorD},
				status: CREATED,
				dependencies: map[Step][]Step{
					errorD: {},
					errorA: {errorD},
					errorB: {errorD},
				},
			},
			args: args{
				ctx: context.TODO(),
			},
			want: map[string]map[string]string{
				"error-a": {},
				"error-b": {},
				"error-d": {},
			},
			wantErr: false,
			wantStatus: map[Step]Status{
				errorA: CANCELLED,
				errorB: CANCELLED,
				errorD: ERROR,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SimpleWorkflow{
				steps:        tt.fields.steps,
				status:       tt.fields.status,
				dependencies: tt.fields.dependencies,
			}
			got, err := s.Execute(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() got = %v, want %v", got, tt.want)
			}
			for step, status := range tt.wantStatus {
				for _, wfStep := range s.steps {
					if wfStep.GetId() == step.GetId() {
						if wfStep.GetStatus().GetCode() != status.GetCode() {
							t.Errorf("Execute() step  %v status got= %v, want %v", wfStep.GetId(), wfStep.GetStatus().GetName(), status.GetName())
						}
						break
					}
				}
			}
		})
	}
}

func TestSimpleWorkflow_Init(t *testing.T) {
	type fields struct {
		steps        []Step
		status       Status
		dependencies map[Step][]Step
	}
	type args struct {
		context context.Context
	}
	A, _ := makeSimpleStep("step-a")
	B, _ := makeSimpleStep("step-b")
	C, _ := makeSimpleStep("step-c")
	D, _ := makeSimpleStep("step-d")

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "With no dependency",
			fields: fields{
				steps: []Step{A, B, C, D},
				dependencies: map[Step][]Step{
					A: {},
					B: {},
					C: {},
					D: {},
				},
			},
			args:    args{},
			wantErr: false,
		},
		{
			name: "With no cycles",
			fields: fields{
				steps: []Step{A, B, C, D},
				dependencies: map[Step][]Step{
					A: {},
					B: {A},
					C: {D, B},
					D: {A},
				},
			},
			args:    args{},
			wantErr: false,
		},
		{
			name: "With cycle",
			fields: fields{
				steps: []Step{A, B, C, D},
				dependencies: map[Step][]Step{
					A: {B},
					B: {C},
					C: {D},
					D: {A},
				},
			},
			args:    args{},
			wantErr: true,
		},
		{
			name: "Without complex cycles",
			fields: fields{
				steps: []Step{A, B, C, D},
				dependencies: map[Step][]Step{
					A: {},
					B: {A, D, C},
					C: {D},
					D: {A},
				},
			},
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SimpleWorkflow{
				steps:        tt.fields.steps,
				status:       tt.fields.status,
				dependencies: tt.fields.dependencies,
			}
			if err := s.Init(tt.args.context); (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSimpleWorkflow_AddStep(t *testing.T) {
	type fields struct {
		steps        []Step
		status       Status
		dependencies map[Step][]Step
	}
	type args struct {
		step         Step
		dependencies []Step
	}
	step, _ := makeSimpleStep("sample-step")

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Add step without error",
			fields: fields{
				steps:        []Step{},
				dependencies: make(map[Step][]Step),
			},
			args:    args{step: step, dependencies: []Step{}},
			wantErr: false,
		},
		{
			name: "Add step that already exists",
			fields: fields{
				steps:        []Step{step},
				dependencies: make(map[Step][]Step),
			},
			args:    args{step: step, dependencies: []Step{}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SimpleWorkflow{
				steps:        tt.fields.steps,
				status:       tt.fields.status,
				dependencies: tt.fields.dependencies,
			}
			if err := s.AddStep(tt.args.step, tt.args.dependencies); (err != nil) != tt.wantErr {
				t.Errorf("AddStep() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSimpleWorkflow_debug(t *testing.T) {
	type fields struct {
		steps        []Step
		status       Status
		dependencies map[Step][]Step
	}
	tests := []struct {
		name           string
		fields         fields
		expectedStdout string
	}{
		{
			name: "with 2 steps workflow",
			fields: fields{
				steps: []Step{newSimpleStep("step-a", CREATED), newSimpleStep("step-b", PENDING)},
			},
			expectedStdout: "[DEBUG]: step-a : CREATED\n[DEBUG]: step-b : PENDING\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.SetFlags(0)
			var buf bytes.Buffer
			log.SetOutput(&buf)
			defer func() {
				log.SetOutput(os.Stderr)
			}()
			s := &SimpleWorkflow{
				steps:        tt.fields.steps,
				status:       tt.fields.status,
				dependencies: tt.fields.dependencies,
			}
			s.debug()
			if buf.String() != tt.expectedStdout {
				t.Errorf("debug() gotStdout = %v, wantStdout = %v", buf.String(), tt.expectedStdout)
			}

		})
	}
}
