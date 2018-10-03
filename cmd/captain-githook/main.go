package main

import (
	"fmt"
	"os/exec"
	"runtime"
	// "github.com/swellaby/captain-githook/pkg/captaingithook"
)

// func main() {
// 	fmt.Println("Ahoy there Matey!")
// 	out, err := captaingithook.Run("git rev-parse --show-toplevel", "c:/dev/vsts-bump/bin")
// 	if err != nil {
// 		fmt.Printf("Crashed and burned with error %s\n", err)
// 		fmt.Printf("Error details: %s\n", out)
// 	} else {
// 		fmt.Printf("The output was: %s", out)
// 		// fmt.Printf("The output was: '%s'\n", string(out[:len(out)-1]))
// 	}
// }

func main() {
	fmt.Println("Ahoy there Matey!")
	var runner, runnerArg string

	if runtime.GOOS == "windows" {
		runner = "cmd.exe"
		runnerArg = "/C"
	} else {
		runner = "sh"
		runnerArg = "-c"
	}

	// hookScript := "git rev-parse --show-toplevel"
	hookScript := "dir"
	cmd := exec.Command(runner, runnerArg, hookScript)
	cmd.Dir = "c:/dev"
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("Crashed and burned with error %s\n", err)
		fmt.Printf("Error details: %s\n", string(out[:len(out)-1]))
	} else {
		fmt.Printf("The output was: %s", string(out[:len(out)-1]))
		// fmt.Printf("The output was: '%s'\n", string(out[:len(out)-1]))
	}
}
