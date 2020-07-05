package ioutils

import (
	"reflect"
	"testing"
	job "timecli/timeutils"
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

func Test_readFile(t *testing.T) {
	tests := []struct {
		name string
		want JSONFile
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readFile(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_writeFile(t *testing.T) {
	type args struct {
		jsonFile JSONFile
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writeFile(tt.args.jsonFile)
		})
	}
}

func TestWriteJob(t *testing.T) {
	type args struct {
		j job.Job
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteJob(tt.args.j)
		})
	}
}

func TestReadJob(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReadJob()
		})
	}
}
