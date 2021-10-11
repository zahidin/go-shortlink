package services

import (
	"context"
	"fmt"
	"shortlink/src/config"
	"shortlink/src/helper/utils"
	"shortlink/src/model"
	"shortlink/src/repositories"
	"time"

	"github.com/sarulabs/di"
)

type ShortLinkService interface {
	AddShortLink(ctx context.Context, linkModel *model.BodyAddShortLink) (string, error)
	GetData(ctx context.Context, id string) (string, error)
}

type ShortLinkServiceImpl struct {
	Repository *repositories.Repository
}

func NewShortLinkService(ioc di.Container) ShortLinkService {
	return &ShortLinkServiceImpl{
		Repository: ioc.Get("repository").(*repositories.Repository),
	}
}

func (s *ShortLinkServiceImpl) GetData(ctx context.Context, id string) (string, error) {
	dataShortLink, err := s.Repository.ShortLink.GetData(ctx, id)
	if err != nil {

		return "", err
	}

	return dataShortLink.Link, nil
}

func (s *ShortLinkServiceImpl) AddShortLink(ctx context.Context, linkModel *model.BodyAddShortLink) (string, error) {
	id := utils.GenerateRandomString(5)

	if linkModel.CustomName != "" {
		_, err := s.Repository.ShortLink.GetData(ctx, linkModel.CustomName)
		if err != nil && err.Error() != "data not found" {
			return "", err
		}

		if err != nil && err.Error() == "data not found" {
			id = linkModel.CustomName
		}
	}

	param := model.ShortLink{
		Id:        id,
		Link:      linkModel.Url,
		CreatedAt: time.Now(),
	}

	dataLink, err := s.Repository.ShortLink.InsertData(ctx, &param)

	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%s/%s", config.GetConfig("BASE_URL"), dataLink.Id)

	return result, nil
}
