package captaingithook

import (
	"bytes"
	"encoding/json"
	"errors"
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
	origGetRepoRoot := getGitRepoRootDirectoryPath
	getGitRepoRootDirectoryPath = func() (string, error) {
		return "", nil
	}
	defer func() { getGitRepoRootDirectoryPath = origGetRepoRoot }()

	initConfigFile("")

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
	origGetRepoRoot := getGitRepoRootDirectoryPath
	getGitRepoRootDirectoryPath = func() (string, error) {
		return "", nil
	}
	defer func() { getGitRepoRootDirectoryPath = origGetRepoRoot }()

	desiredFileName := ".captain-githookrc"

	initConfigFile(desiredFileName)

	if actualFileName != desiredFileName {
		t.Errorf("Attempted to create wrong config file name. Expected: %s, but got: %s.", defaultConfigFileName, actualFileName)
	}

	if !bytes.Equal(actualData, expectedData) {
		t.Errorf("Attempted to create wrong config file contents. Expected: %s, but got: %s.", expectedData, actualData)
	}
}

func TestInitConfigFileReturnsCorrectErrorWhenGitRootNotFound(t *testing.T) {
	origGetRepoRoot := getGitRepoRootDirectoryPath
	getGitRepoRootDirectoryPath = func() (string, error) {
		return "", errors.New("oh no")
	}
	defer func() { getGitRepoRootDirectoryPath = origGetRepoRoot }()
	err := initConfigFile(".captaingithookrc.json")

	if err != errFailedToFindGitRepo {
		t.Errorf("Did not get expected error. Expected: %s, but got: %s.", errFailedToFindGitRepo, err)
	}

	expErrMsg := "encountered a fatal error while trying to determine the root directory of the git repo"
	if err.Error() != expErrMsg {
		t.Errorf("Got wrong error message. Expected: %s, but got: %s.", expErrMsg, err.Error())
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
