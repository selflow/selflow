package workflow

import (
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
