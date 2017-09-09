package execution

import (
	"reflect"

	"github.com/chewxy/gorgonia/tensor/internal/storage"
	"github.com/pkg/errors"
)

/*
GENERATED FILE. DO NOT EDIT
*/

func (e E) Add(t reflect.Type, a *storage.Header, b *storage.Header) (err error) {
	as := isScalar(a)
	bs := isScalar(b)

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()
		switch {
		case as && bs:
			VecAddI(at, bt)
		case as && !bs:
			AddSVI(at[0], bt)
		case !as && bs:
			AddVSI(at, bt[0])
		default:
			VecAddI(at, bt)
		}
		return
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		switch {
		case as && bs:
			VecAddI8(at, bt)
		case as && !bs:
			AddSVI8(at[0], bt)
		case !as && bs:
			AddVSI8(at, bt[0])
		default:
			VecAddI8(at, bt)
		}
		return
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		switch {
		case as && bs:
			VecAddI16(at, bt)
		case as && !bs:
			AddSVI16(at[0], bt)
		case !as && bs:
			AddVSI16(at, bt[0])
		default:
			VecAddI16(at, bt)
		}
		return
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		switch {
		case as && bs:
			VecAddI32(at, bt)
		case as && !bs:
			AddSVI32(at[0], bt)
		case !as && bs:
			AddVSI32(at, bt[0])
		default:
			VecAddI32(at, bt)
		}
		return
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		switch {
		case as && bs:
			VecAddI64(at, bt)
		case as && !bs:
			AddSVI64(at[0], bt)
		case !as && bs:
			AddVSI64(at, bt[0])
		default:
			VecAddI64(at, bt)
		}
		return
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		switch {
		case as && bs:
			VecAddU(at, bt)
		case as && !bs:
			AddSVU(at[0], bt)
		case !as && bs:
			AddVSU(at, bt[0])
		default:
			VecAddU(at, bt)
		}
		return
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		switch {
		case as && bs:
			VecAddU8(at, bt)
		case as && !bs:
			AddSVU8(at[0], bt)
		case !as && bs:
			AddVSU8(at, bt[0])
		default:
			VecAddU8(at, bt)
		}
		return
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		switch {
		case as && bs:
			VecAddU16(at, bt)
		case as && !bs:
			AddSVU16(at[0], bt)
		case !as && bs:
			AddVSU16(at, bt[0])
		default:
			VecAddU16(at, bt)
		}
		return
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		switch {
		case as && bs:
			VecAddU32(at, bt)
		case as && !bs:
			AddSVU32(at[0], bt)
		case !as && bs:
			AddVSU32(at, bt[0])
		default:
			VecAddU32(at, bt)
		}
		return
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		switch {
		case as && bs:
			VecAddU64(at, bt)
		case as && !bs:
			AddSVU64(at[0], bt)
		case !as && bs:
			AddVSU64(at, bt[0])
		default:
			VecAddU64(at, bt)
		}
		return
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		switch {
		case as && bs:
			VecAddF32(at, bt)
		case as && !bs:
			AddSVF32(at[0], bt)
		case !as && bs:
			AddVSF32(at, bt[0])
		default:
			VecAddF32(at, bt)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		switch {
		case as && bs:
			VecAddF64(at, bt)
		case as && !bs:
			AddSVF64(at[0], bt)
		case !as && bs:
			AddVSF64(at, bt[0])
		default:
			VecAddF64(at, bt)
		}
		return
	case Complex64:
		at := a.Complex64s()
		bt := b.Complex64s()
		switch {
		case as && bs:
			VecAddC64(at, bt)
		case as && !bs:
			AddSVC64(at[0], bt)
		case !as && bs:
			AddVSC64(at, bt[0])
		default:
			VecAddC64(at, bt)
		}
		return
	case Complex128:
		at := a.Complex128s()
		bt := b.Complex128s()
		switch {
		case as && bs:
			VecAddC128(at, bt)
		case as && !bs:
			AddSVC128(at[0], bt)
		case !as && bs:
			AddVSC128(at, bt[0])
		default:
			VecAddC128(at, bt)
		}
		return
	case String:
		at := a.Strings()
		bt := b.Strings()
		switch {
		case as && bs:
			VecAddStr(at, bt)
		case as && !bs:
			AddSVStr(at[0], bt)
		case !as && bs:
			AddVSStr(at, bt[0])
		default:
			VecAddStr(at, bt)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Add", t)
	}
}

func (e E) Sub(t reflect.Type, a *storage.Header, b *storage.Header) (err error) {
	as := isScalar(a)
	bs := isScalar(b)

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()
		switch {
		case as && bs:
			VecSubI(at, bt)
		case as && !bs:
			SubSVI(at[0], bt)
		case !as && bs:
			SubVSI(at, bt[0])
		default:
			VecSubI(at, bt)
		}
		return
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		switch {
		case as && bs:
			VecSubI8(at, bt)
		case as && !bs:
			SubSVI8(at[0], bt)
		case !as && bs:
			SubVSI8(at, bt[0])
		default:
			VecSubI8(at, bt)
		}
		return
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		switch {
		case as && bs:
			VecSubI16(at, bt)
		case as && !bs:
			SubSVI16(at[0], bt)
		case !as && bs:
			SubVSI16(at, bt[0])
		default:
			VecSubI16(at, bt)
		}
		return
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		switch {
		case as && bs:
			VecSubI32(at, bt)
		case as && !bs:
			SubSVI32(at[0], bt)
		case !as && bs:
			SubVSI32(at, bt[0])
		default:
			VecSubI32(at, bt)
		}
		return
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		switch {
		case as && bs:
			VecSubI64(at, bt)
		case as && !bs:
			SubSVI64(at[0], bt)
		case !as && bs:
			SubVSI64(at, bt[0])
		default:
			VecSubI64(at, bt)
		}
		return
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		switch {
		case as && bs:
			VecSubU(at, bt)
		case as && !bs:
			SubSVU(at[0], bt)
		case !as && bs:
			SubVSU(at, bt[0])
		default:
			VecSubU(at, bt)
		}
		return
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		switch {
		case as && bs:
			VecSubU8(at, bt)
		case as && !bs:
			SubSVU8(at[0], bt)
		case !as && bs:
			SubVSU8(at, bt[0])
		default:
			VecSubU8(at, bt)
		}
		return
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		switch {
		case as && bs:
			VecSubU16(at, bt)
		case as && !bs:
			SubSVU16(at[0], bt)
		case !as && bs:
			SubVSU16(at, bt[0])
		default:
			VecSubU16(at, bt)
		}
		return
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		switch {
		case as && bs:
			VecSubU32(at, bt)
		case as && !bs:
			SubSVU32(at[0], bt)
		case !as && bs:
			SubVSU32(at, bt[0])
		default:
			VecSubU32(at, bt)
		}
		return
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		switch {
		case as && bs:
			VecSubU64(at, bt)
		case as && !bs:
			SubSVU64(at[0], bt)
		case !as && bs:
			SubVSU64(at, bt[0])
		default:
			VecSubU64(at, bt)
		}
		return
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		switch {
		case as && bs:
			VecSubF32(at, bt)
		case as && !bs:
			SubSVF32(at[0], bt)
		case !as && bs:
			SubVSF32(at, bt[0])
		default:
			VecSubF32(at, bt)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		switch {
		case as && bs:
			VecSubF64(at, bt)
		case as && !bs:
			SubSVF64(at[0], bt)
		case !as && bs:
			SubVSF64(at, bt[0])
		default:
			VecSubF64(at, bt)
		}
		return
	case Complex64:
		at := a.Complex64s()
		bt := b.Complex64s()
		switch {
		case as && bs:
			VecSubC64(at, bt)
		case as && !bs:
			SubSVC64(at[0], bt)
		case !as && bs:
			SubVSC64(at, bt[0])
		default:
			VecSubC64(at, bt)
		}
		return
	case Complex128:
		at := a.Complex128s()
		bt := b.Complex128s()
		switch {
		case as && bs:
			VecSubC128(at, bt)
		case as && !bs:
			SubSVC128(at[0], bt)
		case !as && bs:
			SubVSC128(at, bt[0])
		default:
			VecSubC128(at, bt)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Sub", t)
	}
}

func (e E) Mul(t reflect.Type, a *storage.Header, b *storage.Header) (err error) {
	as := isScalar(a)
	bs := isScalar(b)

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()
		switch {
		case as && bs:
			VecMulI(at, bt)
		case as && !bs:
			MulSVI(at[0], bt)
		case !as && bs:
			MulVSI(at, bt[0])
		default:
			VecMulI(at, bt)
		}
		return
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		switch {
		case as && bs:
			VecMulI8(at, bt)
		case as && !bs:
			MulSVI8(at[0], bt)
		case !as && bs:
			MulVSI8(at, bt[0])
		default:
			VecMulI8(at, bt)
		}
		return
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		switch {
		case as && bs:
			VecMulI16(at, bt)
		case as && !bs:
			MulSVI16(at[0], bt)
		case !as && bs:
			MulVSI16(at, bt[0])
		default:
			VecMulI16(at, bt)
		}
		return
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		switch {
		case as && bs:
			VecMulI32(at, bt)
		case as && !bs:
			MulSVI32(at[0], bt)
		case !as && bs:
			MulVSI32(at, bt[0])
		default:
			VecMulI32(at, bt)
		}
		return
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		switch {
		case as && bs:
			VecMulI64(at, bt)
		case as && !bs:
			MulSVI64(at[0], bt)
		case !as && bs:
			MulVSI64(at, bt[0])
		default:
			VecMulI64(at, bt)
		}
		return
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		switch {
		case as && bs:
			VecMulU(at, bt)
		case as && !bs:
			MulSVU(at[0], bt)
		case !as && bs:
			MulVSU(at, bt[0])
		default:
			VecMulU(at, bt)
		}
		return
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		switch {
		case as && bs:
			VecMulU8(at, bt)
		case as && !bs:
			MulSVU8(at[0], bt)
		case !as && bs:
			MulVSU8(at, bt[0])
		default:
			VecMulU8(at, bt)
		}
		return
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		switch {
		case as && bs:
			VecMulU16(at, bt)
		case as && !bs:
			MulSVU16(at[0], bt)
		case !as && bs:
			MulVSU16(at, bt[0])
		default:
			VecMulU16(at, bt)
		}
		return
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		switch {
		case as && bs:
			VecMulU32(at, bt)
		case as && !bs:
			MulSVU32(at[0], bt)
		case !as && bs:
			MulVSU32(at, bt[0])
		default:
			VecMulU32(at, bt)
		}
		return
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		switch {
		case as && bs:
			VecMulU64(at, bt)
		case as && !bs:
			MulSVU64(at[0], bt)
		case !as && bs:
			MulVSU64(at, bt[0])
		default:
			VecMulU64(at, bt)
		}
		return
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		switch {
		case as && bs:
			VecMulF32(at, bt)
		case as && !bs:
			MulSVF32(at[0], bt)
		case !as && bs:
			MulVSF32(at, bt[0])
		default:
			VecMulF32(at, bt)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		switch {
		case as && bs:
			VecMulF64(at, bt)
		case as && !bs:
			MulSVF64(at[0], bt)
		case !as && bs:
			MulVSF64(at, bt[0])
		default:
			VecMulF64(at, bt)
		}
		return
	case Complex64:
		at := a.Complex64s()
		bt := b.Complex64s()
		switch {
		case as && bs:
			VecMulC64(at, bt)
		case as && !bs:
			MulSVC64(at[0], bt)
		case !as && bs:
			MulVSC64(at, bt[0])
		default:
			VecMulC64(at, bt)
		}
		return
	case Complex128:
		at := a.Complex128s()
		bt := b.Complex128s()
		switch {
		case as && bs:
			VecMulC128(at, bt)
		case as && !bs:
			MulSVC128(at[0], bt)
		case !as && bs:
			MulVSC128(at, bt[0])
		default:
			VecMulC128(at, bt)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Mul", t)
	}
}

func (e E) Div(t reflect.Type, a *storage.Header, b *storage.Header) (err error) {
	as := isScalar(a)
	bs := isScalar(b)

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()
		switch {
		case as && bs:
			VecDivI(at, bt)
		case as && !bs:
			err = DivSVI(at[0], bt)
		case !as && bs:
			err = DivVSI(at, bt[0])
		default:
			err = VecDivI(at, bt)
		}
		return
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		switch {
		case as && bs:
			VecDivI8(at, bt)
		case as && !bs:
			err = DivSVI8(at[0], bt)
		case !as && bs:
			err = DivVSI8(at, bt[0])
		default:
			err = VecDivI8(at, bt)
		}
		return
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		switch {
		case as && bs:
			VecDivI16(at, bt)
		case as && !bs:
			err = DivSVI16(at[0], bt)
		case !as && bs:
			err = DivVSI16(at, bt[0])
		default:
			err = VecDivI16(at, bt)
		}
		return
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		switch {
		case as && bs:
			VecDivI32(at, bt)
		case as && !bs:
			err = DivSVI32(at[0], bt)
		case !as && bs:
			err = DivVSI32(at, bt[0])
		default:
			err = VecDivI32(at, bt)
		}
		return
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		switch {
		case as && bs:
			VecDivI64(at, bt)
		case as && !bs:
			err = DivSVI64(at[0], bt)
		case !as && bs:
			err = DivVSI64(at, bt[0])
		default:
			err = VecDivI64(at, bt)
		}
		return
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		switch {
		case as && bs:
			VecDivU(at, bt)
		case as && !bs:
			err = DivSVU(at[0], bt)
		case !as && bs:
			err = DivVSU(at, bt[0])
		default:
			err = VecDivU(at, bt)
		}
		return
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		switch {
		case as && bs:
			VecDivU8(at, bt)
		case as && !bs:
			err = DivSVU8(at[0], bt)
		case !as && bs:
			err = DivVSU8(at, bt[0])
		default:
			err = VecDivU8(at, bt)
		}
		return
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		switch {
		case as && bs:
			VecDivU16(at, bt)
		case as && !bs:
			err = DivSVU16(at[0], bt)
		case !as && bs:
			err = DivVSU16(at, bt[0])
		default:
			err = VecDivU16(at, bt)
		}
		return
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		switch {
		case as && bs:
			VecDivU32(at, bt)
		case as && !bs:
			err = DivSVU32(at[0], bt)
		case !as && bs:
			err = DivVSU32(at, bt[0])
		default:
			err = VecDivU32(at, bt)
		}
		return
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		switch {
		case as && bs:
			VecDivU64(at, bt)
		case as && !bs:
			err = DivSVU64(at[0], bt)
		case !as && bs:
			err = DivVSU64(at, bt[0])
		default:
			err = VecDivU64(at, bt)
		}
		return
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		switch {
		case as && bs:
			VecDivF32(at, bt)
		case as && !bs:
			DivSVF32(at[0], bt)
		case !as && bs:
			DivVSF32(at, bt[0])
		default:
			VecDivF32(at, bt)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		switch {
		case as && bs:
			VecDivF64(at, bt)
		case as && !bs:
			DivSVF64(at[0], bt)
		case !as && bs:
			DivVSF64(at, bt[0])
		default:
			VecDivF64(at, bt)
		}
		return
	case Complex64:
		at := a.Complex64s()
		bt := b.Complex64s()
		switch {
		case as && bs:
			VecDivC64(at, bt)
		case as && !bs:
			DivSVC64(at[0], bt)
		case !as && bs:
			DivVSC64(at, bt[0])
		default:
			VecDivC64(at, bt)
		}
		return
	case Complex128:
		at := a.Complex128s()
		bt := b.Complex128s()
		switch {
		case as && bs:
			VecDivC128(at, bt)
		case as && !bs:
			DivSVC128(at[0], bt)
		case !as && bs:
			DivVSC128(at, bt[0])
		default:
			VecDivC128(at, bt)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Div", t)
	}
}

func (e E) Pow(t reflect.Type, a *storage.Header, b *storage.Header) (err error) {
	as := isScalar(a)
	bs := isScalar(b)

	switch t {
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		switch {
		case as && bs:
			VecPowF32(at, bt)
		case as && !bs:
			PowSVF32(at[0], bt)
		case !as && bs:
			PowVSF32(at, bt[0])
		default:
			VecPowF32(at, bt)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		switch {
		case as && bs:
			VecPowF64(at, bt)
		case as && !bs:
			PowSVF64(at[0], bt)
		case !as && bs:
			PowVSF64(at, bt[0])
		default:
			VecPowF64(at, bt)
		}
		return
	case Complex64:
		at := a.Complex64s()
		bt := b.Complex64s()
		switch {
		case as && bs:
			VecPowC64(at, bt)
		case as && !bs:
			PowSVC64(at[0], bt)
		case !as && bs:
			PowVSC64(at, bt[0])
		default:
			VecPowC64(at, bt)
		}
		return
	case Complex128:
		at := a.Complex128s()
		bt := b.Complex128s()
		switch {
		case as && bs:
			VecPowC128(at, bt)
		case as && !bs:
			PowSVC128(at[0], bt)
		case !as && bs:
			PowVSC128(at, bt[0])
		default:
			VecPowC128(at, bt)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Pow", t)
	}
}

func (e E) Mod(t reflect.Type, a *storage.Header, b *storage.Header) (err error) {
	as := isScalar(a)
	bs := isScalar(b)

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()
		switch {
		case as && bs:
			VecModI(at, bt)
		case as && !bs:
			ModSVI(at[0], bt)
		case !as && bs:
			ModVSI(at, bt[0])
		default:
			VecModI(at, bt)
		}
		return
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		switch {
		case as && bs:
			VecModI8(at, bt)
		case as && !bs:
			ModSVI8(at[0], bt)
		case !as && bs:
			ModVSI8(at, bt[0])
		default:
			VecModI8(at, bt)
		}
		return
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		switch {
		case as && bs:
			VecModI16(at, bt)
		case as && !bs:
			ModSVI16(at[0], bt)
		case !as && bs:
			ModVSI16(at, bt[0])
		default:
			VecModI16(at, bt)
		}
		return
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		switch {
		case as && bs:
			VecModI32(at, bt)
		case as && !bs:
			ModSVI32(at[0], bt)
		case !as && bs:
			ModVSI32(at, bt[0])
		default:
			VecModI32(at, bt)
		}
		return
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		switch {
		case as && bs:
			VecModI64(at, bt)
		case as && !bs:
			ModSVI64(at[0], bt)
		case !as && bs:
			ModVSI64(at, bt[0])
		default:
			VecModI64(at, bt)
		}
		return
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		switch {
		case as && bs:
			VecModU(at, bt)
		case as && !bs:
			ModSVU(at[0], bt)
		case !as && bs:
			ModVSU(at, bt[0])
		default:
			VecModU(at, bt)
		}
		return
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		switch {
		case as && bs:
			VecModU8(at, bt)
		case as && !bs:
			ModSVU8(at[0], bt)
		case !as && bs:
			ModVSU8(at, bt[0])
		default:
			VecModU8(at, bt)
		}
		return
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		switch {
		case as && bs:
			VecModU16(at, bt)
		case as && !bs:
			ModSVU16(at[0], bt)
		case !as && bs:
			ModVSU16(at, bt[0])
		default:
			VecModU16(at, bt)
		}
		return
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		switch {
		case as && bs:
			VecModU32(at, bt)
		case as && !bs:
			ModSVU32(at[0], bt)
		case !as && bs:
			ModVSU32(at, bt[0])
		default:
			VecModU32(at, bt)
		}
		return
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		switch {
		case as && bs:
			VecModU64(at, bt)
		case as && !bs:
			ModSVU64(at[0], bt)
		case !as && bs:
			ModVSU64(at, bt[0])
		default:
			VecModU64(at, bt)
		}
		return
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		switch {
		case as && bs:
			VecModF32(at, bt)
		case as && !bs:
			ModSVF32(at[0], bt)
		case !as && bs:
			ModVSF32(at, bt[0])
		default:
			VecModF32(at, bt)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		switch {
		case as && bs:
			VecModF64(at, bt)
		case as && !bs:
			ModSVF64(at[0], bt)
		case !as && bs:
			ModVSF64(at, bt[0])
		default:
			VecModF64(at, bt)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Mod", t)
	}
}

func (e E) AddIncr(t reflect.Type, a *storage.Header, b *storage.Header, incr *storage.Header) (err error) {
	as := isScalar(a)
	bs := isScalar(b)
	is := isScalar(incr)
	if ((as && !bs) || (bs && !as)) && is {
		return errors.Errorf("Cannot increment on scalar increment. a: %d, b %d", a.Len(), b.Len())
	}

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()
		it := incr.Ints()

		switch {
		case as && bs:
			VecAddI(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			AddIncrSVI(at[0], bt, it)
		case !as && bs:
			AddIncrVSI(at, bt[0], it)
		default:
			AddIncrI(at, bt, it)
		}
		return
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		it := incr.Int8s()

		switch {
		case as && bs:
			VecAddI8(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			AddIncrSVI8(at[0], bt, it)
		case !as && bs:
			AddIncrVSI8(at, bt[0], it)
		default:
			AddIncrI8(at, bt, it)
		}
		return
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		it := incr.Int16s()

		switch {
		case as && bs:
			VecAddI16(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			AddIncrSVI16(at[0], bt, it)
		case !as && bs:
			AddIncrVSI16(at, bt[0], it)
		default:
			AddIncrI16(at, bt, it)
		}
		return
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		it := incr.Int32s()

		switch {
		case as && bs:
			VecAddI32(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			AddIncrSVI32(at[0], bt, it)
		case !as && bs:
			AddIncrVSI32(at, bt[0], it)
		default:
			AddIncrI32(at, bt, it)
		}
		return
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		it := incr.Int64s()

		switch {
		case as && bs:
			VecAddI64(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			AddIncrSVI64(at[0], bt, it)
		case !as && bs:
			AddIncrVSI64(at, bt[0], it)
		default:
			AddIncrI64(at, bt, it)
		}
		return
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		it := incr.Uints()

		switch {
		case as && bs:
			VecAddU(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			AddIncrSVU(at[0], bt, it)
		case !as && bs:
			AddIncrVSU(at, bt[0], it)
		default:
			AddIncrU(at, bt, it)
		}
		return
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		it := incr.Uint8s()

		switch {
		case as && bs:
			VecAddU8(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			AddIncrSVU8(at[0], bt, it)
		case !as && bs:
			AddIncrVSU8(at, bt[0], it)
		default:
			AddIncrU8(at, bt, it)
		}
		return
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		it := incr.Uint16s()

		switch {
		case as && bs:
			VecAddU16(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			AddIncrSVU16(at[0], bt, it)
		case !as && bs:
			AddIncrVSU16(at, bt[0], it)
		default:
			AddIncrU16(at, bt, it)
		}
		return
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		it := incr.Uint32s()

		switch {
		case as && bs:
			VecAddU32(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			AddIncrSVU32(at[0], bt, it)
		case !as && bs:
			AddIncrVSU32(at, bt[0], it)
		default:
			AddIncrU32(at, bt, it)
		}
		return
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		it := incr.Uint64s()

		switch {
		case as && bs:
			VecAddU64(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			AddIncrSVU64(at[0], bt, it)
		case !as && bs:
			AddIncrVSU64(at, bt[0], it)
		default:
			AddIncrU64(at, bt, it)
		}
		return
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		it := incr.Float32s()

		switch {
		case as && bs:
			VecAddF32(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			AddIncrSVF32(at[0], bt, it)
		case !as && bs:
			AddIncrVSF32(at, bt[0], it)
		default:
			AddIncrF32(at, bt, it)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		it := incr.Float64s()

		switch {
		case as && bs:
			VecAddF64(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			AddIncrSVF64(at[0], bt, it)
		case !as && bs:
			AddIncrVSF64(at, bt[0], it)
		default:
			AddIncrF64(at, bt, it)
		}
		return
	case Complex64:
		at := a.Complex64s()
		bt := b.Complex64s()
		it := incr.Complex64s()

		switch {
		case as && bs:
			VecAddC64(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			AddIncrSVC64(at[0], bt, it)
		case !as && bs:
			AddIncrVSC64(at, bt[0], it)
		default:
			AddIncrC64(at, bt, it)
		}
		return
	case Complex128:
		at := a.Complex128s()
		bt := b.Complex128s()
		it := incr.Complex128s()

		switch {
		case as && bs:
			VecAddC128(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			AddIncrSVC128(at[0], bt, it)
		case !as && bs:
			AddIncrVSC128(at, bt[0], it)
		default:
			AddIncrC128(at, bt, it)
		}
		return
	case String:
		at := a.Strings()
		bt := b.Strings()
		it := incr.Strings()

		switch {
		case as && bs:
			VecAddStr(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			AddIncrSVStr(at[0], bt, it)
		case !as && bs:
			AddIncrVSStr(at, bt[0], it)
		default:
			AddIncrStr(at, bt, it)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Add", t)
	}
}

func (e E) SubIncr(t reflect.Type, a *storage.Header, b *storage.Header, incr *storage.Header) (err error) {
	as := isScalar(a)
	bs := isScalar(b)
	is := isScalar(incr)
	if ((as && !bs) || (bs && !as)) && is {
		return errors.Errorf("Cannot increment on scalar increment. a: %d, b %d", a.Len(), b.Len())
	}

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()
		it := incr.Ints()

		switch {
		case as && bs:
			VecSubI(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			SubIncrSVI(at[0], bt, it)
		case !as && bs:
			SubIncrVSI(at, bt[0], it)
		default:
			SubIncrI(at, bt, it)
		}
		return
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		it := incr.Int8s()

		switch {
		case as && bs:
			VecSubI8(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			SubIncrSVI8(at[0], bt, it)
		case !as && bs:
			SubIncrVSI8(at, bt[0], it)
		default:
			SubIncrI8(at, bt, it)
		}
		return
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		it := incr.Int16s()

		switch {
		case as && bs:
			VecSubI16(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			SubIncrSVI16(at[0], bt, it)
		case !as && bs:
			SubIncrVSI16(at, bt[0], it)
		default:
			SubIncrI16(at, bt, it)
		}
		return
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		it := incr.Int32s()

		switch {
		case as && bs:
			VecSubI32(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			SubIncrSVI32(at[0], bt, it)
		case !as && bs:
			SubIncrVSI32(at, bt[0], it)
		default:
			SubIncrI32(at, bt, it)
		}
		return
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		it := incr.Int64s()

		switch {
		case as && bs:
			VecSubI64(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			SubIncrSVI64(at[0], bt, it)
		case !as && bs:
			SubIncrVSI64(at, bt[0], it)
		default:
			SubIncrI64(at, bt, it)
		}
		return
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		it := incr.Uints()

		switch {
		case as && bs:
			VecSubU(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			SubIncrSVU(at[0], bt, it)
		case !as && bs:
			SubIncrVSU(at, bt[0], it)
		default:
			SubIncrU(at, bt, it)
		}
		return
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		it := incr.Uint8s()

		switch {
		case as && bs:
			VecSubU8(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			SubIncrSVU8(at[0], bt, it)
		case !as && bs:
			SubIncrVSU8(at, bt[0], it)
		default:
			SubIncrU8(at, bt, it)
		}
		return
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		it := incr.Uint16s()

		switch {
		case as && bs:
			VecSubU16(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			SubIncrSVU16(at[0], bt, it)
		case !as && bs:
			SubIncrVSU16(at, bt[0], it)
		default:
			SubIncrU16(at, bt, it)
		}
		return
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		it := incr.Uint32s()

		switch {
		case as && bs:
			VecSubU32(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			SubIncrSVU32(at[0], bt, it)
		case !as && bs:
			SubIncrVSU32(at, bt[0], it)
		default:
			SubIncrU32(at, bt, it)
		}
		return
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		it := incr.Uint64s()

		switch {
		case as && bs:
			VecSubU64(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			SubIncrSVU64(at[0], bt, it)
		case !as && bs:
			SubIncrVSU64(at, bt[0], it)
		default:
			SubIncrU64(at, bt, it)
		}
		return
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		it := incr.Float32s()

		switch {
		case as && bs:
			VecSubF32(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			SubIncrSVF32(at[0], bt, it)
		case !as && bs:
			SubIncrVSF32(at, bt[0], it)
		default:
			SubIncrF32(at, bt, it)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		it := incr.Float64s()

		switch {
		case as && bs:
			VecSubF64(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			SubIncrSVF64(at[0], bt, it)
		case !as && bs:
			SubIncrVSF64(at, bt[0], it)
		default:
			SubIncrF64(at, bt, it)
		}
		return
	case Complex64:
		at := a.Complex64s()
		bt := b.Complex64s()
		it := incr.Complex64s()

		switch {
		case as && bs:
			VecSubC64(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			SubIncrSVC64(at[0], bt, it)
		case !as && bs:
			SubIncrVSC64(at, bt[0], it)
		default:
			SubIncrC64(at, bt, it)
		}
		return
	case Complex128:
		at := a.Complex128s()
		bt := b.Complex128s()
		it := incr.Complex128s()

		switch {
		case as && bs:
			VecSubC128(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			SubIncrSVC128(at[0], bt, it)
		case !as && bs:
			SubIncrVSC128(at, bt[0], it)
		default:
			SubIncrC128(at, bt, it)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Sub", t)
	}
}

func (e E) MulIncr(t reflect.Type, a *storage.Header, b *storage.Header, incr *storage.Header) (err error) {
	as := isScalar(a)
	bs := isScalar(b)
	is := isScalar(incr)
	if ((as && !bs) || (bs && !as)) && is {
		return errors.Errorf("Cannot increment on scalar increment. a: %d, b %d", a.Len(), b.Len())
	}

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()
		it := incr.Ints()

		switch {
		case as && bs:
			VecMulI(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			MulIncrSVI(at[0], bt, it)
		case !as && bs:
			MulIncrVSI(at, bt[0], it)
		default:
			MulIncrI(at, bt, it)
		}
		return
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		it := incr.Int8s()

		switch {
		case as && bs:
			VecMulI8(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			MulIncrSVI8(at[0], bt, it)
		case !as && bs:
			MulIncrVSI8(at, bt[0], it)
		default:
			MulIncrI8(at, bt, it)
		}
		return
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		it := incr.Int16s()

		switch {
		case as && bs:
			VecMulI16(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			MulIncrSVI16(at[0], bt, it)
		case !as && bs:
			MulIncrVSI16(at, bt[0], it)
		default:
			MulIncrI16(at, bt, it)
		}
		return
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		it := incr.Int32s()

		switch {
		case as && bs:
			VecMulI32(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			MulIncrSVI32(at[0], bt, it)
		case !as && bs:
			MulIncrVSI32(at, bt[0], it)
		default:
			MulIncrI32(at, bt, it)
		}
		return
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		it := incr.Int64s()

		switch {
		case as && bs:
			VecMulI64(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			MulIncrSVI64(at[0], bt, it)
		case !as && bs:
			MulIncrVSI64(at, bt[0], it)
		default:
			MulIncrI64(at, bt, it)
		}
		return
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		it := incr.Uints()

		switch {
		case as && bs:
			VecMulU(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			MulIncrSVU(at[0], bt, it)
		case !as && bs:
			MulIncrVSU(at, bt[0], it)
		default:
			MulIncrU(at, bt, it)
		}
		return
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		it := incr.Uint8s()

		switch {
		case as && bs:
			VecMulU8(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			MulIncrSVU8(at[0], bt, it)
		case !as && bs:
			MulIncrVSU8(at, bt[0], it)
		default:
			MulIncrU8(at, bt, it)
		}
		return
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		it := incr.Uint16s()

		switch {
		case as && bs:
			VecMulU16(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			MulIncrSVU16(at[0], bt, it)
		case !as && bs:
			MulIncrVSU16(at, bt[0], it)
		default:
			MulIncrU16(at, bt, it)
		}
		return
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		it := incr.Uint32s()

		switch {
		case as && bs:
			VecMulU32(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			MulIncrSVU32(at[0], bt, it)
		case !as && bs:
			MulIncrVSU32(at, bt[0], it)
		default:
			MulIncrU32(at, bt, it)
		}
		return
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		it := incr.Uint64s()

		switch {
		case as && bs:
			VecMulU64(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			MulIncrSVU64(at[0], bt, it)
		case !as && bs:
			MulIncrVSU64(at, bt[0], it)
		default:
			MulIncrU64(at, bt, it)
		}
		return
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		it := incr.Float32s()

		switch {
		case as && bs:
			VecMulF32(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			MulIncrSVF32(at[0], bt, it)
		case !as && bs:
			MulIncrVSF32(at, bt[0], it)
		default:
			MulIncrF32(at, bt, it)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		it := incr.Float64s()

		switch {
		case as && bs:
			VecMulF64(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			MulIncrSVF64(at[0], bt, it)
		case !as && bs:
			MulIncrVSF64(at, bt[0], it)
		default:
			MulIncrF64(at, bt, it)
		}
		return
	case Complex64:
		at := a.Complex64s()
		bt := b.Complex64s()
		it := incr.Complex64s()

		switch {
		case as && bs:
			VecMulC64(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			MulIncrSVC64(at[0], bt, it)
		case !as && bs:
			MulIncrVSC64(at, bt[0], it)
		default:
			MulIncrC64(at, bt, it)
		}
		return
	case Complex128:
		at := a.Complex128s()
		bt := b.Complex128s()
		it := incr.Complex128s()

		switch {
		case as && bs:
			VecMulC128(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			MulIncrSVC128(at[0], bt, it)
		case !as && bs:
			MulIncrVSC128(at, bt[0], it)
		default:
			MulIncrC128(at, bt, it)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Mul", t)
	}
}

func (e E) DivIncr(t reflect.Type, a *storage.Header, b *storage.Header, incr *storage.Header) (err error) {
	as := isScalar(a)
	bs := isScalar(b)
	is := isScalar(incr)
	if ((as && !bs) || (bs && !as)) && is {
		return errors.Errorf("Cannot increment on scalar increment. a: %d, b %d", a.Len(), b.Len())
	}

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()
		it := incr.Ints()

		switch {
		case as && bs:
			VecDivI(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			DivIncrSVI(at[0], bt, it)
		case !as && bs:
			DivIncrVSI(at, bt[0], it)
		default:
			DivIncrI(at, bt, it)
		}
		return
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		it := incr.Int8s()

		switch {
		case as && bs:
			VecDivI8(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			DivIncrSVI8(at[0], bt, it)
		case !as && bs:
			DivIncrVSI8(at, bt[0], it)
		default:
			DivIncrI8(at, bt, it)
		}
		return
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		it := incr.Int16s()

		switch {
		case as && bs:
			VecDivI16(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			DivIncrSVI16(at[0], bt, it)
		case !as && bs:
			DivIncrVSI16(at, bt[0], it)
		default:
			DivIncrI16(at, bt, it)
		}
		return
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		it := incr.Int32s()

		switch {
		case as && bs:
			VecDivI32(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			DivIncrSVI32(at[0], bt, it)
		case !as && bs:
			DivIncrVSI32(at, bt[0], it)
		default:
			DivIncrI32(at, bt, it)
		}
		return
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		it := incr.Int64s()

		switch {
		case as && bs:
			VecDivI64(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			DivIncrSVI64(at[0], bt, it)
		case !as && bs:
			DivIncrVSI64(at, bt[0], it)
		default:
			DivIncrI64(at, bt, it)
		}
		return
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		it := incr.Uints()

		switch {
		case as && bs:
			VecDivU(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			DivIncrSVU(at[0], bt, it)
		case !as && bs:
			DivIncrVSU(at, bt[0], it)
		default:
			DivIncrU(at, bt, it)
		}
		return
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		it := incr.Uint8s()

		switch {
		case as && bs:
			VecDivU8(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			DivIncrSVU8(at[0], bt, it)
		case !as && bs:
			DivIncrVSU8(at, bt[0], it)
		default:
			DivIncrU8(at, bt, it)
		}
		return
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		it := incr.Uint16s()

		switch {
		case as && bs:
			VecDivU16(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			DivIncrSVU16(at[0], bt, it)
		case !as && bs:
			DivIncrVSU16(at, bt[0], it)
		default:
			DivIncrU16(at, bt, it)
		}
		return
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		it := incr.Uint32s()

		switch {
		case as && bs:
			VecDivU32(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			DivIncrSVU32(at[0], bt, it)
		case !as && bs:
			DivIncrVSU32(at, bt[0], it)
		default:
			DivIncrU32(at, bt, it)
		}
		return
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		it := incr.Uint64s()

		switch {
		case as && bs:
			VecDivU64(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			DivIncrSVU64(at[0], bt, it)
		case !as && bs:
			DivIncrVSU64(at, bt[0], it)
		default:
			DivIncrU64(at, bt, it)
		}
		return
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		it := incr.Float32s()

		switch {
		case as && bs:
			VecDivF32(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			DivIncrSVF32(at[0], bt, it)
		case !as && bs:
			DivIncrVSF32(at, bt[0], it)
		default:
			DivIncrF32(at, bt, it)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		it := incr.Float64s()

		switch {
		case as && bs:
			VecDivF64(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			DivIncrSVF64(at[0], bt, it)
		case !as && bs:
			DivIncrVSF64(at, bt[0], it)
		default:
			DivIncrF64(at, bt, it)
		}
		return
	case Complex64:
		at := a.Complex64s()
		bt := b.Complex64s()
		it := incr.Complex64s()

		switch {
		case as && bs:
			VecDivC64(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			DivIncrSVC64(at[0], bt, it)
		case !as && bs:
			DivIncrVSC64(at, bt[0], it)
		default:
			DivIncrC64(at, bt, it)
		}
		return
	case Complex128:
		at := a.Complex128s()
		bt := b.Complex128s()
		it := incr.Complex128s()

		switch {
		case as && bs:
			VecDivC128(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			DivIncrSVC128(at[0], bt, it)
		case !as && bs:
			DivIncrVSC128(at, bt[0], it)
		default:
			DivIncrC128(at, bt, it)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Div", t)
	}
}

func (e E) PowIncr(t reflect.Type, a *storage.Header, b *storage.Header, incr *storage.Header) (err error) {
	as := isScalar(a)
	bs := isScalar(b)
	is := isScalar(incr)
	if ((as && !bs) || (bs && !as)) && is {
		return errors.Errorf("Cannot increment on scalar increment. a: %d, b %d", a.Len(), b.Len())
	}

	switch t {
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		it := incr.Float32s()

		switch {
		case as && bs:
			VecPowF32(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			PowIncrSVF32(at[0], bt, it)
		case !as && bs:
			PowIncrVSF32(at, bt[0], it)
		default:
			PowIncrF32(at, bt, it)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		it := incr.Float64s()

		switch {
		case as && bs:
			VecPowF64(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			PowIncrSVF64(at[0], bt, it)
		case !as && bs:
			PowIncrVSF64(at, bt[0], it)
		default:
			PowIncrF64(at, bt, it)
		}
		return
	case Complex64:
		at := a.Complex64s()
		bt := b.Complex64s()
		it := incr.Complex64s()

		switch {
		case as && bs:
			VecPowC64(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			PowIncrSVC64(at[0], bt, it)
		case !as && bs:
			PowIncrVSC64(at, bt[0], it)
		default:
			PowIncrC64(at, bt, it)
		}
		return
	case Complex128:
		at := a.Complex128s()
		bt := b.Complex128s()
		it := incr.Complex128s()

		switch {
		case as && bs:
			VecPowC128(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			PowIncrSVC128(at[0], bt, it)
		case !as && bs:
			PowIncrVSC128(at, bt[0], it)
		default:
			PowIncrC128(at, bt, it)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Pow", t)
	}
}

func (e E) ModIncr(t reflect.Type, a *storage.Header, b *storage.Header, incr *storage.Header) (err error) {
	as := isScalar(a)
	bs := isScalar(b)
	is := isScalar(incr)
	if ((as && !bs) || (bs && !as)) && is {
		return errors.Errorf("Cannot increment on scalar increment. a: %d, b %d", a.Len(), b.Len())
	}

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()
		it := incr.Ints()

		switch {
		case as && bs:
			VecModI(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			ModIncrSVI(at[0], bt, it)
		case !as && bs:
			ModIncrVSI(at, bt[0], it)
		default:
			ModIncrI(at, bt, it)
		}
		return
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		it := incr.Int8s()

		switch {
		case as && bs:
			VecModI8(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			ModIncrSVI8(at[0], bt, it)
		case !as && bs:
			ModIncrVSI8(at, bt[0], it)
		default:
			ModIncrI8(at, bt, it)
		}
		return
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		it := incr.Int16s()

		switch {
		case as && bs:
			VecModI16(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			ModIncrSVI16(at[0], bt, it)
		case !as && bs:
			ModIncrVSI16(at, bt[0], it)
		default:
			ModIncrI16(at, bt, it)
		}
		return
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		it := incr.Int32s()

		switch {
		case as && bs:
			VecModI32(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			ModIncrSVI32(at[0], bt, it)
		case !as && bs:
			ModIncrVSI32(at, bt[0], it)
		default:
			ModIncrI32(at, bt, it)
		}
		return
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		it := incr.Int64s()

		switch {
		case as && bs:
			VecModI64(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			ModIncrSVI64(at[0], bt, it)
		case !as && bs:
			ModIncrVSI64(at, bt[0], it)
		default:
			ModIncrI64(at, bt, it)
		}
		return
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		it := incr.Uints()

		switch {
		case as && bs:
			VecModU(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			ModIncrSVU(at[0], bt, it)
		case !as && bs:
			ModIncrVSU(at, bt[0], it)
		default:
			ModIncrU(at, bt, it)
		}
		return
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		it := incr.Uint8s()

		switch {
		case as && bs:
			VecModU8(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			ModIncrSVU8(at[0], bt, it)
		case !as && bs:
			ModIncrVSU8(at, bt[0], it)
		default:
			ModIncrU8(at, bt, it)
		}
		return
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		it := incr.Uint16s()

		switch {
		case as && bs:
			VecModU16(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			ModIncrSVU16(at[0], bt, it)
		case !as && bs:
			ModIncrVSU16(at, bt[0], it)
		default:
			ModIncrU16(at, bt, it)
		}
		return
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		it := incr.Uint32s()

		switch {
		case as && bs:
			VecModU32(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			ModIncrSVU32(at[0], bt, it)
		case !as && bs:
			ModIncrVSU32(at, bt[0], it)
		default:
			ModIncrU32(at, bt, it)
		}
		return
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		it := incr.Uint64s()

		switch {
		case as && bs:
			VecModU64(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			ModIncrSVU64(at[0], bt, it)
		case !as && bs:
			ModIncrVSU64(at, bt[0], it)
		default:
			ModIncrU64(at, bt, it)
		}
		return
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		it := incr.Float32s()

		switch {
		case as && bs:
			VecModF32(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			ModIncrSVF32(at[0], bt, it)
		case !as && bs:
			ModIncrVSF32(at, bt[0], it)
		default:
			ModIncrF32(at, bt, it)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		it := incr.Float64s()

		switch {
		case as && bs:
			VecModF64(at, bt)
			if !is {
				return e.Add(t, incr, a)
			}
			it[0] += at[0]
		case as && !bs:
			ModIncrSVF64(at[0], bt, it)
		case !as && bs:
			ModIncrVSF64(at, bt[0], it)
		default:
			ModIncrF64(at, bt, it)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Mod", t)
	}
}

func (e E) AddIter(t reflect.Type, a *storage.Header, b *storage.Header, ait Iterator, bit Iterator) (err error) {
	as := isScalar(a)
	bs := isScalar(b)

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()
		switch {
		case as && bs:
			VecAddI(at, bt)
		case as && !bs:
			AddIterSVI(at[0], bt, bit)
		case !as && bs:
			AddIterVSI(at, bt[0], ait)
		default:
			AddIterI(at, bt, ait, bit)
		}
		return
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		switch {
		case as && bs:
			VecAddI8(at, bt)
		case as && !bs:
			AddIterSVI8(at[0], bt, bit)
		case !as && bs:
			AddIterVSI8(at, bt[0], ait)
		default:
			AddIterI8(at, bt, ait, bit)
		}
		return
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		switch {
		case as && bs:
			VecAddI16(at, bt)
		case as && !bs:
			AddIterSVI16(at[0], bt, bit)
		case !as && bs:
			AddIterVSI16(at, bt[0], ait)
		default:
			AddIterI16(at, bt, ait, bit)
		}
		return
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		switch {
		case as && bs:
			VecAddI32(at, bt)
		case as && !bs:
			AddIterSVI32(at[0], bt, bit)
		case !as && bs:
			AddIterVSI32(at, bt[0], ait)
		default:
			AddIterI32(at, bt, ait, bit)
		}
		return
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		switch {
		case as && bs:
			VecAddI64(at, bt)
		case as && !bs:
			AddIterSVI64(at[0], bt, bit)
		case !as && bs:
			AddIterVSI64(at, bt[0], ait)
		default:
			AddIterI64(at, bt, ait, bit)
		}
		return
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		switch {
		case as && bs:
			VecAddU(at, bt)
		case as && !bs:
			AddIterSVU(at[0], bt, bit)
		case !as && bs:
			AddIterVSU(at, bt[0], ait)
		default:
			AddIterU(at, bt, ait, bit)
		}
		return
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		switch {
		case as && bs:
			VecAddU8(at, bt)
		case as && !bs:
			AddIterSVU8(at[0], bt, bit)
		case !as && bs:
			AddIterVSU8(at, bt[0], ait)
		default:
			AddIterU8(at, bt, ait, bit)
		}
		return
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		switch {
		case as && bs:
			VecAddU16(at, bt)
		case as && !bs:
			AddIterSVU16(at[0], bt, bit)
		case !as && bs:
			AddIterVSU16(at, bt[0], ait)
		default:
			AddIterU16(at, bt, ait, bit)
		}
		return
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		switch {
		case as && bs:
			VecAddU32(at, bt)
		case as && !bs:
			AddIterSVU32(at[0], bt, bit)
		case !as && bs:
			AddIterVSU32(at, bt[0], ait)
		default:
			AddIterU32(at, bt, ait, bit)
		}
		return
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		switch {
		case as && bs:
			VecAddU64(at, bt)
		case as && !bs:
			AddIterSVU64(at[0], bt, bit)
		case !as && bs:
			AddIterVSU64(at, bt[0], ait)
		default:
			AddIterU64(at, bt, ait, bit)
		}
		return
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		switch {
		case as && bs:
			VecAddF32(at, bt)
		case as && !bs:
			AddIterSVF32(at[0], bt, bit)
		case !as && bs:
			AddIterVSF32(at, bt[0], ait)
		default:
			AddIterF32(at, bt, ait, bit)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		switch {
		case as && bs:
			VecAddF64(at, bt)
		case as && !bs:
			AddIterSVF64(at[0], bt, bit)
		case !as && bs:
			AddIterVSF64(at, bt[0], ait)
		default:
			AddIterF64(at, bt, ait, bit)
		}
		return
	case Complex64:
		at := a.Complex64s()
		bt := b.Complex64s()
		switch {
		case as && bs:
			VecAddC64(at, bt)
		case as && !bs:
			AddIterSVC64(at[0], bt, bit)
		case !as && bs:
			AddIterVSC64(at, bt[0], ait)
		default:
			AddIterC64(at, bt, ait, bit)
		}
		return
	case Complex128:
		at := a.Complex128s()
		bt := b.Complex128s()
		switch {
		case as && bs:
			VecAddC128(at, bt)
		case as && !bs:
			AddIterSVC128(at[0], bt, bit)
		case !as && bs:
			AddIterVSC128(at, bt[0], ait)
		default:
			AddIterC128(at, bt, ait, bit)
		}
		return
	case String:
		at := a.Strings()
		bt := b.Strings()
		switch {
		case as && bs:
			VecAddStr(at, bt)
		case as && !bs:
			AddIterSVStr(at[0], bt, bit)
		case !as && bs:
			AddIterVSStr(at, bt[0], ait)
		default:
			AddIterStr(at, bt, ait, bit)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Add", t)
	}
}

func (e E) SubIter(t reflect.Type, a *storage.Header, b *storage.Header, ait Iterator, bit Iterator) (err error) {
	as := isScalar(a)
	bs := isScalar(b)

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()
		switch {
		case as && bs:
			VecSubI(at, bt)
		case as && !bs:
			SubIterSVI(at[0], bt, bit)
		case !as && bs:
			SubIterVSI(at, bt[0], ait)
		default:
			SubIterI(at, bt, ait, bit)
		}
		return
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		switch {
		case as && bs:
			VecSubI8(at, bt)
		case as && !bs:
			SubIterSVI8(at[0], bt, bit)
		case !as && bs:
			SubIterVSI8(at, bt[0], ait)
		default:
			SubIterI8(at, bt, ait, bit)
		}
		return
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		switch {
		case as && bs:
			VecSubI16(at, bt)
		case as && !bs:
			SubIterSVI16(at[0], bt, bit)
		case !as && bs:
			SubIterVSI16(at, bt[0], ait)
		default:
			SubIterI16(at, bt, ait, bit)
		}
		return
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		switch {
		case as && bs:
			VecSubI32(at, bt)
		case as && !bs:
			SubIterSVI32(at[0], bt, bit)
		case !as && bs:
			SubIterVSI32(at, bt[0], ait)
		default:
			SubIterI32(at, bt, ait, bit)
		}
		return
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		switch {
		case as && bs:
			VecSubI64(at, bt)
		case as && !bs:
			SubIterSVI64(at[0], bt, bit)
		case !as && bs:
			SubIterVSI64(at, bt[0], ait)
		default:
			SubIterI64(at, bt, ait, bit)
		}
		return
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		switch {
		case as && bs:
			VecSubU(at, bt)
		case as && !bs:
			SubIterSVU(at[0], bt, bit)
		case !as && bs:
			SubIterVSU(at, bt[0], ait)
		default:
			SubIterU(at, bt, ait, bit)
		}
		return
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		switch {
		case as && bs:
			VecSubU8(at, bt)
		case as && !bs:
			SubIterSVU8(at[0], bt, bit)
		case !as && bs:
			SubIterVSU8(at, bt[0], ait)
		default:
			SubIterU8(at, bt, ait, bit)
		}
		return
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		switch {
		case as && bs:
			VecSubU16(at, bt)
		case as && !bs:
			SubIterSVU16(at[0], bt, bit)
		case !as && bs:
			SubIterVSU16(at, bt[0], ait)
		default:
			SubIterU16(at, bt, ait, bit)
		}
		return
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		switch {
		case as && bs:
			VecSubU32(at, bt)
		case as && !bs:
			SubIterSVU32(at[0], bt, bit)
		case !as && bs:
			SubIterVSU32(at, bt[0], ait)
		default:
			SubIterU32(at, bt, ait, bit)
		}
		return
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		switch {
		case as && bs:
			VecSubU64(at, bt)
		case as && !bs:
			SubIterSVU64(at[0], bt, bit)
		case !as && bs:
			SubIterVSU64(at, bt[0], ait)
		default:
			SubIterU64(at, bt, ait, bit)
		}
		return
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		switch {
		case as && bs:
			VecSubF32(at, bt)
		case as && !bs:
			SubIterSVF32(at[0], bt, bit)
		case !as && bs:
			SubIterVSF32(at, bt[0], ait)
		default:
			SubIterF32(at, bt, ait, bit)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		switch {
		case as && bs:
			VecSubF64(at, bt)
		case as && !bs:
			SubIterSVF64(at[0], bt, bit)
		case !as && bs:
			SubIterVSF64(at, bt[0], ait)
		default:
			SubIterF64(at, bt, ait, bit)
		}
		return
	case Complex64:
		at := a.Complex64s()
		bt := b.Complex64s()
		switch {
		case as && bs:
			VecSubC64(at, bt)
		case as && !bs:
			SubIterSVC64(at[0], bt, bit)
		case !as && bs:
			SubIterVSC64(at, bt[0], ait)
		default:
			SubIterC64(at, bt, ait, bit)
		}
		return
	case Complex128:
		at := a.Complex128s()
		bt := b.Complex128s()
		switch {
		case as && bs:
			VecSubC128(at, bt)
		case as && !bs:
			SubIterSVC128(at[0], bt, bit)
		case !as && bs:
			SubIterVSC128(at, bt[0], ait)
		default:
			SubIterC128(at, bt, ait, bit)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Sub", t)
	}
}

func (e E) MulIter(t reflect.Type, a *storage.Header, b *storage.Header, ait Iterator, bit Iterator) (err error) {
	as := isScalar(a)
	bs := isScalar(b)

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()
		switch {
		case as && bs:
			VecMulI(at, bt)
		case as && !bs:
			MulIterSVI(at[0], bt, bit)
		case !as && bs:
			MulIterVSI(at, bt[0], ait)
		default:
			MulIterI(at, bt, ait, bit)
		}
		return
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		switch {
		case as && bs:
			VecMulI8(at, bt)
		case as && !bs:
			MulIterSVI8(at[0], bt, bit)
		case !as && bs:
			MulIterVSI8(at, bt[0], ait)
		default:
			MulIterI8(at, bt, ait, bit)
		}
		return
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		switch {
		case as && bs:
			VecMulI16(at, bt)
		case as && !bs:
			MulIterSVI16(at[0], bt, bit)
		case !as && bs:
			MulIterVSI16(at, bt[0], ait)
		default:
			MulIterI16(at, bt, ait, bit)
		}
		return
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		switch {
		case as && bs:
			VecMulI32(at, bt)
		case as && !bs:
			MulIterSVI32(at[0], bt, bit)
		case !as && bs:
			MulIterVSI32(at, bt[0], ait)
		default:
			MulIterI32(at, bt, ait, bit)
		}
		return
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		switch {
		case as && bs:
			VecMulI64(at, bt)
		case as && !bs:
			MulIterSVI64(at[0], bt, bit)
		case !as && bs:
			MulIterVSI64(at, bt[0], ait)
		default:
			MulIterI64(at, bt, ait, bit)
		}
		return
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		switch {
		case as && bs:
			VecMulU(at, bt)
		case as && !bs:
			MulIterSVU(at[0], bt, bit)
		case !as && bs:
			MulIterVSU(at, bt[0], ait)
		default:
			MulIterU(at, bt, ait, bit)
		}
		return
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		switch {
		case as && bs:
			VecMulU8(at, bt)
		case as && !bs:
			MulIterSVU8(at[0], bt, bit)
		case !as && bs:
			MulIterVSU8(at, bt[0], ait)
		default:
			MulIterU8(at, bt, ait, bit)
		}
		return
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		switch {
		case as && bs:
			VecMulU16(at, bt)
		case as && !bs:
			MulIterSVU16(at[0], bt, bit)
		case !as && bs:
			MulIterVSU16(at, bt[0], ait)
		default:
			MulIterU16(at, bt, ait, bit)
		}
		return
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		switch {
		case as && bs:
			VecMulU32(at, bt)
		case as && !bs:
			MulIterSVU32(at[0], bt, bit)
		case !as && bs:
			MulIterVSU32(at, bt[0], ait)
		default:
			MulIterU32(at, bt, ait, bit)
		}
		return
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		switch {
		case as && bs:
			VecMulU64(at, bt)
		case as && !bs:
			MulIterSVU64(at[0], bt, bit)
		case !as && bs:
			MulIterVSU64(at, bt[0], ait)
		default:
			MulIterU64(at, bt, ait, bit)
		}
		return
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		switch {
		case as && bs:
			VecMulF32(at, bt)
		case as && !bs:
			MulIterSVF32(at[0], bt, bit)
		case !as && bs:
			MulIterVSF32(at, bt[0], ait)
		default:
			MulIterF32(at, bt, ait, bit)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		switch {
		case as && bs:
			VecMulF64(at, bt)
		case as && !bs:
			MulIterSVF64(at[0], bt, bit)
		case !as && bs:
			MulIterVSF64(at, bt[0], ait)
		default:
			MulIterF64(at, bt, ait, bit)
		}
		return
	case Complex64:
		at := a.Complex64s()
		bt := b.Complex64s()
		switch {
		case as && bs:
			VecMulC64(at, bt)
		case as && !bs:
			MulIterSVC64(at[0], bt, bit)
		case !as && bs:
			MulIterVSC64(at, bt[0], ait)
		default:
			MulIterC64(at, bt, ait, bit)
		}
		return
	case Complex128:
		at := a.Complex128s()
		bt := b.Complex128s()
		switch {
		case as && bs:
			VecMulC128(at, bt)
		case as && !bs:
			MulIterSVC128(at[0], bt, bit)
		case !as && bs:
			MulIterVSC128(at, bt[0], ait)
		default:
			MulIterC128(at, bt, ait, bit)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Mul", t)
	}
}

func (e E) DivIter(t reflect.Type, a *storage.Header, b *storage.Header, ait Iterator, bit Iterator) (err error) {
	as := isScalar(a)
	bs := isScalar(b)

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()
		switch {
		case as && bs:
			VecDivI(at, bt)
		case as && !bs:
			DivIterSVI(at[0], bt, bit)
		case !as && bs:
			DivIterVSI(at, bt[0], ait)
		default:
			DivIterI(at, bt, ait, bit)
		}
		return
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		switch {
		case as && bs:
			VecDivI8(at, bt)
		case as && !bs:
			DivIterSVI8(at[0], bt, bit)
		case !as && bs:
			DivIterVSI8(at, bt[0], ait)
		default:
			DivIterI8(at, bt, ait, bit)
		}
		return
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		switch {
		case as && bs:
			VecDivI16(at, bt)
		case as && !bs:
			DivIterSVI16(at[0], bt, bit)
		case !as && bs:
			DivIterVSI16(at, bt[0], ait)
		default:
			DivIterI16(at, bt, ait, bit)
		}
		return
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		switch {
		case as && bs:
			VecDivI32(at, bt)
		case as && !bs:
			DivIterSVI32(at[0], bt, bit)
		case !as && bs:
			DivIterVSI32(at, bt[0], ait)
		default:
			DivIterI32(at, bt, ait, bit)
		}
		return
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		switch {
		case as && bs:
			VecDivI64(at, bt)
		case as && !bs:
			DivIterSVI64(at[0], bt, bit)
		case !as && bs:
			DivIterVSI64(at, bt[0], ait)
		default:
			DivIterI64(at, bt, ait, bit)
		}
		return
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		switch {
		case as && bs:
			VecDivU(at, bt)
		case as && !bs:
			DivIterSVU(at[0], bt, bit)
		case !as && bs:
			DivIterVSU(at, bt[0], ait)
		default:
			DivIterU(at, bt, ait, bit)
		}
		return
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		switch {
		case as && bs:
			VecDivU8(at, bt)
		case as && !bs:
			DivIterSVU8(at[0], bt, bit)
		case !as && bs:
			DivIterVSU8(at, bt[0], ait)
		default:
			DivIterU8(at, bt, ait, bit)
		}
		return
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		switch {
		case as && bs:
			VecDivU16(at, bt)
		case as && !bs:
			DivIterSVU16(at[0], bt, bit)
		case !as && bs:
			DivIterVSU16(at, bt[0], ait)
		default:
			DivIterU16(at, bt, ait, bit)
		}
		return
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		switch {
		case as && bs:
			VecDivU32(at, bt)
		case as && !bs:
			DivIterSVU32(at[0], bt, bit)
		case !as && bs:
			DivIterVSU32(at, bt[0], ait)
		default:
			DivIterU32(at, bt, ait, bit)
		}
		return
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		switch {
		case as && bs:
			VecDivU64(at, bt)
		case as && !bs:
			DivIterSVU64(at[0], bt, bit)
		case !as && bs:
			DivIterVSU64(at, bt[0], ait)
		default:
			DivIterU64(at, bt, ait, bit)
		}
		return
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		switch {
		case as && bs:
			VecDivF32(at, bt)
		case as && !bs:
			DivIterSVF32(at[0], bt, bit)
		case !as && bs:
			DivIterVSF32(at, bt[0], ait)
		default:
			DivIterF32(at, bt, ait, bit)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		switch {
		case as && bs:
			VecDivF64(at, bt)
		case as && !bs:
			DivIterSVF64(at[0], bt, bit)
		case !as && bs:
			DivIterVSF64(at, bt[0], ait)
		default:
			DivIterF64(at, bt, ait, bit)
		}
		return
	case Complex64:
		at := a.Complex64s()
		bt := b.Complex64s()
		switch {
		case as && bs:
			VecDivC64(at, bt)
		case as && !bs:
			DivIterSVC64(at[0], bt, bit)
		case !as && bs:
			DivIterVSC64(at, bt[0], ait)
		default:
			DivIterC64(at, bt, ait, bit)
		}
		return
	case Complex128:
		at := a.Complex128s()
		bt := b.Complex128s()
		switch {
		case as && bs:
			VecDivC128(at, bt)
		case as && !bs:
			DivIterSVC128(at[0], bt, bit)
		case !as && bs:
			DivIterVSC128(at, bt[0], ait)
		default:
			DivIterC128(at, bt, ait, bit)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Div", t)
	}
}

func (e E) PowIter(t reflect.Type, a *storage.Header, b *storage.Header, ait Iterator, bit Iterator) (err error) {
	as := isScalar(a)
	bs := isScalar(b)

	switch t {
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		switch {
		case as && bs:
			VecPowF32(at, bt)
		case as && !bs:
			PowIterSVF32(at[0], bt, bit)
		case !as && bs:
			PowIterVSF32(at, bt[0], ait)
		default:
			PowIterF32(at, bt, ait, bit)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		switch {
		case as && bs:
			VecPowF64(at, bt)
		case as && !bs:
			PowIterSVF64(at[0], bt, bit)
		case !as && bs:
			PowIterVSF64(at, bt[0], ait)
		default:
			PowIterF64(at, bt, ait, bit)
		}
		return
	case Complex64:
		at := a.Complex64s()
		bt := b.Complex64s()
		switch {
		case as && bs:
			VecPowC64(at, bt)
		case as && !bs:
			PowIterSVC64(at[0], bt, bit)
		case !as && bs:
			PowIterVSC64(at, bt[0], ait)
		default:
			PowIterC64(at, bt, ait, bit)
		}
		return
	case Complex128:
		at := a.Complex128s()
		bt := b.Complex128s()
		switch {
		case as && bs:
			VecPowC128(at, bt)
		case as && !bs:
			PowIterSVC128(at[0], bt, bit)
		case !as && bs:
			PowIterVSC128(at, bt[0], ait)
		default:
			PowIterC128(at, bt, ait, bit)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Pow", t)
	}
}

func (e E) ModIter(t reflect.Type, a *storage.Header, b *storage.Header, ait Iterator, bit Iterator) (err error) {
	as := isScalar(a)
	bs := isScalar(b)

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()
		switch {
		case as && bs:
			VecModI(at, bt)
		case as && !bs:
			ModIterSVI(at[0], bt, bit)
		case !as && bs:
			ModIterVSI(at, bt[0], ait)
		default:
			ModIterI(at, bt, ait, bit)
		}
		return
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		switch {
		case as && bs:
			VecModI8(at, bt)
		case as && !bs:
			ModIterSVI8(at[0], bt, bit)
		case !as && bs:
			ModIterVSI8(at, bt[0], ait)
		default:
			ModIterI8(at, bt, ait, bit)
		}
		return
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		switch {
		case as && bs:
			VecModI16(at, bt)
		case as && !bs:
			ModIterSVI16(at[0], bt, bit)
		case !as && bs:
			ModIterVSI16(at, bt[0], ait)
		default:
			ModIterI16(at, bt, ait, bit)
		}
		return
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		switch {
		case as && bs:
			VecModI32(at, bt)
		case as && !bs:
			ModIterSVI32(at[0], bt, bit)
		case !as && bs:
			ModIterVSI32(at, bt[0], ait)
		default:
			ModIterI32(at, bt, ait, bit)
		}
		return
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		switch {
		case as && bs:
			VecModI64(at, bt)
		case as && !bs:
			ModIterSVI64(at[0], bt, bit)
		case !as && bs:
			ModIterVSI64(at, bt[0], ait)
		default:
			ModIterI64(at, bt, ait, bit)
		}
		return
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		switch {
		case as && bs:
			VecModU(at, bt)
		case as && !bs:
			ModIterSVU(at[0], bt, bit)
		case !as && bs:
			ModIterVSU(at, bt[0], ait)
		default:
			ModIterU(at, bt, ait, bit)
		}
		return
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		switch {
		case as && bs:
			VecModU8(at, bt)
		case as && !bs:
			ModIterSVU8(at[0], bt, bit)
		case !as && bs:
			ModIterVSU8(at, bt[0], ait)
		default:
			ModIterU8(at, bt, ait, bit)
		}
		return
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		switch {
		case as && bs:
			VecModU16(at, bt)
		case as && !bs:
			ModIterSVU16(at[0], bt, bit)
		case !as && bs:
			ModIterVSU16(at, bt[0], ait)
		default:
			ModIterU16(at, bt, ait, bit)
		}
		return
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		switch {
		case as && bs:
			VecModU32(at, bt)
		case as && !bs:
			ModIterSVU32(at[0], bt, bit)
		case !as && bs:
			ModIterVSU32(at, bt[0], ait)
		default:
			ModIterU32(at, bt, ait, bit)
		}
		return
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		switch {
		case as && bs:
			VecModU64(at, bt)
		case as && !bs:
			ModIterSVU64(at[0], bt, bit)
		case !as && bs:
			ModIterVSU64(at, bt[0], ait)
		default:
			ModIterU64(at, bt, ait, bit)
		}
		return
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		switch {
		case as && bs:
			VecModF32(at, bt)
		case as && !bs:
			ModIterSVF32(at[0], bt, bit)
		case !as && bs:
			ModIterVSF32(at, bt[0], ait)
		default:
			ModIterF32(at, bt, ait, bit)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		switch {
		case as && bs:
			VecModF64(at, bt)
		case as && !bs:
			ModIterSVF64(at[0], bt, bit)
		case !as && bs:
			ModIterVSF64(at, bt[0], ait)
		default:
			ModIterF64(at, bt, ait, bit)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Mod", t)
	}
}

func (e E) AddIterIncr(t reflect.Type, a *storage.Header, b *storage.Header, incr *storage.Header, ait Iterator, bit Iterator, iit Iterator) (err error) {
	as := isScalar(a)
	bs := isScalar(b)
	is := isScalar(incr)

	if ((as && !bs) || (bs && !as)) && is {
		return errors.Errorf("Cannot increment on a scalar increment. len(a): %d, len(b) %d", a.Len(), b.Len())
	}

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()
		it := incr.Ints()
		switch {
		case as && bs:
			VecAddI(at, bt)
			if !is {
				return e.AddIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return AddIterIncrSVI(at[0], bt, it, bit, iit)
		case !as && bs:
			return AddIterIncrVSI(at, bt[0], it, ait, iit)
		default:
			return AddIterIncrI(at, bt, it, ait, bit, iit)
		}
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		it := incr.Int8s()
		switch {
		case as && bs:
			VecAddI8(at, bt)
			if !is {
				return e.AddIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return AddIterIncrSVI8(at[0], bt, it, bit, iit)
		case !as && bs:
			return AddIterIncrVSI8(at, bt[0], it, ait, iit)
		default:
			return AddIterIncrI8(at, bt, it, ait, bit, iit)
		}
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		it := incr.Int16s()
		switch {
		case as && bs:
			VecAddI16(at, bt)
			if !is {
				return e.AddIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return AddIterIncrSVI16(at[0], bt, it, bit, iit)
		case !as && bs:
			return AddIterIncrVSI16(at, bt[0], it, ait, iit)
		default:
			return AddIterIncrI16(at, bt, it, ait, bit, iit)
		}
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		it := incr.Int32s()
		switch {
		case as && bs:
			VecAddI32(at, bt)
			if !is {
				return e.AddIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return AddIterIncrSVI32(at[0], bt, it, bit, iit)
		case !as && bs:
			return AddIterIncrVSI32(at, bt[0], it, ait, iit)
		default:
			return AddIterIncrI32(at, bt, it, ait, bit, iit)
		}
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		it := incr.Int64s()
		switch {
		case as && bs:
			VecAddI64(at, bt)
			if !is {
				return e.AddIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return AddIterIncrSVI64(at[0], bt, it, bit, iit)
		case !as && bs:
			return AddIterIncrVSI64(at, bt[0], it, ait, iit)
		default:
			return AddIterIncrI64(at, bt, it, ait, bit, iit)
		}
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		it := incr.Uints()
		switch {
		case as && bs:
			VecAddU(at, bt)
			if !is {
				return e.AddIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return AddIterIncrSVU(at[0], bt, it, bit, iit)
		case !as && bs:
			return AddIterIncrVSU(at, bt[0], it, ait, iit)
		default:
			return AddIterIncrU(at, bt, it, ait, bit, iit)
		}
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		it := incr.Uint8s()
		switch {
		case as && bs:
			VecAddU8(at, bt)
			if !is {
				return e.AddIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return AddIterIncrSVU8(at[0], bt, it, bit, iit)
		case !as && bs:
			return AddIterIncrVSU8(at, bt[0], it, ait, iit)
		default:
			return AddIterIncrU8(at, bt, it, ait, bit, iit)
		}
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		it := incr.Uint16s()
		switch {
		case as && bs:
			VecAddU16(at, bt)
			if !is {
				return e.AddIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return AddIterIncrSVU16(at[0], bt, it, bit, iit)
		case !as && bs:
			return AddIterIncrVSU16(at, bt[0], it, ait, iit)
		default:
			return AddIterIncrU16(at, bt, it, ait, bit, iit)
		}
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		it := incr.Uint32s()
		switch {
		case as && bs:
			VecAddU32(at, bt)
			if !is {
				return e.AddIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return AddIterIncrSVU32(at[0], bt, it, bit, iit)
		case !as && bs:
			return AddIterIncrVSU32(at, bt[0], it, ait, iit)
		default:
			return AddIterIncrU32(at, bt, it, ait, bit, iit)
		}
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		it := incr.Uint64s()
		switch {
		case as && bs:
			VecAddU64(at, bt)
			if !is {
				return e.AddIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return AddIterIncrSVU64(at[0], bt, it, bit, iit)
		case !as && bs:
			return AddIterIncrVSU64(at, bt[0], it, ait, iit)
		default:
			return AddIterIncrU64(at, bt, it, ait, bit, iit)
		}
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		it := incr.Float32s()
		switch {
		case as && bs:
			VecAddF32(at, bt)
			if !is {
				return e.AddIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return AddIterIncrSVF32(at[0], bt, it, bit, iit)
		case !as && bs:
			return AddIterIncrVSF32(at, bt[0], it, ait, iit)
		default:
			return AddIterIncrF32(at, bt, it, ait, bit, iit)
		}
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		it := incr.Float64s()
		switch {
		case as && bs:
			VecAddF64(at, bt)
			if !is {
				return e.AddIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return AddIterIncrSVF64(at[0], bt, it, bit, iit)
		case !as && bs:
			return AddIterIncrVSF64(at, bt[0], it, ait, iit)
		default:
			return AddIterIncrF64(at, bt, it, ait, bit, iit)
		}
	case Complex64:
		at := a.Complex64s()
		bt := b.Complex64s()
		it := incr.Complex64s()
		switch {
		case as && bs:
			VecAddC64(at, bt)
			if !is {
				return e.AddIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return AddIterIncrSVC64(at[0], bt, it, bit, iit)
		case !as && bs:
			return AddIterIncrVSC64(at, bt[0], it, ait, iit)
		default:
			return AddIterIncrC64(at, bt, it, ait, bit, iit)
		}
	case Complex128:
		at := a.Complex128s()
		bt := b.Complex128s()
		it := incr.Complex128s()
		switch {
		case as && bs:
			VecAddC128(at, bt)
			if !is {
				return e.AddIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return AddIterIncrSVC128(at[0], bt, it, bit, iit)
		case !as && bs:
			return AddIterIncrVSC128(at, bt[0], it, ait, iit)
		default:
			return AddIterIncrC128(at, bt, it, ait, bit, iit)
		}
	case String:
		at := a.Strings()
		bt := b.Strings()
		it := incr.Strings()
		switch {
		case as && bs:
			VecAddStr(at, bt)
			if !is {
				return e.AddIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return AddIterIncrSVStr(at[0], bt, it, bit, iit)
		case !as && bs:
			return AddIterIncrVSStr(at, bt[0], it, ait, iit)
		default:
			return AddIterIncrStr(at, bt, it, ait, bit, iit)
		}
	default:
		return errors.Errorf("Unsupported type %v for Add", t)
	}
}

func (e E) SubIterIncr(t reflect.Type, a *storage.Header, b *storage.Header, incr *storage.Header, ait Iterator, bit Iterator, iit Iterator) (err error) {
	as := isScalar(a)
	bs := isScalar(b)
	is := isScalar(incr)

	if ((as && !bs) || (bs && !as)) && is {
		return errors.Errorf("Cannot increment on a scalar increment. len(a): %d, len(b) %d", a.Len(), b.Len())
	}

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()
		it := incr.Ints()
		switch {
		case as && bs:
			VecSubI(at, bt)
			if !is {
				return e.SubIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return SubIterIncrSVI(at[0], bt, it, bit, iit)
		case !as && bs:
			return SubIterIncrVSI(at, bt[0], it, ait, iit)
		default:
			return SubIterIncrI(at, bt, it, ait, bit, iit)
		}
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		it := incr.Int8s()
		switch {
		case as && bs:
			VecSubI8(at, bt)
			if !is {
				return e.SubIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return SubIterIncrSVI8(at[0], bt, it, bit, iit)
		case !as && bs:
			return SubIterIncrVSI8(at, bt[0], it, ait, iit)
		default:
			return SubIterIncrI8(at, bt, it, ait, bit, iit)
		}
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		it := incr.Int16s()
		switch {
		case as && bs:
			VecSubI16(at, bt)
			if !is {
				return e.SubIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return SubIterIncrSVI16(at[0], bt, it, bit, iit)
		case !as && bs:
			return SubIterIncrVSI16(at, bt[0], it, ait, iit)
		default:
			return SubIterIncrI16(at, bt, it, ait, bit, iit)
		}
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		it := incr.Int32s()
		switch {
		case as && bs:
			VecSubI32(at, bt)
			if !is {
				return e.SubIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return SubIterIncrSVI32(at[0], bt, it, bit, iit)
		case !as && bs:
			return SubIterIncrVSI32(at, bt[0], it, ait, iit)
		default:
			return SubIterIncrI32(at, bt, it, ait, bit, iit)
		}
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		it := incr.Int64s()
		switch {
		case as && bs:
			VecSubI64(at, bt)
			if !is {
				return e.SubIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return SubIterIncrSVI64(at[0], bt, it, bit, iit)
		case !as && bs:
			return SubIterIncrVSI64(at, bt[0], it, ait, iit)
		default:
			return SubIterIncrI64(at, bt, it, ait, bit, iit)
		}
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		it := incr.Uints()
		switch {
		case as && bs:
			VecSubU(at, bt)
			if !is {
				return e.SubIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return SubIterIncrSVU(at[0], bt, it, bit, iit)
		case !as && bs:
			return SubIterIncrVSU(at, bt[0], it, ait, iit)
		default:
			return SubIterIncrU(at, bt, it, ait, bit, iit)
		}
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		it := incr.Uint8s()
		switch {
		case as && bs:
			VecSubU8(at, bt)
			if !is {
				return e.SubIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return SubIterIncrSVU8(at[0], bt, it, bit, iit)
		case !as && bs:
			return SubIterIncrVSU8(at, bt[0], it, ait, iit)
		default:
			return SubIterIncrU8(at, bt, it, ait, bit, iit)
		}
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		it := incr.Uint16s()
		switch {
		case as && bs:
			VecSubU16(at, bt)
			if !is {
				return e.SubIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return SubIterIncrSVU16(at[0], bt, it, bit, iit)
		case !as && bs:
			return SubIterIncrVSU16(at, bt[0], it, ait, iit)
		default:
			return SubIterIncrU16(at, bt, it, ait, bit, iit)
		}
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		it := incr.Uint32s()
		switch {
		case as && bs:
			VecSubU32(at, bt)
			if !is {
				return e.SubIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return SubIterIncrSVU32(at[0], bt, it, bit, iit)
		case !as && bs:
			return SubIterIncrVSU32(at, bt[0], it, ait, iit)
		default:
			return SubIterIncrU32(at, bt, it, ait, bit, iit)
		}
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		it := incr.Uint64s()
		switch {
		case as && bs:
			VecSubU64(at, bt)
			if !is {
				return e.SubIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return SubIterIncrSVU64(at[0], bt, it, bit, iit)
		case !as && bs:
			return SubIterIncrVSU64(at, bt[0], it, ait, iit)
		default:
			return SubIterIncrU64(at, bt, it, ait, bit, iit)
		}
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		it := incr.Float32s()
		switch {
		case as && bs:
			VecSubF32(at, bt)
			if !is {
				return e.SubIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return SubIterIncrSVF32(at[0], bt, it, bit, iit)
		case !as && bs:
			return SubIterIncrVSF32(at, bt[0], it, ait, iit)
		default:
			return SubIterIncrF32(at, bt, it, ait, bit, iit)
		}
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		it := incr.Float64s()
		switch {
		case as && bs:
			VecSubF64(at, bt)
			if !is {
				return e.SubIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return SubIterIncrSVF64(at[0], bt, it, bit, iit)
		case !as && bs:
			return SubIterIncrVSF64(at, bt[0], it, ait, iit)
		default:
			return SubIterIncrF64(at, bt, it, ait, bit, iit)
		}
	case Complex64:
		at := a.Complex64s()
		bt := b.Complex64s()
		it := incr.Complex64s()
		switch {
		case as && bs:
			VecSubC64(at, bt)
			if !is {
				return e.SubIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return SubIterIncrSVC64(at[0], bt, it, bit, iit)
		case !as && bs:
			return SubIterIncrVSC64(at, bt[0], it, ait, iit)
		default:
			return SubIterIncrC64(at, bt, it, ait, bit, iit)
		}
	case Complex128:
		at := a.Complex128s()
		bt := b.Complex128s()
		it := incr.Complex128s()
		switch {
		case as && bs:
			VecSubC128(at, bt)
			if !is {
				return e.SubIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return SubIterIncrSVC128(at[0], bt, it, bit, iit)
		case !as && bs:
			return SubIterIncrVSC128(at, bt[0], it, ait, iit)
		default:
			return SubIterIncrC128(at, bt, it, ait, bit, iit)
		}
	default:
		return errors.Errorf("Unsupported type %v for Sub", t)
	}
}

func (e E) MulIterIncr(t reflect.Type, a *storage.Header, b *storage.Header, incr *storage.Header, ait Iterator, bit Iterator, iit Iterator) (err error) {
	as := isScalar(a)
	bs := isScalar(b)
	is := isScalar(incr)

	if ((as && !bs) || (bs && !as)) && is {
		return errors.Errorf("Cannot increment on a scalar increment. len(a): %d, len(b) %d", a.Len(), b.Len())
	}

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()
		it := incr.Ints()
		switch {
		case as && bs:
			VecMulI(at, bt)
			if !is {
				return e.MulIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return MulIterIncrSVI(at[0], bt, it, bit, iit)
		case !as && bs:
			return MulIterIncrVSI(at, bt[0], it, ait, iit)
		default:
			return MulIterIncrI(at, bt, it, ait, bit, iit)
		}
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		it := incr.Int8s()
		switch {
		case as && bs:
			VecMulI8(at, bt)
			if !is {
				return e.MulIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return MulIterIncrSVI8(at[0], bt, it, bit, iit)
		case !as && bs:
			return MulIterIncrVSI8(at, bt[0], it, ait, iit)
		default:
			return MulIterIncrI8(at, bt, it, ait, bit, iit)
		}
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		it := incr.Int16s()
		switch {
		case as && bs:
			VecMulI16(at, bt)
			if !is {
				return e.MulIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return MulIterIncrSVI16(at[0], bt, it, bit, iit)
		case !as && bs:
			return MulIterIncrVSI16(at, bt[0], it, ait, iit)
		default:
			return MulIterIncrI16(at, bt, it, ait, bit, iit)
		}
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		it := incr.Int32s()
		switch {
		case as && bs:
			VecMulI32(at, bt)
			if !is {
				return e.MulIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return MulIterIncrSVI32(at[0], bt, it, bit, iit)
		case !as && bs:
			return MulIterIncrVSI32(at, bt[0], it, ait, iit)
		default:
			return MulIterIncrI32(at, bt, it, ait, bit, iit)
		}
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		it := incr.Int64s()
		switch {
		case as && bs:
			VecMulI64(at, bt)
			if !is {
				return e.MulIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return MulIterIncrSVI64(at[0], bt, it, bit, iit)
		case !as && bs:
			return MulIterIncrVSI64(at, bt[0], it, ait, iit)
		default:
			return MulIterIncrI64(at, bt, it, ait, bit, iit)
		}
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		it := incr.Uints()
		switch {
		case as && bs:
			VecMulU(at, bt)
			if !is {
				return e.MulIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return MulIterIncrSVU(at[0], bt, it, bit, iit)
		case !as && bs:
			return MulIterIncrVSU(at, bt[0], it, ait, iit)
		default:
			return MulIterIncrU(at, bt, it, ait, bit, iit)
		}
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		it := incr.Uint8s()
		switch {
		case as && bs:
			VecMulU8(at, bt)
			if !is {
				return e.MulIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return MulIterIncrSVU8(at[0], bt, it, bit, iit)
		case !as && bs:
			return MulIterIncrVSU8(at, bt[0], it, ait, iit)
		default:
			return MulIterIncrU8(at, bt, it, ait, bit, iit)
		}
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		it := incr.Uint16s()
		switch {
		case as && bs:
			VecMulU16(at, bt)
			if !is {
				return e.MulIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return MulIterIncrSVU16(at[0], bt, it, bit, iit)
		case !as && bs:
			return MulIterIncrVSU16(at, bt[0], it, ait, iit)
		default:
			return MulIterIncrU16(at, bt, it, ait, bit, iit)
		}
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		it := incr.Uint32s()
		switch {
		case as && bs:
			VecMulU32(at, bt)
			if !is {
				return e.MulIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return MulIterIncrSVU32(at[0], bt, it, bit, iit)
		case !as && bs:
			return MulIterIncrVSU32(at, bt[0], it, ait, iit)
		default:
			return MulIterIncrU32(at, bt, it, ait, bit, iit)
		}
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		it := incr.Uint64s()
		switch {
		case as && bs:
			VecMulU64(at, bt)
			if !is {
				return e.MulIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return MulIterIncrSVU64(at[0], bt, it, bit, iit)
		case !as && bs:
			return MulIterIncrVSU64(at, bt[0], it, ait, iit)
		default:
			return MulIterIncrU64(at, bt, it, ait, bit, iit)
		}
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		it := incr.Float32s()
		switch {
		case as && bs:
			VecMulF32(at, bt)
			if !is {
				return e.MulIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return MulIterIncrSVF32(at[0], bt, it, bit, iit)
		case !as && bs:
			return MulIterIncrVSF32(at, bt[0], it, ait, iit)
		default:
			return MulIterIncrF32(at, bt, it, ait, bit, iit)
		}
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		it := incr.Float64s()
		switch {
		case as && bs:
			VecMulF64(at, bt)
			if !is {
				return e.MulIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return MulIterIncrSVF64(at[0], bt, it, bit, iit)
		case !as && bs:
			return MulIterIncrVSF64(at, bt[0], it, ait, iit)
		default:
			return MulIterIncrF64(at, bt, it, ait, bit, iit)
		}
	case Complex64:
		at := a.Complex64s()
		bt := b.Complex64s()
		it := incr.Complex64s()
		switch {
		case as && bs:
			VecMulC64(at, bt)
			if !is {
				return e.MulIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return MulIterIncrSVC64(at[0], bt, it, bit, iit)
		case !as && bs:
			return MulIterIncrVSC64(at, bt[0], it, ait, iit)
		default:
			return MulIterIncrC64(at, bt, it, ait, bit, iit)
		}
	case Complex128:
		at := a.Complex128s()
		bt := b.Complex128s()
		it := incr.Complex128s()
		switch {
		case as && bs:
			VecMulC128(at, bt)
			if !is {
				return e.MulIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return MulIterIncrSVC128(at[0], bt, it, bit, iit)
		case !as && bs:
			return MulIterIncrVSC128(at, bt[0], it, ait, iit)
		default:
			return MulIterIncrC128(at, bt, it, ait, bit, iit)
		}
	default:
		return errors.Errorf("Unsupported type %v for Mul", t)
	}
}

func (e E) DivIterIncr(t reflect.Type, a *storage.Header, b *storage.Header, incr *storage.Header, ait Iterator, bit Iterator, iit Iterator) (err error) {
	as := isScalar(a)
	bs := isScalar(b)
	is := isScalar(incr)

	if ((as && !bs) || (bs && !as)) && is {
		return errors.Errorf("Cannot increment on a scalar increment. len(a): %d, len(b) %d", a.Len(), b.Len())
	}

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()
		it := incr.Ints()
		switch {
		case as && bs:
			VecDivI(at, bt)
			if !is {
				return e.DivIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return DivIterIncrSVI(at[0], bt, it, bit, iit)
		case !as && bs:
			return DivIterIncrVSI(at, bt[0], it, ait, iit)
		default:
			return DivIterIncrI(at, bt, it, ait, bit, iit)
		}
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		it := incr.Int8s()
		switch {
		case as && bs:
			VecDivI8(at, bt)
			if !is {
				return e.DivIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return DivIterIncrSVI8(at[0], bt, it, bit, iit)
		case !as && bs:
			return DivIterIncrVSI8(at, bt[0], it, ait, iit)
		default:
			return DivIterIncrI8(at, bt, it, ait, bit, iit)
		}
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		it := incr.Int16s()
		switch {
		case as && bs:
			VecDivI16(at, bt)
			if !is {
				return e.DivIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return DivIterIncrSVI16(at[0], bt, it, bit, iit)
		case !as && bs:
			return DivIterIncrVSI16(at, bt[0], it, ait, iit)
		default:
			return DivIterIncrI16(at, bt, it, ait, bit, iit)
		}
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		it := incr.Int32s()
		switch {
		case as && bs:
			VecDivI32(at, bt)
			if !is {
				return e.DivIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return DivIterIncrSVI32(at[0], bt, it, bit, iit)
		case !as && bs:
			return DivIterIncrVSI32(at, bt[0], it, ait, iit)
		default:
			return DivIterIncrI32(at, bt, it, ait, bit, iit)
		}
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		it := incr.Int64s()
		switch {
		case as && bs:
			VecDivI64(at, bt)
			if !is {
				return e.DivIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return DivIterIncrSVI64(at[0], bt, it, bit, iit)
		case !as && bs:
			return DivIterIncrVSI64(at, bt[0], it, ait, iit)
		default:
			return DivIterIncrI64(at, bt, it, ait, bit, iit)
		}
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		it := incr.Uints()
		switch {
		case as && bs:
			VecDivU(at, bt)
			if !is {
				return e.DivIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return DivIterIncrSVU(at[0], bt, it, bit, iit)
		case !as && bs:
			return DivIterIncrVSU(at, bt[0], it, ait, iit)
		default:
			return DivIterIncrU(at, bt, it, ait, bit, iit)
		}
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		it := incr.Uint8s()
		switch {
		case as && bs:
			VecDivU8(at, bt)
			if !is {
				return e.DivIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return DivIterIncrSVU8(at[0], bt, it, bit, iit)
		case !as && bs:
			return DivIterIncrVSU8(at, bt[0], it, ait, iit)
		default:
			return DivIterIncrU8(at, bt, it, ait, bit, iit)
		}
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		it := incr.Uint16s()
		switch {
		case as && bs:
			VecDivU16(at, bt)
			if !is {
				return e.DivIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return DivIterIncrSVU16(at[0], bt, it, bit, iit)
		case !as && bs:
			return DivIterIncrVSU16(at, bt[0], it, ait, iit)
		default:
			return DivIterIncrU16(at, bt, it, ait, bit, iit)
		}
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		it := incr.Uint32s()
		switch {
		case as && bs:
			VecDivU32(at, bt)
			if !is {
				return e.DivIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return DivIterIncrSVU32(at[0], bt, it, bit, iit)
		case !as && bs:
			return DivIterIncrVSU32(at, bt[0], it, ait, iit)
		default:
			return DivIterIncrU32(at, bt, it, ait, bit, iit)
		}
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		it := incr.Uint64s()
		switch {
		case as && bs:
			VecDivU64(at, bt)
			if !is {
				return e.DivIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return DivIterIncrSVU64(at[0], bt, it, bit, iit)
		case !as && bs:
			return DivIterIncrVSU64(at, bt[0], it, ait, iit)
		default:
			return DivIterIncrU64(at, bt, it, ait, bit, iit)
		}
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		it := incr.Float32s()
		switch {
		case as && bs:
			VecDivF32(at, bt)
			if !is {
				return e.DivIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return DivIterIncrSVF32(at[0], bt, it, bit, iit)
		case !as && bs:
			return DivIterIncrVSF32(at, bt[0], it, ait, iit)
		default:
			return DivIterIncrF32(at, bt, it, ait, bit, iit)
		}
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		it := incr.Float64s()
		switch {
		case as && bs:
			VecDivF64(at, bt)
			if !is {
				return e.DivIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return DivIterIncrSVF64(at[0], bt, it, bit, iit)
		case !as && bs:
			return DivIterIncrVSF64(at, bt[0], it, ait, iit)
		default:
			return DivIterIncrF64(at, bt, it, ait, bit, iit)
		}
	case Complex64:
		at := a.Complex64s()
		bt := b.Complex64s()
		it := incr.Complex64s()
		switch {
		case as && bs:
			VecDivC64(at, bt)
			if !is {
				return e.DivIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return DivIterIncrSVC64(at[0], bt, it, bit, iit)
		case !as && bs:
			return DivIterIncrVSC64(at, bt[0], it, ait, iit)
		default:
			return DivIterIncrC64(at, bt, it, ait, bit, iit)
		}
	case Complex128:
		at := a.Complex128s()
		bt := b.Complex128s()
		it := incr.Complex128s()
		switch {
		case as && bs:
			VecDivC128(at, bt)
			if !is {
				return e.DivIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return DivIterIncrSVC128(at[0], bt, it, bit, iit)
		case !as && bs:
			return DivIterIncrVSC128(at, bt[0], it, ait, iit)
		default:
			return DivIterIncrC128(at, bt, it, ait, bit, iit)
		}
	default:
		return errors.Errorf("Unsupported type %v for Div", t)
	}
}

func (e E) PowIterIncr(t reflect.Type, a *storage.Header, b *storage.Header, incr *storage.Header, ait Iterator, bit Iterator, iit Iterator) (err error) {
	as := isScalar(a)
	bs := isScalar(b)
	is := isScalar(incr)

	if ((as && !bs) || (bs && !as)) && is {
		return errors.Errorf("Cannot increment on a scalar increment. len(a): %d, len(b) %d", a.Len(), b.Len())
	}

	switch t {
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		it := incr.Float32s()
		switch {
		case as && bs:
			VecPowF32(at, bt)
			if !is {
				return e.PowIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return PowIterIncrSVF32(at[0], bt, it, bit, iit)
		case !as && bs:
			return PowIterIncrVSF32(at, bt[0], it, ait, iit)
		default:
			return PowIterIncrF32(at, bt, it, ait, bit, iit)
		}
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		it := incr.Float64s()
		switch {
		case as && bs:
			VecPowF64(at, bt)
			if !is {
				return e.PowIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return PowIterIncrSVF64(at[0], bt, it, bit, iit)
		case !as && bs:
			return PowIterIncrVSF64(at, bt[0], it, ait, iit)
		default:
			return PowIterIncrF64(at, bt, it, ait, bit, iit)
		}
	case Complex64:
		at := a.Complex64s()
		bt := b.Complex64s()
		it := incr.Complex64s()
		switch {
		case as && bs:
			VecPowC64(at, bt)
			if !is {
				return e.PowIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return PowIterIncrSVC64(at[0], bt, it, bit, iit)
		case !as && bs:
			return PowIterIncrVSC64(at, bt[0], it, ait, iit)
		default:
			return PowIterIncrC64(at, bt, it, ait, bit, iit)
		}
	case Complex128:
		at := a.Complex128s()
		bt := b.Complex128s()
		it := incr.Complex128s()
		switch {
		case as && bs:
			VecPowC128(at, bt)
			if !is {
				return e.PowIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return PowIterIncrSVC128(at[0], bt, it, bit, iit)
		case !as && bs:
			return PowIterIncrVSC128(at, bt[0], it, ait, iit)
		default:
			return PowIterIncrC128(at, bt, it, ait, bit, iit)
		}
	default:
		return errors.Errorf("Unsupported type %v for Pow", t)
	}
}

func (e E) ModIterIncr(t reflect.Type, a *storage.Header, b *storage.Header, incr *storage.Header, ait Iterator, bit Iterator, iit Iterator) (err error) {
	as := isScalar(a)
	bs := isScalar(b)
	is := isScalar(incr)

	if ((as && !bs) || (bs && !as)) && is {
		return errors.Errorf("Cannot increment on a scalar increment. len(a): %d, len(b) %d", a.Len(), b.Len())
	}

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()
		it := incr.Ints()
		switch {
		case as && bs:
			VecModI(at, bt)
			if !is {
				return e.ModIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return ModIterIncrSVI(at[0], bt, it, bit, iit)
		case !as && bs:
			return ModIterIncrVSI(at, bt[0], it, ait, iit)
		default:
			return ModIterIncrI(at, bt, it, ait, bit, iit)
		}
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		it := incr.Int8s()
		switch {
		case as && bs:
			VecModI8(at, bt)
			if !is {
				return e.ModIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return ModIterIncrSVI8(at[0], bt, it, bit, iit)
		case !as && bs:
			return ModIterIncrVSI8(at, bt[0], it, ait, iit)
		default:
			return ModIterIncrI8(at, bt, it, ait, bit, iit)
		}
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		it := incr.Int16s()
		switch {
		case as && bs:
			VecModI16(at, bt)
			if !is {
				return e.ModIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return ModIterIncrSVI16(at[0], bt, it, bit, iit)
		case !as && bs:
			return ModIterIncrVSI16(at, bt[0], it, ait, iit)
		default:
			return ModIterIncrI16(at, bt, it, ait, bit, iit)
		}
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		it := incr.Int32s()
		switch {
		case as && bs:
			VecModI32(at, bt)
			if !is {
				return e.ModIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return ModIterIncrSVI32(at[0], bt, it, bit, iit)
		case !as && bs:
			return ModIterIncrVSI32(at, bt[0], it, ait, iit)
		default:
			return ModIterIncrI32(at, bt, it, ait, bit, iit)
		}
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		it := incr.Int64s()
		switch {
		case as && bs:
			VecModI64(at, bt)
			if !is {
				return e.ModIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return ModIterIncrSVI64(at[0], bt, it, bit, iit)
		case !as && bs:
			return ModIterIncrVSI64(at, bt[0], it, ait, iit)
		default:
			return ModIterIncrI64(at, bt, it, ait, bit, iit)
		}
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		it := incr.Uints()
		switch {
		case as && bs:
			VecModU(at, bt)
			if !is {
				return e.ModIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return ModIterIncrSVU(at[0], bt, it, bit, iit)
		case !as && bs:
			return ModIterIncrVSU(at, bt[0], it, ait, iit)
		default:
			return ModIterIncrU(at, bt, it, ait, bit, iit)
		}
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		it := incr.Uint8s()
		switch {
		case as && bs:
			VecModU8(at, bt)
			if !is {
				return e.ModIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return ModIterIncrSVU8(at[0], bt, it, bit, iit)
		case !as && bs:
			return ModIterIncrVSU8(at, bt[0], it, ait, iit)
		default:
			return ModIterIncrU8(at, bt, it, ait, bit, iit)
		}
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		it := incr.Uint16s()
		switch {
		case as && bs:
			VecModU16(at, bt)
			if !is {
				return e.ModIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return ModIterIncrSVU16(at[0], bt, it, bit, iit)
		case !as && bs:
			return ModIterIncrVSU16(at, bt[0], it, ait, iit)
		default:
			return ModIterIncrU16(at, bt, it, ait, bit, iit)
		}
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		it := incr.Uint32s()
		switch {
		case as && bs:
			VecModU32(at, bt)
			if !is {
				return e.ModIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return ModIterIncrSVU32(at[0], bt, it, bit, iit)
		case !as && bs:
			return ModIterIncrVSU32(at, bt[0], it, ait, iit)
		default:
			return ModIterIncrU32(at, bt, it, ait, bit, iit)
		}
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		it := incr.Uint64s()
		switch {
		case as && bs:
			VecModU64(at, bt)
			if !is {
				return e.ModIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return ModIterIncrSVU64(at[0], bt, it, bit, iit)
		case !as && bs:
			return ModIterIncrVSU64(at, bt[0], it, ait, iit)
		default:
			return ModIterIncrU64(at, bt, it, ait, bit, iit)
		}
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		it := incr.Float32s()
		switch {
		case as && bs:
			VecModF32(at, bt)
			if !is {
				return e.ModIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return ModIterIncrSVF32(at[0], bt, it, bit, iit)
		case !as && bs:
			return ModIterIncrVSF32(at, bt[0], it, ait, iit)
		default:
			return ModIterIncrF32(at, bt, it, ait, bit, iit)
		}
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		it := incr.Float64s()
		switch {
		case as && bs:
			VecModF64(at, bt)
			if !is {
				return e.ModIter(t, incr, a, iit, ait)
			}
			it[0] += at[0]
			return
		case as && !bs:
			return ModIterIncrSVF64(at[0], bt, it, bit, iit)
		case !as && bs:
			return ModIterIncrVSF64(at, bt[0], it, ait, iit)
		default:
			return ModIterIncrF64(at, bt, it, ait, bit, iit)
		}
	default:
		return errors.Errorf("Unsupported type %v for Mod", t)
	}
}
