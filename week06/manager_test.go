package week06

import (
	"testing"
)

func Test_Manager_Start(t *testing.T) {
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

			m := NewManager()

			err := m.Start(tc.count)

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

	m := NewManager()

	p := &Product{
		Quantity: 2,
	}

	m.Stop()

	err := m.Assign(p)

	exp := ErrManagerStopped{}

	if exp.Error() != err.Error() {
		t.Fatalf("expected : %#v, got : %#v", exp, err)
	}
}

func Test_Manager_Assign_InvalidQuantity(t *testing.T) {
	t.Parallel()

	m := NewManager()

	p := &Product{
		Quantity: 0,
	}

	err := m.Assign(p)

	exp := ErrInvalidQuantity(p.Quantity)

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
	}

	for _, tc := range tcs {

		t.Run(tc.description, func(t *testing.T) {

			m := NewManager()

			err := m.Complete(tc.employee, &tc.product)

			if tc.expected.Error() != err.Error() {
				t.Fatalf("expected : %#v, got : %#v", tc.expected.Error(), err)
			}
		})
	}
}

func Test_Manager_Stop(t *testing.T) {
	t.Parallel()

	m := NewManager()

	m.Stop()

	m.Stop()

	_, ok := <-m.Jobs()

	if ok {
		t.Fatalf("expected : false, got : %#v", ok)
	}
}

func Test_Manager_Assign(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		description string
		employee    Employee
		product     Product
		expected    error
	}{
		{
			description: "ErrInvalidEmployee_Zero_Employee_Number",
			employee:    Employee(0),
			product:     Product{Quantity: 2},
			expected:    ErrInvalidEmployee(0),
		},
		{
			description: "ErrInvalidEmployee_Negative_Employee_Number",
			employee:    Employee(-9),
			product:     Product{Quantity: 2},
			expected:    ErrInvalidEmployee(-9),
		},
	}

	for _, tc := range tcs {

		t.Run(tc.description, func(t *testing.T) {

			m := NewManager()

			e := Employee(tc.employee)

			go e.work(m)

			go func() {
				m.Assign(&tc.product)
			}()

			exp := ErrInvalidEmployee(tc.employee)

			select {
			case err := <-m.Errors():
				if exp.Error() != err.Error() {
					t.Fatalf("expected : %#v, got : %#v", exp, err)
				}
			case <-m.Done():
			}
		})
	}
}

func Test_Manager_Start_Manager_Stop(t *testing.T) {
	t.Parallel()

	m := NewManager()

	m.Start(1)

	m.Stop()

	_, ok := <-m.Jobs()

	if ok {
		t.Fatalf("expected : false, got : %#v", ok)
	}
}
