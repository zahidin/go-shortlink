package services

import "github.com/sarulabs/di"

type Services struct {
	ShortLink ShortLinkService
}

func NewService(ioc di.Container) *Services {
	return &Services{
		ShortLink: NewShortLinkService(ioc),
	}
}
