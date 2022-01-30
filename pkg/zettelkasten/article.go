package zettelkasten

type Article struct {
	Title, FileName, Content string
	Number int
	Tags []*Tag
	Connections []*Article
}

type articleClear struct {
	Title, FileName, Content string
	ArticlePaths []string
	TagPaths     []string
}