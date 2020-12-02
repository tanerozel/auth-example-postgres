package user

import (
	"app"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"templates"

	_ "github.com/lib/pq"
)

type dataview struct{
	Profile interface{} 
	Names []string 
}
 

func UserHandler(w http.ResponseWriter, r *http.Request) {
	
	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data:=&dataview{
		Profile: session.Values["profile"],
		Names :getUser(),
	}
	
 
	templates.RenderTemplate(w, "user", data)
}

func getUser() []string{
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	username := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DATABASE_NAME")
	connection_info := "host=" + host + " port=" + port + " user=" + username + " password=" + password + " dbname=" + dbname + " sslmode=disable"
	 fmt.Print(connection_info)
	db, err := sql.Open("postgres", connection_info)
	if err != nil {
		fmt.Println("Open Error")
	}

	rows, err := db.Query("SELECT * FROM Persons")
	if err != nil {
		fmt.Println(err.Error())
	}
	var data= make([]string,0)

	for rows.Next() {
		var name string
		_ = rows.Scan(&name)
		data=append(data, name)
		
	}
	if len(data) <1{
		
	if _, err := db.Exec("INSERT INTO Persons (name) VALUES ('FATIH SEVER'),('ÇAĞLAR SERT'),('BAHADIR CIVELEK'),('ATKAN GEMICI')"); err != nil {
		fmt.Println(err.Error())
	}
	}
	return data
}