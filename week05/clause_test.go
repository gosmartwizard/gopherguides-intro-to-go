package week05

import (
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
			description: "Golden_Path_With_Strings",
			clause:      Clauses{"NaveenReddy": "IshaanReddy", "KunaReddy": "KarthikeyaReddy"},
			expected:    `"KunaReddy" = "KarthikeyaReddy" and "NaveenReddy" = "IshaanReddy"`,
		},
		{
			description: "Golden_Path_With_Numbers",
			clause:      Clauses{"Naveen": 99, "Kuna": 49},
			expected:    `"Kuna" = '1' and "Naveen" = 'c'`,
		},
		{
			description: "Golden_Path_With_Empty_Strings",
			clause:      Clauses{"Naveen": "", "Kuna": ""},
			expected:    `"Kuna" = "" and "Naveen" = ""`,
		},
		{
			description: "Golden_Path_With_String_Number",
			clause:      Clauses{"NaveenReddy": "Kunareddy", "HNo": 50},
			expected:    `"HNo" = '2' and "NaveenReddy" = "Kunareddy"`,
		},
		{
			description: "Golden_Path_With_Empty_Clauses",
			clause:      Clauses{},
			expected:    "",
		},
	}

	for _, tc := range tcs {

		t.Run(tc.description, func(t *testing.T) {

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
			description: "Golden_Path_With_Strings",
			clause:      Clauses{"NaveenReddy": "IshaanReddy", "KunaReddy": "KarthikeyaReddy"},
			model:       Model{"NaveenReddy": "IshaanReddy", "KunaReddy": "KarthikeyaReddy"},
			expected:    true,
		},
		{
			description: "Error_Path_With_Clause_Model",
			clause:      Clauses{"NaveenReddy": 49, "KunaReddy": 99},
			model:       Model{"NaveenReddy": 49, "KunaReddy": 50},
			expected:    false,
		},
		{
			description: "Golden_Path_With_Empty_Clause_Model",
			clause:      Clauses{},
			model:       Model{},
			expected:    true,
		},
		{
			description: "Golden_Path_With_Empty_String_Number",
			clause:      Clauses{"NaveenReddy": "", "KunaReddy": 99},
			model:       Model{"NaveenReddy": "", "KunaReddy": 99},
			expected:    true,
		},
		{
			description: "Golden_Path_With_Model_Clauses",
			clause:      Clauses{"NaveenReddy": "IshaanReddy", "KunaReddy": "KarthikeyaReddy"},
			model:       Model{"NaveenReddy": "IshaanReddy", "KunaReddy": "KarthikeyaReddy", "Ravanamma": "Kunareddy"},
			expected:    true,
		},
		{
			description: "Error_Path_With_Clauses_Model",
			clause:      Clauses{"NaveenReddy": "IshaanReddy", "KunaReddy": "KarthikeyaReddy", "Ravanamma": "Kunareddy"},
			model:       Model{"NaveenReddy": "IshaanReddy", "KunaReddy": "KarthikeyaReddy"},
			expected:    false,
		},
		{
			description: "Error_Path_With_Empty_Model",
			clause:      Clauses{"NaveenReddy": "IshaanReddy", "KunaReddy": "KarthikeyaReddy", "Ravanamma": "Kunareddy"},
			model:       Model{},
			expected:    false,
		},
		{
			description: "Golden_Path_With_Empty_Clauses",
			clause:      Clauses{},
			model:       Model{"NaveenReddy": "IshaanReddy", "KunaReddy": "KarthikeyaReddy", "Ravanamma": "Kunareddy"},
			expected:    true,
		},
	}

	for _, tc := range tcs {

		t.Run(tc.description, func(t *testing.T) {

			got := tc.clause.Match(tc.model)

			if tc.expected != got {
				t.Fatalf("expected: %#v, got: %#v", tc.expected, got)
			}
		})
	}
}
