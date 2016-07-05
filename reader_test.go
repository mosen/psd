package psd

import (
	"testing"
	"io/ioutil"
	"bytes"
	"fmt"
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
