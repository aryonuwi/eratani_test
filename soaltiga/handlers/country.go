package handlers

import (
	"github.com/aryonuwi/eratani_test/soaltiga/models"
	"github.com/gofiber/fiber/v2"
)

func GetCountry(ctx *fiber.Ctx) error {
	getHighestSpendCountry := models.DB.Raw("SELECT u.country, count(u.country) as count_country,sum(o.total_buy) as buying FROM users u INNER JOIN orders o ON o.id_user =u.id GROUP BY u.country ORDER BY buying DESC")

	return ctx.JSON(getHighestSpendCountry)
}
