package zettelkasten

type Tag struct {
	Title, FileName string
	Articles []*Article
}

type tagClear struct {
	Title, FileName string
	ArticlePaths []string
}

func (t *Tag)Index() string {
	return t.FileName
}