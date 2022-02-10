package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/krls256/knowlage-base-editor/pkg/services/content"
	"log"
	"os"
)

var ghRepo, ghLogin, ghToken string

func init()  {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ghLogin = os.Getenv("GITHUB_LOGIN")
	ghToken = os.Getenv("GITHUB_TOKEN")
	ghRepo = os.Getenv("GITHUB_REPO")
}

func main()  {
	contentService, err := content.New(content.Config{GitHubLogin: ghLogin, GitHubAuthToken: ghToken, GitHubRepoName: ghRepo, PathToStoreRepo: "storage/repo"})
	if err != nil {
		panic(err)
	}
	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		data, err := json.Marshal(contentService.GetArticles())
		if err != nil {
			ctx.SendString(err.Error())
			return ctx.SendStatus(500)
		}
		return ctx.Send(data)
	})
	api := app.Group("/api", func(c *fiber.Ctx) error {
		c.Set("Content-type", "application/json; charset=utf-8")
		return c.Next()
	})
	api.Get("/articles", func(ctx *fiber.Ctx) error {
		body, err := json.Marshal(contentService.GetArticles())
		if err != nil {
			return ctx.SendStatus(500)
		}
		return ctx.Send(body)
	})
	api.Get("/tags", func(ctx *fiber.Ctx) error {
		body, err := json.Marshal(contentService.GetTags())
		if err != nil {
			return ctx.SendStatus(500)
		}
		return ctx.Send(body)
	})
	err = app.Listen(":90")
	if err != nil {
		panic(err)
	}
}
