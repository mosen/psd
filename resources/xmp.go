package resources

import "io"

func DecodeXMP(r io.Reader, id uint16, length uint32) ([]byte, uint32, error) {
	var data []byte = make([]byte, length)

	if _, err := r.Read(data); err != nil {
		return []byte{}, length, nil
	}

	return data, length, nil
}
