package repositories

import "github.com/sarulabs/di"

type Repository struct {
	ShortLink ShortLinkRepository
}

func NewRepository(ioc di.Container) *Repository {
	return &Repository{
		ShortLink: NewShortLinkRepositorty(ioc),
	}
}
