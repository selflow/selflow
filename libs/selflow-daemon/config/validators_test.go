package config

import (
	"testing"
)

type TemplateTesterStruct struct {
	Value interface{} `validate:"template"`
}

type IdentifierTesterStruct struct {
	Value interface{} `validate:"identifier"`
}

func Test_validateGoTemplate(t *testing.T) {
	v := InitValidation()

	tests := []struct {
		name   string
		tester TemplateTesterStruct
		want   bool
	}{
		{
			name:   "String with no template",
			tester: TemplateTesterStruct{"toto"},
			want:   true,
		},
		{
			name:   "Valid template",
			tester: TemplateTesterStruct{"{{ .toto }}"},
			want:   true,
		},
		{
			name:   "Invalid template",
			tester: TemplateTesterStruct{"{{ toto }}"},
			want:   false,
		},
		{
			name:   "Not a string",
			tester: TemplateTesterStruct{884},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := v.Struct(tt.tester)

			if got := err == nil; got != tt.want {
				t.Errorf("validateGoTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidIdentifier(t *testing.T) {
	type args struct {
		identifier string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid identifier",
			args: args{"toto"},
			want: true,
		},
		{
			name: "Invalid identifier",
			args: args{"__888totoXx_"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidIdentifier(tt.args.identifier); got != tt.want {
				t.Errorf("isValidIdentifier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateIdentifier(t *testing.T) {
	v := InitValidation()

	tests := []struct {
		name   string
		tester IdentifierTesterStruct
		want   bool
	}{
		{
			name:   "Valid identifier",
			tester: IdentifierTesterStruct{"toto"},
			want:   true,
		},
		{
			name:   "Invalid identifier",
			tester: IdentifierTesterStruct{"__888totoXx_"},
			want:   false,
		},
		{
			name:   "Not a string",
			tester: IdentifierTesterStruct{884},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := v.Struct(tt.tester)
			if got := err == nil; got != tt.want {
				t.Errorf("validateIdentifier() = %v, want %v", got, tt.want)
			}
		})
	}
}
