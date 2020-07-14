// data structure for shifts
package timeutils

import (
	"time"
)

//Project data structure
type Project struct {
	Name         string    `json:"name"`
	Company      string    `json:"company"`
	Tags         []string  `json:"tags"`
	Description  string    `json:"description"`
	SheduledTime time.Time `json:"sheduledTime"`
	TotalTime    time.Time `json:"totalTime"`
	Jobs         []Job     `json:"job"`
}

// TimeStats represtents the overall hours stats
type TimeStats struct {
	hoursSheduled float32
	hoursDone     float32
	sum           float32
}

type day struct {
	Name string `json:"name"`
	Jobs []Job  `json:"job"`
}

// TimeTable the root file structure
type TimeTable struct {
	Name         string    `json:"name"`
	MonthlyHours int       `json:"monthlyhours"`
	WeeklyHours  int       `json:"weeklyhours"`
	DailyHours   int       `json:"dailyhours"`
	Workdays     bool      `json:"workdays"`
	RunningJob   string    `json:"runningjob"`
	Days         []day     `json:"days"`
	Projects     []Project `json:"projects"`
}

func (tt TimeTable) init(name string, monthlyHours int, weeklyHours int, dailyHours int, workdays bool) {
	tt.Name = name
	tt.MonthlyHours = monthlyHours
	tt.WeeklyHours = weeklyHours
	tt.DailyHours = dailyHours
	tt.Workdays = workdays
}

// AddProject adds new project if it not exists, by unique name
func (tt TimeTable) AddProject(name string) {
	var isNewProject bool = true
	for i := range tt.Projects {
		if tt.Projects[i].Name == name {
			isNewProject = false
		}
	}

	if isNewProject == true {
		p := Project{Name: name}
		tt.Projects = append(tt.Projects, p)
	}
}

// getMonthlyHours returns the hours
func (tt TimeTable) getMonthlyHours() TimeStats {
	var ts TimeStats

	// for i := range tt.Days {
	// 	if true {
	// 		for s := range tt.Days[i].Jobs {
	// 			var j Job = tt.Days[i].Jobs[s]
	// 			j = j.EndTime
	// 		}
	// 	}
	// }

	return ts
}
