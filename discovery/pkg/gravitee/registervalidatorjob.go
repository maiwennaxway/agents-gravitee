package gravitee

import (
	"github.com/Axway/agent-sdk/pkg/jobs"
)

type registerAPIValidatorJob struct {
	jobs.Job
	validatorReady    jobFirstRunDone
	registerValidator func()
}

func newRegisterAPIValidatorJob(apiReady jobFirstRunDone, registerValidator func()) *registerAPIValidatorJob {
	job := &registerAPIValidatorJob{
		validatorReady:    apiReady,
		registerValidator: registerValidator,
	}
	return job
}

func (j *registerAPIValidatorJob) Ready() bool {
	return j.validatorReady()
}

func (j *registerAPIValidatorJob) Status() error {
	return nil
}

func (j *registerAPIValidatorJob) Execute() error {
	j.registerValidator()
	return nil
}
