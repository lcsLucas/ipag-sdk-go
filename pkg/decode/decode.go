package decode

type Decode interface {
	Decode([]byte, any) error
}
