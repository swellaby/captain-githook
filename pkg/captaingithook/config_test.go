package captaingithook

import (
	"testing"
)

const defaultConfigFileName = "captaingithook.json"

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
	foundFile := configFileExists("")

	if foundFile {
		t.Errorf("Got wrong result for found config file. Expected %t, but got %t", false, foundFile)
	}
}

func TestConfigFileExistsReturnsTrueWhenFileFound(t *testing.T) {
	origFileExists := fileExists
	fileExists = func(string) bool { return true }
	defer func() { fileExists = origFileExists }()
	foundFile := configFileExists("")

	if !foundFile {
		t.Errorf("Got wrong result for found config file. Expected %t, but got %t", true, foundFile)
	}
}

func TestCreateConfigFileUsesCorrectDefault(t *testing.T) {
	originalWriteFile := writeFile
	defer func() { writeFile = originalWriteFile }()
	var actualFileName, actualData string
	expectedData := ""
	writeFile = func(fileName, data string) error {
		actualFileName = fileName
		actualData = data
		return nil
	}

	createConfigFile("")

	if actualFileName != defaultConfigFileName {
		t.Errorf("Attempted to create wrong config file name. Expected: %s, but got: %s.", defaultConfigFileName, actualFileName)
	}

	if actualData != expectedData {
		t.Errorf("Attempted to create wrong config file contents. Expected: %s, but got: %s.", expectedData, actualData)
	}
}

func TestCreateConfigFileUsesSpecifiedFileName(t *testing.T) {
	originalWriteFile := writeFile
	defer func() { writeFile = originalWriteFile }()
	var actualFileName, actualData string
	expectedData := ""
	writeFile = func(fileName, data string) error {
		actualFileName = fileName
		actualData = data
		return nil
	}

	desiredFileName := ".captain-githookrc"

	createConfigFile(desiredFileName)

	if actualFileName != desiredFileName {
		t.Errorf("Attempted to create wrong config file name. Expected: %s, but got: %s.", defaultConfigFileName, actualFileName)
	}

	if actualData != expectedData {
		t.Errorf("Attempted to create wrong config file contents. Expected: %s, but got: %s.", expectedData, actualData)
	}
}
