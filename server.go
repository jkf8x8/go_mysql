package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter , r *http.Request){
	fmt.Fprintf(w,"hello")
}

func main(){
	
	mux := http.NewServeMux()
	mux.HandleFunc("/",index)
	
	server :=http.Server{
		Addr:":8080",
		Handler:mux,
	}
	server.ListenAndServe()
	fmt.Println("ok")
}