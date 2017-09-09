package tensor

import "github.com/pkg/errors"

/*
GENERATED FILE. DO NOT EDIT
*/

// Add performs t + other elementwise. Both t and other must have the same shape.
// Acceptable FuncOpts are: UseUnsafe(), WithReuse(T), WithIncr(T)
func (t *Dense) Add(other *Dense, opts ...FuncOpt) (retVal *Dense, err error) {

	var ret Tensor
	if t.oe != nil {
		if ret, err = t.oe.Add(t, other, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do Add()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Add")
		}
		return
	}

	if adder, ok := t.e.(Adder); ok {
		if ret, err = adder.Add(t, other, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do Add()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Add")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support Add()")
}

// Sub performs t - other elementwise. Both t and other must have the same shape.
// Acceptable FuncOpts are: UseUnsafe(), WithReuse(T), WithIncr(T)
func (t *Dense) Sub(other *Dense, opts ...FuncOpt) (retVal *Dense, err error) {

	var ret Tensor
	if t.oe != nil {
		if ret, err = t.oe.Sub(t, other, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do Sub()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Sub")
		}
		return
	}

	if suber, ok := t.e.(Suber); ok {
		if ret, err = suber.Sub(t, other, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do Sub()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Sub")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support Sub()")
}

// Mul performs t × other elementwise. Both t and other must have the same shape.
// Acceptable FuncOpts are: UseUnsafe(), WithReuse(T), WithIncr(T)
func (t *Dense) Mul(other *Dense, opts ...FuncOpt) (retVal *Dense, err error) {

	var ret Tensor
	if t.oe != nil {
		if ret, err = t.oe.Mul(t, other, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do Mul()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Mul")
		}
		return
	}

	if muler, ok := t.e.(Muler); ok {
		if ret, err = muler.Mul(t, other, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do Mul()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Mul")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support Mul()")
}

// Div performs t ÷ other elementwise. Both t and other must have the same shape.
// Acceptable FuncOpts are: UseUnsafe(), WithReuse(T), WithIncr(T)
func (t *Dense) Div(other *Dense, opts ...FuncOpt) (retVal *Dense, err error) {

	var ret Tensor
	if t.oe != nil {
		if ret, err = t.oe.Div(t, other, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do Div()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Div")
		}
		return
	}

	if diver, ok := t.e.(Diver); ok {
		if ret, err = diver.Div(t, other, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do Div()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Div")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support Div()")
}

// Pow performs t ^ other elementwise. Both t and other must have the same shape.
// Acceptable FuncOpts are: UseUnsafe(), WithReuse(T), WithIncr(T)
func (t *Dense) Pow(other *Dense, opts ...FuncOpt) (retVal *Dense, err error) {

	var ret Tensor
	if t.oe != nil {
		if ret, err = t.oe.Pow(t, other, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do Pow()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Pow")
		}
		return
	}

	if power, ok := t.e.(Power); ok {
		if ret, err = power.Pow(t, other, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do Pow()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Pow")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support Pow()")
}

// Mod performs t % other elementwise. Both t and other must have the same shape.
// Acceptable FuncOpts are: UseUnsafe(), WithReuse(T), WithIncr(T)
func (t *Dense) Mod(other *Dense, opts ...FuncOpt) (retVal *Dense, err error) {

	var ret Tensor
	if t.oe != nil {
		if ret, err = t.oe.Mod(t, other, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do Mod()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Mod")
		}
		return
	}

	if moder, ok := t.e.(Moder); ok {
		if ret, err = moder.Mod(t, other, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do Mod()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Mod")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support Mod()")
}

// AddScalar performs t + other elementwise. The leftTensor parameter indicates if the tensor is the left operand. Only scalar types are accepted in other.
// Acceptable FuncOpts are: UseUnsafe(), WithReuse(T), WithIncr(T)
func (t *Dense) AddScalar(other interface{}, leftTensor bool, opts ...FuncOpt) (retVal *Dense, err error) {
	var ret Tensor
	if t.oe != nil {
		if ret, err = t.oe.AddScalar(t, other, leftTensor, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do AddScalar()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "AddScalar")
		}
		return
	}

	if adder, ok := t.e.(Adder); ok {
		if ret, err = adder.AddScalar(t, other, leftTensor, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do AddScalar()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "AddScalar")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support AddScalar()")
}

// SubScalar performs t - other elementwise. The leftTensor parameter indicates if the tensor is the left operand. Only scalar types are accepted in other.
// Acceptable FuncOpts are: UseUnsafe(), WithReuse(T), WithIncr(T)
func (t *Dense) SubScalar(other interface{}, leftTensor bool, opts ...FuncOpt) (retVal *Dense, err error) {
	var ret Tensor
	if t.oe != nil {
		if ret, err = t.oe.SubScalar(t, other, leftTensor, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do SubScalar()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "SubScalar")
		}
		return
	}

	if suber, ok := t.e.(Suber); ok {
		if ret, err = suber.SubScalar(t, other, leftTensor, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do SubScalar()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "SubScalar")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support SubScalar()")
}

// MulScalar performs t × other elementwise. The leftTensor parameter indicates if the tensor is the left operand. Only scalar types are accepted in other.
// Acceptable FuncOpts are: UseUnsafe(), WithReuse(T), WithIncr(T)
func (t *Dense) MulScalar(other interface{}, leftTensor bool, opts ...FuncOpt) (retVal *Dense, err error) {
	var ret Tensor
	if t.oe != nil {
		if ret, err = t.oe.MulScalar(t, other, leftTensor, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do MulScalar()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "MulScalar")
		}
		return
	}

	if muler, ok := t.e.(Muler); ok {
		if ret, err = muler.MulScalar(t, other, leftTensor, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do MulScalar()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "MulScalar")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support MulScalar()")
}

// DivScalar performs t ÷ other elementwise. The leftTensor parameter indicates if the tensor is the left operand. Only scalar types are accepted in other.
// Acceptable FuncOpts are: UseUnsafe(), WithReuse(T), WithIncr(T)
func (t *Dense) DivScalar(other interface{}, leftTensor bool, opts ...FuncOpt) (retVal *Dense, err error) {
	var ret Tensor
	if t.oe != nil {
		if ret, err = t.oe.DivScalar(t, other, leftTensor, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do DivScalar()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "DivScalar")
		}
		return
	}

	if diver, ok := t.e.(Diver); ok {
		if ret, err = diver.DivScalar(t, other, leftTensor, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do DivScalar()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "DivScalar")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support DivScalar()")
}

// PowScalar performs t ^ other elementwise. The leftTensor parameter indicates if the tensor is the left operand. Only scalar types are accepted in other.
// Acceptable FuncOpts are: UseUnsafe(), WithReuse(T), WithIncr(T)
func (t *Dense) PowScalar(other interface{}, leftTensor bool, opts ...FuncOpt) (retVal *Dense, err error) {
	var ret Tensor
	if t.oe != nil {
		if ret, err = t.oe.PowScalar(t, other, leftTensor, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do PowScalar()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "PowScalar")
		}
		return
	}

	if power, ok := t.e.(Power); ok {
		if ret, err = power.PowScalar(t, other, leftTensor, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do PowScalar()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "PowScalar")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support PowScalar()")
}

// ModScalar performs t % other elementwise. The leftTensor parameter indicates if the tensor is the left operand. Only scalar types are accepted in other.
// Acceptable FuncOpts are: UseUnsafe(), WithReuse(T), WithIncr(T)
func (t *Dense) ModScalar(other interface{}, leftTensor bool, opts ...FuncOpt) (retVal *Dense, err error) {
	var ret Tensor
	if t.oe != nil {
		if ret, err = t.oe.ModScalar(t, other, leftTensor, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do ModScalar()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "ModScalar")
		}
		return
	}

	if moder, ok := t.e.(Moder); ok {
		if ret, err = moder.ModScalar(t, other, leftTensor, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do ModScalar()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "ModScalar")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support ModScalar()")
}
