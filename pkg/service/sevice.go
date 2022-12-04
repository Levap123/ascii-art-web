package service

type Service struct {
	Ascii
}

func NewService() *Service {
	return &Service{Ascii: newAsciiService()}
}
