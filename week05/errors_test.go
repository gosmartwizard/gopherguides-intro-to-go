package week05

import "testing"

func Test_TableNotFound_1(t *testing.T) {
	t.Parallel()

	m1 := Model{
		"iPhone": "Iphone5",
	}

	m2 := Model{
		"iPhone": "Iphone6",
	}

	s := Store{}

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
