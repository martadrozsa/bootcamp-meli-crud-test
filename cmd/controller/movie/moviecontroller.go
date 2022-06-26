package movie

import (
	"github.com/gin-gonic/gin"
	"github.com/martadrozsa/bootcamp-meli-crud-test/internal/movie/domain"
	"github.com/martadrozsa/bootcamp-meli-crud-test/pkg/httputil"
	"gopkg.in/errgo.v2/fmt/errors"
	"net/http"
	"strconv"
)

type MovieController struct {
	service domain.MovieService
}

type requestMoviePost struct {
	Id    int64  `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Genre string `json:"genre" binding:"required"`
	Year  int    `json:"year" binding:"required"`
	Award int    `json:"award" binding:"required"`
}

type requestMoviePatch struct {
	Award int `json:"award" binding:"required"`
}

func CreateMovieController(movieService domain.MovieService) *MovieController {
	return &(MovieController{service: movieService})
}

func (c MovieController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		movies, err := c.service.GetAll(ctx.Request.Context())

		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, err)
			return
		}

		httputil.NewResponse(ctx, http.StatusOK, movies)
	}
}

func (c *MovieController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		movieId, err := c.service.GetById(ctx.Request.Context(), id)
		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, err)
			return
		}

		httputil.NewResponse(ctx, http.StatusOK, movieId)
	}
}

func (c *MovieController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var movieDto requestMoviePost
		if err := ctx.ShouldBindJSON(&movieDto); err != nil {
			httputil.NewError(ctx, http.StatusUnprocessableEntity, errors.New("Invalid input. Check the data entered"))
			return
		}
		newMovie, err := c.service.Create(ctx.Request.Context(), movieDto.Name, movieDto.Genre, movieDto.Year, movieDto.Award)

		if err != nil {
			httputil.NewError(ctx, http.StatusConflict, err)
			return
		}
		httputil.NewResponse(ctx, http.StatusCreated, newMovie)
	}
}

func (c *MovieController) UpdateAward() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, errors.New("invalid id"))
			return
		}

		var movieDto requestMoviePatch

		if err := ctx.ShouldBindJSON(&movieDto); err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, err)
			return
		}

		if movieDto.Award < 0 {
			httputil.NewError(ctx, http.StatusBadRequest, errors.New("Award cannot have a negative value"))
			return
		}

		movieUpdate, err := c.service.UpdateAward(ctx.Request.Context(), id, movieDto.Award)

		if err != nil {
			httputil.NewError(ctx, http.StatusNotFound, err)
			return
		}
		httputil.NewResponse(ctx, http.StatusOK, &movieUpdate)
	}
}

func (c *MovieController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
			return
		}

		err = c.service.Delete(ctx.Request.Context(), id)
		if err != nil {
			httputil.NewError(ctx, http.StatusNotFound, err)
			return
		}
		httputil.NewResponse(ctx, http.StatusNoContent, err)
	}
}
