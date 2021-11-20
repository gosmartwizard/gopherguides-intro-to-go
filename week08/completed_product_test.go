package week08

import "testing"

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
			description: "Zero_Product_Materials",
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
		{
			description: "GoldenPath",
			employee:    Employee(1),
			product:     Product{Materials: Materials{Wood: 2, Oil: 3}, builtBy: 1},
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
