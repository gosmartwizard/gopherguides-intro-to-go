package week05

import (
	"reflect"
	"testing"
)

func Test_Clauses_String(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		description string
		clause      Clauses
		expected    string
	}{
		{
			description: "Test_Clauses_String_1",
			clause:      Clauses{"NaveenReddy": "IshaanReddy", "KunaReddy": "KarthikeyaReddy"},
			expected:    `"KunaReddy" = "KarthikeyaReddy" and "NaveenReddy" = "IshaanReddy"`,
		},
		{
			description: "Test_Clauses_String_2",
			clause:      Clauses{"Naveen": 99, "Kuna": 49},
			expected:    `"Kuna" = '1' and "Naveen" = 'c'`,
		},
		{
			description: "Test_Clauses_String_3",
			clause:      Clauses{"Naveen": "", "Kuna": ""},
			expected:    `"Kuna" = "" and "Naveen" = ""`,
		},
		{
			description: "Test_Clauses_String_4",
			clause:      Clauses{"NaveenReddy": "Kunareddy", "HNo": 50},
			expected:    `"HNo" = '2' and "NaveenReddy" = "Kunareddy"`,
		},
		{
			description: "Test_Clauses_String_5",
			clause:      Clauses{},
			expected:    "",
		},
	}

	for _, tc := range tcs {
		tc := tc

		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			got := tc.clause.String()

			if tc.expected != got {
				t.Fatalf("expected: %#v, got: %#v", tc.expected, got)
			}
		})
	}
}

func Test_Clauses_Match(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		description string
		clause      Clauses
		model       Model
		expected    bool
	}{
		{
			description: "Test_Clauses_Match_1",
			clause:      Clauses{"NaveenReddy": "IshaanReddy", "KunaReddy": "KarthikeyaReddy"},
			model:       Model{"NaveenReddy": "IshaanReddy", "KunaReddy": "KarthikeyaReddy"},
			expected:    true,
		},
		{
			description: "Test_Clauses_Match_2",
			clause:      Clauses{"NaveenReddy": 49, "KunaReddy": 99},
			model:       Model{"NaveenReddy": 49, "KunaReddy": 50},
			expected:    false,
		},
		{
			description: "Test_Clauses_Match_3",
			clause:      Clauses{},
			model:       Model{},
			expected:    true,
		},
		{
			description: "Test_Clauses_Match_4",
			clause:      Clauses{"NaveenReddy": "", "KunaReddy": 99},
			model:       Model{"NaveenReddy": "", "KunaReddy": 99},
			expected:    true,
		},
		{
			description: "Test_Clauses_Match_5",
			clause:      Clauses{"NaveenReddy": "IshaanReddy", "KunaReddy": "KarthikeyaReddy"},
			model:       Model{"NaveenReddy": "IshaanReddy", "KunaReddy": "KarthikeyaReddy", "Ravanamma": "Kunareddy"},
			expected:    true,
		},
		{
			description: "Test_Clauses_Match_6",
			clause:      Clauses{"NaveenReddy": "IshaanReddy", "KunaReddy": "KarthikeyaReddy", "Ravanamma": "Kunareddy"},
			model:       Model{"NaveenReddy": "IshaanReddy", "KunaReddy": "KarthikeyaReddy"},
			expected:    false,
		},
		{
			description: "Test_Clauses_Match_7",
			clause:      Clauses{"NaveenReddy": "IshaanReddy", "KunaReddy": "KarthikeyaReddy", "Ravanamma": "Kunareddy"},
			model:       Model{},
			expected:    false,
		},
		{
			description: "Test_Clauses_Match_8",
			clause:      Clauses{},
			model:       Model{"NaveenReddy": "IshaanReddy", "KunaReddy": "KarthikeyaReddy", "Ravanamma": "Kunareddy"},
			expected:    true,
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			got := tc.clause.Match(tc.model)

			if !reflect.DeepEqual(tc.expected, got) {
				t.Fatalf("expected: %#v, got: %#v", tc.expected, got)
			}
		})
	}
}
