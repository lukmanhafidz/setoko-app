package interfaces

import (
	"setokoapp/usecases"

	"github.com/gofiber/fiber/v2"
)

type generateHandler struct {
	generateUsecase usecases.IGenerateUsecase
}

func NewGenerateHandler(generateUsecase usecases.IGenerateUsecase) *generateHandler {
	return &generateHandler{generateUsecase: generateUsecase}
}

func (h *generateHandler) GenerateReceipt(ctx *fiber.Ctx) error {
	response := h.generateUsecase.GenerateReceipt(ctx.Params("trxId"))
	return ctx.Status(fiber.StatusOK).JSON(response)
}
