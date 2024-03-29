package workflow

import (
	"context"
	"reflect"
	"sync"
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

	stepAA := &stepWrapper{makeSimpleStep("step-a-a")}
	stepAB := &stepWrapper{makeSimpleStep("step-a-b")}
	stepAC := &stepWrapper{makeSimpleStep("step-a-c")}

	stepBA := &stepWrapper{makeSimpleStep("step-b-a")}
	stepBB := &stepWrapper{makeSimpleStep("step-b-b")}
	errorBC := &stepWrapper{makeErrorStep("error-b-d")}

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
				steps:        []Step{stepAA, stepAB, stepAC},
				status:       PENDING,
				dependencies: map[Step][]Step{},
			},
			args: args{
				ctx: context.TODO(),
			},
			want: map[string]map[string]string{
				stepAA.GetId(): {},
				stepAB.GetId(): {},
				stepAC.GetId(): {},
			},
			wantErr: false,
			wantStatus: map[Step]Status{
				stepAB: SUCCESS,
				stepAC: SUCCESS,
				stepAA: SUCCESS,
			},
		},
		{
			name: "with cancel",
			fields: fields{
				steps: []Step{stepBA, stepBB, errorBC},
				dependencies: map[Step][]Step{
					errorBC: {},
					stepBA:  {errorBC},
					stepBB:  {errorBC},
				},
			},
			args: args{
				ctx: context.TODO(),
			},
			want: map[string]map[string]string{
				stepBA.GetId():  {},
				stepBB.GetId():  {},
				errorBC.GetId(): {},
			},
			wantErr: false,
			wantStatus: map[Step]Status{
				stepBA:  CANCELLED,
				stepBB:  CANCELLED,
				errorBC: ERROR,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := make(chan map[string]Status)
			s := &SimpleWorkflow{
				steps:        tt.fields.steps,
				Dependencies: tt.fields.dependencies,
				StateCh:      ch,
			}

			wg := sync.WaitGroup{}
			wg.Add(2)

			go func() {
				defer wg.Done()
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
				defer wg.Done()
				for range ch {
				}

			}()

			wg.Wait()
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
