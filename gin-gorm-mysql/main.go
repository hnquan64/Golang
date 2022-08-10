package main

import (
	"gingormsql/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := SetupRouter()
	_ = r.Run(":8080")
	// if ok != nil {
	// 	fmt.Println("Server die")
	// }
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	userRepo := controllers.New()

	r.POST("/users", userRepo.CreateUser)
	// r.GET("/users", userRepo.GetUsers)
	r.GET("/users/:id", userRepo.GetUser)
	r.PUT("/users", userRepo.UpdateUser)
	r.DELETE("/users/:id", userRepo.DeleteUser)

	return r
}
