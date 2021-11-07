package week06

import (
	"testing"
)

func Test_Employee_InValid(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		description string
		employee    Employee
		expected    error
	}{
		{
			description: "Zero_Employee_Number",
			employee:    Employee(0),
			expected:    ErrInvalidEmployee(0),
		},
		{
			description: "Negative_Employee_Number",
			employee:    Employee(-9),
			expected:    ErrInvalidEmployee(-9),
		},
		{
			description: "Positive_Employee_Number",
			employee:    Employee(3),
			expected:    nil,
		},
	}

	for _, tc := range tcs {

		t.Run(tc.description, func(t *testing.T) {

			err := tc.employee.IsValid()

			if err != nil {
				if tc.expected.Error() != err.Error() {
					t.Fatalf("expected : %#v, got : %#v", tc.expected.Error(), err.Error())
				}
			}
		})
	}
}

func Test_Employee_Work(t *testing.T) {
	t.Parallel()

	m := NewManager()

	m.Start(1)

	p := &Product{Quantity: 2}

	go func() {
		m.Assign(p)
	}()

	select {
	case cp := <-m.Completed():

		if cp.Employee != p.BuiltBy() {
			t.Fatalf("expected : %#v, got : %#v", p.builtBy, cp)
		}

	case <-m.Done():
	}

	m.Stop()
}
