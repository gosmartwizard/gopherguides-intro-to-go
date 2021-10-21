package week04

import (
	"fmt"
)

type Music struct {
	Groupname string
	Members   int
	Theme     string
}

func (m Music) Name() string {

	fmt.Println("Group Name : ", m.Groupname)
	return m.Groupname
}

func (m Music) Perform(v Venue) error {

	fmt.Printf("%s is performing %s theme by %d members for %d audience\n", m.Groupname, m.Theme, m.Members, v.Audience)
	return nil
}

func (m Music) Setup(v Venue) error {

	fmt.Println("Setup completed by group : ", m.Groupname)
	return nil
}

func (m Music) Teardown(v Venue) error {
	fmt.Println("TearDown completed by group : ", m.Groupname)
	return nil
}
