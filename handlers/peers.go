package handlers

import (
	"balancer/contracts"
	"balancer/models"
	"balancer/services"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func PeersManager(c *fiber.Ctx) error {
	domains, _ := services.GetAllDomains()
	domain_id := ""
	if len(domains) == 0 {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Please, create at least 1 domain first.",
			"error":   "No domains found",
		})
	}else{
		domain_id = domains[0].Id.String()
	}

	servers, _ := services.GetServersByDomain(domain_id)
	return c.Render("peers_manager", fiber.Map{"servers": servers, "domains":domains})
}

func PeersManagerDomain(c *fiber.Ctx) error {
	json := contracts.PeersManagerRequest{}
	if err := c.BodyParser(&json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
			"error":   err.Error(),
		})
	}

	domains, _ := services.GetAllDomains()
	
	for i := 0; i < len(domains); i++ {
		if domains[i].Id.String() == json.DomainId && i != 0{
			d := domains[0]
			domains[0] = domains[i]
			domains[i] = d
		}
	}


	servers, _ := services.GetServersByDomain(json.DomainId)
	return c.Render("peers_manager", fiber.Map{"servers": servers, "domains":domains})
}

func CreateServer(c *fiber.Ctx) error {
	json := contracts.CreateServerRequest{}
	if err := c.BodyParser(&json); err != nil {
		fmt.Println(err)
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
			"error":   err.Error(),
		})
	}

	server_weight, err := strconv.Atoi(json.ServerWeight)
	if err != nil {
		resp := contracts.Response{
			Code:    400,
			Message: "Incorrect server weight",
		}
		return c.JSON(resp)
	}

	server := models.Server{
		DomainId: json.DomainId,
		Address: json.ServerAddress,
		Label:   json.ServerLabel,
		Weight:  server_weight,
	}

	err = services.CreateServer(server)

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

func DeleteServer(c *fiber.Ctx) error {
	json := contracts.DeleteServerRequest{}
	if err := c.BodyParser(&json); err != nil {
		fmt.Println(err)
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
			"error":   err.Error(),
		})
	}

	serverId := json.ServerId

	err := services.DeleteServer(serverId)
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

func GetAllServers(c *fiber.Ctx) error {
	servers, err := services.GetAllServers()
	if err != nil {
		resp := contracts.Response{
			Code:    400,
			Message: err.Error(),
		}
		return c.JSON(resp)
	}
	return c.JSON(servers)
}
