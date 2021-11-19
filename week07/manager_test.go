package week07

import (
	"context"
	"os/signal"
	"runtime"
	"syscall"
	"testing"
	"time"
)

func Test_Run_GoldenPath(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	count := 1

	cps, err := Run(ctx, count, &Product{Quantity: 25000})

	if err != nil {
		t.Fatalf(err.Error())
	}

	s := len(cps)

	if count != s {
		t.Fatalf("expected : %#v, got : %#v", count, s)
	}
}

func Test_Run__InvalidQuantity(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	_, err := Run(ctx, 1, &Product{Quantity: -2})

	if err == nil {
		t.Fatalf("expected : %#v, got : nil", err.Error())
	}
}

func Test_Run_Timeout(t *testing.T) {
	t.Parallel()

	rootCtx := context.Background()

	ctx, cancel := context.WithTimeout(rootCtx, 10*time.Second)

	defer cancel()

	go Run(ctx, 1, &Product{Quantity: 25000})

	select {
	case <-rootCtx.Done():
	case <-ctx.Done():
	}

	exp := context.DeadlineExceeded.Error()

	if exp != ctx.Err().Error() {
		t.Fatalf("expected : %#v, got : %#v", exp, ctx.Err().Error())
	}
}

func Test_Run_NotifySignal(t *testing.T) {
	t.Parallel()

	if runtime.GOOS == "windows" {
		t.Skip()
	}

	const TEST_SIGNAL = syscall.SIGUSR2

	rootCtx := context.Background()

	sigCtx, cancel := signal.NotifyContext(rootCtx, TEST_SIGNAL)

	defer cancel()

	go Run(sigCtx, 1, &Product{Quantity: 25000})

	go func() {
		time.Sleep(time.Second * 10)
		syscall.Kill(syscall.Getpid(), TEST_SIGNAL)
	}()

	select {
	case <-rootCtx.Done():
	case <-sigCtx.Done():
	}

	cancel()

	exp := context.Canceled.Error()

	if exp != sigCtx.Err().Error() {
		t.Fatalf("expected : %#v, got : %#v", exp, sigCtx.Err().Error())
	}
}

func Test_Manager_Start(t *testing.T) {
	t.Parallel()

	rootCtx := context.Background()

	ctx, cancel := context.WithTimeout(rootCtx, 5*time.Second)

	defer cancel()

	m := NewManager()

	defer m.Stop()

	m.Start(ctx, 5)

	select {
	case <-rootCtx.Done():
	case <-ctx.Done():
	}

	exp := context.DeadlineExceeded.Error()

	if exp != ctx.Err().Error() {
		t.Fatalf("expected : %#v, got : %#v", exp, ctx.Err().Error())
	}
}

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

			m := NewManager()

			defer m.Stop()

			err := m.Start(ctx, tc.count)

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

func Test_Manager_Assign(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		description string
		product     Product
		expected    error
	}{
		{
			description: "ErrInvalidQuantity_Zero_Quantity",
			product:     Product{Quantity: 0},
			expected:    ErrInvalidQuantity(0),
		},
		{
			description: "ErrInvalidQuantity_Negative_Quantity",
			product:     Product{Quantity: -9},
			expected:    ErrInvalidQuantity(-9),
		},
	}

	for _, tc := range tcs {

		t.Run(tc.description, func(t *testing.T) {

			m := NewManager()

			defer m.Stop()

			err := m.Assign(&tc.product)

			exp := ErrInvalidQuantity(tc.product.Quantity)

			if exp.Error() != err.Error() {
				t.Fatalf("expected : %#v, got : %#v", exp, err)
			}
		})
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

			defer m.Stop()

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

func Test_Manager_Start_Manager_Stop(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	m := NewManager()

	m.Start(ctx, 1)

	m.Stop()

	_, ok := <-m.Jobs()

	if ok {
		t.Fatalf("expected : false, got : %#v", ok)
	}
}

func Test_Manager_Complete_ErrManagerStopped(t *testing.T) {
	t.Parallel()

	m := NewManager()

	e := Employee(3)

	p := &Product{Quantity: 2}

	p.Build(e)

	m.Stop()

	err := m.Complete(e, p)

	exp := ErrManagerStopped{}

	if exp.Error() != err.Error() {
		t.Fatalf("expected : %#v, got : %#v", exp.Error(), err.Error())
	}
}
