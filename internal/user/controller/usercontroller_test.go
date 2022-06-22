package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/martadrozsa/bootcamp-meli-crud-test/internal/user/domain"
	"github.com/martadrozsa/bootcamp-meli-crud-test/internal/user/domain/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const ENDPOINT = "/api/users"

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

// Simula uma requisição HTTP para a engine de rotas (gin) com o método HTTP, path e body  passados. Retorna a resposta da API.
func ExecuteTestRequest(router *gin.Engine, method string, path string, body []byte) *httptest.ResponseRecorder {

	request := httptest.NewRequest(method, path, bytes.NewBuffer(body))

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	return response
}

func TestProductController_Create(t *testing.T) {

	mockService := mocks.NewUserService(t)

	expectedUser := domain.User{
		Id:         1,
		Name:       "Ana",
		Age:        25,
		MovieGenre: "Fantasy",
	}

	body := domain.User{
		Name:       "Ana",
		Age:        25,
		MovieGenre: "Fantasy",
	}

	t.Run("create_ok: quando a entrada de dados for bem-sucedida, um código 201 será retornado junto com o objeto inserido", func(t *testing.T) {

		// PREPARAÇÃO
		mockService.
			On("Create", expectedUser.Name, expectedUser.Age, expectedUser.MovieGenre).
			Return(&expectedUser, nil).
			Once()

		controller := CreateUserController(mockService)

		requestBody, _ := json.Marshal(body)

		router := SetUpRouter()
		router.POST(ENDPOINT, controller.Create())

		// EXECUÇÃO
		response := ExecuteTestRequest(router, http.MethodPost, ENDPOINT, requestBody)

		// VALIDAÇÃO
		assert.Equal(t, http.StatusCreated, response.Code)
		assert.JSONEq(t, "{\"data\":{\"id\":1,\"name\":\"Ana\",\"age\":25,\"movie_genre\":\"Fantasy\"}}", response.Body.String())
	})

	t.Run("create_fail: quando o objeto JSON não contiver os campos necessários, um código 422 será retornado.", func(t *testing.T) {

		controller := CreateUserController(nil)

		router := SetUpRouter()
		router.POST(ENDPOINT, controller.Create())

		// EXECUÇÃO
		response := ExecuteTestRequest(router, http.MethodPost, ENDPOINT, []byte{})

		// VALIDAÇÃO
		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
		assert.Equal(t, "{\"code\":422,\"message\":\"invalid input. Check the data entered\"}", response.Body.String())
	})

	t.Run("create_conflict: quando o nome de user já existir, ele retornará um erro 409 Conflict.", func(t *testing.T) {

		// PREPARAÇÃO
		expectedError := errors.New("the name has already been registered")
		mockService.
			On("Create", expectedUser.Name, expectedUser.Age, expectedUser.MovieGenre).
			Return(nil, expectedError).
			Once()

		controller := CreateUserController(mockService)

		requestBody, _ := json.Marshal(body)

		// configura engine de rotas
		router := SetUpRouter()
		router.POST(ENDPOINT, controller.Create())

		// EXECUÇÃO
		response := ExecuteTestRequest(router, http.MethodPost, ENDPOINT, requestBody)

		// VALIDAÇÃO
		assert.Equal(t, http.StatusConflict, response.Code)
		assert.Equal(t, "{\"code\":409,\"message\":\"the name has already been registered\"}{\"data\":null}", response.Body.String())
	})
}

func TestUserController_GetAll(t *testing.T) {

	mockService := mocks.NewUserService(t)
	/*
		expectedUser := []user.User{
			{
				Id:         1,
				Name:       "Ana",
				Age:        25,
				MovieGenre: "Fantasy",
			},
			{
				Id:         2,
				Name:       "Joana",
				Age:        28,
				MovieGenre: "Drama",
			},
		}
	*/
	body := []domain.User{
		{
			Id:         1,
			Name:       "Ana",
			Age:        25,
			MovieGenre: "Fantasy",
		},
		{
			Id:         2,
			Name:       "Joana",
			Age:        28,
			MovieGenre: "Drama",
		},
	}

	t.Run("find_all_internal_server_error: quando a solicitação não for bem-sucedida, o back-end retornará um erro 500 InternalServerError.", func(t *testing.T) {

		// PREPARAÇÃO
		expectedError := errors.New("the request sent to the server is invalid or corrupted")
		mockService.
			On("GetAll").
			Return(nil, expectedError).
			Once()

		controller := CreateUserController(mockService)

		requestBody, _ := json.Marshal(body)

		// configura engine de rotas
		router := SetUpRouter()
		router.GET(ENDPOINT, controller.GetAll())

		response := ExecuteTestRequest(router, http.MethodGet, ENDPOINT, requestBody)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.Equal(t, "{\"code\":500,\"message\":\"the request sent to the server is invalid or corrupted\"}{\"data\":null}", response.Body.String())
	})
}
