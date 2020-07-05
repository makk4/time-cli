package cmd

import (
	"fmt"
	"strings"
	"time"
	io "timecli/ioutils"
	job "timecli/timeutils"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts a new project",
	Long: `Starts a new project long 
multi line descrp`,
	Run: func(cmd *cobra.Command, args []string) {

		var j job.Job
		// var p project.Project
		tags := []string{}

		for i, s := range args {
			if i > 0 && strings.HasPrefix(s, "+") {
				tags = append(tags, string([]rune(s)[1:]))
			}
		}

		job.StartJob(&j, time.Now(), args[0], tags)

		io.WriteJob(j)

		fmt.Println(job.GetProjectName(&j), job.GetTags(&j), "startet @ ", job.GetStartTime(&j))

	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
