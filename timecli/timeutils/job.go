package timeutils

import (
	"time"

	"github.com/segmentio/ksuid"
)

// Job a single job unit structure
type Job struct {
	StartTime   time.Time   `json:"starttime"`
	EndTime     time.Time   `json:"endtime"`
	ID          ksuid.KSUID `json:"id"`
	ProjectName string      `json:"name"`
	Tags        []string    `json:"tags"`
	Description string      `json:"description"`
}

// GetString returns pretty string representation
func GetString(j Job) string {
	return j.ID.String() + "   Project: " + j.ProjectName + "   start: " + j.StartTime.String() + "   end: " + j.EndTime.String()
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
