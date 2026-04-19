package main

import (
	"errors"
	"flag"
	"fmt"
	"g2/cmd"
	"g2/internal"
	"syscall"
	"time"

	"golang.org/x/term"
)

func main() {

	fmt.Println("Enter pwd: ")
	bytePwd, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		panic(err)
	}
	internal.Init(bytePwd)
	newFile := flag.Bool("new", false, "Create a new file")
	open := flag.String("open", "", "open the file for the specified date")

	flag.Parse()

	if *newFile {
		cmd.CreateNewFile()
	}
	if *open == "latest" {
		err := cmd.OpenFileWithDate(time.Now().Format("02-01-2006"))
		if err != nil {
			fmt.Println(err)
		}
	} else if *open != "" {
		err := cmd.OpenFileWithDate(*open)

		if err != nil {
			errors.Join(errors.New("couldnt open the specified file"), err)
		}
	} else {

	}

}
