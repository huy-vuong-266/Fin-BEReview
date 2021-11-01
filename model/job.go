package model

type Job struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type JobManager struct {
	JobQueue chan Job
}
