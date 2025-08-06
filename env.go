package env

import (
	"os"
	"strings"
)

const DefaultSeparator = ","

var (
	TrutyValues = []string{"1", "true", "yes"}
	FalsyValues = []string{"0", "false", "not"}
)

type (
	Parser[T any] = func(s string) (T, error)
	ErrorHandler  = func(err error)
)

type Env struct {
	name         string
	value        string
	sep          string
	required     bool
	errorHandler ErrorHandler
}

func New(name string) *Env {
	return &Env{
		name:  name,
		value: os.Getenv(name),
		sep:   DefaultSeparator,
	}
}

func (env *Env) WithDefault(s string) *Env {
	if env.value == "" {
		env.value = s
	}
	return env
}

func (env *Env) WithRequired(v bool) *Env {
	env.required = v
	return env
}

func (env *Env) Required() *Env {
	env.required = true
	return env
}

func run[T any](env *Env, fn func() (T, error)) T {
	if env.value == "" {
		var v T

		if env.required {
			env.error(&missingError{Name: env.name})
		}

		return v
	}

	v, err := fn()
	if err != nil {
		env.error(err)
	}

	return v
}

func runS[T any](env *Env, parser Parser[T]) []T {
	return run(env, func() ([]T, error) {
		var result []T

		for value := range strings.SplitSeq(env.value, env.sep) {
			ret, err := parser(value)
			if err != nil {
				return nil, err
			}

			result = append(result, ret)
		}

		return result, nil
	})
}

func (env *Env) WithErrorHandler(handler ErrorHandler) *Env {
	env.errorHandler = handler
	return env
}

func (env *Env) error(err error) {
	if env.errorHandler == nil {
		return
	}

	env.errorHandler(&parsingError{
		Name: env.name,
		err:  err,
	})
}
