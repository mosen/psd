package psd

import (
	"github.com/mosen/psd/header"
	"github.com/mosen/psd/resources"
	"image"
	_ "image/jpeg"
)

type PSD struct {
	Header header.Header
	ColorModeDataLength uint32
	ColorModeData []byte
	ResourceLength uint32
	Resources []resources.ImageResource
}

// Retrieve the best available thumbnail (either from PSIR or XMP data)
// If not available the second return parameter is false
func (p *PSD) Thumbnail (image.Image, bool) {

}

// Retrieve the XMP packet as a string
func (p *PSD) XMPString (string, bool) {

}
