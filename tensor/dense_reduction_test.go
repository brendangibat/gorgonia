package tensor

import (
	"testing"

	"github.com/chewxy/gorgonia/tensor/internal/execution"
	"github.com/stretchr/testify/assert"
)

/*
GENERATED FILE. DO NOT EDIT
*/

var denseReductionTests = []struct {
	of   Dtype
	fn   interface{}
	def  interface{}
	axis int

	correct      interface{}
	correctShape Shape
}{
	// int
	{Int, execution.AddI, int(0), 0, []int{6, 8, 10, 12, 14, 16}, Shape{3, 2}},
	{Int, execution.AddI, int(0), 1, []int{6, 9, 24, 27}, Shape{2, 2}},
	{Int, execution.AddI, int(0), 2, []int{1, 5, 9, 13, 17, 21}, Shape{2, 3}},
	// int8
	{Int8, execution.AddI8, int8(0), 0, []int8{6, 8, 10, 12, 14, 16}, Shape{3, 2}},
	{Int8, execution.AddI8, int8(0), 1, []int8{6, 9, 24, 27}, Shape{2, 2}},
	{Int8, execution.AddI8, int8(0), 2, []int8{1, 5, 9, 13, 17, 21}, Shape{2, 3}},
	// int16
	{Int16, execution.AddI16, int16(0), 0, []int16{6, 8, 10, 12, 14, 16}, Shape{3, 2}},
	{Int16, execution.AddI16, int16(0), 1, []int16{6, 9, 24, 27}, Shape{2, 2}},
	{Int16, execution.AddI16, int16(0), 2, []int16{1, 5, 9, 13, 17, 21}, Shape{2, 3}},
	// int32
	{Int32, execution.AddI32, int32(0), 0, []int32{6, 8, 10, 12, 14, 16}, Shape{3, 2}},
	{Int32, execution.AddI32, int32(0), 1, []int32{6, 9, 24, 27}, Shape{2, 2}},
	{Int32, execution.AddI32, int32(0), 2, []int32{1, 5, 9, 13, 17, 21}, Shape{2, 3}},
	// int64
	{Int64, execution.AddI64, int64(0), 0, []int64{6, 8, 10, 12, 14, 16}, Shape{3, 2}},
	{Int64, execution.AddI64, int64(0), 1, []int64{6, 9, 24, 27}, Shape{2, 2}},
	{Int64, execution.AddI64, int64(0), 2, []int64{1, 5, 9, 13, 17, 21}, Shape{2, 3}},
	// uint
	{Uint, execution.AddU, uint(0), 0, []uint{6, 8, 10, 12, 14, 16}, Shape{3, 2}},
	{Uint, execution.AddU, uint(0), 1, []uint{6, 9, 24, 27}, Shape{2, 2}},
	{Uint, execution.AddU, uint(0), 2, []uint{1, 5, 9, 13, 17, 21}, Shape{2, 3}},
	// uint8
	{Uint8, execution.AddU8, uint8(0), 0, []uint8{6, 8, 10, 12, 14, 16}, Shape{3, 2}},
	{Uint8, execution.AddU8, uint8(0), 1, []uint8{6, 9, 24, 27}, Shape{2, 2}},
	{Uint8, execution.AddU8, uint8(0), 2, []uint8{1, 5, 9, 13, 17, 21}, Shape{2, 3}},
	// uint16
	{Uint16, execution.AddU16, uint16(0), 0, []uint16{6, 8, 10, 12, 14, 16}, Shape{3, 2}},
	{Uint16, execution.AddU16, uint16(0), 1, []uint16{6, 9, 24, 27}, Shape{2, 2}},
	{Uint16, execution.AddU16, uint16(0), 2, []uint16{1, 5, 9, 13, 17, 21}, Shape{2, 3}},
	// uint32
	{Uint32, execution.AddU32, uint32(0), 0, []uint32{6, 8, 10, 12, 14, 16}, Shape{3, 2}},
	{Uint32, execution.AddU32, uint32(0), 1, []uint32{6, 9, 24, 27}, Shape{2, 2}},
	{Uint32, execution.AddU32, uint32(0), 2, []uint32{1, 5, 9, 13, 17, 21}, Shape{2, 3}},
	// uint64
	{Uint64, execution.AddU64, uint64(0), 0, []uint64{6, 8, 10, 12, 14, 16}, Shape{3, 2}},
	{Uint64, execution.AddU64, uint64(0), 1, []uint64{6, 9, 24, 27}, Shape{2, 2}},
	{Uint64, execution.AddU64, uint64(0), 2, []uint64{1, 5, 9, 13, 17, 21}, Shape{2, 3}},
	// float32
	{Float32, execution.AddF32, float32(0), 0, []float32{6, 8, 10, 12, 14, 16}, Shape{3, 2}},
	{Float32, execution.AddF32, float32(0), 1, []float32{6, 9, 24, 27}, Shape{2, 2}},
	{Float32, execution.AddF32, float32(0), 2, []float32{1, 5, 9, 13, 17, 21}, Shape{2, 3}},
	// float64
	{Float64, execution.AddF64, float64(0), 0, []float64{6, 8, 10, 12, 14, 16}, Shape{3, 2}},
	{Float64, execution.AddF64, float64(0), 1, []float64{6, 9, 24, 27}, Shape{2, 2}},
	{Float64, execution.AddF64, float64(0), 2, []float64{1, 5, 9, 13, 17, 21}, Shape{2, 3}},
	// complex64
	{Complex64, execution.AddC64, complex64(0), 0, []complex64{6, 8, 10, 12, 14, 16}, Shape{3, 2}},
	{Complex64, execution.AddC64, complex64(0), 1, []complex64{6, 9, 24, 27}, Shape{2, 2}},
	{Complex64, execution.AddC64, complex64(0), 2, []complex64{1, 5, 9, 13, 17, 21}, Shape{2, 3}},
	// complex128
	{Complex128, execution.AddC128, complex128(0), 0, []complex128{6, 8, 10, 12, 14, 16}, Shape{3, 2}},
	{Complex128, execution.AddC128, complex128(0), 1, []complex128{6, 9, 24, 27}, Shape{2, 2}},
	{Complex128, execution.AddC128, complex128(0), 2, []complex128{1, 5, 9, 13, 17, 21}, Shape{2, 3}},
}

func TestDense_Reduce(t *testing.T) {
	assert := assert.New(t)
	for _, drt := range denseReductionTests {
		T := New(WithShape(2, 3, 2), WithBacking(Range(drt.of, 0, 2*3*2)))
		T2, err := T.Reduce(drt.fn, drt.axis, drt.def)
		if err != nil {
			t.Error(err)
			continue
		}
		assert.True(drt.correctShape.Eq(T2.Shape()))
		assert.Equal(drt.correct, T2.Data())

		// stupids:
		_, err = T.Reduce(drt.fn, 1000, drt.def)
		assert.NotNil(err)

		// wrong function type
		var f interface{}
		f = func(a, b float64) float64 { return 0 }
		if drt.of == Float64 {
			f = func(a, b int) int { return 0 }
		}

		_, err = T.Reduce(f, 0, drt.correct)
		assert.NotNil(err)

		// wrong default value type
		var def2 interface{}
		def2 = 3.14
		if drt.of == Float64 {
			def2 = int(1)
		}

		_, err = T.Reduce(drt.fn, 3, def2) // only last axis requires a default value
		assert.NotNil(err)
	}
}

var sumTests = []struct {
	name  string
	of    Dtype
	shape Shape
	along []int

	correctShape Shape
	correct      interface{}
}{
	{"common case: T.Sum() for int", Int, Shape{2, 3}, []int{}, ScalarShape(), int(15)},
	{"A.Sum(0) for int", Int, Shape{2, 3}, []int{0}, Shape{3}, []int{3, 5, 7}},
	{"A.Sum(1) for int", Int, Shape{2, 3}, []int{1}, Shape{2}, []int{3, 12}},
	{"A.Sum(0,1) for int", Int, Shape{2, 3}, []int{0, 1}, ScalarShape(), int(15)},
	{"A.Sum(1,0) for int", Int, Shape{2, 3}, []int{1, 0}, ScalarShape(), int(15)},
	{"3T.Sum(1,2) for int", Int, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []int{66, 210}},
	{"common case: T.Sum() for int8", Int8, Shape{2, 3}, []int{}, ScalarShape(), int8(15)},
	{"A.Sum(0) for int8", Int8, Shape{2, 3}, []int{0}, Shape{3}, []int8{3, 5, 7}},
	{"A.Sum(1) for int8", Int8, Shape{2, 3}, []int{1}, Shape{2}, []int8{3, 12}},
	{"A.Sum(0,1) for int8", Int8, Shape{2, 3}, []int{0, 1}, ScalarShape(), int8(15)},
	{"A.Sum(1,0) for int8", Int8, Shape{2, 3}, []int{1, 0}, ScalarShape(), int8(15)},
	{"3T.Sum(1,2) for int8", Int8, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []int8{66, -46}},
	{"common case: T.Sum() for int16", Int16, Shape{2, 3}, []int{}, ScalarShape(), int16(15)},
	{"A.Sum(0) for int16", Int16, Shape{2, 3}, []int{0}, Shape{3}, []int16{3, 5, 7}},
	{"A.Sum(1) for int16", Int16, Shape{2, 3}, []int{1}, Shape{2}, []int16{3, 12}},
	{"A.Sum(0,1) for int16", Int16, Shape{2, 3}, []int{0, 1}, ScalarShape(), int16(15)},
	{"A.Sum(1,0) for int16", Int16, Shape{2, 3}, []int{1, 0}, ScalarShape(), int16(15)},
	{"3T.Sum(1,2) for int16", Int16, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []int16{66, 210}},
	{"common case: T.Sum() for int32", Int32, Shape{2, 3}, []int{}, ScalarShape(), int32(15)},
	{"A.Sum(0) for int32", Int32, Shape{2, 3}, []int{0}, Shape{3}, []int32{3, 5, 7}},
	{"A.Sum(1) for int32", Int32, Shape{2, 3}, []int{1}, Shape{2}, []int32{3, 12}},
	{"A.Sum(0,1) for int32", Int32, Shape{2, 3}, []int{0, 1}, ScalarShape(), int32(15)},
	{"A.Sum(1,0) for int32", Int32, Shape{2, 3}, []int{1, 0}, ScalarShape(), int32(15)},
	{"3T.Sum(1,2) for int32", Int32, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []int32{66, 210}},
	{"common case: T.Sum() for int64", Int64, Shape{2, 3}, []int{}, ScalarShape(), int64(15)},
	{"A.Sum(0) for int64", Int64, Shape{2, 3}, []int{0}, Shape{3}, []int64{3, 5, 7}},
	{"A.Sum(1) for int64", Int64, Shape{2, 3}, []int{1}, Shape{2}, []int64{3, 12}},
	{"A.Sum(0,1) for int64", Int64, Shape{2, 3}, []int{0, 1}, ScalarShape(), int64(15)},
	{"A.Sum(1,0) for int64", Int64, Shape{2, 3}, []int{1, 0}, ScalarShape(), int64(15)},
	{"3T.Sum(1,2) for int64", Int64, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []int64{66, 210}},
	{"common case: T.Sum() for uint", Uint, Shape{2, 3}, []int{}, ScalarShape(), uint(15)},
	{"A.Sum(0) for uint", Uint, Shape{2, 3}, []int{0}, Shape{3}, []uint{3, 5, 7}},
	{"A.Sum(1) for uint", Uint, Shape{2, 3}, []int{1}, Shape{2}, []uint{3, 12}},
	{"A.Sum(0,1) for uint", Uint, Shape{2, 3}, []int{0, 1}, ScalarShape(), uint(15)},
	{"A.Sum(1,0) for uint", Uint, Shape{2, 3}, []int{1, 0}, ScalarShape(), uint(15)},
	{"3T.Sum(1,2) for uint", Uint, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []uint{66, 210}},
	{"common case: T.Sum() for uint8", Uint8, Shape{2, 3}, []int{}, ScalarShape(), uint8(15)},
	{"A.Sum(0) for uint8", Uint8, Shape{2, 3}, []int{0}, Shape{3}, []uint8{3, 5, 7}},
	{"A.Sum(1) for uint8", Uint8, Shape{2, 3}, []int{1}, Shape{2}, []uint8{3, 12}},
	{"A.Sum(0,1) for uint8", Uint8, Shape{2, 3}, []int{0, 1}, ScalarShape(), uint8(15)},
	{"A.Sum(1,0) for uint8", Uint8, Shape{2, 3}, []int{1, 0}, ScalarShape(), uint8(15)},
	{"3T.Sum(1,2) for uint8", Uint8, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []uint8{66, 210}},
	{"common case: T.Sum() for uint16", Uint16, Shape{2, 3}, []int{}, ScalarShape(), uint16(15)},
	{"A.Sum(0) for uint16", Uint16, Shape{2, 3}, []int{0}, Shape{3}, []uint16{3, 5, 7}},
	{"A.Sum(1) for uint16", Uint16, Shape{2, 3}, []int{1}, Shape{2}, []uint16{3, 12}},
	{"A.Sum(0,1) for uint16", Uint16, Shape{2, 3}, []int{0, 1}, ScalarShape(), uint16(15)},
	{"A.Sum(1,0) for uint16", Uint16, Shape{2, 3}, []int{1, 0}, ScalarShape(), uint16(15)},
	{"3T.Sum(1,2) for uint16", Uint16, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []uint16{66, 210}},
	{"common case: T.Sum() for uint32", Uint32, Shape{2, 3}, []int{}, ScalarShape(), uint32(15)},
	{"A.Sum(0) for uint32", Uint32, Shape{2, 3}, []int{0}, Shape{3}, []uint32{3, 5, 7}},
	{"A.Sum(1) for uint32", Uint32, Shape{2, 3}, []int{1}, Shape{2}, []uint32{3, 12}},
	{"A.Sum(0,1) for uint32", Uint32, Shape{2, 3}, []int{0, 1}, ScalarShape(), uint32(15)},
	{"A.Sum(1,0) for uint32", Uint32, Shape{2, 3}, []int{1, 0}, ScalarShape(), uint32(15)},
	{"3T.Sum(1,2) for uint32", Uint32, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []uint32{66, 210}},
	{"common case: T.Sum() for uint64", Uint64, Shape{2, 3}, []int{}, ScalarShape(), uint64(15)},
	{"A.Sum(0) for uint64", Uint64, Shape{2, 3}, []int{0}, Shape{3}, []uint64{3, 5, 7}},
	{"A.Sum(1) for uint64", Uint64, Shape{2, 3}, []int{1}, Shape{2}, []uint64{3, 12}},
	{"A.Sum(0,1) for uint64", Uint64, Shape{2, 3}, []int{0, 1}, ScalarShape(), uint64(15)},
	{"A.Sum(1,0) for uint64", Uint64, Shape{2, 3}, []int{1, 0}, ScalarShape(), uint64(15)},
	{"3T.Sum(1,2) for uint64", Uint64, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []uint64{66, 210}},
	{"common case: T.Sum() for float32", Float32, Shape{2, 3}, []int{}, ScalarShape(), float32(15)},
	{"A.Sum(0) for float32", Float32, Shape{2, 3}, []int{0}, Shape{3}, []float32{3, 5, 7}},
	{"A.Sum(1) for float32", Float32, Shape{2, 3}, []int{1}, Shape{2}, []float32{3, 12}},
	{"A.Sum(0,1) for float32", Float32, Shape{2, 3}, []int{0, 1}, ScalarShape(), float32(15)},
	{"A.Sum(1,0) for float32", Float32, Shape{2, 3}, []int{1, 0}, ScalarShape(), float32(15)},
	{"3T.Sum(1,2) for float32", Float32, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []float32{66, 210}},
	{"common case: T.Sum() for float64", Float64, Shape{2, 3}, []int{}, ScalarShape(), float64(15)},
	{"A.Sum(0) for float64", Float64, Shape{2, 3}, []int{0}, Shape{3}, []float64{3, 5, 7}},
	{"A.Sum(1) for float64", Float64, Shape{2, 3}, []int{1}, Shape{2}, []float64{3, 12}},
	{"A.Sum(0,1) for float64", Float64, Shape{2, 3}, []int{0, 1}, ScalarShape(), float64(15)},
	{"A.Sum(1,0) for float64", Float64, Shape{2, 3}, []int{1, 0}, ScalarShape(), float64(15)},
	{"3T.Sum(1,2) for float64", Float64, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []float64{66, 210}},
	{"common case: T.Sum() for complex64", Complex64, Shape{2, 3}, []int{}, ScalarShape(), complex64(15)},
	{"A.Sum(0) for complex64", Complex64, Shape{2, 3}, []int{0}, Shape{3}, []complex64{3, 5, 7}},
	{"A.Sum(1) for complex64", Complex64, Shape{2, 3}, []int{1}, Shape{2}, []complex64{3, 12}},
	{"A.Sum(0,1) for complex64", Complex64, Shape{2, 3}, []int{0, 1}, ScalarShape(), complex64(15)},
	{"A.Sum(1,0) for complex64", Complex64, Shape{2, 3}, []int{1, 0}, ScalarShape(), complex64(15)},
	{"3T.Sum(1,2) for complex64", Complex64, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []complex64{66, 210}},
	{"common case: T.Sum() for complex128", Complex128, Shape{2, 3}, []int{}, ScalarShape(), complex128(15)},
	{"A.Sum(0) for complex128", Complex128, Shape{2, 3}, []int{0}, Shape{3}, []complex128{3, 5, 7}},
	{"A.Sum(1) for complex128", Complex128, Shape{2, 3}, []int{1}, Shape{2}, []complex128{3, 12}},
	{"A.Sum(0,1) for complex128", Complex128, Shape{2, 3}, []int{0, 1}, ScalarShape(), complex128(15)},
	{"A.Sum(1,0) for complex128", Complex128, Shape{2, 3}, []int{1, 0}, ScalarShape(), complex128(15)},
	{"3T.Sum(1,2) for complex128", Complex128, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []complex128{66, 210}},
}

func TestDense_Sum(t *testing.T) {
	assert := assert.New(t)
	var T, T2 *Dense
	var err error

	for _, sts := range sumTests {
		T = New(WithShape(sts.shape...), WithBacking(Range(sts.of, 0, sts.shape.TotalSize())))
		if T2, err = T.Sum(sts.along...); err != nil {
			t.Error(err)
			continue
		}
		assert.True(sts.correctShape.Eq(T2.Shape()))
		assert.Equal(sts.correct, T2.Data())
	}

	// idiots
	_, err = T.Sum(1000)
	assert.NotNil(err)
}

var maxTests = []struct {
	name  string
	of    Dtype
	shape Shape
	along []int

	correctShape Shape
	correct      interface{}
}{
	{"common case: T.Max() for int", Int, Shape{2, 3}, []int{}, ScalarShape(), int(5)},
	{"A.Max(0)", Int, Shape{2, 3}, []int{0}, Shape{3}, []int{3, 4, 5}},
	{"A.Max(1)", Int, Shape{2, 3}, []int{1}, Shape{2}, []int{2, 5}},
	{"A.Max(0,1)", Int, Shape{2, 3}, []int{0, 1}, ScalarShape(), int(5)},
	{"A.Max(1,0)", Int, Shape{2, 3}, []int{1, 0}, ScalarShape(), int(5)},
	{"3T.Max(1,2)", Int, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []int{11, 23}},
	{"common case: T.Max() for int8", Int8, Shape{2, 3}, []int{}, ScalarShape(), int8(5)},
	{"A.Max(0)", Int8, Shape{2, 3}, []int{0}, Shape{3}, []int8{3, 4, 5}},
	{"A.Max(1)", Int8, Shape{2, 3}, []int{1}, Shape{2}, []int8{2, 5}},
	{"A.Max(0,1)", Int8, Shape{2, 3}, []int{0, 1}, ScalarShape(), int8(5)},
	{"A.Max(1,0)", Int8, Shape{2, 3}, []int{1, 0}, ScalarShape(), int8(5)},
	{"3T.Max(1,2)", Int8, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []int8{11, 23}},
	{"common case: T.Max() for int16", Int16, Shape{2, 3}, []int{}, ScalarShape(), int16(5)},
	{"A.Max(0)", Int16, Shape{2, 3}, []int{0}, Shape{3}, []int16{3, 4, 5}},
	{"A.Max(1)", Int16, Shape{2, 3}, []int{1}, Shape{2}, []int16{2, 5}},
	{"A.Max(0,1)", Int16, Shape{2, 3}, []int{0, 1}, ScalarShape(), int16(5)},
	{"A.Max(1,0)", Int16, Shape{2, 3}, []int{1, 0}, ScalarShape(), int16(5)},
	{"3T.Max(1,2)", Int16, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []int16{11, 23}},
	{"common case: T.Max() for int32", Int32, Shape{2, 3}, []int{}, ScalarShape(), int32(5)},
	{"A.Max(0)", Int32, Shape{2, 3}, []int{0}, Shape{3}, []int32{3, 4, 5}},
	{"A.Max(1)", Int32, Shape{2, 3}, []int{1}, Shape{2}, []int32{2, 5}},
	{"A.Max(0,1)", Int32, Shape{2, 3}, []int{0, 1}, ScalarShape(), int32(5)},
	{"A.Max(1,0)", Int32, Shape{2, 3}, []int{1, 0}, ScalarShape(), int32(5)},
	{"3T.Max(1,2)", Int32, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []int32{11, 23}},
	{"common case: T.Max() for int64", Int64, Shape{2, 3}, []int{}, ScalarShape(), int64(5)},
	{"A.Max(0)", Int64, Shape{2, 3}, []int{0}, Shape{3}, []int64{3, 4, 5}},
	{"A.Max(1)", Int64, Shape{2, 3}, []int{1}, Shape{2}, []int64{2, 5}},
	{"A.Max(0,1)", Int64, Shape{2, 3}, []int{0, 1}, ScalarShape(), int64(5)},
	{"A.Max(1,0)", Int64, Shape{2, 3}, []int{1, 0}, ScalarShape(), int64(5)},
	{"3T.Max(1,2)", Int64, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []int64{11, 23}},
	{"common case: T.Max() for uint", Uint, Shape{2, 3}, []int{}, ScalarShape(), uint(5)},
	{"A.Max(0)", Uint, Shape{2, 3}, []int{0}, Shape{3}, []uint{3, 4, 5}},
	{"A.Max(1)", Uint, Shape{2, 3}, []int{1}, Shape{2}, []uint{2, 5}},
	{"A.Max(0,1)", Uint, Shape{2, 3}, []int{0, 1}, ScalarShape(), uint(5)},
	{"A.Max(1,0)", Uint, Shape{2, 3}, []int{1, 0}, ScalarShape(), uint(5)},
	{"3T.Max(1,2)", Uint, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []uint{11, 23}},
	{"common case: T.Max() for uint8", Uint8, Shape{2, 3}, []int{}, ScalarShape(), uint8(5)},
	{"A.Max(0)", Uint8, Shape{2, 3}, []int{0}, Shape{3}, []uint8{3, 4, 5}},
	{"A.Max(1)", Uint8, Shape{2, 3}, []int{1}, Shape{2}, []uint8{2, 5}},
	{"A.Max(0,1)", Uint8, Shape{2, 3}, []int{0, 1}, ScalarShape(), uint8(5)},
	{"A.Max(1,0)", Uint8, Shape{2, 3}, []int{1, 0}, ScalarShape(), uint8(5)},
	{"3T.Max(1,2)", Uint8, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []uint8{11, 23}},
	{"common case: T.Max() for uint16", Uint16, Shape{2, 3}, []int{}, ScalarShape(), uint16(5)},
	{"A.Max(0)", Uint16, Shape{2, 3}, []int{0}, Shape{3}, []uint16{3, 4, 5}},
	{"A.Max(1)", Uint16, Shape{2, 3}, []int{1}, Shape{2}, []uint16{2, 5}},
	{"A.Max(0,1)", Uint16, Shape{2, 3}, []int{0, 1}, ScalarShape(), uint16(5)},
	{"A.Max(1,0)", Uint16, Shape{2, 3}, []int{1, 0}, ScalarShape(), uint16(5)},
	{"3T.Max(1,2)", Uint16, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []uint16{11, 23}},
	{"common case: T.Max() for uint32", Uint32, Shape{2, 3}, []int{}, ScalarShape(), uint32(5)},
	{"A.Max(0)", Uint32, Shape{2, 3}, []int{0}, Shape{3}, []uint32{3, 4, 5}},
	{"A.Max(1)", Uint32, Shape{2, 3}, []int{1}, Shape{2}, []uint32{2, 5}},
	{"A.Max(0,1)", Uint32, Shape{2, 3}, []int{0, 1}, ScalarShape(), uint32(5)},
	{"A.Max(1,0)", Uint32, Shape{2, 3}, []int{1, 0}, ScalarShape(), uint32(5)},
	{"3T.Max(1,2)", Uint32, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []uint32{11, 23}},
	{"common case: T.Max() for uint64", Uint64, Shape{2, 3}, []int{}, ScalarShape(), uint64(5)},
	{"A.Max(0)", Uint64, Shape{2, 3}, []int{0}, Shape{3}, []uint64{3, 4, 5}},
	{"A.Max(1)", Uint64, Shape{2, 3}, []int{1}, Shape{2}, []uint64{2, 5}},
	{"A.Max(0,1)", Uint64, Shape{2, 3}, []int{0, 1}, ScalarShape(), uint64(5)},
	{"A.Max(1,0)", Uint64, Shape{2, 3}, []int{1, 0}, ScalarShape(), uint64(5)},
	{"3T.Max(1,2)", Uint64, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []uint64{11, 23}},
	{"common case: T.Max() for float32", Float32, Shape{2, 3}, []int{}, ScalarShape(), float32(5)},
	{"A.Max(0)", Float32, Shape{2, 3}, []int{0}, Shape{3}, []float32{3, 4, 5}},
	{"A.Max(1)", Float32, Shape{2, 3}, []int{1}, Shape{2}, []float32{2, 5}},
	{"A.Max(0,1)", Float32, Shape{2, 3}, []int{0, 1}, ScalarShape(), float32(5)},
	{"A.Max(1,0)", Float32, Shape{2, 3}, []int{1, 0}, ScalarShape(), float32(5)},
	{"3T.Max(1,2)", Float32, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []float32{11, 23}},
	{"common case: T.Max() for float64", Float64, Shape{2, 3}, []int{}, ScalarShape(), float64(5)},
	{"A.Max(0)", Float64, Shape{2, 3}, []int{0}, Shape{3}, []float64{3, 4, 5}},
	{"A.Max(1)", Float64, Shape{2, 3}, []int{1}, Shape{2}, []float64{2, 5}},
	{"A.Max(0,1)", Float64, Shape{2, 3}, []int{0, 1}, ScalarShape(), float64(5)},
	{"A.Max(1,0)", Float64, Shape{2, 3}, []int{1, 0}, ScalarShape(), float64(5)},
	{"3T.Max(1,2)", Float64, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []float64{11, 23}},
}

func TestDense_Max(t *testing.T) {
	assert := assert.New(t)
	var T, T2 *Dense
	var err error

	for _, mts := range maxTests {
		T = New(WithShape(mts.shape...), WithBacking(Range(mts.of, 0, mts.shape.TotalSize())))
		if T2, err = T.Max(mts.along...); err != nil {
			t.Error(err)
			continue
		}
		assert.True(mts.correctShape.Eq(T2.Shape()))
		assert.Equal(mts.correct, T2.Data())
	}
	/* IDIOT TESTING TIME */
	_, err = T.Max(1000)
	assert.NotNil(err)
}

var minTests = []struct {
	name  string
	of    Dtype
	shape Shape
	along []int

	correctShape Shape
	correct      interface{}
}{
	{"common case: T.Min() for int", Int, Shape{2, 3}, []int{}, ScalarShape(), int(0)},
	{"A.Min(0)", Int, Shape{2, 3}, []int{0}, Shape{3}, []int{0, 1, 2}},
	{"A.Min(1)", Int, Shape{2, 3}, []int{1}, Shape{2}, []int{0, 3}},
	{"A.Min(0,1)", Int, Shape{2, 3}, []int{0, 1}, ScalarShape(), int(0)},
	{"A.Min(1,0)", Int, Shape{2, 3}, []int{1, 0}, ScalarShape(), int(0)},
	{"3T.Min(1,2)", Int, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []int{0, 12}},
	{"common case: T.Min() for int8", Int8, Shape{2, 3}, []int{}, ScalarShape(), int8(0)},
	{"A.Min(0)", Int8, Shape{2, 3}, []int{0}, Shape{3}, []int8{0, 1, 2}},
	{"A.Min(1)", Int8, Shape{2, 3}, []int{1}, Shape{2}, []int8{0, 3}},
	{"A.Min(0,1)", Int8, Shape{2, 3}, []int{0, 1}, ScalarShape(), int8(0)},
	{"A.Min(1,0)", Int8, Shape{2, 3}, []int{1, 0}, ScalarShape(), int8(0)},
	{"3T.Min(1,2)", Int8, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []int8{0, 12}},
	{"common case: T.Min() for int16", Int16, Shape{2, 3}, []int{}, ScalarShape(), int16(0)},
	{"A.Min(0)", Int16, Shape{2, 3}, []int{0}, Shape{3}, []int16{0, 1, 2}},
	{"A.Min(1)", Int16, Shape{2, 3}, []int{1}, Shape{2}, []int16{0, 3}},
	{"A.Min(0,1)", Int16, Shape{2, 3}, []int{0, 1}, ScalarShape(), int16(0)},
	{"A.Min(1,0)", Int16, Shape{2, 3}, []int{1, 0}, ScalarShape(), int16(0)},
	{"3T.Min(1,2)", Int16, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []int16{0, 12}},
	{"common case: T.Min() for int32", Int32, Shape{2, 3}, []int{}, ScalarShape(), int32(0)},
	{"A.Min(0)", Int32, Shape{2, 3}, []int{0}, Shape{3}, []int32{0, 1, 2}},
	{"A.Min(1)", Int32, Shape{2, 3}, []int{1}, Shape{2}, []int32{0, 3}},
	{"A.Min(0,1)", Int32, Shape{2, 3}, []int{0, 1}, ScalarShape(), int32(0)},
	{"A.Min(1,0)", Int32, Shape{2, 3}, []int{1, 0}, ScalarShape(), int32(0)},
	{"3T.Min(1,2)", Int32, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []int32{0, 12}},
	{"common case: T.Min() for int64", Int64, Shape{2, 3}, []int{}, ScalarShape(), int64(0)},
	{"A.Min(0)", Int64, Shape{2, 3}, []int{0}, Shape{3}, []int64{0, 1, 2}},
	{"A.Min(1)", Int64, Shape{2, 3}, []int{1}, Shape{2}, []int64{0, 3}},
	{"A.Min(0,1)", Int64, Shape{2, 3}, []int{0, 1}, ScalarShape(), int64(0)},
	{"A.Min(1,0)", Int64, Shape{2, 3}, []int{1, 0}, ScalarShape(), int64(0)},
	{"3T.Min(1,2)", Int64, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []int64{0, 12}},
	{"common case: T.Min() for uint", Uint, Shape{2, 3}, []int{}, ScalarShape(), uint(0)},
	{"A.Min(0)", Uint, Shape{2, 3}, []int{0}, Shape{3}, []uint{0, 1, 2}},
	{"A.Min(1)", Uint, Shape{2, 3}, []int{1}, Shape{2}, []uint{0, 3}},
	{"A.Min(0,1)", Uint, Shape{2, 3}, []int{0, 1}, ScalarShape(), uint(0)},
	{"A.Min(1,0)", Uint, Shape{2, 3}, []int{1, 0}, ScalarShape(), uint(0)},
	{"3T.Min(1,2)", Uint, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []uint{0, 12}},
	{"common case: T.Min() for uint8", Uint8, Shape{2, 3}, []int{}, ScalarShape(), uint8(0)},
	{"A.Min(0)", Uint8, Shape{2, 3}, []int{0}, Shape{3}, []uint8{0, 1, 2}},
	{"A.Min(1)", Uint8, Shape{2, 3}, []int{1}, Shape{2}, []uint8{0, 3}},
	{"A.Min(0,1)", Uint8, Shape{2, 3}, []int{0, 1}, ScalarShape(), uint8(0)},
	{"A.Min(1,0)", Uint8, Shape{2, 3}, []int{1, 0}, ScalarShape(), uint8(0)},
	{"3T.Min(1,2)", Uint8, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []uint8{0, 12}},
	{"common case: T.Min() for uint16", Uint16, Shape{2, 3}, []int{}, ScalarShape(), uint16(0)},
	{"A.Min(0)", Uint16, Shape{2, 3}, []int{0}, Shape{3}, []uint16{0, 1, 2}},
	{"A.Min(1)", Uint16, Shape{2, 3}, []int{1}, Shape{2}, []uint16{0, 3}},
	{"A.Min(0,1)", Uint16, Shape{2, 3}, []int{0, 1}, ScalarShape(), uint16(0)},
	{"A.Min(1,0)", Uint16, Shape{2, 3}, []int{1, 0}, ScalarShape(), uint16(0)},
	{"3T.Min(1,2)", Uint16, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []uint16{0, 12}},
	{"common case: T.Min() for uint32", Uint32, Shape{2, 3}, []int{}, ScalarShape(), uint32(0)},
	{"A.Min(0)", Uint32, Shape{2, 3}, []int{0}, Shape{3}, []uint32{0, 1, 2}},
	{"A.Min(1)", Uint32, Shape{2, 3}, []int{1}, Shape{2}, []uint32{0, 3}},
	{"A.Min(0,1)", Uint32, Shape{2, 3}, []int{0, 1}, ScalarShape(), uint32(0)},
	{"A.Min(1,0)", Uint32, Shape{2, 3}, []int{1, 0}, ScalarShape(), uint32(0)},
	{"3T.Min(1,2)", Uint32, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []uint32{0, 12}},
	{"common case: T.Min() for uint64", Uint64, Shape{2, 3}, []int{}, ScalarShape(), uint64(0)},
	{"A.Min(0)", Uint64, Shape{2, 3}, []int{0}, Shape{3}, []uint64{0, 1, 2}},
	{"A.Min(1)", Uint64, Shape{2, 3}, []int{1}, Shape{2}, []uint64{0, 3}},
	{"A.Min(0,1)", Uint64, Shape{2, 3}, []int{0, 1}, ScalarShape(), uint64(0)},
	{"A.Min(1,0)", Uint64, Shape{2, 3}, []int{1, 0}, ScalarShape(), uint64(0)},
	{"3T.Min(1,2)", Uint64, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []uint64{0, 12}},
	{"common case: T.Min() for float32", Float32, Shape{2, 3}, []int{}, ScalarShape(), float32(0)},
	{"A.Min(0)", Float32, Shape{2, 3}, []int{0}, Shape{3}, []float32{0, 1, 2}},
	{"A.Min(1)", Float32, Shape{2, 3}, []int{1}, Shape{2}, []float32{0, 3}},
	{"A.Min(0,1)", Float32, Shape{2, 3}, []int{0, 1}, ScalarShape(), float32(0)},
	{"A.Min(1,0)", Float32, Shape{2, 3}, []int{1, 0}, ScalarShape(), float32(0)},
	{"3T.Min(1,2)", Float32, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []float32{0, 12}},
	{"common case: T.Min() for float64", Float64, Shape{2, 3}, []int{}, ScalarShape(), float64(0)},
	{"A.Min(0)", Float64, Shape{2, 3}, []int{0}, Shape{3}, []float64{0, 1, 2}},
	{"A.Min(1)", Float64, Shape{2, 3}, []int{1}, Shape{2}, []float64{0, 3}},
	{"A.Min(0,1)", Float64, Shape{2, 3}, []int{0, 1}, ScalarShape(), float64(0)},
	{"A.Min(1,0)", Float64, Shape{2, 3}, []int{1, 0}, ScalarShape(), float64(0)},
	{"3T.Min(1,2)", Float64, Shape{2, 3, 4}, []int{1, 2}, Shape{2}, []float64{0, 12}},
}

func TestDense_Min(t *testing.T) {
	assert := assert.New(t)
	var T, T2 *Dense
	var err error

	for _, mts := range minTests {
		T = New(WithShape(mts.shape...), WithBacking(Range(mts.of, 0, mts.shape.TotalSize())))
		if T2, err = T.Min(mts.along...); err != nil {
			t.Error(err)
			continue
		}
		assert.True(mts.correctShape.Eq(T2.Shape()))
		assert.Equal(mts.correct, T2.Data())
	}

	/* IDIOT TESTING TIME */
	_, err = T.Min(1000)
	assert.NotNil(err)
}
