package cli

import (
	"errors"
	"fmt"
	"testing"
)

func TestInitCommandConfiguresUseCorrectly(t *testing.T) {
	expUse := "init"
	use := initCmd.Use
	if use != expUse {
		t.Errorf("Did not set correct use value for init command. Expected: %s, but got: %s", expUse, use)
	}
}

func TestInitializeCallsDefaultInitMethodWhenConfigFileFlagNotSet(t *testing.T) {
	expError := errors.New("foobar")
	expGreeting := "Ahoy there matey!"
	expLogMessage := "Initializing your repository with the default config file name."
	actGreeting := ""
	actLogMessage := ""
	logCallCount := 0
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

	if err != expError {
		t.Errorf("Did not get correct error. Expected: %s, but got: %s", expError, err)
	}

	if actGreeting != expGreeting {
		t.Errorf("Did not get correct greeting. Expected: %s, but got: %s", expGreeting, string(actGreeting))
	}

	if actLogMessage != expLogMessage {
		t.Errorf("Did not get correct log message. Expected: %s, but got: %s", expLogMessage, actLogMessage)
	}
}

func TestInitializeCallsDefaultInitMethodWhenConfigFileFlagSet(t *testing.T) {
	expError := errors.New("barfoo")
	actFileName := ""
	expFileName := ".captaingithook.json"
	expGreeting := "Ahoy there matey!"
	expLogMessage := fmt.Sprintf("Initializing your repository with the your requested config file name: '%s'.", expFileName)
	actGreeting := ""
	actLogMessage := ""

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

	if err != expError {
		t.Errorf("Did not get correct error. Expected: %s, but got: %s", expError, err)
	}

	if actFileName != expFileName {
		t.Errorf("Did not pass correct desired config file name. Expected %s, but got: %s", expFileName, actFileName)
	}

	if actGreeting != expGreeting {
		t.Errorf("Did not get correct greeting. Expected: %s, but got: %s", expGreeting, string(actGreeting))
	}

	if actLogMessage != expLogMessage {
		t.Errorf("Did not get correct log message. Expected: %s, but got: %s", expLogMessage, actLogMessage)
	}
}
