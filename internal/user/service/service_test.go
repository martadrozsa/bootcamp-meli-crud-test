package service_test

import (
	"fmt"
	"github.com/martadrozsa/bootcamp-meli-crud-test/internal/user/domain"
	"github.com/martadrozsa/bootcamp-meli-crud-test/internal/user/domain/mocks"
	"github.com/martadrozsa/bootcamp-meli-crud-test/internal/user/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserService_Create(t *testing.T) {
	mockRepo := mocks.NewUserRepository(t)

	// INICIALIZACAO / PREPARACAO
	// Testa se o user foi criado com sucesso
	t.Run("create_ok: cria o user com os campos necessários", func(t *testing.T) {
		expectedUser := &domain.User{
			Name:       "Ana",
			Age:        25,
			MovieGenre: "Fantasy",
		}

		// Configura o mock para o teste
		mockRepo.
			// quando o mockRepo/mock for chamado a função "Create" com os argumentos exatos "Ana", 25, "Fantasy"
			On("Create", "Ana", 25, "Fantasy").
			//  Deve retornar (expectedProduct, nil)
			Return(expectedUser, nil).
			// Apenas uma vez (na primeira vez em que for chamado).
			Once()

		// Cria instância do service usando o mockRepo de mock.
		service := service.CreateUserService(mockRepo)

		// EXECUCAO DO TESTE
		// invoca a funcao a  ser testada
		userCreate, err := service.Create("Ana", 25, "Fantasy")

		// VALIDACAO
		assert.Nil(t, err)
		assert.Equal(t, userCreate, expectedUser)
	})

	// Testa se houve algum conflito ao criar o user — Name não pode ser igual
	t.Run("create_conflict: quando o nome já existir, o usuário não será criado", func(t *testing.T) {

		mockRepo.
			On("Create", "Ana", 25, "Fantasy").
			Return(nil, fmt.Errorf("the name has already been registered")).
			Once()

		service := service.CreateUserService(mockRepo)

		expectedUser, err := service.Create("Ana", 25, "Fantasy")

		assert.NotNil(t, err)
		assert.Nil(t, expectedUser)
		assert.Equal(t, err.Error(), "the name has already been registered")
	})
}

func TestUserService_GetAll(t *testing.T) {
	mockRepo := mocks.NewUserRepository(t)

	// Testa se todos os users foram retornados da base
	t.Run("find_all: deve retornar todos os usuários existentes na base", func(t *testing.T) {
		expectedUserList := []domain.User{
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

		mockRepo.
			On("GetAll").
			Return(expectedUserList, nil).
			Once()

		service := service.CreateUserService(mockRepo)

		userList, err := service.GetAll()

		assert.Nil(t, err)
		assert.Equal(t, expectedUserList, userList)

	})

	// Testa se ocorreu algum erro ao buscar por todos os produtos
	t.Run("find_all_err: quando não encontrar todos os usuários, retornará um erro", func(t *testing.T) {

		mockRepo.
			On("GetAll").
			Return([]domain.User{}, fmt.Errorf("error: users not found")).
			Once()

		service := service.CreateUserService(mockRepo)

		_, err := service.GetAll()
		assert.Error(t, err)
	})
}

func TestUserService_GetById(t *testing.T) {
	mockRepo := mocks.NewUserRepository(t)

	// Testa se o ‘id’ do user procurado não for encontrado
	t.Run("find_by_id_non_existent: quando o usuário for procurado por id que não existir, retorna null", func(t *testing.T) {

		mockRepo.
			On("GetById", int64(1)).
			Return(nil, fmt.Errorf("user id was not found")).
			Once()

		service := service.CreateUserService(mockRepo)

		userId, err := service.GetById(int64(1))

		assert.Nil(t, userId)
		assert.NotNil(t, err)
	})

	// Testa se o id do user procurado foi encontrado
	t.Run("find_by_id_existent: quando o user procurado por id existir, retorna as informações do user solicitado", func(t *testing.T) {
		expectedUserList := []*domain.User{
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

		mockRepo.
			On("GetById", int64(1)).
			Return(expectedUserList[1], nil).
			Once()

		service := service.CreateUserService(mockRepo)

		resultUser, err := service.GetById(int64(1))

		assert.Nil(t, err)
		assert.Equal(t, expectedUserList[1], resultUser)

	})
}

func TestUserService_UpdateAge(t *testing.T) {
	mockRepo := mocks.NewUserRepository(t)

	// Testa se a atualização do user atualizado foi bem sucedida
	t.Run("update_existent: Quando a atualização dos dados for bem sucedida, o user será devolvido com as informações atualizadas", func(t *testing.T) {
		expectedUser := &domain.User{
			Id:         1,
			Name:       "Ana",
			Age:        25,
			MovieGenre: "Fantasy",
		}

		mockRepo.
			On("UpdateAge", int64(1), 26).
			Return(expectedUser, nil).
			Once()

		service := service.CreateUserService(mockRepo)

		userUpdate, err := service.UpdateAge(int64(1), 26)

		assert.Nil(t, err)
		assert.Equal(t, userUpdate, expectedUser)

	})

	// Testa se user pesquisado para ser atualizado não existe na base
	t.Run("update_non_existent: quando o user a ser atualizado não existir, será retornado null.", func(t *testing.T) {

		mockRepo.
			On("UpdateAge", int64(1), 26).
			Return(nil, fmt.Errorf("user was not found")).
			Once()

		service := service.CreateUserService(mockRepo)

		userUpdate, err := service.UpdateAge(int64(1), 26)

		assert.Nil(t, userUpdate)
		assert.NotNil(t, err)
	})
}

func TestUserService_Delete(t *testing.T) {
	mockRepo := mocks.NewUserRepository(t)

	// Testa se o ‘id’ do user a ser deletado não for encontrado
	t.Run("delete_non_existent: quando o user não existe, null será retornado.", func(t *testing.T) {

		mockRepo.
			On("Delete", int64(1)).
			Return(fmt.Errorf("user was not found")).
			Once()

		service := service.CreateUserService(mockRepo)

		err := service.Delete(int64(1))

		assert.NotNil(t, err)
	})

	// Testa se o ‘id’ do user a ser deletado for encontrado com sucesso no delete
	t.Run("delete_ok: quando a exclusão for bem-sucedida, o user não aparecerá na lista.", func(t *testing.T) {

		mockRepo.
			On("Delete", int64(1)).
			Return(nil).
			Once()

		service := service.CreateUserService(mockRepo)

		err := service.Delete(int64(1))

		assert.Nil(t, err)
	})
}
