package psd

import (
	"io"
	"github.com/mosen/psd/header"
	"encoding/binary"
	"errors"
	"github.com/mosen/psd/resources"
	"fmt"
)

const DEBUG = false

type decoder struct {
	reader io.Reader
}

// http://www.adobe.com/devnet-apps/photoshop/fileformatashtml/#50577409_71638
func readColorData(r io.Reader) (uint32, []byte, error) {
	var colorModeDataLength uint32
	err := binary.Read(r, binary.BigEndian, &colorModeDataLength)

	if err != nil {
		return 0, nil, err
	}

	if colorModeDataLength == 0 {
		return 0, nil, nil
	} else {
		colorModeData := make([]byte, colorModeDataLength)
		if err := binary.Read(r, binary.BigEndian, &colorModeData); err != nil {
			return 0, nil, err
		}

		return colorModeDataLength, colorModeData[:], nil
	}
}

func decodeImageResources(r io.Reader) (uint32, []resources.ImageResource, error) {
  	var imageResourceLength uint32
	err := binary.Read(r, binary.BigEndian, &imageResourceLength)

	if err != nil {
		return 0, nil, err
	}

	fmt.Println("resource length:", imageResourceLength)

	var bytesRead uint32
	var items []resources.ImageResource
	for bytesRead < imageResourceLength {
	 	resource, n, err := resources.Decode(r)
		bytesRead += n

		if DEBUG {
			fmt.Printf("Bytes read: %d\n", bytesRead)
		}

		if err != nil {
			return 0, nil, err
		}

		items = append(items, *resource)
	}

	return imageResourceLength, items, nil
}

func (d *decoder) decode(r io.Reader) (*PSD, error) {
	header, err := header.Decode(r)
	if err != nil {
		return nil, err
	}

	psd := &PSD{Header: *header}

	colorModeDataLength, colorModeData, err := readColorData(r)

	if err != nil {
		return nil, errors.New("could not read color mode data")
	}

	psd.ColorModeDataLength = colorModeDataLength
	psd.ColorModeData = colorModeData

	resourceLength, resources, err := decodeImageResources(r)
	if err != nil {
		return nil, err
	}

	psd.ResourceLength = resourceLength
	psd.Resources = resources

	return psd, nil
}

func Decode(r io.Reader) (*PSD, error) {
	var d decoder

	psd, err := d.decode(r)

	if err != nil {
		return nil, err
	}

	return psd, nil
}

