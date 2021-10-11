package repositories

import (
	"context"
	"errors"
	"shortlink/src/helper/database/firebase"
	"shortlink/src/model"

	"github.com/sarulabs/di"
)

type ShortLinkRepository interface {
	InsertData(ctx context.Context, model *model.ShortLink) (*model.ShortLink, error)
	GetData(ctx context.Context, id string) (*model.ShortLink, error)
}

type ShortLinkRepositoryImpl struct{}

func NewShortLinkRepositorty(ioc di.Container) ShortLinkRepository {
	return &ShortLinkRepositoryImpl{}
}

func (r *ShortLinkRepositoryImpl) GetData(ctx context.Context, id string) (*model.ShortLink, error) {
	var result model.ShortLink

	db, err := firebase.FirebaseDatabase(ctx)
	if err != nil {
		return nil, err
	}

	if err := db.NewRef("link/"+id).Get(ctx, &result); err != nil {
		return nil, err
	}

	if result.Id == "" {
		return nil, errors.New("data not found")
	}

	return &result, nil
}

func (r *ShortLinkRepositoryImpl) InsertData(ctx context.Context, model *model.ShortLink) (*model.ShortLink, error) {

	db, err := firebase.FirebaseDatabase(ctx)
	if err != nil {
		return nil, err
	}

	if err := db.NewRef("link/"+model.Id).Set(ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}
