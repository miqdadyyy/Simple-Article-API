package main

import (
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(app *fiber.App) {
	h := InitHandler()

	app.Get("/article", h.HandleGetArticles)
	app.Post("/article", h.HandleCreateArticle)
	app.Get("/article/:slug", h.HandleGetArticleDetail)
	app.Post("/article/:slug/comment", h.HandleCreateArticleComment)
}
