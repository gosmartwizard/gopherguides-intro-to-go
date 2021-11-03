package week05

import (
	"errors"
	"fmt"
	"testing"
)

func Test_ErrTableNotFound_TableNotFound(t *testing.T) {
	t.Parallel()

	m1 := Model{
		"iPhone": "Iphone5",
	}

	m2 := Model{
		"iPhone": "Iphone6",
	}

	s := &Store{}

	s.Insert("Mobiles", m1, m2)

	tnf := ErrTableNotFound{
		table: "Laptops",
	}

	tb := tnf.TableNotFound()

	data := s.db()

	_, ok := data[tb]

	if ok {
		t.Fatal("expected: false, got: true")
	}
}

func Test_ErrNoRows_Error(t *testing.T) {
	t.Parallel()

	enr := &errNoRows{
		table: "Mobiles",
	}

	err := enr.Error()

	exp := fmt.Sprintf("[%s] no rows found\nquery: ", enr.table)

	if exp != err {
		t.Fatalf("expected: %#v, got: %#v", exp, err)
	}
}

func Test_ErrNoRows_RowNotFound(t *testing.T) {
	t.Parallel()

	enr := &errNoRows{
		table: "iPad",
		clauses: Clauses{
			"name":    "iPadAir2",
			"version": "ios10",
		},
	}

	table, clauses := enr.RowNotFound()

	exp := "iPad"
	act := table

	if exp != act {
		t.Fatalf("expected : %#v, got : %#v", exp, act)
	}

	expLen := len(enr.clauses)
	actLen := len(clauses)

	if expLen != actLen {
		fmt.Printf("expected: %#v, got : %#v", expLen, actLen)
	}
}

func Test_AsErrNoRows_1(t *testing.T) {
	t.Parallel()

	err := errors.New("Other errrr")

	_, ok := AsErrNoRows(err)

	if ok {
		t.Fatalf("expected false, got: true")
	}
}

func Test_AsErrNoRows_2(t *testing.T) {
	t.Parallel()

	err := &errNoRows{}

	e, ok := AsErrNoRows(err)

	if !ok {
		t.Fatalf("expected true, got: false")
	}

	ok = IsErrNoRows(e)

	if !ok {
		t.Fatalf("expected true, got: false")
	}
}
