package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ranon-rat/simple-blog/client/src"
)

func main() {
	fmt.Println(src.ListOfPosts())
	json.NewEncoder(os.Stdout).Encode(src.ListOfPosts())
	src.ConnectWS()
}
