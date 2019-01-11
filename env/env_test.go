package env

import (
	"os"
	"testing"
)

func TestEnvGet(t *testing.T) {
	key := "FOO"
	t.Fatal(os.Getenv(key))
}
