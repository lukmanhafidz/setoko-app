package model

import (
	"net/http"
	"setokoapp/constants"

	"github.com/gofiber/fiber/v2"
)

var Config = struct {
	Port int `env:"port"`
}{}

type BaseResp struct {
	ResponseCode string      `json:"responseCode"`
	ResponseDesc string      `json:"responseDesc"`
	Data         interface{} `json:"data"`
}

func (br *BaseResp) OK(data interface{}) BaseResp {
	return BaseResp{
		ResponseCode: constants.RC_SUCCESS,
		ResponseDesc: constants.RD_SUCCESS,
		Data:         data,
	}
}

func ResponseOk(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(http.StatusOK).JSON(new(BaseResp).OK(nil))
}
