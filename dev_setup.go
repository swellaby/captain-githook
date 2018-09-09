// +build mage,vscode

package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func toolInstalled(tool string) bool {
	var runner, runnerArg string
	if runtime.GOOS == "windows" {
		runner = "cmd.exe"
		runnerArg = "/C"
	} else {
		runner = "sh"
		runnerArg = "-c"
	}

	cmd := exec.Command(runner, runnerArg, tool)
	if err := cmd.Run(); err != nil {
		fmt.Printf("Tool: '%s' not found.\n", tool)
		return false
	}

	return true
}

func installTool(tool, toolPath string) (string, error) {
	if isInstalled := toolInstalled(tool); isInstalled == true {
		fmt.Printf("Tool: '%s' is already installed\n", tool)
		return "", nil
	}

	cmd := exec.Command("go", "get", "-u", toolPath)
	cmd.Dir = os.TempDir();
	fmt.Printf("Installing %s...\n", tool)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func main() {
	out, err := installTool("mage", "github.com/magefile/mage")

	if err != nil {
		fmt.Printf("An error occurred while installing mage %s\n", err)
		fmt.Printf("Error details: %s\n", string(out))
	} else {
		fmt.Println("Mage is setup successfully. Please run `mage` to view the available targets.")
	}
}