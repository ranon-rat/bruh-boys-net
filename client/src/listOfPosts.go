package src

import (
	"os"
	"sort"
)

func ListOfPosts() (posts []string) {
	var publications []Posts
	a, _ := os.ReadDir("./posts")
	for _, f := range a {
		i, _ := f.Info()
		publications = append(publications, Posts{i.ModTime().Unix(), i.Name()})
	}
	sort.Slice(publications, func(i, j int) bool {
		return publications[i].Date > publications[j].Date
	})
	for _, f := range publications {
		posts = append(posts, f.Name)

	}
	return
}
