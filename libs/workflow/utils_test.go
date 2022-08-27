package workflow

import (
	"errors"
	"reflect"
	"testing"
)

func Test_mergeStringStringStringMaps(t *testing.T) {
	type args struct {
		destination map[string]map[string]string
		maps        []map[string]map[string]string
	}
	tests := []struct {
		name string
		args args
		want map[string]map[string]string
	}{
		{
			name: "single map",
			args: args{
				destination: map[string]map[string]string{
					"aaa": {
						"bbb": "ccc",
					},
				},
				maps: []map[string]map[string]string{},
			},
			want: map[string]map[string]string{
				"aaa": {
					"bbb": "ccc",
				},
			},
		},
		{
			name: "two maps",
			args: args{
				destination: map[string]map[string]string{
					"aaa": {
						"bbb": "ccc",
					},
				},
				maps: []map[string]map[string]string{
					{
						"ddd": {
							"eee": "fff",
						},
					},
				},
			},
			want: map[string]map[string]string{
				"aaa": {
					"bbb": "ccc",
				},
				"ddd": {
					"eee": "fff",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeStringStringStringMaps(tt.args.destination, tt.args.maps...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeStringStringStringMaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appendErrorList(t *testing.T) {
	type args struct {
		errorLst []error
		err      error
	}
	tests := []struct {
		name string
		args args
		want []error
	}{
		{
			name: "should return unchanged list if err is nil",
			args: args{
				errorLst: []error{errors.New("errA"), errors.New("errB")},
				err:      nil,
			},
			want: []error{errors.New("errA"), errors.New("errB")},
		},
		{
			name: "should add err to list",
			args: args{
				errorLst: []error{errors.New("errA"), errors.New("errB")},
				err:      errors.New("errC"),
			},
			want: []error{errors.New("errA"), errors.New("errB"), errors.New("errC")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := appendErrorList(tt.args.errorLst, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendErrorList() = %v, want %v", got, tt.want)
			}
		})
	}
}
