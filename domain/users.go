package domain

import "time"

type UserStruct struct {
	Id          string
	Username    string
	Email       string
	Phone       string
	DateOfBirth *time.Time
}

type UserDBInterface interface {
	GetAllDb() ([]UserStruct, error)
	InserDb(user *UserStruct) (string, error)
	GetByIdDB(id string) (*UserStruct, error)
	UpdateDb(user *UserStruct) (string, error)
	DeleteUserDB(id string) (string, error)
}
