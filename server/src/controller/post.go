package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"text/template"
)

func tmp(content string, w http.ResponseWriter) {
	file, _ := os.ReadFile("templates/post.html")
	tmpl := template.Must(template.New("a").Parse(string(file)))

	if err := tmpl.Execute(w, map[string]string{
		"content": content,
	}); err != nil {
		fmt.Println(err)
	}

}
func PostTemplate(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	name := r.URL.Query().Get("name")
	content, exists := Posts[name]
	id := rand.Int()
	wg.Add(1)
	go func() {
		if !exists && len(peers) > 0 {
			postsChan[id] = make(chan string)
			GiveMe <- Protocol{
				ID:   id,
				Kind: POST,
				Name: name,
			}
		}
		wg.Done()
	}()
	wg.Wait()
	if _, e := postsChan[id]; e {
		content = <-postsChan[id]
		Posts[name] = content
		delete(postsChan, id)
	}
	tmp(content, w)
}
