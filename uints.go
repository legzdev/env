package env

import (
	"strconv"
)

type Unsigned interface {
	uint | uint8 | uint16 | uint32 | uint64
}

func parseUint[T Unsigned](s string, bitSize int) (T, error) {
	i, err := strconv.ParseUint(s, 10, bitSize)
	if err != nil {
		return 0, err
	}

	return T(i), nil
}

func (env *Env) Uint() uint {
	return run(env, func() (uint, error) {
		return parseUint[uint](env.value, 64)
	})
}

func (env *Env) Uint8() uint8 {
	return run(env, func() (uint8, error) {
		return parseUint[uint8](env.value, 8)
	})
}

func (env *Env) Uint16() uint16 {
	return run(env, func() (uint16, error) {
		return parseUint[uint16](env.value, 16)
	})
}

func (env *Env) Uint32() uint32 {
	return run(env, func() (uint32, error) {
		return parseUint[uint32](env.value, 32)
	})
}

func (env *Env) Uint64() uint64 {
	return run(env, func() (uint64, error) {
		return parseUint[uint64](env.value, 64)
	})
}

func (env *Env) UintParse(base int, bitSize int) uint64 {
	return run(env, func() (uint64, error) {
		return strconv.ParseUint(env.value, base, bitSize)
	})
}

func (env *Env) UintSlice() []uint {
	return runS(env, func(s string) (uint, error) {
		return parseUint[uint](s, 64)
	})
}

func (env *Env) Uint8Slice() []uint8 {
	return runS(env, func(s string) (uint8, error) {
		return parseUint[uint8](s, 8)
	})
}

func (env *Env) Uint16Slice() []uint16 {
	return runS(env, func(s string) (uint16, error) {
		return parseUint[uint16](s, 16)
	})
}

func (env *Env) Uint32Slice() []uint32 {
	return runS(env, func(s string) (uint32, error) {
		return parseUint[uint32](s, 32)
	})
}

func (env *Env) Uint64Slice() []uint64 {
	return runS(env, func(s string) (uint64, error) {
		return parseUint[uint64](s, 64)
	})
}

func (env *Env) Uint64SliceParse(base int, bitSize int) []uint64 {
	return runS(env, func(s string) (uint64, error) {
		return strconv.ParseUint(s, base, bitSize)
	})
}
