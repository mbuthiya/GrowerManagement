package main

import(
	"log"
	"net/http"
	"database/sql"
	_"github.com/lib/pq"
	"github.com/gorilla/mux"
	
)

var Db *sql.DB

// Login Handler
func loginHandler(w http.ResponseWriter, r *http.Request){}

// Signup Handler
func signupHandler(w http.ResponseWriter, r *http.Request){}


func init(){
	var err error

	Db,err = sql.Open("postgres","user=growManage dbname=growManage password=fruitylupoooo1 sslmode=disable")
	if err != nil{
		log.Fatal("Error Creating Database Connection",err)
	}

}

func main(){

	router := mux.NewRouter()

	// Handling routes with mux
	router.HandleFunc("/login",loginHandler)
	router.HandleFunc("/signup",signupHandler)

	server:= http.Server{
		Addr:"127.0.0.1:8080",
		Handler:router,
	}

	err:=server.ListenAndServe()
	
	if err != nil {
		log.Fatal("Error starting up server",err)
	}
}