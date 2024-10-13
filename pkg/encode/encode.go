package encode

type Encode interface {
	Encode(data any) ([]byte, error)
	ContentType() string
}
