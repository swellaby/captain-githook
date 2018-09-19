// +build mage

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

const (
	testResultsDirectory     = ".testresults/"
	coverageResultsDirectory = ".coverage/"
	junitXmlTestResultsFile  = testResultsDirectory + "junit.xml"
	jsonTestResultsFile      = testResultsDirectory + "unit.json"
	coverageOutFile          = coverageResultsDirectory + "coverage.out"
	coberturaCoverageFile    = coverageResultsDirectory + "cobertura.xml"
	htmlCoverageFile		 = coverageResultsDirectory + "index.html"
	goCovJsonFile			 = coverageResultsDirectory + "gocov.json"
	goVetResultsFile         = testResultsDirectory + "govet.out"
	goLintResultsFile        = testResultsDirectory + "golint.out"
)

func goGetTool(tool string) {
	cmd := exec.Command("go", "get", tool)
	cmd.Dir = os.TempDir()
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
	fmt.Println("Installing gocov...")
	goGetTool("github.com/axw/gocov/gocov")
	fmt.Println("Installing gocovxml...")
	goGetTool("github.com/AlekSi/gocov-xml")
	fmt.Println("Installing gocov-html...")
	goGetTool("github.com/matm/gocov-html")
	fmt.Println("Installing golangci-lint...")
	goGetTool("github.com/golangci/golangci-lint/cmd/golangci-lint")
}

func createTestOutputDirectories() {
	os.MkdirAll(testResultsDirectory, os.ModePerm)
	os.MkdirAll(coverageResultsDirectory, os.ModePerm)
}

func cleanTestResultFiles() {
	os.RemoveAll(testResultsDirectory)
	os.RemoveAll(coverageResultsDirectory)
}

func Setup() {
	installDevTools()
	fmt.Println("Creating directories for test/coverage/etc. output files...")
	createTestOutputDirectories()
}

func getPackageNames(pkgPath string) string {
	cmd := exec.Command("go", "list", pkgPath)
	out, _ := cmd.CombinedOutput()
	result := strings.TrimSuffix(string(out), "\n")
	return result

	// packages := strings.Split(string(out), "\n")
	// for _, pkg := range packages {
	// 	if len(pkg) < 1 {
	// 		fmt.Println("nope, last one")
	// 	} else {
	// 		fmt.Printf("%s\n", pkg)
	// 	}
	// }
}

func createCoberturaCodeCoverageReport() {
	cmd := exec.Command("gocov-xml")
	inFile, _ := ioutil.ReadFile(goCovJsonFile)
	cmd.Stdin = bytes.NewBuffer(inFile)
	outFile, _ := os.Create(coberturaCoverageFile)
	defer outFile.Close()
	cmd.Stdout = outFile
	cmd.Run()
}

func createHtmlCodeCoverageReport() {
	cmd := exec.Command("gocov-html", goCovJsonFile)
	outFile, _ := os.Create(htmlCoverageFile)
	defer outFile.Close()
	cmd.Stdout = outFile
	cmd.Run()
}

func createGoCovJsonFile() {
	cmd := exec.Command("gocov", "convert", coverageOutFile)
	outFile, _ := os.Create(goCovJsonFile)
	defer outFile.Close()
	cmd.Stdout = outFile
	cmd.Run()
}

func createCoverageReports() {
	createGoCovJsonFile()
	createCoberturaCodeCoverageReport()
	createHtmlCodeCoverageReport()
}

// Test Runs the unit tests
func Test() error {
	cleanTestResultFiles()
	createTestOutputDirectories()
	fmt.Println("Running tests...")
	pkg := getPackageNames("./pkg/...")
	coverProfile := "-coverprofile=" + coverageOutFile
	testCmd := exec.Command("go", "test", pkg, "-v", coverProfile, "-covermode", "count")
	testOutput, err := testCmd.CombinedOutput()
	fmt.Printf(string(testOutput))

	// Create Standard JSON Result File
	jsonCmd := exec.Command("go", "tool", "test2json", "-t", "-p", pkg)
	jsonCmd.Stdin = bytes.NewBuffer(testOutput)
	jsonOutFile, err := os.Create(jsonTestResultsFile)
	jsonCmd.Stdout = jsonOutFile
	jsonCmd.Run()
	jsonOutFile.Close()

	// Create JUnit XML Result File
	junitCmd := exec.Command("go-junit-report")
	junitCmd.Stdin = bytes.NewBuffer(testOutput)
	outfile, err := os.Create(junitXmlTestResultsFile)

	junitCmd.Stdout = outfile
	junitCmd.Run()
	outfile.Close()

	createCoverageReports()

	return err
}

// Lint Runs the linter
func Lint() error {
	fmt.Println("Running golint...")
	cmd := exec.Command("golint", "./...")
	outfile, err := os.Create(goLintResultsFile)
	defer outfile.Close()
	cmd.Stdout = outfile
	output, err := cmd.CombinedOutput()
	fmt.Printf(string(output))
	return err
}

// Lint Runs Vet
func Vet() error {
	fmt.Println("Running go vet...")
	cmd := exec.Command("go", "vet", "./...")
	outfile, err := os.Create(goVetResultsFile)
	defer outfile.Close()
	cmd.Stdout = outfile
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
	Vet()
}

func Clean() {
	fmt.Println("Cleaning...")
	//os.RemoveAll("MyApp")
}
