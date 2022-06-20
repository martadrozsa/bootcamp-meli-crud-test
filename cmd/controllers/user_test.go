package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/martadrozsa/bootcamp-meli-crud-test/cmd/controllers"
	"github.com/martadrozsa/bootcamp-meli-crud-test/internal/domains/user"
	"github.com/martadrozsa/bootcamp-meli-crud-test/internal/domains/user/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

// simula uma requisição HTTP para a engine de rotas (gin) com o método HTTP, path e body  passados. Retorna a resposta da API.
func ExecuteTestRequest(router *gin.Engine, method string, path string, body []byte) *httptest.ResponseRecorder {

	request := httptest.NewRequest(method, path, bytes.NewBuffer(body))

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	return response
}

func TestProductController_Create(t *testing.T) {

	mockService := mocks.NewUserService(t)

	expectedUser := user.User{
		Id:         1,
		Name:       "Ana",
		Age:        25,
		MovieGenre: "Fantasy",
	}

	body := user.User{
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

		controller := controllers.CreateUserController(mockService)

		requestBody, _ := json.Marshal(body)

		router := SetUpRouter()
		router.POST("/api/users", controller.Create())

		// EXECUÇÃO
		response := ExecuteTestRequest(router, "POST", "/api/users", requestBody)

		// VALIDAÇÃO
		assert.Equal(t, http.StatusCreated, response.Code)

	})

	t.Run("create_fail: quando o objeto JSON não contiver os campos necessários, um código 422 será retornado.", func(t *testing.T) {

		controller := controllers.CreateUserController(nil)

		router := SetUpRouter()
		router.POST("/api/users", controller.Create())

		// EXECUÇÃO
		response := ExecuteTestRequest(router, "POST", "/api/users", []byte{})

		// VALIDAÇÃO
		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
		assert.Equal(t, "{\"message\":\"invalid imput. Check the data entered\"}", response.Body.String())
	})

	t.Run("create_conflict: quando o nome de user já existir, ele retornará um erro 409 Conflict.", func(t *testing.T) {

		// PREPARAÇÃO
		expectedError := errors.New("the name has already been registered")
		mockService.
			On("Create", expectedUser.Name, expectedUser.Age, expectedUser.MovieGenre).
			Return(nil, expectedError).
			Once()

		controller := controllers.CreateUserController(mockService)

		requestBody, _ := json.Marshal(body)

		// configura engine de rotas
		router := SetUpRouter()
		router.POST("/api/users", controller.Create())

		// EXECUÇÃO
		response := ExecuteTestRequest(router, "POST", "/api/users", requestBody)

		// VALIDAÇÃO
		assert.Equal(t, http.StatusConflict, response.Code)
		//assert.Equal(t, "{\"code\":409,\"message\":\"the name has already been registered\"}", response.Body.String())
	})
}
