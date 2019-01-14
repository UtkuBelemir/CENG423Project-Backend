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
	r.Handle("/auth/profile", combineWithMiddleware(handlers.ProfilePutHandler)).Methods(http.MethodPut, http.MethodOptions)
	r.Handle("/auth/profile", combineWithMiddleware(handlers.ProfileGetHandler)).Methods(http.MethodGet, http.MethodOptions)
	r.Handle("/advertisement", combineWithMiddleware(handlers.SaveAdvertisementHandler)).Methods(http.MethodPost, http.MethodOptions)
	r.Handle("/advertisement", combineWithMiddleware(handlers.GetAdvertisementsHandler)).Methods(http.MethodGet, http.MethodOptions)
	r.Handle("/advertisement/{record_id}", combineWithMiddleware(handlers.GetAdvertisementHandler)).Methods(http.MethodGet, http.MethodOptions)
	r.Handle("/advertisement/{record_id}", combineWithMiddleware(handlers.UpdateAdvertisementHandler)).Methods(http.MethodPut, http.MethodOptions)
	r.Handle("/advertisement/{record_id}", combineWithMiddleware(handlers.DeleteAdvertisementHandler)).Methods(http.MethodDelete, http.MethodOptions)
	r.Handle("/advertisement/type/{type}", combineWithMiddleware(handlers.GetAdvertisementsHandler)).Methods(http.MethodGet, http.MethodOptions)
	r.Handle("/advertisement/type/{type}/{dashboard}", combineWithMiddleware(handlers.GetAdvertisementsHandler)).Methods(http.MethodGet, http.MethodOptions)
	r.Handle("/advertisement/owner/{owner}", combineWithMiddleware(handlers.GetAdvertisementOwnerHandler)).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/advertisement/{record_id}/image", handlers.GetAdvertisementImageHandler).Methods(http.MethodGet, http.MethodOptions)
	err = http.ListenAndServe(":"+PORT, r)
	if err != nil {
		fmt.Println("Error on listening port : " + PORT + err.Error())
	}
}
func combineWithMiddleware(handler http.HandlerFunc) *negroni.Negroni {
	return negroni.New(negroni.HandlerFunc(handlers.TokenVerifyMiddleware), negroni.WrapFunc(handler))
}
