package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

func ConnectWS(w http.ResponseWriter, r *http.Request) {
	// I have to change this but you are not going to know the actual password if this is a close service
	// maybe i will generate a hash just for avoiding anything new but Im testing this first
	if r.Header.Get("pass") != password {
		http.Error(w, "you are not allowed to access to this service", http.StatusNotAcceptable)

		return
	}

	websocket.Handler(func(c *websocket.Conn) {
		peers[c] = true
		json.NewEncoder(c).Encode(Protocol{
			ID:   0,
			Kind: LIST,
			Name: "",
		})
		ReceiveInfo(c)
	}).ServeHTTP(w, r)
}
func SendInfoToEveryone() {
	for {

		info := <-GiveMe
		for c := range peers {
			json.NewEncoder(c).Encode(info)

		}
	}
}
func CheckThePosts() {
	for {
		for c := range peers {
			json.NewEncoder(c).Encode(Protocol{
				ID:   0,
				Kind: LIST,
				Name: "",
			})
		}
		time.Sleep(time.Minute * 10)
		postAvaible = []string{}
	}
}
func ReceiveInfo(c *websocket.Conn) {
	reader := json.NewDecoder(c)
	for {
		var resp Response
		if err := reader.Decode(&resp); err != nil {
			delete(peers, c)
			return
		}
		switch resp.Kind {
		case POST:
			if resp.Error {
				close(postsChan[resp.ID])
				delete(postsChan, resp.ID)
				continue
			}
			postsChan[resp.ID] <- resp.Post
		case LIST:

			postAvaible = append(postAvaible, resp.List...)
			fmt.Println(postAvaible, resp.List, resp)

		}

	}
}
