package config

import (
	"os"
	"testing"
)

func TestGetEnv_Success(t *testing.T) {
	key := "TEST_KEY"
	expected := "test_value"
	os.Setenv(key, expected)
	defer os.Unsetenv("TEST_KEY")

	val := GetEnv(key)
	if val != expected {
		t.Errorf("Expected %s, got %s", expected, val)
	}
}
