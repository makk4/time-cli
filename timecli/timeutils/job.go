package timeutils

import (
	"time"

	"github.com/segmentio/ksuid"
)

//Job test
type Job struct {
	StartTime   time.Time   `json:"starttime"`
	EndTime     time.Time   `json:"endtime"`
	ID          ksuid.KSUID `json:"id"`
	ProjectName string      `json:"name"`
	Tags        []string    `json:"tags"`
	Description string      `json:"description"`
}

// StartJob test
func StartJob(j *Job, startTime time.Time, projectName string, tags []string) {
	j.ID = ksuid.New()
	j.StartTime = startTime
	j.ProjectName = projectName
	j.Tags = tags
}

// EndJob test
func EndJob(j *Job, endTime time.Time, projectName string, tags []string) {
	j.ID = ksuid.New()
	j.EndTime = endTime
	j.ProjectName = projectName
	j.Tags = tags
}
