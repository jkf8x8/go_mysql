package main

import (
	"fmt"
	"net/http"
)

//~~~~~~~~~~~基础服务
// func index(w http.ResponseWriter , r *http.Request){
// 	fmt.Fprintf(w,"hello")
// }

// func main(){
	
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/",index)
	
// 	server :=http.Server{
// 		Addr:":8080",
// 		Handler:mux,
// 	}
// 	server.ListenAndServe()
// 	fmt.Println("ok")
// }


//~~~~~~~~~~~自定义处理器
type CustomHandler struct{}

func (cus *CustomHandler) ServeHTTP(w http.ResponseWriter , r *http.Request){
	fmt.Fprintf(w, "custome handle")
}

func main(){
	Custom := CustomHandler{}
	server := http.Server{
		Addr:":8080",
		Handler: &Custom,
	}

	server.ListenAndServe()
}