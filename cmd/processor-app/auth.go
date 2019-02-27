package main

import(
	"log"
	"net/http"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	_"github.com/lib/pq"
	"github.com/satori/go.uuid"
	"time"
	
)

var Db *sql.DB

func loginWithEmail(w http.ResponseWriter, usr Users) (error){

	// Query the database
	userQueryResults:= Db.QueryRow("select id,password from users where email=$1",usr.email)

	// New instance for users
	storedUserCred := &Users{}
	
	err := userQueryResults.Scan(&storedUserCred.id,&storedUserCred.password)
	if err != nil {
		return err
	}

	// Compare hash and password
	err = bcrypt.CompareHashAndPassword([]byte(storedUserCred.password),[]byte(usr.password))
	if err != nil {
		return err
	}
	
	// Setting up redis session
	sessionToken := uuid.Must(uuid.NewV4()).String()
	
	_,err =cache.Do("SETEX",sessionToken,"86400",storedUserCred.id)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(86400 * time.Second),
	})

	return nil
}

func signUpWithEmail(w http.ResponseWriter, usr Users) error{

	// Hashing the password and second argument is the cost of hashing
	hashedPassword,err:= bcrypt.GenerateFromPassword([]byte(usr.password),8)

	// Insert to database
	_,err =Db.Query("insert into Users (email,password) values($1,$2)",usr.email,string(hashedPassword))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error signupHandler:",err)
		return err
	}

	return nil
}