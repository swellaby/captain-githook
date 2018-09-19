package captaingithook

import (
	"testing"
)

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
		"captaingithook.json",
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
