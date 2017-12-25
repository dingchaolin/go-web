package main

import (
	"net/http"
	"net/http/httptest"
)

type ModifierMiddleware struct{
	handler http.Handler
}

func ( this *ModifierMiddleware) ServeHTTP( w http.ResponseWriter, r *http.Request ){
	rec := httptest.NewRecorder()

	this.handler.ServeHTTP( rec, r )

	for k, v := range rec.Header(){
		w.Header()[k] = v
	}

	w.Header().Set("go-web-function", "vip")
	w.WriteHeader( 418 )
	w.Write([]byte("hey, this is middleware"))
	w.Write(rec.Body.Bytes())
}

func myHandler4( w http.ResponseWriter, r *http.Request ){
	w.Write([]byte("hello world"))
}

func main(){
	mid := &ModifierMiddleware{http.HandlerFunc( myHandler4)}

	http.ListenAndServe( ":8080", mid )
}