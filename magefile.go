// +build mage

package main

import (
	"fmt"
	"os/exec"
)

// var Default = Build

// Test is
func Test() error {
	fmt.Println("Running tests...")
	cmd := exec.Command("go", "test", "-v", "./...")
	output, err := cmd.CombinedOutput()
	fmt.Printf(string(output))
	return err
}

// Hello world
func Lint() error {
	fmt.Println("Running golint")
	// cmd := exec.Command("golint", "", "-v", "./...")
	return nil
}

// A build step that requires additional params, or platform specific steps for example
// func Build() error {
// 	fmt.Println("Building...")
// 	cmd := exec.Command("go", "build", "-o", "MyApp", ".")
// 	return cmd.Run()
// }

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	//os.RemoveAll("MyApp")
}

