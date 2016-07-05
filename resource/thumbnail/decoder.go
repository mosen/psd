package resource

import (
	"encoding/binary"
	_ "image/jpeg"
	"io"
	"github.com/mosen/psd/resource"
)

func Decode(r io.Reader, id uint16, length uint32) (interface{}, uint32, error) {
	var bytesRead uint32

	var format uint32
	if err := binary.Read(r, binary.BigEndian, &format); err != nil {
		return nil, bytesRead, err
	}
	bytesRead += 4

	var width uint32
	if err := binary.Read(r, binary.BigEndian, &width); err != nil {
		return nil, bytesRead, err
	}
	bytesRead += 4

	var height uint32
	if err := binary.Read(r, binary.BigEndian, &height); err != nil {
		return nil, bytesRead, err
	}
	bytesRead += 4

	var widthBytes uint32
	if err := binary.Read(r, binary.BigEndian, &widthBytes); err != nil {
		return nil, bytesRead, err
	}
	bytesRead += 4

	var size uint32
	if err := binary.Read(r, binary.BigEndian, &size); err != nil {
		return nil, bytesRead, err
	}
	bytesRead += 4

	var compressedSize uint32
	if err := binary.Read(r, binary.BigEndian, &compressedSize); err != nil {
		return nil, bytesRead, err
	}
	bytesRead += 4

	var bitDepth uint16
	if err := binary.Read(r, binary.BigEndian, &bitDepth); err != nil {
		return nil, bytesRead, err
	}
	bytesRead += 2

	var planeCount uint16
	if err := binary.Read(r, binary.BigEndian, &planeCount); err != nil {
		return nil, bytesRead, err
	}
	bytesRead += 2

	var data []byte = make([]byte, compressedSize)
	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, bytesRead, err
	}
	bytesRead += compressedSize

	return &Thumbnail{
		format,
		width,
		height,
		widthBytes,
		size,
		compressedSize,
		bitDepth,
		planeCount,
		data,
	}, bytesRead, nil
}

func init() {
	resource.RegisterID(RESOURCE_ID, Decode)
}
