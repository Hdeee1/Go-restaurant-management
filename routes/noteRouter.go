package routes

import (
	"github.com/Hdeee1/go-restaurant-management/controllers"
	"github.com/Hdeee1/go-restaurant-management/middleware"
	"github.com/gin-gonic/gin"
)

func NoteRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/notes", middleware.Authentication(), middleware.CheckRole("admin"), controllers.CreateNote())
	incomingRoutes.GET("/notes", middleware.Authentication(), controllers.GetNotes())
	incomingRoutes.GET("/notes/:note_id", middleware.Authentication(), controllers.GetNote())
	incomingRoutes.PATCH("/notes/:note_id", middleware.Authentication(), controllers.UpdateNote())
}
