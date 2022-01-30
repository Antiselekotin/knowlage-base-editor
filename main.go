package main

import (
	"context"
	"fmt"
	"github.com/krls256/knowlage-base-editor/pkg/github"
	"github.com/joho/godotenv"
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
	ctx := context.Background()
	user := &github.User{Login: ghLogin, AuthToken: ghToken}
	repo := &github.Repository{Name: ghRepo}
	err := github.WriteRepoToDisk(ctx, user, repo, "storage/repo")
	fmt.Println(err)
}
