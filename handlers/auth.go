package handlers

import (
	"../dbPkg"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

var dbConn = dbPkg.New()

func SignUpHandler(wri http.ResponseWriter, req *http.Request) {
	SetCORS(&wri)

	var tempUserData dbPkg.Users
	resp := dbPkg.ResponseModel{Success: false, Data: nil, Message: "Error when creating user."}
	err := json.NewDecoder(req.Body).Decode(&tempUserData)
	if err != nil {
		fmt.Println("Error when decoding user data in SignUpHandler : " + err.Error())
		_ = json.NewEncoder(wri).Encode(resp)
		return
	}
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(tempUserData.Password), 10)
	tempUserData.Password = string(passwordHash)
	tempUserData.Role = "student"
	err = dbConn.DB.Create(&tempUserData).Error

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			resp.Message = "Username and email should be unique"
		}
		fmt.Println("Error when insert user data to DB in SignUpHandler : " + err.Error())
		_ = json.NewEncoder(wri).Encode(resp)
		return
	}
	resp.Success = true
	resp.Message = "User created successfully"
	_ = json.NewEncoder(wri).Encode(resp)
}
func LoginHandler(wri http.ResponseWriter, req *http.Request) {
	SetCORS(&wri)

	var loginInfo, userData dbPkg.Users
	resp := dbPkg.ResponseModel{Success: false, Data: nil, Message: "Username or password is wrong"}
	err := json.NewDecoder(req.Body).Decode(&loginInfo)
	if err != nil {
		fmt.Println("Error when decoding user data in LoginHandler : " + err.Error())
		_ = json.NewEncoder(wri).Encode(resp)
		return
	}
	err = dbConn.DB.Where("username = ?", loginInfo.Username).Find(&userData).Error
	if err != nil {
		fmt.Println("Error when reading user data from DB in LoginHandler : " + err.Error())
		_ = json.NewEncoder(wri).Encode(resp)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(loginInfo.Password))
	if err != nil {
		resp.Message = "Username or password is wrong"
		_ = json.NewEncoder(wri).Encode(resp)
		return
	}
	token, err := CreateJwtToken(userData.Username)
	if err != nil {
		fmt.Println("Error when creating JWT token in LoginHandler : " + err.Error())
		_ = json.NewEncoder(wri).Encode(resp)
		return
	}
	resp.Success = true
	resp.Message = ""
	resp.Data = dbPkg.UserClaims{
		Username:  userData.Username,
		Email:     userData.Email,
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Gender:    userData.Gender,
		Role:      userData.Role,
		Token:     token,
	}
	_ = json.NewEncoder(wri).Encode(resp)
	return
}
func RenewPasswordHandler(wri http.ResponseWriter, req *http.Request) {
	SetCORS(&wri)
	var loginInfo, userData dbPkg.Users
	resp := dbPkg.ResponseModel{Success: false, Data: nil, Message: "Please check your information deatils"}
	err := json.NewDecoder(req.Body).Decode(&loginInfo)
	if err != nil {
		fmt.Println("Error when decoding user data in RenewPasswordHandler : " + err.Error())
		_ = json.NewEncoder(wri).Encode(resp)
		return
	}
	err = dbConn.DB.Where("username = ? and first_name = ? and email = ?", loginInfo.Username, loginInfo.FirstName, loginInfo.Email).Find(&userData).Error
	if err != nil {
		fmt.Println("Error when reading user data from DB in RenewPasswordHandler : " + err.Error())
		_ = json.NewEncoder(wri).Encode(resp)
		return
	}
	if userData == (dbPkg.Users{}) {
		_ = json.NewEncoder(wri).Encode(resp)
		return
	}
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(loginInfo.Password), 10)
	userData.Password = string(passwordHash)
	err = dbConn.DB.Save(&userData).Error
	if err != nil {
		fmt.Println("Error when updating user data from DB in RenewPasswordHandler : " + err.Error())
		resp.Message = "Error when updating password"
		_ = json.NewEncoder(wri).Encode(resp)
		return
	}
	resp.Success = true
	resp.Message = "Password updated successfully"
	_ = json.NewEncoder(wri).Encode(resp)
	return
}
func CookieLoginHandler(wri http.ResponseWriter, req *http.Request) {
	SetCORS(&wri)
	var userData dbPkg.Users
	resp := dbPkg.ResponseModel{Success: false, Data: nil, Message: "Cookie is not valid"}
	jwtData := GetDataJWT(wri, req)
	err := dbConn.DB.Where("username = ?", jwtData["username"]).Find(&userData).Error
	if err != nil {
		fmt.Println("Error when reading user from DB in CookieLoginHandler: " + err.Error())
		_ = json.NewEncoder(wri).Encode(resp)
		return
	}
	token, err := CreateJwtToken(userData.Username)
	if err != nil {
		fmt.Println("Error when creating JWT token in CookieLoginHandler : " + err.Error())
		_ = json.NewEncoder(wri).Encode(resp)
		return
	}
	resp.Success = true
	resp.Message = ""
	resp.Data = dbPkg.UserClaims{
		Username:  userData.Username,
		Email:     userData.Email,
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Gender:    userData.Gender,
		Role:      userData.Role,
		Token:     token,
	}
	_ = json.NewEncoder(wri).Encode(resp)
}
