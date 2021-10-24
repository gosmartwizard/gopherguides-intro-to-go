package week04

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_Dance_Perform(t *testing.T) {

	t.Parallel()

	d := Dance{
		Groupname: "Dhee2021",
		Members:   9,
		Theme:     "Rock",
	}

	var buf bytes.Buffer
	var v Venue

	v.Log = &buf

	v.Audience = 50

	err := d.Perform(v)

	if err != nil {
		t.Fatal(err)
	}

	exp := "Dhee2021 has performed for 50 people. \n"

	act := fmt.Sprint(v.Log)

	if exp != act {
		t.Fatalf("expected : %s , actual : %s", exp, act)
	}
}

func Test_Dance_Name(t *testing.T) {

	t.Parallel()

	d := Dance{
		Groupname: "Dhee2021",
		Members:   9,
		Theme:     "Rock",
	}

	var buf bytes.Buffer
	var v Venue

	v.Log = &buf

	exp := d.Groupname

	act := d.Name()

	if exp != act {
		t.Fatalf("expected : %s, Actual : %s", exp, act)
	}
}

func Test_Dance_Teardown(t *testing.T) {

	t.Parallel()

	d := Dance{
		Groupname: "Dhee2021",
		Members:   9,
		Theme:     "Rock",
	}

	var buf bytes.Buffer
	var v Venue

	v.Log = &buf

	v.Audience = 50

	err := d.Teardown(v)

	if err != nil {
		t.Fatal(err)
	}

	exp := "Dhee2021 has completed teardown. \n"

	act := fmt.Sprint(v.Log)

	if exp != act {
		t.Fatalf("expected : %s , actual : %s", exp, act)
	}
}
