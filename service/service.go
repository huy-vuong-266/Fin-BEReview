package service

type UserServiceInterface interface {
	AddFund(user interface{}, amount int64) error
	Withdraw(user interface{}, amount int64) error
	GetUserByID(userID string) (interface{}, error)
}

type Service struct {
	UserService UserServiceInterface
}
