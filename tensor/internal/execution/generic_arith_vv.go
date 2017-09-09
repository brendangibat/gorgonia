package execution

/*
GENERATED FILE. DO NOT EDIT
*/

import (
	"math"
	"math/cmplx"
	_ "unsafe"

	"github.com/chewxy/math32"
	_ "github.com/chewxy/vecf32"
	_ "github.com/chewxy/vecf64"
)

func VecAddI(a []int, b []int) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] + b[i]
	}
}

func VecAddI8(a []int8, b []int8) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] + b[i]
	}
}

func VecAddI16(a []int16, b []int16) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] + b[i]
	}
}

func VecAddI32(a []int32, b []int32) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] + b[i]
	}
}

func VecAddI64(a []int64, b []int64) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] + b[i]
	}
}

func VecAddU(a []uint, b []uint) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] + b[i]
	}
}

func VecAddU8(a []uint8, b []uint8) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] + b[i]
	}
}

func VecAddU16(a []uint16, b []uint16) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] + b[i]
	}
}

func VecAddU32(a []uint32, b []uint32) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] + b[i]
	}
}

func VecAddU64(a []uint64, b []uint64) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] + b[i]
	}
}

//go:linkname VecAddF32 github.com/chewxy/vecf32.Add
func VecAddF32(a []float32, b []float32)

//go:linkname VecAddF64 github.com/chewxy/vecf64.Add
func VecAddF64(a []float64, b []float64)

func VecAddC64(a []complex64, b []complex64) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] + b[i]
	}
}

func VecAddC128(a []complex128, b []complex128) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] + b[i]
	}
}

func VecAddStr(a []string, b []string) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] + b[i]
	}
}

func VecSubI(a []int, b []int) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] - b[i]
	}
}

func VecSubI8(a []int8, b []int8) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] - b[i]
	}
}

func VecSubI16(a []int16, b []int16) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] - b[i]
	}
}

func VecSubI32(a []int32, b []int32) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] - b[i]
	}
}

func VecSubI64(a []int64, b []int64) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] - b[i]
	}
}

func VecSubU(a []uint, b []uint) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] - b[i]
	}
}

func VecSubU8(a []uint8, b []uint8) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] - b[i]
	}
}

func VecSubU16(a []uint16, b []uint16) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] - b[i]
	}
}

func VecSubU32(a []uint32, b []uint32) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] - b[i]
	}
}

func VecSubU64(a []uint64, b []uint64) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] - b[i]
	}
}

//go:linkname VecSubF32 github.com/chewxy/vecf32.Sub
func VecSubF32(a []float32, b []float32)

//go:linkname VecSubF64 github.com/chewxy/vecf64.Sub
func VecSubF64(a []float64, b []float64)

func VecSubC64(a []complex64, b []complex64) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] - b[i]
	}
}

func VecSubC128(a []complex128, b []complex128) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] - b[i]
	}
}

func VecMulI(a []int, b []int) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] * b[i]
	}
}

func VecMulI8(a []int8, b []int8) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] * b[i]
	}
}

func VecMulI16(a []int16, b []int16) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] * b[i]
	}
}

func VecMulI32(a []int32, b []int32) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] * b[i]
	}
}

func VecMulI64(a []int64, b []int64) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] * b[i]
	}
}

func VecMulU(a []uint, b []uint) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] * b[i]
	}
}

func VecMulU8(a []uint8, b []uint8) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] * b[i]
	}
}

func VecMulU16(a []uint16, b []uint16) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] * b[i]
	}
}

func VecMulU32(a []uint32, b []uint32) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] * b[i]
	}
}

func VecMulU64(a []uint64, b []uint64) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] * b[i]
	}
}

//go:linkname VecMulF32 github.com/chewxy/vecf32.Mul
func VecMulF32(a []float32, b []float32)

//go:linkname VecMulF64 github.com/chewxy/vecf64.Mul
func VecMulF64(a []float64, b []float64)

func VecMulC64(a []complex64, b []complex64) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] * b[i]
	}
}

func VecMulC128(a []complex128, b []complex128) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] * b[i]
	}
}

func VecDivI(a []int, b []int) (err error) {
	a = a[:]
	b = b[:len(a)]
	var errs errorIndices
	for i := range a {
		if b[i] == 0 {
			errs = append(errs, i)
			a[i] = 0
			continue
		}
		a[i] = a[i] / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func VecDivI8(a []int8, b []int8) (err error) {
	a = a[:]
	b = b[:len(a)]
	var errs errorIndices
	for i := range a {
		if b[i] == 0 {
			errs = append(errs, i)
			a[i] = 0
			continue
		}
		a[i] = a[i] / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func VecDivI16(a []int16, b []int16) (err error) {
	a = a[:]
	b = b[:len(a)]
	var errs errorIndices
	for i := range a {
		if b[i] == 0 {
			errs = append(errs, i)
			a[i] = 0
			continue
		}
		a[i] = a[i] / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func VecDivI32(a []int32, b []int32) (err error) {
	a = a[:]
	b = b[:len(a)]
	var errs errorIndices
	for i := range a {
		if b[i] == 0 {
			errs = append(errs, i)
			a[i] = 0
			continue
		}
		a[i] = a[i] / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func VecDivI64(a []int64, b []int64) (err error) {
	a = a[:]
	b = b[:len(a)]
	var errs errorIndices
	for i := range a {
		if b[i] == 0 {
			errs = append(errs, i)
			a[i] = 0
			continue
		}
		a[i] = a[i] / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func VecDivU(a []uint, b []uint) (err error) {
	a = a[:]
	b = b[:len(a)]
	var errs errorIndices
	for i := range a {
		if b[i] == 0 {
			errs = append(errs, i)
			a[i] = 0
			continue
		}
		a[i] = a[i] / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func VecDivU8(a []uint8, b []uint8) (err error) {
	a = a[:]
	b = b[:len(a)]
	var errs errorIndices
	for i := range a {
		if b[i] == 0 {
			errs = append(errs, i)
			a[i] = 0
			continue
		}
		a[i] = a[i] / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func VecDivU16(a []uint16, b []uint16) (err error) {
	a = a[:]
	b = b[:len(a)]
	var errs errorIndices
	for i := range a {
		if b[i] == 0 {
			errs = append(errs, i)
			a[i] = 0
			continue
		}
		a[i] = a[i] / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func VecDivU32(a []uint32, b []uint32) (err error) {
	a = a[:]
	b = b[:len(a)]
	var errs errorIndices
	for i := range a {
		if b[i] == 0 {
			errs = append(errs, i)
			a[i] = 0
			continue
		}
		a[i] = a[i] / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func VecDivU64(a []uint64, b []uint64) (err error) {
	a = a[:]
	b = b[:len(a)]
	var errs errorIndices
	for i := range a {
		if b[i] == 0 {
			errs = append(errs, i)
			a[i] = 0
			continue
		}
		a[i] = a[i] / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

//go:linkname VecDivF32 github.com/chewxy/vecf32.Div
func VecDivF32(a []float32, b []float32)

//go:linkname VecDivF64 github.com/chewxy/vecf64.Div
func VecDivF64(a []float64, b []float64)

func VecDivC64(a []complex64, b []complex64) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] / b[i]
	}
}

func VecDivC128(a []complex128, b []complex128) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] / b[i]
	}
}

//go:linkname VecPowF32 github.com/chewxy/vecf32.Pow
func VecPowF32(a []float32, b []float32)

//go:linkname VecPowF64 github.com/chewxy/vecf64.Pow
func VecPowF64(a []float64, b []float64)

func VecPowC64(a []complex64, b []complex64) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = complex64(cmplx.Pow(complex128(a[i]), complex128(b[i])))
	}
}

func VecPowC128(a []complex128, b []complex128) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = cmplx.Pow(a[i], b[i])
	}
}

func VecModI(a []int, b []int) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] % b[i]
	}
}

func VecModI8(a []int8, b []int8) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] % b[i]
	}
}

func VecModI16(a []int16, b []int16) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] % b[i]
	}
}

func VecModI32(a []int32, b []int32) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] % b[i]
	}
}

func VecModI64(a []int64, b []int64) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] % b[i]
	}
}

func VecModU(a []uint, b []uint) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] % b[i]
	}
}

func VecModU8(a []uint8, b []uint8) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] % b[i]
	}
}

func VecModU16(a []uint16, b []uint16) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] % b[i]
	}
}

func VecModU32(a []uint32, b []uint32) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] % b[i]
	}
}

func VecModU64(a []uint64, b []uint64) {
	a = a[:]
	b = b[:len(a)]
	for i := range a {
		a[i] = a[i] % b[i]
	}
}

//go:linkname VecModF32 github.com/chewxy/vecf32.Mod
func VecModF32(a []float32, b []float32)

//go:linkname VecModF64 github.com/chewxy/vecf64.Mod
func VecModF64(a []float64, b []float64)

func AddIncrI(a []int, b []int, incr []int) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] + b[i]
	}
}

func AddIncrI8(a []int8, b []int8, incr []int8) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] + b[i]
	}
}

func AddIncrI16(a []int16, b []int16, incr []int16) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] + b[i]
	}
}

func AddIncrI32(a []int32, b []int32, incr []int32) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] + b[i]
	}
}

func AddIncrI64(a []int64, b []int64, incr []int64) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] + b[i]
	}
}

func AddIncrU(a []uint, b []uint, incr []uint) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] + b[i]
	}
}

func AddIncrU8(a []uint8, b []uint8, incr []uint8) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] + b[i]
	}
}

func AddIncrU16(a []uint16, b []uint16, incr []uint16) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] + b[i]
	}
}

func AddIncrU32(a []uint32, b []uint32, incr []uint32) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] + b[i]
	}
}

func AddIncrU64(a []uint64, b []uint64, incr []uint64) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] + b[i]
	}
}

//go:linkname AddIncrF32 github.com/chewxy/vecf32.IncrAdd
func AddIncrF32(a []float32, b []float32, incr []float32)

//go:linkname AddIncrF64 github.com/chewxy/vecf64.IncrAdd
func AddIncrF64(a []float64, b []float64, incr []float64)

func AddIncrC64(a []complex64, b []complex64, incr []complex64) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] + b[i]
	}
}

func AddIncrC128(a []complex128, b []complex128, incr []complex128) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] + b[i]
	}
}

func AddIncrStr(a []string, b []string, incr []string) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] + b[i]
	}
}

func SubIncrI(a []int, b []int, incr []int) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] - b[i]
	}
}

func SubIncrI8(a []int8, b []int8, incr []int8) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] - b[i]
	}
}

func SubIncrI16(a []int16, b []int16, incr []int16) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] - b[i]
	}
}

func SubIncrI32(a []int32, b []int32, incr []int32) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] - b[i]
	}
}

func SubIncrI64(a []int64, b []int64, incr []int64) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] - b[i]
	}
}

func SubIncrU(a []uint, b []uint, incr []uint) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] - b[i]
	}
}

func SubIncrU8(a []uint8, b []uint8, incr []uint8) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] - b[i]
	}
}

func SubIncrU16(a []uint16, b []uint16, incr []uint16) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] - b[i]
	}
}

func SubIncrU32(a []uint32, b []uint32, incr []uint32) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] - b[i]
	}
}

func SubIncrU64(a []uint64, b []uint64, incr []uint64) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] - b[i]
	}
}

//go:linkname SubIncrF32 github.com/chewxy/vecf32.IncrSub
func SubIncrF32(a []float32, b []float32, incr []float32)

//go:linkname SubIncrF64 github.com/chewxy/vecf64.IncrSub
func SubIncrF64(a []float64, b []float64, incr []float64)

func SubIncrC64(a []complex64, b []complex64, incr []complex64) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] - b[i]
	}
}

func SubIncrC128(a []complex128, b []complex128, incr []complex128) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] - b[i]
	}
}

func MulIncrI(a []int, b []int, incr []int) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] * b[i]
	}
}

func MulIncrI8(a []int8, b []int8, incr []int8) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] * b[i]
	}
}

func MulIncrI16(a []int16, b []int16, incr []int16) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] * b[i]
	}
}

func MulIncrI32(a []int32, b []int32, incr []int32) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] * b[i]
	}
}

func MulIncrI64(a []int64, b []int64, incr []int64) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] * b[i]
	}
}

func MulIncrU(a []uint, b []uint, incr []uint) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] * b[i]
	}
}

func MulIncrU8(a []uint8, b []uint8, incr []uint8) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] * b[i]
	}
}

func MulIncrU16(a []uint16, b []uint16, incr []uint16) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] * b[i]
	}
}

func MulIncrU32(a []uint32, b []uint32, incr []uint32) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] * b[i]
	}
}

func MulIncrU64(a []uint64, b []uint64, incr []uint64) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] * b[i]
	}
}

//go:linkname MulIncrF32 github.com/chewxy/vecf32.IncrMul
func MulIncrF32(a []float32, b []float32, incr []float32)

//go:linkname MulIncrF64 github.com/chewxy/vecf64.IncrMul
func MulIncrF64(a []float64, b []float64, incr []float64)

func MulIncrC64(a []complex64, b []complex64, incr []complex64) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] * b[i]
	}
}

func MulIncrC128(a []complex128, b []complex128, incr []complex128) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] * b[i]
	}
}

func DivIncrI(a []int, b []int, incr []int) (err error) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	var errs errorIndices
	for i := range incr {
		if b[i] == 0 {
			errs = append(errs, i)
			incr[i] = 0
			continue
		}
		incr[i] += a[i] / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIncrI8(a []int8, b []int8, incr []int8) (err error) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	var errs errorIndices
	for i := range incr {
		if b[i] == 0 {
			errs = append(errs, i)
			incr[i] = 0
			continue
		}
		incr[i] += a[i] / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIncrI16(a []int16, b []int16, incr []int16) (err error) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	var errs errorIndices
	for i := range incr {
		if b[i] == 0 {
			errs = append(errs, i)
			incr[i] = 0
			continue
		}
		incr[i] += a[i] / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIncrI32(a []int32, b []int32, incr []int32) (err error) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	var errs errorIndices
	for i := range incr {
		if b[i] == 0 {
			errs = append(errs, i)
			incr[i] = 0
			continue
		}
		incr[i] += a[i] / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIncrI64(a []int64, b []int64, incr []int64) (err error) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	var errs errorIndices
	for i := range incr {
		if b[i] == 0 {
			errs = append(errs, i)
			incr[i] = 0
			continue
		}
		incr[i] += a[i] / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIncrU(a []uint, b []uint, incr []uint) (err error) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	var errs errorIndices
	for i := range incr {
		if b[i] == 0 {
			errs = append(errs, i)
			incr[i] = 0
			continue
		}
		incr[i] += a[i] / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIncrU8(a []uint8, b []uint8, incr []uint8) (err error) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	var errs errorIndices
	for i := range incr {
		if b[i] == 0 {
			errs = append(errs, i)
			incr[i] = 0
			continue
		}
		incr[i] += a[i] / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIncrU16(a []uint16, b []uint16, incr []uint16) (err error) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	var errs errorIndices
	for i := range incr {
		if b[i] == 0 {
			errs = append(errs, i)
			incr[i] = 0
			continue
		}
		incr[i] += a[i] / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIncrU32(a []uint32, b []uint32, incr []uint32) (err error) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	var errs errorIndices
	for i := range incr {
		if b[i] == 0 {
			errs = append(errs, i)
			incr[i] = 0
			continue
		}
		incr[i] += a[i] / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIncrU64(a []uint64, b []uint64, incr []uint64) (err error) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	var errs errorIndices
	for i := range incr {
		if b[i] == 0 {
			errs = append(errs, i)
			incr[i] = 0
			continue
		}
		incr[i] += a[i] / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

//go:linkname DivIncrF32 github.com/chewxy/vecf32.IncrDiv
func DivIncrF32(a []float32, b []float32, incr []float32)

//go:linkname DivIncrF64 github.com/chewxy/vecf64.IncrDiv
func DivIncrF64(a []float64, b []float64, incr []float64)

func DivIncrC64(a []complex64, b []complex64, incr []complex64) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] / b[i]
	}
}

func DivIncrC128(a []complex128, b []complex128, incr []complex128) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] / b[i]
	}
}

//go:linkname PowIncrF32 github.com/chewxy/vecf32.IncrPow
func PowIncrF32(a []float32, b []float32, incr []float32)

//go:linkname PowIncrF64 github.com/chewxy/vecf64.IncrPow
func PowIncrF64(a []float64, b []float64, incr []float64)

func PowIncrC64(a []complex64, b []complex64, incr []complex64) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += complex64(cmplx.Pow(complex128(a[i]), complex128(b[i])))
	}
}

func PowIncrC128(a []complex128, b []complex128, incr []complex128) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += cmplx.Pow(a[i], b[i])
	}
}

func ModIncrI(a []int, b []int, incr []int) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] % b[i]
	}
}

func ModIncrI8(a []int8, b []int8, incr []int8) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] % b[i]
	}
}

func ModIncrI16(a []int16, b []int16, incr []int16) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] % b[i]
	}
}

func ModIncrI32(a []int32, b []int32, incr []int32) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] % b[i]
	}
}

func ModIncrI64(a []int64, b []int64, incr []int64) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] % b[i]
	}
}

func ModIncrU(a []uint, b []uint, incr []uint) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] % b[i]
	}
}

func ModIncrU8(a []uint8, b []uint8, incr []uint8) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] % b[i]
	}
}

func ModIncrU16(a []uint16, b []uint16, incr []uint16) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] % b[i]
	}
}

func ModIncrU32(a []uint32, b []uint32, incr []uint32) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] % b[i]
	}
}

func ModIncrU64(a []uint64, b []uint64, incr []uint64) {
	a = a[:]
	b = b[:len(a)]
	incr = incr[:len(a)]
	for i := range incr {
		incr[i] += a[i] % b[i]
	}
}

//go:linkname ModIncrF32 github.com/chewxy/vecf32.IncrMod
func ModIncrF32(a []float32, b []float32, incr []float32)

//go:linkname ModIncrF64 github.com/chewxy/vecf64.IncrMod
func ModIncrF64(a []float64, b []float64, incr []float64)

func AddIterI(a []int, b []int, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] + b[j]
		}
	}
	return
}

func AddIterI8(a []int8, b []int8, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] + b[j]
		}
	}
	return
}

func AddIterI16(a []int16, b []int16, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] + b[j]
		}
	}
	return
}

func AddIterI32(a []int32, b []int32, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] + b[j]
		}
	}
	return
}

func AddIterI64(a []int64, b []int64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] + b[j]
		}
	}
	return
}

func AddIterU(a []uint, b []uint, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] + b[j]
		}
	}
	return
}

func AddIterU8(a []uint8, b []uint8, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] + b[j]
		}
	}
	return
}

func AddIterU16(a []uint16, b []uint16, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] + b[j]
		}
	}
	return
}

func AddIterU32(a []uint32, b []uint32, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] + b[j]
		}
	}
	return
}

func AddIterU64(a []uint64, b []uint64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] + b[j]
		}
	}
	return
}

func AddIterF32(a []float32, b []float32, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] + b[j]
		}
	}
	return
}

func AddIterF64(a []float64, b []float64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] + b[j]
		}
	}
	return
}

func AddIterC64(a []complex64, b []complex64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] + b[j]
		}
	}
	return
}

func AddIterC128(a []complex128, b []complex128, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] + b[j]
		}
	}
	return
}

func AddIterStr(a []string, b []string, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] + b[j]
		}
	}
	return
}

func SubIterI(a []int, b []int, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] - b[j]
		}
	}
	return
}

func SubIterI8(a []int8, b []int8, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] - b[j]
		}
	}
	return
}

func SubIterI16(a []int16, b []int16, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] - b[j]
		}
	}
	return
}

func SubIterI32(a []int32, b []int32, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] - b[j]
		}
	}
	return
}

func SubIterI64(a []int64, b []int64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] - b[j]
		}
	}
	return
}

func SubIterU(a []uint, b []uint, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] - b[j]
		}
	}
	return
}

func SubIterU8(a []uint8, b []uint8, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] - b[j]
		}
	}
	return
}

func SubIterU16(a []uint16, b []uint16, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] - b[j]
		}
	}
	return
}

func SubIterU32(a []uint32, b []uint32, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] - b[j]
		}
	}
	return
}

func SubIterU64(a []uint64, b []uint64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] - b[j]
		}
	}
	return
}

func SubIterF32(a []float32, b []float32, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] - b[j]
		}
	}
	return
}

func SubIterF64(a []float64, b []float64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] - b[j]
		}
	}
	return
}

func SubIterC64(a []complex64, b []complex64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] - b[j]
		}
	}
	return
}

func SubIterC128(a []complex128, b []complex128, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] - b[j]
		}
	}
	return
}

func MulIterI(a []int, b []int, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] * b[j]
		}
	}
	return
}

func MulIterI8(a []int8, b []int8, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] * b[j]
		}
	}
	return
}

func MulIterI16(a []int16, b []int16, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] * b[j]
		}
	}
	return
}

func MulIterI32(a []int32, b []int32, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] * b[j]
		}
	}
	return
}

func MulIterI64(a []int64, b []int64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] * b[j]
		}
	}
	return
}

func MulIterU(a []uint, b []uint, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] * b[j]
		}
	}
	return
}

func MulIterU8(a []uint8, b []uint8, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] * b[j]
		}
	}
	return
}

func MulIterU16(a []uint16, b []uint16, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] * b[j]
		}
	}
	return
}

func MulIterU32(a []uint32, b []uint32, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] * b[j]
		}
	}
	return
}

func MulIterU64(a []uint64, b []uint64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] * b[j]
		}
	}
	return
}

func MulIterF32(a []float32, b []float32, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] * b[j]
		}
	}
	return
}

func MulIterF64(a []float64, b []float64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] * b[j]
		}
	}
	return
}

func MulIterC64(a []complex64, b []complex64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] * b[j]
		}
	}
	return
}

func MulIterC128(a []complex128, b []complex128, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] * b[j]
		}
	}
	return
}

func DivIterI(a []int, b []int, ait Iterator, bit Iterator) (err error) {
	var errs errorIndices
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if b[j] == 0 {
				errs = append(errs, i)
				a[i] = 0
				continue
			}
			a[i] = a[i] / b[j]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterI8(a []int8, b []int8, ait Iterator, bit Iterator) (err error) {
	var errs errorIndices
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if b[j] == 0 {
				errs = append(errs, i)
				a[i] = 0
				continue
			}
			a[i] = a[i] / b[j]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterI16(a []int16, b []int16, ait Iterator, bit Iterator) (err error) {
	var errs errorIndices
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if b[j] == 0 {
				errs = append(errs, i)
				a[i] = 0
				continue
			}
			a[i] = a[i] / b[j]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterI32(a []int32, b []int32, ait Iterator, bit Iterator) (err error) {
	var errs errorIndices
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if b[j] == 0 {
				errs = append(errs, i)
				a[i] = 0
				continue
			}
			a[i] = a[i] / b[j]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterI64(a []int64, b []int64, ait Iterator, bit Iterator) (err error) {
	var errs errorIndices
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if b[j] == 0 {
				errs = append(errs, i)
				a[i] = 0
				continue
			}
			a[i] = a[i] / b[j]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterU(a []uint, b []uint, ait Iterator, bit Iterator) (err error) {
	var errs errorIndices
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if b[j] == 0 {
				errs = append(errs, i)
				a[i] = 0
				continue
			}
			a[i] = a[i] / b[j]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterU8(a []uint8, b []uint8, ait Iterator, bit Iterator) (err error) {
	var errs errorIndices
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if b[j] == 0 {
				errs = append(errs, i)
				a[i] = 0
				continue
			}
			a[i] = a[i] / b[j]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterU16(a []uint16, b []uint16, ait Iterator, bit Iterator) (err error) {
	var errs errorIndices
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if b[j] == 0 {
				errs = append(errs, i)
				a[i] = 0
				continue
			}
			a[i] = a[i] / b[j]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterU32(a []uint32, b []uint32, ait Iterator, bit Iterator) (err error) {
	var errs errorIndices
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if b[j] == 0 {
				errs = append(errs, i)
				a[i] = 0
				continue
			}
			a[i] = a[i] / b[j]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterU64(a []uint64, b []uint64, ait Iterator, bit Iterator) (err error) {
	var errs errorIndices
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			if b[j] == 0 {
				errs = append(errs, i)
				a[i] = 0
				continue
			}
			a[i] = a[i] / b[j]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterF32(a []float32, b []float32, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] / b[j]
		}
	}
	return
}

func DivIterF64(a []float64, b []float64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] / b[j]
		}
	}
	return
}

func DivIterC64(a []complex64, b []complex64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] / b[j]
		}
	}
	return
}

func DivIterC128(a []complex128, b []complex128, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] / b[j]
		}
	}
	return
}

func PowIterF32(a []float32, b []float32, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = math32.Pow(a[i], b[j])
		}
	}
	return
}

func PowIterF64(a []float64, b []float64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = math.Pow(a[i], b[j])
		}
	}
	return
}

func PowIterC64(a []complex64, b []complex64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = complex64(cmplx.Pow(complex128(a[i]), complex128(b[j])))
		}
	}
	return
}

func PowIterC128(a []complex128, b []complex128, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = cmplx.Pow(a[i], b[j])
		}
	}
	return
}

func ModIterI(a []int, b []int, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] % b[j]
		}
	}
	return
}

func ModIterI8(a []int8, b []int8, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] % b[j]
		}
	}
	return
}

func ModIterI16(a []int16, b []int16, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] % b[j]
		}
	}
	return
}

func ModIterI32(a []int32, b []int32, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] % b[j]
		}
	}
	return
}

func ModIterI64(a []int64, b []int64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] % b[j]
		}
	}
	return
}

func ModIterU(a []uint, b []uint, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] % b[j]
		}
	}
	return
}

func ModIterU8(a []uint8, b []uint8, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] % b[j]
		}
	}
	return
}

func ModIterU16(a []uint16, b []uint16, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] % b[j]
		}
	}
	return
}

func ModIterU32(a []uint32, b []uint32, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] % b[j]
		}
	}
	return
}

func ModIterU64(a []uint64, b []uint64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = a[i] % b[j]
		}
	}
	return
}

func ModIterF32(a []float32, b []float32, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = math32.Mod(a[i], b[j])
		}
	}
	return
}

func ModIterF64(a []float64, b []float64, ait Iterator, bit Iterator) (err error) {
	var i, j int
	var validi, validj bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj {
			a[i] = math.Mod(a[i], b[j])
		}
	}
	return
}

func AddIterIncrI(a []int, b []int, incr []int, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] + b[j]
		}
	}
	return
}

func AddIterIncrI8(a []int8, b []int8, incr []int8, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] + b[j]
		}
	}
	return
}

func AddIterIncrI16(a []int16, b []int16, incr []int16, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] + b[j]
		}
	}
	return
}

func AddIterIncrI32(a []int32, b []int32, incr []int32, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] + b[j]
		}
	}
	return
}

func AddIterIncrI64(a []int64, b []int64, incr []int64, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] + b[j]
		}
	}
	return
}

func AddIterIncrU(a []uint, b []uint, incr []uint, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] + b[j]
		}
	}
	return
}

func AddIterIncrU8(a []uint8, b []uint8, incr []uint8, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] + b[j]
		}
	}
	return
}

func AddIterIncrU16(a []uint16, b []uint16, incr []uint16, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] + b[j]
		}
	}
	return
}

func AddIterIncrU32(a []uint32, b []uint32, incr []uint32, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] + b[j]
		}
	}
	return
}

func AddIterIncrU64(a []uint64, b []uint64, incr []uint64, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] + b[j]
		}
	}
	return
}

func AddIterIncrF32(a []float32, b []float32, incr []float32, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] + b[j]
		}
	}
	return
}

func AddIterIncrF64(a []float64, b []float64, incr []float64, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] + b[j]
		}
	}
	return
}

func AddIterIncrC64(a []complex64, b []complex64, incr []complex64, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] + b[j]
		}
	}
	return
}

func AddIterIncrC128(a []complex128, b []complex128, incr []complex128, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] + b[j]
		}
	}
	return
}

func AddIterIncrStr(a []string, b []string, incr []string, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] + b[j]
		}
	}
	return
}

func SubIterIncrI(a []int, b []int, incr []int, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] - b[j]
		}
	}
	return
}

func SubIterIncrI8(a []int8, b []int8, incr []int8, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] - b[j]
		}
	}
	return
}

func SubIterIncrI16(a []int16, b []int16, incr []int16, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] - b[j]
		}
	}
	return
}

func SubIterIncrI32(a []int32, b []int32, incr []int32, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] - b[j]
		}
	}
	return
}

func SubIterIncrI64(a []int64, b []int64, incr []int64, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] - b[j]
		}
	}
	return
}

func SubIterIncrU(a []uint, b []uint, incr []uint, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] - b[j]
		}
	}
	return
}

func SubIterIncrU8(a []uint8, b []uint8, incr []uint8, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] - b[j]
		}
	}
	return
}

func SubIterIncrU16(a []uint16, b []uint16, incr []uint16, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] - b[j]
		}
	}
	return
}

func SubIterIncrU32(a []uint32, b []uint32, incr []uint32, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] - b[j]
		}
	}
	return
}

func SubIterIncrU64(a []uint64, b []uint64, incr []uint64, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] - b[j]
		}
	}
	return
}

func SubIterIncrF32(a []float32, b []float32, incr []float32, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] - b[j]
		}
	}
	return
}

func SubIterIncrF64(a []float64, b []float64, incr []float64, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] - b[j]
		}
	}
	return
}

func SubIterIncrC64(a []complex64, b []complex64, incr []complex64, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] - b[j]
		}
	}
	return
}

func SubIterIncrC128(a []complex128, b []complex128, incr []complex128, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] - b[j]
		}
	}
	return
}

func MulIterIncrI(a []int, b []int, incr []int, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] * b[j]
		}
	}
	return
}

func MulIterIncrI8(a []int8, b []int8, incr []int8, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] * b[j]
		}
	}
	return
}

func MulIterIncrI16(a []int16, b []int16, incr []int16, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] * b[j]
		}
	}
	return
}

func MulIterIncrI32(a []int32, b []int32, incr []int32, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] * b[j]
		}
	}
	return
}

func MulIterIncrI64(a []int64, b []int64, incr []int64, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] * b[j]
		}
	}
	return
}

func MulIterIncrU(a []uint, b []uint, incr []uint, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] * b[j]
		}
	}
	return
}

func MulIterIncrU8(a []uint8, b []uint8, incr []uint8, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] * b[j]
		}
	}
	return
}

func MulIterIncrU16(a []uint16, b []uint16, incr []uint16, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] * b[j]
		}
	}
	return
}

func MulIterIncrU32(a []uint32, b []uint32, incr []uint32, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] * b[j]
		}
	}
	return
}

func MulIterIncrU64(a []uint64, b []uint64, incr []uint64, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] * b[j]
		}
	}
	return
}

func MulIterIncrF32(a []float32, b []float32, incr []float32, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] * b[j]
		}
	}
	return
}

func MulIterIncrF64(a []float64, b []float64, incr []float64, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] * b[j]
		}
	}
	return
}

func MulIterIncrC64(a []complex64, b []complex64, incr []complex64, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] * b[j]
		}
	}
	return
}

func MulIterIncrC128(a []complex128, b []complex128, incr []complex128, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] * b[j]
		}
	}
	return
}

func DivIterIncrI(a []int, b []int, incr []int, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var errs errorIndices
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			if b[j] == 0 {
				errs = append(errs, i)
				incr[i] = 0
				continue
			}
			incr[k] += a[i] / b[j]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterIncrI8(a []int8, b []int8, incr []int8, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var errs errorIndices
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			if b[j] == 0 {
				errs = append(errs, i)
				incr[i] = 0
				continue
			}
			incr[k] += a[i] / b[j]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterIncrI16(a []int16, b []int16, incr []int16, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var errs errorIndices
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			if b[j] == 0 {
				errs = append(errs, i)
				incr[i] = 0
				continue
			}
			incr[k] += a[i] / b[j]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterIncrI32(a []int32, b []int32, incr []int32, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var errs errorIndices
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			if b[j] == 0 {
				errs = append(errs, i)
				incr[i] = 0
				continue
			}
			incr[k] += a[i] / b[j]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterIncrI64(a []int64, b []int64, incr []int64, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var errs errorIndices
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			if b[j] == 0 {
				errs = append(errs, i)
				incr[i] = 0
				continue
			}
			incr[k] += a[i] / b[j]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterIncrU(a []uint, b []uint, incr []uint, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var errs errorIndices
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			if b[j] == 0 {
				errs = append(errs, i)
				incr[i] = 0
				continue
			}
			incr[k] += a[i] / b[j]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterIncrU8(a []uint8, b []uint8, incr []uint8, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var errs errorIndices
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			if b[j] == 0 {
				errs = append(errs, i)
				incr[i] = 0
				continue
			}
			incr[k] += a[i] / b[j]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterIncrU16(a []uint16, b []uint16, incr []uint16, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var errs errorIndices
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			if b[j] == 0 {
				errs = append(errs, i)
				incr[i] = 0
				continue
			}
			incr[k] += a[i] / b[j]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterIncrU32(a []uint32, b []uint32, incr []uint32, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var errs errorIndices
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			if b[j] == 0 {
				errs = append(errs, i)
				incr[i] = 0
				continue
			}
			incr[k] += a[i] / b[j]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterIncrU64(a []uint64, b []uint64, incr []uint64, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var errs errorIndices
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			if b[j] == 0 {
				errs = append(errs, i)
				incr[i] = 0
				continue
			}
			incr[k] += a[i] / b[j]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterIncrF32(a []float32, b []float32, incr []float32, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] / b[j]
		}
	}
	return
}

func DivIterIncrF64(a []float64, b []float64, incr []float64, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] / b[j]
		}
	}
	return
}

func DivIterIncrC64(a []complex64, b []complex64, incr []complex64, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] / b[j]
		}
	}
	return
}

func DivIterIncrC128(a []complex128, b []complex128, incr []complex128, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] / b[j]
		}
	}
	return
}

func PowIterIncrF32(a []float32, b []float32, incr []float32, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += math32.Pow(a[i], b[j])
		}
	}
	return
}

func PowIterIncrF64(a []float64, b []float64, incr []float64, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += math.Pow(a[i], b[j])
		}
	}
	return
}

func PowIterIncrC64(a []complex64, b []complex64, incr []complex64, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += complex64(cmplx.Pow(complex128(a[i]), complex128(b[j])))
		}
	}
	return
}

func PowIterIncrC128(a []complex128, b []complex128, incr []complex128, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += cmplx.Pow(a[i], b[j])
		}
	}
	return
}

func ModIterIncrI(a []int, b []int, incr []int, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] % b[j]
		}
	}
	return
}

func ModIterIncrI8(a []int8, b []int8, incr []int8, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] % b[j]
		}
	}
	return
}

func ModIterIncrI16(a []int16, b []int16, incr []int16, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] % b[j]
		}
	}
	return
}

func ModIterIncrI32(a []int32, b []int32, incr []int32, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] % b[j]
		}
	}
	return
}

func ModIterIncrI64(a []int64, b []int64, incr []int64, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] % b[j]
		}
	}
	return
}

func ModIterIncrU(a []uint, b []uint, incr []uint, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] % b[j]
		}
	}
	return
}

func ModIterIncrU8(a []uint8, b []uint8, incr []uint8, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] % b[j]
		}
	}
	return
}

func ModIterIncrU16(a []uint16, b []uint16, incr []uint16, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] % b[j]
		}
	}
	return
}

func ModIterIncrU32(a []uint32, b []uint32, incr []uint32, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] % b[j]
		}
	}
	return
}

func ModIterIncrU64(a []uint64, b []uint64, incr []uint64, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += a[i] % b[j]
		}
	}
	return
}

func ModIterIncrF32(a []float32, b []float32, incr []float32, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += math32.Mod(a[i], b[j])
		}
	}
	return
}

func ModIterIncrF64(a []float64, b []float64, incr []float64, ait Iterator, bit Iterator, iit Iterator) (err error) {
	var i, j, k int
	var validi, validj, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if j, validj, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validj && validk {
			incr[k] += math.Mod(a[i], b[j])
		}
	}
	return
}
