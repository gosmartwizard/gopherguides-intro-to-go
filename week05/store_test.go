package week05

import (
	"reflect"
	"testing"
)

func Test_Insert_1(t *testing.T) {
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

	for k, v := range s.data {
		if k != "Mobiles" {
			t.Fatalf("expected: %#v,and got: %#v", "Mobiles", k)
		}

		if !reflect.DeepEqual(ms, v) {
			t.Fatalf("expected: %#v,and got: %#v", ms, v)
		}
	}

}

func Test_All_1(t *testing.T) {
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

	if !reflect.DeepEqual(ms, mods) {
		t.Fatalf("expected: %#v,and got: %#v", ms, mods)
	}
}

func Test_All_2(t *testing.T) {
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

	if ok := IsErrTableNotFound(err); !ok {
		t.Fatalf("expected: %#v,and got: %#v", "ErrTableNotFound", err)
	}

	exp := &ErrTableNotFound{
		table: "Laptops",
	}

	if !reflect.DeepEqual(exp, err) {
		t.Fatalf("expected: %#v,and got: %#v", exp, err)
	}
}

func Test_All_3(t *testing.T) {
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

	if ok := IsErrTableNotFound(err); !ok {
		t.Fatalf("expected: %#v,and got: %#v", "ErrTableNotFound", err)
	}

	exp := &ErrTableNotFound{
		table: "Laptops",
	}

	if !reflect.DeepEqual(exp.Error(), err.Error()) {
		t.Fatalf("expected: %#v,and got: %#v", exp, err)
	}
}

func Test_Len_1(t *testing.T) {
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
	if !reflect.DeepEqual(exp, len) {
		t.Fatalf("expected: %#v,and got: %#v", exp, len)
	}
}

func Test_Len_2(t *testing.T) {
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

	if ok := IsErrTableNotFound(err); !ok {
		t.Fatalf("expected: %#v,and got: %#v", "ErrTableNotFound", err)
	}

	if !reflect.DeepEqual(0, len) {
		t.Fatalf("expected: %#v,and got: %#v", 0, len)
	}
}

func Test_Select_1(t *testing.T) {
	t.Parallel()

	m := Model{
		"iPhone": "Iphone5",
	}

	c := Clauses{
		"iPhone": "Iphone5",
	}

	s := &Store{}

	s.Insert("Mobiles", m)

	ms, err := s.Select("Mobiles", c)

	if err != nil {
		t.Fatalf("expected: nil, got error: %#v", err)
	}

	exp := Models{m}

	if !reflect.DeepEqual(exp, ms) {
		t.Fatalf("expected: %#v,and got: %#v", exp, ms)
	}
}

func Test_Select_2(t *testing.T) {
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

func Test_Select_3(t *testing.T) {
	t.Parallel()

	m := Model{
		"iPhone": "Iphone5",
	}

	c := Clauses{}

	s := &Store{}

	s.Insert("Mobiles", m)

	ms, err := s.Select("Mobiles", c)

	if err != nil {
		t.Fatalf("expected: nil, got error: %#v", err)
	}

	exp := Models{m}

	if !reflect.DeepEqual(exp, ms) {
		t.Fatalf("expected: %#v,and got: %#v", exp, ms)
	}
}

func Test_Select_4(t *testing.T) {
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

	exp := &errNoRows{
		clauses: c,
		table:   "Mobiles",
	}

	if !reflect.DeepEqual(exp, err) {
		t.Fatalf("expected: %#v,and got: %#v", exp, err)
	}
}
