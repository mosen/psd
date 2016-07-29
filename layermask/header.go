package layermask

import (
	"io"
	"encoding/binary"
)

type Header struct {
	Length uint32

}

type LayerMasks struct {
	Header Header
	LayerInfo []LayerInfo
}

func DecodeLayerMasks(r io.Reader) (uint32, *LayerMasks, error) {
	var bytesRead uint32
	var length uint32
	if err := binary.Read(r, binary.BigEndian, &length); err != nil {
		return bytesRead, nil, err
	}

	var header Header = Header{length}
	//var layerMasks = LayerMasks{Header: header}


}