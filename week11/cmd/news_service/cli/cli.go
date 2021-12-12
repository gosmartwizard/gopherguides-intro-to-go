package cli

import (
	"fmt"
	"io"
	"os"
	//"github.com/gosmartwizard/gopherguides-intro-to-go/week11/pkg"
)

type App struct {
}

func (app *App) Main(args []string) error {
	if app == nil {
		return fmt.Errorf("app is nil")
	}

	if len(args) == 0 {
		return GetNewserviceStats()
	}

	if args[0] == "-h" {
		return app.Usage(os.Stdout)
	}

	switch args[0] {
	case "stream":
		return HandleStream(args[1:])
	case "read":
		return HandleRead(args[1:])
	case "clear":
		return HandleClear(args[1:])
	default:
		return fmt.Errorf("%v command not supported", args[0])
	}
}

func (app *App) Usage(w io.Writer) error {
	fmt.Fprintln(w, "Usage: news_service <command> [options] [<args>...]")

	fmt.Fprintln(w, "---------------")

	fmt.Fprintln(w, "news_service stream -f /tmp/NewsServiceBackup.json -j -o /tmp/stream.json sports")

	fmt.Fprintln(w, "news_service read -o /tmp/articles.json -j -f /tmp/NewsServiceBackup.json 1 2 3")

	fmt.Fprintln(w, "news_service clear -f /tmp/NewsServiceBackup.json")

	return nil
}
