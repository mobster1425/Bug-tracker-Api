package controls

import (
	"feyin/bug-tracker/config"
	"feyin/bug-tracker/models"
	"feyin/bug-tracker/utils"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type MemberDTO struct {
	MembersID      uint      `json:"members_id" gorm:"column:id"`
	JoinedAt       time.Time `json:"joined_at" gorm:"column:joined_at"`
	MemberID       uint      `json:"member_id" gorm:"column:id"`
	MemberUserName string    `json:"member_user_name" gorm:"column:user_name"`
}

type request struct {
	Members []string `json:"members" binding:"required"`
}

// @Summary Add members to a project
// @Description Add members to a project by project ID
// @Tags Members
// @ID add-project-members
// @Param projectId path int true "Project ID"
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param request body request true "Member IDs"
// @Success 200 {array} MemberDTO "Added members"
// @Failure 400 {object} string "Bad Request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 403 {object} string "Forbidden"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /projects/{projectId}/members [post]
func AddProjectMembers(c *gin.Context) {
	// Get member IDs from request body

	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}
	// Validate project name
	MemberError := utils.ProjectMembersError(req.Members)
	if MemberError != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": MemberError})
		return
	}
	// Get project ID from request params
	projectIDStr := c.Param("projectId")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	// Fetch existing members
	var existingMembers []models.Member
	if err := config.DB.Where("project_id = ?", projectID).Find(&existingMembers).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid project ID"})
		return
	}

	// Fetch target project
	var targetProject models.Project
	if err := config.DB.Preload("Members").Where("id = ?", projectID).First(&targetProject).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid project ID"})
		return
	}

	// Get user ID from request
	userIDStr := c.GetString("userid")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Check if the requesting user is the project creator
	if targetProject.UserId != uint(userID) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access is denied."})
		return
	}

	// Get current member IDs
	/*
		currentMembers := make(map[string]bool)
		for _, member := range targetProject.Members {
			currentMembers[member.Member.UserName] = true
		}
	*/

	// Create a map to check for existing members
	currentMembers := make(map[uint]bool)
	for _, member := range existingMembers {
		currentMembers[member.UserId] = true
	}

	// Validate new member IDs
	for _, memberIDStr := range req.Members {
		memberID, err := strconv.Atoi(memberIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid member ID"})
			return
		}

		if currentMembers[uint(memberID)] {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Member is already added."})
			return
		}
	}

	// Validate new member IDs
	/*
		for _, memberID := range req.Members {
			if currentMembers[memberID] {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Member is already added."})
				return
			}
		}
	*/
	// Insert new members
	var membersToInsert []models.Member
	for _, memberIDStr := range req.Members {
		memberID, err := strconv.Atoi(memberIDStr)
		if err != nil {
			// Handle the error, e.g., by returning an error response
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid member ID"})
			return
		}

		member := models.Member{
			JoinedAt: time.Now(),
			//Project:  models.Project{ID: uint(projectID)},
			ProjectId: uint(projectID),
			UserId:    uint(memberID),
			//Member:   models.User{ID: uint(memberID)}, // Use memberID here
		}
		membersToInsert = append(membersToInsert, member)
	}

	db := config.DB

	tx := db.Begin()

	if err := tx.Create(&membersToInsert).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add members"})
		return
	}

	tx.Commit()

	var memberDTO []MemberDTO
	// Fetch updated members
	//	var updatedMembers []models.Member
	if err := db.
		Table("members").
		Joins("LEFT JOIN users ON members.user_id = users.id").
		Where("project_id = ?", projectID).
		Select("members.id, members.joined_at, users.id , users.user_name ").
		//Scan(&updatedMembers).Error; err != nil {
		Scan(&memberDTO).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch added members"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"members": memberDTO})

}

// @Summary Remove a project member
// @Description Remove a member from a project by project ID and member ID
// @Tags Members
// @ID remove-project-member
// @Param projectId path int true "Project ID"
// @Param memberId path int true "Member ID"
// @Security ApiKeyAuth
// @Produce json
// @Success 400 {object} string "Project Member removed successfully"
// @Failure 400 {object} string "Bad Request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 403 {object} string "Forbidden"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /projects/{projectId}/members/{memberId} [delete]
func RemoveProjectMember(c *gin.Context) {
	// Get project and member IDs from request params
	projectIDStr := c.Param("projectId")
	memberIDStr := c.Param("memberId")

	// Convert projectID and memberID to uint
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}
	memberID, err := strconv.Atoi(memberIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid member ID"})
		return
	}

	// Extract user ID from the request
	userIDStr := c.GetString("userid")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Find target project with members
	var targetProject models.Project
	if err := config.DB.Preload("Members").Where("id = ?", projectID).First(&targetProject).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Invalid project ID"})
		return
	}

	if targetProject.UserId != uint(userID) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Access is denied."})
		return
	}

	// Check if current user is the project creator
	if targetProject.UserId == uint(memberID) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Project creator can't be removed."})
		return
	}

	db := config.DB
	tx := db.Begin()

	// Find the member to be deleted
	var memberToDelete models.Member
	if err := tx.Where("project_id = ? AND user_id = ?", projectID, memberID).First(&memberToDelete).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find project member to delete"})
		return
	}

	// Delete the member from the members table
	if err := tx.Delete(&memberToDelete).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove project member from members table"})
		return
	}

	tx.Commit()
	fmt.Println("Last step")

	c.JSON(200, gin.H{
		"Message": "Project Member removed successfully",
	})
}

// @Summary Leave a project as a member
// @Description Leave a project as a member by project ID
// @Tags Members
// @ID leave-project-as-member
// @Param projectId path int true "Project ID"
// @Security ApiKeyAuth
// @Produce json
// @Success 400 {object} string "Project Member left successfully"
// @Failure 400 {object} string "Bad Request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 403 {object} string "Forbidden"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /projects/{projectId}/members/leave [post]
func LeaveProjectAsMember(c *gin.Context) {
	projectIDStr := c.Param("projectId")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	// Extract user ID from the request
	userIDStr := c.GetString("userid")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Find the target project with members using GORM
	var targetProject models.Project
	if err := config.DB.Preload("Members").Where("id = ?", projectID).First(&targetProject).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Project not found"})
		return
	}

	if targetProject.UserId == uint(userID) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Project creator can't leave."})
		return
	}

	isMember := false
	for _, member := range targetProject.Members {
		// if member.Member.ID == uint(userID) {
		if member.UserId == uint(userID) {
			isMember = true
			break
		}
	}

	if !isMember {
		c.JSON(http.StatusNotFound, gin.H{"message": "You're not a member of the project."})
		return
	}

	db := config.DB

	// Find the member to delete
	var targetMember models.Member
	if err := db.
		Where("project_id = ? AND user_id = ?", projectID, userID).
		First(&targetMember).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
		return
	}

	tx := db.Begin()

	// Check if the current user is the project creator (assuming targetMember.Project.UserId is the project creator ID)
	if targetMember.Project.UserId == uint(userID) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Project creator can't leave."})
		return
	}

	// Delete the member
	if err := tx.Delete(&targetMember).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to leave the project"})
		return
	}

	tx.Commit()

	// c.JSON(http.StatusNoContent)
	c.JSON(400, gin.H{
		"Message": "Project Member left successfully",
	})
}
