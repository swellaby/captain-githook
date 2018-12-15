package captaingithook

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitializeReturnsCorrectErrorOnUnknownGitRoot(t *testing.T) {
	expectedErrMsg := "unexpected error encountered while trying to determine the git repo path."
	origFunc := getGitRepoRootDirectoryPath
	defer func() { getGitRepoRootDirectoryPath = origFunc }()
	getGitRepoRootDirectoryPath = func() (string, error) {
		return "", errors.New(expectedErrMsg)
	}
	assert.Error(t, Initialize(), expectedErrMsg)
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

	assert.Error(t, Initialize(), expectedErrMsg)
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

	assert.Equal(t, errInvalidGitHooksDirectoryPath, Initialize())
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

	assert.Nil(t, Initialize())
}

func TestInitializeWithFileNameReturnsCorrectErrorOnUnknownGitRoot(t *testing.T) {
	expectedErrMsg := "unexpected error encountered while trying to determine the git repo path."
	origFunc := getGitRepoRootDirectoryPath
	defer func() { getGitRepoRootDirectoryPath = origFunc }()
	getGitRepoRootDirectoryPath = func() (string, error) {
		return "", errors.New(expectedErrMsg)
	}

	assert.Error(t, InitializeWithFileName(""), expectedErrMsg)
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

	assert.Error(t, InitializeWithFileName(""), expectedErrMsg)
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

	assert.Equal(t, errInvalidGitHooksDirectoryPath, InitializeWithFileName(""))
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
	assert.Nil(t, InitializeWithFileName(expFileName))
	assert.Equal(t, expFileName, actualFileName, "Did not get correct config file name. Expected: %s, but got %s", expFileName, actualFileName)
}

func TestRunHookReturnsCorrectErrorOnUnknownGitRoot(t *testing.T) {
	expectedErr := errors.New("0978234")
	origFunc := getGitRepoRootDirectoryPath
	defer func() { getGitRepoRootDirectoryPath = origFunc }()
	getGitRepoRootDirectoryPath = func() (string, error) {
		return "", expectedErr
	}

	output, err := RunHook("")
	assert.Equal(t, expectedErr, err)
	assert.Equal(t, "", output, "Output was not the expected empty string. Output: %s", output)
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
	assert.Equal(t, expError, err)
	assert.Equal(t, "", output, "Output was not the expected empty string. Output: %s", output)
	assert.Equal(t, expDirPath, actDirPath, "Did not use correct directory path to locate config. Expected: %s, but got %s", expDirPath, actDirPath)
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
	assert.Equal(t, expError, err)
	assert.Equal(t, expOutput, output, "Output was not the expected value of: %s, but instead was: %s", expOutput, output)
	assert.Equal(t, expDirPath, actDirPath, "Did not use correct directory path to locate config. Expected: %s, but got %s", expDirPath, actDirPath)
	assert.Equal(t, expHookName, actHookName, "Did not use correct hook name. Expected: %s, but got %s", expHookName, actHookName)
	assert.Equal(t, *runnerConfig, *actConfig, "Did not use correct captain-githook config. Expected: %v, but got %v", *runnerConfig, *actConfig)
}
