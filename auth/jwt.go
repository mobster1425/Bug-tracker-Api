package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// getting secret env from .env
var Jwtkey = []byte(os.Getenv("SECRET"))

func TokenGeneration(id string) string {
	/*

		Here, a new JWT token is created using jwt.NewWithClaims. It specifies the signing method (jwt.SigningMethodHS256)
		and the claims to be included in the token. In this case, two claims are added:

		"sub" (subject): This claim contains the user identifier (id) that is passed to the function.
		"exp" (expiration time): This claim sets an expiration time for the token. In this example,
		 the token will expire in 30 days from the current time.
	*/
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	/*
		The SignedString method is used to sign the token using the secret key (Jwtkey).
		 If signing the token encounters an error, it will result in a panic with the error message.

	*/
	tokenString, err := token.SignedString(Jwtkey)
	if err != nil {
		panic(err)
	}
	return tokenString

}
