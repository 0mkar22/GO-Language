package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type NoteController struct{}

func (n *NoteController) InitNoteContollerRouter(router *gin.Engine) {
	notes := router.Group("/notes")
	notes.GET("/get", n.GetNotes)
	notes.POST("/post", n.CreateNote)
}

func (n *NoteController) GetNotes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Get all notes"})

}

func (n *NoteController) CreateNote(c *gin.Context) {
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Create a new note"})

}
