package main

import (
	"fmt"
	"os"

	"github.com/ranon-rat/simple-blog/client/src"
)

func main() {
	fmt.Println(`
    ______            _       _                                  _    
    | ___ \          | |     | |                                | |   
    | |_/ /_ __ _   _| |__   | |__   ___  _   _ ___   _ __   ___| |_  
    | ___ \ '__| | | | '_ \  | '_ \ / _ \| | | / __| | '_ \ / _ \ __|       ____()()
    | |_/ / |  | |_| | | | | | |_) | (_) | |_| \__ \ | | | |  __/ |_       /      @@
    \____/|_|   \__,_|_| |_| |_.__/ \___/ \__, |___/ |_| |_|\___|\__| ~~~~~\_;m__m._>o
                                           __/ |                      
                                          |___/            
	A project made by bruh boys.


 `)
	src.ConnectWS(os.Args[1])
}
