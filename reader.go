package psd

import (
	"io"
	"github.com/mosen/psd/header"
)

type decoder struct {
	reader io.Reader
}

func (d *decoder) decode(r io.Reader) error {
	header := header.Header.Decode(r)
}

func Decode(r io.Reader) (PSD, error) {
	var d decoder
	if err := d.decode(r); err != nil {
		return nil, err
	}

	psd := &PSD{}

	return psd, nil
}

