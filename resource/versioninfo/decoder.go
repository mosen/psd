package versioninfo

import (
	"encoding/binary"
	"io"
	"github.com/mosen/psd/resource"
	psdstr "github.com/mosen/psd/strings"
	"fmt"
)

// (Photoshop 6.0) Version Info.
//
// Hex: 0x0421
// Dec: 1057
const RESOURCE_ID = 0x421


type VersionInfo struct {
	HasRealMergedData []byte
	WriterName string
	ReaderName string
	Version uint32
}

func (v VersionInfo) String() string {
	return fmt.Sprintf("Writer: %s, Reader: %s, Version: %d", v.WriterName, v.ReaderName, v.Version)
}

func Decode(r io.Reader, id uint16, length uint32) (interface{}, uint32, error) {
	var bytesRead uint32

	var hasRealMergedData []byte = make([]byte, 1)
	if err := binary.Read(r, binary.BigEndian, &hasRealMergedData); err != nil {
		return nil, bytesRead, err
	}
	bytesRead += 1

	writerName, strBytes, err := psdstr.DecodeString(r)
	if err != nil {
		return nil, bytesRead, err
	}
	bytesRead += strBytes

	readerName, strBytes, err := psdstr.DecodeString(r)
	if err != nil {
		return nil, bytesRead, err
	}
	bytesRead += strBytes

	var version uint32
	if err := binary.Read(r, binary.BigEndian, &version); err != nil {
		return nil, bytesRead, err
	}

	vi := &VersionInfo {
		HasRealMergedData: hasRealMergedData,
		WriterName: writerName,
		ReaderName: readerName,
		Version: version,
	}

	fmt.Println(vi)


	return vi, bytesRead, nil
}

func init() {
	resource.RegisterID(RESOURCE_ID, Decode)
}