package timeutils

// TimeTable test
type TimeTable struct {
	Name         string    `json:"name"`
	MonthlyHours int       `json:"monthlyhours"`
	WeeklyHours  int       `json:"weeklyhours"`
	DailyHours   int       `json:"dailyhours"`
	Workdays     bool      `json:"workdays"`
	RunningJob   string    `json:"runningjob"`
	Projects     []Project `json:"projects"`
}
