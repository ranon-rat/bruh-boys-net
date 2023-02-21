package src

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"golang.org/x/net/websocket"
)

func ConnectWS() {
	origin, _ := url.Parse("http://" + host)
	u, _ := url.Parse("ws://" + host)
	conn, err := websocket.DialConfig(&websocket.Config{
		Origin:    origin,
		Location:  u,
		Version:   websocket.ProtocolVersionHybi13,
		TlsConfig: nil,
		Header: http.Header{
			"pass": {password}, // later i will use a token btw
		},
	})
	if err != nil {
		panic(err)
	}
	CommunicateToServer(conn)

}

func CommunicateToServer(conn *websocket.Conn) {

	for {
		var info Protocol
		if err := json.NewDecoder(conn).Decode(&info); err != nil {
			panic(err)
		}
		fmt.Println(info)
		switch info.Kind {
		case POST:
			content, err := GetPost(info.Name)
			json.NewEncoder(conn).Encode(Response{
				ID:    info.ID,
				Kind:  info.Kind,
				Error: err != nil,
				List:  nil,
				Post:  content})
		case LISTOFPOSTS:
			posts := ListOfPosts()
			json.NewEncoder(conn).Encode(Response{
				ID:    0,
				Kind:  LISTOFPOSTS,
				Error: false,
				List:  posts,
				Post:  ""})
		}
	}
}
