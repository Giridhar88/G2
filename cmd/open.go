package cmd

import (
	"fmt"
	"g2/internal"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

// create temp file in tmp dir to write the decrypted content to temporarily, func accepts the bytes contnet to write in the tmp file and returns the path of the temp file
func createTempFile(data []byte) (string, error) {
	tmp, err := os.CreateTemp("", "diary*")
	if err != nil {
		return "", err
	}
	tmpPath := tmp.Name()
	tmp.WriteString(string(data))
	return tmpPath, nil
}

// open the decrypted content in nvim, func accetps the path of tmp file
func openFileInNvim(path string) error {
	defer os.Remove(path)
	cmd := exec.Command("nvim", path)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func OpenFileWithDate(dateStr string) error {
	g2path, err := checkg2Dir()
	if err != nil {
		return err
	}
	path := filepath.Join(g2path, fmt.Sprintf("%s.enc", dateStr))
	err = OpenFile(path)
	if err != nil {
		return err
	}
	return nil
}

// Opens the specified file in nvim after decrypting and storing it in a temp file, func accepts the path of the enc file
func OpenFile(path string) error {

	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	content, err := internal.Decrypt(data)
	if err != nil {
		return err
	}
	tmpPath, err := createTempFile(content)
	if err != nil {
		return err
	}
	err = openFileInNvim(tmpPath)
	if err != nil {
		return err
	}
	return nil
}
