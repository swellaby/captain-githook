package captaingithook

import (
	"fmt"
	"os/exec"
	// "runtime"
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

func newCommand(directory, cmdName string, args ...string) command {
	// runner, runnerArg := getRunnerInfo(runtime.GOOS)
	// cmdArgs := append([]string{runnerArg, cmdName}, args...)
	cmd := osCommand(cmdName, args...)

	if len(directory) > 0 {
		cmd.Dir = directory
	}

	return cmd
}

func run(command string, commandArgs ...string) (resultOutput string, err error) {
	return runInDir("", command, commandArgs...)
}

func runInDir(directory, command string, commandArgs ...string) (resultOutput string, err error) {
	cmd := createCommand(directory, command, commandArgs...)

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Got an error when running command: '%s' with args: '%v'. Error: %s", command, commandArgs, err)
	}
	resultOutput = string(out)

	return resultOutput, err
}
