package main

import (
	 "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"net/http"
	"html/template"
)

type Users struct{
	id int
	uuid int
	name string
	email string
	created_at string
}

func index( w http.ResponseWriter , request *http.Request){
	// fmt.Println(request.RequestURI)
	if request.RequestURI == "/favicon.ico" {
		return 
	}
	// _,err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/chsat")
	db,_:=sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/chat")
	defer db.Close()

	//删除数据
	// stmt,_ := db.Prepare("DELETE FROM users WHERE id>?")
	// res,_ := stmt.Exec( 2 )
	// aff,_ := res.RowsAffected()
	// defer stmt.Close()
	// fmt.Println(aff)


	//修改数据
	// stmt,_ := db.Prepare("UPDATE users SET name=? WHERE id>?")
	// res,_ := stmt.Exec("fei",2)
	// aff,_ := res.RowsAffected()
	// fmt.Println(aff)
	
	//插入数据
	// stmt,_ := db.Prepare(" INSERT INTO users(id,uuid) VALUES(?,?) ")
	// res,_ := stmt.Exec(4,4)
	// id,_ := res.LastInsertId()
	// fmt.Println(id)

	//查询数据
	rows,serr := db.Query("SELECT * FROM users")
	if serr != nil{
		fmt.Printf("查询错误...%s",serr)
	}
	// var users []Users
	for rows.Next(){
		var id int
		var uuid int
		var name string
		var email string
		var created_at string

		rows.Scan(&id,&uuid,&name,&email,&created_at)
		fmt.Printf("id:%d uuid:%d name:%s email:%s created_at:%s \n",id,uuid,name,email,created_at)
	}
	rows.Close() 

	// fmt.Println(users)
	
	
	temp := template.Must( template.ParseFiles("template/index.html","template/nav.html") )
	temp.ExecuteTemplate(w , "layout", nil)

}


func main(){

	mux := http.NewServeMux()

	mux.HandleFunc("/",index)

	server := http.Server{
		Addr:":8080",
		Handler:mux,
	}

	server.ListenAndServe() 

	fmt.Println(" ok ")


}