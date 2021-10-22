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

	fmt.Printf("%s performance started \n", m.Groupname)
	fmt.Printf("%s is performing %s theme by %d members \n", m.Groupname, m.Theme, m.Members)
	fmt.Printf("%s performance is successfully completed \n", m.Groupname)

	return nil
}

func (m Music) Setup(v Venue) error {

	fmt.Println("Setup started by ", m.Groupname)
	fmt.Println("Setup completed by ", m.Groupname)

	return nil
}
