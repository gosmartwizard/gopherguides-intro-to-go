package week04

import (
	"bytes"
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

	err := d.Perform(v)

	if err != nil {
		t.Error(err)
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

	exp := d.Name()

	if exp != d.Groupname {
		t.Errorf("Expected : %s, act : %s", exp, d.Groupname)
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

	err := d.Perform(v)

	if err != nil {
		t.Error(err)
	}
}
