package internal

import (
	"reflect"
	"testing"
	timetable "timecli/timeutils"
)

func Test_getUserPath(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getUserPath(); got != tt.want {
				t.Errorf("getUserPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadFile(t *testing.T) {
	tests := []struct {
		name string
		want timetable.TimeTable
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadFile(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteFile(t *testing.T) {
	type args struct {
		jsonFile timetable.TimeTable
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteFile(tt.args.jsonFile)
		})
	}
}
