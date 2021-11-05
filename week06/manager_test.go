package week06

import (
	"testing"
)

// ATTENTION: YOU ARE NOT ALLOWED TO SUBMIT THIS
// TEST AS PART OF YOUR ASSIGNMENT!!
//
// This test is meant to demonstrate how to use the
// application.
//
// YOU MUST DELETE THIS TEST !!BEFORE!! YOU SUBMIT!!

func Test_Manager_1(t *testing.T) {
	t.Parallel()

	m := NewManager()

	err := m.Start(0)

	exp := ErrInvalidEmployeeCount(0)

	if exp.Error() != err.Error() {
		t.Fatalf("expected : %#v, got : %#v", exp, err)
	}
}

func Test_Manager_2(t *testing.T) {
	t.Parallel()

	m := NewManager()

	err := m.Start(-9)

	exp := ErrInvalidEmployeeCount(-9)

	if exp.Error() != err.Error() {
		t.Fatalf("expected : %#v, got : %#v", exp, err)
	}
}

func Test_Manager_Demonstration(t *testing.T) {
	t.Parallel()

	// ALWAYS use the NewManager function to
	// create a new Manager. This ensures that
	// the Manager is always in a consistent state.
	// Failure to do so will result in a panic.
	m := NewManager()

	// start the manager with 5 employees
	err := m.Start(5)
	if err != nil {
		t.Fatal(err)
	}

	// buildCount is the amount of products we
	// want to build
	buildCount := 5

	// launch buildCount number of goroutines
	// each goroutine will push a new product
	// into the system.
	for i := 0; i < buildCount; i++ {

		// for each new product we will launch
		// a new goroutine to push it into the
		// system.
		go func() {

			// push a new product, with a quantity of 2,
			// in the system.
			// this will block until the manager is able to
			// to assign the product to an employee.
			err := m.Assign(&Product{Quantity: 2})
			if err != nil {
				m.Errors() <- err
			}

		}() // go

	} // for

	// actual completed products we receive.
	var act []CompletedProduct

	// launch a goroutine that will receive the completed
	// products and add them to the act slice.
	go func() {

		// keep receiving completed products until
		// the completed channel is closed.
		for cp := range m.Completed() {

			// add the completed product to the act slice
			act = append(act, cp)

			// if we have received all the products we
			// want to receive, we can stop the manager.
			if len(act) == buildCount {
				m.Stop()
			}

		} // for

	}() // go

	select {
	case err := <-m.Errors():
		// if the manager received an error, we will
		// mark the test as failed.
		t.Fatal(err)
	case <-m.Done():
		// the manager has been stopped, we can
		// continue and check the results.
	}

	if len(act) != buildCount {
		t.Fatalf("expected %d, got %d", buildCount, len(act))
	}

}
