package zettelkasten

import (
	"os"
	"path/filepath"
	"strings"
)

var (
	articlesDir, tagsDir = "base", "tags"
	articlePathSubstring = "(../base/"
	tagsPathSubstring = "(../../tags/"
	connPathSubstring = "(../"
)

func (b *Base) ParseFromDisk(pathToRoot string) error {
	absPath, err := filepath.Abs(pathToRoot)
	if err != nil {
		return err
	}
	dirs, err := os.ReadDir(absPath)
	if err != nil {
		return err
	}
	dirMap := make(map[string]bool)
	for _, dir := range dirs {
		if dir.IsDir() {
			dirMap[dir.Name()] = true
		}
	}
	if !dirMap[articlesDir] || !dirMap[tagsDir] {
		return ErrNoReqDirs
	}
	tagsClear, err := parseAllTags(absPath)
	if err != nil {
		return err
	}
	articlesClear, err := parseAllArticles(absPath)
	if err != nil {
		return err
	}
	for _, tag := range tagsClear {
		b.AddTag(tag)
	}
	for _, article := range articlesClear {
		b.AddArticle(article)
	}
	return nil
}

func parseAllTags(absPath string) ([]tagClear, error) {
	tagPath := filepath.Join(absPath, tagsDir)
	tagsDirContent, err := os.ReadDir(tagPath)
	if err != nil {
		return nil, err
	}

	tagsContent := make([]tagClear, 0, len(tagsDirContent))
	for _, file := range tagsDirContent {
		tag, err := parseTagFile(file, tagPath)
		if err != nil {
			return nil, err
		}
		tagsContent = append(tagsContent, tag)
	}
	return tagsContent, nil
}

func parseAllArticles(absPath string) ([]articleClear, error) {
	articlePath := filepath.Join(absPath, articlesDir)
	articleDirContent, err := os.ReadDir(articlePath)
	if err != nil {
		return nil, err
	}

	articlesContent := make([]articleClear, 0, len(articleDirContent))
	for _, dir := range articleDirContent {
		article, err := parseArticleFile(dir, articlePath)
		if err != nil {
			return nil, err
		}
		articlesContent = append(articlesContent, article)
	}
	return articlesContent, nil
}

func parseTagFile(file os.DirEntry, tagPath string) (tagClear, error) {
	tagContent := tagClear{
		Title:        "",
		FileName:     file.Name(),
		ArticlePaths: make([]string, 0),
	}
	if !file.IsDir() {
		data, err := os.ReadFile(filepath.Join(filepath.Join(tagPath, file.Name())))
		if err != nil {
			return tagContent, err
		}
		lines := strings.Split(string(data), "\n")
		if len(lines) == 0 {
			return tagContent, ErrTagFileNotCorrect
		}
		if strings.HasPrefix(lines[0], "# ") {
			tagContent.Title = lines[0][2:]
			for _, line := range lines[1:] {
				if cut, ok := cutPathFromLine(line, articlePathSubstring); ok {
					tagContent.ArticlePaths = append(tagContent.ArticlePaths, cut)
				}
			}
			return tagContent, nil
		}
	}
	return tagContent, ErrTagFileNotCorrect
}

func parseArticleFile(dir os.DirEntry, articlePath string) (articleClear, error) {
	articleContent := articleClear{
		Title:        "",
		FileName:     dir.Name(),
		ArticlePaths: make([]string, 0),
		TagPaths:     make([]string, 0),
	}

	if dir.IsDir() {
		article, err := os.ReadFile(filepath.Join(articlePath, dir.Name(), "README.md"))
		if err != nil {
			return articleContent, err
		}
		lines := strings.Split(string(article), "\n")

		if len(lines) == 0 {
			return articleContent, ErrArticleFileNotCorrect
		}

		if strings.HasPrefix(lines[0], "# ") {
			articleContent.Title = lines[0][2:]
			iterStat := map[string]bool{}
			for lineIndex, line := range lines[1:] {

				if strings.HasPrefix(line, "## ") {
					if strings.HasPrefix(line[3:], "Теги") {
						iterStat["isTagSection"], iterStat["isConnectionSection"] = true, false
						continue
					} else if strings.HasPrefix(line[3:], "Связи") {
						iterStat["isTagSection"], iterStat["isConnectionSection"] = false, true
						continue
					} else {
						articleContent.Content = strings.Join(lines[lineIndex:], "\n")
						break
					}
				}
				if iterStat["isTagSection"] {
					if cut, ok := cutPathFromLine(line, tagsPathSubstring); ok {
						articleContent.TagPaths = append(articleContent.TagPaths, cut)
					}
				}
				if iterStat["isConnectionSection"] {
					if cut, ok := cutPathFromLine(line, connPathSubstring); ok {
						articleContent.ArticlePaths = append(articleContent.ArticlePaths, cut)
					}
				}
			}
			return articleContent, nil
		}
	}
	return articleContent, ErrArticleFileNotCorrect
}

func cutPathFromLine(line, subString string) (cut string, ok bool) {
	index := strings.Index(line, subString)
	if index != -1 {
		return strings.ReplaceAll(line[index+len(subString):len(line)-1], "/", ""), true
	}
	return
}