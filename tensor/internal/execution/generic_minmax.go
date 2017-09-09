package execution

/*
GENERATED FILE. DO NOT EDIT
*/

func VecMinI(a, b []int) {
	a = a[:]
	b = b[:len(a)]
	for i, v := range a {
		bv := b[i]
		if bv < v {
			a[i] = bv
		}
	}
}
func VecMaxI(a, b []int) {
	a = a[:]
	b = b[:len(a)]
	for i, v := range a {
		bv := b[i]
		if bv > v {
			a[i] = bv
		}
	}
}
func VecMinI8(a, b []int8) {
	a = a[:]
	b = b[:len(a)]
	for i, v := range a {
		bv := b[i]
		if bv < v {
			a[i] = bv
		}
	}
}
func VecMaxI8(a, b []int8) {
	a = a[:]
	b = b[:len(a)]
	for i, v := range a {
		bv := b[i]
		if bv > v {
			a[i] = bv
		}
	}
}
func VecMinI16(a, b []int16) {
	a = a[:]
	b = b[:len(a)]
	for i, v := range a {
		bv := b[i]
		if bv < v {
			a[i] = bv
		}
	}
}
func VecMaxI16(a, b []int16) {
	a = a[:]
	b = b[:len(a)]
	for i, v := range a {
		bv := b[i]
		if bv > v {
			a[i] = bv
		}
	}
}
func VecMinI32(a, b []int32) {
	a = a[:]
	b = b[:len(a)]
	for i, v := range a {
		bv := b[i]
		if bv < v {
			a[i] = bv
		}
	}
}
func VecMaxI32(a, b []int32) {
	a = a[:]
	b = b[:len(a)]
	for i, v := range a {
		bv := b[i]
		if bv > v {
			a[i] = bv
		}
	}
}
func VecMinI64(a, b []int64) {
	a = a[:]
	b = b[:len(a)]
	for i, v := range a {
		bv := b[i]
		if bv < v {
			a[i] = bv
		}
	}
}
func VecMaxI64(a, b []int64) {
	a = a[:]
	b = b[:len(a)]
	for i, v := range a {
		bv := b[i]
		if bv > v {
			a[i] = bv
		}
	}
}
func VecMinU(a, b []uint) {
	a = a[:]
	b = b[:len(a)]
	for i, v := range a {
		bv := b[i]
		if bv < v {
			a[i] = bv
		}
	}
}
func VecMaxU(a, b []uint) {
	a = a[:]
	b = b[:len(a)]
	for i, v := range a {
		bv := b[i]
		if bv > v {
			a[i] = bv
		}
	}
}
func VecMinU8(a, b []uint8) {
	a = a[:]
	b = b[:len(a)]
	for i, v := range a {
		bv := b[i]
		if bv < v {
			a[i] = bv
		}
	}
}
func VecMaxU8(a, b []uint8) {
	a = a[:]
	b = b[:len(a)]
	for i, v := range a {
		bv := b[i]
		if bv > v {
			a[i] = bv
		}
	}
}
func VecMinU16(a, b []uint16) {
	a = a[:]
	b = b[:len(a)]
	for i, v := range a {
		bv := b[i]
		if bv < v {
			a[i] = bv
		}
	}
}
func VecMaxU16(a, b []uint16) {
	a = a[:]
	b = b[:len(a)]
	for i, v := range a {
		bv := b[i]
		if bv > v {
			a[i] = bv
		}
	}
}
func VecMinU32(a, b []uint32) {
	a = a[:]
	b = b[:len(a)]
	for i, v := range a {
		bv := b[i]
		if bv < v {
			a[i] = bv
		}
	}
}
func VecMaxU32(a, b []uint32) {
	a = a[:]
	b = b[:len(a)]
	for i, v := range a {
		bv := b[i]
		if bv > v {
			a[i] = bv
		}
	}
}
func VecMinU64(a, b []uint64) {
	a = a[:]
	b = b[:len(a)]
	for i, v := range a {
		bv := b[i]
		if bv < v {
			a[i] = bv
		}
	}
}
func VecMaxU64(a, b []uint64) {
	a = a[:]
	b = b[:len(a)]
	for i, v := range a {
		bv := b[i]
		if bv > v {
			a[i] = bv
		}
	}
}
func VecMinF32(a, b []float32) {
	a = a[:]
	b = b[:len(a)]
	for i, v := range a {
		bv := b[i]
		if bv < v {
			a[i] = bv
		}
	}
}
func VecMaxF32(a, b []float32) {
	a = a[:]
	b = b[:len(a)]
	for i, v := range a {
		bv := b[i]
		if bv > v {
			a[i] = bv
		}
	}
}
func VecMinF64(a, b []float64) {
	a = a[:]
	b = b[:len(a)]
	for i, v := range a {
		bv := b[i]
		if bv < v {
			a[i] = bv
		}
	}
}
func VecMaxF64(a, b []float64) {
	a = a[:]
	b = b[:len(a)]
	for i, v := range a {
		bv := b[i]
		if bv > v {
			a[i] = bv
		}
	}
}
func VecMinStr(a, b []string) {
	a = a[:]
	b = b[:len(a)]
	for i, v := range a {
		bv := b[i]
		if bv < v {
			a[i] = bv
		}
	}
}
func VecMaxStr(a, b []string) {
	a = a[:]
	b = b[:len(a)]
	for i, v := range a {
		bv := b[i]
		if bv > v {
			a[i] = bv
		}
	}
}
func MinI(a, b int) (c int) {
	if a < b {
		return a
	}
	return b
}

func MaxI(a, b int) (c int) {
	if a > b {
		return a
	}
	return b
}
func MinI8(a, b int8) (c int8) {
	if a < b {
		return a
	}
	return b
}

func MaxI8(a, b int8) (c int8) {
	if a > b {
		return a
	}
	return b
}
func MinI16(a, b int16) (c int16) {
	if a < b {
		return a
	}
	return b
}

func MaxI16(a, b int16) (c int16) {
	if a > b {
		return a
	}
	return b
}
func MinI32(a, b int32) (c int32) {
	if a < b {
		return a
	}
	return b
}

func MaxI32(a, b int32) (c int32) {
	if a > b {
		return a
	}
	return b
}
func MinI64(a, b int64) (c int64) {
	if a < b {
		return a
	}
	return b
}

func MaxI64(a, b int64) (c int64) {
	if a > b {
		return a
	}
	return b
}
func MinU(a, b uint) (c uint) {
	if a < b {
		return a
	}
	return b
}

func MaxU(a, b uint) (c uint) {
	if a > b {
		return a
	}
	return b
}
func MinU8(a, b uint8) (c uint8) {
	if a < b {
		return a
	}
	return b
}

func MaxU8(a, b uint8) (c uint8) {
	if a > b {
		return a
	}
	return b
}
func MinU16(a, b uint16) (c uint16) {
	if a < b {
		return a
	}
	return b
}

func MaxU16(a, b uint16) (c uint16) {
	if a > b {
		return a
	}
	return b
}
func MinU32(a, b uint32) (c uint32) {
	if a < b {
		return a
	}
	return b
}

func MaxU32(a, b uint32) (c uint32) {
	if a > b {
		return a
	}
	return b
}
func MinU64(a, b uint64) (c uint64) {
	if a < b {
		return a
	}
	return b
}

func MaxU64(a, b uint64) (c uint64) {
	if a > b {
		return a
	}
	return b
}
func MinF32(a, b float32) (c float32) {
	if a < b {
		return a
	}
	return b
}

func MaxF32(a, b float32) (c float32) {
	if a > b {
		return a
	}
	return b
}
func MinF64(a, b float64) (c float64) {
	if a < b {
		return a
	}
	return b
}

func MaxF64(a, b float64) (c float64) {
	if a > b {
		return a
	}
	return b
}
func MinStr(a, b string) (c string) {
	if a < b {
		return a
	}
	return b
}

func MaxStr(a, b string) (c string) {
	if a > b {
		return a
	}
	return b
}
