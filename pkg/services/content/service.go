package content

import (
	"context"
	"github.com/krls256/knowlage-base-editor/pkg/github"
	"github.com/krls256/knowlage-base-editor/pkg/zettelkasten"
)

type service struct {
	base *zettelkasten.Base
}

type Config struct {
	GitHubLogin, GitHubAuthToken, GitHubRepoName, PathToStoreRepo string
}

func New(conf Config) (*service, error)  {
	ctx := context.Background()
	user := &github.User{Login: conf.GitHubLogin, AuthToken: conf.GitHubAuthToken}
	repo := &github.Repository{Name: conf.GitHubRepoName}
	err := github.WriteRepoToDisk(ctx, user, repo, conf.PathToStoreRepo)
	if err != nil {
		return nil, err
	}
	base := zettelkasten.NewBase()
	base.ParseFromDisk(conf.PathToStoreRepo)
	return &service{base: base}, nil
}

func (s *service)GetArticles() []*zettelkasten.Article {
	return s.base.Articles()
}

func (s *service)GetTags() []*zettelkasten.Tag {
	return s.base.Tags()
}