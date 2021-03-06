package cmd

import (
	"fmt"
	"os"
	"time"

	io "github.com/makk4/time-cli/internal/ioutils"
	job "github.com/makk4/time-cli/internal/timeutils"

	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops a running job",
	Long:  `Stops a running job`,
	Run: func(cmd *cobra.Command, args []string) {

		var jsonFile = io.ReadFile()
		var j job.Job

		if jsonFile.RunningJob == "" {
			fmt.Println("No Job startet")
			os.Exit(0)
		}

		jsonFile.RunningJob = ""

		for i := range jsonFile.Projects {
			for k := range jsonFile.Projects[i].Jobs {
				if jsonFile.Projects[i].Jobs[k].ID == jsonFile.RunningJob {
					j = jsonFile.Projects[i].Jobs[k]
					jsonFile.Projects[i].Jobs[k].EndTime = time.Now().String()
				}
			}
		}

		io.WriteFile(jsonFile)

		fmt.Println(j.ProjectName, j.Tags, "stoped @ ", j.StartTime)

	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
