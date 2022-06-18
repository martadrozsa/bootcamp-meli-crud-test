package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/martadrozsa/bootcamp-meli-crud-test/internal/domains/user"
	"net/http"
	"strconv"
)

type requestUserPost struct {
	Name       string `json:"name" binding:"required"`
	Age        int    `json:"age" binding:"required"`
	MovieGenre string `json:"movie_genre" binding:"required"`
}

type requestUserPatch struct {
	Age int `json:"age" binding:"required"`
}

type UserController struct {
	service user.UserService
}

func CreateUserController(userService user.UserService) *UserController {
	return &(UserController{service: userService})
}

func (c UserController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users, err := c.service.GetAll()

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
		}
		ctx.JSON(http.StatusOK, users)
	}
}

func (c UserController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
		}

		user, err := c.service.GetById(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, err)
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}

func (c UserController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var requestUser requestUserPost

		if err := ctx.ShouldBindJSON(&requestUser); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": "invalid imput. Check the data entered",
			})
			return
		}

		newUser, err := c.service.Create(requestUser.Name, requestUser.Age, requestUser.MovieGenre)

		if err != nil {
			ctx.JSON(http.StatusConflict, err)
		}

		ctx.JSON(http.StatusCreated, newUser)
	}
}