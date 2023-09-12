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

type BugRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Priority    string `json:"priority" binding:"required"`
}

type BugUpdateData struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Priority    string `json:"priority" binding:"required"`
}

type BugDTO struct {
	ID                 uint      `json:"id" gorm:"column:id"`
	Title              string    `json:"title" gorm:"column:title"`
	Description        string    `json:"description" gorm:"column:description"`
	Priority           string    `json:"priority" gorm:"column:priority"`
	IsResolved         bool      `json:"is_resolved" gorm:"column:is_resolved"`
	CreatedAt          time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt          time.Time `json:"updated_at" gorm:"column:updated_at"`
	ClosedAt           time.Time `json:"closed_at" gorm:"column:closed_at"`
	ReopenedAt         time.Time `json:"reopened_at" gorm:"column:reopened_at"`
	CreatedByID        uint      `json:"created_by_id" gorm:"column:id"`
	CreatedByUserName  string    `json:"created_by_user_name" gorm:"column:user_name"`
	UpdatedByID        uint      `json:"updated_by_id" gorm:"column:id"`
	UpdatedByUserName  string    `json:"updated_by_user_name" gorm:"column:user_name"`
	ClosedByID         uint      `json:"closed_by_id" gorm:"column:id"`
	ClosedByUserName   string    `json:"closed_by_user_name" gorm:"column:user_name"`
	ReopenedByID       uint      `json:"reopened_by_id" gorm:"column:id"`
	ReopenedByUserName string    `json:"reopened_by_user_name" gorm:"column:user_name"`
	NoteID             uint      `json:"note_id" gorm:"column:id"`
	// BugID              uint      `json:"bug_id" gorm:"column:id"`
	NoteBody       string    `json:"note_body" gorm:"column:body"`
	NoteCreatedAt  time.Time `json:"note_created_at" gorm:"column:created_at"`
	NoteUpdatedAt  time.Time `json:"note_updated_at" gorm:"column:updated_at"`
	AuthorID       uint      `json:"author_id" gorm:"column:id"`
	AuthorUserName string    `json:"author_user_name" gorm:"column:user_name"`
}

type BugDTO1 struct {
	ID                 uint      `json:"id" gorm:"column:id"`
	Title              string    `json:"title" gorm:"column:title"`
	Description        string    `json:"description" gorm:"column:description"`
	Priority           string    `json:"priority" gorm:"column:priority"`
	IsResolved         bool      `json:"is_resolved" gorm:"column:is_resolved"`
	CreatedAt          time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt          time.Time `json:"updated_at" gorm:"column:updated_at"`
	ClosedAt           time.Time `json:"closed_at" gorm:"column:closed_at"`
	ReopenedAt         time.Time `json:"reopened_at" gorm:"column:reopened_at"`
	CreatedByID        uint      `json:"created_by_id" gorm:"column:id"`
	CreatedByUserName  string    `json:"created_by_user_name" gorm:"column:user_name"`
	UpdatedByID        uint      `json:"updated_by_id" gorm:"column:id"`
	UpdatedByUserName  string    `json:"updated_by_user_name" gorm:"column:user_name"`
	ClosedByID         uint      `json:"closed_by_id" gorm:"column:id"`
	ClosedByUserName   string    `json:"closed_by_user_name" gorm:"column:user_name"`
	ReopenedByID       uint      `json:"reopened_by_id" gorm:"column:id"`
	ReopenedByUserName string    `json:"reopened_by_user_name" gorm:"column:user_name"`
	NoteID             uint      `json:"note_id" gorm:"column:id"`
	// BugID              uint      `json:"bug_id" gorm:"column:id"`
	NoteBody       string    `json:"note_body" gorm:"column:body"`
	NoteCreatedAt  time.Time `json:"note_created_at" gorm:"column:created_at"`
	NoteUpdatedAt  time.Time `json:"note_updated_at" gorm:"column:updated_at"`
	AuthorUserID   uint      `json:"author_user_id" gorm:"column:id"`
	AuthorUserName string    `json:"author_user" gorm:"column:user_name"`
}

type BugDTO3 struct {
	ID                 uint      `json:"id" gorm:"column:id"`
	Title              string    `json:"title" gorm:"column:title"`
	Description        string    `json:"description" gorm:"column:description"`
	Priority           string    `json:"priority" gorm:"column:priority"`
	IsResolved         bool      `json:"is_resolved" gorm:"column:is_resolved"`
	CreatedAt          time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt          time.Time `json:"updated_at" gorm:"column:updated_at"`
	ClosedAt           time.Time `json:"closed_at" gorm:"column:closed_at"`
	ReopenedAt         time.Time `json:"reopened_at" gorm:"column:reopened_at"`
	CreatedByID        uint      `json:"created_by_id" gorm:"column:id"`
	CreatedByUserName  string    `json:"created_by_user_name" gorm:"column:user_name"`
	UpdatedByID        uint      `json:"updated_by_id" gorm:"column:id"`
	UpdatedByUserName  string    `json:"updated_by_user_name" gorm:"column:user_name"`
	ClosedByID         uint      `json:"closed_by_id" gorm:"column:id"`
	ClosedByUserName   string    `json:"closed_by_user_name" gorm:"column:user_name"`
	ReopenedByID       uint      `json:"reopened_by_id" gorm:"column:id"`
	ReopenedByUserName string    `json:"reopened_by_user_name" gorm:"column:user_name"`
	NoteID             uint      `json:"note_id" gorm:"column:id"`
	//	BugID              uint      `json:"bug_id" gorm:"column:id"`
	NoteBody       string    `json:"note_body" gorm:"column:body"`
	NoteCreatedAt  time.Time `json:"note_created_at" gorm:"column:created_at"`
	NoteUpdatedAt  time.Time `json:"note_updated_at" gorm:"column:updated_at"`
	AuthorUserID   uint      `json:"author_user_id" gorm:"column:id"`
	AuthorUserName string    `json:"author_user" gorm:"column:user_name"`
}

type BugDTO4 struct {
	ID                 uint      `json:"id" gorm:"column:id"`
	Title              string    `json:"title" gorm:"column:title"`
	Description        string    `json:"description" gorm:"column:description"`
	Priority           string    `json:"priority" gorm:"column:priority"`
	IsResolved         bool      `json:"is_resolved" gorm:"column:is_resolved"`
	CreatedAt          time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt          time.Time `json:"updated_at" gorm:"column:updated_at"`
	ClosedAt           time.Time `json:"closed_at" gorm:"column:closed_at"`
	ReopenedAt         time.Time `json:"reopened_at" gorm:"column:reopened_at"`
	CreatedByID        uint      `json:"created_by_id" gorm:"column:id"`
	CreatedByUserName  string    `json:"created_by_user_name" gorm:"column:user_name"`
	UpdatedByID        uint      `json:"updated_by_id" gorm:"column:id"`
	UpdatedByUserName  string    `json:"updated_by_user_name" gorm:"column:user_name"`
	ClosedByID         uint      `json:"closed_by_id" gorm:"column:id"`
	ClosedByUserName   string    `json:"closed_by_user_name" gorm:"column:user_name"`
	ReopenedByID       uint      `json:"reopened_by_id" gorm:"column:id"`
	ReopenedByUserName string    `json:"reopened_by_user_name" gorm:"column:user_name"`
	NoteID             uint      `json:"note_id" gorm:"column:id"`
	//	BugID              uint      `json:"bug_id" gorm:"column:id"`
	NoteBody       string    `json:"note_body" gorm:"column:body"`
	NoteCreatedAt  time.Time `json:"note_created_at" gorm:"column:created_at"`
	NoteUpdatedAt  time.Time `json:"note_updated_at" gorm:"column:updated_at"`
	AuthorUserID   uint      `json:"author_user_id" gorm:"column:id"`
	AuthorUserName string    `json:"author_user" gorm:"column:user_name"`
}

type BugDTO5 struct {
	ID                 uint      `json:"id" gorm:"column:id"`
	Title              string    `json:"title" gorm:"column:title"`
	Description        string    `json:"description" gorm:"column:description"`
	Priority           string    `json:"priority" gorm:"column:priority"`
	IsResolved         bool      `json:"is_resolved" gorm:"column:is_resolved"`
	CreatedAt          time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt          time.Time `json:"updated_at" gorm:"column:updated_at"`
	ClosedAt           time.Time `json:"closed_at" gorm:"column:closed_at"`
	ReopenedAt         time.Time `json:"reopened_at" gorm:"column:reopened_at"`
	CreatedByID        uint      `json:"created_by_id" gorm:"column:id"`
	CreatedByUserName  string    `json:"created_by_user_name" gorm:"column:user_name"`
	UpdatedByID        uint      `json:"updated_by_id" gorm:"column:id"`
	UpdatedByUserName  string    `json:"updated_by_user_name" gorm:"column:user_name"`
	ClosedByID         uint      `json:"closed_by_id" gorm:"column:id"`
	ClosedByUserName   string    `json:"closed_by_user_name" gorm:"column:user_name"`
	ReopenedByID       uint      `json:"reopened_by_id" gorm:"column:id"`
	ReopenedByUserName string    `json:"reopened_by_user_name" gorm:"column:user_name"`
	NoteID             uint      `json:"note_id" gorm:"column:id"`
	//	BugID              uint      `json:"bug_id" gorm:"column:id"`
	NoteBody       string    `json:"note_body" gorm:"column:body"`
	NoteCreatedAt  time.Time `json:"note_created_at" gorm:"column:created_at"`
	NoteUpdatedAt  time.Time `json:"note_updated_at" gorm:"column:updated_at"`
	AuthorUserID   uint      `json:"author_user_id" gorm:"column:id"`
	AuthorUserName string    `json:"author_user" gorm:"column:user_name"`
}

// GetBugs godoc
// @Summary Get all bugs in a project
// @Description Get all bugs in a project with the specified project ID.
// @Tags Bugs
// @Accept json
// @Produce json
// @Param projectId path int true "ID of the project to fetch bugs from"
// @Success 200 {array} BugDTO "List of bugs in the project"
// @Failure 400 {object} string "Invalid user ID, project ID, or request data"
// @Failure 401 {object} string "Access is denied"
// @Failure 500 {object} string "Failed to fetch bugs"
// @Router /projects/{projectId}/bugs [get]
func GetBugs(c *gin.Context) {
	// Extract user ID from the request
	userIDStr := c.GetString("userid")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Extract project ID from the params
	projectIDStr := c.Param("projectId")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}
	db := config.DB
	// Fetch project members for the given project
	var projectMembers []models.Member
	db.Where("project_id = ?", projectID).Preload("Member").Find(&projectMembers)

	// Check if the requesting user is a member of the project
	isMember := false
	for _, member := range projectMembers {
		if int(member.UserId) == userID {
			isMember = true
			break
		}
	}

	if !isMember {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access is denied."})
		return
	}

	var bugDTO []BugDTO

	if err := db.Table("bugs").
		Select("bugs.id, bugs.title, bugs.description, bugs.priority, bugs.is_resolved, bugs.created_at, bugs.updated_at, bugs.closed_at, bugs.reopened_at,created_by.id , created_by.user_name, updated_by.id , updated_by.user_name, closed_by.id, closed_by.user_name , reopened_by.id, reopened_by.user_name, note.id , note.body, note.created_at , note.updated_at , Author.id ,Author.user_name").
		Joins("LEFT JOIN users as created_by ON bugs.created_by_id = created_by.id").
		Joins("LEFT JOIN users as updated_by ON bugs.updated_by_id = updated_by.id").
		Joins("LEFT JOIN users as closed_by ON bugs.closed_by_id = closed_by.id").
		Joins("LEFT JOIN users as reopened_by ON bugs.reopened_by_id = reopened_by.id").
		Joins("LEFT JOIN notes as note ON bugs.id = note.bug_id").
		Joins("LEFT JOIN users as Author ON note.user_id = Author.id").
		Where("bugs.project_id = ?", projectID).
		Find(&bugDTO).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch all bugs"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"bugs": bugDTO})
}

// CreateBug godoc
// @Summary Create a new bug in a project
// @Description Create a new bug in a project with the specified project ID.
// @Tags Bugs
// @Accept json
// @Produce json
// @Param projectId path int true "ID of the project to create the bug in"
// @Param bugRequest body BugRequest true "Bug details"
// @Success 201 {object} BugDTO1 "Created bug details"
// @Failure 400 {object} string "Invalid user ID, project ID, or request data"
// @Failure 401 {object} string "Access is denied"
// @Failure 500 {object} string "Failed to create bug"
// @Router /projects/{projectId}/bugs [post]
func CreateBug(c *gin.Context) {
	var bugRequest BugRequest
	if err := c.ShouldBindJSON(&bugRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	fmt.Println("step1")

	// Validate bug fields
	errors := utils.ValidateBugFields(bugRequest.Title, bugRequest.Description, bugRequest.Priority)
	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}
	fmt.Println("step2")

	projectIDStr := c.Param("projectId")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}
	fmt.Println("step3")

	userIDStr := c.GetString("userid")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	db := config.DB
	// Fetch project members for the given project
	var projectMembers []models.Member
	db.Where("project_id = ?", projectID).Preload("Member").Find(&projectMembers)

	fmt.Println("step4")

	// Check if the requesting user is a member of the project
	isMember := false
	for _, member := range projectMembers {
		if int(member.UserId) == userID {
			isMember = true
			break
		}
	}
	fmt.Println("step5", isMember)
	if !isMember {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access is denied."})
		return
	}
	// Create a new bug instance
	createdByID := uint(userID)
	newBug := models.Bug{
		Title:       bugRequest.Title,
		Description: bugRequest.Description,
		Priority:    models.Priority(bugRequest.Priority),
		ProjectId:   uint(projectID),
		// ProjectId:    &createdByID,
		CreatedById:  &createdByID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		ReopenedById: nil,
		ClosedById:   nil,
		UpdatedById:  nil,
	}

	fmt.Println("step6 new bug=", newBug)
	// Insert the bug into the database
	// db.Create(&newBug)
	// Save project and its members to the database
	tx := db.Begin()
	if err := tx.Create(&newBug).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Bug"})
		return
	}

	tx.Commit()

	fmt.Println("step7, after creaing bug")

	var bugDTO1 BugDTO1

	if err := db.Table("bugs").
		Select("bugs.id, bugs.title, bugs.description, bugs.priority, bugs.is_resolved, bugs.created_at, bugs.updated_at, bugs.closed_at, bugs.reopened_at,created_by.id , created_by.user_name, updated_by.id , updated_by.user_name, closed_by.id, closed_by.user_name , reopened_by.id, reopened_by.user_name, note.id, note.body, note.created_at , note.updated_at , Author.id ,Author.user_name ").
		Joins("LEFT JOIN users as created_by ON bugs.created_by_id = created_by.id").
		Joins("LEFT JOIN users as updated_by ON bugs.updated_by_id = updated_by.id").
		Joins("LEFT JOIN users as closed_by ON bugs.closed_by_id = closed_by.id").
		Joins("LEFT JOIN users as reopened_by ON bugs.reopened_by_id = reopened_by.id").
		Joins("LEFT JOIN notes as note ON bugs.id = note.bug_id").
		Joins("LEFT JOIN users as Author ON note.user_id = Author.id").
		Where("bugs.id = ?", newBug.ID).
		First(&bugDTO1).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch created bugss"})
		return
	}

	fmt.Println("step7, after fetchiing created bug , =", bugDTO1)

	c.JSON(http.StatusCreated, gin.H{"Created bug": bugDTO1})

}

// UpdateBug godoc
// @Summary Update a bug in a project
// @Description Update a bug in a project with the specified project and bug IDs.
// @Tags Bugs
// @Accept json
// @Produce json
// @Param projectId path int true "ID of the project containing the bug"
// @Param bugId path int true "ID of the bug to update"
// @Param bugUpdateData body BugUpdateData true "Updated bug details"
// @Success 200 {object} BugDTO3 "Updated bug details"
// @Failure 400 {object} string "Invalid user ID, project ID, bug ID, or request data"
// @Failure 401 {object} string "Access is denied"
// @Failure 500 {object} string "Failed to update bug"
// @Router /projects/{projectId}/bugs/{bugId} [put]
func UpdateBug(c *gin.Context) {

	// Extract bug ID and project ID from the request params
	bugIDStr := c.Param("bugId")
	projectIDStr := c.Param("projectId")
	bugID, err := strconv.Atoi(bugIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bug ID"})
		return
	}
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	// Bind request body to input data struct
	var inputData BugUpdateData
	if err := c.ShouldBindJSON(&inputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Validate bug data
	validationErrors := utils.ValidateBugFields(inputData.Title, inputData.Description, inputData.Priority)

	if len(validationErrors) > 0 {
		errorMessage := ""
		for _, errMessage := range validationErrors {
			errorMessage = errMessage
			break
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	/*
		if len(validationErrors) > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErrors[0]})
			return
		}
	*/

	// Extract user ID from the request
	userIDStr := c.GetString("userid")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Fetch project members for the given project
	var projectMembers []models.Member
	db := config.DB
	db.Where("project_id = ?", projectID).Preload("Member").Find(&projectMembers)

	// Check if the requesting user is a member of the project
	isMember := false
	for _, member := range projectMembers {
		if int(member.UserId) == userID {
			isMember = true
			break
		}
	}

	if !isMember {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access is denied."})
		return
	}

	// Fetch the target bug to update
	var targetBug models.Bug

	if err := db.First(&targetBug, bugID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bug ID"})
		return
	}

	// Update bug data
	targetBug.Title = inputData.Title
	targetBug.Description = inputData.Description
	targetBug.Priority = models.Priority(inputData.Priority)
	targetBug.UpdatedBy.ID = uint(userID)
	targetBug.UpdatedAt = time.Now()

	tx := db.Begin()
	// Save the updated bug
	if err := tx.Save(&targetBug).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update bug"})
		return
	}
	tx.Commit()

	// Fetch the updated bug with related data
	// var updatedBug models.Bug
	var bugDTO3 BugDTO3
	if err := db.Table("bugs").
		Select("bugs.id, bugs.title, bugs.description, bugs.priority, bugs.is_resolved, bugs.created_at, bugs.updated_at, bugs.closed_at, bugs.reopened_at,created_by.id , created_by.user_name, updated_by.id , updated_by.user_name, closed_by.id, closed_by.user_name , reopened_by.id, reopened_by.user_name, note.id  , note.body, note.created_at , note.updated_at , Author.id ,Author.user_name ").
		Joins("LEFT JOIN users as created_by ON bugs.created_by_id = created_by.id").
		Joins("LEFT JOIN users as updated_by ON bugs.updated_by_id = updated_by.id").
		Joins("LEFT JOIN users as closed_by ON bugs.closed_by_id = closed_by.id").
		Joins("LEFT JOIN users as reopened_by ON bugs.reopened_by_id = reopened_by.id").
		Joins("LEFT JOIN notes as note ON bugs.id = note.bug_id").
		Joins("LEFT JOIN users as Author ON note.user_id = Author.id").
		Where("bugs.id = ?", targetBug.ID).
		First(&bugDTO3).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated bugs"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Updated bug": bugDTO3})

}

// @Summary Delete a bug
// @Description Delete a bug by bug ID
// @Tags Bugs
// @ID delete-bug
// @Param projectId path int true "Project ID"
// @Param bugId path int true "Bug ID"
// @Security ApiKeyAuth
// @Produce json
// @Success 204 "No Content"
// @Failure 400 {object} string "Bad Request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 403 {object} string "Forbidden"
// @Failure 500 {object} string "Internal Server Error"
// @Router /projects/{projectId}/bugs/{bugId} [delete]
func DeleteBug(c *gin.Context) {
	// Extract project ID and bug ID from the request params
	projectIDStr := c.Param("projectId")
	bugIDStr := c.Param("bugId")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}
	bugID, err := strconv.Atoi(bugIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bug ID"})
		return
	}

	db := config.DB

	// Find the target project using project ID
	var targetProject models.Project
	if err := db.First(&targetProject, projectID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	// Find the target bug using bug ID
	var targetBug models.Bug
	if err := db.First(&targetBug, bugID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bug ID"})
		return
	}

	// Extract user ID from the request
	userIDStr := c.GetString("userid")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Check if the user has permission to delete the bug
	if targetProject.UserId != uint(userID) && *targetBug.CreatedById != uint(userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access is denied."})
		return
	}

	tx := db.Begin()

	// Delete the related notes using the bug ID
	if err := tx.Where("bug_id = ?", bugID).Delete(models.Note{}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete related notes"})
		return
	}

	// Delete the target bug
	if err := tx.Delete(&targetBug).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete bug"})
		return
	}
	tx.Commit()

	c.JSON(http.StatusNoContent, nil)
	// c.JSON(http.StatusNoContent, "Bug has been deleted")
}

// @Summary Close a bug
// @Description Close a bug by bug ID
// @Tags Bugs
// @ID close-bug
// @Param projectId path int true "Project ID"
// @Param bugId path int true "Bug ID"
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} BugDTO4 "Closed bug"
// @Failure 400 {object} string "Bad Request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 403 {object} string "Forbidden"
// @Failure 500 {object} string "Internal Server Error"
// @Router /projects/{projectId}/bugs/{bugId}/close [post]
func CloseBug(c *gin.Context) {
	// Extract project ID and bug ID from the request params
	projectIDStr := c.Param("projectId")
	bugIDStr := c.Param("bugId")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}
	bugID, err := strconv.Atoi(bugIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bug ID"})
		return
	}

	db := config.DB

	// Fetch project members for the given project
	var projectMembers []models.Member
	db.Where("project_id = ?", projectID).Preload("Member").Find(&projectMembers)

	// Extract user ID from the request
	userIDStr := c.GetString("userid")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Check if the user is a member of the project
	isMember := false
	for _, member := range projectMembers {
		if int(member.UserId) == userID {
			isMember = true
			break
		}
	}

	if !isMember {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access is denied."})
		return
	}

	// Find the target bug using bug ID
	var targetBug models.Bug
	if err := db.First(&targetBug, bugID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bug ID"})
		return
	}

	// Check if the bug is already marked as closed
	if targetBug.IsResolved {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bug is already marked as closed."})
		return
	}

	closedid := uint(userID)
	time1 := time.Now()
	// Update bug information to mark it as closed
	targetBug.IsResolved = true
	targetBug.ClosedById = &closedid
	targetBug.ClosedAt = &time1
	// targetBug.ReopenedById = models.User{}
	targetBug.ReopenedById = nil
	// targetBug.ReopenedAt = time.Time{}
	targetBug.ReopenedAt = nil

	tx := db.Begin()

	// Save the updated bug
	if err := tx.Save(&targetBug).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update bug"})
		return
	}
	tx.Commit()

	var bugDTO4 BugDTO4

	if err := db.Table("bugs").
		Select("bugs.id, bugs.title, bugs.description, bugs.priority, bugs.is_resolved, bugs.created_at, bugs.updated_at, bugs.closed_at, bugs.reopened_at,created_by.id , created_by.user_name, updated_by.id , updated_by.user_name, closed_by.id, closed_by.user_name , reopened_by.id, reopened_by.user_name, note.id , note.body, note.created_at , note.updated_at , Author.id ,Author.user_name ").
		Joins("LEFT JOIN users as created_by ON bugs.created_by_id = created_by.id").
		Joins("LEFT JOIN users as updated_by ON bugs.updated_by_id = updated_by.id").
		Joins("LEFT JOIN users as closed_by ON bugs.closed_by_id = closed_by.id").
		Joins("LEFT JOIN users as reopened_by ON bugs.reopened_by_id = reopened_by.id").
		Joins("LEFT JOIN notes as note ON bugs.id = note.bug_id").
		Joins("LEFT JOIN users as Author ON note.user_id = Author.id").
		Where("bugs.id = ?", bugID).
		First(&bugDTO4).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch closed bugs"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Closed bug": bugDTO4})
}

// @Summary Reopen a bug
// @Description Reopen a bug by bug ID
// @Tags Bugs
// @ID reopen-bug
// @Param projectId path int true "Project ID"
// @Param bugId path int true "Bug ID"
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} BugDTO5 "Reopened bug"
// @Failure 400 {object} string "Bad Request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 403 {object} string "Forbidden"
// @Failure 500 {object} string "Internal Server Error"
// @Router /projects/{projectId}/bugs/{bugId}/reopen [post]
func ReOpenBug(c *gin.Context) {
	// Extract project ID and bug ID from the request params
	projectIDStr := c.Param("projectId")
	bugIDStr := c.Param("bugId")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}
	bugID, err := strconv.Atoi(bugIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bug ID"})
		return
	}

	db := config.DB

	// Fetch project members for the given project
	var projectMembers []models.Member
	db.Where("project_id = ?", projectID).Preload("Member").Find(&projectMembers)

	// Extract user ID from the request
	userIDStr := c.GetString("userid")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Check if the user is a member of the project
	isMember := false
	for _, member := range projectMembers {
		if int(member.UserId) == userID {
			isMember = true
			break
		}
	}

	if !isMember {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access is denied."})
		return
	}

	// Find the target bug using bug ID
	var targetBug models.Bug
	if err := db.First(&targetBug, bugID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bug ID"})
		return
	}

	// Check if the bug is already marked as opened
	if !targetBug.IsResolved {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bug is already marked as opened."})
		return
	}

	// Update bug information to mark it as opened
	reopenedid := uint(userID)
	time1 := time.Now()
	targetBug.IsResolved = false
	targetBug.ClosedById = nil
	targetBug.ClosedAt = nil
	targetBug.ReopenedById = &reopenedid
	targetBug.ReopenedAt = &time1

	tx := db.Begin()
	// Save the updated bug
	if err := tx.Save(&targetBug).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update bug"})
		return
	}

	tx.Commit()
	var bugDTO5 BugDTO5
	if err := db.Table("bugs").
		Select("bugs.id, bugs.title, bugs.description, bugs.priority, bugs.is_resolved, bugs.created_at, bugs.updated_at, bugs.closed_at, bugs.reopened_at,created_by.id , created_by.user_name, updated_by.id , updated_by.user_name, closed_by.id, closed_by.user_name , reopened_by.id, reopened_by.user_name, note.id , note.body, note.created_at , note.updated_at , Author.id ,Author.user_name ").
		Joins("LEFT JOIN users as created_by ON bugs.created_by_id = created_by.id").
		Joins("LEFT JOIN users as updated_by ON bugs.updated_by_id = updated_by.id").
		Joins("LEFT JOIN users as closed_by ON bugs.closed_by_id = closed_by.id").
		Joins("LEFT JOIN users as reopened_by ON bugs.reopened_by_id = reopened_by.id").
		Joins("LEFT JOIN notes as note ON bugs.id = note.bug_id").
		Joins("LEFT JOIN users as Author ON note.user_id = Author.id").
		Where("bugs.id = ?", bugID).
		First(&bugDTO5).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reopened bugs"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Reopened bug": bugDTO5})
}
