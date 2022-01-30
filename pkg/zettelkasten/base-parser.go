package zettelkasten

import (
	"os"
	"path/filepath"
	"strings"
)

var articlesDir, tagsDir = "base", "tags"

func (b *base) ParseFromDisk(pathToRoot string) error {
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
	tagsClear, err := parseTagFiles(absPath)
	if err != nil {
		return err
	}
	articlesClear, err := parseArticleFiles(absPath)
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

func parseTagFiles(absPath string) ([]tagClear, error) {
	tagPath := filepath.Join(absPath, tagsDir)
	tagsDirContent, err := os.ReadDir(tagPath)
	if err != nil {
		return nil, err
	}
	tagsContent := make([]tagClear, 0)
	articlePathSubstring := "(../base/"
	for _, file := range tagsDirContent {
		if !file.IsDir() {
			data, err := os.ReadFile(filepath.Join(filepath.Join(tagPath, file.Name())))
			if err != nil {
				return nil, err
			}
			strData := string(data)
			fileLines := strings.Split(strData, "\n")
			if len(fileLines) == 0 {
				return nil, ErrTagFileNotCorrect
			}
			if strings.HasPrefix(fileLines[0], "# ") {
				tagContent := tagClear{
					Title:        fileLines[0][2:],
					FileName:     file.Name(),
					ArticlePaths: make([]string, 0),
				}
				for _, line := range fileLines {
					index := strings.Index(line, articlePathSubstring)
					if index != -1 {
						tagContent.ArticlePaths = append(tagContent.ArticlePaths,
							strings.ReplaceAll(line[index+len(articlePathSubstring):len(line)-1], "/", ""))
					}
				}
				tagsContent = append(tagsContent, tagContent)
			} else {
				return nil, ErrTagFileNotCorrect
			}
		}
	}
	return tagsContent, nil
}

func parseArticleFiles(absPath string) ([]articleClear, error) {
	articlePath := filepath.Join(absPath, articlesDir)
	articleDirContent, err := os.ReadDir(articlePath)
	tagsPathSubstring := "(../../tags/"
	connPathSubstring := "(../"
	articlesContent := make([]articleClear, 0)
	if err != nil {
		return nil, err
	}
	for _, dir := range articleDirContent {
		if dir.IsDir() {
			articleBt, err := os.ReadFile(filepath.Join(articlePath, dir.Name(), "README.md"))
			if err != nil {
				return nil, err
			}
			articleStr := string(articleBt)
			lines := strings.Split(articleStr, "\n")
			if len(lines) == 0 {
				return nil, ErrArticleFileNotCorrect
			}
			if strings.HasPrefix(lines[0], "# ") {
				articleContent := articleClear{
					Title:        lines[0][2:],
					FileName:     dir.Name(),
					ArticlePaths: make([]string, 0),
					TagPaths:     make([]string, 0),
				}
				iterStat := map[string]bool{}
				for lineIndex, line := range lines {
					if strings.HasPrefix(line, "## ") {
						if strings.HasPrefix(line[3:], "Теги") {
							iterStat["isTagSection"] = true
							iterStat["isConnectionSection"] = false
							continue
						} else if strings.HasPrefix(line[3:], "Связи") {
							iterStat["isTagSection"] = false
							iterStat["isConnectionSection"] = true
							continue
						} else {
							articleContent.Content = strings.Join(lines[lineIndex:], "\n")
							break
						}
					}
					if iterStat["isTagSection"] {
						index := strings.Index(line, tagsPathSubstring)
						if index != -1 {
							articleContent.TagPaths = append(articleContent.TagPaths,
								line[index+len(tagsPathSubstring):len(line)-1])
						}
					}
					if iterStat["isConnectionSection"] {
						index := strings.Index(line, connPathSubstring)
						if index != -1 {
							articleContent.ArticlePaths = append(articleContent.ArticlePaths,
								line[index+len(connPathSubstring):len(line)-1])
						}
					}
				}
				articlesContent = append(articlesContent, articleContent)
			} else {
				return nil, ErrArticleFileNotCorrect
			}
		}
	}
	return articlesContent, nil
}
