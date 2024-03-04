package controller

import (
	"fmt"

	"github.com/ValSpp/ber1taskanic/database"
	"github.com/ValSpp/ber1taskanic/models"
	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {
	var beritapost models.Berita
	if err := c.BodyParser(&beritapost); err != nil {
		fmt.Println("Unable to parse body")
	}
	if err := database.DB.Create(&beritapost).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "muatan tidak valid",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Selamat!, Postingan Anda sudah aktif",
	})
}

func AllPost(c *fiber.Ctx) {
	
}
