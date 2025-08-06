package env

import (
	"errors"
	"net/url"
	"slices"
	"strings"
	"time"
)

func (env *Env) String() string {
	return env.value
}

func (env *Env) StringSlice() []string {
	return run(env, func() ([]string, error) {
		return strings.Split(env.value, env.sep), nil
	})
}

func (env *Env) Bool() bool {
	return run(env, func() (bool, error) {
		value := strings.ToLower(env.value)

		if slices.Contains(TrutyValues, value) {
			return true, nil
		}

		if slices.Contains(FalsyValues, value) {
			return false, nil
		}

		err := &parsingError{
			Name: env.name,
			err:  errors.New("unrecognized value '" + env.value + "'"),
		}

		return false, err
	})
}

func (env *Env) Duration() time.Duration {
	return run(env, func() (time.Duration, error) {
		return time.ParseDuration(env.value)
	})
}

func (env *Env) Time(layout string) time.Time {
	return run(env, func() (time.Time, error) {
		return time.Parse(layout, env.value)
	})
}

func (env *Env) URL() *url.URL {
	return run(env, func() (*url.URL, error) {
		return url.Parse(env.value)
	})
}
