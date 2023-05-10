package workflow

import (
	"context"
	"reflect"
	"testing"
)

type cancelStep struct {
	*SimpleStep
}

func (s *cancelStep) Execute(_ context.Context) (map[string]string, error) {
	s.SetStatus(CANCELLED)
	return map[string]string{}, nil
}

func Test_stepWrapper_Execute(t *testing.T) {
	stepA := &errorStep{makeSimpleStep("step-a")}
	stepB := makeSimpleStep("step-a")
	stepC := &cancelStep{makeSimpleStep("step-a")}

	type fields struct {
		Step Step
	}
	tests := []struct {
		name       string
		fields     fields
		want       map[string]string
		wantErr    bool
		wantStatus Status
	}{
		{
			name:       "step failed",
			fields:     fields{Step: stepA},
			wantErr:    true,
			want:       map[string]string{},
			wantStatus: ERROR,
		},
		{
			name:       "step succeeded",
			fields:     fields{Step: stepB},
			want:       map[string]string{},
			wantStatus: SUCCESS,
		},
		{
			name:       "step already finished",
			fields:     fields{Step: stepC},
			want:       map[string]string{},
			wantStatus: CANCELLED,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stepWrapper{
				Step: tt.fields.Step,
			}
			got, err := s.Execute(context.TODO())
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() got = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(s.GetStatus(), tt.wantStatus) {
				t.Errorf("Execute() status got = %v, want %v", s.GetStatus(), tt.wantStatus)
			}
		})
	}
}
