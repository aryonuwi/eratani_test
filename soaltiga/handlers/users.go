package handlers

import (
	"github.com/aryonuwi/eratani_test/soaltiga/models"
	"github.com/gofiber/fiber/v2"
)

func GetUser(ctx *fiber.Ctx) error {

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"get": "data get",
	})
}

func PostUser(ctx *fiber.Ctx) error {
	postData := new(models.RegsiterdUsers)

	if err := ctx.BodyParser(postData); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "FAILED",
			"result": err,
		})
	}

	setSaveData := models.Users{
		Country:        postData.Country,
		CreditCardType: postData.CreditCardType,
		CreditCard:     postData.CreditCard,
		FirstName:      postData.FirstName,
		LastName:       postData.LastName,
	}

	saveDataError := models.DB.Create(&setSaveData).Error

	if saveDataError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "FAILED",
			"result": saveDataError,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "SUCCESS",
		"result": postData,
	})
}
