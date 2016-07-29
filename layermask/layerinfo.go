package layermask

import (
	"io"
	"encoding/binary"
)

type PSBLayerInfo struct {
	Length uint64

}

// Layer and Mask Information Section
// http://www.adobe.com/devnet-apps/photoshop/fileformatashtml/#50577409_pgfId-1031423
type LayerInfo struct {
	Length uint32
	Count int16
	Records []LayerRecord
	ChannelImageData []Channel
}

// Decode layer info from reader into pointer to layerinfo structure, returns bytes read and/or error
func DecodeLayerInfo(r io.Reader, info *LayerInfo) (uint32, error) {
	var bytesRead uint32
	if err := binary.Read(r, binary.BigEndian, info.Length); err != nil {
		return 0, err
	}
	bytesRead += 4

	if err := binary.Read(r, binary.BigEndian, info.Count); err != nil {
		return bytesRead, err
	}
	bytesRead += 2

	var recordBytesRead int
	for recordBytesRead < info.Length - 6 {

	}


}
