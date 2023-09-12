package routes

import (
	"feyin/bug-tracker/controls"
	"feyin/bug-tracker/middleware"

	"github.com/gin-gonic/gin"
)

// the purpose of the UserRouts function is to create a route group for user-related operations under the /user path

func ProjectRoute(c *gin.Engine) {
	project := c.Group("/projects")
	{

		//bug routes

		project.GET("/:projectId/bugs", middleware.UserAuth, controls.GetBugs)
		project.POST("/:projectId/bugs", middleware.UserAuth, controls.CreateBug)
		project.PUT("/:projectId/bugs/:bugId", middleware.UserAuth, controls.UpdateBug)
		project.DELETE("/:projectId/bugs/:bugId", middleware.UserAuth, controls.DeleteBug)
		project.POST("/:projectId/bugs/:bugId/close", middleware.UserAuth, controls.CloseBug)
		project.POST("/:projectId/bugs/:bugId/reopen", middleware.UserAuth, controls.ReOpenBug)

		//project routes

		project.GET("/", middleware.UserAuth, controls.GetProjects)
		project.POST("/", middleware.UserAuth, controls.CreateProject)
		project.PUT("/:projectId", middleware.UserAuth, controls.EditProjectByName)
		project.DELETE("/:projectId", middleware.UserAuth, controls.DeleteProject)

		// note routes

		project.POST("/:projectId/bugs/:bugId/notes", middleware.UserAuth, controls.PostNote)
		project.DELETE("/:projectId/notes/:noteId", middleware.UserAuth, controls.DeleteNote)
		project.PUT("/:projectId/notes/:noteId", middleware.UserAuth, controls.UpdateNote)

		// members routes

		project.POST("/:projectId/members", middleware.UserAuth, controls.AddProjectMembers)
		project.DELETE("/:projectId/members/:memberId", middleware.UserAuth, controls.RemoveProjectMember)
		project.POST("/:projectId/members/leave", middleware.UserAuth, controls.LeaveProjectAsMember)

	}

}
