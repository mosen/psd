package psd

import (
	"io"
	"encoding/binary"
	"fmt"
)

// Decode a Photoshop String returning the string, number of bytes read including the length field, and an error
func DecodeString(r io.Reader) (string, uint32, error) {
	var bytesRead uint32

	// The number of characters, unicode characters in this case take up 2 bytes
	var lengthInChars uint32
	if err := binary.Read(r, binary.BigEndian, &lengthInChars); err != nil {
		return "", 0, err
	}
	bytesRead += 4
	var lengthBytes uint32 = lengthInChars * 2

	fmt.Println("Reading string of length", lengthInChars)

	var rawValue []byte = make([]byte, lengthBytes)
	if err := binary.Read(r, binary.BigEndian, &rawValue); err != nil {
		return "", 0, err
	}
	bytesRead += lengthBytes

	strValue := string(rawValue)
	fmt.Println(strValue)

	return strValue, bytesRead, nil
}
