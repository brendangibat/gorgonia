package execution

import "unsafe"

/*
GENERATED FILE. DO NOT EDIT
*/

func MapB(fn func(bool) bool, a []bool) {
	for i := range a {
		a[i] = fn(a[i])
	}
	return
}

func MapI(fn func(int) int, a []int) {
	for i := range a {
		a[i] = fn(a[i])
	}
	return
}

func MapI8(fn func(int8) int8, a []int8) {
	for i := range a {
		a[i] = fn(a[i])
	}
	return
}

func MapI16(fn func(int16) int16, a []int16) {
	for i := range a {
		a[i] = fn(a[i])
	}
	return
}

func MapI32(fn func(int32) int32, a []int32) {
	for i := range a {
		a[i] = fn(a[i])
	}
	return
}

func MapI64(fn func(int64) int64, a []int64) {
	for i := range a {
		a[i] = fn(a[i])
	}
	return
}

func MapU(fn func(uint) uint, a []uint) {
	for i := range a {
		a[i] = fn(a[i])
	}
	return
}

func MapU8(fn func(uint8) uint8, a []uint8) {
	for i := range a {
		a[i] = fn(a[i])
	}
	return
}

func MapU16(fn func(uint16) uint16, a []uint16) {
	for i := range a {
		a[i] = fn(a[i])
	}
	return
}

func MapU32(fn func(uint32) uint32, a []uint32) {
	for i := range a {
		a[i] = fn(a[i])
	}
	return
}

func MapU64(fn func(uint64) uint64, a []uint64) {
	for i := range a {
		a[i] = fn(a[i])
	}
	return
}

func MapUintptr(fn func(uintptr) uintptr, a []uintptr) {
	for i := range a {
		a[i] = fn(a[i])
	}
	return
}

func MapF32(fn func(float32) float32, a []float32) {
	for i := range a {
		a[i] = fn(a[i])
	}
	return
}

func MapF64(fn func(float64) float64, a []float64) {
	for i := range a {
		a[i] = fn(a[i])
	}
	return
}

func MapC64(fn func(complex64) complex64, a []complex64) {
	for i := range a {
		a[i] = fn(a[i])
	}
	return
}

func MapC128(fn func(complex128) complex128, a []complex128) {
	for i := range a {
		a[i] = fn(a[i])
	}
	return
}

func MapStr(fn func(string) string, a []string) {
	for i := range a {
		a[i] = fn(a[i])
	}
	return
}

func MapUnsafePointer(fn func(unsafe.Pointer) unsafe.Pointer, a []unsafe.Pointer) {
	for i := range a {
		a[i] = fn(a[i])
	}
	return
}

func MapErrB(fn func(bool) (bool, error), a []bool) (err error) {
	for i := range a {
		if a[i], err = fn(a[i]); handleNoOp(err) != nil {
			return
		}
	}
	return
}

func MapErrI(fn func(int) (int, error), a []int) (err error) {
	for i := range a {
		if a[i], err = fn(a[i]); handleNoOp(err) != nil {
			return
		}
	}
	return
}

func MapErrI8(fn func(int8) (int8, error), a []int8) (err error) {
	for i := range a {
		if a[i], err = fn(a[i]); handleNoOp(err) != nil {
			return
		}
	}
	return
}

func MapErrI16(fn func(int16) (int16, error), a []int16) (err error) {
	for i := range a {
		if a[i], err = fn(a[i]); handleNoOp(err) != nil {
			return
		}
	}
	return
}

func MapErrI32(fn func(int32) (int32, error), a []int32) (err error) {
	for i := range a {
		if a[i], err = fn(a[i]); handleNoOp(err) != nil {
			return
		}
	}
	return
}

func MapErrI64(fn func(int64) (int64, error), a []int64) (err error) {
	for i := range a {
		if a[i], err = fn(a[i]); handleNoOp(err) != nil {
			return
		}
	}
	return
}

func MapErrU(fn func(uint) (uint, error), a []uint) (err error) {
	for i := range a {
		if a[i], err = fn(a[i]); handleNoOp(err) != nil {
			return
		}
	}
	return
}

func MapErrU8(fn func(uint8) (uint8, error), a []uint8) (err error) {
	for i := range a {
		if a[i], err = fn(a[i]); handleNoOp(err) != nil {
			return
		}
	}
	return
}

func MapErrU16(fn func(uint16) (uint16, error), a []uint16) (err error) {
	for i := range a {
		if a[i], err = fn(a[i]); handleNoOp(err) != nil {
			return
		}
	}
	return
}

func MapErrU32(fn func(uint32) (uint32, error), a []uint32) (err error) {
	for i := range a {
		if a[i], err = fn(a[i]); handleNoOp(err) != nil {
			return
		}
	}
	return
}

func MapErrU64(fn func(uint64) (uint64, error), a []uint64) (err error) {
	for i := range a {
		if a[i], err = fn(a[i]); handleNoOp(err) != nil {
			return
		}
	}
	return
}

func MapErrUintptr(fn func(uintptr) (uintptr, error), a []uintptr) (err error) {
	for i := range a {
		if a[i], err = fn(a[i]); handleNoOp(err) != nil {
			return
		}
	}
	return
}

func MapErrF32(fn func(float32) (float32, error), a []float32) (err error) {
	for i := range a {
		if a[i], err = fn(a[i]); handleNoOp(err) != nil {
			return
		}
	}
	return
}

func MapErrF64(fn func(float64) (float64, error), a []float64) (err error) {
	for i := range a {
		if a[i], err = fn(a[i]); handleNoOp(err) != nil {
			return
		}
	}
	return
}

func MapErrC64(fn func(complex64) (complex64, error), a []complex64) (err error) {
	for i := range a {
		if a[i], err = fn(a[i]); handleNoOp(err) != nil {
			return
		}
	}
	return
}

func MapErrC128(fn func(complex128) (complex128, error), a []complex128) (err error) {
	for i := range a {
		if a[i], err = fn(a[i]); handleNoOp(err) != nil {
			return
		}
	}
	return
}

func MapErrStr(fn func(string) (string, error), a []string) (err error) {
	for i := range a {
		if a[i], err = fn(a[i]); handleNoOp(err) != nil {
			return
		}
	}
	return
}

func MapErrUnsafePointer(fn func(unsafe.Pointer) (unsafe.Pointer, error), a []unsafe.Pointer) (err error) {
	for i := range a {
		if a[i], err = fn(a[i]); handleNoOp(err) != nil {
			return
		}
	}
	return
}

func MapIterB(fn func(bool) bool, a []bool, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = fn(a[i])
		}
	}
	return
}

func MapIterI(fn func(int) int, a []int, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = fn(a[i])
		}
	}
	return
}

func MapIterI8(fn func(int8) int8, a []int8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = fn(a[i])
		}
	}
	return
}

func MapIterI16(fn func(int16) int16, a []int16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = fn(a[i])
		}
	}
	return
}

func MapIterI32(fn func(int32) int32, a []int32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = fn(a[i])
		}
	}
	return
}

func MapIterI64(fn func(int64) int64, a []int64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = fn(a[i])
		}
	}
	return
}

func MapIterU(fn func(uint) uint, a []uint, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = fn(a[i])
		}
	}
	return
}

func MapIterU8(fn func(uint8) uint8, a []uint8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = fn(a[i])
		}
	}
	return
}

func MapIterU16(fn func(uint16) uint16, a []uint16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = fn(a[i])
		}
	}
	return
}

func MapIterU32(fn func(uint32) uint32, a []uint32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = fn(a[i])
		}
	}
	return
}

func MapIterU64(fn func(uint64) uint64, a []uint64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = fn(a[i])
		}
	}
	return
}

func MapIterUintptr(fn func(uintptr) uintptr, a []uintptr, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = fn(a[i])
		}
	}
	return
}

func MapIterF32(fn func(float32) float32, a []float32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = fn(a[i])
		}
	}
	return
}

func MapIterF64(fn func(float64) float64, a []float64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = fn(a[i])
		}
	}
	return
}

func MapIterC64(fn func(complex64) complex64, a []complex64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = fn(a[i])
		}
	}
	return
}

func MapIterC128(fn func(complex128) complex128, a []complex128, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = fn(a[i])
		}
	}
	return
}

func MapIterStr(fn func(string) string, a []string, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = fn(a[i])
		}
	}
	return
}

func MapIterUnsafePointer(fn func(unsafe.Pointer) unsafe.Pointer, a []unsafe.Pointer, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = fn(a[i])
		}
	}
	return
}

func MapIterErrB(fn func(bool) (bool, error), a []bool, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i], err = fn(a[i]); handleNoOp(err) != nil {
				return
			}
		}
	}
	return
}

func MapIterErrI(fn func(int) (int, error), a []int, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i], err = fn(a[i]); handleNoOp(err) != nil {
				return
			}
		}
	}
	return
}

func MapIterErrI8(fn func(int8) (int8, error), a []int8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i], err = fn(a[i]); handleNoOp(err) != nil {
				return
			}
		}
	}
	return
}

func MapIterErrI16(fn func(int16) (int16, error), a []int16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i], err = fn(a[i]); handleNoOp(err) != nil {
				return
			}
		}
	}
	return
}

func MapIterErrI32(fn func(int32) (int32, error), a []int32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i], err = fn(a[i]); handleNoOp(err) != nil {
				return
			}
		}
	}
	return
}

func MapIterErrI64(fn func(int64) (int64, error), a []int64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i], err = fn(a[i]); handleNoOp(err) != nil {
				return
			}
		}
	}
	return
}

func MapIterErrU(fn func(uint) (uint, error), a []uint, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i], err = fn(a[i]); handleNoOp(err) != nil {
				return
			}
		}
	}
	return
}

func MapIterErrU8(fn func(uint8) (uint8, error), a []uint8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i], err = fn(a[i]); handleNoOp(err) != nil {
				return
			}
		}
	}
	return
}

func MapIterErrU16(fn func(uint16) (uint16, error), a []uint16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i], err = fn(a[i]); handleNoOp(err) != nil {
				return
			}
		}
	}
	return
}

func MapIterErrU32(fn func(uint32) (uint32, error), a []uint32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i], err = fn(a[i]); handleNoOp(err) != nil {
				return
			}
		}
	}
	return
}

func MapIterErrU64(fn func(uint64) (uint64, error), a []uint64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i], err = fn(a[i]); handleNoOp(err) != nil {
				return
			}
		}
	}
	return
}

func MapIterErrUintptr(fn func(uintptr) (uintptr, error), a []uintptr, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i], err = fn(a[i]); handleNoOp(err) != nil {
				return
			}
		}
	}
	return
}

func MapIterErrF32(fn func(float32) (float32, error), a []float32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i], err = fn(a[i]); handleNoOp(err) != nil {
				return
			}
		}
	}
	return
}

func MapIterErrF64(fn func(float64) (float64, error), a []float64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i], err = fn(a[i]); handleNoOp(err) != nil {
				return
			}
		}
	}
	return
}

func MapIterErrC64(fn func(complex64) (complex64, error), a []complex64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i], err = fn(a[i]); handleNoOp(err) != nil {
				return
			}
		}
	}
	return
}

func MapIterErrC128(fn func(complex128) (complex128, error), a []complex128, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i], err = fn(a[i]); handleNoOp(err) != nil {
				return
			}
		}
	}
	return
}

func MapIterErrStr(fn func(string) (string, error), a []string, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i], err = fn(a[i]); handleNoOp(err) != nil {
				return
			}
		}
	}
	return
}

func MapIterErrUnsafePointer(fn func(unsafe.Pointer) (unsafe.Pointer, error), a []unsafe.Pointer, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i], err = fn(a[i]); handleNoOp(err) != nil {
				return
			}
		}
	}
	return
}

func MapIncrI(fn func(int) int, a []int) {
	for i := range a {
		a[i] += fn(a[i])
	}
	return
}

func MapIncrI8(fn func(int8) int8, a []int8) {
	for i := range a {
		a[i] += fn(a[i])
	}
	return
}

func MapIncrI16(fn func(int16) int16, a []int16) {
	for i := range a {
		a[i] += fn(a[i])
	}
	return
}

func MapIncrI32(fn func(int32) int32, a []int32) {
	for i := range a {
		a[i] += fn(a[i])
	}
	return
}

func MapIncrI64(fn func(int64) int64, a []int64) {
	for i := range a {
		a[i] += fn(a[i])
	}
	return
}

func MapIncrU(fn func(uint) uint, a []uint) {
	for i := range a {
		a[i] += fn(a[i])
	}
	return
}

func MapIncrU8(fn func(uint8) uint8, a []uint8) {
	for i := range a {
		a[i] += fn(a[i])
	}
	return
}

func MapIncrU16(fn func(uint16) uint16, a []uint16) {
	for i := range a {
		a[i] += fn(a[i])
	}
	return
}

func MapIncrU32(fn func(uint32) uint32, a []uint32) {
	for i := range a {
		a[i] += fn(a[i])
	}
	return
}

func MapIncrU64(fn func(uint64) uint64, a []uint64) {
	for i := range a {
		a[i] += fn(a[i])
	}
	return
}

func MapIncrF32(fn func(float32) float32, a []float32) {
	for i := range a {
		a[i] += fn(a[i])
	}
	return
}

func MapIncrF64(fn func(float64) float64, a []float64) {
	for i := range a {
		a[i] += fn(a[i])
	}
	return
}

func MapIncrC64(fn func(complex64) complex64, a []complex64) {
	for i := range a {
		a[i] += fn(a[i])
	}
	return
}

func MapIncrC128(fn func(complex128) complex128, a []complex128) {
	for i := range a {
		a[i] += fn(a[i])
	}
	return
}

func MapIncrStr(fn func(string) string, a []string) {
	for i := range a {
		a[i] += fn(a[i])
	}
	return
}

func MapIncrErrI(fn func(int) (int, error), a []int) (err error) {
	for i := range a {
		var x int
		if x, err = fn(a[i]); err != nil {
			if err = handleNoOp(err); err != nil {
				return
			}
		}
		a[i] = x
	}
	return
}

func MapIncrErrI8(fn func(int8) (int8, error), a []int8) (err error) {
	for i := range a {
		var x int8
		if x, err = fn(a[i]); err != nil {
			if err = handleNoOp(err); err != nil {
				return
			}
		}
		a[i] = x
	}
	return
}

func MapIncrErrI16(fn func(int16) (int16, error), a []int16) (err error) {
	for i := range a {
		var x int16
		if x, err = fn(a[i]); err != nil {
			if err = handleNoOp(err); err != nil {
				return
			}
		}
		a[i] = x
	}
	return
}

func MapIncrErrI32(fn func(int32) (int32, error), a []int32) (err error) {
	for i := range a {
		var x int32
		if x, err = fn(a[i]); err != nil {
			if err = handleNoOp(err); err != nil {
				return
			}
		}
		a[i] = x
	}
	return
}

func MapIncrErrI64(fn func(int64) (int64, error), a []int64) (err error) {
	for i := range a {
		var x int64
		if x, err = fn(a[i]); err != nil {
			if err = handleNoOp(err); err != nil {
				return
			}
		}
		a[i] = x
	}
	return
}

func MapIncrErrU(fn func(uint) (uint, error), a []uint) (err error) {
	for i := range a {
		var x uint
		if x, err = fn(a[i]); err != nil {
			if err = handleNoOp(err); err != nil {
				return
			}
		}
		a[i] = x
	}
	return
}

func MapIncrErrU8(fn func(uint8) (uint8, error), a []uint8) (err error) {
	for i := range a {
		var x uint8
		if x, err = fn(a[i]); err != nil {
			if err = handleNoOp(err); err != nil {
				return
			}
		}
		a[i] = x
	}
	return
}

func MapIncrErrU16(fn func(uint16) (uint16, error), a []uint16) (err error) {
	for i := range a {
		var x uint16
		if x, err = fn(a[i]); err != nil {
			if err = handleNoOp(err); err != nil {
				return
			}
		}
		a[i] = x
	}
	return
}

func MapIncrErrU32(fn func(uint32) (uint32, error), a []uint32) (err error) {
	for i := range a {
		var x uint32
		if x, err = fn(a[i]); err != nil {
			if err = handleNoOp(err); err != nil {
				return
			}
		}
		a[i] = x
	}
	return
}

func MapIncrErrU64(fn func(uint64) (uint64, error), a []uint64) (err error) {
	for i := range a {
		var x uint64
		if x, err = fn(a[i]); err != nil {
			if err = handleNoOp(err); err != nil {
				return
			}
		}
		a[i] = x
	}
	return
}

func MapIncrErrF32(fn func(float32) (float32, error), a []float32) (err error) {
	for i := range a {
		var x float32
		if x, err = fn(a[i]); err != nil {
			if err = handleNoOp(err); err != nil {
				return
			}
		}
		a[i] = x
	}
	return
}

func MapIncrErrF64(fn func(float64) (float64, error), a []float64) (err error) {
	for i := range a {
		var x float64
		if x, err = fn(a[i]); err != nil {
			if err = handleNoOp(err); err != nil {
				return
			}
		}
		a[i] = x
	}
	return
}

func MapIncrErrC64(fn func(complex64) (complex64, error), a []complex64) (err error) {
	for i := range a {
		var x complex64
		if x, err = fn(a[i]); err != nil {
			if err = handleNoOp(err); err != nil {
				return
			}
		}
		a[i] = x
	}
	return
}

func MapIncrErrC128(fn func(complex128) (complex128, error), a []complex128) (err error) {
	for i := range a {
		var x complex128
		if x, err = fn(a[i]); err != nil {
			if err = handleNoOp(err); err != nil {
				return
			}
		}
		a[i] = x
	}
	return
}

func MapIncrErrStr(fn func(string) (string, error), a []string) (err error) {
	for i := range a {
		var x string
		if x, err = fn(a[i]); err != nil {
			if err = handleNoOp(err); err != nil {
				return
			}
		}
		a[i] = x
	}
	return
}

func MapIterIncrI(fn func(int) int, a []int, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] += fn(a[i])
		}
	}
	return
}

func MapIterIncrI8(fn func(int8) int8, a []int8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] += fn(a[i])
		}
	}
	return
}

func MapIterIncrI16(fn func(int16) int16, a []int16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] += fn(a[i])
		}
	}
	return
}

func MapIterIncrI32(fn func(int32) int32, a []int32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] += fn(a[i])
		}
	}
	return
}

func MapIterIncrI64(fn func(int64) int64, a []int64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] += fn(a[i])
		}
	}
	return
}

func MapIterIncrU(fn func(uint) uint, a []uint, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] += fn(a[i])
		}
	}
	return
}

func MapIterIncrU8(fn func(uint8) uint8, a []uint8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] += fn(a[i])
		}
	}
	return
}

func MapIterIncrU16(fn func(uint16) uint16, a []uint16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] += fn(a[i])
		}
	}
	return
}

func MapIterIncrU32(fn func(uint32) uint32, a []uint32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] += fn(a[i])
		}
	}
	return
}

func MapIterIncrU64(fn func(uint64) uint64, a []uint64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] += fn(a[i])
		}
	}
	return
}

func MapIterIncrF32(fn func(float32) float32, a []float32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] += fn(a[i])
		}
	}
	return
}

func MapIterIncrF64(fn func(float64) float64, a []float64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] += fn(a[i])
		}
	}
	return
}

func MapIterIncrC64(fn func(complex64) complex64, a []complex64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] += fn(a[i])
		}
	}
	return
}

func MapIterIncrC128(fn func(complex128) complex128, a []complex128, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] += fn(a[i])
		}
	}
	return
}

func MapIterIncrStr(fn func(string) string, a []string, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] += fn(a[i])
		}
	}
	return
}

func MapIterIncrErrI(fn func(int) (int, error), a []int, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			var x int
			if x, err = fn(a[i]); err != nil {
				if err = handleNoOp(err); err != nil {
					return
				}
			}
			a[i] = x
		}
	}
	return
}

func MapIterIncrErrI8(fn func(int8) (int8, error), a []int8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			var x int8
			if x, err = fn(a[i]); err != nil {
				if err = handleNoOp(err); err != nil {
					return
				}
			}
			a[i] = x
		}
	}
	return
}

func MapIterIncrErrI16(fn func(int16) (int16, error), a []int16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			var x int16
			if x, err = fn(a[i]); err != nil {
				if err = handleNoOp(err); err != nil {
					return
				}
			}
			a[i] = x
		}
	}
	return
}

func MapIterIncrErrI32(fn func(int32) (int32, error), a []int32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			var x int32
			if x, err = fn(a[i]); err != nil {
				if err = handleNoOp(err); err != nil {
					return
				}
			}
			a[i] = x
		}
	}
	return
}

func MapIterIncrErrI64(fn func(int64) (int64, error), a []int64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			var x int64
			if x, err = fn(a[i]); err != nil {
				if err = handleNoOp(err); err != nil {
					return
				}
			}
			a[i] = x
		}
	}
	return
}

func MapIterIncrErrU(fn func(uint) (uint, error), a []uint, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			var x uint
			if x, err = fn(a[i]); err != nil {
				if err = handleNoOp(err); err != nil {
					return
				}
			}
			a[i] = x
		}
	}
	return
}

func MapIterIncrErrU8(fn func(uint8) (uint8, error), a []uint8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			var x uint8
			if x, err = fn(a[i]); err != nil {
				if err = handleNoOp(err); err != nil {
					return
				}
			}
			a[i] = x
		}
	}
	return
}

func MapIterIncrErrU16(fn func(uint16) (uint16, error), a []uint16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			var x uint16
			if x, err = fn(a[i]); err != nil {
				if err = handleNoOp(err); err != nil {
					return
				}
			}
			a[i] = x
		}
	}
	return
}

func MapIterIncrErrU32(fn func(uint32) (uint32, error), a []uint32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			var x uint32
			if x, err = fn(a[i]); err != nil {
				if err = handleNoOp(err); err != nil {
					return
				}
			}
			a[i] = x
		}
	}
	return
}

func MapIterIncrErrU64(fn func(uint64) (uint64, error), a []uint64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			var x uint64
			if x, err = fn(a[i]); err != nil {
				if err = handleNoOp(err); err != nil {
					return
				}
			}
			a[i] = x
		}
	}
	return
}

func MapIterIncrErrF32(fn func(float32) (float32, error), a []float32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			var x float32
			if x, err = fn(a[i]); err != nil {
				if err = handleNoOp(err); err != nil {
					return
				}
			}
			a[i] = x
		}
	}
	return
}

func MapIterIncrErrF64(fn func(float64) (float64, error), a []float64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			var x float64
			if x, err = fn(a[i]); err != nil {
				if err = handleNoOp(err); err != nil {
					return
				}
			}
			a[i] = x
		}
	}
	return
}

func MapIterIncrErrC64(fn func(complex64) (complex64, error), a []complex64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			var x complex64
			if x, err = fn(a[i]); err != nil {
				if err = handleNoOp(err); err != nil {
					return
				}
			}
			a[i] = x
		}
	}
	return
}

func MapIterIncrErrC128(fn func(complex128) (complex128, error), a []complex128, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			var x complex128
			if x, err = fn(a[i]); err != nil {
				if err = handleNoOp(err); err != nil {
					return
				}
			}
			a[i] = x
		}
	}
	return
}

func MapIterIncrErrStr(fn func(string) (string, error), a []string, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			var x string
			if x, err = fn(a[i]); err != nil {
				if err = handleNoOp(err); err != nil {
					return
				}
			}
			a[i] = x
		}
	}
	return
}
