package resources

type ImageResource struct {
    Magic [4]byte
    Id uint16
    Name string
    DataLength uint32
    Data []byte
}


