package handlers

import (
	"balancer/contracts"
	"balancer/models"
	"balancer/services"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Domains(c *fiber.Ctx) error {
	domains, err := services.GetAllDomains()
	if err != nil {
		resp := contracts.Response{
			Code:    400,
			Message: err.Error(),
		}
		return c.JSON(resp)
	}

	return c.Render("domains", fiber.Map{"domains": domains})
}

func CreateDomain(c *fiber.Ctx) error {
	json := contracts.CreateDomainRequest{}
	if err := c.BodyParser(&json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
			"error":   err.Error(),
		})
	}

	domain := models.Domain{
		Label:   json.DomainLabel,
	}

	err := services.CreateDomain(domain)

	if err != nil {
		resp := contracts.Response{
			Code:    400,
			Message: err.Error(),
		}
		return c.JSON(resp)
	}

	resp := contracts.Response{
		Code:    200,
		Message: "success",
	}

	return c.JSON(resp)
}

func DeleteDomain(c *fiber.Ctx) error {
	json := contracts.DeleteDomainRequest{}
	if err := c.BodyParser(&json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
			"error":   err.Error(),
		})
	}

	domainId := json.DomainId
	fmt.Println(domainId)

	err := services.DeleteDomain(domainId)
	if err != nil {
		resp := contracts.Response{
			Code:    400,
			Message: err.Error(),
		}
		return c.JSON(resp)
	}

	resp := contracts.Response{
		Code:    200,
		Message: "success",
	}

	return c.JSON(resp)
}

func GetAllDomains(c *fiber.Ctx) error {
	domains, err := services.GetAllDomains()
	if err != nil {
		resp := contracts.Response{
			Code:    400,
			Message: err.Error(),
		}
		return c.JSON(resp)
	}
	return c.JSON(domains)
}
