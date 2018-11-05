package main

import (
	"io/ioutil"
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter , r *http.Request ,_ httprouter.Params){
	fmt.Println(r.ContentLength)
	getData := make([]byte,r.ContentLength)
	r.Body.Read(getData)
	fmt.Fprint(w , getData)
}

//接收表单数据
func indexPost(w http.ResponseWriter , r *http.Request , _ httprouter.Params){
	//解析 enctype="application/x-www-form-urlencoded"
	// r.ParseForm()
	// fmt.Fprint(w , r.Form)

	//解析 enctype="multipart/form-data"
	// r.ParseMultipartForm(1024)
	// fmt.Fprint(w , r.MultipartForm)

	//接收指定 键值对 r.FormValue  r.PostFormValue
	// fmt.Fprint(w,r.FormValue("gugu"))
	// fmt.Fprint(w,r.PostFormValue("gugu"))
	// fmt.Fprint(w,r.PostForm)

	// r.Form  r.PostForm

}

//文件上传
func fileUp(w http.ResponseWriter , r *http.Request ,_ httprouter.Params){
	// r.ParseMultipartForm(1024)
	// fileHeader := r.MultipartForm.File["pop"][0]
	// file,err := fileHeader.Open()
	// if err ==nil{
	// 	data,err :=ioutil.ReadAll(file)
	// 	if err ==nil{
	// 		fmt.Fprintln(w,string(data))
	// 	}
	// }

	//FormFile
	file,_,err := r.FormFile("pop")
	if err == nil{
		data,err := ioutil.ReadAll(file)
		if err==nil{
			fmt.Fprintln(w , string(data))
		}
	}
}

func main(){

	router := httprouter.New()

	router.GET("/",index)
	// router.POST("/",indexPost)
	//文件上传
	router.POST("/",fileUp)

	server := http.Server{
		Addr:":8080",
		Handler:router,
	}

	server.ListenAndServe()
}