package captaingithook

import (
    "fmt"
    "os/exec"
	"runtime"
)

func getRunnerInfo () (runner, runnerArg string) {
    if runtime.GOOS == "windows" {
		runner = "cmd.exe"
		runnerArg = "/C"
	} else {
		runner = "sh"
		runnerArg = "-c"
    }

    return runner, runnerArg
}

func runCommand (command string, directory string) (resultOutput string, err error) {
    runner, runnerArg := getRunnerInfo()

    cmd := exec.Command(runner, runnerArg, command)

    if len(directory) > 0 {
        cmd.Dir = directory
    }

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