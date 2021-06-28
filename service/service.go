package service

import "cassandra_rest_api_users/domain"

type UserRepository struct {
	repo domain.UserDBInterface
}

type UserService interface {
	GetAll() ([]domain.UserStruct, error)
	Insert(user *domain.UserStruct) (string, error)
	GetById(id string) (*domain.UserStruct, error)
	Update(user *domain.UserStruct) (string, error)
	Delete(id string) (string, error)
}

func (r UserRepository) GetAll() ([]domain.UserStruct, error) {
	return r.repo.GetAllDb()
}

func (r UserRepository) Insert(user *domain.UserStruct) (string, error) {
	return r.repo.InserDb(user)
}

func (r UserRepository) GetById(id string) (*domain.UserStruct, error) {
	return r.repo.GetByIdDB(id)
}

func (r UserRepository) Update(user *domain.UserStruct) (string, error) {
	return r.repo.UpdateDb(user)
}

func (r UserRepository) Delete(id string) (string, error) {
	return r.repo.DeleteUserDB(id)
}

func NewUserService(repository domain.UserDBInterface) UserRepository {
	return UserRepository{repo: repository}
}
