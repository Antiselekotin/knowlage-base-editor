package zettelkasten

import (
	"strconv"
	"strings"
)

type Base struct {
	articles map[string]*Article
	tags map[string]*Tag
}

func NewBase() *Base {
	return &Base{
		tags: make(map[string]*Tag),
		articles: make(map[string]*Article),
	}
}

func (b *Base)AddTag(tag tagClear) {
	newTag := Tag{Title: tag.Title, FileName: tag.FileName, Articles: make([]*Article, 0)}
	for _, articlePath := range tag.ArticlePaths {
		article, ok := b.articles[articlePath]
		if ok {
			newTag.Articles = append(newTag.Articles, article)
			article.Tags = append(article.Tags, &newTag)
		}
	}
	b.tags[newTag.Index()] = &newTag
}

func (b *Base)AddArticle(article articleClear) {
	fName := article.FileName
	num := 0
	numberStr := strings.Split(fName, "-")
	if len(numberStr) != 0 {
		tmp, err := strconv.Atoi(numberStr[0])
		if err == nil {
			num = tmp
		}
	}
	newArticle := Article{Title: article.Title, FileName: fName, Content: article.Content, Number: num,
		Tags: make([]*Tag, 0), Connections: make([]*Article, 0)}
		for _, articlePath := range article.ArticlePaths {
			article, ok := b.articles[articlePath]
			if ok {
				newArticle.Connections = append(newArticle.Connections, article)
				article.Connections = append(article.Connections, &newArticle)
			}
		}
	for _, tagPath := range article.TagPaths {
		tag, ok := b.tags[tagPath]
		if ok {
			newArticle.Tags = append(newArticle.Tags, tag)
			tag.Articles = append(tag.Articles , &newArticle)
		}
	}
	b.articles[newArticle.Index()] = &newArticle
}

func (b *Base)Tags() []*Tag {
	tags := make([]*Tag, len(b.tags))
	i := 0
	for _, tag := range b.tags {
		tags[i] = tag
		i++
	}
	return tags
}

func (b *Base)Articles() []*Article {
	articles := make([]*Article, len(b.articles))
	i := 0
	for _, article := range b.articles {
		articles[i] = article
		i++
	}
	return articles
}