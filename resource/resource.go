package resource

import (
	"encoding/binary"
	"fmt"
	"io"
)

type ImageResource struct {
	Magic      [4]byte
	Id         uint16
	Name       string
	DataLength uint32
	Data       []byte
}

type ImageResourceData interface {
	Id() uint16
}

type ImageResourceDecoder struct {
	id uint16
	decode func(io.Reader, uint16, uint32) (interface{}, uint32, error)
}

var decoders []ImageResourceDecoder

// Fetch decoder for given resource id, bool is false if none was found
func decoderForId(id uint16) (*ImageResourceDecoder, bool) {
	for _, v := range decoders {
		if v.id == id {
			return &v, true
		}
	}
	return nil, false
}

// Register a photoshop image resource decoder which returns its own resource instance, number of bytes read, error
func RegisterID(id uint16, decode func(io.Reader, uint16, uint32) (interface{}, uint32, error)) {
	decoders = append(decoders, ImageResourceDecoder{id, decode})
}

// Decode the name of the image resource section, which is often just null
func decodeName(r io.Reader) (string, uint32, error) {
	var strlen uint8
	var bytesRead uint32

	if err := binary.Read(r, binary.BigEndian, &strlen); err != nil {
		return "", 0, err
	}
	bytesRead += 1

	if DEBUG {
		fmt.Printf("Name is %d bytes\n", strlen)
	}

	if strlen%2 == 1 {
		strlen += 1
	}

	if DEBUG {
		fmt.Printf("Name will actually read %d bytes\n", strlen)
	}

	if strlen == 0 {
		// Even a zero length string is nul terminated, so we have to read the following zero
		discard := make([]byte, 1)
		r.Read(discard)
		bytesRead += 1
		return "", bytesRead, nil
	} else {
		var name []byte = make([]byte, strlen)

		if err := binary.Read(r, binary.BigEndian, &name); err != nil {
			return "", bytesRead, err
		}
		bytesRead += uint32(strlen)

		fmt.Printf("Name: %s", string(name))

		return string(name), bytesRead, nil
	}

	return "", bytesRead, nil
}

// Returns number of bytes read or error
func decodeData(r io.Reader, id uint16, length uint32) (uint32, error) {
	if length%2 == 1 {
		length += 1
		if DEBUG {
			fmt.Printf("Length padded to: %d\n", length)
		}
	}

	decoder, ok := decoderForId(id)

	if ok {
		data, byteCount, err := decoder.decode(r, id, length)
		if err != nil {
			return byteCount, err
		}

		fmt.Printf("%v\n", data)
		return byteCount, nil
	} else {
		var crap []byte = make([]byte, length)
		_, err := r.Read(crap)

		if err != nil {
			return length, err
		}

		return length, nil
	}
}

func Decode(r io.Reader) (*ImageResource, uint32, error) {
	var bytesRead uint32
	var magic [4]byte
	if err := binary.Read(r, binary.BigEndian, &magic); err != nil {
		return nil, 0, err
	}
	bytesRead += 4

	var id uint16
	if err := binary.Read(r, binary.BigEndian, &id); err != nil {
		return nil, 0, err
	}
	bytesRead += 2

	if DEBUG {
		fmt.Printf("Image resource section id is 0x%04x\n", id)
	}

	name, n, err := decodeName(r)
	if err != nil {
		return nil, 0, err
	}
	bytesRead += n

	var imageResourceLength uint32
	if err := binary.Read(r, binary.BigEndian, &imageResourceLength); err != nil {
		return nil, 0, err
	}
	bytesRead += 4

	if DEBUG {
		fmt.Printf("Image resource section is %d bytes\n", imageResourceLength)
	}

	dataLen, err := decodeData(r, id, imageResourceLength)

	if err != nil {
		return nil, bytesRead, err
	}

	bytesRead += dataLen

	return &ImageResource{
		magic,
		id,
		name,
		imageResourceLength,
		[]byte{},
	}, bytesRead, nil
}


