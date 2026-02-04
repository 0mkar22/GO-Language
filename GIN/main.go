package main

import (
	"example/Gin_Web_Framework/controllers"

	"github.com/gin-gonic/gin"
)

// func Pong(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, gin.H{"message": "pong"})

// }

// func number(c *gin.Context) {
// 	id := c.Param("id")
// 	new_id := c.Param("new_id")
// 	c.IndentedJSON(http.StatusOK, gin.H{"number": id})
// 	c.IndentedJSON(http.StatusOK, gin.H{"new_number": new_id})

// }

// func email(c *gin.Context) {

// 	type Email struct {
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 	}
// 	var email Email
// 	if err := c.BindJSON(&email); err != nil {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
// 		return
// 	}
// 	c.IndentedJSON(http.StatusOK, gin.H{"email": email.Email, "password": email.Password})

// }

// func remove(c *gin.Context) {
// 	id := c.Param("id")

// 	if id == "" {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id parameter is required"})
// 		return
// 	}
// 	c.IndentedJSON(http.StatusOK, gin.H{"removed_id": id})
// }

func main() {
	// router.GET("/number/:id/:new_id", number)
	// router.GET("/ping", Pong)
	// router.POST("/email", email)
	// router.PUT("/replacewhole", email)
	// router.PATCH("/replacesome", email)
	// router.DELETE("/delete/:id", remove)
	router := gin.Default()

	noteContoller := &controllers.NoteController{}
	noteContoller.InitNoteContollerRouter(router)
	router.Run(":8080")
}
