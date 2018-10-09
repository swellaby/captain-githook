package cli

import (
	"github.com/spf13/cobra"
	"github.com/swellaby/captain-githook/captaingithook"
)

// Runner describes a CLI runner
type Runner interface {
	Execute() error
}

var rootCmd = &cobra.Command{
	Use:     "captain-githook",
	Version: captaingithook.Version,
}

// GetRunner returns the runner
func GetRunner() Runner {
	return rootCmd
}
