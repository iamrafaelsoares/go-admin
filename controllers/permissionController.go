package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iamrafaelsoares/go-admin/database"
	"github.com/iamrafaelsoares/go-admin/models"
)

func AllPermissions(c *fiber.Ctx) error {
	var permissions []models.Permission

	database.DB.Find(&permissions)

	return c.JSON(permissions)
}
