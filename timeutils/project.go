package timeutils

import (
	"time"
)

//Project a 2Project structure
type Project2 struct {
	Name         string    `json:"name"`
	Company      string    `json:"company"`
	Tags         []string  `json:"tags"`
	Description  string    `json:"description"`
	SheduledTime time.Time `json:"sheduledTime"`
	TotalTime    time.Time `json:"totalTime"`
	Jobs         []Job     `json:"job"`
}
