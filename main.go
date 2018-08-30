package main

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	privKey = "key/app.key"
	pubKey  = "key/app.key.pub"
)

var (
	signKey  *rsa.PrivateKey
	verifKey *rsa.PublicKey
)

func init() {
	signBytes, err := ioutil.ReadFile(privKey)
	if err != nil {
		panic(err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		panic(err)
	}
}

func init() {
	verifBytes, err := ioutil.ReadFile(pubKey)
	if err != nil {
		panic(err)
	}

	verifKey, err = jwt.ParseRSAPublicKeyFromPEM(verifBytes)
	if err != nil {
		panic(err)
	}
}

func main() {
	tokenStr, err := generateTokenString()
	if err != nil {
		panic(err)
	}

	fmt.Println("Encoded token:", tokenStr)

	claims, err := parseTokenString(tokenStr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Decoded claims:", *claims)
}

func parseTokenString(tokenStr string) (*jwt.MapClaims, error) {
	var claims jwt.MapClaims
	_, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		return verifKey, nil
	})
	if err != nil {
		return nil, err
	}
	return &claims, nil
}

func generateTokenString() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.MapClaims{
		"onamae": "hoge fuga",
		"myid":   12345,
	})
	return token.SignedString(signKey)
}
