package main

import (
	"net/http"
	"time"
	"os"
	"os/signal"
	"log"
)

func main1() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	http.HandleFunc("/say", say)
	log.Println("server is starting ")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func say(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("say hi"))
	http.Error(w,"error",http.StatusNotFound)
}

func main() {
	server := http.Server{
		WriteTimeout: 2 * time.Second,
		Addr:":8080",
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/say", say)
	mux.Handle("/", &myHandler{})
	server.Handler = mux

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
		if err := server.Close(); err != nil {
			log.Fatal("Close server:", err)
		}

	}()

	log.Println("server is starting...")
	err := server.ListenAndServe()
	if err == http.ErrServerClosed {
		log.Println("server is closed under request")
	} else if (err != nil) {
		log.Println("server is closed unexpected")
	}

}

type myHandler struct{}

func (handler *myHandler)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi,this my handler"))
}




