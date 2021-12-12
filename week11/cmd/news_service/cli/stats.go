package cli

import (
	"fmt"
	"os"
	"week11/newsservice"
)

func GetNewserviceStats() error {

	ns := newsservice.NewNewService()

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
