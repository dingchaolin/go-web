package main

import (
	"net/http"
)
type AppendMiddlerware struct{
	handler http.Handler
}




func (this *AppendMiddlerware ) ServeHTTP( w http.ResponseWriter, r *http.Request){
	this.handler.ServeHTTP( w, r )
	w.Write([]byte("hey, this is middleware"))

}

func myHandler2( w http.ResponseWriter, r *http.Request ){
	w.Write([]byte("hello world"))
}

func main(){
	mid := &AppendMiddlerware{http.HandlerFunc(myHandler2)}
	http.ListenAndServe( ":8080", mid )
}