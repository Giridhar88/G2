package main

import (
	"flag"
	"fmt"
	"g2/cmd"
)

func main() {

	newFile := flag.Bool("new", false, "Create a new file")
	fetchLatestFile := flag.Bool("fetch", false, "Fetch the latest created file")

	flag.Parse()

	if *newFile {
		cmd.CreateNewFile()
	}

	if *fetchLatestFile {
		fmt.Println("fetching the latest file")
	}

}
