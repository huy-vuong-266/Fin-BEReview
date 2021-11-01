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

func (s *ServiceUser) AddFund(userID string, amount int64) error {
	userRes, err := s.GetUserByID(userID)
	if err != nil {
		return err
	}
	user := userRes.(*model.User)
	err = storage.DB.Debug().Table("users").Where("user_id = ?", userID).Update("budget", user.Budget+amount).Error
	return err
}
func (s *ServiceUser) Withdraw(userID string, amount int64) error {
	userRes, err := s.GetUserByID(userID)
	if err != nil {
		return err
	}
	u := userRes.(*model.User)
	err = storage.DB.Table("users").Where("user_id = ?", userID).Update("budget", u.Budget-amount).Error
	return err
}
func (s *ServiceUser) GetUserByID(userID string) (interface{}, error) {
	var user model.User

	err := storage.DB.Where("user_id = ?", userID).Find(&user).Error

	return &user, err
}
