package main

import (
	"fmt"
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
	fmt.Println(contentService.GetArticles())
}
