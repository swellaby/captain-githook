package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/swellaby/captain-githook/captaingithook"
)

var initializeCaptainGithook = captaingithook.Initialize
var initializeCaptainGithookWithConfigName = captaingithook.InitializeWithFileName
var initConfigFileName string
var log = fmt.Println
var logf = fmt.Printf

var initCmd = &cobra.Command{
	Use:  "init",
	RunE: initialize,
}

func init() {
	initCmd.Flags().StringVarP(&initConfigFileName, "config-filename", "f", "", "Your desired captain-githook config file name")
	rootCmd.AddCommand(initCmd)
}

func initialize(cmd *cobra.Command, args []string) error {
	log("Ahoy there matey!")
	if len(initConfigFileName) < 1 {
		log("Initializing your repository with the default config file name.")
		return initializeCaptainGithook()
	}

	logf("Initializing your repository with the your requested config file name: '%s'.", initConfigFileName)
	return initializeCaptainGithookWithConfigName(initConfigFileName)
}
