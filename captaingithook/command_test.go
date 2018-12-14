package captaingithook

import (
	"github.com/stretchr/testify/assert"
	"os/exec"
	"testing"
)

type MockCommand struct {
	CombinedOutputFunc func() ([]byte, error)
}

func (m MockCommand) CombinedOutput() ([]byte, error) {
	if m.CombinedOutputFunc != nil {
		return m.CombinedOutputFunc()
	}
	return nil, nil
}

const echoScript = "echo foobar"

var echoCmd = [...]string{"sh", "-c", echoScript}
var expNumArgs = len(echoCmd)

func TestGetRunnerInfoReturnsCorrectValueOnWindows(t *testing.T) {
	const expectedRunner = "cmd.exe"
	const expectedRunnerArg = "/C"
	runner, runnerArg := getRunnerInfo("windows")
	assert.Equal(t, expectedRunner, runner)
	assert.Equal(t, expectedRunnerArg, runnerArg)
}

func TestGetRunnerInfoReturnsCorrectValueOnNonWindows(t *testing.T) {
	const expectedRunner = "sh"
	const expectedRunnerArg = "-c"
	nonWindowsOperatingSystems := []string{
		"linux",
		"darwin",
		"freebsd",
	}

	for _, os := range nonWindowsOperatingSystems {
		runner, runnerArg := getRunnerInfo(os)
		assert.Equal(t, expectedRunner, runner, "Runner was incorrect for OS: %s. Expected: %s, but got: %s.", os, expectedRunner, runner)
		assert.Equal(t, expectedRunnerArg, runnerArg, "Runner Arg was incorrect for OS: %s. Expected: %s, but got: %s.", os, expectedRunnerArg, runnerArg)
	}
}

func TestNewCommandUsesDirectoryWhenSpecified(t *testing.T) {
	dir := "/usr/some/other/directory"
	mockCmd := &exec.Cmd{}
	osCommand = func(name string, arg ...string) *exec.Cmd {
		mockCmd.Path = name
		mockCmd.Args = append([]string{name}, arg...)
		return mockCmd
	}
	defer func() { osCommand = exec.Command }()
	assert.NotNil(t, newCommand(dir, echoScript))
	assert.Equal(t, dir, mockCmd.Dir, "Target directory for command was incorrect. Expected: %s, but got: %s.", dir, mockCmd.Dir)
	numArgs := len(mockCmd.Args)
	assert.Equal(t, expNumArgs, numArgs, "Did not get correct number of command args. Expected: %d, but got: %d", expNumArgs, numArgs)
}

func TestNewCommandUsesCallingProcDirectoryWhenNotSpecified(t *testing.T) {
	dir := ""
	mockCmd := &exec.Cmd{}
	osCommand = func(name string, arg ...string) *exec.Cmd {
		mockCmd.Path = name
		mockCmd.Args = append([]string{name}, arg...)
		return mockCmd
	}
	defer func() { osCommand = exec.Command }()
	assert.NotNil(t, newCommand(dir, echoScript))
	assert.Equal(t, dir, mockCmd.Dir, "Target directory for command was incorrect. Expected: %s, but got: %s.", dir, mockCmd.Dir)
	numArgs := len(mockCmd.Args)
	assert.Equal(t, expNumArgs, numArgs, "Did not get correct number of command args. Expected: %d, but got: %d", expNumArgs, numArgs)
}

func TestRunReturnsCorrectResults(t *testing.T) {
	mockBytes := []byte("foobar")
	createCommand = func(directory, script string) command {
		return &MockCommand{CombinedOutputFunc: func() ([]byte, error) {
			return mockBytes, nil
		}}
	}
	defer func() { createCommand = newCommand }()

	result, err := run("")
	assert.Nil(t, err)
	assert.Equal(t, string(mockBytes), result, "Result from run was incorrect. Expected: %s, but got: %s.", result, string(mockBytes))
}
