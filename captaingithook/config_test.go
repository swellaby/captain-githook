package captaingithook

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
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
	assert.False(t, isValidConfigFileName(fileName), "File name: %s should have been invalid", fileName)
}

func TestIsValidConfigFileNameReturnsFalseOnEmptyName(t *testing.T) {
	assert.False(t, isValidConfigFileName(""), "Empty string for file name should have been invalid")
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
		assert.True(t, isValidConfigFileName(fileName), "File name: %s should have been valid", fileName)
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
	assert.False(t, configFileExists(""))
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
	assert.True(t, configFileExists("captaingithook.json"))
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
	assert.Equal(t, defaultConfigFileName, actualFileName)
	assertionErrMsg := "Attempted to create wrong config file contents. Expected: %s, but got: %s."
	assert.Equal(t, expectedData, actualData, assertionErrMsg, expectedData, actualData)
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
	assert.Equal(t, desiredFileName, actualFileName)
	assertionErrMsg := "Attempted to create wrong config file contents. Expected: %s, but got: %s."
	assert.Equal(t, expectedData, actualData, assertionErrMsg, expectedData, actualData)
}

func TestInitConfigFileReturnsCorrectErrorOnWriteError(t *testing.T) {
	originalWriteFile := writeFile
	defer func() { writeFile = originalWriteFile }()
	errMsgDetails := "ouch"
	expectedErrMsg := "unexpected error encountered while trying to create the config file. Error details: " + errMsgDetails
	writeFile = func(fileName string, data []byte) error {
		return errors.New(errMsgDetails)
	}
	assert.Error(t, initConfigFile("", ".captaingithookrc.json"), expectedErrMsg)
}

func TestGetDefaultConfigJSONContentReturnsCorrectValues(t *testing.T) {
	expectedJSONContent := getDefaultJSON()
	actualJSONContent, err := getDefaultConfigJSONContent()
	assert.Nil(t, err)
	assert.Equal(t, expectedJSONContent, actualJSONContent)
}

func TestGetRepoConfigReturnsCorrectErrorWhenConfigFileDoesNotExist(t *testing.T) {
	origFunc := readFile
	defer func() { readFile = origFunc }()
	readFile = func(string) ([]byte, error) {
		return nil, errors.New("")
	}
	config, err := getRepoConfig("")
	assert.Nil(t, config)
	assert.Equal(t, errConfigFileNotFound, err)
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
	assertionErrMsg := "Did not get correct config file path. Expected: %s, but got: %s"
	assert.Equal(t, expConfigFilePath, actConfigFilePath, assertionErrMsg, expConfigFilePath, actConfigFilePath)
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
	assert.Nil(t, config)
	assert.Equal(t, errConfigFileParseFailed, err)
	assertionErrMsg := "Did not pass correct bytes to unmarshall function. Expected: %v, but got: %v"
	assert.Equal(t, configFileContents, actFileContents, assertionErrMsg, configFileContents, actFileContents)
}

func TestGetRepoConfigReturnsCorrectConfig(t *testing.T) {
	configFileContents := getDefaultJSON()
	origReadFunc := readFile

	defer func() { readFile = origReadFunc }()
	readFile = func(filePath string) ([]byte, error) {
		return configFileContents, nil
	}

	actConfig, err := getRepoConfig("")
	assert.Equal(t, *config, *actConfig)
	assert.Nil(t, err)
}
