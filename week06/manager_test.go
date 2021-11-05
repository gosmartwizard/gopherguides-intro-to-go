package week06

import (
	"testing"
)

func Test_Manager_1(t *testing.T) {
	t.Parallel()

	m := NewManager()

	err := m.Start(0)

	exp := ErrInvalidEmployeeCount(0)

	if exp.Error() != err.Error() {
		t.Fatalf("expected : %#v, got : %#v", exp, err)
	}
}

func Test_Manager_2(t *testing.T) {
	t.Parallel()

	m := NewManager()

	err := m.Start(-9)

	exp := ErrInvalidEmployeeCount(-9)

	if exp.Error() != err.Error() {
		t.Fatalf("expected : %#v, got : %#v", exp, err)
	}
}

func Test_Manager_3(t *testing.T) {
	t.Parallel()

	m := NewManager()

	e := Employee(-9)

	p := &Product{
		Quantity: 2,
		builtBy:  1,
	}

	err := m.Complete(e, p)

	exp := ErrInvalidEmployee(-9)

	if exp.Error() != err.Error() {
		t.Fatalf("expected : %#v, got : %#v", exp, err)
	}
}

func Test_Manager_4(t *testing.T) {
	t.Parallel()

	m := NewManager()

	e := Employee(1)

	p := &Product{
		Quantity: 0,
		builtBy:  e,
	}

	err := m.Complete(e, p)

	exp := ErrInvalidQuantity(p.Quantity)

	if exp.Error() != err.Error() {
		t.Fatalf("expected : %#v, got : %#v", exp, err)
	}
}

func Test_Manager_5(t *testing.T) {
	t.Parallel()

	m := NewManager()

	e := Employee(1)

	p := &Product{
		Quantity: 2,
		builtBy:  e,
	}

	m.Stop()

	err := m.Assign(p)

	exp := ErrManagerStopped{}

	if exp.Error() != err.Error() {
		t.Fatalf("expected : %#v, got : %#v", exp, err)
	}
}

func Test_Manager_6(t *testing.T) {
	t.Parallel()

	m := NewManager()

	m.Stop()

	m.Stop()
}

func Test_Manager_7(t *testing.T) {
	t.Parallel()

	m := NewManager()

	e := Employee(1)

	p := &Product{
		Quantity: 0,
		builtBy:  e,
	}

	err := m.Assign(p)

	exp := ErrInvalidQuantity(p.Quantity)

	if exp.Error() != err.Error() {
		t.Fatalf("expected : %#v, got : %#v", exp, err)
	}
}
