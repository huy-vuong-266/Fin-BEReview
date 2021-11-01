package service

type UserServiceInterface interface {
	AddFund(userID string, amount int64) error
	Withdraw(userID string, amount int64) error
	GetUserByID(userID string) (interface{}, error)
}
type JobServiceInterface interface {
	AddJob(interface{}) error
	ProcessJob(s Service) error
}

type Service struct {
	UserService UserServiceInterface
	JobService  JobServiceInterface
}
