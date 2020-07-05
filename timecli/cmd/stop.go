package cmd

import (
	"fmt"
	"os"
	"time"
	io "timecli/ioutils"
	job "timecli/timeutils"

	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
				if jsonFile.Projects[i].Jobs[k].ID.String() == jsonFile.RunningJob {
					j = jsonFile.Projects[i].Jobs[k]
					jsonFile.Projects[i].Jobs[k].EndTime = time.Now()
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
