package zettelkasten

import (
	"strconv"
	"strings"
)

type base struct {
	Articles map[string]*Article
	Tags map[string]*Tag
}

func (b *base)AddTag(tag tagClear) {
	newTag := Tag{Title: tag.Title, FileName: tag.FileName, Articles: make([]*Article, 0)}
	for _, articleTitle := range tag.ArticlePaths {
		article, ok := b.Articles[articleTitle]
		if ok {
			newTag.Articles = append(newTag.Articles, article)
			article.Tags = append(article.Tags, &newTag)
		}
	}
}

func (b *base)AddArticle(article articleClear) {
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
		for _, articleTitle := range article.ArticlePaths {
			article, ok := b.Articles[articleTitle]
			if ok {
				newArticle.Connections = append(newArticle.Connections, article)
				article.Connections = append(article.Connections, &newArticle)
			}
		}
	for _, tagTitle := range article.TagPaths {
		tag, ok := b.Tags[tagTitle]
		if ok {
			newArticle.Tags = append(newArticle.Tags, tag)
			tag.Articles = append(tag.Articles , &newArticle)
		}
	}
}

func NewBase() *base {
	return &base{
		Tags: make(map[string]*Tag),
		Articles: make(map[string]*Article),
	}
}
