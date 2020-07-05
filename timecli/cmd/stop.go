package cmd

import (
	io "timecli/ioutils"

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
		// var j job.Job

		io.ReadJob()
		// io.WriteJob(&j)

		// fmt.Println(job.GetName(&j), job.GetTags(&j), "stoped @ ", job.GetStartTime(&j))

	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
