package week08

import (
	"context"
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

	ctx := context.Background()

	m := &Manager{}

	defer m.Stop()

	ctx, err := m.Start(ctx, 1)

	if err != nil {
		t.Fatalf("expected : nil, got : %#v", err)
	}

	p := &Product{
		Materials: Materials{
			Wood: 2,
			Oil:  3,
		},
	}

	go func() {
		m.Assign(p)
	}()

	select {
	case cp := <-m.Completed():

		if cp.Employee != p.BuiltBy() {
			t.Fatalf("expected : %#v, got : %#v", p.builtBy, cp)
		}

	case <-ctx.Done():
	}
}

func Test_Employee_Work_Errors(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	m := &Manager{}

	defer m.Stop()

	ctx, err := m.Start(ctx, 1)

	if err != nil {
		t.Fatalf("expected : nil, got : %#v", err)
	}

	p := &Product{
		Materials: Materials{},
	}

	go func() {
		m.Assign(p)
	}()

	select {
	case err := <-m.Errors():

		if err == nil {
			t.Fatalf("expected : %#v, got : nil", err.Error())
		}
	case <-ctx.Done():
	}
}

func Test_Employee_Work_Close_Jobs(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	m := &Manager{}

	defer m.Stop()

	_, err := m.Start(ctx, 1)

	if err != nil {
		t.Fatalf("expected : nil, got : %#v", err)
	}

	if m.Jobs() != nil {
		close(m.Jobs())
	}
	_, ok := <-m.Jobs()

	if ok {
		t.Fatal("expected : false, got : true")
	}
}

func Test_Employee_Work_ErrInvalidMaterials(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	m := &Manager{}

	defer m.Stop()

	ctx, err := m.Start(ctx, 1)

	if err != nil {
		t.Fatalf("expected : nil, got : %#v", err)
	}
	e := Employee(0)

	go e.work(ctx, m)

	p := &Product{
		Materials: Materials{},
	}

	m.Jobs() <- p

	select {
	case err := <-m.Errors():

		if err == nil {
			t.Fatalf("expected : %#v, got : nil", err.Error())
		}
	case <-ctx.Done():
	}
}
