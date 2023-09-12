package controls

import (
	"feyin/bug-tracker/config"
	"fmt"
	"net/http"
	"strconv"

	//	"feyin/bug-tracker/models"

	"github.com/gin-gonic/gin"
	//"golang.org/x/crypto/bcrypt"
)

type UsersDTO struct {
	UserID   uint   `json:"user_id" gorm:"column:id"`
	UserName string `json:"user_name" gorm:"column:user_name"`
}

// GetAllUsers godoc
// @Summary Get all users except the authenticated user
// @Description Get a list of all users except the authenticated user.
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} UsersDTO "List of users"
// @Failure 400 {object} string "Invalid user ID"
// @Failure 500 {object} string "Failed to fetch users"
// @Router /user/users [get]
func GetAllUsers(c *gin.Context) {
	fmt.Println("Get all users function started")

	userIDStr := c.GetString("userid")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	db := config.DB
	//	var users []models.User

	var usersDTO []UsersDTO

	if err := db.
		Table("users").
		Where("users.id != ?", userID).
		Select("users.id, users.user_name ").
		Scan(&usersDTO).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	fmt.Println("Get all users function ended suceesfully")
	//	c.JSON(http.StatusOK, users)
	c.JSON(http.StatusOK, gin.H{"Users": usersDTO})

}
