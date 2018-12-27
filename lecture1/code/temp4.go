package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

// 用map字典做路由分发
var mux map[string]func(w http.ResponseWriter, r *http.Request)

func main() {
	server := http.Server{
		Addr:        ":8004",
		Handler:     &myHandler2{},
		ReadTimeout: 5 * time.Second,
	}
	mux = make(map[string]func(w http.ResponseWriter, r *http.Request))
	mux["hello"] = sayHello
	mux["bye"] = sayBye
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

type myHandler2 struct{}

func (this *myHandler2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}
	io.WriteString(w, "URL: "+r.URL.String())
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World! This is version3")
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "BYE BYE This is version3")
}