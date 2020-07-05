package timeutils

import (
	"reflect"
	"testing"
	"time"
)

func TestGetProjectName(t *testing.T) {
	type args struct {
		j *Job
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetProjectName(tt.args.j); got != tt.want {
				t.Errorf("GetProjectName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTags(t *testing.T) {
	type args struct {
		j *Job
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTags(tt.args.j); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStartTime(t *testing.T) {
	type args struct {
		j *Job
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStartTime(tt.args.j); got != tt.want {
				t.Errorf("GetStartTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
