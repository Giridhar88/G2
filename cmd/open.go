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
	defer tmp.Close()
	tmpPath := tmp.Name()
	tmp.WriteString(string(data))
	return tmpPath, nil
}

// open the decrypted content in nvim, func accetps the path of tmp file
func openFileInNvim(path string) error {
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

// opens the file with the specified date, attaches the hidden dir to the path and opens it using the OpenFile func
func OpenFileWithDate(dateStr string, pwd []byte) error {
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

// Opens the specified file in nvim after decrypting and storing it in a temp file, func accepts the path of the enc file
func OpenFile(path string, pwd []byte) error {

	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	content, err := internal.Decrypt(data, pwd)
	if err != nil {
		return err
	}
	tmpPath, err := createTempFile(content)
	if err != nil {
		return err
	}
	defer os.Remove(tmpPath)
	err = openFileInNvim(tmpPath)
	if err != nil {
		return err
	}
	newContents, err := os.ReadFile(tmpPath)
	if err != nil {
		return err
	}
	file2write, err := os.OpenFile(path, os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file2write.Close()

	encryptedContents, err := internal.Encrypt(newContents, pwd)
	if err != nil {
		return err
	}
	_, err = file2write.Write(encryptedContents)
	if err != nil {
		return err
	}

	return nil
}
