package resource

import (
	"io"
	"encoding/binary"
)

const (
	_ = iota
	UNIT_INCHES
	UNIT_CM
	UNIT_POINTS
	UNIT_PICAS
	UNIT_COLUMNS
)

type ResolutionInfo struct {
	HorizontalPPI uint32
	HorizontalUnit int16 // 1 ppi, 1 pixels per cm
	WidthUnit int16 // From consts above
	VerticalPPI uint32
	VerticalUnit int16 // 1 ppi, 1 pixels per cm
	HeightUnit int16 // From consts above
}

func DecodeResolutionInfo(r io.Reader, id uint16, length uint32) (*ResolutionInfo, uint32, error) {
	resInfo := ResolutionInfo{}
	if err := binary.Read(r, binary.BigEndian, &resInfo); err != nil {
		return nil, 0, err
	}

	return &resInfo, 16, nil
}