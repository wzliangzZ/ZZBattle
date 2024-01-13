package service

import "github.com/wzliangzZ/ZZBatle/apps/database/models"

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetUserById(id string) (user *models.User, err error) {
	return
}
func (s *UserService) AddUser(user *models.User) (err error) {
	return
}

func (s *UserService) UpdatUserById(id string, userUpdate *models.User) (user *models.User, err error) {
	return
} 
