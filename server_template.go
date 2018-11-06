package main

import (
	"fmt"
	"net/http"
	"html/template"
)

func index( w http.ResponseWriter , r *http.Request){

	//创建模板
	// tem := template.New("./template/index.html")
	//解析模板
	// tem.Parse("./template/index.html")

	// parseFiles 返回分析模板 和错误结果 相当于上面代码的简写方式
	// tem,_ := template.ParseFiles("./template/index.html")

	//template.Must 处理解析模板时出现的错误
	// tem := template.Must( template.ParseFiles("./template/index.html") )

	//执行模板
	// tem.Execute(w,"hello world")
	//解析多个模板时执行指定模板
	tem , _ := template.ParseFiles("./template/index.html","./template/index2.html")
	// fmt.Println(tem.Name())

	//模板语法
	// 迭代
	rangeData := []string{"a","b","c","d"}
	tem.ExecuteTemplate(w,"index2.html",rangeData)

	fmt.Fprint(w,"ok")
}


func main(){

	server := http.Server{
		Addr:":8080",
	}

	http.HandleFunc("/",index)

	server.ListenAndServe()
}