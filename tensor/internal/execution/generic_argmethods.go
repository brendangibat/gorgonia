package execution

import (
	"math"

	"github.com/chewxy/math32"
)

/*
GENERATED FILE. DO NOT EDIT
*/

func ArgmaxI(a []int) int {
	var set bool
	var f int
	var max int
	for i := range a {
		v := a[i]
		if !set {
			f = v
			max = i
			set = true
			continue
		}
		if v > f {
			max = i
			f = v
		}
	}
	return max
}

func ArgmaxI8(a []int8) int {
	var set bool
	var f int8
	var max int
	for i := range a {
		v := a[i]
		if !set {
			f = v
			max = i
			set = true
			continue
		}
		if v > f {
			max = i
			f = v
		}
	}
	return max
}

func ArgmaxI16(a []int16) int {
	var set bool
	var f int16
	var max int
	for i := range a {
		v := a[i]
		if !set {
			f = v
			max = i
			set = true
			continue
		}
		if v > f {
			max = i
			f = v
		}
	}
	return max
}

func ArgmaxI32(a []int32) int {
	var set bool
	var f int32
	var max int
	for i := range a {
		v := a[i]
		if !set {
			f = v
			max = i
			set = true
			continue
		}
		if v > f {
			max = i
			f = v
		}
	}
	return max
}

func ArgmaxI64(a []int64) int {
	var set bool
	var f int64
	var max int
	for i := range a {
		v := a[i]
		if !set {
			f = v
			max = i
			set = true
			continue
		}
		if v > f {
			max = i
			f = v
		}
	}
	return max
}

func ArgmaxU(a []uint) int {
	var set bool
	var f uint
	var max int
	for i := range a {
		v := a[i]
		if !set {
			f = v
			max = i
			set = true
			continue
		}
		if v > f {
			max = i
			f = v
		}
	}
	return max
}

func ArgmaxU8(a []uint8) int {
	var set bool
	var f uint8
	var max int
	for i := range a {
		v := a[i]
		if !set {
			f = v
			max = i
			set = true
			continue
		}
		if v > f {
			max = i
			f = v
		}
	}
	return max
}

func ArgmaxU16(a []uint16) int {
	var set bool
	var f uint16
	var max int
	for i := range a {
		v := a[i]
		if !set {
			f = v
			max = i
			set = true
			continue
		}
		if v > f {
			max = i
			f = v
		}
	}
	return max
}

func ArgmaxU32(a []uint32) int {
	var set bool
	var f uint32
	var max int
	for i := range a {
		v := a[i]
		if !set {
			f = v
			max = i
			set = true
			continue
		}
		if v > f {
			max = i
			f = v
		}
	}
	return max
}

func ArgmaxU64(a []uint64) int {
	var set bool
	var f uint64
	var max int
	for i := range a {
		v := a[i]
		if !set {
			f = v
			max = i
			set = true
			continue
		}
		if v > f {
			max = i
			f = v
		}
	}
	return max
}

func ArgmaxF32(a []float32) int {
	var set bool
	var f float32
	var max int
	for i := range a {
		v := a[i]
		if !set {
			f = v
			max = i
			set = true
			continue
		}
		if math32.IsNaN(v) || math32.IsInf(v, 1) {
			max = i
			return max
		}
		if v > f {
			max = i
			f = v
		}
	}
	return max
}

func ArgmaxF64(a []float64) int {
	var set bool
	var f float64
	var max int
	for i := range a {
		v := a[i]
		if !set {
			f = v
			max = i
			set = true
			continue
		}
		if math.IsNaN(v) || math.IsInf(v, 1) {
			max = i
			return max
		}
		if v > f {
			max = i
			f = v
		}
	}
	return max
}

func ArgmaxStr(a []string) int {
	var set bool
	var f string
	var max int
	for i := range a {
		v := a[i]
		if !set {
			f = v
			max = i
			set = true
			continue
		}
		if v > f {
			max = i
			f = v
		}
	}
	return max
}

func ArgmaxMaskedI(a []int, mask []bool) int {
	var set bool
	var f int
	var max int
	for i := range a {
		if mask[i] {
			continue
		}
		v := a[i]
		if !set {
			f = v
			max = i
			set = true
			continue
		}
		if v > f {
			max = i
			f = v
		}
	}
	return max
}

func ArgmaxMaskedI8(a []int8, mask []bool) int {
	var set bool
	var f int8
	var max int
	for i := range a {
		if mask[i] {
			continue
		}
		v := a[i]
		if !set {
			f = v
			max = i
			set = true
			continue
		}
		if v > f {
			max = i
			f = v
		}
	}
	return max
}

func ArgmaxMaskedI16(a []int16, mask []bool) int {
	var set bool
	var f int16
	var max int
	for i := range a {
		if mask[i] {
			continue
		}
		v := a[i]
		if !set {
			f = v
			max = i
			set = true
			continue
		}
		if v > f {
			max = i
			f = v
		}
	}
	return max
}

func ArgmaxMaskedI32(a []int32, mask []bool) int {
	var set bool
	var f int32
	var max int
	for i := range a {
		if mask[i] {
			continue
		}
		v := a[i]
		if !set {
			f = v
			max = i
			set = true
			continue
		}
		if v > f {
			max = i
			f = v
		}
	}
	return max
}

func ArgmaxMaskedI64(a []int64, mask []bool) int {
	var set bool
	var f int64
	var max int
	for i := range a {
		if mask[i] {
			continue
		}
		v := a[i]
		if !set {
			f = v
			max = i
			set = true
			continue
		}
		if v > f {
			max = i
			f = v
		}
	}
	return max
}

func ArgmaxMaskedU(a []uint, mask []bool) int {
	var set bool
	var f uint
	var max int
	for i := range a {
		if mask[i] {
			continue
		}
		v := a[i]
		if !set {
			f = v
			max = i
			set = true
			continue
		}
		if v > f {
			max = i
			f = v
		}
	}
	return max
}

func ArgmaxMaskedU8(a []uint8, mask []bool) int {
	var set bool
	var f uint8
	var max int
	for i := range a {
		if mask[i] {
			continue
		}
		v := a[i]
		if !set {
			f = v
			max = i
			set = true
			continue
		}
		if v > f {
			max = i
			f = v
		}
	}
	return max
}

func ArgmaxMaskedU16(a []uint16, mask []bool) int {
	var set bool
	var f uint16
	var max int
	for i := range a {
		if mask[i] {
			continue
		}
		v := a[i]
		if !set {
			f = v
			max = i
			set = true
			continue
		}
		if v > f {
			max = i
			f = v
		}
	}
	return max
}

func ArgmaxMaskedU32(a []uint32, mask []bool) int {
	var set bool
	var f uint32
	var max int
	for i := range a {
		if mask[i] {
			continue
		}
		v := a[i]
		if !set {
			f = v
			max = i
			set = true
			continue
		}
		if v > f {
			max = i
			f = v
		}
	}
	return max
}

func ArgmaxMaskedU64(a []uint64, mask []bool) int {
	var set bool
	var f uint64
	var max int
	for i := range a {
		if mask[i] {
			continue
		}
		v := a[i]
		if !set {
			f = v
			max = i
			set = true
			continue
		}
		if v > f {
			max = i
			f = v
		}
	}
	return max
}

func ArgmaxMaskedF32(a []float32, mask []bool) int {
	var set bool
	var f float32
	var max int
	for i := range a {
		if mask[i] {
			continue
		}
		v := a[i]
		if !set {
			f = v
			max = i
			set = true
			continue
		}
		if math32.IsNaN(v) || math32.IsInf(v, 1) {
			max = i
			return max
		}
		if v > f {
			max = i
			f = v
		}
	}
	return max
}

func ArgmaxMaskedF64(a []float64, mask []bool) int {
	var set bool
	var f float64
	var max int
	for i := range a {
		if mask[i] {
			continue
		}
		v := a[i]
		if !set {
			f = v
			max = i
			set = true
			continue
		}
		if math.IsNaN(v) || math.IsInf(v, 1) {
			max = i
			return max
		}
		if v > f {
			max = i
			f = v
		}
	}
	return max
}

func ArgmaxMaskedStr(a []string, mask []bool) int {
	var set bool
	var f string
	var max int
	for i := range a {
		if mask[i] {
			continue
		}
		v := a[i]
		if !set {
			f = v
			max = i
			set = true
			continue
		}
		if v > f {
			max = i
			f = v
		}
	}
	return max
}

func ArgminI(a []int) int {
	var set bool
	var f int
	var min int
	for i := range a {
		v := a[i]
		if !set {
			f = v
			min = i
			set = true
			continue
		}
		if v < f {
			min = i
			f = v
		}
	}
	return min
}

func ArgminI8(a []int8) int {
	var set bool
	var f int8
	var min int
	for i := range a {
		v := a[i]
		if !set {
			f = v
			min = i
			set = true
			continue
		}
		if v < f {
			min = i
			f = v
		}
	}
	return min
}

func ArgminI16(a []int16) int {
	var set bool
	var f int16
	var min int
	for i := range a {
		v := a[i]
		if !set {
			f = v
			min = i
			set = true
			continue
		}
		if v < f {
			min = i
			f = v
		}
	}
	return min
}

func ArgminI32(a []int32) int {
	var set bool
	var f int32
	var min int
	for i := range a {
		v := a[i]
		if !set {
			f = v
			min = i
			set = true
			continue
		}
		if v < f {
			min = i
			f = v
		}
	}
	return min
}

func ArgminI64(a []int64) int {
	var set bool
	var f int64
	var min int
	for i := range a {
		v := a[i]
		if !set {
			f = v
			min = i
			set = true
			continue
		}
		if v < f {
			min = i
			f = v
		}
	}
	return min
}

func ArgminU(a []uint) int {
	var set bool
	var f uint
	var min int
	for i := range a {
		v := a[i]
		if !set {
			f = v
			min = i
			set = true
			continue
		}
		if v < f {
			min = i
			f = v
		}
	}
	return min
}

func ArgminU8(a []uint8) int {
	var set bool
	var f uint8
	var min int
	for i := range a {
		v := a[i]
		if !set {
			f = v
			min = i
			set = true
			continue
		}
		if v < f {
			min = i
			f = v
		}
	}
	return min
}

func ArgminU16(a []uint16) int {
	var set bool
	var f uint16
	var min int
	for i := range a {
		v := a[i]
		if !set {
			f = v
			min = i
			set = true
			continue
		}
		if v < f {
			min = i
			f = v
		}
	}
	return min
}

func ArgminU32(a []uint32) int {
	var set bool
	var f uint32
	var min int
	for i := range a {
		v := a[i]
		if !set {
			f = v
			min = i
			set = true
			continue
		}
		if v < f {
			min = i
			f = v
		}
	}
	return min
}

func ArgminU64(a []uint64) int {
	var set bool
	var f uint64
	var min int
	for i := range a {
		v := a[i]
		if !set {
			f = v
			min = i
			set = true
			continue
		}
		if v < f {
			min = i
			f = v
		}
	}
	return min
}

func ArgminF32(a []float32) int {
	var set bool
	var f float32
	var min int
	for i := range a {
		v := a[i]
		if !set {
			f = v
			min = i
			set = true
			continue
		}
		if math32.IsNaN(v) || math32.IsInf(v, -1) {
			min = i
			return min
		}
		if v < f {
			min = i
			f = v
		}
	}
	return min
}

func ArgminF64(a []float64) int {
	var set bool
	var f float64
	var min int
	for i := range a {
		v := a[i]
		if !set {
			f = v
			min = i
			set = true
			continue
		}
		if math.IsNaN(v) || math.IsInf(v, -1) {
			min = i
			return min
		}
		if v < f {
			min = i
			f = v
		}
	}
	return min
}

func ArgminStr(a []string) int {
	var set bool
	var f string
	var min int
	for i := range a {
		v := a[i]
		if !set {
			f = v
			min = i
			set = true
			continue
		}
		if v < f {
			min = i
			f = v
		}
	}
	return min
}

func ArgminMaskedI(a []int, mask []bool) int {
	var set bool
	var f int
	var min int
	for i := range a {
		if mask[i] {
			continue
		}
		v := a[i]
		if !set {
			f = v
			min = i
			set = true
			continue
		}
		if v < f {
			min = i
			f = v
		}
	}
	return min
}

func ArgminMaskedI8(a []int8, mask []bool) int {
	var set bool
	var f int8
	var min int
	for i := range a {
		if mask[i] {
			continue
		}
		v := a[i]
		if !set {
			f = v
			min = i
			set = true
			continue
		}
		if v < f {
			min = i
			f = v
		}
	}
	return min
}

func ArgminMaskedI16(a []int16, mask []bool) int {
	var set bool
	var f int16
	var min int
	for i := range a {
		if mask[i] {
			continue
		}
		v := a[i]
		if !set {
			f = v
			min = i
			set = true
			continue
		}
		if v < f {
			min = i
			f = v
		}
	}
	return min
}

func ArgminMaskedI32(a []int32, mask []bool) int {
	var set bool
	var f int32
	var min int
	for i := range a {
		if mask[i] {
			continue
		}
		v := a[i]
		if !set {
			f = v
			min = i
			set = true
			continue
		}
		if v < f {
			min = i
			f = v
		}
	}
	return min
}

func ArgminMaskedI64(a []int64, mask []bool) int {
	var set bool
	var f int64
	var min int
	for i := range a {
		if mask[i] {
			continue
		}
		v := a[i]
		if !set {
			f = v
			min = i
			set = true
			continue
		}
		if v < f {
			min = i
			f = v
		}
	}
	return min
}

func ArgminMaskedU(a []uint, mask []bool) int {
	var set bool
	var f uint
	var min int
	for i := range a {
		if mask[i] {
			continue
		}
		v := a[i]
		if !set {
			f = v
			min = i
			set = true
			continue
		}
		if v < f {
			min = i
			f = v
		}
	}
	return min
}

func ArgminMaskedU8(a []uint8, mask []bool) int {
	var set bool
	var f uint8
	var min int
	for i := range a {
		if mask[i] {
			continue
		}
		v := a[i]
		if !set {
			f = v
			min = i
			set = true
			continue
		}
		if v < f {
			min = i
			f = v
		}
	}
	return min
}

func ArgminMaskedU16(a []uint16, mask []bool) int {
	var set bool
	var f uint16
	var min int
	for i := range a {
		if mask[i] {
			continue
		}
		v := a[i]
		if !set {
			f = v
			min = i
			set = true
			continue
		}
		if v < f {
			min = i
			f = v
		}
	}
	return min
}

func ArgminMaskedU32(a []uint32, mask []bool) int {
	var set bool
	var f uint32
	var min int
	for i := range a {
		if mask[i] {
			continue
		}
		v := a[i]
		if !set {
			f = v
			min = i
			set = true
			continue
		}
		if v < f {
			min = i
			f = v
		}
	}
	return min
}

func ArgminMaskedU64(a []uint64, mask []bool) int {
	var set bool
	var f uint64
	var min int
	for i := range a {
		if mask[i] {
			continue
		}
		v := a[i]
		if !set {
			f = v
			min = i
			set = true
			continue
		}
		if v < f {
			min = i
			f = v
		}
	}
	return min
}

func ArgminMaskedF32(a []float32, mask []bool) int {
	var set bool
	var f float32
	var min int
	for i := range a {
		if mask[i] {
			continue
		}
		v := a[i]
		if !set {
			f = v
			min = i
			set = true
			continue
		}
		if math32.IsNaN(v) || math32.IsInf(v, -1) {
			min = i
			return min
		}
		if v < f {
			min = i
			f = v
		}
	}
	return min
}

func ArgminMaskedF64(a []float64, mask []bool) int {
	var set bool
	var f float64
	var min int
	for i := range a {
		if mask[i] {
			continue
		}
		v := a[i]
		if !set {
			f = v
			min = i
			set = true
			continue
		}
		if math.IsNaN(v) || math.IsInf(v, -1) {
			min = i
			return min
		}
		if v < f {
			min = i
			f = v
		}
	}
	return min
}

func ArgminMaskedStr(a []string, mask []bool) int {
	var set bool
	var f string
	var min int
	for i := range a {
		if mask[i] {
			continue
		}
		v := a[i]
		if !set {
			f = v
			min = i
			set = true
			continue
		}
		if v < f {
			min = i
			f = v
		}
	}
	return min
}
