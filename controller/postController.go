package controller

import (
	"fmt"
	"math"
	"strconv"

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

func AllPost(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 5
	offset := (page - 1) * limit
	var total int64
	var ambilberita []models.Berita
	database.DB.Preload("User").Offset(offset).Limit(limit).Find(&ambilberita)
	database.DB.Model(&models.Berita{}).Count(&total)
	return c.JSON(fiber.Map{
		"data": ambilberita,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})
}

func DetailPost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var postberita models.Berita
	database.DB.Where("id=?", id).Preload("User").First(&postberita)
	return c.JSON(fiber.Map{
		"data": postberita,
	})
}

func UpdatePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	berita := models.Berita{
		Id: uint(id),
	}
	if err := c.BodyParser(&berita); err != nil {
		fmt.Println("Unable to parse body")
	}
	database.DB.Model(&berita).Updates(berita)
	return c.JSON(berita)
}
