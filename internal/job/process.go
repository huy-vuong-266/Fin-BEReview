package job

import (
	"Fin-BEReview/constants"
	"Fin-BEReview/model"
	"Fin-BEReview/service"
	"Fin-BEReview/storage"
	"encoding/json"
	"fmt"
)

type ServiceJob struct {
	JobRing model.JobManager
}

func NewServiceJob() service.JobServiceInterface {
	return &ServiceJob{
		JobRing: model.JobManager{
			JobQueue: make(chan model.Job, 10),
		},
	}
}

func (j *ServiceJob) AddJob(jobRequest interface{}) error {
	job, ok := jobRequest.(*model.Job)
	if !ok {
		return fmt.Errorf("can't parse job")
	}
	select {
	case j.JobRing.JobQueue <- model.Job{
		Key:   job.Key,
		Value: job.Value,
	}:
		return nil
	default:
		redisResponse := storage.Redis.LPush(job.Key, job.Value)
		ok, err := redisResponse.Result()
		if ok != 1 || err != nil {
			return err
		}
		return nil
	}

}

func (j *ServiceJob) ProcessJob(s service.Service) error {
	job := <-j.JobRing.JobQueue
	switch job.Key {
	case constants.FinAddFundJobKey:
		var request model.FinRequest
		err := json.Unmarshal([]byte(job.Value), &request)
		if err != nil {
			return nil
		}

		err = s.UserService.AddFund(request.UserID, request.Amount)
		return err
	case constants.FinWithdrawJobKey:
		var request model.FinRequest
		err := json.Unmarshal([]byte(job.Value), &request)
		if err != nil {
			return nil
		}
		err = s.UserService.Withdraw(request.UserID, request.Amount)
		return err
	default:
		return fmt.Errorf("invalid job key")
	}
}
