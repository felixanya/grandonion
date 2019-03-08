package jobque

type Job struct {
	Payload 	Payload
}

var JobQueue chan Job

func (j *Job) doJob() string {
	return ">>>do jobs>>>"
}
