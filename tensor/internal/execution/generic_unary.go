package execution

import (
	"math"
	"math/cmplx"

	"github.com/chewxy/math32"
)

/*
GENERATED FILE. DO NOT EDIT
*/

func NegI(a []int) {
	for i := range a {
		a[i] = -a[i]
	}
}

func NegI8(a []int8) {
	for i := range a {
		a[i] = -a[i]
	}
}

func NegI16(a []int16) {
	for i := range a {
		a[i] = -a[i]
	}
}

func NegI32(a []int32) {
	for i := range a {
		a[i] = -a[i]
	}
}

func NegI64(a []int64) {
	for i := range a {
		a[i] = -a[i]
	}
}

func NegU(a []uint) {
	for i := range a {
		a[i] = -a[i]
	}
}

func NegU8(a []uint8) {
	for i := range a {
		a[i] = -a[i]
	}
}

func NegU16(a []uint16) {
	for i := range a {
		a[i] = -a[i]
	}
}

func NegU32(a []uint32) {
	for i := range a {
		a[i] = -a[i]
	}
}

func NegU64(a []uint64) {
	for i := range a {
		a[i] = -a[i]
	}
}

func NegF32(a []float32) {
	for i := range a {
		a[i] = -a[i]
	}
}

func NegF64(a []float64) {
	for i := range a {
		a[i] = -a[i]
	}
}

func NegC64(a []complex64) {
	for i := range a {
		a[i] = -a[i]
	}
}

func NegC128(a []complex128) {
	for i := range a {
		a[i] = -a[i]
	}
}

func InvI(a []int) {
	for i := range a {
		a[i] = 1 / a[i]
	}
}

func InvI8(a []int8) {
	for i := range a {
		a[i] = 1 / a[i]
	}
}

func InvI16(a []int16) {
	for i := range a {
		a[i] = 1 / a[i]
	}
}

func InvI32(a []int32) {
	for i := range a {
		a[i] = 1 / a[i]
	}
}

func InvI64(a []int64) {
	for i := range a {
		a[i] = 1 / a[i]
	}
}

func InvU(a []uint) {
	for i := range a {
		a[i] = 1 / a[i]
	}
}

func InvU8(a []uint8) {
	for i := range a {
		a[i] = 1 / a[i]
	}
}

func InvU16(a []uint16) {
	for i := range a {
		a[i] = 1 / a[i]
	}
}

func InvU32(a []uint32) {
	for i := range a {
		a[i] = 1 / a[i]
	}
}

func InvU64(a []uint64) {
	for i := range a {
		a[i] = 1 / a[i]
	}
}

func InvF32(a []float32) {
	for i := range a {
		a[i] = 1 / a[i]
	}
}

func InvF64(a []float64) {
	for i := range a {
		a[i] = 1 / a[i]
	}
}

func InvC64(a []complex64) {
	for i := range a {
		a[i] = 1 / a[i]
	}
}

func InvC128(a []complex128) {
	for i := range a {
		a[i] = 1 / a[i]
	}
}

func SquareI(a []int) {
	for i := range a {
		a[i] = a[i] * a[i]
	}
}

func SquareI8(a []int8) {
	for i := range a {
		a[i] = a[i] * a[i]
	}
}

func SquareI16(a []int16) {
	for i := range a {
		a[i] = a[i] * a[i]
	}
}

func SquareI32(a []int32) {
	for i := range a {
		a[i] = a[i] * a[i]
	}
}

func SquareI64(a []int64) {
	for i := range a {
		a[i] = a[i] * a[i]
	}
}

func SquareU(a []uint) {
	for i := range a {
		a[i] = a[i] * a[i]
	}
}

func SquareU8(a []uint8) {
	for i := range a {
		a[i] = a[i] * a[i]
	}
}

func SquareU16(a []uint16) {
	for i := range a {
		a[i] = a[i] * a[i]
	}
}

func SquareU32(a []uint32) {
	for i := range a {
		a[i] = a[i] * a[i]
	}
}

func SquareU64(a []uint64) {
	for i := range a {
		a[i] = a[i] * a[i]
	}
}

func SquareF32(a []float32) {
	for i := range a {
		a[i] = a[i] * a[i]
	}
}

func SquareF64(a []float64) {
	for i := range a {
		a[i] = a[i] * a[i]
	}
}

func SquareC64(a []complex64) {
	for i := range a {
		a[i] = a[i] * a[i]
	}
}

func SquareC128(a []complex128) {
	for i := range a {
		a[i] = a[i] * a[i]
	}
}

func CubeI(a []int) {
	for i := range a {
		a[i] = a[i] * a[i] * a[i]
	}
}

func CubeI8(a []int8) {
	for i := range a {
		a[i] = a[i] * a[i] * a[i]
	}
}

func CubeI16(a []int16) {
	for i := range a {
		a[i] = a[i] * a[i] * a[i]
	}
}

func CubeI32(a []int32) {
	for i := range a {
		a[i] = a[i] * a[i] * a[i]
	}
}

func CubeI64(a []int64) {
	for i := range a {
		a[i] = a[i] * a[i] * a[i]
	}
}

func CubeU(a []uint) {
	for i := range a {
		a[i] = a[i] * a[i] * a[i]
	}
}

func CubeU8(a []uint8) {
	for i := range a {
		a[i] = a[i] * a[i] * a[i]
	}
}

func CubeU16(a []uint16) {
	for i := range a {
		a[i] = a[i] * a[i] * a[i]
	}
}

func CubeU32(a []uint32) {
	for i := range a {
		a[i] = a[i] * a[i] * a[i]
	}
}

func CubeU64(a []uint64) {
	for i := range a {
		a[i] = a[i] * a[i] * a[i]
	}
}

func CubeF32(a []float32) {
	for i := range a {
		a[i] = a[i] * a[i] * a[i]
	}
}

func CubeF64(a []float64) {
	for i := range a {
		a[i] = a[i] * a[i] * a[i]
	}
}

func CubeC64(a []complex64) {
	for i := range a {
		a[i] = a[i] * a[i] * a[i]
	}
}

func CubeC128(a []complex128) {
	for i := range a {
		a[i] = a[i] * a[i] * a[i]
	}
}

func ExpF32(a []float32) {
	for i := range a {
		a[i] = math32.Exp(a[i])
	}
}

func ExpF64(a []float64) {
	for i := range a {
		a[i] = math.Exp(a[i])
	}
}

func ExpC64(a []complex64) {
	for i := range a {
		a[i] = complex64(cmplx.Exp(complex128(a[i])))
	}
}

func ExpC128(a []complex128) {
	for i := range a {
		a[i] = cmplx.Exp(a[i])
	}
}

func TanhF32(a []float32) {
	for i := range a {
		a[i] = math32.Tanh(a[i])
	}
}

func TanhF64(a []float64) {
	for i := range a {
		a[i] = math.Tanh(a[i])
	}
}

func TanhC64(a []complex64) {
	for i := range a {
		a[i] = complex64(cmplx.Tanh(complex128(a[i])))
	}
}

func TanhC128(a []complex128) {
	for i := range a {
		a[i] = cmplx.Tanh(a[i])
	}
}

func LogF32(a []float32) {
	for i := range a {
		a[i] = math32.Log(a[i])
	}
}

func LogF64(a []float64) {
	for i := range a {
		a[i] = math.Log(a[i])
	}
}

func LogC64(a []complex64) {
	for i := range a {
		a[i] = complex64(cmplx.Log(complex128(a[i])))
	}
}

func LogC128(a []complex128) {
	for i := range a {
		a[i] = cmplx.Log(a[i])
	}
}

func Log2F32(a []float32) {
	for i := range a {
		a[i] = math32.Log2(a[i])
	}
}

func Log2F64(a []float64) {
	for i := range a {
		a[i] = math.Log2(a[i])
	}
}

func Log10F32(a []float32) {
	for i := range a {
		a[i] = math32.Log10(a[i])
	}
}

func Log10F64(a []float64) {
	for i := range a {
		a[i] = math.Log10(a[i])
	}
}

func Log10C64(a []complex64) {
	for i := range a {
		a[i] = complex64(cmplx.Log10(complex128(a[i])))
	}
}

func Log10C128(a []complex128) {
	for i := range a {
		a[i] = cmplx.Log10(a[i])
	}
}

func SqrtF32(a []float32) {
	for i := range a {
		a[i] = math32.Sqrt(a[i])
	}
}

func SqrtF64(a []float64) {
	for i := range a {
		a[i] = math.Sqrt(a[i])
	}
}

func SqrtC64(a []complex64) {
	for i := range a {
		a[i] = complex64(cmplx.Sqrt(complex128(a[i])))
	}
}

func SqrtC128(a []complex128) {
	for i := range a {
		a[i] = cmplx.Sqrt(a[i])
	}
}

func CbrtF32(a []float32) {
	for i := range a {
		a[i] = math32.Cbrt(a[i])
	}
}

func CbrtF64(a []float64) {
	for i := range a {
		a[i] = math.Cbrt(a[i])
	}
}

func InvSqrtF32(a []float32) {
	for i := range a {
		a[i] = float32(1) / math32.Sqrt(a[i])
	}
}

func InvSqrtF64(a []float64) {
	for i := range a {
		a[i] = float64(1) / math.Sqrt(a[i])
	}
}

func NegIterI(a []int, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = -a[i]
		}
	}
	return
}

func NegIterI8(a []int8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = -a[i]
		}
	}
	return
}

func NegIterI16(a []int16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = -a[i]
		}
	}
	return
}

func NegIterI32(a []int32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = -a[i]
		}
	}
	return
}

func NegIterI64(a []int64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = -a[i]
		}
	}
	return
}

func NegIterU(a []uint, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = -a[i]
		}
	}
	return
}

func NegIterU8(a []uint8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = -a[i]
		}
	}
	return
}

func NegIterU16(a []uint16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = -a[i]
		}
	}
	return
}

func NegIterU32(a []uint32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = -a[i]
		}
	}
	return
}

func NegIterU64(a []uint64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = -a[i]
		}
	}
	return
}

func NegIterF32(a []float32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = -a[i]
		}
	}
	return
}

func NegIterF64(a []float64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = -a[i]
		}
	}
	return
}

func NegIterC64(a []complex64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = -a[i]
		}
	}
	return
}

func NegIterC128(a []complex128, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = -a[i]
		}
	}
	return
}

func InvIterI(a []int, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = 1 / a[i]
		}
	}
	return
}

func InvIterI8(a []int8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = 1 / a[i]
		}
	}
	return
}

func InvIterI16(a []int16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = 1 / a[i]
		}
	}
	return
}

func InvIterI32(a []int32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = 1 / a[i]
		}
	}
	return
}

func InvIterI64(a []int64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = 1 / a[i]
		}
	}
	return
}

func InvIterU(a []uint, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = 1 / a[i]
		}
	}
	return
}

func InvIterU8(a []uint8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = 1 / a[i]
		}
	}
	return
}

func InvIterU16(a []uint16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = 1 / a[i]
		}
	}
	return
}

func InvIterU32(a []uint32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = 1 / a[i]
		}
	}
	return
}

func InvIterU64(a []uint64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = 1 / a[i]
		}
	}
	return
}

func InvIterF32(a []float32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = 1 / a[i]
		}
	}
	return
}

func InvIterF64(a []float64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = 1 / a[i]
		}
	}
	return
}

func InvIterC64(a []complex64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = 1 / a[i]
		}
	}
	return
}

func InvIterC128(a []complex128, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = 1 / a[i]
		}
	}
	return
}

func SquareIterI(a []int, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * a[i]
		}
	}
	return
}

func SquareIterI8(a []int8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * a[i]
		}
	}
	return
}

func SquareIterI16(a []int16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * a[i]
		}
	}
	return
}

func SquareIterI32(a []int32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * a[i]
		}
	}
	return
}

func SquareIterI64(a []int64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * a[i]
		}
	}
	return
}

func SquareIterU(a []uint, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * a[i]
		}
	}
	return
}

func SquareIterU8(a []uint8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * a[i]
		}
	}
	return
}

func SquareIterU16(a []uint16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * a[i]
		}
	}
	return
}

func SquareIterU32(a []uint32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * a[i]
		}
	}
	return
}

func SquareIterU64(a []uint64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * a[i]
		}
	}
	return
}

func SquareIterF32(a []float32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * a[i]
		}
	}
	return
}

func SquareIterF64(a []float64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * a[i]
		}
	}
	return
}

func SquareIterC64(a []complex64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * a[i]
		}
	}
	return
}

func SquareIterC128(a []complex128, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * a[i]
		}
	}
	return
}

func CubeIterI(a []int, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * a[i] * a[i]
		}
	}
	return
}

func CubeIterI8(a []int8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * a[i] * a[i]
		}
	}
	return
}

func CubeIterI16(a []int16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * a[i] * a[i]
		}
	}
	return
}

func CubeIterI32(a []int32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * a[i] * a[i]
		}
	}
	return
}

func CubeIterI64(a []int64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * a[i] * a[i]
		}
	}
	return
}

func CubeIterU(a []uint, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * a[i] * a[i]
		}
	}
	return
}

func CubeIterU8(a []uint8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * a[i] * a[i]
		}
	}
	return
}

func CubeIterU16(a []uint16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * a[i] * a[i]
		}
	}
	return
}

func CubeIterU32(a []uint32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * a[i] * a[i]
		}
	}
	return
}

func CubeIterU64(a []uint64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * a[i] * a[i]
		}
	}
	return
}

func CubeIterF32(a []float32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * a[i] * a[i]
		}
	}
	return
}

func CubeIterF64(a []float64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * a[i] * a[i]
		}
	}
	return
}

func CubeIterC64(a []complex64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * a[i] * a[i]
		}
	}
	return
}

func CubeIterC128(a []complex128, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * a[i] * a[i]
		}
	}
	return
}

func ExpIterF32(a []float32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = math32.Exp(a[i])
		}
	}
	return
}

func ExpIterF64(a []float64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = math.Exp(a[i])
		}
	}
	return
}

func ExpIterC64(a []complex64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = complex64(cmplx.Exp(complex128(a[i])))
		}
	}
	return
}

func ExpIterC128(a []complex128, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = cmplx.Exp(a[i])
		}
	}
	return
}

func TanhIterF32(a []float32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = math32.Tanh(a[i])
		}
	}
	return
}

func TanhIterF64(a []float64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = math.Tanh(a[i])
		}
	}
	return
}

func TanhIterC64(a []complex64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = complex64(cmplx.Tanh(complex128(a[i])))
		}
	}
	return
}

func TanhIterC128(a []complex128, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = cmplx.Tanh(a[i])
		}
	}
	return
}

func LogIterF32(a []float32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = math32.Log(a[i])
		}
	}
	return
}

func LogIterF64(a []float64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = math.Log(a[i])
		}
	}
	return
}

func LogIterC64(a []complex64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = complex64(cmplx.Log(complex128(a[i])))
		}
	}
	return
}

func LogIterC128(a []complex128, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = cmplx.Log(a[i])
		}
	}
	return
}

func Log2IterF32(a []float32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = math32.Log2(a[i])
		}
	}
	return
}

func Log2IterF64(a []float64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = math.Log2(a[i])
		}
	}
	return
}

func Log10IterF32(a []float32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = math32.Log10(a[i])
		}
	}
	return
}

func Log10IterF64(a []float64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = math.Log10(a[i])
		}
	}
	return
}

func Log10IterC64(a []complex64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = complex64(cmplx.Log10(complex128(a[i])))
		}
	}
	return
}

func Log10IterC128(a []complex128, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = cmplx.Log10(a[i])
		}
	}
	return
}

func SqrtIterF32(a []float32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = math32.Sqrt(a[i])
		}
	}
	return
}

func SqrtIterF64(a []float64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = math.Sqrt(a[i])
		}
	}
	return
}

func SqrtIterC64(a []complex64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = complex64(cmplx.Sqrt(complex128(a[i])))
		}
	}
	return
}

func SqrtIterC128(a []complex128, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = cmplx.Sqrt(a[i])
		}
	}
	return
}

func CbrtIterF32(a []float32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = math32.Cbrt(a[i])
		}
	}
	return
}

func CbrtIterF64(a []float64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = math.Cbrt(a[i])
		}
	}
	return
}

func InvSqrtIterF32(a []float32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = float32(1) / math32.Sqrt(a[i])
		}
	}
	return
}

func InvSqrtIterF64(a []float64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = float64(1) / math.Sqrt(a[i])
		}
	}
	return
}

func AbsI(a []int) {
	for i := range a {
		if a[i] < 0 {
			a[i] = -a[i]
		}
	}
}

func AbsI8(a []int8) {
	for i := range a {
		if a[i] < 0 {
			a[i] = -a[i]
		}
	}
}

func AbsI16(a []int16) {
	for i := range a {
		if a[i] < 0 {
			a[i] = -a[i]
		}
	}
}

func AbsI32(a []int32) {
	for i := range a {
		if a[i] < 0 {
			a[i] = -a[i]
		}
	}
}

func AbsI64(a []int64) {
	for i := range a {
		if a[i] < 0 {
			a[i] = -a[i]
		}
	}
}

func AbsF32(a []float32) {
	for i := range a {
		a[i] = math32.Abs(a[i])
	}
}

func AbsF64(a []float64) {
	for i := range a {
		a[i] = math.Abs(a[i])
	}
}

func SignI(a []int) {
	for i := range a {
		if a[i] < 0 {
			a[i] = -1
		} else if a[i] > 0 {
			a[i] = 1
		}
	}
}

func SignI8(a []int8) {
	for i := range a {
		if a[i] < 0 {
			a[i] = -1
		} else if a[i] > 0 {
			a[i] = 1
		}
	}
}

func SignI16(a []int16) {
	for i := range a {
		if a[i] < 0 {
			a[i] = -1
		} else if a[i] > 0 {
			a[i] = 1
		}
	}
}

func SignI32(a []int32) {
	for i := range a {
		if a[i] < 0 {
			a[i] = -1
		} else if a[i] > 0 {
			a[i] = 1
		}
	}
}

func SignI64(a []int64) {
	for i := range a {
		if a[i] < 0 {
			a[i] = -1
		} else if a[i] > 0 {
			a[i] = 1
		}
	}
}

func SignF32(a []float32) {
	for i := range a {
		if a[i] < 0 {
			a[i] = -1
		} else if a[i] > 0 {
			a[i] = 1
		}
	}
}

func SignF64(a []float64) {
	for i := range a {
		if a[i] < 0 {
			a[i] = -1
		} else if a[i] > 0 {
			a[i] = 1
		}
	}
}

func AbsIterI(a []int, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < 0 {
				a[i] = -a[i]
			}
		}
	}
	return
}

func AbsIterI8(a []int8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < 0 {
				a[i] = -a[i]
			}
		}
	}
	return
}

func AbsIterI16(a []int16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < 0 {
				a[i] = -a[i]
			}
		}
	}
	return
}

func AbsIterI32(a []int32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < 0 {
				a[i] = -a[i]
			}
		}
	}
	return
}

func AbsIterI64(a []int64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < 0 {
				a[i] = -a[i]
			}
		}
	}
	return
}

func AbsIterF32(a []float32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = math32.Abs(a[i])
		}
	}
	return
}

func AbsIterF64(a []float64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = math.Abs(a[i])
		}
	}
	return
}

func SignIterI(a []int, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < 0 {
				a[i] = -1
			} else if a[i] > 0 {
				a[i] = 1
			}
		}
	}
	return
}

func SignIterI8(a []int8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < 0 {
				a[i] = -1
			} else if a[i] > 0 {
				a[i] = 1
			}
		}
	}
	return
}

func SignIterI16(a []int16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < 0 {
				a[i] = -1
			} else if a[i] > 0 {
				a[i] = 1
			}
		}
	}
	return
}

func SignIterI32(a []int32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < 0 {
				a[i] = -1
			} else if a[i] > 0 {
				a[i] = 1
			}
		}
	}
	return
}

func SignIterI64(a []int64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < 0 {
				a[i] = -1
			} else if a[i] > 0 {
				a[i] = 1
			}
		}
	}
	return
}

func SignIterF32(a []float32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < 0 {
				a[i] = -1
			} else if a[i] > 0 {
				a[i] = 1
			}
		}
	}
	return
}

func SignIterF64(a []float64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < 0 {
				a[i] = -1
			} else if a[i] > 0 {
				a[i] = 1
			}
		}
	}
	return
}

func ClampI(a []int, min int, max int) {
	for i := range a {
		if a[i] < min {
			a[i] = min
			continue
		}
		if a[i] > max {
			a[i] = max
		}
	}
}

func ClampI8(a []int8, min int8, max int8) {
	for i := range a {
		if a[i] < min {
			a[i] = min
			continue
		}
		if a[i] > max {
			a[i] = max
		}
	}
}

func ClampI16(a []int16, min int16, max int16) {
	for i := range a {
		if a[i] < min {
			a[i] = min
			continue
		}
		if a[i] > max {
			a[i] = max
		}
	}
}

func ClampI32(a []int32, min int32, max int32) {
	for i := range a {
		if a[i] < min {
			a[i] = min
			continue
		}
		if a[i] > max {
			a[i] = max
		}
	}
}

func ClampI64(a []int64, min int64, max int64) {
	for i := range a {
		if a[i] < min {
			a[i] = min
			continue
		}
		if a[i] > max {
			a[i] = max
		}
	}
}

func ClampU(a []uint, min uint, max uint) {
	for i := range a {
		if a[i] < min {
			a[i] = min
			continue
		}
		if a[i] > max {
			a[i] = max
		}
	}
}

func ClampU8(a []uint8, min uint8, max uint8) {
	for i := range a {
		if a[i] < min {
			a[i] = min
			continue
		}
		if a[i] > max {
			a[i] = max
		}
	}
}

func ClampU16(a []uint16, min uint16, max uint16) {
	for i := range a {
		if a[i] < min {
			a[i] = min
			continue
		}
		if a[i] > max {
			a[i] = max
		}
	}
}

func ClampU32(a []uint32, min uint32, max uint32) {
	for i := range a {
		if a[i] < min {
			a[i] = min
			continue
		}
		if a[i] > max {
			a[i] = max
		}
	}
}

func ClampU64(a []uint64, min uint64, max uint64) {
	for i := range a {
		if a[i] < min {
			a[i] = min
			continue
		}
		if a[i] > max {
			a[i] = max
		}
	}
}

func ClampF32(a []float32, min float32, max float32) {
	for i := range a {
		if a[i] < min || math32.IsInf(a[i], -1) {
			a[i] = min
			continue
		}
		if a[i] > max || math32.IsInf(a[i], 1) {
			a[i] = max
		}
	}
}

func ClampF64(a []float64, min float64, max float64) {
	for i := range a {
		if a[i] < min || math.IsInf(a[i], -1) {
			a[i] = min
			continue
		}
		if a[i] > max || math.IsInf(a[i], 1) {
			a[i] = max
		}
	}
}

func ClampIterI(a []int, ait Iterator, min int, max int) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < min {
				a[i] = min
				continue
			}
			if a[i] > max {
				a[i] = max
			}
		}
	}
	return
}

func ClampIterI8(a []int8, ait Iterator, min int8, max int8) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < min {
				a[i] = min
				continue
			}
			if a[i] > max {
				a[i] = max
			}
		}
	}
	return
}

func ClampIterI16(a []int16, ait Iterator, min int16, max int16) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < min {
				a[i] = min
				continue
			}
			if a[i] > max {
				a[i] = max
			}
		}
	}
	return
}

func ClampIterI32(a []int32, ait Iterator, min int32, max int32) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < min {
				a[i] = min
				continue
			}
			if a[i] > max {
				a[i] = max
			}
		}
	}
	return
}

func ClampIterI64(a []int64, ait Iterator, min int64, max int64) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < min {
				a[i] = min
				continue
			}
			if a[i] > max {
				a[i] = max
			}
		}
	}
	return
}

func ClampIterU(a []uint, ait Iterator, min uint, max uint) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < min {
				a[i] = min
				continue
			}
			if a[i] > max {
				a[i] = max
			}
		}
	}
	return
}

func ClampIterU8(a []uint8, ait Iterator, min uint8, max uint8) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < min {
				a[i] = min
				continue
			}
			if a[i] > max {
				a[i] = max
			}
		}
	}
	return
}

func ClampIterU16(a []uint16, ait Iterator, min uint16, max uint16) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < min {
				a[i] = min
				continue
			}
			if a[i] > max {
				a[i] = max
			}
		}
	}
	return
}

func ClampIterU32(a []uint32, ait Iterator, min uint32, max uint32) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < min {
				a[i] = min
				continue
			}
			if a[i] > max {
				a[i] = max
			}
		}
	}
	return
}

func ClampIterU64(a []uint64, ait Iterator, min uint64, max uint64) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < min {
				a[i] = min
				continue
			}
			if a[i] > max {
				a[i] = max
			}
		}
	}
	return
}

func ClampIterF32(a []float32, ait Iterator, min float32, max float32) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < min || math32.IsInf(a[i], -1) {
				a[i] = min
				continue
			}
			if a[i] > max || math32.IsInf(a[i], 1) {
				a[i] = max
			}
		}
	}
	return
}

func ClampIterF64(a []float64, ait Iterator, min float64, max float64) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < min || math.IsInf(a[i], -1) {
				a[i] = min
				continue
			}
			if a[i] > max || math.IsInf(a[i], 1) {
				a[i] = max
			}
		}
	}
	return
}
