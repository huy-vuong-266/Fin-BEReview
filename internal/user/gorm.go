package user

import (
	model "Fin-BEReview/model"
	"Fin-BEReview/service"
	"Fin-BEReview/storage"
)

type ServiceUser struct{}

func NewServiceUser() service.UserServiceInterface {
	return &ServiceUser{}
}

func (s *ServiceUser) AddFund(user interface{}, amount int64) error {
	u := user.(*model.User)
	err := storage.DB.Table("users").Where("user_id = ?", u.UserID).Update("budget", u.Budget+amount).Error
	return err
}
func (s *ServiceUser) Withdraw(user interface{}, amount int64) error {
	u := user.(*model.User)
	err := storage.DB.Table("users").Where("user_id = ?", u.UserID).Update("budget", u.Budget-amount).Error
	return err
}
func (s *ServiceUser) GetUserByID(userID string) (interface{}, error) {
	var user model.User

	err := storage.DB.Where("user_id = ?", userID).Find(&user).Error

	return &user, err
}
