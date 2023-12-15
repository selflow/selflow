package workflow

import (
	"reflect"
	"testing"
)

func Test_makeErrorStep(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want *errorStep
	}{
		{
			name: "should create errorStep",
			args: args{"step-a"},
			want: &errorStep{&SimpleStep{
				Id:     "step-a",
				Status: PENDING,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeErrorStep(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeErrorStep() got = %v, want %v", got, tt.want)
			}
		})
	}
}
