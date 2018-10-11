package cli

import (
	"fmt"
	"testing"
)

func TestRunCommandConfiguresUseCorrectly(t *testing.T) {
	expUse := "run"
	use := runCmd.Use
	if use != expUse {
		t.Errorf("Did not set correct use value for run command. Expected: %s, but got: %s", expUse, use)
	}
}

func TestRunCommandExecutesCorrectly(t *testing.T) {
	actHookName := ""
	expHookName := "commit-msg"
	hookName = expHookName
	expOutput := "hello world"
	actLogOutput := ""
	actHookLogfOutput := ""
	expHookLogfOutput := fmt.Sprintf("Running hook: '%s'...\n", expHookName)
	expErr := fmt.Errorf("foobar")
	origLogFunc := log
	defer func() { log = origLogFunc }()
	log = func(contents ...interface{}) (int, error) {
		actLogOutput = fmt.Sprint(contents[0])
		return 0, nil
	}
	origLogfFunc := logf
	defer func() { logf = origLogfFunc }()
	logf = func(template string, contents ...interface{}) (int, error) {
		actHookLogfOutput = fmt.Sprintf(template, contents[0])
		return 0, nil
	}
	origRunFunc := runGitHook
	defer func() { runGitHook = origRunFunc }()
	runGitHook = func(hook string) (string, error) {
		actHookName = hook
		return expOutput, expErr
	}

	actErr := runHook(nil, nil)

	if actHookLogfOutput != expHookLogfOutput {
		t.Errorf("Did not get correct hook script output. Expected: %s, but got: %s", expOutput, actLogOutput)
	}

	if actLogOutput != expOutput {
		t.Errorf("Did not get correct hook script output. Expected: %s, but got: %s", expOutput, actLogOutput)
	}

	if actErr != expErr {
		t.Errorf("Did not get correct error from hook run. Expected: %s, but got: %s", expErr, actErr)
	}

	if actHookName != expHookName {
		t.Errorf("Did not use correct hook name. Expected: %s, but got: %s", expHookName, actHookName)
	}
}

func TestRunHookDoesNotLogEmptyOutput(t *testing.T) {
	logCalled := false
	loggedContent := ""
	origLogFunc := log
	defer func() { log = origLogFunc }()
	log = func(contents ...interface{}) (int, error) {
		logCalled = true
		loggedContent = fmt.Sprint(contents[0])
		return 0, nil
	}
	origRunFunc := runGitHook
	defer func() { runGitHook = origRunFunc }()
	runGitHook = func(hook string) (string, error) {
		return "", nil
	}

	runHook(nil, nil)

	if logCalled {
		t.Errorf("Log was not supposed to be called but was. Log output: %s", loggedContent)
	}
}
