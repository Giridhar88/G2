package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func CreateNewFile() (bool, error) {
	currDate := time.Now()
	dateStr := currDate.Format("02-01-2006")

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return false, err
	}

	//ensure .g2files exists inside homedir
	dir := filepath.Join(homeDir, ".g2files")

	err = os.MkdirAll(dir, 0700)
	if err != nil {
		return false, err
	}

	//create the .enc file to store the encrypted content
	path := filepath.Join(homeDir, ".g2files", fmt.Sprintf("%s.enc", dateStr))
	if _, err := os.Stat(path); err == nil {
		fmt.Println("File already exists")
		// return true, nil
	}

	fmt.Printf("creating a file at %s\n", path)

	newFile, err := os.Create(path)
	if err != nil {
		return false, err
	}
	defer newFile.Close()

	//create the tmp file to open it from the editor
	tmp, err := os.CreateTemp("", "diary-*")
	tmpPath := tmp.Name()
	if err != nil {
		return false, err
	}
	tmp.WriteString("Random ass shit")
	tmp.Close()
	defer os.Remove(tmp.Name())

	//open the editor
	cmd := exec.Command("nvim", tmpPath)
	//copy the output to the terminal
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println("error opening the editor")
		return false, err
	}
	return true, nil
}
