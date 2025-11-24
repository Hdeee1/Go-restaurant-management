package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Hdeee1/go-restaurant-management/controllers"
)

func NoteRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/notes", controllers.CreateNote())
	incomingRoutes.GET("/notes", controllers.GetNotes())
	incomingRoutes.GET("/notes/:note_id", controllers.GetNote())
	incomingRoutes.PATCH("/notes/:note_id", controllers.UpdateNote())
}