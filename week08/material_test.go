package week08

import (
	"testing"
	"time"
)

func Test_Materials_Duration(t *testing.T) {
	t.Parallel()

	m := Materials{
		Metal:   1,
		Oil:     2,
		Plastic: 3,
		Wood:    4,
	}

	d := m.Duration()

	exp := time.Duration(48000000)
	if exp != d {
		t.Fatalf("expected : %#v, got : %#v", exp, d)
	}
}
