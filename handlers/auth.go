package handlers

import (
	"balancer/contracts"
	"balancer/services"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func LoginView(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{})
}

func Login(c *fiber.Ctx) error {
	json := new(contracts.LoginRequest)
	if err := c.BodyParser(json); err != nil {
		fmt.Println(err)
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
			"error":   err.Error(),
		})
	}

	err := services.LoginValidation(json)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Wrong email or password",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"code":    200,
		"message": "Success",
	})
}
