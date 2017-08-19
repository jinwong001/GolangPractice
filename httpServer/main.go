package main

import (
	"net/http"
	"time"
	"os"
	"os/signal"
	"log"
	"fmt"
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
	////Form：存储了post、put和get参数，在使用之前需要调用ParseForm方法。并且以post参数为先
	////http.request在解析参数的时候会将同名的参数都放进同一个数组里，所以这里要用[0]获取到第一个。
	////PostForm：存储了post、put参数，在使用之前需要调用ParseForm方法。
	////MultipartForm：存储了包含了文件上传的表单的post参数，在使用前需要调用ParseMultipartForm方法。

	//
	////r.ParseForm()
	////if len(r.Form["id"]) > 0 {
	////	fmt.Fprintln(w, r.Form["id"][0])
	////}
	//
	////所以表单enctype要设置成multipart/form-data。此时无法通过PostFormValue来获取id的值
	//
	//r.PostForm()
	//
	//r.PostFormValue("Id")
	//
	//
	//
	//queryForm, err := url.ParseQuery(r.URL.RawQuery)
	//if err == nil && len(queryForm["id"]) > 0 {
	//	fmt.Fprintln(w, queryForm["id"][0])
	//}
	//
	//
	//
	//r.ParseForm()
	//r.PostForm("Id")
	//
	//
	//
	//r.ParseMultipartForm(32 << 20)
	//if r.MultipartForm != nil {
	//	values := r.MultipartForm.Value["id"]
	//	if len(values) > 0 {
	//		fmt.Fprintln(w, values[0])
	//	}
	//}



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




