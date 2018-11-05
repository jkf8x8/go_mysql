package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter , r *http.Request ,_ httprouter.Params){
	fmt.Fprint(w , "hello")
}

func main(){

	router := httprouter.New()


	router.GET("/",index)

	server := http.Server{
		Addr:":8080",
		Handler:router,
	}

	server.ListenAndServe()
}