package main

import(
	
	"log"
	"net/http"
	"html/template"
	"os"
	"database/sql"
	_"github.com/lib/pq"
	"github.com/gorilla/mux"
	"github.com/gomodule/redigo/redis"
	
)



var BaseURL string = "http://localhost:8090"
var cache redis.Conn

type Users struct{
	id int
	email string
	password string
}


// Login Handler
func loginHandler(w http.ResponseWriter, r *http.Request){
	usr := &Users{}
	baseTemplates,_:= os.Getwd()
	
	usr.email = r.PostFormValue("email")
	usr.password =r.PostFormValue("password")

	_,err:=r.Cookie("session_token")
	
	if err ==nil{
		http.Redirect(w,r,BaseURL+"/welcome",http.StatusSeeOther)
		return
	}

	if usr.email =="" || usr.password == ""{

		t,_ := template.ParseFiles(baseTemplates+"/templates/login.html")
		t.Execute(w,nil)

	}else{

		err:=loginWithEmail(w,*usr)
		if err != nil {
			log.Println("Error")
		}
		http.Redirect(w,r,BaseURL+"/welcome",http.StatusSeeOther)
	}





}

// Signup Handler
func signupHandler(w http.ResponseWriter, r *http.Request){

	usr := &Users{}
	baseTemplates,_:= os.Getwd()
	
	usr.email = r.PostFormValue("email")
	usr.password =r.PostFormValue("password")

	_,err:=r.Cookie("session_token")
	
	if err ==nil{
		http.Redirect(w,r,BaseURL+"/welcome",http.StatusSeeOther)
		return
	}

	if usr.email =="" || usr.password == ""{

		t,_ := template.ParseFiles(baseTemplates+"/templates/signup.html")
		t.Execute(w,nil)

	}else{

		err:=signUpWithEmail(w,*usr)
		if err != nil {
			log.Println("Error")
		}
		http.Redirect(w,r,BaseURL+"/welcome",http.StatusSeeOther)
	}
	
	
}

func welcomePageHandler(w http.ResponseWriter, r *http.Request){
	baseTemplates,_:= os.Getwd()
	t,_ := template.ParseFiles(baseTemplates+"/templates/welcome.html")
	t.Execute(w,nil)
}

func init(){
	var err error

	Db,err = sql.Open("postgres","user=growManage dbname=growManage password=fruitylupoooo1 sslmode=disable")
	if err != nil{
		log.Fatal("Error Creating Database Connection",err)
	}

	// Initialize reddis connection 
	conn, err := redis.DialURL("redis://localhost")
	if err != nil {
		log.Fatal("Redis Connection Error: ",err)
	}
	// Assign the connection to the package level `cache` variable
	cache = conn

}

func main(){

	router := mux.NewRouter()

	// Handling routes with mux
	router.HandleFunc("/login",loginHandler)
	router.HandleFunc("/signup",signupHandler)
	router.HandleFunc("/welcome",welcomePageHandler)

	server:= http.Server{
		Addr:"127.0.0.1:8090",
		Handler:router,
	}

	err:=server.ListenAndServe()
	
	if err != nil {
		log.Fatal("Error starting up server",err)
	}
}