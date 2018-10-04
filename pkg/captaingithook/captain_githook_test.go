package captaingithook

import (
	"testing"
)

func TestInitializeCorrectlyAddsConfigAndHookFiles(t *testing.T) {
	origFunc := initializeCaptainGithookConfigFile
	defer func() { initializeCaptainGithookConfigFile = origFunc }()
	actualFileName := ""
	initializeCaptainGithookConfigFile = func(fileName string) error {
		actualFileName = fileName
		return nil
	}
	expFileName := ".captaingithookrc.json"
	err := Initialize(expFileName)

	if err != nil {
		t.Errorf("Error was not nil. Error value: %s", err)
	}

	if expFileName != actualFileName {
		t.Errorf("Did not get correct config file name. Expected: %s, but got %s", expFileName, actualFileName)
	}
}

func TestInitializeReturnsCorrectErrorOnUnknownGitRoot(t *testing.T) {
	origFunc := initializeCaptainGithookConfigFile
	defer func() { initializeCaptainGithookConfigFile = origFunc }()
	initializeCaptainGithookConfigFile = func(fileName string) error {
		return errFailedToFindGitRepo
	}

	err := Initialize("")

	if errFailedToFindGitRepo != err {
		t.Errorf("Did not get correct error. Expected: %s, but got %s", errFailedToFindGitRepo, err)
	}
}
