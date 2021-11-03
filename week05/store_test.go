package week05

import (
	"errors"
	"testing"
)

func assertModel(t *testing.T, act Model, exp Model) {
	for i, m := range exp {
		if m != act[i] {
			t.Fatalf("expected : %#v, got : %#v", m, act[i])
		}
	}
}

func Test_Store_Insert(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		description string
		models      Models
		tn          string
	}{
		{
			description: "Golden_Path",
			models: Models{
				Model{"iPhone": "Iphone5"},
				Model{"iPhone": "Iphone6"},
			},
			tn: "Mobiles",
		},
		{
			description: "No_Models",
			models:      Models{},
			tn:          "Mobiles",
		},
	}

	for _, tc := range tcs {

		t.Run(tc.description, func(t *testing.T) {

			s := &Store{}

			for _, m := range tc.models {
				s.Insert(tc.tn, m)
			}

			for k, v1 := range s.data {
				if k != tc.tn {
					t.Fatalf("expected: %#v,and got: %#v", "Mobiles", k)
				}

				for i, v2 := range tc.models {
					assertModel(t, v1[i], v2)
				}
			}
		})
	}
}

func Test_Store_All(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		description string
		models      Models
		itn         string
		atn         string
		expected    Models
		err         error
	}{
		{
			description: "Golden_Path",
			models: Models{
				Model{"iPhone": "Iphone5"},
				Model{"iPhone": "Iphone6"},
			},
			itn: "Mobiles",
			atn: "Mobiles",
			expected: Models{
				Model{"iPhone": "Iphone5"},
				Model{"iPhone": "Iphone6"},
			},
			err: nil,
		},
		{
			description: "Error_Table_Not_Found",
			models: Models{
				Model{"iPhone": "Iphone5"},
				Model{"iPhone": "Iphone6"},
			},
			itn: "Mobiles",
			atn: "Laptops",
			expected: Models{
				Model{"iPhone": "Iphone5"},
				Model{"iPhone": "Iphone6"},
			},
			err: ErrTableNotFound{table: "Laptops"},
		},
	}

	for _, tc := range tcs {

		t.Run(tc.description, func(t *testing.T) {

			s := &Store{}

			for _, m := range tc.models {
				s.Insert(tc.itn, m)
			}

			mods, err := s.All(tc.atn)

			if err != nil {
				b := errors.Is(err, tc.err)
				if !b {
					t.Fatalf("expected : %#v, got : %#v", tc.err, err)
				}

				if ok := IsErrTableNotFound(err); !ok {
					t.Fatalf("expected: %#v,and got: %#v", tc.err, err)
				}

				if tc.err.Error() != err.Error() {
					t.Fatalf("expected: %#v,and got: %#v", tc.err, err)
				}

				return
			}

			for i, m := range mods {
				assertModel(t, m, tc.expected[i])
			}
		})
	}
}

func Test_Store_Len(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		description string
		models      Models
		itn         string
		ltn         string
		expected    int
		err         error
	}{
		{
			description: "Golden_Path",
			models: Models{
				Model{"iPhone": "Iphone5"},
				Model{"iPhone": "Iphone6"},
			},
			itn:      "Mobiles",
			ltn:      "Mobiles",
			expected: 2,
			err:      nil,
		},
		{
			description: "Error_Table_Not_Found",
			models: Models{
				Model{"iPhone": "Iphone5"},
				Model{"iPhone": "Iphone6"},
			},
			itn:      "Mobiles",
			ltn:      "Laptops",
			expected: 0,
			err:      ErrTableNotFound{table: "Laptops"},
		},
	}

	for _, tc := range tcs {

		t.Run(tc.description, func(t *testing.T) {

			s := Store{}

			for _, m := range tc.models {
				s.Insert(tc.itn, m)
			}

			len, err := s.Len(tc.ltn)

			if err != nil {
				b := errors.Is(err, tc.err)
				if !b {
					t.Fatalf("expected : %#v, got : %#v", tc.err, err)
				}
				return
			}

			if tc.expected != len {
				t.Fatalf("expected: %#v, got: %#v", tc.expected, len)
			}
		})
	}
}

func Test_Store_Select(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		description string
		model       Model
		clause      Clauses
		itn         string
		stn         string
		models      Models
		err         error
	}{
		{
			description: "Golden_Path",
			model:       Model{"iPhone": "Iphone5"},
			clause:      Clauses{"iPhone": "Iphone5"},
			itn:         "Mobiles",
			stn:         "Mobiles",
			models:      Models{Model{"iPhone": "Iphone5"}},
			err:         nil,
		},
		{
			description: "Error_Table_Not_Found",
			model:       Model{"iPhone": "Iphone5"},
			clause:      Clauses{"iPhone": "Iphone5"},
			itn:         "Mobiles",
			stn:         "Laptops",
			models:      Models{Model{"iPhone": "Iphone5"}},
			err:         ErrTableNotFound{table: "Laptops"},
		},
		{
			description: "Empty_Clauses",
			model:       Model{"iPhone": "Iphone5"},
			clause:      Clauses{},
			itn:         "Mobiles",
			stn:         "Mobiles",
			models:      Models{Model{"iPhone": "Iphone5"}},
			err:         nil,
		},
		{
			description: "Error_No_Rows",
			model:       Model{"iPhone": "Iphone5"},
			clause:      Clauses{"iPhone": "Iphone6"},
			itn:         "Mobiles",
			stn:         "Mobiles",
			models:      Models{Model{"iPhone": "Iphone5"}},
			err:         &errNoRows{},
		},
	}

	for _, tc := range tcs {

		t.Run(tc.description, func(t *testing.T) {

			s := &Store{}

			s.Insert(tc.itn, tc.model)

			mods, err := s.Select(tc.stn, tc.clause)

			if tc.err != nil {
				b := errors.Is(err, tc.err)
				if !b {
					t.Fatalf("expected : %#v, got : %#v", tc.err, err)
				}
				return
			}

			for i, m := range mods {
				assertModel(t, m, tc.models[i])
			}
		})
	}
}
