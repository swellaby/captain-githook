// +build mage

package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

const (
	testResultsDirectory = ".testresults/"
	coverageResultsDirectory = ".coverage/"
	junitXmlTestResultsFileName = "junit.xml"
	junitXmlTestResultsFile = testResultsDirectory + junitXmlTestResultsFileName
	goTestResultsJsonFileName = "unit.json"
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
	os.MkdirAll(testResultsDirectory, os.ModePerm)
	os.MkdirAll(coverageResultsDirectory, os.ModePerm)
}

func Td() {
	createTestOutputDirectories()
	testArgs := fmt.Sprintf("-v | go-junit-report > %s%s", testResultsDirectory, junitXmlTestResultsFile)
	fmt.Printf("%s", testArgs)
}

func Setup() {
	installDevTools()
	fmt.Println("Creating directories for test/coverage/etc. output files...")
	createTestOutputDirectories()
}

// Test Runs the unit tests
func Test() error {
	createTestOutputDirectories()
	fmt.Println("Running tests...")
	testCmd := exec.Command("go", "test", "./pkg/...", "-v")
	testOutput, err := testCmd.CombinedOutput()
	fmt.Printf(string(testOutput))

	// Create JUnitXML Formatted Results
	junitCmd := exec.Command("go-junit-report")
	junitCmd.Stdin = bytes.NewBuffer(testOutput)
	outfile, err := os.Create(junitXmlTestResultsFile)
	defer outfile.Close()
	junitCmd.Stdout = outfile
	junitCmd.Run()

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
