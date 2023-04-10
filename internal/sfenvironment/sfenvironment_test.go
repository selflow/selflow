package sfenvironment

import (
	"os"
	"testing"
)

func TestGetDaemonBaseDirectory(t *testing.T) {
	tests := []struct {
		name string
		env  string
		want string
	}{
		{
			name: "default",
			want: DaemonBaseDirectoryEnvDefaultValue,
		},
		{
			name: "specific env",
			env:  "toto",
			want: "toto",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.env != "" {
				_ = os.Setenv(DaemonBaseDirectoryEnvKey, tt.env)
				defer func() {
					_ = os.Unsetenv(DaemonBaseDirectoryEnvKey)
				}()
			}
			if got := GetDaemonBaseDirectory(); got != tt.want {
				t.Errorf("GetDaemonBaseDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDaemonDebugPort(t *testing.T) {
	tests := []struct {
		name string
		env  string
		want string
	}{
		{
			name: "default",
			want: DaemonDebugPortEnvDefaultValue,
		},
		{
			name: "specific env",
			env:  "toto",
			want: "toto",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.env != "" {
				_ = os.Setenv(DaemonDebugPortEnvKey, tt.env)
				defer func() {
					_ = os.Unsetenv(DaemonDebugPortEnvKey)
				}()
			}

			if got := GetDaemonDebugPort(); got != tt.want {
				t.Errorf("GetDaemonDebugPort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDaemonHostBaseDirectory(t *testing.T) {
	tests := []struct {
		name        string
		baseEnv     string
		hostBaseEnv string
		want        string
	}{
		{
			name: "default",
			want: DaemonBaseDirectoryEnvDefaultValue,
		},
		{
			name:    "specific base env",
			baseEnv: "toto",
			want:    "toto",
		},
		{
			name:    "specific host base env",
			baseEnv: "toto",
			want:    "toto",
		},
		{
			name:        "specific host base and base env",
			baseEnv:     "toto",
			hostBaseEnv: "tata",
			want:        "tata",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.baseEnv != "" {
				_ = os.Setenv(DaemonBaseDirectoryEnvKey, tt.baseEnv)
				defer func() {
					_ = os.Unsetenv(DaemonBaseDirectoryEnvKey)
				}()
			}

			if tt.hostBaseEnv != "" {
				_ = os.Setenv(DaemonHostBaseDirectoryEnvKey, tt.hostBaseEnv)
				defer func() {
					_ = os.Unsetenv(DaemonHostBaseDirectoryEnvKey)
				}()
			}

			if got := GetDaemonHostBaseDirectory(); got != tt.want {
				t.Errorf("GetDaemonHostBaseDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDaemonImage(t *testing.T) {
	tests := []struct {
		name string
		env  string
		want string
	}{
		{
			name: "default",
			want: DaemonImageEnvDefaultValue,
		},
		{
			name: "specific env",
			env:  "toto",
			want: "toto",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.env != "" {
				_ = os.Setenv(DaemonImageEnvKey, tt.env)
				defer func() {
					_ = os.Unsetenv(DaemonImageEnvKey)
				}()
			}
			if got := GetDaemonImage(); got != tt.want {
				t.Errorf("GetDaemonImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDaemonName(t *testing.T) {
	tests := []struct {
		name string
		env  string
		want string
	}{
		{
			name: "default",
			want: DaemonNameEnvDefaultValue,
		},
		{
			name: "specific env",
			env:  "toto",
			want: "toto",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.env != "" {
				_ = os.Setenv(DaemonNameEnvKey, tt.env)
				defer func() {
					_ = os.Unsetenv(DaemonNameEnvKey)
				}()
			}
			if got := GetDaemonName(); got != tt.want {
				t.Errorf("GetDaemonName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDaemonNetwork(t *testing.T) {
	tests := []struct {
		name       string
		nameEnv    string
		networkEnv string
		want       string
	}{
		{
			name: "default",
			want: DaemonNameEnvDefaultValue,
		},
		{
			name:    "specific name env",
			nameEnv: "toto",
			want:    "toto",
		},
		{
			name:       "specific network env",
			networkEnv: "toto",
			want:       "toto",
		},
		{
			name:       "specific network and name env",
			nameEnv:    "tata",
			networkEnv: "toto",
			want:       "toto",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.nameEnv != "" {
				_ = os.Setenv(DaemonNameEnvKey, tt.nameEnv)
				defer func() {
					_ = os.Unsetenv(DaemonNameEnvKey)
				}()
			}
			if tt.networkEnv != "" {
				_ = os.Setenv(DaemonNetworkEnvKey, tt.networkEnv)
				defer func() {
					_ = os.Unsetenv(DaemonNetworkEnvKey)
				}()
			}
			if got := GetDaemonNetwork(); got != tt.want {
				t.Errorf("GetDaemonNetwork() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDaemonPort(t *testing.T) {
	tests := []struct {
		name string
		env  string
		want string
	}{
		{
			name: "default",
			want: DaemonPortEnvDefaultValue,
		},
		{
			name: "specific env",
			env:  "toto",
			want: "toto",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.env != "" {
				_ = os.Setenv(DaemonPortEnvKey, tt.env)
				defer func() {
					_ = os.Unsetenv(DaemonPortEnvKey)
				}()
			}
			if got := GetDaemonPort(); got != tt.want {
				t.Errorf("GetDaemonPort() = %v, want %v", got, tt.want)
			}
		})
	}
}
