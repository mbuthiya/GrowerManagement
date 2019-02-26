package main

import(
	
	"log"
	"net/http"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	_"github.com/lib/pq"
	"github.com/gorilla/mux"
	
)


var Db *sql.DB

type Users struct{
	id int
	email string
	password string
}

// Login Handler
func loginHandler(w http.ResponseWriter, r *http.Request){}

// Signup Handler
func signupHandler(w http.ResponseWriter, r *http.Request){
	
	usr := &Users{}

	usr.email = r.PostFormValue("email")
	usr.password =r.PostFormValue("password")

	// Hashing the password and second argument is the cost of hashing
	hashedPassword,err:= bcrypt.GenerateFromPassword([]byte(usr.password),8)

	// Insert to database
	_,err =Db.Query("insert into Users (email,password) values($1,$2)",usr.email,string(hashedPassword))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal("Error signupHandler:",err)
		return
	}

	w.Write([]byte("Success"))
	
	
}

func welcomePageHandler(w http.ResponseWriter, r *http.Request){}

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
		Addr:"127.0.0.1:8000",
		Handler:router,
	}

	err:=server.ListenAndServe()
	
	if err != nil {
		log.Fatal("Error starting up server",err)
	}
}