package zettelkasten

import "encoding/json"

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

func (a *Article)MarshalJSON() ([]byte, error) {
	tagsIds := make([]int, 0, len(a.Tags))
	connIds := make([]int, 0, len(a.Connections))
	for _, tag := range a.Tags {
		tagsIds = append(tagsIds, tag.TempId)
	}
	for _, conn := range a.Connections {
		connIds = append(connIds, conn.Number)
	}
	return json.Marshal(struct {
		Title string `json:"title"`
		FileName string `json:"file_name"`
		Content string `json:"content"`
		Number int `json:"number"`
		TempId int `json:"temp_id"`
		Tags []int `json:"tags"`
		Connections []int `json:"connections"`
	}{
		Title: a.Title,
		FileName: a.FileName,
		Content: a.Content,
		Number: a.Number,
		TempId: a.Number,
		Tags: tagsIds,
		Connections: connIds,
	})
}