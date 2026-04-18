package main

import (
	"errors"
	"flag"
	"fmt"
	"g2/cmd"
)

func main() {

	newFile := flag.Bool("new", false, "Create a new file")
	fetchLatestFile := flag.Bool("fetch", false, "Fetch the latest created file")
	open := flag.String("open", "", "open the file for the specified date")

	flag.Parse()

	if *newFile {
		cmd.CreateNewFile()
	}
	if *open != "" {
		err := cmd.OpenFile(*open)

		if err != nil {
			errors.Join(errors.New("couldnt open the specified file"), err)
		}
	}
	if *fetchLatestFile {
		fmt.Println("fetching the latest file")
	}

}
