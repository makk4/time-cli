package cmd

import (
	"fmt"
	io "timecli/ioutils"
	job "timecli/timeutils"

	"github.com/spf13/cobra"
)

// reportCmd represents the report command
var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Prints a report",
	Long:  `Prints a report on a given timespan`,
	Run: func(cmd *cobra.Command, args []string) {
		var jsonFile = io.ReadFile()
		for i := range jsonFile.Projects {
			for k := range jsonFile.Projects[i].Jobs {
				var j job.Job = jsonFile.Projects[i].Jobs[k]
				fmt.Println(job.GetString(j))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(reportCmd)
}
