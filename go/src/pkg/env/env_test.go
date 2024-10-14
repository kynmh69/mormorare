package env

import (
	"os"
	"testing"
)

func TestFindEnvVariableExists(t *testing.T) {
	os.Setenv("EXISTING_KEY", "value")
	defer os.Unsetenv("EXISTING_KEY")

	if got := FindEnv("EXISTING_KEY"); got != "value" {
		t.Errorf("FindEnv() = %v, want %v", got, "value")
	}
}

func TestFindEnvVariableDoesNotExist(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("FindEnv() did not panic")
		}
	}()

	FindEnv("NON_EXISTING_KEY")
}
