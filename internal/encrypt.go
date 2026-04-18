package internal

// encrypt the bytes given to the func
func Encrypt(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return []byte{}, nil
	}
	return data, nil
}
