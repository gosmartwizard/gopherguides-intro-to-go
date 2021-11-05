package week06

import (
	"fmt"
	"testing"
)

func Test_Completed_Product_1(t *testing.T) {
	t.Parallel()

	cp := CompletedProduct{
		Employee: Employee(-8),
	}

	err := cp.IsValid()

	exp := ErrInvalidEmployee(-8)

	if exp.Error() != err.Error() {
		t.Fatalf("expected : %#v, got : %#v", exp, err)
	}
}

func Test_Completed_Product_2(t *testing.T) {
	t.Parallel()

	cp := CompletedProduct{
		Employee: Employee(1),
		Product: Product{
			Quantity: -9,
			builtBy:  Employee(1),
		},
	}

	err := cp.IsValid()

	exp := ErrInvalidQuantity(-9)

	if exp.Error() != err.Error() {
		t.Fatalf("expected : %#v, got : %#v", exp, err)
	}

}

func Test_Completed_Product_3(t *testing.T) {
	t.Parallel()

	cp := CompletedProduct{
		Employee: Employee(1),
		Product: Product{
			Quantity: 2,
			builtBy:  0,
		},
	}

	err := cp.IsValid()

	exp := ErrProductNotBuilt(fmt.Sprintf("product is not built: %v", cp.Product))

	if exp.Error() != err.Error() {
		t.Fatalf("expected : %#v, got : %#v ", exp.Error(), err.Error())
	}
}
