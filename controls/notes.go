package controls

import (
	"feyin/bug-tracker/config"
	"feyin/bug-tracker/models"
	"feyin/bug-tracker/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Extract request data
type requestBody struct {
	Body string `json:"body" binding:"required"`
}

type NoteDTO struct {
	NoteID         uint      `json:"note_id" gorm:"column:id"`
	BugID          uint      `json:"bug_id" gorm:"column:id"`
	Body           string    `json:"note_body" gorm:"column:body"`
	CreatedAt      time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"column:updated_at"`
	AuthhorID      uint      `json:"author_id" gorm:"column:id"`
	AuthorUserName string    `json:"author_user_name" gorm:"column:user_name"`
}

// @Summary Create a new note
// @Description Create a new note for a bug
// @Tags Notes
// @ID create-note
// @Param projectId path int true "Project ID"
// @Param bugId path int true "Bug ID"
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param request body requestBody true "Note details"
// @Success 201 {object} NoteDTO "Created note"
// @Failure 400 {object} string "Bad Request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 403 {object} string "Forbidden"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /projects/{projectId}/bugs/{bugId}/notes [post]
func PostNote(c *gin.Context) {
	// Extract user ID from the request
	userIDStr := c.GetString("userid")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var req requestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	projectID := c.Param("projectId")
	bugID := c.Param("bugId")

	// Check if the body is empty
	if req.Body == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Note body field must not be empty."})
		return
	}

	db := config.DB
	// Find the project
	var targetProject models.Project
	if err := db.Where("id = ?", projectID).Preload("Members").First(&targetProject).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Project not found"})
		return
	}

	// Check if the current user is a member of the project
	var isMember bool
	for _, member := range targetProject.Members {
		if member.UserId == uint(userID) {
			isMember = true
			break
		}
	}
	if !isMember {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Access is denied. Not a member of the project."})
		return
	}

	// Convert bugID string to uint
	bugIDInt, err := strconv.Atoi(bugID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bug ID"})
		return
	}

	// Create and save the new note
	newNote := models.Note{
		Body: req.Body,
		// Author: models.User{ID: uint(userID)},
		UserId: uint(userID),
		BugId:  uint(bugIDInt),
		//	Bug:    models.Bug{ID: uint(bugIDInt)},
	}

	tx := db.Begin()

	if err := tx.Create(&newNote).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create note"})
		return
	}
	tx.Commit()

	var noteDTO NoteDTO

	if err := db.
		Table("notes").
		Joins("LEFT JOIN users ON notes.user_id = users.id").
		Joins("LEFT JOIN bugs ON notes.bug_id = bugs.id").
		Where("notes.id = ?", newNote.ID).
		Select("notes.id,bugs.id,notes.body,notes.created_at,notes.updated_at, users.id , users.user_name ").
		Scan(&noteDTO).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch created note"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"note": noteDTO})
}

// @Summary Delete a note
// @Description Delete a note by note ID
// @Tags Notes
// @ID delete-note
// @Param projectId path int true "Project ID"
// @Param noteId path int true "Note ID"
// @Security ApiKeyAuth
// @Produce json
// @Success 400 {object} string "Note Deleted successfully"
// @Failure 400 {object} string "Bad Request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 403 {object} string "Forbidden"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /projects/{projectId}/notes/{noteId} [delete]
func DeleteNote(c *gin.Context) {
	projectIDStr := c.Param("projectId")
	noteIDStr := c.Param("noteId")

	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	noteID, err := strconv.Atoi(noteIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	db := config.DB

	var targetProject models.Project
	if err := db.
		Where("id = ?", projectID).
		Preload("Members").
		First(&targetProject).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Invalid project ID"})
		return
	}

	memberIDs := make([]uint, len(targetProject.Members))
	for i, member := range targetProject.Members {
		// memberIDs[i] = member.Member.ID
		memberIDs[i] = member.UserId
	}

	userIDStr := c.GetString("userid")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if !utils.ContainsUint(memberIDs, uint(userID)) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Access is denied. Not a member of the project."})
		return
	}

	var targetNote models.Note
	if err := db.
		Where("id = ?", noteID).
		First(&targetNote).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Invalid note ID"})
		return
	}

	if targetNote.UserId != uint(userID) && targetProject.UserId != uint(userID) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Access is denied."})
		return
	}

	tx := db.Begin()

	if err := tx.Delete(&targetNote).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete note"})
		return
	}
	tx.Commit()

	//	c.Status(http.StatusNoContent)
	c.JSON(400, gin.H{
		"Message": "Note Deleted successfully",
	})
}

type NoteDTO1 struct {
	NoteID         uint      `json:"note_id" gorm:"column:id"`
	BugID          uint      `json:"bug_id" gorm:"column:id"`
	Body           string    `json:"note_body" gorm:"column:body"`
	CreatedAt      time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"column:updated_at"`
	AuthorID       uint      `json:"author_id" gorm:"column:id"`
	AuthorUserName string    `json:"author_user_name" gorm:"column:user_name"`
}

type requestBody1 struct {
	Body string `json:"body" binding:"required"`
}

// @Summary Update a note
// @Description Update a note by note ID
// @Tags Notes
// @ID update-note
// @Param projectId path int true "Project ID"
// @Param noteId path int true "Note ID"
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param request body requestBody1 true "Note details"
// @Success 200 {object} NoteDTO1 "Updated note"
// @Failure 400 {object} string "Bad Request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 403 {object} string "Forbidden"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /projects/{projectId}/notes/{noteId} [put]
func UpdateNote(c *gin.Context) {

	var req requestBody1

	projectIDStr := c.Param("projectId")
	noteIDStr := c.Param("noteId")

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.DB

	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	noteID, err := strconv.Atoi(noteIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	var targetProject models.Project
	if err := db.
		Where("id = ?", projectID).
		Preload("Members").
		First(&targetProject).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Invalid project ID"})
		return
	}

	memberIDs := make([]uint, len(targetProject.Members))
	for i, member := range targetProject.Members {
		//	memberIDs[i] = member.Member.ID
		memberIDs[i] = member.UserId
	}

	userIDStr := c.GetString("userid")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if !utils.ContainsUint(memberIDs, uint(userID)) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Access is denied. Not a member of the project."})
		return
	}

	var targetNote models.Note
	if err := db.
		Where("id = ?", noteID).
		First(&targetNote).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Invalid note ID"})
		return
	}

	if targetNote.UserId != uint(userID) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Access is denied."})
		return
	}

	tx := db.Begin()

	targetNote.Body = req.Body

	if err := tx.Save(&targetNote).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update note"})
		return
	}

	tx.Commit()

	var noteDTO1 NoteDTO1

	if err := db.
		Table("notes").
		Joins("LEFT JOIN users ON notes.user_id = users.id").
		Joins("LEFT JOIN bugs ON notes.bug_id = bugs.id").
		Where("notes.id = ?", targetNote.ID).
		Select("notes.id,bugs.id,notes.body,notes.created_at,notes.updated_at, users.id , users.user_name ").
		Scan(&noteDTO1).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated note"})
		return
	}

	//	c.JSON(http.StatusOK, targetNote)
	c.JSON(http.StatusOK, gin.H{"Updated Note": noteDTO1})
}
