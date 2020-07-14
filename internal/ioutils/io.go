package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"

	timetable "github.com/makk4/time-cli/internal/timeutils"
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

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	var jsonFile timetable.TimeTable

	err = json.Unmarshal(byteValue, &jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	return jsonFile

}

// WriteFile writes TimeTable to file
func WriteFile(jsonFile timetable.TimeTable) {
	file, err := json.Marshal(&jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile(getUserPath(), file, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
