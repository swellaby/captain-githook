package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
)

func installMage() {
	cmd := exec.Command("go", "get", "github.com/magefile/mage")
	cmd.Dir = os.TempDir()
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("An error occurred while installing mage %s\n", err)
		fmt.Printf("Error details: %s\n", string(out))
		os.Exit(1)
	} else {
		fmt.Println("Mage successfully installed.")
	}
}

func runMageSetupTarget() {
	cmd := exec.Command("mage", "setup")
	_, currentFilePath, _, _ := runtime.Caller(0)
	cmd.Dir = filepath.Join(path.Dir(currentFilePath), "..")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error encountered while running `mage setup`. Error details: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s", string(out))
	os.Exit(0)
}

func main() {
	fmt.Println("Ensuring Mage is installed and available...")
	installMage()
	fmt.Println("Running `mage setup` to configure workspace...")
	runMageSetupTarget()
}
