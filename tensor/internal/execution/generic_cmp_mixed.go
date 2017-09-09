package execution

import "unsafe"

/*
GENERATED FILE. DO NOT EDIT
*/

func GtSVI(a int, b []int, retVal []bool) {
	for i := range retVal {
		retVal[i] = a > b[i]
	}
}

func GtSVI8(a int8, b []int8, retVal []bool) {
	for i := range retVal {
		retVal[i] = a > b[i]
	}
}

func GtSVI16(a int16, b []int16, retVal []bool) {
	for i := range retVal {
		retVal[i] = a > b[i]
	}
}

func GtSVI32(a int32, b []int32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a > b[i]
	}
}

func GtSVI64(a int64, b []int64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a > b[i]
	}
}

func GtSVU(a uint, b []uint, retVal []bool) {
	for i := range retVal {
		retVal[i] = a > b[i]
	}
}

func GtSVU8(a uint8, b []uint8, retVal []bool) {
	for i := range retVal {
		retVal[i] = a > b[i]
	}
}

func GtSVU16(a uint16, b []uint16, retVal []bool) {
	for i := range retVal {
		retVal[i] = a > b[i]
	}
}

func GtSVU32(a uint32, b []uint32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a > b[i]
	}
}

func GtSVU64(a uint64, b []uint64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a > b[i]
	}
}

func GtSVF32(a float32, b []float32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a > b[i]
	}
}

func GtSVF64(a float64, b []float64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a > b[i]
	}
}

func GtSVStr(a string, b []string, retVal []bool) {
	for i := range retVal {
		retVal[i] = a > b[i]
	}
}

func GteSVI(a int, b []int, retVal []bool) {
	for i := range retVal {
		retVal[i] = a >= b[i]
	}
}

func GteSVI8(a int8, b []int8, retVal []bool) {
	for i := range retVal {
		retVal[i] = a >= b[i]
	}
}

func GteSVI16(a int16, b []int16, retVal []bool) {
	for i := range retVal {
		retVal[i] = a >= b[i]
	}
}

func GteSVI32(a int32, b []int32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a >= b[i]
	}
}

func GteSVI64(a int64, b []int64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a >= b[i]
	}
}

func GteSVU(a uint, b []uint, retVal []bool) {
	for i := range retVal {
		retVal[i] = a >= b[i]
	}
}

func GteSVU8(a uint8, b []uint8, retVal []bool) {
	for i := range retVal {
		retVal[i] = a >= b[i]
	}
}

func GteSVU16(a uint16, b []uint16, retVal []bool) {
	for i := range retVal {
		retVal[i] = a >= b[i]
	}
}

func GteSVU32(a uint32, b []uint32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a >= b[i]
	}
}

func GteSVU64(a uint64, b []uint64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a >= b[i]
	}
}

func GteSVF32(a float32, b []float32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a >= b[i]
	}
}

func GteSVF64(a float64, b []float64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a >= b[i]
	}
}

func GteSVStr(a string, b []string, retVal []bool) {
	for i := range retVal {
		retVal[i] = a >= b[i]
	}
}

func LtSVI(a int, b []int, retVal []bool) {
	for i := range retVal {
		retVal[i] = a < b[i]
	}
}

func LtSVI8(a int8, b []int8, retVal []bool) {
	for i := range retVal {
		retVal[i] = a < b[i]
	}
}

func LtSVI16(a int16, b []int16, retVal []bool) {
	for i := range retVal {
		retVal[i] = a < b[i]
	}
}

func LtSVI32(a int32, b []int32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a < b[i]
	}
}

func LtSVI64(a int64, b []int64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a < b[i]
	}
}

func LtSVU(a uint, b []uint, retVal []bool) {
	for i := range retVal {
		retVal[i] = a < b[i]
	}
}

func LtSVU8(a uint8, b []uint8, retVal []bool) {
	for i := range retVal {
		retVal[i] = a < b[i]
	}
}

func LtSVU16(a uint16, b []uint16, retVal []bool) {
	for i := range retVal {
		retVal[i] = a < b[i]
	}
}

func LtSVU32(a uint32, b []uint32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a < b[i]
	}
}

func LtSVU64(a uint64, b []uint64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a < b[i]
	}
}

func LtSVF32(a float32, b []float32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a < b[i]
	}
}

func LtSVF64(a float64, b []float64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a < b[i]
	}
}

func LtSVStr(a string, b []string, retVal []bool) {
	for i := range retVal {
		retVal[i] = a < b[i]
	}
}

func LteSVI(a int, b []int, retVal []bool) {
	for i := range retVal {
		retVal[i] = a <= b[i]
	}
}

func LteSVI8(a int8, b []int8, retVal []bool) {
	for i := range retVal {
		retVal[i] = a <= b[i]
	}
}

func LteSVI16(a int16, b []int16, retVal []bool) {
	for i := range retVal {
		retVal[i] = a <= b[i]
	}
}

func LteSVI32(a int32, b []int32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a <= b[i]
	}
}

func LteSVI64(a int64, b []int64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a <= b[i]
	}
}

func LteSVU(a uint, b []uint, retVal []bool) {
	for i := range retVal {
		retVal[i] = a <= b[i]
	}
}

func LteSVU8(a uint8, b []uint8, retVal []bool) {
	for i := range retVal {
		retVal[i] = a <= b[i]
	}
}

func LteSVU16(a uint16, b []uint16, retVal []bool) {
	for i := range retVal {
		retVal[i] = a <= b[i]
	}
}

func LteSVU32(a uint32, b []uint32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a <= b[i]
	}
}

func LteSVU64(a uint64, b []uint64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a <= b[i]
	}
}

func LteSVF32(a float32, b []float32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a <= b[i]
	}
}

func LteSVF64(a float64, b []float64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a <= b[i]
	}
}

func LteSVStr(a string, b []string, retVal []bool) {
	for i := range retVal {
		retVal[i] = a <= b[i]
	}
}

func EqSVB(a bool, b []bool, retVal []bool) {
	for i := range retVal {
		retVal[i] = a == b[i]
	}
}

func EqSVI(a int, b []int, retVal []bool) {
	for i := range retVal {
		retVal[i] = a == b[i]
	}
}

func EqSVI8(a int8, b []int8, retVal []bool) {
	for i := range retVal {
		retVal[i] = a == b[i]
	}
}

func EqSVI16(a int16, b []int16, retVal []bool) {
	for i := range retVal {
		retVal[i] = a == b[i]
	}
}

func EqSVI32(a int32, b []int32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a == b[i]
	}
}

func EqSVI64(a int64, b []int64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a == b[i]
	}
}

func EqSVU(a uint, b []uint, retVal []bool) {
	for i := range retVal {
		retVal[i] = a == b[i]
	}
}

func EqSVU8(a uint8, b []uint8, retVal []bool) {
	for i := range retVal {
		retVal[i] = a == b[i]
	}
}

func EqSVU16(a uint16, b []uint16, retVal []bool) {
	for i := range retVal {
		retVal[i] = a == b[i]
	}
}

func EqSVU32(a uint32, b []uint32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a == b[i]
	}
}

func EqSVU64(a uint64, b []uint64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a == b[i]
	}
}

func EqSVUintptr(a uintptr, b []uintptr, retVal []bool) {
	for i := range retVal {
		retVal[i] = a == b[i]
	}
}

func EqSVF32(a float32, b []float32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a == b[i]
	}
}

func EqSVF64(a float64, b []float64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a == b[i]
	}
}

func EqSVC64(a complex64, b []complex64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a == b[i]
	}
}

func EqSVC128(a complex128, b []complex128, retVal []bool) {
	for i := range retVal {
		retVal[i] = a == b[i]
	}
}

func EqSVStr(a string, b []string, retVal []bool) {
	for i := range retVal {
		retVal[i] = a == b[i]
	}
}

func EqSVUnsafePointer(a unsafe.Pointer, b []unsafe.Pointer, retVal []bool) {
	for i := range retVal {
		retVal[i] = a == b[i]
	}
}

func NeSVB(a bool, b []bool, retVal []bool) {
	for i := range retVal {
		retVal[i] = a != b[i]
	}
}

func NeSVI(a int, b []int, retVal []bool) {
	for i := range retVal {
		retVal[i] = a != b[i]
	}
}

func NeSVI8(a int8, b []int8, retVal []bool) {
	for i := range retVal {
		retVal[i] = a != b[i]
	}
}

func NeSVI16(a int16, b []int16, retVal []bool) {
	for i := range retVal {
		retVal[i] = a != b[i]
	}
}

func NeSVI32(a int32, b []int32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a != b[i]
	}
}

func NeSVI64(a int64, b []int64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a != b[i]
	}
}

func NeSVU(a uint, b []uint, retVal []bool) {
	for i := range retVal {
		retVal[i] = a != b[i]
	}
}

func NeSVU8(a uint8, b []uint8, retVal []bool) {
	for i := range retVal {
		retVal[i] = a != b[i]
	}
}

func NeSVU16(a uint16, b []uint16, retVal []bool) {
	for i := range retVal {
		retVal[i] = a != b[i]
	}
}

func NeSVU32(a uint32, b []uint32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a != b[i]
	}
}

func NeSVU64(a uint64, b []uint64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a != b[i]
	}
}

func NeSVUintptr(a uintptr, b []uintptr, retVal []bool) {
	for i := range retVal {
		retVal[i] = a != b[i]
	}
}

func NeSVF32(a float32, b []float32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a != b[i]
	}
}

func NeSVF64(a float64, b []float64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a != b[i]
	}
}

func NeSVC64(a complex64, b []complex64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a != b[i]
	}
}

func NeSVC128(a complex128, b []complex128, retVal []bool) {
	for i := range retVal {
		retVal[i] = a != b[i]
	}
}

func NeSVStr(a string, b []string, retVal []bool) {
	for i := range retVal {
		retVal[i] = a != b[i]
	}
}

func NeSVUnsafePointer(a unsafe.Pointer, b []unsafe.Pointer, retVal []bool) {
	for i := range retVal {
		retVal[i] = a != b[i]
	}
}

func GtSameSVI(a int, b []int) {
	for i := range b {
		if a > b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func GtSameSVI8(a int8, b []int8) {
	for i := range b {
		if a > b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func GtSameSVI16(a int16, b []int16) {
	for i := range b {
		if a > b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func GtSameSVI32(a int32, b []int32) {
	for i := range b {
		if a > b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func GtSameSVI64(a int64, b []int64) {
	for i := range b {
		if a > b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func GtSameSVU(a uint, b []uint) {
	for i := range b {
		if a > b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func GtSameSVU8(a uint8, b []uint8) {
	for i := range b {
		if a > b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func GtSameSVU16(a uint16, b []uint16) {
	for i := range b {
		if a > b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func GtSameSVU32(a uint32, b []uint32) {
	for i := range b {
		if a > b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func GtSameSVU64(a uint64, b []uint64) {
	for i := range b {
		if a > b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func GtSameSVF32(a float32, b []float32) {
	for i := range b {
		if a > b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func GtSameSVF64(a float64, b []float64) {
	for i := range b {
		if a > b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func GtSameSVStr(a string, b []string) {
	for i := range b {
		if a > b[i] {
			b[i] = "true"
		} else {
			b[i] = "false"
		}
	}
}

func GteSameSVI(a int, b []int) {
	for i := range b {
		if a >= b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func GteSameSVI8(a int8, b []int8) {
	for i := range b {
		if a >= b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func GteSameSVI16(a int16, b []int16) {
	for i := range b {
		if a >= b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func GteSameSVI32(a int32, b []int32) {
	for i := range b {
		if a >= b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func GteSameSVI64(a int64, b []int64) {
	for i := range b {
		if a >= b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func GteSameSVU(a uint, b []uint) {
	for i := range b {
		if a >= b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func GteSameSVU8(a uint8, b []uint8) {
	for i := range b {
		if a >= b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func GteSameSVU16(a uint16, b []uint16) {
	for i := range b {
		if a >= b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func GteSameSVU32(a uint32, b []uint32) {
	for i := range b {
		if a >= b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func GteSameSVU64(a uint64, b []uint64) {
	for i := range b {
		if a >= b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func GteSameSVF32(a float32, b []float32) {
	for i := range b {
		if a >= b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func GteSameSVF64(a float64, b []float64) {
	for i := range b {
		if a >= b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func GteSameSVStr(a string, b []string) {
	for i := range b {
		if a >= b[i] {
			b[i] = "true"
		} else {
			b[i] = "false"
		}
	}
}

func LtSameSVI(a int, b []int) {
	for i := range b {
		if a < b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func LtSameSVI8(a int8, b []int8) {
	for i := range b {
		if a < b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func LtSameSVI16(a int16, b []int16) {
	for i := range b {
		if a < b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func LtSameSVI32(a int32, b []int32) {
	for i := range b {
		if a < b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func LtSameSVI64(a int64, b []int64) {
	for i := range b {
		if a < b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func LtSameSVU(a uint, b []uint) {
	for i := range b {
		if a < b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func LtSameSVU8(a uint8, b []uint8) {
	for i := range b {
		if a < b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func LtSameSVU16(a uint16, b []uint16) {
	for i := range b {
		if a < b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func LtSameSVU32(a uint32, b []uint32) {
	for i := range b {
		if a < b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func LtSameSVU64(a uint64, b []uint64) {
	for i := range b {
		if a < b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func LtSameSVF32(a float32, b []float32) {
	for i := range b {
		if a < b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func LtSameSVF64(a float64, b []float64) {
	for i := range b {
		if a < b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func LtSameSVStr(a string, b []string) {
	for i := range b {
		if a < b[i] {
			b[i] = "true"
		} else {
			b[i] = "false"
		}
	}
}

func LteSameSVI(a int, b []int) {
	for i := range b {
		if a <= b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func LteSameSVI8(a int8, b []int8) {
	for i := range b {
		if a <= b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func LteSameSVI16(a int16, b []int16) {
	for i := range b {
		if a <= b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func LteSameSVI32(a int32, b []int32) {
	for i := range b {
		if a <= b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func LteSameSVI64(a int64, b []int64) {
	for i := range b {
		if a <= b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func LteSameSVU(a uint, b []uint) {
	for i := range b {
		if a <= b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func LteSameSVU8(a uint8, b []uint8) {
	for i := range b {
		if a <= b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func LteSameSVU16(a uint16, b []uint16) {
	for i := range b {
		if a <= b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func LteSameSVU32(a uint32, b []uint32) {
	for i := range b {
		if a <= b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func LteSameSVU64(a uint64, b []uint64) {
	for i := range b {
		if a <= b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func LteSameSVF32(a float32, b []float32) {
	for i := range b {
		if a <= b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func LteSameSVF64(a float64, b []float64) {
	for i := range b {
		if a <= b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func LteSameSVStr(a string, b []string) {
	for i := range b {
		if a <= b[i] {
			b[i] = "true"
		} else {
			b[i] = "false"
		}
	}
}

func EqSameSVB(a bool, b []bool) {
	for i := range b {
		if a == b[i] {
			b[i] = true
		} else {
			b[i] = false
		}
	}
}

func EqSameSVI(a int, b []int) {
	for i := range b {
		if a == b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func EqSameSVI8(a int8, b []int8) {
	for i := range b {
		if a == b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func EqSameSVI16(a int16, b []int16) {
	for i := range b {
		if a == b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func EqSameSVI32(a int32, b []int32) {
	for i := range b {
		if a == b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func EqSameSVI64(a int64, b []int64) {
	for i := range b {
		if a == b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func EqSameSVU(a uint, b []uint) {
	for i := range b {
		if a == b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func EqSameSVU8(a uint8, b []uint8) {
	for i := range b {
		if a == b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func EqSameSVU16(a uint16, b []uint16) {
	for i := range b {
		if a == b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func EqSameSVU32(a uint32, b []uint32) {
	for i := range b {
		if a == b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func EqSameSVU64(a uint64, b []uint64) {
	for i := range b {
		if a == b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func EqSameSVUintptr(a uintptr, b []uintptr) {
	for i := range b {
		if a == b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func EqSameSVF32(a float32, b []float32) {
	for i := range b {
		if a == b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func EqSameSVF64(a float64, b []float64) {
	for i := range b {
		if a == b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func EqSameSVC64(a complex64, b []complex64) {
	for i := range b {
		if a == b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func EqSameSVC128(a complex128, b []complex128) {
	for i := range b {
		if a == b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func EqSameSVStr(a string, b []string) {
	for i := range b {
		if a == b[i] {
			b[i] = "true"
		} else {
			b[i] = "false"
		}
	}
}

func NeSameSVB(a bool, b []bool) {
	for i := range b {
		if a != b[i] {
			b[i] = true
		} else {
			b[i] = false
		}
	}
}

func NeSameSVI(a int, b []int) {
	for i := range b {
		if a != b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func NeSameSVI8(a int8, b []int8) {
	for i := range b {
		if a != b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func NeSameSVI16(a int16, b []int16) {
	for i := range b {
		if a != b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func NeSameSVI32(a int32, b []int32) {
	for i := range b {
		if a != b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func NeSameSVI64(a int64, b []int64) {
	for i := range b {
		if a != b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func NeSameSVU(a uint, b []uint) {
	for i := range b {
		if a != b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func NeSameSVU8(a uint8, b []uint8) {
	for i := range b {
		if a != b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func NeSameSVU16(a uint16, b []uint16) {
	for i := range b {
		if a != b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func NeSameSVU32(a uint32, b []uint32) {
	for i := range b {
		if a != b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func NeSameSVU64(a uint64, b []uint64) {
	for i := range b {
		if a != b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func NeSameSVUintptr(a uintptr, b []uintptr) {
	for i := range b {
		if a != b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func NeSameSVF32(a float32, b []float32) {
	for i := range b {
		if a != b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func NeSameSVF64(a float64, b []float64) {
	for i := range b {
		if a != b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func NeSameSVC64(a complex64, b []complex64) {
	for i := range b {
		if a != b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func NeSameSVC128(a complex128, b []complex128) {
	for i := range b {
		if a != b[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
}

func NeSameSVStr(a string, b []string) {
	for i := range b {
		if a != b[i] {
			b[i] = "true"
		} else {
			b[i] = "false"
		}
	}
}

func GtIterSVI(a int, b []int, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a > b[i]
		}
	}
	return
}

func GtIterSVI8(a int8, b []int8, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a > b[i]
		}
	}
	return
}

func GtIterSVI16(a int16, b []int16, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a > b[i]
		}
	}
	return
}

func GtIterSVI32(a int32, b []int32, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a > b[i]
		}
	}
	return
}

func GtIterSVI64(a int64, b []int64, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a > b[i]
		}
	}
	return
}

func GtIterSVU(a uint, b []uint, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a > b[i]
		}
	}
	return
}

func GtIterSVU8(a uint8, b []uint8, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a > b[i]
		}
	}
	return
}

func GtIterSVU16(a uint16, b []uint16, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a > b[i]
		}
	}
	return
}

func GtIterSVU32(a uint32, b []uint32, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a > b[i]
		}
	}
	return
}

func GtIterSVU64(a uint64, b []uint64, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a > b[i]
		}
	}
	return
}

func GtIterSVF32(a float32, b []float32, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a > b[i]
		}
	}
	return
}

func GtIterSVF64(a float64, b []float64, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a > b[i]
		}
	}
	return
}

func GtIterSVStr(a string, b []string, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a > b[i]
		}
	}
	return
}

func GteIterSVI(a int, b []int, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a >= b[i]
		}
	}
	return
}

func GteIterSVI8(a int8, b []int8, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a >= b[i]
		}
	}
	return
}

func GteIterSVI16(a int16, b []int16, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a >= b[i]
		}
	}
	return
}

func GteIterSVI32(a int32, b []int32, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a >= b[i]
		}
	}
	return
}

func GteIterSVI64(a int64, b []int64, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a >= b[i]
		}
	}
	return
}

func GteIterSVU(a uint, b []uint, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a >= b[i]
		}
	}
	return
}

func GteIterSVU8(a uint8, b []uint8, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a >= b[i]
		}
	}
	return
}

func GteIterSVU16(a uint16, b []uint16, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a >= b[i]
		}
	}
	return
}

func GteIterSVU32(a uint32, b []uint32, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a >= b[i]
		}
	}
	return
}

func GteIterSVU64(a uint64, b []uint64, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a >= b[i]
		}
	}
	return
}

func GteIterSVF32(a float32, b []float32, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a >= b[i]
		}
	}
	return
}

func GteIterSVF64(a float64, b []float64, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a >= b[i]
		}
	}
	return
}

func GteIterSVStr(a string, b []string, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a >= b[i]
		}
	}
	return
}

func LtIterSVI(a int, b []int, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a < b[i]
		}
	}
	return
}

func LtIterSVI8(a int8, b []int8, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a < b[i]
		}
	}
	return
}

func LtIterSVI16(a int16, b []int16, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a < b[i]
		}
	}
	return
}

func LtIterSVI32(a int32, b []int32, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a < b[i]
		}
	}
	return
}

func LtIterSVI64(a int64, b []int64, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a < b[i]
		}
	}
	return
}

func LtIterSVU(a uint, b []uint, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a < b[i]
		}
	}
	return
}

func LtIterSVU8(a uint8, b []uint8, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a < b[i]
		}
	}
	return
}

func LtIterSVU16(a uint16, b []uint16, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a < b[i]
		}
	}
	return
}

func LtIterSVU32(a uint32, b []uint32, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a < b[i]
		}
	}
	return
}

func LtIterSVU64(a uint64, b []uint64, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a < b[i]
		}
	}
	return
}

func LtIterSVF32(a float32, b []float32, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a < b[i]
		}
	}
	return
}

func LtIterSVF64(a float64, b []float64, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a < b[i]
		}
	}
	return
}

func LtIterSVStr(a string, b []string, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a < b[i]
		}
	}
	return
}

func LteIterSVI(a int, b []int, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a <= b[i]
		}
	}
	return
}

func LteIterSVI8(a int8, b []int8, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a <= b[i]
		}
	}
	return
}

func LteIterSVI16(a int16, b []int16, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a <= b[i]
		}
	}
	return
}

func LteIterSVI32(a int32, b []int32, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a <= b[i]
		}
	}
	return
}

func LteIterSVI64(a int64, b []int64, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a <= b[i]
		}
	}
	return
}

func LteIterSVU(a uint, b []uint, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a <= b[i]
		}
	}
	return
}

func LteIterSVU8(a uint8, b []uint8, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a <= b[i]
		}
	}
	return
}

func LteIterSVU16(a uint16, b []uint16, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a <= b[i]
		}
	}
	return
}

func LteIterSVU32(a uint32, b []uint32, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a <= b[i]
		}
	}
	return
}

func LteIterSVU64(a uint64, b []uint64, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a <= b[i]
		}
	}
	return
}

func LteIterSVF32(a float32, b []float32, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a <= b[i]
		}
	}
	return
}

func LteIterSVF64(a float64, b []float64, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a <= b[i]
		}
	}
	return
}

func LteIterSVStr(a string, b []string, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a <= b[i]
		}
	}
	return
}

func EqIterSVB(a bool, b []bool, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a == b[i]
		}
	}
	return
}

func EqIterSVI(a int, b []int, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a == b[i]
		}
	}
	return
}

func EqIterSVI8(a int8, b []int8, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a == b[i]
		}
	}
	return
}

func EqIterSVI16(a int16, b []int16, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a == b[i]
		}
	}
	return
}

func EqIterSVI32(a int32, b []int32, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a == b[i]
		}
	}
	return
}

func EqIterSVI64(a int64, b []int64, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a == b[i]
		}
	}
	return
}

func EqIterSVU(a uint, b []uint, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a == b[i]
		}
	}
	return
}

func EqIterSVU8(a uint8, b []uint8, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a == b[i]
		}
	}
	return
}

func EqIterSVU16(a uint16, b []uint16, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a == b[i]
		}
	}
	return
}

func EqIterSVU32(a uint32, b []uint32, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a == b[i]
		}
	}
	return
}

func EqIterSVU64(a uint64, b []uint64, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a == b[i]
		}
	}
	return
}

func EqIterSVUintptr(a uintptr, b []uintptr, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a == b[i]
		}
	}
	return
}

func EqIterSVF32(a float32, b []float32, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a == b[i]
		}
	}
	return
}

func EqIterSVF64(a float64, b []float64, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a == b[i]
		}
	}
	return
}

func EqIterSVC64(a complex64, b []complex64, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a == b[i]
		}
	}
	return
}

func EqIterSVC128(a complex128, b []complex128, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a == b[i]
		}
	}
	return
}

func EqIterSVStr(a string, b []string, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a == b[i]
		}
	}
	return
}

func EqIterSVUnsafePointer(a unsafe.Pointer, b []unsafe.Pointer, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a == b[i]
		}
	}
	return
}

func NeIterSVB(a bool, b []bool, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a != b[i]
		}
	}
	return
}

func NeIterSVI(a int, b []int, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a != b[i]
		}
	}
	return
}

func NeIterSVI8(a int8, b []int8, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a != b[i]
		}
	}
	return
}

func NeIterSVI16(a int16, b []int16, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a != b[i]
		}
	}
	return
}

func NeIterSVI32(a int32, b []int32, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a != b[i]
		}
	}
	return
}

func NeIterSVI64(a int64, b []int64, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a != b[i]
		}
	}
	return
}

func NeIterSVU(a uint, b []uint, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a != b[i]
		}
	}
	return
}

func NeIterSVU8(a uint8, b []uint8, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a != b[i]
		}
	}
	return
}

func NeIterSVU16(a uint16, b []uint16, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a != b[i]
		}
	}
	return
}

func NeIterSVU32(a uint32, b []uint32, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a != b[i]
		}
	}
	return
}

func NeIterSVU64(a uint64, b []uint64, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a != b[i]
		}
	}
	return
}

func NeIterSVUintptr(a uintptr, b []uintptr, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a != b[i]
		}
	}
	return
}

func NeIterSVF32(a float32, b []float32, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a != b[i]
		}
	}
	return
}

func NeIterSVF64(a float64, b []float64, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a != b[i]
		}
	}
	return
}

func NeIterSVC64(a complex64, b []complex64, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a != b[i]
		}
	}
	return
}

func NeIterSVC128(a complex128, b []complex128, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a != b[i]
		}
	}
	return
}

func NeIterSVStr(a string, b []string, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a != b[i]
		}
	}
	return
}

func NeIterSVUnsafePointer(a unsafe.Pointer, b []unsafe.Pointer, retVal []bool, bit Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a != b[i]
		}
	}
	return
}

func GtSameIterSVI(a int, b []int, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a > b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func GtSameIterSVI8(a int8, b []int8, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a > b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func GtSameIterSVI16(a int16, b []int16, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a > b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func GtSameIterSVI32(a int32, b []int32, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a > b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func GtSameIterSVI64(a int64, b []int64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a > b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func GtSameIterSVU(a uint, b []uint, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a > b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func GtSameIterSVU8(a uint8, b []uint8, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a > b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func GtSameIterSVU16(a uint16, b []uint16, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a > b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func GtSameIterSVU32(a uint32, b []uint32, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a > b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func GtSameIterSVU64(a uint64, b []uint64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a > b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func GtSameIterSVF32(a float32, b []float32, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a > b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func GtSameIterSVF64(a float64, b []float64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a > b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func GtSameIterSVStr(a string, b []string, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a > b[i] {
				b[i] = "true"
			} else {
				b[i] = "false"
			}
		}
	}
	return
}

func GteSameIterSVI(a int, b []int, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a >= b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func GteSameIterSVI8(a int8, b []int8, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a >= b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func GteSameIterSVI16(a int16, b []int16, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a >= b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func GteSameIterSVI32(a int32, b []int32, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a >= b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func GteSameIterSVI64(a int64, b []int64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a >= b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func GteSameIterSVU(a uint, b []uint, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a >= b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func GteSameIterSVU8(a uint8, b []uint8, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a >= b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func GteSameIterSVU16(a uint16, b []uint16, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a >= b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func GteSameIterSVU32(a uint32, b []uint32, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a >= b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func GteSameIterSVU64(a uint64, b []uint64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a >= b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func GteSameIterSVF32(a float32, b []float32, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a >= b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func GteSameIterSVF64(a float64, b []float64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a >= b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func GteSameIterSVStr(a string, b []string, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a >= b[i] {
				b[i] = "true"
			} else {
				b[i] = "false"
			}
		}
	}
	return
}

func LtSameIterSVI(a int, b []int, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a < b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func LtSameIterSVI8(a int8, b []int8, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a < b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func LtSameIterSVI16(a int16, b []int16, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a < b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func LtSameIterSVI32(a int32, b []int32, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a < b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func LtSameIterSVI64(a int64, b []int64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a < b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func LtSameIterSVU(a uint, b []uint, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a < b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func LtSameIterSVU8(a uint8, b []uint8, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a < b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func LtSameIterSVU16(a uint16, b []uint16, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a < b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func LtSameIterSVU32(a uint32, b []uint32, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a < b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func LtSameIterSVU64(a uint64, b []uint64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a < b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func LtSameIterSVF32(a float32, b []float32, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a < b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func LtSameIterSVF64(a float64, b []float64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a < b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func LtSameIterSVStr(a string, b []string, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a < b[i] {
				b[i] = "true"
			} else {
				b[i] = "false"
			}
		}
	}
	return
}

func LteSameIterSVI(a int, b []int, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a <= b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func LteSameIterSVI8(a int8, b []int8, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a <= b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func LteSameIterSVI16(a int16, b []int16, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a <= b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func LteSameIterSVI32(a int32, b []int32, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a <= b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func LteSameIterSVI64(a int64, b []int64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a <= b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func LteSameIterSVU(a uint, b []uint, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a <= b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func LteSameIterSVU8(a uint8, b []uint8, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a <= b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func LteSameIterSVU16(a uint16, b []uint16, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a <= b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func LteSameIterSVU32(a uint32, b []uint32, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a <= b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func LteSameIterSVU64(a uint64, b []uint64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a <= b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func LteSameIterSVF32(a float32, b []float32, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a <= b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func LteSameIterSVF64(a float64, b []float64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a <= b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func LteSameIterSVStr(a string, b []string, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a <= b[i] {
				b[i] = "true"
			} else {
				b[i] = "false"
			}
		}
	}
	return
}

func EqSameIterSVB(a bool, b []bool, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a == b[i] {
				b[i] = true
			} else {
				b[i] = false
			}
		}
	}
	return
}

func EqSameIterSVI(a int, b []int, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a == b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func EqSameIterSVI8(a int8, b []int8, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a == b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func EqSameIterSVI16(a int16, b []int16, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a == b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func EqSameIterSVI32(a int32, b []int32, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a == b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func EqSameIterSVI64(a int64, b []int64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a == b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func EqSameIterSVU(a uint, b []uint, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a == b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func EqSameIterSVU8(a uint8, b []uint8, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a == b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func EqSameIterSVU16(a uint16, b []uint16, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a == b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func EqSameIterSVU32(a uint32, b []uint32, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a == b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func EqSameIterSVU64(a uint64, b []uint64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a == b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func EqSameIterSVUintptr(a uintptr, b []uintptr, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a == b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func EqSameIterSVF32(a float32, b []float32, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a == b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func EqSameIterSVF64(a float64, b []float64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a == b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func EqSameIterSVC64(a complex64, b []complex64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a == b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func EqSameIterSVC128(a complex128, b []complex128, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a == b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func EqSameIterSVStr(a string, b []string, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a == b[i] {
				b[i] = "true"
			} else {
				b[i] = "false"
			}
		}
	}
	return
}

func NeSameIterSVB(a bool, b []bool, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a != b[i] {
				b[i] = true
			} else {
				b[i] = false
			}
		}
	}
	return
}

func NeSameIterSVI(a int, b []int, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a != b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func NeSameIterSVI8(a int8, b []int8, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a != b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func NeSameIterSVI16(a int16, b []int16, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a != b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func NeSameIterSVI32(a int32, b []int32, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a != b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func NeSameIterSVI64(a int64, b []int64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a != b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func NeSameIterSVU(a uint, b []uint, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a != b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func NeSameIterSVU8(a uint8, b []uint8, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a != b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func NeSameIterSVU16(a uint16, b []uint16, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a != b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func NeSameIterSVU32(a uint32, b []uint32, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a != b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func NeSameIterSVU64(a uint64, b []uint64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a != b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func NeSameIterSVUintptr(a uintptr, b []uintptr, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a != b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func NeSameIterSVF32(a float32, b []float32, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a != b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func NeSameIterSVF64(a float64, b []float64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a != b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func NeSameIterSVC64(a complex64, b []complex64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a != b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func NeSameIterSVC128(a complex128, b []complex128, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a != b[i] {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
	return
}

func NeSameIterSVStr(a string, b []string, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a != b[i] {
				b[i] = "true"
			} else {
				b[i] = "false"
			}
		}
	}
	return
}

func GtVSI(a []int, b int, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] > b
	}
}

func GtVSI8(a []int8, b int8, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] > b
	}
}

func GtVSI16(a []int16, b int16, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] > b
	}
}

func GtVSI32(a []int32, b int32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] > b
	}
}

func GtVSI64(a []int64, b int64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] > b
	}
}

func GtVSU(a []uint, b uint, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] > b
	}
}

func GtVSU8(a []uint8, b uint8, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] > b
	}
}

func GtVSU16(a []uint16, b uint16, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] > b
	}
}

func GtVSU32(a []uint32, b uint32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] > b
	}
}

func GtVSU64(a []uint64, b uint64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] > b
	}
}

func GtVSF32(a []float32, b float32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] > b
	}
}

func GtVSF64(a []float64, b float64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] > b
	}
}

func GtVSStr(a []string, b string, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] > b
	}
}

func GteVSI(a []int, b int, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] >= b
	}
}

func GteVSI8(a []int8, b int8, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] >= b
	}
}

func GteVSI16(a []int16, b int16, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] >= b
	}
}

func GteVSI32(a []int32, b int32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] >= b
	}
}

func GteVSI64(a []int64, b int64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] >= b
	}
}

func GteVSU(a []uint, b uint, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] >= b
	}
}

func GteVSU8(a []uint8, b uint8, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] >= b
	}
}

func GteVSU16(a []uint16, b uint16, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] >= b
	}
}

func GteVSU32(a []uint32, b uint32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] >= b
	}
}

func GteVSU64(a []uint64, b uint64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] >= b
	}
}

func GteVSF32(a []float32, b float32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] >= b
	}
}

func GteVSF64(a []float64, b float64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] >= b
	}
}

func GteVSStr(a []string, b string, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] >= b
	}
}

func LtVSI(a []int, b int, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] < b
	}
}

func LtVSI8(a []int8, b int8, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] < b
	}
}

func LtVSI16(a []int16, b int16, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] < b
	}
}

func LtVSI32(a []int32, b int32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] < b
	}
}

func LtVSI64(a []int64, b int64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] < b
	}
}

func LtVSU(a []uint, b uint, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] < b
	}
}

func LtVSU8(a []uint8, b uint8, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] < b
	}
}

func LtVSU16(a []uint16, b uint16, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] < b
	}
}

func LtVSU32(a []uint32, b uint32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] < b
	}
}

func LtVSU64(a []uint64, b uint64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] < b
	}
}

func LtVSF32(a []float32, b float32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] < b
	}
}

func LtVSF64(a []float64, b float64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] < b
	}
}

func LtVSStr(a []string, b string, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] < b
	}
}

func LteVSI(a []int, b int, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] <= b
	}
}

func LteVSI8(a []int8, b int8, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] <= b
	}
}

func LteVSI16(a []int16, b int16, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] <= b
	}
}

func LteVSI32(a []int32, b int32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] <= b
	}
}

func LteVSI64(a []int64, b int64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] <= b
	}
}

func LteVSU(a []uint, b uint, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] <= b
	}
}

func LteVSU8(a []uint8, b uint8, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] <= b
	}
}

func LteVSU16(a []uint16, b uint16, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] <= b
	}
}

func LteVSU32(a []uint32, b uint32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] <= b
	}
}

func LteVSU64(a []uint64, b uint64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] <= b
	}
}

func LteVSF32(a []float32, b float32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] <= b
	}
}

func LteVSF64(a []float64, b float64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] <= b
	}
}

func LteVSStr(a []string, b string, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] <= b
	}
}

func EqVSB(a []bool, b bool, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] == b
	}
}

func EqVSI(a []int, b int, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] == b
	}
}

func EqVSI8(a []int8, b int8, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] == b
	}
}

func EqVSI16(a []int16, b int16, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] == b
	}
}

func EqVSI32(a []int32, b int32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] == b
	}
}

func EqVSI64(a []int64, b int64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] == b
	}
}

func EqVSU(a []uint, b uint, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] == b
	}
}

func EqVSU8(a []uint8, b uint8, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] == b
	}
}

func EqVSU16(a []uint16, b uint16, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] == b
	}
}

func EqVSU32(a []uint32, b uint32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] == b
	}
}

func EqVSU64(a []uint64, b uint64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] == b
	}
}

func EqVSUintptr(a []uintptr, b uintptr, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] == b
	}
}

func EqVSF32(a []float32, b float32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] == b
	}
}

func EqVSF64(a []float64, b float64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] == b
	}
}

func EqVSC64(a []complex64, b complex64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] == b
	}
}

func EqVSC128(a []complex128, b complex128, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] == b
	}
}

func EqVSStr(a []string, b string, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] == b
	}
}

func EqVSUnsafePointer(a []unsafe.Pointer, b unsafe.Pointer, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] == b
	}
}

func NeVSB(a []bool, b bool, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] != b
	}
}

func NeVSI(a []int, b int, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] != b
	}
}

func NeVSI8(a []int8, b int8, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] != b
	}
}

func NeVSI16(a []int16, b int16, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] != b
	}
}

func NeVSI32(a []int32, b int32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] != b
	}
}

func NeVSI64(a []int64, b int64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] != b
	}
}

func NeVSU(a []uint, b uint, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] != b
	}
}

func NeVSU8(a []uint8, b uint8, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] != b
	}
}

func NeVSU16(a []uint16, b uint16, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] != b
	}
}

func NeVSU32(a []uint32, b uint32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] != b
	}
}

func NeVSU64(a []uint64, b uint64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] != b
	}
}

func NeVSUintptr(a []uintptr, b uintptr, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] != b
	}
}

func NeVSF32(a []float32, b float32, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] != b
	}
}

func NeVSF64(a []float64, b float64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] != b
	}
}

func NeVSC64(a []complex64, b complex64, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] != b
	}
}

func NeVSC128(a []complex128, b complex128, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] != b
	}
}

func NeVSStr(a []string, b string, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] != b
	}
}

func NeVSUnsafePointer(a []unsafe.Pointer, b unsafe.Pointer, retVal []bool) {
	for i := range retVal {
		retVal[i] = a[i] != b
	}
}

func GtSameVSI(a []int, b int) {
	for i := range a {
		if a[i] > b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GtSameVSI8(a []int8, b int8) {
	for i := range a {
		if a[i] > b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GtSameVSI16(a []int16, b int16) {
	for i := range a {
		if a[i] > b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GtSameVSI32(a []int32, b int32) {
	for i := range a {
		if a[i] > b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GtSameVSI64(a []int64, b int64) {
	for i := range a {
		if a[i] > b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GtSameVSU(a []uint, b uint) {
	for i := range a {
		if a[i] > b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GtSameVSU8(a []uint8, b uint8) {
	for i := range a {
		if a[i] > b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GtSameVSU16(a []uint16, b uint16) {
	for i := range a {
		if a[i] > b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GtSameVSU32(a []uint32, b uint32) {
	for i := range a {
		if a[i] > b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GtSameVSU64(a []uint64, b uint64) {
	for i := range a {
		if a[i] > b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GtSameVSF32(a []float32, b float32) {
	for i := range a {
		if a[i] > b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GtSameVSF64(a []float64, b float64) {
	for i := range a {
		if a[i] > b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GtSameVSStr(a []string, b string) {
	for i := range a {
		if a[i] > b {
			a[i] = "true"
		} else {
			a[i] = "false"
		}
	}
}

func GteSameVSI(a []int, b int) {
	for i := range a {
		if a[i] >= b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GteSameVSI8(a []int8, b int8) {
	for i := range a {
		if a[i] >= b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GteSameVSI16(a []int16, b int16) {
	for i := range a {
		if a[i] >= b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GteSameVSI32(a []int32, b int32) {
	for i := range a {
		if a[i] >= b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GteSameVSI64(a []int64, b int64) {
	for i := range a {
		if a[i] >= b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GteSameVSU(a []uint, b uint) {
	for i := range a {
		if a[i] >= b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GteSameVSU8(a []uint8, b uint8) {
	for i := range a {
		if a[i] >= b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GteSameVSU16(a []uint16, b uint16) {
	for i := range a {
		if a[i] >= b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GteSameVSU32(a []uint32, b uint32) {
	for i := range a {
		if a[i] >= b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GteSameVSU64(a []uint64, b uint64) {
	for i := range a {
		if a[i] >= b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GteSameVSF32(a []float32, b float32) {
	for i := range a {
		if a[i] >= b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GteSameVSF64(a []float64, b float64) {
	for i := range a {
		if a[i] >= b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func GteSameVSStr(a []string, b string) {
	for i := range a {
		if a[i] >= b {
			a[i] = "true"
		} else {
			a[i] = "false"
		}
	}
}

func LtSameVSI(a []int, b int) {
	for i := range a {
		if a[i] < b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LtSameVSI8(a []int8, b int8) {
	for i := range a {
		if a[i] < b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LtSameVSI16(a []int16, b int16) {
	for i := range a {
		if a[i] < b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LtSameVSI32(a []int32, b int32) {
	for i := range a {
		if a[i] < b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LtSameVSI64(a []int64, b int64) {
	for i := range a {
		if a[i] < b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LtSameVSU(a []uint, b uint) {
	for i := range a {
		if a[i] < b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LtSameVSU8(a []uint8, b uint8) {
	for i := range a {
		if a[i] < b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LtSameVSU16(a []uint16, b uint16) {
	for i := range a {
		if a[i] < b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LtSameVSU32(a []uint32, b uint32) {
	for i := range a {
		if a[i] < b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LtSameVSU64(a []uint64, b uint64) {
	for i := range a {
		if a[i] < b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LtSameVSF32(a []float32, b float32) {
	for i := range a {
		if a[i] < b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LtSameVSF64(a []float64, b float64) {
	for i := range a {
		if a[i] < b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LtSameVSStr(a []string, b string) {
	for i := range a {
		if a[i] < b {
			a[i] = "true"
		} else {
			a[i] = "false"
		}
	}
}

func LteSameVSI(a []int, b int) {
	for i := range a {
		if a[i] <= b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LteSameVSI8(a []int8, b int8) {
	for i := range a {
		if a[i] <= b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LteSameVSI16(a []int16, b int16) {
	for i := range a {
		if a[i] <= b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LteSameVSI32(a []int32, b int32) {
	for i := range a {
		if a[i] <= b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LteSameVSI64(a []int64, b int64) {
	for i := range a {
		if a[i] <= b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LteSameVSU(a []uint, b uint) {
	for i := range a {
		if a[i] <= b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LteSameVSU8(a []uint8, b uint8) {
	for i := range a {
		if a[i] <= b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LteSameVSU16(a []uint16, b uint16) {
	for i := range a {
		if a[i] <= b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LteSameVSU32(a []uint32, b uint32) {
	for i := range a {
		if a[i] <= b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LteSameVSU64(a []uint64, b uint64) {
	for i := range a {
		if a[i] <= b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LteSameVSF32(a []float32, b float32) {
	for i := range a {
		if a[i] <= b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LteSameVSF64(a []float64, b float64) {
	for i := range a {
		if a[i] <= b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func LteSameVSStr(a []string, b string) {
	for i := range a {
		if a[i] <= b {
			a[i] = "true"
		} else {
			a[i] = "false"
		}
	}
}

func EqSameVSB(a []bool, b bool) {
	for i := range a {
		if a[i] == b {
			a[i] = true
		} else {
			a[i] = false
		}
	}
}

func EqSameVSI(a []int, b int) {
	for i := range a {
		if a[i] == b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func EqSameVSI8(a []int8, b int8) {
	for i := range a {
		if a[i] == b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func EqSameVSI16(a []int16, b int16) {
	for i := range a {
		if a[i] == b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func EqSameVSI32(a []int32, b int32) {
	for i := range a {
		if a[i] == b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func EqSameVSI64(a []int64, b int64) {
	for i := range a {
		if a[i] == b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func EqSameVSU(a []uint, b uint) {
	for i := range a {
		if a[i] == b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func EqSameVSU8(a []uint8, b uint8) {
	for i := range a {
		if a[i] == b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func EqSameVSU16(a []uint16, b uint16) {
	for i := range a {
		if a[i] == b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func EqSameVSU32(a []uint32, b uint32) {
	for i := range a {
		if a[i] == b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func EqSameVSU64(a []uint64, b uint64) {
	for i := range a {
		if a[i] == b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func EqSameVSUintptr(a []uintptr, b uintptr) {
	for i := range a {
		if a[i] == b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func EqSameVSF32(a []float32, b float32) {
	for i := range a {
		if a[i] == b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func EqSameVSF64(a []float64, b float64) {
	for i := range a {
		if a[i] == b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func EqSameVSC64(a []complex64, b complex64) {
	for i := range a {
		if a[i] == b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func EqSameVSC128(a []complex128, b complex128) {
	for i := range a {
		if a[i] == b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func EqSameVSStr(a []string, b string) {
	for i := range a {
		if a[i] == b {
			a[i] = "true"
		} else {
			a[i] = "false"
		}
	}
}

func NeSameVSB(a []bool, b bool) {
	for i := range a {
		if a[i] != b {
			a[i] = true
		} else {
			a[i] = false
		}
	}
}

func NeSameVSI(a []int, b int) {
	for i := range a {
		if a[i] != b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func NeSameVSI8(a []int8, b int8) {
	for i := range a {
		if a[i] != b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func NeSameVSI16(a []int16, b int16) {
	for i := range a {
		if a[i] != b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func NeSameVSI32(a []int32, b int32) {
	for i := range a {
		if a[i] != b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func NeSameVSI64(a []int64, b int64) {
	for i := range a {
		if a[i] != b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func NeSameVSU(a []uint, b uint) {
	for i := range a {
		if a[i] != b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func NeSameVSU8(a []uint8, b uint8) {
	for i := range a {
		if a[i] != b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func NeSameVSU16(a []uint16, b uint16) {
	for i := range a {
		if a[i] != b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func NeSameVSU32(a []uint32, b uint32) {
	for i := range a {
		if a[i] != b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func NeSameVSU64(a []uint64, b uint64) {
	for i := range a {
		if a[i] != b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func NeSameVSUintptr(a []uintptr, b uintptr) {
	for i := range a {
		if a[i] != b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func NeSameVSF32(a []float32, b float32) {
	for i := range a {
		if a[i] != b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func NeSameVSF64(a []float64, b float64) {
	for i := range a {
		if a[i] != b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func NeSameVSC64(a []complex64, b complex64) {
	for i := range a {
		if a[i] != b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func NeSameVSC128(a []complex128, b complex128) {
	for i := range a {
		if a[i] != b {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}

func NeSameVSStr(a []string, b string) {
	for i := range a {
		if a[i] != b {
			a[i] = "true"
		} else {
			a[i] = "false"
		}
	}
}

func GtIterVSI(a []int, b int, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] > b
		}
	}
	return
}

func GtIterVSI8(a []int8, b int8, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] > b
		}
	}
	return
}

func GtIterVSI16(a []int16, b int16, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] > b
		}
	}
	return
}

func GtIterVSI32(a []int32, b int32, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] > b
		}
	}
	return
}

func GtIterVSI64(a []int64, b int64, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] > b
		}
	}
	return
}

func GtIterVSU(a []uint, b uint, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] > b
		}
	}
	return
}

func GtIterVSU8(a []uint8, b uint8, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] > b
		}
	}
	return
}

func GtIterVSU16(a []uint16, b uint16, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] > b
		}
	}
	return
}

func GtIterVSU32(a []uint32, b uint32, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] > b
		}
	}
	return
}

func GtIterVSU64(a []uint64, b uint64, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] > b
		}
	}
	return
}

func GtIterVSF32(a []float32, b float32, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] > b
		}
	}
	return
}

func GtIterVSF64(a []float64, b float64, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] > b
		}
	}
	return
}

func GtIterVSStr(a []string, b string, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] > b
		}
	}
	return
}

func GteIterVSI(a []int, b int, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] >= b
		}
	}
	return
}

func GteIterVSI8(a []int8, b int8, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] >= b
		}
	}
	return
}

func GteIterVSI16(a []int16, b int16, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] >= b
		}
	}
	return
}

func GteIterVSI32(a []int32, b int32, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] >= b
		}
	}
	return
}

func GteIterVSI64(a []int64, b int64, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] >= b
		}
	}
	return
}

func GteIterVSU(a []uint, b uint, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] >= b
		}
	}
	return
}

func GteIterVSU8(a []uint8, b uint8, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] >= b
		}
	}
	return
}

func GteIterVSU16(a []uint16, b uint16, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] >= b
		}
	}
	return
}

func GteIterVSU32(a []uint32, b uint32, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] >= b
		}
	}
	return
}

func GteIterVSU64(a []uint64, b uint64, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] >= b
		}
	}
	return
}

func GteIterVSF32(a []float32, b float32, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] >= b
		}
	}
	return
}

func GteIterVSF64(a []float64, b float64, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] >= b
		}
	}
	return
}

func GteIterVSStr(a []string, b string, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] >= b
		}
	}
	return
}

func LtIterVSI(a []int, b int, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] < b
		}
	}
	return
}

func LtIterVSI8(a []int8, b int8, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] < b
		}
	}
	return
}

func LtIterVSI16(a []int16, b int16, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] < b
		}
	}
	return
}

func LtIterVSI32(a []int32, b int32, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] < b
		}
	}
	return
}

func LtIterVSI64(a []int64, b int64, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] < b
		}
	}
	return
}

func LtIterVSU(a []uint, b uint, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] < b
		}
	}
	return
}

func LtIterVSU8(a []uint8, b uint8, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] < b
		}
	}
	return
}

func LtIterVSU16(a []uint16, b uint16, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] < b
		}
	}
	return
}

func LtIterVSU32(a []uint32, b uint32, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] < b
		}
	}
	return
}

func LtIterVSU64(a []uint64, b uint64, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] < b
		}
	}
	return
}

func LtIterVSF32(a []float32, b float32, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] < b
		}
	}
	return
}

func LtIterVSF64(a []float64, b float64, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] < b
		}
	}
	return
}

func LtIterVSStr(a []string, b string, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] < b
		}
	}
	return
}

func LteIterVSI(a []int, b int, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] <= b
		}
	}
	return
}

func LteIterVSI8(a []int8, b int8, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] <= b
		}
	}
	return
}

func LteIterVSI16(a []int16, b int16, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] <= b
		}
	}
	return
}

func LteIterVSI32(a []int32, b int32, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] <= b
		}
	}
	return
}

func LteIterVSI64(a []int64, b int64, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] <= b
		}
	}
	return
}

func LteIterVSU(a []uint, b uint, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] <= b
		}
	}
	return
}

func LteIterVSU8(a []uint8, b uint8, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] <= b
		}
	}
	return
}

func LteIterVSU16(a []uint16, b uint16, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] <= b
		}
	}
	return
}

func LteIterVSU32(a []uint32, b uint32, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] <= b
		}
	}
	return
}

func LteIterVSU64(a []uint64, b uint64, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] <= b
		}
	}
	return
}

func LteIterVSF32(a []float32, b float32, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] <= b
		}
	}
	return
}

func LteIterVSF64(a []float64, b float64, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] <= b
		}
	}
	return
}

func LteIterVSStr(a []string, b string, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] <= b
		}
	}
	return
}

func EqIterVSB(a []bool, b bool, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] == b
		}
	}
	return
}

func EqIterVSI(a []int, b int, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] == b
		}
	}
	return
}

func EqIterVSI8(a []int8, b int8, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] == b
		}
	}
	return
}

func EqIterVSI16(a []int16, b int16, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] == b
		}
	}
	return
}

func EqIterVSI32(a []int32, b int32, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] == b
		}
	}
	return
}

func EqIterVSI64(a []int64, b int64, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] == b
		}
	}
	return
}

func EqIterVSU(a []uint, b uint, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] == b
		}
	}
	return
}

func EqIterVSU8(a []uint8, b uint8, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] == b
		}
	}
	return
}

func EqIterVSU16(a []uint16, b uint16, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] == b
		}
	}
	return
}

func EqIterVSU32(a []uint32, b uint32, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] == b
		}
	}
	return
}

func EqIterVSU64(a []uint64, b uint64, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] == b
		}
	}
	return
}

func EqIterVSUintptr(a []uintptr, b uintptr, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] == b
		}
	}
	return
}

func EqIterVSF32(a []float32, b float32, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] == b
		}
	}
	return
}

func EqIterVSF64(a []float64, b float64, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] == b
		}
	}
	return
}

func EqIterVSC64(a []complex64, b complex64, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] == b
		}
	}
	return
}

func EqIterVSC128(a []complex128, b complex128, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] == b
		}
	}
	return
}

func EqIterVSStr(a []string, b string, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] == b
		}
	}
	return
}

func EqIterVSUnsafePointer(a []unsafe.Pointer, b unsafe.Pointer, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] == b
		}
	}
	return
}

func NeIterVSB(a []bool, b bool, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] != b
		}
	}
	return
}

func NeIterVSI(a []int, b int, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] != b
		}
	}
	return
}

func NeIterVSI8(a []int8, b int8, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] != b
		}
	}
	return
}

func NeIterVSI16(a []int16, b int16, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] != b
		}
	}
	return
}

func NeIterVSI32(a []int32, b int32, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] != b
		}
	}
	return
}

func NeIterVSI64(a []int64, b int64, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] != b
		}
	}
	return
}

func NeIterVSU(a []uint, b uint, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] != b
		}
	}
	return
}

func NeIterVSU8(a []uint8, b uint8, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] != b
		}
	}
	return
}

func NeIterVSU16(a []uint16, b uint16, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] != b
		}
	}
	return
}

func NeIterVSU32(a []uint32, b uint32, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] != b
		}
	}
	return
}

func NeIterVSU64(a []uint64, b uint64, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] != b
		}
	}
	return
}

func NeIterVSUintptr(a []uintptr, b uintptr, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] != b
		}
	}
	return
}

func NeIterVSF32(a []float32, b float32, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] != b
		}
	}
	return
}

func NeIterVSF64(a []float64, b float64, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] != b
		}
	}
	return
}

func NeIterVSC64(a []complex64, b complex64, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] != b
		}
	}
	return
}

func NeIterVSC128(a []complex128, b complex128, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] != b
		}
	}
	return
}

func NeIterVSStr(a []string, b string, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] != b
		}
	}
	return
}

func NeIterVSUnsafePointer(a []unsafe.Pointer, b unsafe.Pointer, retVal []bool, ait Iterator, rit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = rit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			retVal[k] = a[i] != b
		}
	}
	return
}

func GtSameIterVSI(a []int, b int, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] > b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GtSameIterVSI8(a []int8, b int8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] > b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GtSameIterVSI16(a []int16, b int16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] > b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GtSameIterVSI32(a []int32, b int32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] > b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GtSameIterVSI64(a []int64, b int64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] > b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GtSameIterVSU(a []uint, b uint, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] > b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GtSameIterVSU8(a []uint8, b uint8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] > b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GtSameIterVSU16(a []uint16, b uint16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] > b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GtSameIterVSU32(a []uint32, b uint32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] > b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GtSameIterVSU64(a []uint64, b uint64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] > b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GtSameIterVSF32(a []float32, b float32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] > b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GtSameIterVSF64(a []float64, b float64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] > b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GtSameIterVSStr(a []string, b string, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] > b {
				a[i] = "true"
			} else {
				a[i] = "false"
			}
		}
	}
	return
}

func GteSameIterVSI(a []int, b int, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] >= b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GteSameIterVSI8(a []int8, b int8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] >= b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GteSameIterVSI16(a []int16, b int16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] >= b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GteSameIterVSI32(a []int32, b int32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] >= b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GteSameIterVSI64(a []int64, b int64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] >= b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GteSameIterVSU(a []uint, b uint, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] >= b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GteSameIterVSU8(a []uint8, b uint8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] >= b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GteSameIterVSU16(a []uint16, b uint16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] >= b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GteSameIterVSU32(a []uint32, b uint32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] >= b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GteSameIterVSU64(a []uint64, b uint64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] >= b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GteSameIterVSF32(a []float32, b float32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] >= b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GteSameIterVSF64(a []float64, b float64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] >= b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func GteSameIterVSStr(a []string, b string, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] >= b {
				a[i] = "true"
			} else {
				a[i] = "false"
			}
		}
	}
	return
}

func LtSameIterVSI(a []int, b int, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LtSameIterVSI8(a []int8, b int8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LtSameIterVSI16(a []int16, b int16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LtSameIterVSI32(a []int32, b int32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LtSameIterVSI64(a []int64, b int64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LtSameIterVSU(a []uint, b uint, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LtSameIterVSU8(a []uint8, b uint8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LtSameIterVSU16(a []uint16, b uint16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LtSameIterVSU32(a []uint32, b uint32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LtSameIterVSU64(a []uint64, b uint64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LtSameIterVSF32(a []float32, b float32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LtSameIterVSF64(a []float64, b float64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LtSameIterVSStr(a []string, b string, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] < b {
				a[i] = "true"
			} else {
				a[i] = "false"
			}
		}
	}
	return
}

func LteSameIterVSI(a []int, b int, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] <= b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LteSameIterVSI8(a []int8, b int8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] <= b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LteSameIterVSI16(a []int16, b int16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] <= b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LteSameIterVSI32(a []int32, b int32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] <= b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LteSameIterVSI64(a []int64, b int64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] <= b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LteSameIterVSU(a []uint, b uint, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] <= b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LteSameIterVSU8(a []uint8, b uint8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] <= b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LteSameIterVSU16(a []uint16, b uint16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] <= b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LteSameIterVSU32(a []uint32, b uint32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] <= b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LteSameIterVSU64(a []uint64, b uint64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] <= b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LteSameIterVSF32(a []float32, b float32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] <= b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LteSameIterVSF64(a []float64, b float64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] <= b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func LteSameIterVSStr(a []string, b string, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] <= b {
				a[i] = "true"
			} else {
				a[i] = "false"
			}
		}
	}
	return
}

func EqSameIterVSB(a []bool, b bool, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] == b {
				a[i] = true
			} else {
				a[i] = false
			}
		}
	}
	return
}

func EqSameIterVSI(a []int, b int, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] == b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func EqSameIterVSI8(a []int8, b int8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] == b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func EqSameIterVSI16(a []int16, b int16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] == b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func EqSameIterVSI32(a []int32, b int32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] == b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func EqSameIterVSI64(a []int64, b int64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] == b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func EqSameIterVSU(a []uint, b uint, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] == b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func EqSameIterVSU8(a []uint8, b uint8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] == b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func EqSameIterVSU16(a []uint16, b uint16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] == b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func EqSameIterVSU32(a []uint32, b uint32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] == b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func EqSameIterVSU64(a []uint64, b uint64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] == b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func EqSameIterVSUintptr(a []uintptr, b uintptr, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] == b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func EqSameIterVSF32(a []float32, b float32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] == b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func EqSameIterVSF64(a []float64, b float64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] == b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func EqSameIterVSC64(a []complex64, b complex64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] == b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func EqSameIterVSC128(a []complex128, b complex128, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] == b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func EqSameIterVSStr(a []string, b string, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] == b {
				a[i] = "true"
			} else {
				a[i] = "false"
			}
		}
	}
	return
}

func NeSameIterVSB(a []bool, b bool, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] != b {
				a[i] = true
			} else {
				a[i] = false
			}
		}
	}
	return
}

func NeSameIterVSI(a []int, b int, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] != b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func NeSameIterVSI8(a []int8, b int8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] != b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func NeSameIterVSI16(a []int16, b int16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] != b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func NeSameIterVSI32(a []int32, b int32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] != b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func NeSameIterVSI64(a []int64, b int64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] != b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func NeSameIterVSU(a []uint, b uint, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] != b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func NeSameIterVSU8(a []uint8, b uint8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] != b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func NeSameIterVSU16(a []uint16, b uint16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] != b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func NeSameIterVSU32(a []uint32, b uint32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] != b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func NeSameIterVSU64(a []uint64, b uint64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] != b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func NeSameIterVSUintptr(a []uintptr, b uintptr, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] != b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func NeSameIterVSF32(a []float32, b float32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] != b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func NeSameIterVSF64(a []float64, b float64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] != b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func NeSameIterVSC64(a []complex64, b complex64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] != b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func NeSameIterVSC128(a []complex128, b complex128, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] != b {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
	}
	return
}

func NeSameIterVSStr(a []string, b string, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if a[i] != b {
				a[i] = "true"
			} else {
				a[i] = "false"
			}
		}
	}
	return
}
