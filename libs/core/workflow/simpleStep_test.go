package workflow

import (
	"context"
	"reflect"
	"testing"
)

func TestSimpleStep_ValidStep(t *testing.T) {
	step := &SimpleStep{}
	var i interface{} = step

	_, ok := i.(Step)

	if !ok {
		t.Errorf("Simple step is an invalid step")
	}
}

func TestSimpleStep_Execute(t *testing.T) {
	type fields struct {
		id     string
		status Status
	}
	type args struct {
		context context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name:    "execution",
			fields:  fields{status: PENDING},
			args:    args{},
			want:    map[string]string{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SimpleStep{
				Id:     tt.fields.id,
				Status: tt.fields.status,
			}
			got, err := s.Execute(tt.args.context)
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleStep_GetId(t *testing.T) {
	type fields struct {
		id     string
		status Status
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Step with id",
			fields: fields{"my-step", PENDING},
			want:   "my-step",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SimpleStep{
				Id:     tt.fields.id,
				Status: tt.fields.status,
			}
			if got := s.GetId(); got != tt.want {
				t.Errorf("GetId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleStep_GetStatus(t *testing.T) {
	type fields struct {
		id     string
		status Status
	}
	tests := []struct {
		name   string
		fields fields
		want   Status
	}{
		{
			name:   "with PENDING Status",
			fields: fields{"my-step", PENDING},
			want:   PENDING,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SimpleStep{
				Id:     tt.fields.id,
				Status: tt.fields.status,
			}
			if got := s.GetStatus(); got != tt.want {
				t.Errorf("GetStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeSimpleStep(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want *SimpleStep
	}{
		{
			"create step with id",
			args{"my-step"},
			&SimpleStep{
				Id:     "my-step",
				Status: PENDING,
			},
		},
		{
			"create step with empty id",
			args{""},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeSimpleStep(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeSimpleStep() got = %v, want %v", got, tt.want)
			}
		})
	}
}
