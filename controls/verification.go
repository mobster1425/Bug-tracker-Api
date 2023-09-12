package controls

import (
	"crypto/rand"
	"fmt"
	"os"

	"net/http"
	"net/smtp"

	"feyin/bug-tracker/config"
	"feyin/bug-tracker/models"
	"math/big"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func VerifyOTP(email string) string {
	Otp, err := getRandNum()
	if err != nil {
		panic(err)
	}

	sendMail(email, Otp)
	return Otp
}

// smtp server protocol is used to send outgoing email from a users mail account to receivers emails server, in here the receiver email is
// a slice of string, because the receivers would b more than one
func sendMail(email string, otp string) {
	fmt.Println("Sending email to:", email)

	/*

		- This section of code prints out the email address and OTP to the console for debugging purposes.
		 It helps verify that the correct email and OTP are being used.

		- The `os.Getenv` function is used to retrieve environment variables named `"EMAIL"` and `"PASSWORD"`.
		These variables are expected to contain the sender's email address and password for authentication.
	*/
	fmt.Println("Email : ", email, " otp :", otp)
	// Sender data.
	from := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")

	// Receiver email address.
	//This defines the receiver's email address (the
	//email address where the OTP will be sent) as a string slice containing a single element,
	//which is the `email` parameter passed to the function.
	to := []string{
		email,
	}

	// smtp server configuration.
	//These lines define the SMTP server host and port for sending the email. In this case, it's set up for Gmail's SMTP server.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.

	// Authentication.
	//This line creates an authentication mechanism using the sender's
	//email address (`from`) and password. The `smtp.PlainAuth` function is used here to establish plain text authentication.
	//Your email client provides your Gmail username (email address) and password to authenticate with Gmail's SMTP server.
	//This ensures that you're authorized to use the server for sending emails.

	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.

	/*
		This section uses the `smtp.SendMail` function to send the email. It takes several parameters:
		  - The SMTP server host and port.
		  - The authentication mechanism created earlier.
		  - The sender's email address.
		  - The receiver's email address (stored in the `to` slice).
		  - The email body, which contains the OTP.
	*/
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(otp+" is your otp"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email sent successfully to:", email)
}

func getRandNum() (string, error) {
	nBig, e := rand.Int(rand.Reader, big.NewInt(8999))
	if e != nil {
		return "", e
	}
	return strconv.FormatInt(nBig.Int64()+1000, 10), nil
}

// -------Otp validatioin------------->
// the OtpValidation function is responsible for validating the OTP provided by the user during registration

type User_otp struct {
	Otp   string
	Email string
}

// OtpValidation godoc
// @Summary Validate OTP
// @Description Validate OTP provided by the user during registration
// @Tags users
// @Accept json
// @Produce json
// @Param data body User_otp true "User data"
// @Success 200 {object} string "New User Successfully Registered"
// @Failure 400 {object} string "Could not bind the JSON Data"
// @Failure 404 {object} string "User not found"
// @Failure 422 {object} string "Wrong OTP, Retry again"
// @Router /user/signup/otpvalidate [post]
func OtpValidation(c *gin.Context) {
	fmt.Println("Starting OTP Validation...")

	var user_otp User_otp
	var userData models.User
	if c.Bind(&user_otp) != nil {
		c.JSON(400, gin.H{
			"Error": "Could not bind the JSON Data",
		})
		return
	}
	db := config.DB
	result := db.First(&userData, "otp LIKE ? AND email LIKE ?", user_otp.Otp, user_otp.Email)

	//If the query result has an error, a JSON response with a 404 status code is sent, and the error message from the result is included.
	//Additionally, a JSON response with a 422 status code is sent with an error message indicating that the OTP is wrong and the user
	//should try registering again.
	if result.Error != nil {
		fmt.Println("Error:", result.Error)
		c.JSON(404, gin.H{
			"Error": result.Error.Error(),
		})
		c.JSON(422, gin.H{
			"Error":   "Wrong OTP ,Retry again",
			"Message": "Goto /signup/otpvalidate",
		})
		return
	}
	fmt.Println("User successfully registered.")
	c.JSON(200, gin.H{
		"Message": "New User Successfully Registered",
	})
}

/*
STEP 1

# SIGNUP

what happens in the sign up function is that, after validating users input in the checkuserdata struct,we create a hashed password and
validate other users input
we store into the database,after we get otp from the verifyotp , the verifyotp function calls the sendmail func before returning the otp
the send mail func sends the otp to the users email and then we store the otp into the database
after we then direct user to go validate otp

STEP 2
OTP VALIDATION

after that in the otpvalidate function, we are sending the user email and otp received from the users mail,
so we are checking if the otp sent by the cleint and the email sent by client are the same with the otp,email stored in the database
during signup function, then the user is successfully registered, else we tell the user to re-enter the otp
*/

//it is the same step for  both of these 2 function = generateotpforforgetpasswprd and userchangepassword

// Generating otp for forgot password

type UserEnterData2 struct {
	Email string
}

// GenerateOtpForForgotPassword godoc
// @Summary Generate OTP for forgot password
// @Description Generate OTP for the user's forgot password request.
// @Tags users
// @Accept json
// @Produce json
// @Param data body UserEnterData2 true "User data"
// @Success 200 {object} string "OTP sent successfully"
// @Failure 400 {object} string "Data binding error"
// @Failure 409 {object} string "User not exist"
// @Router /user/forgotpassword [put]
func GenerateOtpForForgotPassword(c *gin.Context) {
	fmt.Println("GenerateOtpForForgotPassword controller called")

	var data UserEnterData2
	if c.Bind(&data) != nil {
		c.JSON(400, gin.H{
			"Error": "Error when the data binding",
		})
		fmt.Println("Data binding error")
		return
	}
	otp := VerifyOTP(data.Email)

	db := config.DB
	var userData models.User

	result := db.Model(userData).Where("email = ?", data.Email).Update("otp", otp)
	if result.Error != nil {
		c.JSON(409, gin.H{
			"Error": "User not exist",
		})
		fmt.Println("User not found in the database")
		return
	}
	c.JSON(200, gin.H{
		"Message": "otp send go to /user/forgotpassword/changepassword",
	})
	fmt.Println("OTP sent successfully")
}

// Reseting the password after forgot password

type userEnterData1 struct {
	Email           string
	Otp             string
	Password        string
	ConfirmPassword string
}

// UserChangePassword godoc
// @Summary Change user's password
// @Description Change a user's password by providing the email, OTP, and new password.
// @Tags users
// @Accept json
// @Produce json
// @Param data body userEnterData1 true "User data"
// @Success 200 {object} string "Password changed successfully"
// @Failure 400 {object} string "Data binding error"
// @Failure 400 {object} string "Password not match"
// @Failure 400 {object} string "Hashing password error"
// @Failure 409 {object} string "User not exist"
// @Failure 400 {object} string "Invalid OTP"
// @Router /user/forgotpassword/changepassword [post]
func UserChangePassword(c *gin.Context) {
	fmt.Println("UserChangePassword controller called")

	var data userEnterData1
	var userData models.User
	if c.Bind(&data) != nil {
		c.JSON(400, gin.H{
			"Error": "Error when the data binding",
		})
		fmt.Println("Data binding error")
		return
	}
	if data.Password != data.ConfirmPassword {
		c.JSON(400, gin.H{
			"Error": "Password not match",
		})
		fmt.Println("Password mismatch")
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": "False",
			"Error":  "Hashing password error",
		})
		fmt.Println("Hashing password error:", err)
		return
	}
	db := config.DB
	result := db.Find(&userData, "email = ?", data.Email)
	if result.Error != nil {
		c.JSON(409, gin.H{
			"Error": "User not exist",
		})
		fmt.Println("User not found in the database")
		return
	}
	if data.Otp != userData.Otp {
		c.JSON(400, gin.H{
			"Error": "Invalide otp",
		})
		fmt.Println("Invalid OTP")
		return
	}
	result = db.Model(&userData).Where("email = ?", data.Email).Update("password", hash)
	if result.Error != nil {
		c.JSON(409, gin.H{
			"Error": "User not exist",
		})
		fmt.Println("User not found in the database")
		return
	}
	c.JSON(200, gin.H{
		"Message": "Password Change successfully",
	})
	fmt.Println("User forgot his password, so it has been chagned succesffuly,Password changed successfully")
}
