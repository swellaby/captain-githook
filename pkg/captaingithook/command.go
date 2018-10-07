package captaingithook

import (
	"os/exec"
	"runtime"
)

var osCommand = exec.Command
var createCommand = newCommand
var runCommand = run
var runCommandInDir = runInDir

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

func newCommand(directory, command string) command {
	runner, runnerArg := getRunnerInfo(runtime.GOOS)
	cmdArgs := []string{runnerArg, command}
	cmd := osCommand(runner, cmdArgs...)

	if len(directory) > 0 {
		cmd.Dir = directory
	}

	return cmd
}

func run(command string) (resultOutput string, err error) {
	return runInDir("", command)
}

func runInDir(directory, command string) (resultOutput string, err error) {
	cmd := createCommand(directory, command)

	out, err := cmd.CombinedOutput()
	resultOutput = string(out)

	return resultOutput, err
}
