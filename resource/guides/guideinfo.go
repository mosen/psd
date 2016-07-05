package guides

// Grid and guides information
// Hex: 0x0408
// Dec: 1032

type GuideInfo struct {
	// Header
	Version uint32
	Future [8]byte
	GuideCount uint32

	Guides []GuideResource
}

const (
	GUIDE_DIRECTION_VERTICAL = iota
	GUIDE_DIRECTION_HORIZONTAL
)

type GuideResource struct {
	Location uint32
	Direction uint8
}
