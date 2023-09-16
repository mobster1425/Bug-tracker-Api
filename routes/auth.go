package routes

import (
	"feyin/bug-tracker/controls"
	"feyin/bug-tracker/middleware"

	"github.com/gin-gonic/gin"
)

// the purpose of the UserRouts function is to create a route group for user-related operations under the /user path

func UserRoutes(c *gin.Engine) {
	User := c.Group("/user")
	{

		//User rountes
		User.POST("/login", controls.Login)
		User.POST("/signup", controls.SignUp)
		User.POST("/signup/otpvalidate", controls.OtpValidation)
		User.GET("/logout", middleware.UserAuth, controls.Signout)
		User.GET("/users", middleware.UserAuth, controls.GetAllUsers)
		User.POST("/userchangepassword", middleware.UserAuth, controls.ChangePassword)
		User.PUT("/userchangepassword/updatepassword", middleware.UserAuth, controls.UpdatePassword)

		//Forgot Password
		//	User.PUT("/forgotpassword", middleware.UserAuth, controls.GenerateOtpForForgotPassword)
		User.PUT("/forgotpassword", controls.GenerateOtpForForgotPassword)
		// 	User.POST("/forgotpassword/changepassword", middleware.UserAuth, controls.ForgotPassword)
		User.POST("/forgotpassword/changepassword", controls.ForgotPassword)
	}

}
