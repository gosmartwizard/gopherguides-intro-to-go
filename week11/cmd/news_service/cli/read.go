package cli

import (
	"flag"
)

//var _ IOCommander = &ListCmd{}

type ReadCmd struct {
	//IO
	Name string
	//Book *notes.Book // book to list

	// flags:
	BackupFile string // location of the book database file
	JSON       bool   // output in JSON format
	Output     string // output file
	//Verbose bool   // verbose output

	flags *flag.FlagSet
}

// snippet: type

// snippet: flags
func (cmd *ReadCmd) Flags() *flag.FlagSet {

	if len(cmd.BackupFile) == 0 {
		cmd.BackupFile = "/tmp/news.json"
	}

	if cmd.flags != nil {
		return cmd.flags
	}

	// create a new flag set
	cmd.flags = flag.NewFlagSet(cmd.Name, flag.ContinueOnError)

	// add the flags to the flag set
	cmd.flags.BoolVar(&cmd.JSON, "j", cmd.JSON, "output in json format")
	cmd.flags.StringVar(&cmd.BackupFile, "f", cmd.BackupFile, "location of the backupfile")
	cmd.flags.StringVar(&cmd.Output, "o", cmd.Output, "output results to a file")

	return cmd.flags
}

// snippet: flags

// snippet: main
/*
func (cmd *ListCmd) Main(ctx context.Context, pwd string, args []string) error {
	if err := cmd.init(pwd, args); err != nil {
		return err
	}

	return cmd.print(cmd.Book)
} */

// snippet: main

// snippet: setio
/* func (cmd *ListCmd) SetIO(oi IO) {
	cmd.IO = oi
} */

// snippet: setio

// snippet: init
/* func (cmd *ReadCmd) init(pwd string, args []string) error {
	// set a default location for the database file
	// if the user hasn't specified one
	if cmd.BackupFile == "" {
		cmd.BackupFile = "/tmp/news.json"
	}

	// parse the flags with the given arguments
	if err := cmd.Flags().Parse(args); err != nil {
		return err
	}

	if cmd.Book == nil {
		cmd.Book = &notes.Book{}
		// open the database file
		err := OpenBook(filepath.Join(pwd, cmd.DB), cmd.Book)
		if err != nil {
			return err
		}

	}

	return nil
}
*/
// snippet: init

// snippet: print
/* func (cmd *ReadCmd) print(book *notes.Book) error {
	notes, err := book.Select()
	if err != nil {
		return err
	}

	// if the user has specified an output file
	// write the results to the file
	if len(cmd.Output) > 0 {
		os.MkdirAll(filepath.Dir(cmd.Output), 0755)
		f, err := os.Create(cmd.Output)
		if err != nil {
			return err
		}
		defer f.Close()

		// set the Stdout to the file
		cmd.Out = f
	}

	// if the user has specified JSON output
	// marshal the results to JSON
	// and print to Stdout (or file)
	if cmd.JSON {
		return json.NewEncoder(cmd.Stdout()).Encode(notes)
	}

	// otherwise, print the results to Stdout
	for _, note := range notes {
		fmt.Fprintf(cmd.Stdout(), "%d\t%s\n", note.ID(), note.Short(50))
	}

	return nil
}
*/
// snippet: print
