package src

import (
	"errors"
	"os"

	"github.com/gomarkdown/markdown"
)

func GetPost(name string) (post string, err error) {
	postHash := make(map[string]bool)
	for _, v := range ListOfPosts() {
		postHash[v] = true
	}
	if !postHash[name] {
		err = errors.New("file is outside of posts")
		return
	}
	content, err := os.ReadFile("./posts/" + name)
	if err != nil {

		return
	}
	if name[len(name)-2:] == "md" {
		post = string(markdown.ToHTML(content, nil, nil))
		return
	}

	return

}
