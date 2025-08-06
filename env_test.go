package env_test

import (
	"crypto/rand"
	"encoding/hex"
	"testing"
)

func GetEnvName() string {
	random := make([]byte, 6)
	rand.Read(random)
	return hex.EncodeToString(random)
}

func envError(t *testing.T, value string, expected any, result any) {
	t.Errorf("invalid result for %T(%s): expected '%v' got '%v'",
		expected, value, expected, result,
	)
}
