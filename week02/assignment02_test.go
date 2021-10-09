package week02

import (
	"reflect"
	"testing"
)

func TestArray1(t *testing.T) {

	exp := [4]string{"John", "Paul", "George", "Ringo"}
	act := [4]string{}

	for i, v := range exp {
		act[i] = v
	}

	for i, v := range act {
		if act[i] != exp[i] {
			t.Errorf("act[%d] : %s is not equal to exp[%d] : %s ", i, v, i, exp[i])
		}
	}
}

func TestArray2(t *testing.T) {

	exp := [4]string{"John", "Paul", "George", "Ringo"}
	act := [4]string{}

	for i, v := range exp {
		act[i] = v
	}

	if b := reflect.DeepEqual(act, exp); !b {
		t.Error("act contents : ", act, " and exp contents : ", exp, "are not equal")
	}
}

func TestArray3(t *testing.T) {

	exp := [4]string{"John", "Paul", "George", "Ringo"}
	act := [4]string{}

	for i, v := range exp {
		act[i] = v
	}

	if len(act) != len(exp) {
		t.Errorf("len of act : %d is not equal to len of exp : %d ", len(act), len(exp))
	}
}

func TestSlice1(t *testing.T) {

	exp := []string{"John", "Paul", "George", "Ringo"}
	var act []string

	for _, v := range exp {
		act = append(act, v)
	}

	for i, v := range act {
		if act[i] != exp[i] {
			t.Errorf("act[%d] : %s is not equal to exp[%d] : %s ", i, v, i, exp[i])
		}
	}
}

func TestSlice2(t *testing.T) {

	exp := []string{"John", "Paul", "George", "Ringo"}
	var act []string

	for _, v := range exp {
		act = append(act, v)
	}

	if b := reflect.DeepEqual(act, exp); !b {
		t.Error("act contents : ", act, " and exp contents : ", exp, "are not equal")
	}
}

func TestSlice3(t *testing.T) {

	exp := []string{"John", "Paul", "George", "Ringo"}
	var act []string

	for _, v := range exp {
		act = append(act, v)
	}

	if len(act) != len(exp) {
		t.Errorf("len of act : %d is not equal to len of exp : %d ", len(act), len(exp))
	}
}

func TestMap1(t *testing.T) {

	exp := map[string]string{
		"John":   "john@gmail.com",
		"Paul":   "paul@gmail.com",
		"George": "george@gmail.com",
		"Ringo":  "ringo@gmail.com",
	}

	act := map[string]string{}

	for k, v := range exp {
		act[k] = v
	}

	for ka, va := range act {
		ve, ok := exp[ka]
		if ok {
			if va != ve {
				t.Errorf("act[%q] : %s is not equal to exp[%q] : %s ", ka, va, ka, ve)
			}
		}
	}
}

func TestMap2(t *testing.T) {

	exp := map[string]string{
		"John":   "john@gmail.com",
		"Paul":   "paul@gmail.com",
		"George": "george@gmail.com",
		"Ringo":  "ringo@gmail.com",
	}

	act := map[string]string{}


	for k, v := range exp {
		act[k] = v
	}

	if b := reflect.DeepEqual(act, exp); !b {
		t.Error("act contents : ", act, " and exp contents : ", exp, "are not equal")
	}
}

func TestMap3(t *testing.T) {

	exp := map[string]string{
		"John":   "john@gmail.com",
		"Paul":   "paul@gmail.com",
		"George": "george@gmail.com",
		"Ringo":  "ringo@gmail.com",
	}

	act := map[string]string{}

	for k, v := range exp {
		act[k] = v
	}

	if len(act) != len(exp) {
		t.Errorf("len of act : %d is not equal to len of exp : %d ", len(act), len(exp))
	}
}
