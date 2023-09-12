package controls

import (
	"feyin/bug-tracker/auth"
	"feyin/bug-tracker/config"
	"feyin/bug-tracker/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type checkUserData struct {
	UserName        string
	Email           string
	Password        string
	ConfirmPassword string
	Otp             string
}

// user sign up

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

// SignUp godoc
// @Summary Create a new user
// @Description Create a new user with the provided data.
// @Tags users
// @Accept json
// @Produce json
// @Param user body checkUserData true "User data"
// @Success 202 {string} string "Go to /signup/otpvalidate"
// @Failure 400 {object} ErrorResponse "Data binding error"
// @Failure 409 {object} ErrorResponse "User already Exist"
// @Router /user/signup [post]

func SignUp(c *gin.Context) {

	// the checkuserdata struct is used for validating user inputs from the client and perform operation before we
	// insert into the database
	var Data checkUserData

	if c.Bind(&Data) != nil {
		c.JSON(400, gin.H{
			"error": "Data binding error",
		})
		return
	}

	var temp_user models.User
	hash, err := bcrypt.GenerateFromPassword([]byte(Data.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": "False",
			"Error":  "Hashing password error",
		})
		return
	}

	db := config.DB
	result := db.First(&temp_user, "email LIKE ?", Data.Email)
	// result := db.Where("email = ?", Data.Email).First(&temp_user)

	// Debug print statement
	fmt.Println("Result Error:", result.Error)

	// if there is an error that means no such user in the database , so we insert into it
	if result.Error != nil {
		//storing user data in the model
		user := models.User{

			UserName: Data.UserName,
			Email:    Data.Email,
			Password: string(hash),
		}

		// getting otp from verifyotp after sending a mail to the receiptant email
		otp := VerifyOTP(Data.Email)

		// Debug print statement
		fmt.Println("Generated OTP:", otp)

		//insert user into database
		result2 := db.Create(&user)

		// if there was an error with the insertion
		// else if no error insert otp into database
		if result2.Error != nil {
			c.JSON(500, gin.H{
				"Status": "False",
				"Error":  "User data creating error",
			})
		} else {
			db.Model(&user).Where("email LIKE ?", user.Email).Update("otp", otp)

			// Debug print statement
			fmt.Println("User inserted and OTP updated")

			// direct the user to validate the otp
			c.JSON(202, gin.H{
				"message": "Go to /signup/otpvalidate",
			})
		}
		// if there is no error i.e result.error == nil, then a user was found
	} else {
		c.JSON(409, gin.H{
			"Error": "User already Exist",
		})
		return
	}

}

// user login
type userData struct {
	Email    string
	Password string
}

// Login godoc
// @Summary Log in a user
// @Description Log in a user with the provided email and password.
// @Tags users
// @Accept json
// @Produce json
// @Param user body userData true "User data"
// @Success 200 {object} string "User login successfully"
// @Failure 400 {object} string "Data binding error"
// @Failure 404 {object} string "User not found"
// @Failure 400 {object} string "Password is incorrect"
// @Router /user/login [post]
func Login(c *gin.Context) {
	fmt.Println("Starting Login...")

	var user userData
	if c.Bind(&user) != nil {
		c.JSON(400, gin.H{
			"error": "Data binding error",
		})
		return
	}
	var checkUser models.User
	db := config.DB

	result := db.First(&checkUser, "email LIKE ?", user.Email)
	if result.Error != nil {
		c.JSON(404, gin.H{
			"Status":  "false",
			"Message": result.Error.Error(),
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": "false",
			"error":  "Password is incorrect",
		})
		return
	}

	//>>>>>>>>>>>>>>>>> Generating a JWT-tokent <<<<<<<<<<<<<<<//
	//strconv.Itoa is a function from the strconv package that converts an integer to its string representation.
	str := strconv.Itoa(int(checkUser.ID))

	tokenString := auth.TokenGeneration(str)
	/*
		This line sets the same-site attribute of cookies to "Lax" mode.
		It restricts cookies from being sent in cross-site requests that change the state of the user,
		 such as a POST request from an external site.
	*/
	c.SetSameSite(http.SameSiteLaxMode)

	/*
	   This line sets a cookie named "UserAutherization" with the tokenString value.
	   The third argument, 3600*24*30, is the expiration time of the cookie, set to 30 days in seconds.
	   The remaining arguments set the domain, path, secure flag, and HttpOnly flag for the cookie.
	*/
	c.SetCookie("UserAutherization", tokenString, 3600*24*30, "", "", false, true)
	fmt.Println("User login successfully")
	c.JSON(http.StatusOK, gin.H{
		"message": "User login successfully",
	})

}

// Signout godoc
// @Summary Log out a user
// @Description Log out a user by clearing the authentication cookie.
// @Tags users
// @Produce json
// @Success 200 {object} string "User Successfully Log Out"
// @Router /user/logout [get]
func Signout(c *gin.Context) {
	// User logout
	fmt.Println("Starting Logout...")
	c.SetCookie("UserAutherization", "", -1, "", "", false, false)
	c.JSON(200, gin.H{
		"Message": "User Successfully  Log Out",
	})
	fmt.Println("LogOut finished...")
}

// validate
func Validate(c *gin.Context) {
	c.Get("user")
	c.JSON(200, gin.H{
		"message": "User login successfully",
	})
}

// ChangePassword godoc
// @Summary Change a user's password
// @Description Change a user's password by providing the current password and the new password.
// @Tags users
// @Accept json
// @Produce json
// @Param userEnterData body checkUserData true "User data"
// @Success 200 {object} string "Password changed successfully"
// @Failure 400 {object} string "Data binding error"
// @Failure 400 {object} string "Password not match"
// @Failure 500 {object} string "Error in string conversion"
// @Failure 409 {object} string "User not exist"
// @Failure 400 {object} string "Password is incorrect"
// @Router /user/userchangepassword [post]
func ChangePassword(c *gin.Context) {

	/*
	   UserChangePassword: Validates whether a provided password matches the stored password for a user.
	    This function helps determine whether a user has entered the correct password before proceeding with further actions.

	   UpdatePassword: Updates the password in the database for a specific user.

	   In summary, both functions have distinct purposes – one for validation and the other for updating the password –
	   and they can remain separate for clarity and maintainability.
	*/

	// change user password, here user is not forgetting the password, but he is chagning it himself

	fmt.Println("ChangePassword controller called")

	var userEnterData checkUserData
	if c.Bind(&userEnterData) != nil {
		c.JSON(400, gin.H{
			"error": "Data binding error",
		})
		fmt.Println("Data binding error")
		return
	}

	if userEnterData.Password != userEnterData.ConfirmPassword {
		c.JSON(400, gin.H{
			"Message": "Password not match",
		})
		fmt.Println("Password and ConfirmPassword do not match")
		return
	}

	id, err := strconv.Atoi(c.GetString("userid"))

	if err != nil {
		c.JSON(500, gin.H{
			"Error": "Error in string conversion",
		})
		fmt.Println("Error in string conversion:", err)
		return
	}

	var userData models.User
	db := config.DB
	result := db.First(&userData, "id = ?", id)
	if result.Error != nil {
		c.JSON(409, gin.H{
			"Error": "User not exist",
		})
		fmt.Println("User not found in the database")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(userEnterData.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "false",
			"error":  "Password is incorrect",
		})
		fmt.Println("Password is incorrect")
		return
	} else {
		c.JSON(200, gin.H{
			"message": "Go to /userchangepassword/updatepassword",
		})
		// fmt.Println("Password changed successfully")
		fmt.Println("the user has changed his password successfully, he did not forgrt it,he wants to change it,after we update")
	}

}

// UpdatePassword godoc
// @Summary Update a user's password
// @Description Update a user's password by providing the user's current password and the new password.
// @Tags users
// @Accept json
// @Produce json
// @Param userEnterData body checkUserData true "User data"
// @Success 202 {object} string "Successfully updated password"
// @Failure 400 {object} string "Data binding error"
// @Failure 400 {object} string "Hashing password error"
// @Failure 400 {object} string "Error in string conversion"
// @Failure 409 {object} string "User not exist"
// @Router /user/userchangepassword/updatepassword [post]
func UpdatePassword(c *gin.Context) {
	// Update password
	fmt.Println("UpdatePassword controller called")

	var userEnterData checkUserData
	var userData models.User
	if c.Bind(&userEnterData) != nil {
		c.JSON(400, gin.H{
			"error": "Data binding error",
		})
		fmt.Println("Data binding error")
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(userEnterData.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": "False",
			"Error":  "Hashing password error",
		})
		fmt.Println("Hashing password error:", err)
		return
	}

	id, err := strconv.Atoi(c.GetString("userid"))
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Error in string conversion",
		})
		fmt.Println("Error in string conversion:", err)
		return
	}
	db := config.DB
	result := db.First(&userData, "id = ?", id)
	if result.Error != nil {
		c.JSON(409, gin.H{
			"Error": "User not exist",
		})
		fmt.Println("User not found in the database")
		return
	}

	db.Model(&userData).Where("id = ?", id).Update("password", hash)
	c.JSON(202, gin.H{
		"message": "Successfully updated password",
	})
	fmt.Println("Password updated successfully")

}
