package main

import (
	"./dbPkg"
	"./handlers"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"os"
)

const PORT = "3001"

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()

	dbPkg.New()
	handlers.InitializeJWTKeys(dir)
	r.HandleFunc("/", handlers.IndexHandler)
	r.HandleFunc("/auth/signup", handlers.SignUpHandler)
	r.HandleFunc("/auth/login", handlers.LoginHandler)
	r.HandleFunc("/auth/renew", handlers.RenewPasswordHandler)
	r.Handle("/auth/cookie-login", combineWithMiddleware(handlers.CookieLoginHandler))
	err = http.ListenAndServe(":"+PORT, r)
	if err != nil {
		fmt.Println("Error on listening port : " + PORT + err.Error())
	}
}
func combineWithMiddleware(handler http.HandlerFunc) *negroni.Negroni {
	return negroni.New(negroni.HandlerFunc(handlers.TokenVerifyMiddleware), negroni.WrapFunc(handler))
}
