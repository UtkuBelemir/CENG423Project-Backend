package handlers

import (
	"../dbPkg"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"math"
	"net/http"
	"strconv"
)

func GetAdvertisementsHandler(wri http.ResponseWriter, req *http.Request) {
	SetCORSFile(&wri)
	resp := dbPkg.ResponseModel{Success: false, Data: nil, Message: "Error when getting advertisements"}
	var tempAdvertisements []dbPkg.AdvertisementsResponse
	var err error
	vars := mux.Vars(req)
	if len(vars["type"]) == 0 {
		_ = json.NewEncoder(wri).Encode(&resp)
		return
	}
	jwtData := GetDataJWT(wri, req)
	userName := jwtData["username"].(string)
	if vars["dashboard"] == "1" && userName != "admin" {
		err = dbConn.DB.Debug().Select("record_id,title,description,price,category,owner,type,CASE WHEN image IS NOT NULL AND length(image) != 0 THEN 1 ELSE 0 END AS image").Where("status != 0 AND type = ? AND owner = ?", vars["type"], userName).Find(&tempAdvertisements).Error
	} else {
		if userName != "admin" {
			err = dbConn.DB.Debug().Select("record_id,title,description,price,category,owner,type,status,CASE WHEN image IS NOT NULL AND length(image) != 0 THEN 1 ELSE 0 END AS image").Where("type = ? AND status != 0", vars["type"]).Find(&tempAdvertisements).Error
		} else {
			err = dbConn.DB.Debug().Select("record_id,title,description,price,category,owner,type,status,CASE WHEN image IS NOT NULL AND length(image) != 0 THEN 1 ELSE 0 END AS image").Where("type = ?", vars["type"]).Find(&tempAdvertisements).Error
		}
	}
	if err != nil {
		_ = json.NewEncoder(wri).Encode(&resp)
		return
	}
	resp = dbPkg.ResponseModel{Success: true, Data: tempAdvertisements}
	_ = json.NewEncoder(wri).Encode(&resp)
	return
}

func GetAdvertisementImageHandler(wri http.ResponseWriter, req *http.Request) {
	SetCORSFile(&wri)
	wri.Header().Set("Content-Type", "image/jpeg")
	vars := mux.Vars(req)
	var tempAdvertisement dbPkg.Advertisements
	err := dbConn.DB.Debug().Select("image").Where("record_id = ?", vars["record_id"]).Find(&tempAdvertisement).Error
	if err != nil {
		fmt.Println("Error when reading image from DB in GetAdvertisementImageHandler. ", err.Error())
		return
	}
	wri.Write(tempAdvertisement.Image)
	return
}
func GetAdvertisementHandler(wri http.ResponseWriter, req *http.Request) {
	SetCORS(&wri)
	vars := mux.Vars(req)
	userName := GetDataJWT(wri, req)["username"].(string)
	var tempAdvertisement dbPkg.AdvertisementsResponse
	var err error
	resp := dbPkg.ResponseModel{Success: false, Message: "Error when getting advertisement detail"}
	if userName == "admin" {
		err = dbConn.DB.Debug().Select("record_id,title,description,price,category,owner,type,CASE WHEN image IS NOT NULL AND length(image) != 0 THEN 1 ELSE 0 END AS image").Where("record_id = ?", vars["record_id"]).Find(&tempAdvertisement).Error
	} else {
		err = dbConn.DB.Debug().Select("record_id,title,description,price,category,owner,type,CASE WHEN image IS NOT NULL AND length(image) != 0 THEN 1 ELSE 0 END AS image").Where("record_id = ? and status != 0", vars["record_id"]).Find(&tempAdvertisement).Error
	}

	if err != nil {
		fmt.Println("Error when reading image from DB in GetAdvertisementImageHandler. ", err.Error())
		_ = json.NewEncoder(wri).Encode(&resp)
		return
	}
	resp = dbPkg.ResponseModel{Success: true, Data: tempAdvertisement}
	_ = json.NewEncoder(wri).Encode(&resp)
	return
}

func GetAdvertisementOwnerHandler(wri http.ResponseWriter, req *http.Request) {
	SetCORS(&wri)
	vars := mux.Vars(req)
	var tempUser dbPkg.Users
	resp := dbPkg.ResponseModel{Success: false, Message: "Error when getting owner details"}
	if len(vars["owner"]) == 0 {
		_ = json.NewEncoder(wri).Encode(&resp)
		return
	}
	err := dbConn.DB.Debug().Select("first_name,last_name,email").Where("username = ?", vars["owner"]).Find(&tempUser).Error
	if err != nil {
		fmt.Println("Error when reading image from DB in GetAdvertisementOwnerHandler. ", err.Error())
		_ = json.NewEncoder(wri).Encode(&resp)
		return
	}
	resp = dbPkg.ResponseModel{Success: true, Data: tempUser}
	_ = json.NewEncoder(wri).Encode(&resp)
	return
}

func SaveAdvertisementHandler(wri http.ResponseWriter, req *http.Request) {
	SetCORSFile(&wri)
	var fileBuffer bytes.Buffer
	var tempAdvertisement dbPkg.Advertisements
	resp := dbPkg.ResponseModel{Success: false, Data: nil, Message: "Error when creating advertisement"}
	file, _, _ := req.FormFile("image")
	if file != nil {
		defer file.Close()
		io.Copy(&fileBuffer, file)
	}
	tempPrice := req.FormValue("price")
	if len(tempPrice) == 0 {
		tempPrice = "0"
	}

	price, err := strconv.ParseFloat(tempPrice, 64)
	if err != nil {
		fmt.Println("Error when parsing string to float in SaveAdvertisementHandler: " + err.Error())
		_ = json.NewEncoder(wri).Encode(&resp)
		return
	}
	jwtData := GetDataJWT(wri, req)
	tempAdvertisement = dbPkg.Advertisements{
		Title:       req.FormValue("title"),
		Description: req.FormValue("description"),
		Category:    req.FormValue("category"),
		Price:       math.Round(price*100) / 100,
		Owner:       jwtData["username"].(string),
		Type:        req.FormValue("type"),
		Image:       fileBuffer.Bytes(),
	}
	err = dbConn.DB.Debug().Create(&tempAdvertisement).Error
	if err != nil {
		fmt.Println("Error when inserting advertisement to DB in SaveAdvertisementHandler: " + err.Error())
		_ = json.NewEncoder(wri).Encode(&resp)
		return
	}
	resp = dbPkg.ResponseModel{Success: true, Data: nil, Message: "Advertisement created successfully"}
	_ = json.NewEncoder(wri).Encode(&resp)
	return
}
func DeleteAdvertisementHandler(wri http.ResponseWriter, req *http.Request) {
	SetCORS(&wri)
	vars := mux.Vars(req)
	resp := dbPkg.ResponseModel{Success: false, Data: nil, Message: "Error when deleting advertisement"}
	err := dbConn.DB.Debug().Where("record_id = ?", vars["record_id"]).Delete(dbPkg.Advertisements{}).Error
	if err != nil {
		_ = json.NewEncoder(wri).Encode(&resp)
		return
	}
	resp.Success = true
	resp.Message = "Advertisement deleted successfully."
	_ = json.NewEncoder(wri).Encode(&resp)
	return
}
func UpdateAdvertisementHandler(wri http.ResponseWriter, req *http.Request) {
	SetCORSFile(&wri)
	var fileBuffer bytes.Buffer
	var tempAdvertisement dbPkg.Advertisements
	vars := mux.Vars(req)
	resp := dbPkg.ResponseModel{Success: false, Data: nil, Message: "Error when updating advertisement"}
	omitColumns := "record_id,type,owner"
	file, _, _ := req.FormFile("image")
	if file != nil {
		defer file.Close()
		io.Copy(&fileBuffer, file)
	} else {
		if req.FormValue("image") == "1" {
			omitColumns = "record_id,type,owner,image"
		}
	}
	fmt.Println("FİLLEEEE", req.FormValue("status"))
	tempPrice := req.FormValue("price")
	if len(tempPrice) == 0 {
		tempPrice = "0"
	}

	price, err := strconv.ParseFloat(tempPrice, 64)
	if err != nil {
		fmt.Println("Error when parsing string to float in UpdateAdvertisementHandler: " + err.Error())
		_ = json.NewEncoder(wri).Encode(&resp)
		return
	}
	tempAdvertisement = dbPkg.Advertisements{
		RecordId:    vars["record_id"],
		Title:       req.FormValue("title"),
		Description: req.FormValue("description"),
		Category:    req.FormValue("category"),
		Price:       math.Round(price*100) / 100,
		Image:       fileBuffer.Bytes(),
	}
	err = dbConn.DB.Model(&tempAdvertisement).Debug().Omit(omitColumns).Updates(&tempAdvertisement).Error
	if len(req.FormValue("status")) != 0 {
		status, errToInt := strconv.Atoi(req.FormValue("status"))
		if errToInt == nil {
			err = dbConn.DB.Model(&tempAdvertisement).Debug().Where("record_id = ?", tempAdvertisement.RecordId).UpdateColumn("status", status).Error
		}
	}
	if err != nil {
		fmt.Println("Error when updateing advertisement to DB in UpdateAdvertisementHandler: " + err.Error())
		_ = json.NewEncoder(wri).Encode(&resp)
		return
	}
	resp = dbPkg.ResponseModel{Success: true, Data: nil, Message: "Advertisement updated successfully"}
	_ = json.NewEncoder(wri).Encode(&resp)
	return
}
