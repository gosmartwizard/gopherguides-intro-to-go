package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"syscall"
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

	filePath := "/tmp/newsarticles1.json"

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

	filePath := "/tmp/newsarticles2.json"

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

	filePath := "/tmp/newsarticles3.json"

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

	filePath := "/tmp/newsarticles4.json"

	generateBackupData(filePath)

	var args []string
	args = append(args, "read")
	args = append(args, "-f")
	args = append(args, filePath)
	args = append(args, "-o")
	args = append(args, "/tmp/readcmd_output_stream.json")
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
	t.Parallel()

	filePath := "/tmp/newsarticles5.json"

	generateBackupData(filePath)

	go func() {
		time.Sleep(15 * time.Second)
		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
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

/*
func Test_Cli_Stream_F_Flag(t *testing.T) {
	t.Parallel()

	filePath := "/tmp/newsarticles6.json"

	generateBackupData(filePath)

	go func() {
		time.Sleep(10 * time.Second)
		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
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
	t.Parallel()

	filePath := "/tmp/newsarticles7.json"

	generateBackupData(filePath)

	go func() {
		time.Sleep(10 * time.Second)
		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
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
	t.Parallel()

	filePath := "/tmp/newsarticles8.json"

	generateBackupData(filePath)

	go func() {
		time.Sleep(10 * time.Second)
		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	}()

	var args []string
	args = append(args, "stream")
	args = append(args, "-f")
	args = append(args, filePath)
	args = append(args, "-o")
	args = append(args, "/tmp/streamcmd_output_stream.json")
	args = append(args, "Movies")
	args = append(args, "Tech")
	args = append(args, "Sports")

	app := &App{}
	err := app.Main(args)

	if err != nil {
		t.Fatalf("expected : nil, got : %#v", err)
	}
}
*/
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

	filePath := "/tmp/newsarticles9.json"

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

	filePath := "/tmp/NewsServiceBackup.json"

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
