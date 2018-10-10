package captaingithook

import (
	"bytes"
	"encoding/json"
	"errors"
	// "fmt"
	"path/filepath"
	"testing"
)

const defaultConfigFileName = "captaingithook.json"

var hooksConfig = &HooksConfig{
	PreCommit: "go test ./...",
}
var config = &Config{Hooks: *hooksConfig}

func getDefaultJSON() []byte {
	data, _ := json.MarshalIndent(config, "", "  ")
	return data
}

func TestIsValidConfigFileNameReturnsFalseOnInvalidName(t *testing.T) {
	fileName := "captaingithook.yml"
	isValid := isValidConfigFileName(fileName)

	if isValid {
		t.Errorf("Validity was wrong for file name: %s. Expected: true, but got: %t.", fileName, isValid)
	}
}

func TestIsValidConfigFileNameReturnsFalseOnEmptyName(t *testing.T) {
	fileName := ""
	isValid := isValidConfigFileName(fileName)

	if isValid {
		t.Errorf("Validity was wrong for file name: %s. Expected: true, but got: %t.", fileName, isValid)
	}
}

func TestIsValidConfigFileNameReturnsTrueOnValidNames(t *testing.T) {
	validConfigFileNames := []string{
		defaultConfigFileName,
		".captaingithook.json",
		"captaingithookrc",
		".captaingithookrc",
		"captaingithookrc.json",
		".captaingithookrc.json",
		"captain-githook.json",
		".captain-githook.json",
		"captain-githookrc",
		".captain-githookrc",
		"captain-githookrc.json",
		".captain-githookrc.json",
	}

	for _, fileName := range validConfigFileNames {
		isValid := isValidConfigFileName(fileName)

		if !isValid {
			t.Errorf("Validity was wrong for file name: %s. Expected: true, but got: %t.", fileName, isValid)
		}
	}
}

func TestConfigFileExistsReturnsFalseWhenNoFilesFound(t *testing.T) {
	origFileExists := fileExists
	fileExists = func(string) bool { return false }
	defer func() { fileExists = origFileExists }()
	origGetRepoRoot := getGitRepoRootDirectoryPath
	getGitRepoRootDirectoryPath = func() (string, error) {
		return "", nil
	}
	defer func() { getGitRepoRootDirectoryPath = origGetRepoRoot }()
	foundFile := configFileExists("")

	if foundFile {
		t.Errorf("Got wrong result for found config file. Expected %t, but got %t", false, foundFile)
	}
}

func TestConfigFileExistsReturnsTrueWhenFileFound(t *testing.T) {
	origFileExists := fileExists
	fileExists = func(string) bool { return true }
	defer func() { fileExists = origFileExists }()
	origGetRepoRoot := getGitRepoRootDirectoryPath
	getGitRepoRootDirectoryPath = func() (string, error) {
		return "", nil
	}
	defer func() { getGitRepoRootDirectoryPath = origGetRepoRoot }()
	foundFile := configFileExists("")

	if !foundFile {
		t.Errorf("Got wrong result for found config file. Expected %t, but got %t", true, foundFile)
	}
}

func TestInitConfigFileUsesCorrectDefault(t *testing.T) {
	originalWriteFile := writeFile
	defer func() { writeFile = originalWriteFile }()
	var actualFileName string
	var actualData []byte
	expectedData := getDefaultJSON()
	writeFile = func(fileName string, data []byte) error {
		actualFileName = fileName
		actualData = data
		return nil
	}

	initConfigFile("", "")

	if actualFileName != defaultConfigFileName {
		t.Errorf("Attempted to create wrong config file name. Expected: %s, but got: %s.", defaultConfigFileName, actualFileName)
	}

	if !bytes.Equal(actualData, expectedData) {
		t.Errorf("Attempted to create wrong config file contents. Expected: %s, but got: %s.", expectedData, actualData)
	}
}

func TestInitConfigFileUsesSpecifiedFileName(t *testing.T) {
	originalWriteFile := writeFile
	defer func() { writeFile = originalWriteFile }()
	var actualFileName string
	var actualData []byte
	expectedData := getDefaultJSON()
	writeFile = func(fileName string, data []byte) error {
		actualFileName = fileName
		actualData = data
		return nil
	}

	desiredFileName := ".captain-githookrc"

	initConfigFile("", desiredFileName)

	if actualFileName != desiredFileName {
		t.Errorf("Attempted to create wrong config file name. Expected: %s, but got: %s.", defaultConfigFileName, actualFileName)
	}

	if !bytes.Equal(actualData, expectedData) {
		t.Errorf("Attempted to create wrong config file contents. Expected: %s, but got: %s.", expectedData, actualData)
	}
}

func TestInitConfigFileReturnsCorrectErrorOnWriteError(t *testing.T) {
	originalWriteFile := writeFile
	defer func() { writeFile = originalWriteFile }()
	errMsgDetails := "ouch"
	expectedErrMsg := "unexpected error encountered while trying to create the config file. Error details: " + errMsgDetails
	writeFile = func(fileName string, data []byte) error {
		return errors.New(errMsgDetails)
	}

	err := initConfigFile("", ".captaingithookrc.json")
	if err.Error() != expectedErrMsg {
		t.Errorf("Got wrong error message. Expected: %s, but got: %s.", expectedErrMsg, err.Error())
	}
}

func TestGetDefaultConfigJSONContentReturnsCorrectValues(t *testing.T) {
	expectedJSONContent := getDefaultJSON()
	actualJSONContent, err := getDefaultConfigJSONContent()

	if err != nil {
		t.Errorf("Error should have been nil but was not. Error was: %s", err)
	}

	if !bytes.Equal(expectedJSONContent, actualJSONContent) {
		t.Errorf("Did not get correct JSON data for default config object. Expected: %s, but got: %s", string(expectedJSONContent), string(actualJSONContent))
	}
}

func TestGetRepoConfigReturnsCorrectErrorWhenConfigFileDoesNotExist(t *testing.T) {
	origFunc := readFile
	defer func() { readFile = origFunc }()
	readFile = func(string) ([]byte, error) {
		return nil, errors.New("")
	}
	config, err := getRepoConfig("")

	if config != nil {
		t.Errorf("Config was not nil. Config was: %v", config)
	}

	if err != errConfigFileNotFound {
		t.Errorf("Did not get correct error. Expected: %s, but got: %s", errConfigFileNotFound, err)
	}
}

func TestGetRepoConfigScansForConfigFilesInSpecifiedDirectory(t *testing.T) {
	repoPath := "/usr/me/foo"
	expConfigFilePath := filepath.Join(repoPath, defaultConfigFileName)
	actConfigFilePath := ""
	origFunc := readFile
	counter := 0
	defer func() { readFile = origFunc }()
	readFile = func(filePath string) ([]byte, error) {
		if counter == 0 {
			actConfigFilePath = filePath
		}
		counter++
		return nil, errors.New("")
	}
	getRepoConfig(repoPath)

	if actConfigFilePath != expConfigFilePath {
		t.Errorf("Did not get correct config file path. Expected: %s, but got: %s", expConfigFilePath, actConfigFilePath)
	}
}

func TestGetRepoConfigReturnsCorrectErrorWhenConfigFileParsingFails(t *testing.T) {
	configFileContents := getDefaultJSON()
	var actFileContents []byte
	origReadFunc := readFile
	defer func() { readFile = origReadFunc }()
	readFile = func(filePath string) ([]byte, error) {
		return configFileContents, nil
	}
	origUnmarshallFunc := jsonUnmarshall
	defer func() { jsonUnmarshall = origUnmarshallFunc }()
	jsonUnmarshall = func(contents []byte, v interface{}) error {
		actFileContents = contents
		return errors.New("")
	}

	config, err := getRepoConfig("")

	if config != nil {
		t.Errorf("Config was not nil. Config was: %v", config)
	}

	if err != errConfigFileParseFailed {
		t.Errorf("Did not get correct error. Expected: %s, but got: %s", errConfigFileNotFound, err)
	}

	if !bytes.Equal(actFileContents, configFileContents) {
		t.Errorf("Did not pass correct bytes to unmarshall function. Expected: %v, but got: %v", configFileContents, actFileContents)
	}
}

func TestGetRepoConfigReturnsCorrectConfig(t *testing.T) {
	configFileContents := getDefaultJSON()
	origReadFunc := readFile

	defer func() { readFile = origReadFunc }()
	readFile = func(filePath string) ([]byte, error) {
		return configFileContents, nil
	}

	actConfig, err := getRepoConfig("")

	if *actConfig != *config {
		t.Errorf("Did not get correct config. Expected: %v, but got: %v", config, actConfig)
	}

	if err != nil {
		t.Errorf("Error was not nil. Error was: %s", err)
	}
}

// func TestFoo(t *testing.T) {
// 	config, err := getRepoConfig("c:/dev/captain-githook")
// 	if err != nil {
// 		t.Errorf("Error was not nil. Error: %s", err)
// 	}
// 	hooks := config.Hooks
// 	fmt.Printf("Config value: %v\n", hooks)
// 	fmt.Printf("PreCommit hook value: %s\n", hooks.PreCommit)

// }
