package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"testing"
	"time"
)

func generateBackupData(filePath string) error {

	articles := make(map[int]Article)

	article := Article{}

	article.Source = "FileBasedSource"
	article.Category = "Sports"
	article.Description = "Brian Lara"

	articles[1] = article

	article.Source = "FileBasedSource"
	article.Category = "Movies"
	article.Description = "Dear Comrade"

	articles[2] = article

	article.Source = "FileBasedSource"
	article.Category = "Tech"
	article.Description = "Kubernetes"

	articles[3] = article

	fileBytes, err := json.Marshal(articles)

	if err != nil {
		return err
	}

	os.Remove(filePath)

	ioutil.WriteFile(filePath, fileBytes, 0644)

	return nil
}

func Test_Cli_Read(t *testing.T) {
	t.Parallel()

	filePath := "./testdata/cli/newsarticles1.json"

	generateBackupData(filePath)

	var args []string
	args = append(args, "read")
	args = append(args, "-f")
	args = append(args, filePath)
	args = append(args, "1")
	args = append(args, "2")
	args = append(args, "3")

	app := &App{}
	err := app.Main(args)

	if err != nil {
		t.Fatalf("expected : nil, got : %#v", err)
	}
}

func Test_Cli_Read_F_Flag(t *testing.T) {
	t.Parallel()

	filePath := "./testdata/cli/newsarticles2.json"

	generateBackupData(filePath)

	var args []string
	args = append(args, "read")
	args = append(args, "-f")
	args = append(args, filePath)
	args = append(args, "1")
	args = append(args, "2")
	args = append(args, "3")

	app := &App{}
	err := app.Main(args)

	if err != nil {
		t.Fatalf("expected : nil, got : %#v", err)
	}
}

func Test_Cli_Read_J_Flag(t *testing.T) {
	t.Parallel()

	filePath := "./testdata/cli/newsarticles3.json"

	generateBackupData(filePath)

	var args []string
	args = append(args, "read")
	args = append(args, "-f")
	args = append(args, filePath)
	args = append(args, "-j")
	args = append(args, "1")
	args = append(args, "2")
	args = append(args, "3")

	app := &App{}
	err := app.Main(args)

	if err != nil {
		t.Fatalf("expected : nil, got : %#v", err)
	}
}

func Test_Cli_Read_O_Flag(t *testing.T) {
	t.Parallel()

	filePath := "./testdata/cli/newsarticles4.json"

	generateBackupData(filePath)

	var args []string
	args = append(args, "read")
	args = append(args, "-f")
	args = append(args, filePath)
	args = append(args, "-o")
	args = append(args, "./testdata/cli/readcmd_output_stream.json")
	args = append(args, "1")
	args = append(args, "2")
	args = append(args, "3")

	app := &App{}
	err := app.Main(args)

	if err != nil {
		t.Fatalf("expected : nil, got : %#v", err)
	}
}

func Test_Cli_Read_Empty(t *testing.T) {
	t.Parallel()

	var args []string
	args = append(args, "read")

	app := &App{}
	err := app.Main(args)

	if err == nil {
		t.Fatalf("expected : %#v, got : nil", err)
	}
}

func Test_Cli_Stream(t *testing.T) {
	//t.Parallel()

	if runtime.GOOS == "windows" {
		t.Skip()
	}

	filePath := "./testdata/cli/newsarticles5.json"

	generateBackupData(filePath)

	go func() {
		time.Sleep(15 * time.Second)
		//syscall.Kill(syscall.Getpid(), syscall.SIGINT)
		p, err := os.FindProcess(os.Getpid())

		if err != nil {
			return
		}
		p.Signal(os.Interrupt)
	}()

	var args []string
	args = append(args, "stream")
	args = append(args, "-f")
	args = append(args, filePath)
	args = append(args, "Movies")
	args = append(args, "Tech")
	args = append(args, "Sports")

	app := &App{}
	err := app.Main(args)

	if err != nil {
		t.Fatalf("expected : nil, got : %#v", err)
	}
}

func Test_Cli_Stream_F_Flag(t *testing.T) {
	t.Parallel()

	if runtime.GOOS == "windows" {
		t.Skip()
	}

	filePath := "./testdata/cli/newsarticles6.json"

	generateBackupData(filePath)

	go func() {
		time.Sleep(10 * time.Second)
		//syscall.Kill(syscall.Getpid(), syscall.SIGINT)
		p, err := os.FindProcess(os.Getpid())

		if err != nil {
			return
		}
		p.Signal(os.Interrupt)
	}()

	var args []string
	args = append(args, "stream")
	args = append(args, "-f")
	args = append(args, filePath)
	args = append(args, "Movies")
	args = append(args, "Tech")
	args = append(args, "Sports")

	app := &App{}
	err := app.Main(args)

	if err != nil {
		t.Fatalf("expected : nil, got : %#v", err)
	}
}

func Test_Cli_Stream_J_Flag(t *testing.T) {
	//t.Parallel()

	if runtime.GOOS == "windows" {
		t.Skip()
	}

	filePath := "./testdata/cli/newsarticles7.json"

	generateBackupData(filePath)

	go func() {
		time.Sleep(10 * time.Second)
		//syscall.Kill(syscall.Getpid(), syscall.SIGINT)

		p, err := os.FindProcess(os.Getpid())

		if err != nil {
			return
		}
		p.Signal(os.Interrupt)
	}()

	var args []string
	args = append(args, "stream")
	args = append(args, "-f")
	args = append(args, filePath)
	args = append(args, "-j")
	args = append(args, "Movies")
	args = append(args, "Tech")
	args = append(args, "Sports")

	app := &App{}
	err := app.Main(args)

	if err != nil {
		t.Fatalf("expected : nil, got : %#v", err)
	}
}

func Test_Cli_Stream_O_Flag(t *testing.T) {
	//t.Parallel()

	fmt.Printf("TESTING GOOS : %v", runtime.GOOS)

	if runtime.GOOS == "windows" {
		t.Skip()
	}

	filePath := "./testdata/cli/newsarticles8.json"

	generateBackupData(filePath)

	go func() {
		time.Sleep(10 * time.Second)
		//syscall.Kill(syscall.Getpid(), syscall.SIGINT)
		p, err := os.FindProcess(os.Getpid())

		if err != nil {
			return
		}
		p.Signal(os.Interrupt)
	}()

	var args []string
	args = append(args, "stream")
	args = append(args, "-f")
	args = append(args, filePath)
	args = append(args, "-o")
	args = append(args, "./testdata/cli/streamcmd_output_stream.json")
	args = append(args, "Movies")
	args = append(args, "Tech")
	args = append(args, "Sports")

	app := &App{}
	err := app.Main(args)

	if err != nil {
		t.Fatalf("expected : nil, got : %#v", err)
	}
}

func Test_Cli_Stream_Empty(t *testing.T) {
	t.Parallel()

	var args []string
	args = append(args, "stream")

	app := &App{}
	err := app.Main(args)

	if err == nil {
		t.Fatalf("expected : %#v, got : nil", err)
	}
}
func Test_Cli_Clear(t *testing.T) {
	t.Parallel()

	filePath := "./testdata/cli/newsarticles9.json"

	generateBackupData(filePath)

	var args []string
	args = append(args, "clear")
	args = append(args, "-f")
	args = append(args, filePath)

	app := &App{}
	err := app.Main(args)

	if err != nil {
		t.Fatalf("expected : nil, got : %#v", err)
	}
}

func Test_Cli_Root(t *testing.T) {
	t.Parallel()

	filePath := "./NewsServiceBackup.json"

	generateBackupData(filePath)

	var args []string

	app := &App{}
	err := app.Main(args)

	if err != nil {
		t.Fatalf("expected : nil, got : %#v", err)
	}
}

func Test_Cli_Usage(t *testing.T) {
	var args []string
	args = append(args, "-h")

	app := &App{}
	err := app.Main(args)

	if err != nil {
		t.Fatalf("expected : nil, got : %#v", err)
	}
}
