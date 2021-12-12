package main

import (
	"flag"
)

func HandleClear(args []string) error {

	var BackupFile string

	clearCmd := flag.NewFlagSet("clear", flag.ExitOnError)

	clearCmd.StringVar(&BackupFile, "f", "/tmp/NewsServiceBackup.json", "location of the backupfile")

	clearCmd.Parse(args)

	err := clear(BackupFile)

	if err != nil {
		return err
	}

	return nil
}

func clear(backupFile string) error {

	ns := NewNewService()

	err := ns.Start()

	if err != nil {
		return err
	}

	err = ns.Clear(backupFile)

	if err != nil {
		return err
	}

	return nil
}
