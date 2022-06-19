package main

import (
	"github.com/gin-gonic/gin"
	"github.com/martadrozsa/bootcamp-meli-crud-test/cmd/controller"
	"github.com/martadrozsa/bootcamp-meli-crud-test/internal/domains/user"
)

func main() {

	router := gin.Default()
	group := router.Group("api/")

	userRepository := user.CreateUserRepository()
	userService := user.CreateUserService(userRepository)
	userController := controller.CreateUserController(userService)

	userGroup := group.Group("/users")
	userGroup.GET("/", userController.GetAll())
	userGroup.GET("/:id", userController.GetById())
	userGroup.POST("/", userController.Create())
	userGroup.PATCH("/:id", userController.UpdateAge())
	userGroup.DELETE("/:id", userController.Delete())

	router.Run()
}
