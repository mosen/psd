package psd

import (
	"github.com/mosen/psd/header"
	"github.com/mosen/psd/resource"
	thumbnail "github.com/mosen/psd/resource/thumbnail"
	"errors"
	"fmt"
)

type PSD struct {
	Header              header.Header
	ColorModeDataLength uint32
	ColorModeData       []byte
	ResourceLength      uint32
	Resources           []resource.ImageResource
}

func (p *PSD) ResourceById(id uint16) (*resource.ImageResource, bool) {
	for _, v := range p.Resources {
		if v.Id == id {
			return &v, true
		}
	}

	return nil, false
}


// Get a list of resource id's in this PSD file
func (p *PSD) ResourceIds() ([]int) {
	var ids []int = []int{}

	for _, v := range p.Resources {
		ids = append(ids, int(v.Id))
	}

	return ids
}


// Retrieve the best available thumbnail (either from PSIR or XMP data)
// If not available the second return parameter is false
func (p *PSD) Thumbnail() ([]byte, error) {
	resource, found := p.ResourceById(1036)

	if !found {
		return nil, errors.New("no resource found with thumbnail ID 1036")
	}

	thumb, ok := resource.Data.(*thumbnail.Thumbnail)

	if !ok {
		fmt.Printf("%v\n", resource.Data)
		return nil, errors.New("found resource but could not assert type is Thumbnail")
	}

	return thumb.Data, nil
}

// Retrieve the XMP packet as a string
func (p *PSD) XMPString() (string, error) {
   	resource, found := p.ResourceById(1060)

	if !found {
		return "", errors.New("no resource found with xmp metadata id 1060")
	}

	if xmpData, ok := resource.Data.([]byte); ok {
		return string(xmpData), nil
	} else {
		fmt.Printf("%v\n", resource.Data)
		return "", errors.New("found xmp resource but could not assert type is string")
	}
}
