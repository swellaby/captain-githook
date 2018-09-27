package captaingithook

import (
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
	resultOutput = string(out)

	return resultOutput, err
}
