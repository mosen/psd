package resource

import (
	"io"
	"github.com/mosen/psd/resource"
)

// (Photoshop 7.0) XMP metadata. File info as XML description. See http://www.adobe.com/devnet/xmp/
// Hex: 0x0424
// Dec: 1060
const RESOURCE_ID = 0x424

func Decode(r io.Reader, id uint16, length uint32) ([]byte, uint32, error) {
	var data []byte = make([]byte, length)

	if _, err := r.Read(data); err != nil {
		return []byte{}, length, nil
	}

	return data, length, nil
}

func init() {
	resource.RegisterID(RESOURCE_ID, Decode)
}