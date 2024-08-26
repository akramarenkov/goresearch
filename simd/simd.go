package simd

// https://gcc.gnu.org/onlinedocs/gcc/x86-Built-in-Functions.html

// The absence of options such as -mavx -mavx2 that enable built-in functions that
// reference the necessary processor instructions will result in the error
// "inlining failed in call to "always_inline"" when compiling

/*
#cgo CFLAGS: -mavx -mavx2
#include "simd.h"
*/
import "C"

func SlowMultiplication(first uint64, second uint64) uint64 {
	product := uint64(0)

	for range second {
		product += first
	}

	return product
}

// Only for numbers without fractional parts.
func SlowMultiplicationF64(first float64, second float64) float64 {
	product := float64(0)

	for range uint64(second) {
		product += first
	}

	return product
}

func SlowMultiplicationSIMD(first uint64, second uint64) uint64 {
	return uint64(C.SlowMultiplication(C.ulonglong(first), C.ulonglong(second)))
}

// Only for numbers without fractional parts.
func SlowMultiplicationF64SIMD(first float64, second float64) float64 {
	return float64(C.SlowMultiplicationF64(C.double(first), C.double(second)))
}
