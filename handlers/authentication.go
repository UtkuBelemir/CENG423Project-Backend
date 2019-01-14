package handlers

import (
	"../dbPkg"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	privateKeyPath = "/keys/priv.key"
	publicKeyPath  = "/keys/pub.key"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

type TokenClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func InitializeJWTKeys(workPath string) {
	signBytes, err := ioutil.ReadFile(workPath + privateKeyPath)
	if err != nil {
		panic("Error on reading private key ! : " + err.Error())
	}
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		panic("Error on parsing private key ! : " + err.Error())
	}
	verifyBytes, err := ioutil.ReadFile(workPath + publicKeyPath)
	if err != nil {
		panic("Error on reading verify key ! : " + err.Error())
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		panic("Error on parsing verify key ! : " + err.Error())
	}
}

func CreateJwtToken(userName string) (string, error) {
	claims := TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(), //TODO: ZAMANI AYARLA
		},
		userName,
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokent, err := rawToken.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return tokent, nil
}

func TokenVerifyMiddleware(wri http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	SetCORS(&wri)
	if IsTokenAcceptable(req) {
		next(wri, req)
		return
	}
	_ = json.NewEncoder(wri).Encode(dbPkg.ResponseModel{Success: false, Data: nil, Message: "Unauthorized Access"})
}

func IsTokenAcceptable(req *http.Request) bool {
	cook := req.Header.Get("Authorization")
	token, err := jwt.Parse(cook, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})
	if err == nil && token.Valid {
		return true
	}
	return false
}

func JWTData(curToken string) (jwt.MapClaims, error) {
	usersToken, err := jwt.Parse(curToken, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})
	if err == nil && usersToken.Valid {
		fmt.Println(usersToken.Claims.(jwt.MapClaims))
		return usersToken.Claims.(jwt.MapClaims), nil
	}
	return nil, errors.New("Token is not valid")
}
func GetDataJWT(wri http.ResponseWriter, req *http.Request) jwt.MapClaims {
	resp := dbPkg.ResponseModel{Success: false, Data: nil, Message: "Cookie is not valid"}
	cook := req.Header.Get("Authorization")
	if len(cook) == 0 {
		_ = json.NewEncoder(wri).Encode(resp)
		return nil
	}
	jwtData, err := JWTData(cook)
	if err != nil {
		fmt.Println("Error when parsing JWT in CookieLoginHandler: " + err.Error())
		_ = json.NewEncoder(wri).Encode(resp)
		return nil
	}
	return jwtData
}

/*
Private Key Generation

ssh-keygen -t rsa -b 4096 -f jwtRS256.key
# Don't add passphrase
openssl rsa -in jwtRS256.key -pubout -outform PEM -out jwtRS256.key.pub
cat jwtRS256.key
cat jwtRS256.key.pub
*/
