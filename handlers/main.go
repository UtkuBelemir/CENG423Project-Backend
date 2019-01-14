package handlers

import (
	"fmt"
	"net/http"
)

func IndexHandler(wri http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(wri, "ok")
}
func SetCORS(wri *http.ResponseWriter) {

	(*wri).Header().Set("Content-Type", "application/json; charset=utf-8")
	(*wri).Header().Set("Access-Control-Allow-Methods", "*")
	(*wri).Header().Set("Access-Control-Allow-Origin", "*")
	//Access-Control-Allow-Origin
	(*wri).Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept,Authorization")
}
func SetCORSFile(wri *http.ResponseWriter) {

	(*wri).Header().Set("Content-Type", "multipart/form-data")
	(*wri).Header().Set("Access-Control-Allow-Methods", "*")
	(*wri).Header().Set("Access-Control-Allow-Origin", "*")
	//Access-Control-Allow-Origin
	(*wri).Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept,Authorization")
}
