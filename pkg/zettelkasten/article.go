package zettelkasten

type Article struct {
	Title, FileName, Content string
	Number int
	Tags []*Tag `json:"-"`
	Connections []*Article `json:"-"`
}

type articleClear struct {
	Title, FileName, Content string
	ArticlePaths []string
	TagPaths     []string
}

func (a *Article)Index() string {
	return a.FileName
}
