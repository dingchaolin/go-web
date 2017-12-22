package main

import (
	"io"
	"log"
	"net/http"
	"os"
)
// 静态文件必须使用绝对路径
func main(){

	mux := http.NewServeMux()

	mux.Handle( "/", &myHandler{} )
	mux.HandleFunc( "/hello", sayHello1 )

	wd, err := os.Getwd() //当前工作目录
	if err != nil{
		log.Fatal(err)
	}

	//http.StripPrefix("/static/" 干掉前缀 可能不需要
	//静态文件服务
	mux.Handle( "/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(wd))))

	err = http.ListenAndServe(":8080", mux )
	if err != nil {
		log.Fatal( err )
	}

}

type myHandler struct{}

func (*myHandler) ServeHTTP( w http.ResponseWriter, r *http.Request){
	io.WriteString( w, "URL: " + r.URL.String() )
}

func sayHello1( w http.ResponseWriter, r *http.Request ){
	io.WriteString( w, "hello world, this v 2")
}