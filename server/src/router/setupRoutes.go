package router

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ranon-rat/bruh-boys-net/server/src/controller"
)

func SetupRoutes() {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.HandleFunc("/ws", controller.ConnectWS)
	http.HandleFunc("/post", controller.PostTemplate)
	http.HandleFunc("/", controller.HomeTemplate)

	go controller.SendInfoToEveryone()
	go controller.CheckThePosts()
	port, _ := os.LookupEnv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("service avaible in http://localhost:" + port)
	log.Println(http.ListenAndServe(":"+port, nil))
}
