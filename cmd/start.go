package cmd

import (
	"fmt"
	"os"
	io "timecli/internals"
	job "timecli/timeutils"
	project "timecli/timeutils"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts a new project",
	Long: `Starts a new project long 
multi line descrp`,
	Run: func(cmd *cobra.Command, args []string) {

		var jsonFile = io.ReadFile()

		if jsonFile.RunningJob != "" {
			fmt.Println("Job already startet")
			os.Exit(0)
		}

		var j job.Job

		jsonFile.RunningJob = j.StartJob(args)

		var newProject bool = true
		for i := range jsonFile.Projects {
			if jsonFile.Projects[i].Name == j.ProjectName {
				jsonFile.Projects[i].Jobs = append(jsonFile.Projects[i].Jobs, j)
				newProject = false
			}
		}

		if newProject == true {
			var p project.Project
			p.Name = j.ProjectName
			p.Jobs = append(p.Jobs, j)
			jsonFile.Projects = append(jsonFile.Projects, p)
		}
		jsonFile.AddProject()
		io.WriteFile(jsonFile)

		fmt.Println(j.ProjectName, j.Tags, "startet @ ", j.StartTime)

	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
