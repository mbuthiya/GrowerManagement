package main

import(
	"net/http"
	"database/sql"
	_"github.com/lib/pq"
	"github.com/gorilla/mux"
	
)

var Db *sql.DB

// Login Handler
func login(w http.ResponseWriter, r *http.Request){}

func init(){
	var err error

	Db,err = sql.Open("postgres","user=growManage dbname=growManage password=fruitylupoooo1 sslmode=disable")
	if err != nil{
		panic(err)
	}

}