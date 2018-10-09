package main

import (
	// "fmt"
	"github.com/spf13/cobra"
	"github.com/swellaby/captain-githook/captaingithook"
)

var rootCmd *cobra.Command
var initCmd *cobra.Command
var runCmd *cobra.Command
var removeCmd *cobra.Command

var initializeCaptainGithook = captaingithook.Initialize
var initializeCaptainGithookWithConfigName = captaingithook.InitializeWithFileName

func runRootCmd(cmd *cobra.Command, args []string) {

}

func setup(cmd *cobra.Command, args []string) {

}

func runHook(cmd *cobra.Command, args []string) {

}

func remove(cmd *cobra.Command, args []string) {

}

func init() {
	rootCmd = &cobra.Command{
		Version: captaingithook.Version,
		Run:     runRootCmd,
	}
	initCmd = &cobra.Command{
		Run: setup,
	}
	runCmd = &cobra.Command{
		Run: runHook,
	}
	removeCmd = &cobra.Command{
		Run: remove,
	}
}
