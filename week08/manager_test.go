package week08

import (
	"context"
	"testing"
)

func Test_Manager_Start_ErrInvalidEmployeeCount(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		description string
		count       int
		expected    error
	}{
		{
			description: "Zero_Employee_Count",
			count:       0,
			expected:    ErrInvalidEmployeeCount(0),
		},
		{
			description: "Negative_Employee_Count",
			count:       -9,
			expected:    ErrInvalidEmployeeCount(-9),
		},
		{
			description: "Positive_Employee_Count",
			count:       3,
			expected:    nil,
		},
	}

	for _, tc := range tcs {

		t.Run(tc.description, func(t *testing.T) {

			ctx := context.Background()

			m := &Manager{}

			_, err := m.Start(ctx, tc.count)

			if err != nil {
				if tc.expected.Error() != err.Error() {
					t.Fatalf("expected : %#v, got : %#v", tc.expected.Error(), err)
				}
			}
		})
	}
}

func Test_Manager_Assign_ManagerStopped(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	m := &Manager{}

	_, err := m.Start(ctx, 1)
	if err != nil {
		t.Fatalf("expected : nil, got : %#v", err)
	}

	p := &Product{}

	m.Stop()

	err = m.Assign(p)

	exp := ErrManagerStopped{}

	if exp.Error() != err.Error() {
		t.Fatalf("expected : %#v, got : %#v", exp, err)
	}
}

func Test_Manager_Complete(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		description string
		employee    Employee
		product     Product
		expected    error
	}{
		{
			description: "Zero_Employee_Count",
			employee:    Employee(0),
			product:     Product{},
			expected:    ErrInvalidEmployee(0),
		},
		{
			description: "Negative_Employee_Count",
			employee:    Employee(-9),
			product:     Product{},
			expected:    ErrInvalidEmployee(-9),
		},
		{
			description: "Error_Invalid_Materials",
			employee:    Employee(1),
			product:     Product{Materials: Materials{}},
			expected:    ErrInvalidMaterials(0),
		},
		{
			description: "Error_Product_Not_Built",
			employee:    Employee(1),
			product:     Product{Materials: Materials{Wood: 2, Oil: 3}, builtBy: 0},
			expected:    ErrProductNotBuilt("product is not built: [{oil:3x}, {wood:2x}]"),
		},
	}

	for _, tc := range tcs {

		t.Run(tc.description, func(t *testing.T) {

			ctx := context.Background()

			m := &Manager{}

			_, err := m.Start(ctx, 1)

			if err != nil {
				t.Fatalf("expected : nil, got : %#v", err)
			}

			defer m.Stop()

			err = m.Complete(tc.employee, &tc.product)

			if tc.expected.Error() != err.Error() {
				t.Fatalf("expected : %#v, got : %#v", tc.expected.Error(), err)
			}
		})
	}
}
