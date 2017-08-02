package main

import (
	"net/http"
	"log"
)

func main() {


	http.Handle("/", http.FileServer(http.Dir("fileServer/static")))
	log.Println("server is starting...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
