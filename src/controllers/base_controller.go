package controllers

import "github.com/sarulabs/di"

type Controllers struct {
	Shortlink ShortLinkController
}

func NewController(ioc di.Container) *Controllers {
	return &Controllers{
		Shortlink: NewShortLinkController(ioc),
	}
}
