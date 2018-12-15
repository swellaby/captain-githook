package cli

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunCommandConfiguresUseCorrectly(t *testing.T) {
	assert.Equal(t, "run", runCmd.Use)
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
	assert.Equal(t, expHookLogfOutput, actHookLogfOutput, "Did not get correct hook script output. Expected: %s, but got: %s", expOutput, actLogOutput)
	assert.Equal(t, expOutput, actLogOutput, "Did not get correct hook script output. Expected: %s, but got: %s", expOutput, actLogOutput)
	assert.Equal(t, expErr, actErr, "Did not get correct error from hook run. Expected: %s, but got: %s", expErr, actErr)
	assert.Equal(t, expHookName, actHookName, "Did not use correct hook name. Expected: %s, but got: %s", expHookName, actHookName)
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
	assert.False(t, logCalled, "Log was not supposed to be called but was. Log output: %s", loggedContent)
}
