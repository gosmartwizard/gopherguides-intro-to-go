package week04

import (
	"bytes"
	"testing"
)

func Test_Music_Perform(t *testing.T) {

	t.Parallel()

	m := Music{
		Groupname: "Raaga2021",
		Members:   5,
		Theme:     "Melody",
	}

	var buf bytes.Buffer
	var v Venue
	v.Log = &buf
	v.Audience = 50

	err := m.Perform(v)

	if err != nil {
		t.Error(err)
	}

	exp := "Raaga2021 has performed for 50 people. \n"

	act := buf.String()

	if exp != act {
		t.Fatalf("expected : %s , actual : %s", exp, act)
	}
}

func Test_Music_Name(t *testing.T) {

	t.Parallel()

	m := Music{
		Groupname: "Raaga2021",
		Members:   5,
		Theme:     "Melody",
	}

	var buf bytes.Buffer
	var v Venue

	v.Log = &buf

	exp := m.Groupname

	act := m.Name()

	if exp != act {
		t.Fatalf("expected : %s, Actual : %s", exp, act)
	}
}

func Test_Music_Setup(t *testing.T) {

	t.Parallel()

	m := Music{
		Groupname: "Raaga2021",
		Members:   5,
		Theme:     "Melody",
	}

	var buf bytes.Buffer
	var v Venue

	v.Log = &buf

	v.Audience = 50

	err := m.Setup(v)

	if err != nil {
		t.Fatal(err)
	}

	exp := "Raaga2021 has completed setup. \n"

	act := buf.String()

	if exp != act {
		t.Fatalf("expected : %s , actual : %s", exp, act)
	}
}
