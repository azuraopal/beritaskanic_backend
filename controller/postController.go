package controller

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/ValSpp/ber1taskanic/database"
	"github.com/ValSpp/ber1taskanic/models"
	"github.com/ValSpp/ber1taskanic/util"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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
	// if err := c.BodyParser(&berita); err != nil {
	// 	fmt.Println("Unable to parse body")
	// }
	database.DB.Where("id=?", id).Preload("User").First(&berita)
	if berita.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Gagal Mengupdate, Silahkan coba lagi!",
		})
	}
	database.DB.Model(&berita).Updates(berita)
	return c.JSON(fiber.Map{
		"message": "post updated successfully",
	})
}

func UniquePost(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, _ := util.Parsejwt(cookie)
	var berita []models.Berita
	database.DB.Model(&berita).Where("user_id=?", id).Preload("User").Find(&berita)

	return c.JSON(berita)
}

func DeletePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	berita := models.Berita{
		Id: uint(id),
	}

	var postberita models.Berita
	database.DB.Where("id=?", id).Preload("User").First(&postberita)
	if postberita.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Maaf, Postingan tidak ada",
		})
	}

	deleteQuery := database.DB.Delete(&berita)

	if errors.Is(deleteQuery.Error, gorm.ErrRecordNotFound) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Maaf, Postingan tidak ada",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Postingan berhasil dihapus!",
	})

}
