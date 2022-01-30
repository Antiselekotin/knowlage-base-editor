package zettelkasten

import "errors"

var (
	ErrNoReqDirs = errors.New("in given path not enough required directories (" + articlesDir + " & " + tagsDir + ")")
	ErrTagFileNotCorrect = errors.New("tag file is not correct")
	ErrArticleFileNotCorrect = errors.New("article file is not correct")
)