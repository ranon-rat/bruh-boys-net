package controller

import "golang.org/x/net/websocket"

type Protocol struct {
	ID   int    `json:"id"`
	Kind int    `json:"kind"` //1=list of posts |2=Post
	Name string `json:"name"`
}
type Response struct {
	ID    int      `json:"id"`
	Kind  int      `json:"kind"` //1=list of posts |2=Post
	Error bool     `json:"error"`
	List  []string `json:"listOfPosts"`
	Post  string   `json:"content"`
}

const (
	POST = 1
	LIST = 2
)

var (
	password    = "123" // change this, dont be retarded
	GiveMe      = make(chan Protocol)
	peers       = make(map[*websocket.Conn]bool)
	postAvaible []string
	Posts       = make(map[string]string)
	postsChan   = make(map[int]chan string)
)
