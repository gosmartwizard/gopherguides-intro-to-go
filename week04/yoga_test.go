package week04

import (
	"bytes"
	"testing"
)

func Test_Yoga_Perform(t *testing.T) {

	t.Parallel()

	y := Yoga{
		Groupname: "Isha2021",
		Members:   1,
		Theme:     "Angamardhana",
	}

	var buf bytes.Buffer
	var v Venue

	v.Log = &buf

	v.Audience = 50

	err := y.Perform(v)

	if err != nil {
		t.Fatal(err)
	}

	exp := "Isha2021 has performed for 50 people. \n"

	act := buf.String()

	if exp != act {
		t.Fatalf("expected : %s , actual : %s", exp, act)
	}
}

func Test_Yoga_Name(t *testing.T) {

	t.Parallel()

	y := Yoga{
		Groupname: "Isha2021",
		Members:   1,
		Theme:     "Angamardhana",
	}

	var buf bytes.Buffer
	var v Venue

	v.Log = &buf

	exp := y.Groupname

	act := y.Name()

	if exp != act {
		t.Fatalf("expected : %s, Actual : %s", exp, act)
	}
}

func Test_Yoga_Teardown(t *testing.T) {

	t.Parallel()

	y := Yoga{
		Groupname: "Isha2021",
		Members:   1,
		Theme:     "Angamardhana",
	}

	var buf bytes.Buffer
	var v Venue

	v.Log = &buf

	v.Audience = 50

	err := y.Teardown(v)

	if err != nil {
		t.Fatal(err)
	}

	exp := "Isha2021 has completed teardown. \n"

	act := buf.String()

	if exp != act {
		t.Fatalf("expected : %s , actual : %s", exp, act)
	}
}
