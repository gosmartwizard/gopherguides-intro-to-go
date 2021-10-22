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

	err := y.Perform(v)

	if err != nil {
		t.Error(err)
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

	exp := y.Name()

	if exp != y.Groupname {
		t.Errorf("Expected : %s, act : %s", exp, y.Groupname)
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

	err := y.Perform(v)

	if err != nil {
		t.Error(err)
	}
}
