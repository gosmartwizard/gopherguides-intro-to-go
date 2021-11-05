package week06

import (
	"testing"
)

func Test_Employee_1(t *testing.T) {
	t.Parallel()

	e := Employee(2)

	err := e.IsValid()

	if err != nil {
		t.Fatalf("expected : nil, got : %#v", err)
	}
}

func Test_Employee_2(t *testing.T) {
	t.Parallel()

	e := Employee(0)

	err := e.IsValid()

	exp := ErrInvalidEmployee(0)

	if exp.Error() != err.Error() {
		t.Fatalf("expected : %#v, got : %#v", exp, err)
	}
}

func Test_Employee_3(t *testing.T) {
	t.Parallel()

	e := Employee(-9)

	err := e.IsValid()

	exp := ErrInvalidEmployee(-9)

	if exp.Error() != err.Error() {
		t.Fatalf("expected : %#v, got : %#v", exp, err)
	}
}
