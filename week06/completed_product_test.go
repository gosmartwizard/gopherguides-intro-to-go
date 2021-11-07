package week06

import (
	"testing"
)

func Test_CompletedProduct_IsValid(t *testing.T) {
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
			product:     Product{},
			expected:    ErrInvalidEmployee(0),
		},
		{
			description: "Negative_Employee_Number",
			employee:    Employee(-9),
			product:     Product{},
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
			description: "Error_Product_Not_Built",
			employee:    Employee(1),
			product:     Product{Quantity: 2, builtBy: 0},
			expected:    ErrProductNotBuilt("product is not built: {2 0}"),
		},
		{
			description: "Positive_Employee_Count_Product_Quantity",
			employee:    Employee(1),
			product:     Product{Quantity: 2, builtBy: 1},
			expected:    nil,
		},
	}

	for _, tc := range tcs {

		t.Run(tc.description, func(t *testing.T) {

			cp := CompletedProduct{
				Employee: tc.employee,
				Product:  tc.product,
			}

			err := cp.IsValid()

			if err != nil {
				if tc.expected.Error() != err.Error() {
					t.Fatalf("expected : %#v, got : %#v", tc.expected.Error(), err.Error())
				}
			}
		})
	}
}
