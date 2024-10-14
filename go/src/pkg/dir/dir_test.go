package dir

import (
	"os"
	"testing"
)

func TestGetProjectRootWithoutGoMod(t *testing.T) {
	tempDir := t.TempDir()
	os.Chdir(tempDir)
	defer os.Chdir("..")

	tests := []struct {
		name string
		want string
	}{
		{"GoModDoesNotExist", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetProjectRoot(); got != tt.want {
				t.Errorf("GetProjectRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetProjectRoot(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"GoModExists", "/Users/Hiroki/Applications/GoProjects/mormorare/go/src"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetProjectRoot(); got != tt.want {
				t.Errorf("GetProjectRoot() = %v, want %v", got, tt.want)
			}
		})

	}
}
