package layermask

import (
	"image"
	"io"
	"encoding/binary"
)

const (
	CHID_REAL_USER_MASK = iota - 3
	CHID_USER_MASK
	CHID_TRANS_MASK
	RED
	GREEN
	BLUE
)

const (
	LAYER_FLAG_TRANSPARENCY_PROTECTED = iota
	LAYER_FLAG_VISIBLE
	LAYER_FLAG_OBSOLETE
	LAYER_FLAG_USE_BIT_4
	LAYER_FLAG_BIT_4_DATA
)

const (
	MASK_FLAG_POSITION_RELATIVE_TO_LAYER = iota
	MASK_FLAG_LAYER_MASK_DISABLED
	MASK_FLAG_INVERT_LAYER_MASK
	MASK_FLAG_FROM_OTHER_DATA
	MASK_FLAG_HAS_PARAMS
)

const (
	MASK_PARAM_USER_DENSITY = iota
	MASK_PARAM_USER_FEATHER
	MASK_PARAM_VECTOR_DENSITY
	MASK_PARAM_VECTOR_FEATHER
)

var BlendModes map[string]string = map[string]string{
	"SoCo": "Solid Color",
	"GdFl": "Gradient",
	"PtFl": "Pattern",
	"brit": "Brightness/Contrast",
	"levl": "Levels",
	"curv": "Curves",
	"expA": "Exposure",
	"vibA": "Vibrance",
	"hue ": "Hue/Saturation", // PS 4.0
	"hue2": "Hue/Saturation",
	"blnc": "Color Balance",
	"blwh": "Black and White",
	"phfl": "Photo Filter",
	"mixr": "Channel Mixer",
	"clrL": "Color Lookup",
	"nvrt": "Invert",
	"post": "Posterize",
	"thrs": "Threshold",
	"grdm": "Gradient Map",
	"selc": "Selective Color",
}

type PSBChannel struct {
	Id         uint16
	DataLength uint64
}

type Channel struct {
	Id         uint16
	DataLength uint32
}


// This also applies to adjustment layers as well as mask layers
type Mask struct {
	Length       uint32
	Rect         image.Rectangle
	DefaultColor uint8
	Flags        uint8
	Parameters   uint8
	Padding      [2]byte
	RealFlags    uint8
	Rect2        image.Rectangle
}

type BlendingRange struct {
	SourceRange uint32
	DestRange   uint32
}

type BlendingRanges struct {
	Length uint32
	Ranges []BlendingRange
}

// http://www.adobe.com/devnet-apps/photoshop/fileformatashtml/#50577409_13084
type LayerRecord struct {
	Rect           image.Rectangle
	ChannelCount   uint16
	Channels       []Channel
	Signature      [4]byte // '8BIM'
	BlendMode      [4]byte // 4 char blend mode string
	Opacity        uint8
	Clipping       uint8
	Flags          uint8
	Filler         uint8
	ExtraLength    uint32
	Masks          []Mask
	BlendingRanges BlendingRanges
	Name           string
}

func DecodeLayerRecord(r io.Reader, record *LayerRecord) (uint32, error) {
	var bytesRead int

	var top, left, bottom, right uint32
	if err := binary.Read(r, binary.BigEndian, &top); err != nil {
		return bytesRead, err
	}

	if err := binary.Read(r, binary.BigEndian, &left); err != nil {
		return bytesRead, err
	}

	if err := binary.Read(r, binary.BigEndian, &bottom); err != nil {
		return bytesRead, err
	}

	if err := binary.Read(r, binary.BigEndian, &right); err != nil {
		return bytesRead, err
	}
	record.Rect = image.Rect(left, top, right, bottom)
	bytesRead += 16

	if err := binary.Read(r, binary.BigEndian, &record.ChannelCount); err != nil {
		return bytesRead, err
	}
	bytesRead += 2



}
