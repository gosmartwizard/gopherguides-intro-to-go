package week06

import (
	"testing"
)

func Test_Product_BuiltBy(t *testing.T) {
	t.Parallel()

	e := Employee(0)

	m := NewManager()

	m.Start(1)

	p := &Product{Quantity: 1}

	go m.Assign(p)

	go func() {
		for cp := range m.Completed() {
			e = cp.Employee
			err := cp.IsValid()
			if err == nil {
				m.Stop()
			}
		}
	}()

	select {
	case err := <-m.Errors():
		t.Fatal(err)
	case <-m.Done():
	}

	exp := p.BuiltBy()

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
			product:     Product{Quantity: 2},
			expected:    ErrInvalidEmployee(0),
		},
		{
			description: "Negative_Employee_Number",
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
