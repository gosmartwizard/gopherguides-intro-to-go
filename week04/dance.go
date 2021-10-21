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

	fmt.Printf("%s is performing %s theme by %d members for %d audience\n", d.Groupname, d.Theme, d.Members, v.Audience)
	return nil
}

func (d Dance) Setup(v Venue) error {

	fmt.Println("Setup completed by group : ", d.Groupname)
	return nil
}

func (d Dance) Teardown(v Venue) error {
	fmt.Println("TearDown completed by group : ", d.Groupname)
	return nil
}
