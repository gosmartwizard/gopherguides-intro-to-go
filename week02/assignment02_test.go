package week02

import (
	"reflect"
	"testing"
)

func Test_Array_1(t *testing.T) {

	exp := [4]string{"John", "Paul", "George", "Ringo"}
	act := [4]string{}

	for i, v := range exp {
		act[i] = v
	}

	for i, v := range act {
		if v != exp[i] {
			t.Errorf("act[%d] : %s is not equal to exp[%d] : %s ", i, v, i, exp[i])
		}
	}
}

func Test_Array_2(t *testing.T) {

	exp := [4]string{"John", "Paul", "George", "Ringo"}
	act := [4]string{}

	for i, v := range exp {
		act[i] = v
	}

	if b := reflect.DeepEqual(act, exp); !b {
		t.Error("act contents : ", act, " and exp contents : ", exp, "are not equal")
	}
}

func Test_Array_3(t *testing.T) {

	exp := [4]string{"John", "Paul", "George", "Ringo"}
	act := [4]string{}

	for i, v := range exp {
		act[i] = v
	}

	la := len(act)
	le := len(exp)
	if la != le {
		t.Errorf("len of act : %d is not equal to len of exp : %d ", la, le)
	}
}

func Test_Slice_1(t *testing.T) {

	exp := []string{"John", "Paul", "George", "Ringo"}
	act := make([]string, 0, len(exp))

	for _, v := range exp {
		act = append(act, v)
	}

	for i, v := range act {
		if v != exp[i] {
			t.Errorf("act[%d] : %s is not equal to exp[%d] : %s ", i, v, i, exp[i])
		}
	}
}

func Test_Slice_2(t *testing.T) {

	exp := []string{"John", "Paul", "George", "Ringo"}
	var act []string

	for _, v := range exp {
		act = append(act, v)
	}

	if b := reflect.DeepEqual(act, exp); !b {
		t.Error("act contents : ", act, " and exp contents : ", exp, "are not equal")
	}
}

func Test_Slice_3(t *testing.T) {

	exp := []string{"John", "Paul", "George", "Ringo"}
	var act []string

	for _, v := range exp {
		act = append(act, v)
	}

	la := len(act)
	le := len(exp)
	if la != le {
		t.Errorf("len of act : %d is not equal to len of exp : %d ", la, le)
	}
}

func Test_Slice_4(t *testing.T) {

	exp := []string{"John", "Paul", "George", "Ringo"}
	var act []string

	for _, v := range exp {
		act = append(act, v)
	}

	la := len(act)
	le := len(exp)
	if la != le {
		t.Errorf("len of act : %d is not equal to len of exp : %d ", la, le)
	} else {
		for i, v := range act {
			if act[i] != exp[i] {
				t.Errorf("act[%d] : %s is not equal to exp[%d] : %s ", i, v, i, exp[i])
			}
		}
	}
}

func Test_Map_1(t *testing.T) {

	exp := map[string]string{
		"John":   "john@gmail.com",
		"Paul":   "paul@gmail.com",
		"George": "george@gmail.com",
		"Ringo":  "ringo@gmail.com",
	}

	act := make(map[string]string, len(exp))

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

func Test_Map_2(t *testing.T) {

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

func Test_Map_3(t *testing.T) {

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

	la := len(act)
	le := len(exp)
	if la != le {
		t.Errorf("len of act : %d is not equal to len of exp : %d ", la, le)
	}
}

func Test_Map_4(t *testing.T) {

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

	la := len(act)
	le := len(exp)
	if la != le {
		t.Errorf("len of act : %d is not equal to len of exp : %d ", la, le)
	} else {
		for ka, va := range act {
			ve, ok := exp[ka]
			if ok {
				if va != ve {
					t.Errorf("act[%q] : %s is not equal to exp[%q] : %s ", ka, va, ka, ve)
				}
			}
		}
	}
}

