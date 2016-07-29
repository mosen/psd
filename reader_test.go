package psd

import (
	"testing"
	"io/ioutil"
	"bytes"
	"fmt"
	_ "github.com/mosen/psd/resource/thumbnail"
	_ "github.com/mosen/psd/resource/xmp"
	//_ "github.com/mosen/psd/resource/versioninfo"
)

func TestDecode(t *testing.T) {
	data, err := ioutil.ReadFile("./bluesquare/Spotty180C.psd")
	reader := bytes.NewReader(data)

	if err != nil {
		t.Error(err)
	}

	psd, err := Decode(reader)

	if err != nil {
		t.Error(err)
	}

	fmt.Printf("Channels: %d\n", psd.Header.Channels)
	fmt.Printf("Bit depth: %d\n", psd.Header.Depth)
}


func TestThumbnail(t *testing.T) {
	data, err := ioutil.ReadFile("./bluesquare/Glass.psd")
	reader := bytes.NewReader(data)

	if err != nil {
		t.Error(err)
	}

	psd, err := Decode(reader)

	if err != nil {
		t.Error(err)
	}

	ids := psd.ResourceIds()
	fmt.Printf("%v\n", ids)

	thumbnail, err := psd.Thumbnail()
	if err != nil {
		t.Errorf("Thumbnail not found in fixture: %s", err)
	} else {
		ioutil.WriteFile("./bluesquare/thumbnail.jpg", thumbnail, 0644)
	}
}

func TestXMPString(t *testing.T) {
	data, err := ioutil.ReadFile("./bluesquare/Glass.psd")
	reader := bytes.NewReader(data)

	if err != nil {
		t.Error(err)
	}

	psd, err := Decode(reader)

	if err != nil {
		t.Error(err)
	}

	xmpString, err := psd.XMPString()

	if err != nil {
		t.Error(err)
	}

	if xmpString[:9] != "<?xpacket" {
		t.Errorf("does not seem like the start of an XMP string: %s\n", xmpString[:9])
	}
}