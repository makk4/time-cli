package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	timetable "timecli/timeutils"
)

func getUserPath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return filepath.Join(usr.HomeDir, "timeapp.json")
}

// ReadFile reads the file and converts to TimeTable
func ReadFile() timetable.TimeTable {
	file, err := os.Open(getUserPath())
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var jsonFile timetable.TimeTable

	json.Unmarshal(byteValue, &jsonFile)

	return jsonFile

}

// WriteFile writes TimeTable to file
func WriteFile(jsonFile timetable.TimeTable) {
	file, err := json.Marshal(&jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	_ = ioutil.WriteFile(getUserPath(), file, 0644)
}
