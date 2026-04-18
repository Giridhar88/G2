package internal

// decrypt the bytes given to the func
func Decrypt(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return []byte{}, nil
	}
	return data, nil
}
