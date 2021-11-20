package week08

import (
	"testing"
)

func Test_Product_BuiltBy(t *testing.T) {
	t.Parallel()

	p := &Product{
		Materials: Materials{
			Wood: 2,
			Oil:  3,
		},
		builtBy: 23,
	}

	e := p.BuiltBy()

	exp := Employee(23)

	if exp != e {
		t.Fatalf("expected : %#v, got : %#v", exp, e)
	}
}

func Test_Product_Build(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		description string
		employee    Employee
		product     Product
		expected    error
	}{
		{
			description: "Zero_Employee_Number",
			employee:    Employee(0),
			product:     Product{Materials: Materials{Wood: 2, Oil: 3}},
			expected:    ErrInvalidEmployee(0),
		},
		{
			description: "Negative_Employee_Number",
			employee:    Employee(-9),
			product:     Product{Materials: Materials{Wood: 2, Oil: 3}},
			expected:    ErrInvalidEmployee(-9),
		},
		{
			description: "Error_Invalid_Materials_",
			employee:    Employee(1),
			product:     Product{Materials: Materials{}},
			expected:    ErrInvalidMaterials(0),
		},
	}

	for _, tc := range tcs {

		t.Run(tc.description, func(t *testing.T) {

			err := tc.product.Build(tc.employee, &Warehouse{})

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
			description: "Error_Invalid_Materials",
			product:     Product{Materials: Materials{}},
			expected:    ErrInvalidMaterials(0),
		},
		{
			description: "Golden_Path",
			product:     Product{Materials: Materials{Wood: 2, Oil: 3}},
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
			description: "Error_Invalid_Materials",
			product:     Product{Materials: Materials{}},
			expected:    ErrInvalidMaterials(0),
		},
		{
			description: "Error_Product_Not_Built",
			product:     Product{Materials: Materials{Wood: 2, Oil: 3}, builtBy: 0},
			expected:    ErrProductNotBuilt("product is not built: [{oil:3x}, {wood:2x}]"),
		},
		{
			description: "Golden_Path",
			product:     Product{Materials: Materials{Wood: 2, Oil: 3}, builtBy: 1},
			expected:    nil,
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
