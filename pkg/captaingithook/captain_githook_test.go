package captaingithook

import (
	"errors"
	"testing"
)

func TestInitializeReturnsCorrectErrorOnUnknownGitRoot(t *testing.T) {
	expectedErrMsg := "unexpected error encountered while trying to determine the git repo path."
	origFunc := getGitRepoRootDirectoryPath
	defer func() { getGitRepoRootDirectoryPath = origFunc }()
	getGitRepoRootDirectoryPath = func() (string, error) {
		return "", errors.New(expectedErrMsg)
	}

	err := Initialize()

	if actErrMsg := err.Error(); actErrMsg != expectedErrMsg {
		t.Errorf("Did not get correct error. Expected: %s, but got %s", expectedErrMsg, actErrMsg)
	}
}

func TestInitializeReturnsCorrectErrorOnFailedConfigCreation(t *testing.T) {
	origGitFunc := getGitRepoRootDirectoryPath
	defer func() { getGitRepoRootDirectoryPath = origGitFunc }()
	getGitRepoRootDirectoryPath = func() (string, error) {
		return "", nil
	}
	expectedErrMsg := "unexpected error encountered while trying to create config file."
	origInitFunc := initializeCaptainGithookConfigFile
	defer func() { initializeCaptainGithookConfigFile = origInitFunc }()
	initializeCaptainGithookConfigFile = func(string, string) error {
		return errors.New(expectedErrMsg)
	}

	err := Initialize()

	if actErrMsg := err.Error(); actErrMsg != expectedErrMsg {
		t.Errorf("Did not get correct error. Expected: %s, but got %s", expectedErrMsg, actErrMsg)
	}
}

func TestInitializeCorrectlyAddsConfigAndHookFiles(t *testing.T) {
	origGitFunc := getGitRepoRootDirectoryPath
	defer func() { getGitRepoRootDirectoryPath = origGitFunc }()
	getGitRepoRootDirectoryPath = func() (string, error) {
		return "", nil
	}
	origFunc := initializeCaptainGithookConfigFile
	defer func() { initializeCaptainGithookConfigFile = origFunc }()
	initializeCaptainGithookConfigFile = func(path, fileName string) error {
		return nil
	}

	err := Initialize()

	if err != nil {
		t.Errorf("Error was not nil. Error value: %s", err)
	}
}

func TestInitializeWithFileNameReturnsCorrectErrorOnUnknownGitRoot(t *testing.T) {
	expectedErrMsg := "unexpected error encountered while trying to determine the git repo path."
	origFunc := getGitRepoRootDirectoryPath
	defer func() { getGitRepoRootDirectoryPath = origFunc }()
	getGitRepoRootDirectoryPath = func() (string, error) {
		return "", errors.New(expectedErrMsg)
	}

	err := InitializeWithFileName("")

	if actErrMsg := err.Error(); actErrMsg != expectedErrMsg {
		t.Errorf("Did not get correct error. Expected: %s, but got %s", expectedErrMsg, actErrMsg)
	}
}

func TestInitializeWithFileNameReturnsCorrectErrorOnFailedConfigCreation(t *testing.T) {
	origGitFunc := getGitRepoRootDirectoryPath
	defer func() { getGitRepoRootDirectoryPath = origGitFunc }()
	getGitRepoRootDirectoryPath = func() (string, error) {
		return "", nil
	}
	expectedErrMsg := "unexpected error encountered while trying to create config file."
	origInitFunc := initializeCaptainGithookConfigFile
	defer func() { initializeCaptainGithookConfigFile = origInitFunc }()
	initializeCaptainGithookConfigFile = func(string, string) error {
		return errors.New(expectedErrMsg)
	}

	err := InitializeWithFileName("")

	if actErrMsg := err.Error(); actErrMsg != expectedErrMsg {
		t.Errorf("Did not get correct error. Expected: %s, but got %s", expectedErrMsg, actErrMsg)
	}
}

func TestInitializeWithFileNameCorrectlyAddsConfigAndHookFiles(t *testing.T) {
	origGitFunc := getGitRepoRootDirectoryPath
	defer func() { getGitRepoRootDirectoryPath = origGitFunc }()
	getGitRepoRootDirectoryPath = func() (string, error) {
		return "", nil
	}
	origFunc := initializeCaptainGithookConfigFile
	defer func() { initializeCaptainGithookConfigFile = origFunc }()
	actualFileName := ""
	initializeCaptainGithookConfigFile = func(path, fileName string) error {
		actualFileName = fileName
		return nil
	}
	expFileName := ".captaingithookrc.json"
	err := InitializeWithFileName(expFileName)

	if err != nil {
		t.Errorf("Error was not nil. Error value: %s", err)
	}

	if expFileName != actualFileName {
		t.Errorf("Did not get correct config file name. Expected: %s, but got %s", expFileName, actualFileName)
	}
}

func TestFoo(t *testing.T) {
	err := Initialize()
	if err != nil {
		t.Errorf("Error was not nil. Error: %s", err)
	}
}
