package week06

import (
	"testing"
)

func Test_Product_1(t *testing.T) {
	t.Parallel()

	p := Product{
		Quantity: 2,
		builtBy:  Employee(1),
	}

	err := p.IsValid()

	if err != nil {
		t.Fatalf("expected : nil, got : %#v", err)
	}
}

func Test_Product_2(t *testing.T) {
	t.Parallel()

	p := Product{
		Quantity: 0,
		builtBy:  Employee(1),
	}

	err := p.IsValid()

	exp := ErrInvalidQuantity(p.Quantity)

	if exp.Error() != err.Error() {
		t.Fatalf("expected : %#v, got : %#v", exp, err)
	}
}

func Test_Product_3(t *testing.T) {
	t.Parallel()

	p := Product{
		Quantity: -9,
		builtBy:  Employee(1),
	}

	err := p.IsValid()

	exp := ErrInvalidQuantity(p.Quantity)

	if exp.Error() != err.Error() {
		t.Fatalf("expected : %#v, got : %#v", exp, err)
	}
}

func Test_Product_4(t *testing.T) {
	t.Parallel()

	p := Product{
		Quantity: 2,
		builtBy:  1,
	}

	e := p.BuiltBy()

	exp := Employee(1)

	if exp != e {
		t.Fatalf("expected : %#v, got : %#v", exp, e)
	}
}

func Test_Product_5(t *testing.T) {
	t.Parallel()

	e := Employee(-9)

	p := Product{
		Quantity: 2,
		builtBy:  1,
	}

	err := p.Build(e)

	exp := ErrInvalidEmployee(-9)

	if exp.Error() != err.Error() {
		t.Fatalf("expected : %#v, got : %#v", exp, err)
	}
}

func Test_Product_6(t *testing.T) {
	t.Parallel()

	e := Employee(1)

	p := Product{
		Quantity: 0,
		builtBy:  e,
	}

	err := p.Build(e)

	exp := ErrInvalidQuantity(p.Quantity)

	if exp.Error() != err.Error() {
		t.Fatalf("expected : %#v, got : %#v", exp, err)
	}
}

func Test_Product_7(t *testing.T) {
	t.Parallel()

	e := Employee(1)

	p := Product{
		Quantity: -9,
		builtBy:  e,
	}

	err := p.Build(e)

	exp := ErrInvalidQuantity(p.Quantity)

	if exp.Error() != err.Error() {
		t.Fatalf("expected : %#v, got : %#v", exp, err)
	}
}
