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

	err := m.Perform(v)

	if err != nil {
		t.Error(err)
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

	exp := m.Name()

	if exp != m.Groupname {
		t.Errorf("Expected : %s, act : %s", exp, m.Groupname)
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

	err := m.Setup(v)

	if err != nil {
		t.Error(err)
	}
}
