package week06

import (
	"testing"
)

/* func Test_Product_BuiltBy(t *testing.T) {
	t.Parallel()

	p := Product{
		Quantity: 2,
		builtBy:  1,
	}

	e := p.BuiltBy()

	exp := Employee(1)

	if exp != e {
		t.Fatalf("expected : %#v, got : %#v", exp, e)
	}
} */

func Test_Product_Build(t *testing.T) {
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
			product:     Product{Quantity: 2},
			expected:    ErrInvalidEmployee(0),
		},
		{
			description: "Negative_Employee_Count",
			employee:    Employee(-9),
			product:     Product{Quantity: 2},
			expected:    ErrInvalidEmployee(-9),
		},
		{
			description: "Zero_Product_Quantity",
			employee:    Employee(1),
			product:     Product{Quantity: 0},
			expected:    ErrInvalidQuantity(0),
		},
		{
			description: "Negative_Product_Quantity",
			employee:    Employee(1),
			product:     Product{Quantity: -9},
			expected:    ErrInvalidQuantity(-9),
		},
		{
			description: "Positive_Employee_Count_Product_Quantity",
			employee:    Employee(1),
			product:     Product{Quantity: 2},
			expected:    nil,
		},
	}

	for _, tc := range tcs {

		t.Run(tc.description, func(t *testing.T) {

			err := tc.product.Build(tc.employee)

			if err != nil {
				if tc.expected.Error() != err.Error() {
					t.Fatalf("expected : %#v, got : %#v", tc.expected.Error(), err.Error())
				}
			}

		})
	}
}

func Test_Product_IsValid(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		description string
		product     Product
		expected    error
	}{
		{
			description: "Zero_Product_Quantity",
			product:     Product{Quantity: 0},
			expected:    ErrInvalidQuantity(0),
		},
		{
			description: "Negative_Product_Quantity",
			product:     Product{Quantity: -9},
			expected:    ErrInvalidQuantity(-9),
		},
		{
			description: "Positive_Product_Quantity",
			product:     Product{Quantity: 2},
			expected:    nil,
		},
	}

	for _, tc := range tcs {

		t.Run(tc.description, func(t *testing.T) {

			err := tc.product.IsValid()

			if err != nil {
				if tc.expected.Error() != err.Error() {
					t.Fatalf("expected : %#v, got : %#v", tc.expected.Error(), err)
				}
			}
		})
	}
}

func Test_Product_IsBuilt(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		description string
		product     Product
		expected    error
	}{
		{
			description: "Zero_Product_Quantity",
			product:     Product{Quantity: 0},
			expected:    ErrInvalidQuantity(0),
		},
		{
			description: "Negative_Product_Quantity",
			product:     Product{Quantity: -9},
			expected:    ErrInvalidQuantity(-9),
		},
		{
			description: "Positive_Product_Quantity",
			product:     Product{Quantity: 2},
			expected:    ErrProductNotBuilt("product is not built: {2 0}"),
		},
	}

	for _, tc := range tcs {

		t.Run(tc.description, func(t *testing.T) {

			err := tc.product.IsBuilt()

			if err != nil {
				if tc.expected.Error() != err.Error() {
					t.Fatalf("expected : %#v, got : %#v", tc.expected.Error(), err)
				}
			}
		})
	}
}
