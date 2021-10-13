package controllers

import (
	"fmt"
	"net/http"
	"shortlink/src/helper/utils"
	"shortlink/src/model"
	"shortlink/src/services"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sarulabs/di"
)

type ShortLinkController interface {
	AddShortLink(ctx *fiber.Ctx) error
	GetLink(ctx *fiber.Ctx) error
	GetAllData(ctx *fiber.Ctx) error
}

type ShortLinkControllerImpl struct {
	Service *services.Services
}

func NewShortLinkController(ioc di.Container) ShortLinkController {
	return &ShortLinkControllerImpl{
		Service: ioc.Get("service").(*services.Services),
	}
}

func (c *ShortLinkControllerImpl) GetLink(ctx *fiber.Ctx) error {
	params := new(model.ParamGetLink)
	params.Id = ctx.Params("id")
	validationErr := utils.Validation(*params)
	if validationErr != nil {
		return utils.ResponseError(ctx, *validationErr[0], http.StatusBadRequest)
	}

	result, err := c.Service.ShortLink.GetData(ctx.Context(), params.Id)
	if err != nil && err.Error() != "data not found" {
		return utils.ResponseError(ctx, err, http.StatusInternalServerError)
	}

	if err != nil && err.Error() == "data not found" {
		return utils.ResponseError(ctx, err, http.StatusNotFound)

	}

	return utils.ResponseSuccess(ctx, result, "Success Get Link", http.StatusOK)
}

func (c *ShortLinkControllerImpl) AddShortLink(ctx *fiber.Ctx) error {
	body := new(model.BodyAddShortLink)

	if err := ctx.BodyParser(body); err != nil {
		return utils.ResponseError(ctx, err, http.StatusInternalServerError)
	}

	validationErr := utils.Validation(*body)
	if validationErr != nil {
		return utils.ResponseError(ctx, *validationErr[0], http.StatusBadRequest)
	}

	fmt.Println(body)

	result, err := c.Service.ShortLink.AddShortLink(ctx.Context(), body)
	if err != nil {
		return utils.ResponseError(ctx, err, http.StatusInternalServerError)
	}

	return utils.ResponseSuccess(ctx, result, "Success Add Short Link", http.StatusCreated)
}

func (c *ShortLinkControllerImpl) GetAllData(ctx *fiber.Ctx) error {
	// Testing Response Pagination
	data := []model.ShortLink{{Id: "123", Link: "http://googla.com", CreatedAt: time.Now()}}
	meta := utils.ResponseMeta{
		Page:        1,
		Limit:       10,
		TotalPage:   5,
		TotalRecord: len(data),
	}

	return utils.ResponsePagination(ctx, data, "Success Get All Data", meta, http.StatusOK)
}
