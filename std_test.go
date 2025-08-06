package env_test

import (
	"errors"
	"os"
	"testing"

	"github.com/legzdev/env"
)

func TestBool(t *testing.T) {
	boolTests := map[string]bool{
		"":     false,
		"0":    false,
		"true": true,
		"1":    true,
	}

	envName := GetEnvName()

	for value, expexted := range boolTests {
		os.Setenv(envName, value)

		result := env.New(envName).
			WithErrorHandler(func(err error) { t.Error(err) }).
			Bool()

		if result != expexted {
			envError(t, value, expexted, result)
		}
	}

	os.Setenv(envName, "-")

	env.New(envName).WithErrorHandler(func(err error) {
		if !errors.Is(err, env.ErrParsingFailed) {
			t.Error("accepted invalid value")
		}
	})
}
