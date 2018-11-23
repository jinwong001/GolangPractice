package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"log"
	"runtime"
)

/**
 *
 *   参考：href=https://www.jianshu.com/p/8ade70e51210
 *
 *
 *
 */
func main() {
	//TODO
	r := mux.NewRouter()
	r.HandleFunc("/home", homeHandler)
	r.HandleFunc("/panic", panicHandler)
	r.Use(recoverMiddleware)                     // mux 支持中间件
	r.Use(loggingMiddleware, loggingMiddleware1) // mux 支持中间件

	s := r.PathPrefix("/v1").Subrouter()
	s.HandleFunc("/home", homeHandlerV1)
	s.HandleFunc("/homepost", homeHandlerV1).Methods("POST","GET")

	http.Handle("/", r) // 不用再加规则，"/" 最好
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "this is home")
}

func homeHandlerV1(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "this is home:v1")
}

func panicHandler(w http.ResponseWriter, r *http.Request) {
	panic("test panic")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func loggingMiddleware1(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi, loggingMiddleware1\n"))
		next.ServeHTTP(w, r)
		w.Write([]byte("\nhi, loggingMiddleware1 end\n"))
	})
}

func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("painc:", err)
				printStack()
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("system error"))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func printStack() {
	buf := make([]byte, 1<<16)
	runtime.Stack(buf, false) // true 可以调用其他线程 stack traces,否则当前线程
	fmt.Printf("[start all stack]----------------  %s   ----------------[all stack end]", buf)
}
