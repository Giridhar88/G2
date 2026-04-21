package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// checks if g2 dir exists in the home dir
func checkg2Dir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Join(errors.New("error finding the home dir"), err)
	}
	dir := filepath.Join(homeDir, ".g2files")

	err = os.MkdirAll(dir, 0700)

	if err != nil {
		return dir, err
	}

	return dir, nil
}

// create new file with the curr date. creates the enc file and opens the raw content in a tmp file, if enc file already exists opens it in tmp file after decrypting it
func CreateNewFile(pwd []byte) error {
	currDate := time.Now()
	dateStr := currDate.Format("02-01-2006")

	g2path, err := checkg2Dir()
	if err != nil {
		return err
	}

	path := filepath.Join(g2path, fmt.Sprintf("%s.enc", dateStr))
	err = OpenFile(path, pwd)
	if err != nil {
		return err
	}

	return nil
}
