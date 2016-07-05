package resource

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	_ "image/jpeg"
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
