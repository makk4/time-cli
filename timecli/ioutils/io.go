package ioutils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	job "timecli/timeutils"
	project "timecli/timeutils"
)

// JSONFile test
type JSONFile struct {
	Name         string            `json:"name"`
	MonthlyHours int               `json:"monthlyhours"`
	WeeklyHours  int               `json:"weeklyhours"`
	DailyHours   int               `json:"dailyhours"`
	Workdays     bool              `json:"workdays"`
	RunningJob   string            `json:"runningjob"`
	Jobs         []job.Job         `json:"jobs"`
	Projects     []project.Project `json:"projects"`
}

func getUserPath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return filepath.Join(usr.HomeDir, "timeapp.json")
}

func readFile() JSONFile {
	file, err := os.Open(getUserPath())
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var jsonFile JSONFile

	json.Unmarshal(byteValue, &jsonFile)

	fmt.Println(jsonFile)

	return jsonFile

}

func writeFile(jsonFile JSONFile) {
	file, err := json.MarshalIndent(jsonFile, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	_ = ioutil.WriteFile(getUserPath(), file, 0644)
}

// WriteJob test
func WriteJob(j job.Job) {
	var jsonFile = readFile()
	jsonFile.Jobs = append(jsonFile.Jobs, j)
	writeFile(jsonFile)
}

// ReadJob test
func ReadJob() {
	readFile()
}
