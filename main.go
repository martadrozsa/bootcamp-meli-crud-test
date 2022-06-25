package main

import (
	"github.com/gin-gonic/gin"
	"github.com/martadrozsa/bootcamp-meli-crud-test/cmd/controller/user"
	modules2 "github.com/martadrozsa/bootcamp-meli-crud-test/internal/user/repository/mysql"
	"github.com/martadrozsa/bootcamp-meli-crud-test/internal/user/service"
)

func main() {

	router := gin.Default()
	group := router.Group("api/")

	userRepository := modules2.CreateUserRepository()
	userService := service.CreateUserService(userRepository)
	userController := user.CreateUserController(userService)

	userGroup := group.Group("/users")
	userGroup.GET("/", userController.GetAll())
	userGroup.GET("/:id", userController.GetById())
	userGroup.POST("/", userController.Create())
	userGroup.PATCH("/:id", userController.UpdateAge())
	userGroup.DELETE("/:id", userController.Delete())

	router.Run()
}
