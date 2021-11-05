package week06

import (
	"testing"
	"time"
)

func Test_Employee_1(t *testing.T) {
	t.Parallel()

	e := Employee(2)

	err := e.IsValid()

	if err != nil {
		t.Fatalf("expected : nil, got : %#v", err)
	}
}

func Test_Employee_2(t *testing.T) {
	t.Parallel()

	e := Employee(0)

	err := e.IsValid()

	exp := ErrInvalidEmployee(0)

	if exp.Error() != err.Error() {
		t.Fatalf("expected : %#v, got : %#v", exp, err)
	}
}

func Test_Employee_3(t *testing.T) {
	t.Parallel()

	e := Employee(-9)

	err := e.IsValid()

	exp := ErrInvalidEmployee(-9)

	if exp.Error() != err.Error() {
		t.Fatalf("expected : %#v, got : %#v", exp, err)
	}
}

func Test_Employee_4(t *testing.T) {
	t.Parallel()

	m := NewManager()

	e := Employee(-1)

	go e.work(m)

	go func() {
		m.Assign(&Product{Quantity: 2})
	}()

	exp := ErrInvalidEmployee(-1)

	select {
	case err := <-m.Errors():
		if exp.Error() != err.Error() {
			t.Fatalf("expected : %#v, got : %#v", exp, err)
		}
	case <-m.Done():
	}
}

func Test_Employee_5(t *testing.T) {
	t.Parallel()

	m := NewManager()

	e := Employee(0)

	go e.work(m)

	go func() {
		m.Assign(&Product{Quantity: 2})
	}()

	exp := ErrInvalidEmployee(0)

	select {
	case err := <-m.Errors():
		if exp.Error() != err.Error() {
			t.Fatalf("expected : %#v, got : %#v", exp, err)
		}
	case <-m.Done():
	}
}

func Test_Employee_6(t *testing.T) {
	t.Parallel()

	m := NewManager()

	e := Employee(1)

	go e.work(m)

	go func(m *Manager) {
		m.Assign(&Product{Quantity: 2})
	}(m)

	var act []CompletedProduct

	go func() {
		for cp := range m.Completed() {
			act = append(act, cp)
			if len(act) == 1 {
				close(m.Jobs())
				time.Sleep(time.Millisecond * 1000)
				close(m.Errors())
			}
		}
	}()
}
