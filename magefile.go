// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
)

func goGetTool(tool string) (string, error) {
	cmd := exec.Command("go", "get", "-u", tool)
	cmd.Dir = os.TempDir();
	out, err := cmd.CombinedOutput()
	return string(out), err
}

// var Default = Build

func Install() error {
	goGetTool("github.com/jstemmer/go-junit-report")
	return nil
	// _, err := installTool("go-junit-report", "github.com/jstemmer/go-junit-report")
	// return nil
}

// Test Runs the unit tests.
func Test() error {
	fmt.Println("Running tests...")
	cmd := exec.Command("go", "test", "-v", "./pkg/...")
	output, err := cmd.CombinedOutput()
	fmt.Printf(string(output))
	return err
}

// Lint Runs the linter.
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
