package header

const headerlength = 30
const magic string = "8BPS"
const versionPsd = 1
const versionPsb = 2

const (
	bitmap  = iota
	grayscale
	indexed
	rgb
	cmyk
	multichannel
	duotone
	lab
)

