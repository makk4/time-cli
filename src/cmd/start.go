package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
	"path/filepath"
	"strings"
	"time"

	"github.com/segmentio/ksuid"
	"github.com/spf13/cobra"
)

//Project test
type Project struct {
	startTime   time.Time
	endTime     time.Time
	id          ksuid.KSUID
	name        string
	tags        []string
	description string
}

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts a new project",
	Long: `Starts a new project long 
multi line descrp`,
	Run: func(cmd *cobra.Command, args []string) {

		var project Project

		project.startTime = time.Now()
		project.name = args[0]
		project.id = ksuid.New()

		for i, s := range args {
			if i > 0 && strings.HasPrefix(s, "+") {
				project.tags = append(project.tags, string([]rune(s)[1:]))
			}
		}

		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}

		file, _ := json.MarshalIndent(project, "", " ")

		_ = ioutil.WriteFile(filepath.Join(usr.HomeDir, "test.json"), file, 0644)

		fmt.Println(project.name, project.tags, "startet @ ", project.startTime.Format("2006-01-02 15:04:05"))

	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
