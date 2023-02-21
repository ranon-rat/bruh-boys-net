package src

const (
	POST        = 1
	LISTOFPOSTS = 2
)

var password = "123" // change it when you done
var host = "localhost:8080/ws"

type Posts struct {
	Date int64  `json:"-"`
	Name string `json:"name"`
}
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
