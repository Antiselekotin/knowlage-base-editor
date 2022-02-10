package zettelkasten

import "encoding/json"

type Tag struct {
	TempId int
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

func (t *Tag)MarshalJSON() ([]byte, error) {
	articlesIds := make([]int, 0, len(t.Articles))
	for _, tag := range t.Articles {
		articlesIds = append(articlesIds, tag.Number)
	}
	return json.Marshal(struct {
		Title string `json:"title"`
		FileName string `json:"file_name"`
		TempId int `json:"temp_id"`
		Articles []int `json:"articles"`
	}{
		Title: t.Title,
		FileName: t.FileName,
		TempId: t.TempId,
		Articles: articlesIds,
	})
}