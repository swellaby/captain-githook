package cli

import (
	"github.com/spf13/cobra"
	"github.com/swellaby/captain-githook/captaingithook"
	"github.com/swellaby/ci-detective/cidetective"
)

var initializeCaptainGithook = captaingithook.Initialize
var initializeCaptainGithookWithConfigName = captaingithook.InitializeWithFileName
var isCI = cidetective.IsCI
var initConfigFileName string

var initCmd = &cobra.Command{
	Use:  "init",
	RunE: initialize,
}

func init() {
	initCmd.Flags().StringVarP(&initConfigFileName, "config-filename", "f", "", "Your desired captain-githook config file name")
	rootCmd.AddCommand(initCmd)
}

func initialize(*cobra.Command, []string) error {
	log("Ahoy there matey!")
	if isCI() {
		log("CI environment detected. Skipping git hook install.")
		return nil
	}
	if len(initConfigFileName) < 1 {
		log("Initializing your repository with the default config file name.")
		return initializeCaptainGithook()
	}

	logf("Initializing your repository with the your requested config file name: '%s'.", initConfigFileName)
	return initializeCaptainGithookWithConfigName(initConfigFileName)
}
