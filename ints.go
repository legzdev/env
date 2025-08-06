package env

import (
	"strconv"
)

type Signed interface {
	int | int8 | int16 | int32 | int64
}

func parseInt[T Signed](s string, bitSize int) (T, error) {
	i, err := strconv.ParseInt(s, 10, bitSize)
	if err != nil {
		return 0, err
	}

	return T(i), nil
}

func (env *Env) Int() int {
	return run(env, func() (int, error) {
		return strconv.Atoi(env.value)
	})
}

func (env *Env) Int8() int8 {
	return run(env, func() (int8, error) {
		return parseInt[int8](env.value, 8)
	})
}

func (env *Env) Int16() int16 {
	return run(env, func() (int16, error) {
		return parseInt[int16](env.value, 16)
	})
}

func (env *Env) Int32() int32 {
	return run(env, func() (int32, error) {
		return parseInt[int32](env.value, 32)
	})
}

func (env *Env) Int64() int64 {
	return run(env, func() (int64, error) {
		return parseInt[int64](env.value, 64)
	})
}

func (env *Env) IntParse(base int, bitSize int) int64 {
	return run(env, func() (int64, error) {
		return strconv.ParseInt(env.value, base, bitSize)
	})
}

func (env *Env) IntSlice() []int {
	return runS(env, func(s string) (int, error) {
		return strconv.Atoi(s)
	})
}

func (env *Env) Int8Slice() []int8 {
	return runS(env, func(s string) (int8, error) {
		return parseInt[int8](s, 8)
	})
}

func (env *Env) Int16Slice() []int16 {
	return runS(env, func(s string) (int16, error) {
		return parseInt[int16](s, 16)
	})
}

func (env *Env) Int32Slice() []int32 {
	return runS(env, func(s string) (int32, error) {
		return parseInt[int32](s, 32)
	})
}

func (env *Env) Int64Slice() []int64 {
	return runS(env, func(s string) (int64, error) {
		return parseInt[int64](s, 64)
	})
}

func (env *Env) Int64SliceParse(base int, bitSize int) []int64 {
	return runS(env, func(s string) (int64, error) {
		return strconv.ParseInt(s, base, bitSize)
	})
}
