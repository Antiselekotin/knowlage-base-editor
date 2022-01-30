package zettelkasten

type Tag struct {
	Title, FileName string
	Articles []*Article
}

type tagClear struct {
	Title, FileName string
	ArticlePaths []string
}