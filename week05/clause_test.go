package week05

import (
	"reflect"
	"testing"
)

func Test_String_1(t *testing.T) {
	t.Parallel()

	c := Clauses{
		"NaveenReddy": "IshaanReddy",
		"KunaReddy":   "KarthikeyaReddy",
	}

	exp := `"KunaReddy" = "KarthikeyaReddy" and "NaveenReddy" = "IshaanReddy"`
	act := c.String()

	if exp != act {
		t.Fatalf("expected : %s  and actual : %s", exp, act)
	}
}

func Test_String_2(t *testing.T) {
	t.Parallel()

	c := Clauses{
		"Naveen": 99,
		"Kuna":   49,
	}

	exp := `"Kuna" = '1' and "Naveen" = 'c'`
	act := c.String()

	if exp != act {
		t.Fatalf("expected : %s  and actual : %s", exp, act)
	}
}

func Test_String_3(t *testing.T) {
	t.Parallel()

	c := Clauses{
		"Naveen": "",
		"Kuna":   "",
	}

	exp := `"Kuna" = "" and "Naveen" = ""`
	act := c.String()

	if exp != act {
		t.Fatalf("expected : %s  and actual : %s", exp, act)
	}
}

func Test_String_4(t *testing.T) {
	t.Parallel()

	c := Clauses{
		"NaveenReddy": "Kunareddy",
		"HNo":         50,
	}

	exp := `"HNo" = '2' and "NaveenReddy" = "Kunareddy"`
	act := c.String()

	if exp != act {
		t.Fatalf("expected : %s  and actual : %s", exp, act)
	}
}

func Test_String_5(t *testing.T) {
	t.Parallel()

	c := Clauses{}

	exp := ""
	act := c.String()

	if exp != act {
		t.Fatalf("expected : %s  and actual : %s", exp, act)
	}
}

func Test_TableDrivenTests_String(t *testing.T) {
	t.Parallel()

	table := []struct {
		description string
		clause      Clauses
		expected    string
	}{
		{
			description: "Test_String_1",
			clause:      Clauses{"NaveenReddy": "IshaanReddy", "KunaReddy": "KarthikeyaReddy"},
			expected:    `"KunaReddy" = "KarthikeyaReddy" and "NaveenReddy" = "IshaanReddy"`,
		},
		{
			description: "Test_String_2",
			clause:      Clauses{"Naveen": 99, "Kuna": 49},
			expected:    `"Kuna" = '1' and "Naveen" = 'c'`,
		},
		{
			description: "Test_String_3",
			clause:      Clauses{"Naveen": "", "Kuna": ""},
			expected:    `"Kuna" = "" and "Naveen" = ""`,
		},
		{
			description: "Test_String_4",
			clause:      Clauses{"NaveenReddy": "Kunareddy", "HNo": 50},
			expected:    `"HNo" = '2' and "NaveenReddy" = "Kunareddy"`,
		},
		{
			description: "Test_String_5",
			clause:      Clauses{},
			expected:    "",
		},
	}

	for _, tt := range table {
		t.Run(tt.description, func(t *testing.T) {
			got := tt.clause.String()

			if tt.expected != got {
				t.Fatalf("expected: %#v, got: %#v", tt.expected, got)
			}

			if !reflect.DeepEqual(tt.expected, got) {
				t.Fatalf("expected: %#v, got: %#v", tt.expected, got)
			}
		})
	}
}

func Test_Match_1(t *testing.T) {
	t.Parallel()

	c := Clauses{
		"NaveenReddy": "IshaanReddy",
		"KunaReddy":   "KarthikeyaReddy",
	}

	m := Model{
		"NaveenReddy": "IshaanReddy",
		"KunaReddy":   "KarthikeyaReddy",
	}

	exp := true
	act := c.Match(m)

	if exp != act {
		t.Fatalf("expected : %v  and actual : %v", exp, act)
	}
}

func Test_Match_2(t *testing.T) {
	t.Parallel()

	c := Clauses{
		"NaveenReddy": 49,
		"KunaReddy":   99,
	}

	m := Model{
		"NaveenReddy": 49,
		"KunaReddy":   50,
	}

	exp := false
	act := c.Match(m)

	if exp != act {
		t.Fatalf("expected : %v  and actual : %v", exp, act)
	}
}

func Test_Match_3(t *testing.T) {
	t.Parallel()

	c := Clauses{}

	m := Model{}

	exp := true
	act := c.Match(m)

	if exp != act {
		t.Fatalf("expected : %v  and actual : %v", exp, act)
	}
}

func Test_Match_4(t *testing.T) {
	t.Parallel()

	c := Clauses{
		"NaveenReddy": "",
		"KunaReddy":   99,
	}

	m := Model{
		"NaveenReddy": "",
		"KunaReddy":   99,
	}

	exp := true
	act := c.Match(m)

	if exp != act {
		t.Fatalf("expected : %v  and actual : %v", exp, act)
	}
}

func Test_Match_5(t *testing.T) {
	t.Parallel()

	c := Clauses{
		"NaveenReddy": "IshaanReddy",
		"KunaReddy":   "KarthikeyaReddy",
	}

	m := Model{
		"NaveenReddy": "IshaanReddy",
		"KunaReddy":   "KarthikeyaReddy",
		"Ravanamma":   "Kunareddy",
	}

	exp := true
	act := c.Match(m)

	if exp != act {
		t.Fatalf("expected : %v  and actual : %v", exp, act)
	}
}

func Test_Match_6(t *testing.T) {
	t.Parallel()

	c := Clauses{
		"NaveenReddy": "IshaanReddy",
		"KunaReddy":   "KarthikeyaReddy",
		"Ravanamma":   "Kunareddy",
	}

	m := Model{
		"NaveenReddy": "IshaanReddy",
		"KunaReddy":   "KarthikeyaReddy",
	}

	exp := false
	act := c.Match(m)

	if exp != act {
		t.Fatalf("expected : %v  and actual : %v", exp, act)
	}
}

func Test_Match_7(t *testing.T) {
	t.Parallel()

	c := Clauses{
		"NaveenReddy": "IshaanReddy",
		"KunaReddy":   "KarthikeyaReddy",
		"Ravanamma":   "Kunareddy",
	}

	m := Model{}

	exp := false
	act := c.Match(m)

	if exp != act {
		t.Fatalf("expected : %v  and actual : %v", exp, act)
	}
}

func Test_Match_8(t *testing.T) {
	t.Parallel()

	c := Clauses{}

	m := Model{
		"NaveenReddy": "IshaanReddy",
		"KunaReddy":   "KarthikeyaReddy",
		"Ravanamma":   "Kunareddy",
	}

	exp := true
	act := c.Match(m)

	if exp != act {
		t.Fatalf("expected : %v  and actual : %v", exp, act)
	}
}
