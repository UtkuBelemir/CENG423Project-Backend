package handlers

import (
	"../dbPkg"
	"encoding/json"
	"fmt"
	"net/http"
)

func ProfilePutHandler(wri http.ResponseWriter, req *http.Request) {
	SetCORS(&wri)
	var newUserData dbPkg.Users
	resp := dbPkg.ResponseModel{Success: false, Data: nil, Message: "Error when updating profile"}
	err := json.NewDecoder(req.Body).Decode(&newUserData)
	if err != nil {
		fmt.Println("Error when parsing new data in ProfileHandler: " + err.Error())
		_ = json.NewEncoder(wri).Encode(resp)
		return
	}
	jwtData, err := JWTData(req.Header.Get("Authorization"))
	if err != nil {
		fmt.Println("Error when parsing JWT in ProfileHandler: " + err.Error())
		_ = json.NewEncoder(wri).Encode(resp)
		return
	}
	newUserData.Username = jwtData["username"].(string)
	err = dbConn.DB.Save(&newUserData).Error
	if err != nil {
		fmt.Println("Error when updating profile information to DB in ProfileHandler: " + err.Error())
		_ = json.NewEncoder(wri).Encode(resp)
		return
	}
	resp.Success = true
	resp.Message = ""
	resp.Data = dbPkg.UserClaims{
		Username:  newUserData.Username,
		Email:     newUserData.Email,
		FirstName: newUserData.FirstName,
		LastName:  newUserData.LastName,
		Gender:    newUserData.Gender,
		Role:      newUserData.Role,
	}
	_ = json.NewEncoder(wri).Encode(resp)
}
func ProfileGetHandler(wri http.ResponseWriter, req *http.Request) {
	SetCORS(&wri)
	fmt.Fprintln(wri, "ok")
	return
}
