package timeutils

import (
	"testing"
	"time"
)

func TestStartJob(t *testing.T) {
	type args struct {
		j           *Job
		startTime   time.Time
		projectName string
		tags        []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StartJob(tt.args.j, tt.args.startTime, tt.args.projectName, tt.args.tags)
		})
	}
}

func TestEndJob(t *testing.T) {
	type args struct {
		j           *Job
		endTime     time.Time
		projectName string
		tags        []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EndJob(tt.args.j, tt.args.endTime, tt.args.projectName, tt.args.tags)
		})
	}
}
