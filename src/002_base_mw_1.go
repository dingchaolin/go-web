package main

import (
	"net/http"
)
type SingleHost struct{
	handler http.Handler
	allowedHost string
}

func (this *SingleHost)ServeHTTP( w http.ResponseWriter, r *http.Request){
	println(r.Host)
	if r.Host == this.allowedHost{
		this.handler.ServeHTTP( w, r)
	}else{
		w.WriteHeader(403)
	}
}

func MyHandler( w http.ResponseWriter, r *http.Request ){
	w.Write([]byte("hello world"))
}

func main(){
	single := &SingleHost{
		handler:http.HandlerFunc( MyHandler ),
		allowedHost:"localhost:8080",
	}
	http.ListenAndServe( ":8080", single )
}