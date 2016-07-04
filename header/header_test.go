package header

import (
	"testing"
	"io/ioutil"
	"fmt"
	"bytes"
)

func TestDecode(t *testing.T) {
	data, err := ioutil.ReadFile("../bluesquare/BlueSquare.psd")
	reader := bytes.NewReader(data)

	if err != nil {
		t.Error(err)
	}

	header, err := Decode(reader)

	if err != nil {
		t.Error(err)
	}

	fmt.Println("%v", header)
}
