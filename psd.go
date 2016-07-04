package psd

import (
	"github.com/mosen/psd/header"
	"github.com/mosen/psd/resources"
)

type PSD struct {
	Header header.Header
	ColorMode [4]byte
	ResourceLength uint32
	Resources []resources.ImageResource

}