package main

import (
	"fmt"
	"os"
)

func GetNewserviceStats() error {

	ns := NewNewService()

	err := ns.Start()

	if err != nil {
		return err
	}

	stats, err := ns.GetNewsServiceStats("/tmp/NewsServiceBackup.json")

	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stdout, stats)

	return nil
}
