package week04

import (
	"fmt"
)

type Yoga struct {
	Groupname string
	Members   int
	Theme     string
}

func (y Yoga) Name() string {

	fmt.Println("Group Name : ", y.Groupname)

	return y.Groupname
}

func (y Yoga) Perform(v Venue) error {

	fmt.Printf("%s performance started \n", y.Groupname)
	fmt.Printf("%s is performing %s theme by %d members \n", y.Groupname, y.Theme, y.Members)
	fmt.Printf("%s performance is successfully completed \n", y.Groupname)

	return nil

}

func (y Yoga) Teardown(v Venue) error {

	fmt.Println("Setup started by ", y.Groupname)
	fmt.Println("Setup completed by ", y.Groupname)

	return nil
}
