package workflow

import (
	"bytes"
	"context"
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
	stepA := makeSimpleStep("step-a")
	stepSuccessB := makeSimpleStep("step-b")
	stepSuccessB.Status = SUCCESS
	stepSuccessC := makeSimpleStep("step-c")
	stepSuccessC.Status = SUCCESS
	stepPendingD := makeSimpleStep("step-c")
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

func TestSimpleWorkflow_Execute(t *testing.T) {
	type fields struct {
		steps        []Step
		status       Status
		dependencies map[Step][]Step
	}
	stepA := &stepWrapper{makeSimpleStep("step-a")}
	stepB := &stepWrapper{makeSimpleStep("step-b")}
	stepC := &stepWrapper{makeSimpleStep("step-c")}

	stepD := &stepWrapper{makeSimpleStep("step-d")}
	stepE := &stepWrapper{makeSimpleStep("step-e")}
	stepF := &stepWrapper{makeSimpleStep("step-f")}

	errorA := &stepWrapper{makeSimpleStep("error-a")}
	errorB := &stepWrapper{makeSimpleStep("error-b")}
	errorD := &stepWrapper{makeErrorStep("error-d")}

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
				steps: []Step{errorA, errorB, errorD},
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
			ch := make(chan map[string]Status)
			s := &SimpleWorkflow{
				steps:        tt.fields.steps,
				Dependencies: tt.fields.dependencies,
				StateCh:      make(chan map[string]Status, 1),
			}

			go func() {
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
			}()

			go func() {
				for range ch {
				}
			}()
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
	A := makeSimpleStep("step-a")
	B := makeSimpleStep("step-b")
	C := makeSimpleStep("step-c")
	D := makeSimpleStep("step-d")

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
				Dependencies: tt.fields.dependencies,
			}
			if err := s.Init(); (err != nil) != tt.wantErr {
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
	step := makeSimpleStep("sample-step")

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
				Dependencies: tt.fields.dependencies,
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
				Dependencies: tt.fields.dependencies,
			}
			s.debug()
			if buf.String() != tt.expectedStdout {
				t.Errorf("debug() gotStdout = %v, wantStdout = %v", buf.String(), tt.expectedStdout)
			}

		})
	}
}
