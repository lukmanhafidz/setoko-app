package model

import (
	"net/http"
	"setokoapp/constants"

	"github.com/gofiber/fiber/v2"
)

var Config = struct {
	Port     int `env:"port"`
	Postgres struct {
		Host     string `env:"host"`
		User     string `env:"user"`
		Password string `env:"password"`
		Name     string `env:"name"`
		Port     string `env:"port"`
	} `env:"postgres"`
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

func (br *BaseResp) Error(responseCode, responseDesc string) BaseResp {
	return BaseResp{
		ResponseCode: responseCode,
		ResponseDesc: responseDesc,
		Data:         nil,
	}
}

func ResponseOk(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(http.StatusOK).JSON(new(BaseResp).OK(nil))
}
