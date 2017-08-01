package main

import (
	"net/http"
	"log"
	"net/http/httptest"
)

type SingleHost struct {
	handler   http.Handler
	allowHost string
}

func (this SingleHost)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rec:=httptest.NewRecorder()
	this.handler.ServeHTTP(w,r)
	for k,v := range rec.Header(){
		w.Header()[k]=v
	}

	w.Header().Set("go web","vip")
	w.WriteHeader(418)
	w.Write(rec.Body.Bytes())
	w.Write([]byte("hey this middler\\"))
}

func singleHost(handler http.Handler, allowHost string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Host)
		if allowHost == r.Host {
			handler.ServeHTTP(w, r)
		} else {
			w.WriteHeader(403)
		}
	}
	return http.HandlerFunc(fn)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

type AppendMiddleware struct {
	handler http.Handler
}

func (this *AppendMiddleware)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	this.handler.ServeHTTP(w, r)
	w.Write([]byte("Hey is append middleware"))
}

func main() {
	singleHost := SingleHost{
		handler:http.HandlerFunc(myHandler),
		allowHost:"localhost:1234",
	}

	http.ListenAndServe(":1234",singleHost)
	//http.ListenAndServe(":1234", singleHost(http.HandlerFunc(myHandler),"localhost:123e4"))

	//http.ListenAndServe(":1234", &AppendMiddleware{http.HandlerFunc(myHandler)})
	//http.ListenAndServe(":1234",MiddleWare(http.HandlerFunc(myHandler)))

}

func MiddleWare(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec:=httptest.NewRecorder()
		handler.ServeHTTP(w,r)
		for k,v := range rec.Header(){
			w.Header()[k]=v
		}

		w.Header().Set("go web","vip")
		w.WriteHeader(418)
		w.Write(rec.Body.Bytes())
		w.Write([]byte("hey this middler\\"))
	})
}
