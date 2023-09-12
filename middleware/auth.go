package middleware

import (
	"fmt"

	"net/http"
	"os"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func UserAuth(c *gin.Context) {
	fmt.Println("Starting UserAuth...")
	//Get the cookie off req
	tokenString, err := c.Cookie("UserAutherization")

	if err != nil {
		fmt.Println("Invalid access, User logout")
		c.JSON(401, gin.H{
			"Massage": "Invalid access, User logout",
		})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Decode/validate it
	//Verification is done here to check if the token from client and env token are the same and follow same signingmethod
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		c.JSON(500, gin.H{
			"Status": "False",
			"Error":  "Error occured while token genaration",
		})
	}

	//Decoding is done here , we are extracting the claims, the payloads i mean
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("userid", claims["sub"])

		c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)

	}
	fmt.Println("User authenticated successfully")
}

/*

1.`func UserAuth(c *gin.Context) {`: This function is a middleware function in the Gin framework, designed to handle user
 authentication before allowing access to certain routes.

2. `tokenString, err := c.Cookie("UserAutherization")`: This line retrieves the value of the "UserAutherization"
cookie from the incoming HTTP request's headers. If the cookie doesn't exist or there's an error while retrieving it,
the `err` variable will be set. If the cookie exists, its value will be stored in the `tokenString` variable.

3. `if err != nil { ... }`: This block of code checks if there was an error while retrieving the cookie.
If an error occurred (indicating that the cookie is not present or there was a problem reading it),
a JSON response with a 401 Unauthorized status is sent, indicating that the access is invalid. The `c.AbortWithStatus`
 function terminates the request and stops further execution of the middleware.

4. Next, the code decodes and validates the token using the `jwt.Parse` function from the JSON Web Token (JWT) library.


 `token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {`: This line parses and validates the JWT token
 stored in `tokenString`. The `jwt.Parse` function requires a token and a validation function.

2. The validation function (`func(token *jwt.Token) (interface{}, error) { ... }`) is defined inline.
It checks whether the signing method of the token is of type `jwt.SigningMethodHMAC`, which indicates that the token was signed
using HMAC (Hash-based Message Authentication Code).

3. If the signing method is not as expected, an error is returned, indicating an unexpected signing method. Otherwise,
the function returns the secret key used for token validation.

4. The secret key for validation is retrieved from the environment variable `SECERET` (likely a typo, should be `SECRET`),
which is read using `os.Getenv("SECRET")`.


The purpose of the `jwt.Parse` function with the callback function is to verify the authenticity and integrity of the JWT.






*/
