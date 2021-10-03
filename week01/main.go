package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	varString := "GO"
	varInteger := 42
	varBoolean := true

	fmt.Printf("Printing, %T (%#v)! \n", varString, varString)
	fmt.Printf("Printing, %T (%#v)! \n", varInteger, varInteger)
	fmt.Printf("Printing, %T (%#v)! \n", varBoolean, varBoolean)
}
