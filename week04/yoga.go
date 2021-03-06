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

	if v.Audience == 0 {
		return fmt.Errorf("there are %d audience to perform", v.Audience)
	}

	fmt.Printf("%s performance started \n", y.Groupname)
	fmt.Printf("%s is performing %s theme by %d members \n", y.Groupname, y.Theme, y.Members)
	fmt.Printf("%s performance is successfully completed \n", y.Groupname)

	fmt.Fprintf(v.Log, "%s has performed for %d people. \n", y.Groupname, v.Audience)

	return nil

}

func (y Yoga) Teardown(v Venue) error {

	fmt.Println("Setup started by ", y.Groupname)
	fmt.Println("Setup completed by ", y.Groupname)

	fmt.Fprintf(v.Log, "%s has completed teardown. \n", y.Groupname)

	return nil
}
