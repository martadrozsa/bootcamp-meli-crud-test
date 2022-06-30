package main

import (
	"github.com/gin-gonic/gin"
	"github.com/martadrozsa/bootcamp-meli-crud-test/config"

	"github.com/martadrozsa/bootcamp-meli-crud-test/cmd/controller/movie"
	mysqlMovie "github.com/martadrozsa/bootcamp-meli-crud-test/internal/movie/repository/mysql"
	serviceMovie "github.com/martadrozsa/bootcamp-meli-crud-test/internal/movie/service"
)

func main() {

	db := config.ConnectDb()
	defer db.Close()

	router := gin.Default()
	group := router.Group("api/")

	//userRepository := mysql.CreateUserRepository(db)
	//userService := service.CreateUserService(userRepository)
	//userController := user.CreateUserController(userService)

	//userGroup := group.Group("/users")
	//userGroup.GET("/", userController.GetAll())
	//userGroup.GET("/:id", userController.GetById())
	//userGroup.POST("/", userController.Create())
	//userGroup.PATCH("/:id", userController.UpdateAge())
	//userGroup.DELETE("/:id", userController.Delete())

	movieRepository := mysqlMovie.CreateMovieRepository(db)
	movieService := serviceMovie.CreateMovieService(movieRepository)
	movieController := movie.CreateMovieController(movieService)

	movieGroup := group.Group("/movies")
	movieGroup.GET("/", movieController.GetAll())
	movieGroup.GET("/:id", movieController.GetById())
	movieGroup.POST("/", movieController.Create())
	movieGroup.PATCH("/:id", movieController.UpdateAward())
	movieGroup.DELETE("/:id", movieController.Delete())

	router.Run()
}
