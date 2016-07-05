package psd

import (
	"testing"
	"io/ioutil"
	"bytes"
	"fmt"
	_ "github.com/mosen/psd/resource/thumbnail"
	_ "github.com/mosen/psd/resource/xmp"
)

func TestDecode(t *testing.T) {
	data, err := ioutil.ReadFile("./bluesquare/Glass.psd")
	reader := bytes.NewReader(data)

	if err != nil {
		t.Error(err)
	}

	psd, err := Decode(reader)

	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%v\n", psd)
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

	fmt.Printf("%v\n", xmpString)
}