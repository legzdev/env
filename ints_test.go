package env_test

import (
	"os"
	"slices"
	"testing"

	"github.com/legzdev/env"
)

func TestInt(t *testing.T) {
	intTests := map[string]int{
		"":    0,
		"0":   0,
		"-87": -87,
		"34":  34,
	}

	envName := GetEnvName()

	for value, expected := range intTests {
		os.Setenv(envName, value)

		result := env.New(envName).
			WithErrorHandler(func(err error) { t.Error(err) }).
			Int()

		if result != expected {
			envError(t, value, expected, result)
		}
	}
}

func TestIntSlice(t *testing.T) {
	intSliceTests := map[string][]int{
		"":          {},
		"0":         {0},
		"9,1,293":   {9, 1, 293},
		"-87,34,25": {-87, 34, 25},
	}

	envName := GetEnvName()

	for value, expected := range intSliceTests {
		os.Setenv(envName, value)

		result := env.New(envName).
			WithErrorHandler(func(err error) { t.Error(err) }).
			IntSlice()

		if !slices.Equal(result, expected) {
			envError(t, value, expected, result)
		}
	}
}
