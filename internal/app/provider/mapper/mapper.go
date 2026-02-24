package mapper

type mapper struct{}

type Mapper interface {
}

func New() Mapper {
	return &mapper{}
}
