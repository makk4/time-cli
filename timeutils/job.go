package timeutils

import (
	"strings"
	"time"

	"github.com/segmentio/ksuid"
)

// Job a single job unit structure
type Job struct {
	StartTime   string   `json:"starttime"`
	EndTime     string   `json:"endtime"`
	ID          string   `json:"id"`
	ProjectName string   `json:"name"`
	Tags        []string `json:"tags"`
	Description string   `json:"description"`
}

// GetString returns pretty string representation
func GetString(j Job) string {
	return j.ID + "   Project: " + j.ProjectName + "   start: " + j.StartTime + "   end: " + j.EndTime
}

// StartJob test
func (j Job) StartJob(args []string) Job {
	tags := []string{}

	for i, s := range args {
		if i > 0 && strings.HasPrefix(s, "+") {
			tags = append(tags, string([]rune(s)[1:]))
		}
	}

	j.ID = ksuid.New().String()
	j.StartTime = time.Now().Format("Mon Jan _2 15:04:05 MST 2006")
	j.ProjectName = args[0]
	j.Tags = tags
	return j
}

// EndJob test
func (j Job) EndJob() {
	j.EndTime = time.Now().Format("Mon Jan _2 15:04:05 MST 2006")
}
