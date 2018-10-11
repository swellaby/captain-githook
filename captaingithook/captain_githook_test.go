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

func TestInitializeReturnsCorrectErrorOnFailedHookFileCreation(t *testing.T) {
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
	origHookFunc := initializeGitHookFiles
	defer func() { initializeGitHookFiles = origHookFunc }()
	initializeGitHookFiles = func() error {
		return errInvalidGitHooksDirectoryPath
	}

	if err := Initialize(); err != errInvalidGitHooksDirectoryPath {
		t.Errorf("Did not get correct error. Expected: %s, but got %s", errInvalidGitHooksDirectoryPath, err)
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
	origHookFunc := initializeGitHookFiles
	defer func() { initializeGitHookFiles = origHookFunc }()
	initializeGitHookFiles = func() error {
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

func TestInitializeWithFileNameReturnsCorrectErrorOnFailedHookFileCreation(t *testing.T) {
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
	origHookFunc := initializeGitHookFiles
	defer func() { initializeGitHookFiles = origHookFunc }()
	initializeGitHookFiles = func() error {
		return errInvalidGitHooksDirectoryPath
	}

	err := InitializeWithFileName("")

	if err != errInvalidGitHooksDirectoryPath {
		t.Errorf("Did not get correct error. Expected: %s, but got %s", errInvalidGitHooksDirectoryPath, err)
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
	origHookFunc := initializeGitHookFiles
	defer func() { initializeGitHookFiles = origHookFunc }()
	initializeGitHookFiles = func() error {
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

func TestRunHookReturnsCorrectErrorOnUnknownGitRoot(t *testing.T) {
	expectedErr := errors.New("0978234")
	origFunc := getGitRepoRootDirectoryPath
	defer func() { getGitRepoRootDirectoryPath = origFunc }()
	getGitRepoRootDirectoryPath = func() (string, error) {
		return "", expectedErr
	}

	output, err := RunHook("")

	if err != expectedErr {
		t.Errorf("Did not get correct error. Expected: %s, but got %s", expectedErr, err)
	}

	if output != "" {
		t.Errorf("Output was not the expected empty string. Output: %s", output)
	}
}

func TestRunHookReturnsCorrectErrorOnConfigLoadError(t *testing.T) {
	expDirPath := "/foo"
	actDirPath := ""
	expError := errors.New("adfadf")
	origGitFunc := getGitRepoRootDirectoryPath
	defer func() { getGitRepoRootDirectoryPath = origGitFunc }()
	getGitRepoRootDirectoryPath = func() (string, error) {
		return expDirPath, nil
	}
	origConfigFunc := getCaptainGithookConfig
	defer func() { getCaptainGithookConfig = origConfigFunc }()
	getCaptainGithookConfig = func(path string) (*Config, error) {
		actDirPath = path
		return nil, expError
	}

	output, err := RunHook("")

	if err != expError {
		t.Errorf("Did not get correct error. Expected: %s, but got %s", expError, err)
	}

	if output != "" {
		t.Errorf("Output was not the expected empty string. Output: %s", output)
	}

	if actDirPath != expDirPath {
		t.Errorf("Did not use correct directory path to locate config. Expected: %s, but got %s", expDirPath, actDirPath)
	}
}

func TestRunHookReturnsCorrectResultOfScriptExecution(t *testing.T) {
	var actConfig *Config
	expDirPath := "/usr/bar"
	actDirPath := ""
	expHookName := "commit-msg"
	actHookName := ""
	expError := errors.New("abc123")
	expOutput := "golint ./..."
	origGitFunc := getGitRepoRootDirectoryPath
	defer func() { getGitRepoRootDirectoryPath = origGitFunc }()
	getGitRepoRootDirectoryPath = func() (string, error) {
		return expDirPath, nil
	}
	origConfigFunc := getCaptainGithookConfig
	defer func() { getCaptainGithookConfig = origConfigFunc }()
	getCaptainGithookConfig = func(path string) (*Config, error) {
		return runnerConfig, nil
	}
	origRunFunc := runHookScript
	defer func() { runHookScript = origRunFunc }()
	runHookScript = func(cfg *Config, hn, dir string) (string, error) {
		actConfig = cfg
		actHookName = hn
		actDirPath = dir
		return expOutput, expError
	}

	output, err := RunHook(expHookName)

	if err != expError {
		t.Errorf("Did not get correct error. Expected: %s, but got %s", expError, err)
	}

	if output != expOutput {
		t.Errorf("Output was not the expected empty string. Output: %s", output)
	}

	if actDirPath != expDirPath {
		t.Errorf("Did not use correct directory path to run hook. Expected: %s, but got %s", expDirPath, actDirPath)
	}

	if actHookName != expHookName {
		t.Errorf("Did not use correct hook name. Expected: %s, but got %s", expHookName, actHookName)
	}

	if *actConfig != *runnerConfig {
		t.Errorf("Did not use correct captain-githook config. Expected: %v, but got %v", *runnerConfig, *actConfig)
	}
}
