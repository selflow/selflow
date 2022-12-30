package workflow

import (
	"reflect"
	"testing"
)

func TestErrorStep_Cancel(t *testing.T) {
	type fields struct {
		id     string
		status Status
	}
	tests := []struct {
		name       string
		fields     fields
		wantFields fields
		wantErr    bool
	}{
		{
			name: "simple-case",
			fields: fields{
				id:     "step-a",
				status: CREATED,
			},
			wantFields: fields{
				"step-a",
				CANCELLED,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ErrorStep{
				id:     tt.fields.id,
				status: tt.fields.status,
			}
			if err := s.Cancel(); (err != nil) != tt.wantErr {
				t.Errorf("Cancel() error = %v, wantErr %v", err, tt.wantErr)
			}
			if s.id != tt.wantFields.id || s.status != tt.wantFields.status {
				t.Errorf("Cancel() got = %v, want %v", *s, tt.wantFields)
			}
		})
	}
}

func Test_makeErrorStep(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *ErrorStep
		wantErr bool
	}{
		{
			name:    "should create ErrorStep",
			args:    args{"step-a"},
			want:    &ErrorStep{"step-a", CREATED},
			wantErr: false,
		},
		{
			name:    "should return an error if empty id",
			args:    args{""},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := makeErrorStep(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("makeErrorStep() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeErrorStep() got = %v, want %v", got, tt.want)
			}
		})
	}
}
