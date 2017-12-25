package main

import (
	"net/http"
)
func SingleHost1( handler http.Handler, allowedHost string) http.Handler {
	fn := func( w http.ResponseWriter, r *http.Request ){
		println(r.Host)
		if r.Host == allowedHost{
			handler.ServeHTTP( w, r)
		}else{
			w.WriteHeader(403)
		}
	}

	return http.HandlerFunc(fn)
}


func MyHandler1( w http.ResponseWriter, r *http.Request ){
	w.Write([]byte("hello world"))
}

func main(){
	single := SingleHost1( http.HandlerFunc(MyHandler1),"dcl.com")
	http.ListenAndServe( ":8080", single )
}