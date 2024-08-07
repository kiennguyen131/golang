package service

import "ecommerce-backend-api/init/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewNewService() *UserService {
	return &UserService{
		userRepo: repo.NewNewRepo(),
	}
}

// user service u
func (us *UserService) GetInfoUser() string {
	return us.userRepo.GetInfoUser()
}
