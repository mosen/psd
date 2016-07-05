package resource

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"image"
	_ "image/jpeg"
	"io"
	"github.com/mosen/psd/resource"
)

// (Photoshop 5.0) Thumbnail resource (supersedes resource 1033).
// Hex: 0x040C
// Dec: 1036
// http://www.adobe.com/devnet-apps/photoshop/fileformatashtml/#50577409_74450
const RESOURCE_ID = 0x40C

const (
	FORMAT_RAWRGB = iota
	FORMAT_JPEGRGB
)

type Thumbnail struct {
	Format         uint32
	Width          uint32
	Height         uint32
	WidthBytes     uint32 // Padded row bytes
	Size           uint32
	CompressedSize uint32 // After compression
	BitDepth       uint16
	PlaneCount     uint16

	Data []byte
}

func (t *Thumbnail) String() string {
	width := fmt.Sprintf("Width: %d", t.Width)
	height := fmt.Sprintf("Height: %d", t.Height)
	depth := fmt.Sprintf("Depth: %d", t.BitDepth)
	size := fmt.Sprintf("Size (Compressed): %d", t.CompressedSize)

	return fmt.Sprintf("Thumbnail (%s, %s, %s, %s)", width, height, depth, size)
}

func (t *Thumbnail) Image() (image.Image, error) {
	switch t.Format {
	case FORMAT_JPEGRGB:
		r := bytes.NewReader(t.Data)
		img, _, err := image.Decode(r)

		return img, err

	default:
		return nil, errors.New("cannot decode this type of thumbnail")
	}
}

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
