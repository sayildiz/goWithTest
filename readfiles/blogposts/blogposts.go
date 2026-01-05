package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

const (
	titleSeperator       = "Title: "
	descriptionSeperator = "Description: "
	tagsSeperator        = "Tags: "
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post

	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())

		if err != nil {
			return nil, err //todo: needs cllarification, should fail if one file fails
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	return Post{Title: readMetaLine(titleSeperator),
		Description: readMetaLine(descriptionSeperator),
		Tags:        getTags(readMetaLine(tagsSeperator)),
		//Tags:        strings.Split(readMetaLine(tagsSeperator), ", "), // propably easier
		Body: readBody(scanner),
	}, nil
}

func readBody(scanner *bufio.Scanner) string {
	//ignore a line
	scanner.Scan()
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n")
}

func getTags(tagString string) []string {
	var res []string
	for tag := range strings.SplitSeq(tagString, ",") {
		res = append(res, strings.TrimSpace(tag))
	}
	return res
}
