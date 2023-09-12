package controls

import (
	"feyin/bug-tracker/config"
	"feyin/bug-tracker/models"
	"feyin/bug-tracker/utils"

	// "fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UserEnterData struct {
	Name    string   `json:"name" binding:"required"`
	Members []string `json:"members"`
}

type Result struct {
	ProjectID         uint      `gorm:"column:id"`
	ProjectName       string    `gorm:"column:project_name"`
	CreatedAt         time.Time `gorm:"column:created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at"`
	CreatedByID       uint      `gorm:"column:id"`
	CreatedByUserName string    `gorm:"column:user_name"` // Rename this field to match the column name
	MemberID          uint      `gorm:"column:id"`        // Rename this field to match the column name
	JoinedAt          time.Time `gorm:"column:joined_at"` // Rename this field to match the column name
	MemberUserID      uint      `gorm:"column:id"`        // Rename this field to match the column name
	MemberUserName    string    `gorm:"column:user_name"` // Rename this field to match the column name
	BugID             uint      `gorm:"column:id"`        // Rename this field to match the column name
}

// GetProjects godoc
// @Summary Get projects associated with the authenticated user
// @Description Get a list of projects associated with the authenticated user.
// @Tags projects
// @Accept json
// @Produce json
// @Success 200 {object} Result "List of projects"
// @Failure 400 {object} string "Invalid user ID"
// @Failure 500 {object} string "Failed to fetch projects"
// @Router /projects/ [get]
func GetProjects(c *gin.Context) {

	// Extract user ID from the request
	userIDStr := c.GetString("userid")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// var projects []models.Project

	db := config.DB
	// Fetch projects with specified fields and related data

	var result []Result

	db.Table("projects").
		Select("projects.id, projects.project_name, projects.created_at, projects.updated_at, created_by.id, created_by.user_name, members.id, members.joined_at, member.id, member.user_name, bugs.id").
		Joins("JOIN users as created_by ON projects.user_id = created_by.id").
		Joins("LEFT JOIN members ON members.project_id = projects.id").
		Joins("LEFT JOIN users as member ON members.user_id = member.id").
		Joins("LEFT JOIN bugs ON bugs.project_id = projects.id").
		Where("members.user_id = ?", userID).
		//	Where("member.id = ?", userID).
		Scan(&result)
		//Find(&result)

	c.JSON(http.StatusOK, gin.H{"projects": result})

}

type Result2 struct {
	ProjectID         uint      `gorm:"column:id"`
	ProjectName       string    `gorm:"column:project_name"`
	CreatedAt         time.Time `gorm:"column:created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at"`
	CreatedByID       uint      `gorm:"column:id"`        // Rename this field to match the column name
	CreatedByUserName string    `gorm:"column:user_name"` // Rename this field to match the column name
	MemberID          uint      `gorm:"column:id"`        // Rename this field to match the column name
	JoinedAt          time.Time `gorm:"column:joined_at"` // Rename this field to match the column name
	MemberUserID      uint      `gorm:"column:id"`        // Rename this field to match the column name
	MemberUserName    string    `gorm:"column:user_name"` // Rename this field to match the column name
	BugID             uint      `gorm:"column:id"`        // Rename this field to match the column name
}

// CreateProject godoc
// @Summary Create a new project
// @Description Create a new project and assign members to it.
// @Tags projects
// @Accept json
// @Produce json
// @Param projectInfo body UserEnterData true "Project details"
// @Success 200 {object} Result2 "Created project details"
// @Failure 400 {object} string "Invalid request data"
// @Failure 500 {object} string "Failed to create project"
// @Router /projects/ [post]
func CreateProject(c *gin.Context) {

	var inputData UserEnterData
	if err := c.ShouldBindJSON(&inputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Extract user ID from the request
	userIDStr := c.GetString("userid")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	db := config.DB

	// Create a new project
	newProject := models.Project{
		ProjectName: inputData.Name,
		UserId:      uint(userID), // Set the UserID to the ID of the user creating the project
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Check if the project name is unique

	// Create project members
	var members []models.Member
	for _, memberIDStr := range inputData.Members {
		memberID, err := strconv.Atoi(memberIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid member ID"})
			return
		}

		// Ensure that when creating a new member, the foreign key (ProjectID) is set
		members = append(members, models.Member{
			UserId:    uint(memberID),
			ProjectId: newProject.ID, // Set the ProjectID to the ID of the new project
			JoinedAt:  time.Now(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
	}

	// Save project and its members to the database
	tx := db.Begin()
	if err := tx.Create(&newProject).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project"})
		return
	}

	for i := range members {
		members[i].ProjectId = newProject.ID // Assign the project to the member
		if err := tx.Create(&members[i]).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project members"})
			return
		}
	}

	tx.Commit()

	var result []Result2

	if err := db.Table("projects").
		Select("projects.id, projects.project_name, projects.created_at, projects.updated_at, created_by.id, created_by.user_name, members.id, members.joined_at, member.id, member.user_name, bugs.id").
		Joins("JOIN users as created_by ON projects.user_id = created_by.id").
		Joins("LEFT JOIN members ON members.project_id = projects.id").
		Joins("LEFT JOIN users as member ON members.user_id = member.id").
		Joins("LEFT JOIN bugs ON bugs.project_id = projects.id").
		Where("projects.id = ?", newProject.ID).
		Scan(&result).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch created project"})
		return
	}

	// c.JSON(http.StatusOK, gin.H{"project": fetchedProject})
	c.JSON(http.StatusOK, gin.H{"project": result})
}

// Bind request JSON data
type InputData struct {
	Name string `json:"name" binding:"required"`
}

type EditedProjectResponse struct {
	Message        string  `json:"Message"`
	UpdatedProject Project `json:"Updated Project"`
}

type Project struct {
	ProjectName string    `json:"Project Name"`
	ProjectID   uint      `json:"Project ID"`
	UpdatedAt   time.Time `json:"Updated At"`
	CreatedAt   time.Time `json:"Created At"`
}

// EditProjectByName godoc
// @Summary Edit a project's name
// @Description Edit the name of a project with the specified ID.
// @Tags projects
// @Accept json
// @Produce json
// @Param projectId path int true "ID of the project to edit"
// @Param projectInfo body InputData true "New project name"
// @Success 200 {object} EditedProjectResponse "Successfully updated the project"
// @Failure 400 {object} string "Invalid user ID or request data"
// @Failure 403 {object} string "Permission denied"
// @Failure 404 {object} string "Project not found"
// @Failure 500 {object} string "Failed to update project"
// @Router /projects/{projectId} [put]
func EditProjectByName(c *gin.Context) {
	var userdata InputData
	// Extract user ID from the request
	userIDStr := c.GetString("userid")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Extract project ID from the params
	projectIDStr := c.Param("projectId")
	//  no need to convert it from string gorm functions automatically converts it from any type
	//projectid = uint(projectIDStr)

	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	if err := c.ShouldBindJSON(&userdata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Validate project name
	nameError := utils.ProjectNameError(userdata.Name)
	if nameError != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": nameError})
		return
	}
	db := config.DB
	// Fetch the project from the database
	var project models.Project
	result := db.First(&project, projectID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}
	// Check if the user has permission to edit the project
	if int(project.UserId) != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to edit this project"})
		return
	}

	// Update the project name
	project.ProjectName = userdata.Name
	result = db.Save(&project)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update project"})
		return
	}
	result2 := db.Preload("CreatedBy").
		Preload("Members").
		Preload("Bugs").
		First(&project, projectID)

	if result2.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}
	// c.JSON(http.StatusOK, gin.H{"project": project})
	// c.JSON(http.StatusOK, gin.H{"project": result2})
	/*
		c.JSON(200, gin.H{
			"Message": "Successfully Updated the Project",
			"Updated Project ": gin.H{
				"Project Name": project.ProjectName,
				"Project ID":   project.ID,
				"Updated At":   project.UpdatedAt,
				"CreatedAt":    project.CreatedAt,
			},
		})
	*/

	c.JSON(200, EditedProjectResponse{
		Message: "Successfully Updated the Project",
		UpdatedProject: Project{
			ProjectName: project.ProjectName,
			ProjectID:   project.ID,
			UpdatedAt:   project.UpdatedAt,
			CreatedAt:   project.CreatedAt,
		},
	})

}

// DeleteProject godoc
// @Summary Delete a project
// @Description Delete a project with the specified ID, including its members and bugs.
// @Tags projects
// @Accept json
// @Produce json
// @Param projectId path int true "ID of the project to delete"
// @Success 200 {object} string "Deletion confirmation"
// @Failure 400 {object} string "Invalid user ID or project ID"
// @Failure 403 {object} string "Permission denied"
// @Failure 404 {object} string "Project not found"
// @Router /projects/{projectId} [delete]
func DeleteProject(c *gin.Context) {
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
	// Fetch the project from the database
	var project models.Project
	result := db.First(&project, projectID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	// Check if the user has permission to delete the project
	if project.UserId != uint(userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to delete this project"})
		return
	}

	// Delete related data: members and bugs
	db.Where("project_id = ?", projectID).Delete(models.Member{})
	db.Where("project_id = ?", projectID).Delete(models.Bug{})

	// Delete the project
	db.Delete(&project)
	c.JSON(400, gin.H{
		"Message": "Project Deleted successfully",
	})

}
