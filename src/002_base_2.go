package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

var mux map[string]func( w http.ResponseWriter, r *http.Request)
func main(){

	server := http.Server{
		Addr: ":8080",
		Handler : &myHandler1{},
		ReadTimeout: 5  * time.Second,
	}

	mux = make(map[string]func( w http.ResponseWriter, r *http.Request))
	mux["/hello"] = sayYs
	mux["/bye"] = sayBye

	err := server.ListenAndServe()

	if err!= nil {
		log.Fatal( err )
	}


}

type myHandler1 struct{}

func (*myHandler1) ServeHTTP( w http.ResponseWriter, r *http.Request){

	if h, ok := mux[r.URL.String()]; ok {
		h( w, r)
		return
	}

	io.WriteString( w, "URL: " + r.URL.String() )
}

func sayYs( w http.ResponseWriter, r *http.Request ){
	io.WriteString( w, "hello Ys, this v 3")
}

func sayBye( w http.ResponseWriter, r *http.Request ){
	io.WriteString( w, "bye~, this v 3")
}