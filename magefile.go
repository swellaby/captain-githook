// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
)

func goGetTool(tool string) {
	cmd := exec.Command("go", "get", tool)
	cmd.Dir = os.TempDir();
	out, err := cmd.CombinedOutput()
	fmt.Printf("%s", string(out))

	if err != nil {
		os.Exit(1)
	}
}

func installDevTools() {
	fmt.Println("Installing dev tools...")
	fmt.Println("Installing go-junit-report...")
	goGetTool("github.com/jstemmer/go-junit-report")
}

func createTestOutputDirectories() {
	fmt.Println("Creating directories for test/coverage/etc. output files...")
}

func Setup() {
	installDevTools()
	createTestOutputDirectories()
}

// Test Runs the unit tests
func Test() error {
	fmt.Println("Running tests...")
	cmd := exec.Command("go", "test", "-v", "./pkg/...")
	output, err := cmd.CombinedOutput()
	fmt.Printf(string(output))
	return err
}

// Lint Runs the linter
func Lint() error {
	fmt.Println("Running golint...")
	cmd := exec.Command("golint", "./...")
	output, err := cmd.CombinedOutput()
	fmt.Printf(string(output))
	return err
}

// A build step that requires additional params, or platform specific steps for example
// func Build() error {
// 	fmt.Println("Building...")
// 	cmd := exec.Command("go", "build", "-o", "MyApp", ".")
// 	return cmd.Run()
// }

func Ci() {
	Test()
	fmt.Println()
	Lint()
}

func Clean() {
	fmt.Println("Cleaning...")
	//os.RemoveAll("MyApp")
}
