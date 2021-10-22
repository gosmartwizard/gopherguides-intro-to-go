package week04

import (
	"bytes"
	"testing"
)

func Test_Venue_1(t *testing.T) {

	t.Parallel()

	var v Venue

	err := v.Entertain(50)

	if err == nil {
		t.Errorf("Expected error but returned nil")
	}
}

func Test_Venue_2(t *testing.T) {

	t.Parallel()

	d := Dance{
		Groupname: "Dhee2021",
		Members:   9,
		Theme:     "Rock",
	}

	m := Music{
		Groupname: "Raaga2021",
		Members:   5,
		Theme:     "Melody",
	}

	y := Yoga{
		Groupname: "Isha2021",
		Members:   1,
		Theme:     "Angamardhana",
	}

	var buf bytes.Buffer
	var v Venue
	v.Log = &buf

	err := v.Entertain(100, d, m, y)

	if err != nil {
		t.Error(err)
	}
}

func Test_Play_1(t *testing.T) {

	t.Parallel()

	y := Yoga{
		Groupname: "Isha2021",
		Members:   1,
		Theme:     "Angamardhana",
	}

	var buf bytes.Buffer
	var v Venue
	v.Log = &buf

	err := v.play(y)

	if err != nil {
		t.Error(err)
	}
}

func Test_Play_2(t *testing.T) {

	t.Parallel()

	m := Music{
		Groupname: "Raaga2021",
		Members:   5,
		Theme:     "Melody",
	}

	var buf bytes.Buffer
	var v Venue
	v.Log = &buf

	err := v.play(m)

	if err != nil {
		t.Error(err)
	}
}

func Test_Play_3(t *testing.T) {

	t.Parallel()

	d := Dance{
		Groupname: "Dhee2021",
		Members:   9,
		Theme:     "Rock",
	}

	var buf bytes.Buffer
	var v Venue
	v.Log = &buf

	err := v.play(d)

	if err != nil {
		t.Error(err)
	}
}
