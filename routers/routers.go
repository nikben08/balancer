package routers

import (
	"balancer/handlers"

	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	api := app.Group("")
	auth := api.Group("/auth")
	auth.Get("/login", handlers.LoginView)
	auth.Post("/login", handlers.Login)

	dashboard := api.Group("/dashboard")
	dashboard.Get("/peers_manager", handlers.PeersManager)
	dashboard.Post("/peers_manager", handlers.PeersManagerDomain)
	dashboard.Post("/create_server", handlers.CreateServer)
	dashboard.Post("/delete_server", handlers.DeleteServer)
	dashboard.Get("/get_all_servers", handlers.GetAllServers)

	dashboard.Get("/domains", handlers.Domains)
	dashboard.Post("/create_domain", handlers.CreateDomain)
	dashboard.Post("/delete_domain", handlers.DeleteDomain)
	dashboard.Get("/get_all_domains", handlers.GetAllDomains)

}
