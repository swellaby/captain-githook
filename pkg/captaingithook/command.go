package captaingithook

import (
	"fmt"
	"os/exec"
	"runtime"
)

var osCommand = exec.Command
var createCommand = newCommand
var runCommand = run

type command interface {
	CombinedOutput() ([]byte, error)
}

func getRunnerInfo(operatingSystem string) (runner, runnerArg string) {
	if operatingSystem == "windows" {
		runner = "cmd.exe"
		runnerArg = "/C"
	} else {
		runner = "sh"
		runnerArg = "-c"
	}

	return runner, runnerArg
}

func newCommand(directory, name string, args ...string) command {
	cmd := osCommand(name, args...)

	if len(directory) > 0 {
		cmd.Dir = directory
	}

	return cmd
}

func run(command string, directory string) (resultOutput string, err error) {
	runner, runnerArg := getRunnerInfo(runtime.GOOS)
	cmd := createCommand(directory, runner, runnerArg, command)

	out, err := cmd.CombinedOutput()
	resultOutput = string(out[:len(out)-1])

	if err != nil {
		fmt.Printf("Crashed and burned with error %s\n", err)
		fmt.Printf("Error details: %s\n", resultOutput)
	} else {
		fmt.Printf("The output was: %s", resultOutput)
		// fmt.Printf("The output was: '%s'\n", string(out[:len(out)-1]))
	}

	return resultOutput, err
}
