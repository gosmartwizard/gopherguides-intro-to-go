package week05

import (
	"testing"
)

func assertModel(t *testing.T, act Model, exp Model) {
	for i, m := range exp {
		if m != act[i] {
			t.Fatalf("expected : %#v, got : %#v", m, act[i])
		}
	}
}

func Test_Store_Insert(t *testing.T) {
	t.Parallel()

	m1 := Model{
		"iPhone": "Iphone5",
	}

	m2 := Model{
		"iPhone": "Iphone6",
	}

	ms := Models{m1, m2}

	s := &Store{}

	s.Insert("Mobiles", m1, m2)

	for k, v1 := range s.data {
		if k != "Mobiles" {
			t.Fatalf("expected: %#v,and got: %#v", "Mobiles", k)
		}

		for i, v2 := range ms {
			assertModel(t, v1[i], v2)
		}
	}
}

func Test_Store_All_1(t *testing.T) {
	t.Parallel()

	m1 := Model{
		"iPhone": "Iphone5",
	}

	m2 := Model{
		"iPhone": "Iphone6",
	}

	ms := Models{m1, m2}

	s := &Store{}

	s.Insert("Mobiles", m1, m2)

	mods, err := s.All("Mobiles")

	if err != nil {
		t.Fatalf("expected: nil, got error: %#v", err)
	}

	for i, m := range mods {
		assertModel(t, m, ms[i])
	}
}

func Test_Store_All_2(t *testing.T) {
	t.Parallel()

	m1 := Model{
		"iPhone": "Iphone5",
	}

	m2 := Model{
		"iPhone": "Iphone6",
	}

	s := &Store{}

	s.Insert("Mobiles", m1, m2)

	_, err := s.All("Laptops")

	exp := &ErrTableNotFound{
		table: "Laptops",
	}

	if ok := IsErrTableNotFound(err); !ok {
		t.Fatalf("expected: %#v,and got: %#v", exp, err)
	}
}

func Test_Store_All_3(t *testing.T) {
	t.Parallel()

	m1 := Model{
		"iPhone": "Iphone5",
	}

	m2 := Model{
		"iPhone": "Iphone6",
	}

	s := &Store{}

	s.Insert("Mobiles", m1, m2)

	_, err := s.All("Laptops")

	exp := &ErrTableNotFound{
		table: "Laptops",
	}

	if ok := IsErrTableNotFound(err); !ok {
		t.Fatalf("expected: %#v,and got: %#v", exp, err)
	}

	if exp.Error() != err.Error() {
		t.Fatalf("expected: %#v,and got: %#v", exp, err)
	}
}

func Test_Store_Len_1(t *testing.T) {
	t.Parallel()

	m1 := Model{
		"iPhone": "Iphone5",
	}

	m2 := Model{
		"iPhone": "Iphone6",
	}

	s := Store{}

	s.Insert("Mobiles", m1, m2)

	len, err := s.Len("Mobiles")

	if err != nil {
		t.Fatalf("expected: nil, got error: %#v", err)
	}

	exp := 2

	if exp != len {
		t.Fatalf("expected: %#v,and got: %#v", exp, len)
	}
}

func Test_Store_Len_2(t *testing.T) {
	t.Parallel()

	m1 := Model{
		"iPhone": "Iphone5",
	}

	m2 := Model{
		"iPhone": "Iphone6",
	}

	s := Store{}

	s.Insert("Mobiles", m1, m2)

	len, err := s.Len("Laptops")

	exp := &ErrTableNotFound{
		table: "Laptops",
	}

	if ok := IsErrTableNotFound(err); !ok {
		t.Fatalf("expected: %#v,and got: %#v", exp, err)
	}

	expLen := 0

	if expLen != len {
		t.Fatalf("expected: %#v,and got: %#v", expLen, len)
	}
}

func Test_Store_Select_1(t *testing.T) {
	t.Parallel()

	m := Model{
		"iPhone": "Iphone5",
	}

	c := Clauses{
		"iPhone": "Iphone5",
	}

	s := &Store{}

	s.Insert("Mobiles", m)

	mods, err := s.Select("Mobiles", c)

	if err != nil {
		t.Fatalf("expected: nil, got error: %#v", err)
	}

	exp := Models{m}

	for i, m := range mods {
		assertModel(t, m, exp[i])
	}
}

func Test_Store_Select_2(t *testing.T) {
	t.Parallel()

	m := Model{
		"iPhone": "Iphone5",
	}

	c := Clauses{
		"iPhone": "Iphone5",
	}

	s := &Store{}

	s.Insert("Mobiles", m)

	_, err := s.Select("Laptops", c)

	if err == nil {
		t.Fatal("expected error , got nil")
	}
}

func Test_Store_Select_3(t *testing.T) {
	t.Parallel()

	m := Model{
		"iPhone": "Iphone5",
	}

	c := Clauses{}

	s := &Store{}

	s.Insert("Mobiles", m)

	mods, err := s.Select("Mobiles", c)

	if err != nil {
		t.Fatalf("expected: nil, got error: %#v", err)
	}

	exp := Models{m}

	for i, m := range mods {
		assertModel(t, m, exp[i])
	}
}

func Test_Store_Select_4(t *testing.T) {
	t.Parallel()

	m := Model{
		"iPhone": "Iphone5",
	}

	c := Clauses{
		"iPhone": "Iphone6",
	}

	s := &Store{}

	s.Insert("Mobiles", m)

	_, err := s.Select("Mobiles", c)

	if ok := IsErrNoRows(err); !ok {
		t.Fatalf("expected: %#v,and got: %#v", "IsErrNoRows", err)
	}
}
