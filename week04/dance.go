package week04

import (
	"fmt"
)

type Dance struct {
	Groupname string
	Members   int
	Theme     string
}

func (d Dance) Name() string {

	fmt.Println("Group Name : ", d.Groupname)

	return d.Groupname
}

func (d Dance) Perform(v Venue) error {

	fmt.Printf("%s performance started \n", d.Groupname)
	fmt.Printf("%s is performing %s theme by %d members \n", d.Groupname, d.Theme, d.Members)
	fmt.Printf("%s performance is successfully completed \n", d.Groupname)

	return nil
}

func (d Dance) Teardown(v Venue) error {

	fmt.Println("Setup started by ", d.Groupname)
	fmt.Println("Setup completed by ", d.Groupname)

	return nil
}
