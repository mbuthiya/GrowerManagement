package main

import(
	
	"log"
	"net/http"
	"database/sql"
	_"github.com/lib/pq"
	"github.com/gorilla/mux"
	
)



var BaseURL string = "http://localhost:8090"

type Users struct{
	id int
	email string
	password string
}



// Login Handler
func loginHandler(w http.ResponseWriter, r *http.Request){
	usr := &Users{}

	usr.email = r.PostFormValue("email")
	usr.password =r.PostFormValue("password")

	// Login user with form details
	err := loginWithEmail(w,*usr)
	if err!= nil{
		w.Write([]byte("Please log in"))
	}


}

// Signup Handler
func signupHandler(w http.ResponseWriter, r *http.Request){
	log.Println(r.RemoteAddr)
	usr := &Users{}

	usr.email = r.PostFormValue("email")
	usr.password =r.PostFormValue("password")

	// Signup user
	err := signUpWithEmail(w,*usr)
	if err!= nil{
		w.Write([]byte("Please log in"))
	}
	
	http.Redirect(w,r,BaseURL+"/login",http.StatusSeeOther)
	
	
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
		Addr:"127.0.0.1:8090",
		Handler:router,
	}

	err:=server.ListenAndServe()
	
	if err != nil {
		log.Fatal("Error starting up server",err)
	}
}