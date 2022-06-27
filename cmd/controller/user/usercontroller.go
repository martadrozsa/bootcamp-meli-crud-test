package user

import (
	"github.com/gin-gonic/gin"
	"github.com/martadrozsa/bootcamp-meli-crud-test/internal/user/domain"
	"github.com/martadrozsa/bootcamp-meli-crud-test/pkg/httputil"
	"net/http"
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
		users, err := c.service.GetAll(ctx.Request.Context())

		if err != nil {
			httputil.NewError(ctx, http.StatusInternalServerError, err)
		}
		httputil.NewResponse(ctx, http.StatusOK, users)
	}
}
