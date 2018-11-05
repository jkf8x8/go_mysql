package main

import (
	"fmt"
	"net/http"

)

func index(w http.ResponseWriter , r *http.Request){
	str := `<h2>hehehe</h2>`
	w.Header().Set("content-type","text/html")
	// w.WriteHeader(501)
	w.Write([]byte(str))

	fmt.Println("ok")
}

func main(){

	server := http.Server{
		Addr:":8080",
	}

	http.HandleFunc("/",index)

	server.ListenAndServe()

}