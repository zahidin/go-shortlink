package model

import "time"

type BodyAddShortLink struct {
	Url        string `json:"url" validate:"required,url,excludesall=0x20"`
	CustomName string `json:"customName" validate:"max=10,excludesall=0x20"`
}

type ParamGetLink struct {
	Id string `validate:"required"`
}

type ShortLink struct {
	Id        string    `json:"id"`
	Link      string    `json:"link"`
	CreatedAt time.Time `json:"createdAt"`
}
