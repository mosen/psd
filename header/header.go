package header

import (
	"io"
	"fmt"
	"encoding/binary"
	"errors"
)

type Header struct {
 	Magic [4]byte
	Version uint16
	Channels uint16
	Height uint32
	Width uint32
	Depth uint16
	Mode uint16
}

func (h Header) String() (string) {
	magic := fmt.Sprintf("MAGIC: %s", h.Magic)
	version := fmt.Sprintf("VER: %d", h.Version)
	channels := fmt.Sprintf("CHANNELS: %d", h.Channels)
	height := fmt.Sprintf("HEIGHT: %d", h.Height)
	width := fmt.Sprintf("WIDTH: %d", h.Width)
	depth := fmt.Sprintf("DEPTH: %d", h.Depth)
	mode := fmt.Sprintf("MODE: %d", h.Mode)

	s := fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s", magic, version, channels, height, width, depth, mode)

	return s
}

func Decode(r io.Reader) (*Header, error) {
	var magic [4]byte
	err := binary.Read(r, binary.BigEndian, &magic)

	if err != nil {
		return nil, err
	}

	if magic != [4]byte{56, 66, 80, 83} {
		return nil, errors.New("file does not contain psd magic '8BPS'.")
	}

	header := &Header{
		Magic: magic,
	}

	var version uint16
	if err := binary.Read(r, binary.BigEndian, &version); err != nil {
		return nil, err
	}
	header.Version = version

	// Reserved
	var unused [6]byte
	if err := binary.Read(r, binary.BigEndian, &unused); err != nil {
		return nil, err
	}

	var channels uint16
	if err := binary.Read(r, binary.BigEndian, &channels); err != nil {
		return nil, err
	}
	header.Channels = channels

	var height uint32
	if err := binary.Read(r, binary.BigEndian, &height); err != nil {
		return nil, err
	}
	header.Height = height

	var width uint32
	if err := binary.Read(r, binary.BigEndian, &width); err != nil {
		return nil, err
	}
	header.Width = width

	var depth uint16
	if err := binary.Read(r, binary.BigEndian, &depth); err != nil {
		return nil, err
	}
	header.Depth = depth

	var mode uint16
	if err := binary.Read(r, binary.BigEndian, &mode); err != nil {
		return nil, err
	}
	header.Mode = mode

	return header, nil
}
