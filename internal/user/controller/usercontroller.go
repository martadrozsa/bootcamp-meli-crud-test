package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/martadrozsa/bootcamp-meli-crud-test/internal/user/domain"
	"github.com/martadrozsa/bootcamp-meli-crud-test/pkg/httputil"
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
	service domain.UserService
}

func CreateUserController(userService domain.UserService) *UserController {
	return &(UserController{service: userService})
}

func (c UserController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users, err := c.service.GetAll()

		if err != nil {
			httputil.NewError(ctx, http.StatusInternalServerError, err)
		}
		httputil.NewResponse(ctx, http.StatusOK, users)
	}
}

func (c UserController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, err)
		}

		userId, err := c.service.GetById(id)
		if err != nil {
			httputil.NewError(ctx, http.StatusNotFound, err)
			return
		}

		httputil.NewResponse(ctx, http.StatusOK, userId)
	}
}

func (c UserController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var requestUser requestUserPost

		if err := ctx.ShouldBindJSON(&requestUser); err != nil {
			httputil.NewError(ctx, http.StatusUnprocessableEntity, errors.New("invalid input. Check the data entered"))
			return
		}

		newUser, err := c.service.Create(requestUser.Name, requestUser.Age, requestUser.MovieGenre)

		if err != nil {
			httputil.NewError(ctx, http.StatusConflict, err)
		}

		httputil.NewResponse(ctx, http.StatusCreated, newUser)
	}
}

func (c UserController) UpdateAge() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, err)
		}

		var requestUser requestUserPatch
		if err := ctx.ShouldBindJSON(&requestUser); err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, err)
			return
		}

		if requestUser.Age <= 0 {
			httputil.NewError(ctx, http.StatusBadRequest, errors.New("invalid input. Check the data entered"))
			return
		}

		userUpdate, err := c.service.UpdateAge(id, requestUser.Age)
		if err != nil {
			httputil.NewError(ctx, http.StatusNotFound, err)
			return
		}

		httputil.NewResponse(ctx, http.StatusOK, userUpdate)
	}
}

func (c UserController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, errors.New("invalid id"))
			return
		}

		err = c.service.Delete(id)
		if err != nil {
			httputil.NewError(ctx, http.StatusNotFound, err)
			return
		}
		httputil.NewResponse(ctx, http.StatusNoContent, err)
	}
}
