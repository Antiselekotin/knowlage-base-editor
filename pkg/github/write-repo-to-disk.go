package github

import (
	"archive/zip"
	"bytes"
	"context"
	"errors"
	"github.com/google/go-github/v42/github"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func WriteRepoToDisk(ctx context.Context, user *User, repo *Repository, pathToCopy string) error {
	pathAbs, err := filepath.Abs(pathToCopy)
	if err != nil {
		return err
	}
	pathAbs += "/"
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: user.AuthToken})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	downloadLink, _, err := client.Repositories.GetArchiveLink(ctx, user.Login, repo.Name, github.Zipball, nil, true)
	if err != nil {
		return err
	}
	resp, err := http.Get(downloadLink.String())
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	rB := bytes.NewReader(respBody)
	reader, err := zip.NewReader(rB, rB.Size())
	if err != nil {
		return err
	}
	if len(reader.File) == 0 {
		return nil
	}
	first := reader.File[0]
	prefix := strings.Split(first.Name, "/")[0] + "/"
	isOk, created := createAllFiles(reader.File, prefix, pathAbs)
	if !isOk {
		for i := len(created) - 1; i >=0; i-- {
			_ = os.Remove(created[i])
		}
		return ErrUnZip
	}
	return nil
}

func isDir(str string) bool {
	if len(str) == 0 {
		return false
	}
	return str[len(str) - 1] == '/'
}

func createAllFiles(files []*zip.File, prefix, pathAbs string) (bool, []string) {
	isOk := true
	var created []string
	var err error
	for _, file := range files{
		p := strings.Replace(file.Name, prefix, pathAbs, -1)
		if isDir(p) {
			if _, err = os.Stat(p); errors.Is(err, os.ErrNotExist) {
				err = os.Mkdir(p, 0777)
				if err != nil {
					isOk = false
					break
				}
			}
			created = append(created, p)
		} else {
			read, err := file.Open()
			if err != nil {
				isOk = false
				break
			}
			fileContent, err := ioutil.ReadAll(read)
			if err != nil {
				isOk = false
				break
			}
			err = os.WriteFile(p, fileContent, 0666)
			if err != nil {
				isOk = false
				break
			}
			created = append(created, p)
		}
	}
	return isOk, created
}