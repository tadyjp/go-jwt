package handlers

import (
	"io/ioutil"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// Sign hoge
func Sign(c echo.Context) error {
	signedClaim, err := signClaim()
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, signedClaim)
}

func signClaim() (string, error) {
	signBytes, err := ioutil.ReadFile("configs/private.pem")
	if err != nil {
		panic(err)
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		panic(err)
	}

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}

	// Create the Claims
	claims := MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)

	ss, err := token.SignedString(signKey)
	if err != nil {
		panic(err)
	}
	return ss, nil
}
