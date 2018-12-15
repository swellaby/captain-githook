package cli

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitCommandConfiguresUseCorrectly(t *testing.T) {
	assert.Equal(t, "init", initCmd.Use)
}

func TestInitializeCallsDefaultInitMethodWhenConfigFileFlagNotSet(t *testing.T) {
	expError := errors.New("foobar")
	expGreeting := "Ahoy there matey!"
	expLogMessage := "Initializing your repository with the default config file name."
	actGreeting := ""
	actLogMessage := ""
	logCallCount := 0
	origIsCi := isCI
	defer func() { isCI = origIsCi }()
	isCI = func() bool {
		return false
	}
	origFunc := initializeCaptainGithook
	defer func() { initializeCaptainGithook = origFunc }()
	initializeCaptainGithook = func() error {
		return expError
	}
	origLogFunc := log
	defer func() { log = origLogFunc }()
	log = func(contents ...interface{}) (int, error) {
		if logCallCount == 0 {
			actGreeting = fmt.Sprint(contents[0])
		} else {
			actLogMessage = fmt.Sprint(contents[0])
		}
		logCallCount++
		return 0, nil
	}

	err := initialize(nil, nil)
	assert.Equal(t, expError, err, "Did not get correct error. Expected: %s, but got: %s", expError, err)
	assert.Equal(t, expGreeting, actGreeting, "Did not get correct greeting. Expected: %s, but got: %s", expGreeting, string(actGreeting))
	assert.Equal(t, expLogMessage, actLogMessage, "Did not get correct log message. Expected: %s, but got: %s", expLogMessage, actLogMessage)
}

func TestInitializeCallsDefaultInitMethodWhenConfigFileFlagSet(t *testing.T) {
	expError := errors.New("barfoo")
	actFileName := ""
	expFileName := ".captaingithook.json"
	expGreeting := "Ahoy there matey!"
	expLogMessage := fmt.Sprintf("Initializing your repository with the your requested config file name: '%s'.", expFileName)
	actGreeting := ""
	actLogMessage := ""
	origIsCi := isCI
	defer func() { isCI = origIsCi }()
	isCI = func() bool {
		return false
	}
	initConfigFileName = expFileName
	origFunc := initializeCaptainGithookWithConfigName
	defer func() { initializeCaptainGithookWithConfigName = origFunc }()
	initializeCaptainGithookWithConfigName = func(filename string) error {
		actFileName = filename
		return expError
	}
	origLogFunc := log
	defer func() { log = origLogFunc }()
	log = func(contents ...interface{}) (int, error) {
		actGreeting = fmt.Sprint(contents[0])
		return 0, nil
	}

	origLogfFunc := logf
	defer func() { logf = origLogfFunc }()
	logf = func(format string, contents ...interface{}) (int, error) {
		actLogMessage = fmt.Sprintf(format, contents[0])
		return 0, nil
	}

	err := initialize(nil, nil)
	assert.Equal(t, expError, err, "Did not get correct error. Expected: %s, but got: %s", expError, err)
	assert.Equal(t, expFileName, actFileName, "Did not pass correct desired config file name. Expected %s, but got: %s", expFileName, actFileName)
	assert.Equal(t, expGreeting, actGreeting, "Did not get correct greeting. Expected: %s, but got: %s", expGreeting, string(actGreeting))
	assert.Equal(t, expLogMessage, actLogMessage, "Did not get correct log message. Expected: %s, but got: %s", expLogMessage, actLogMessage)
}

func TestShouldNotInitializeHooksWhenCIDetected(t *testing.T) {
	actLogMessage := ""
	expLogMessage := "CI environment detected. Skipping git hook install."
	initFuncCalled := false
	origIsCi := isCI
	defer func() { isCI = origIsCi }()
	isCI = func() bool {
		return true
	}
	origLogFunc := log
	defer func() { log = origLogFunc }()
	log = func(contents ...interface{}) (int, error) {
		actLogMessage = fmt.Sprint(contents[0])
		return 0, nil
	}
	origInitNameFunc := initializeCaptainGithookWithConfigName
	defer func() { initializeCaptainGithookWithConfigName = origInitNameFunc }()
	initializeCaptainGithookWithConfigName = func(string) error {
		initFuncCalled = true
		return nil
	}
	origInitFunc := initializeCaptainGithook
	defer func() { initializeCaptainGithook = origInitFunc }()
	initializeCaptainGithookWithConfigName = func(string) error {
		initFuncCalled = true
		return nil
	}

	initialize(nil, nil)
	assert.False(t, initFuncCalled, "Initialization function should not have been called")
	assert.Equal(t, expLogMessage, actLogMessage, "Did not get correct log message. Expected: %s, but got: %s", expLogMessage, actLogMessage)
}
