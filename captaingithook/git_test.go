package captaingithook

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

const expGitRootCommand = "git rev-parse --show-toplevel"
const expGitHooksCommand = "git rev-parse --git-path hooks"

func TestGetRootDirectoryPathHandlesErrorCorrectly(t *testing.T) {
	errMsgDetails := "not a git repo"
	expectedErrMsg := "unexpected error encountered while trying to determine the git repo root directory. Error details: " + errMsgDetails
	origRunCommand := runCommand
	runCommand = func(cmd string) (string, error) {
		return "", errors.New(errMsgDetails)
	}
	defer func() { runCommand = origRunCommand }()

	_, err := getRootDirectoryPath()
	assert.Error(t, err, expectedErrMsg)
}

func TestGetRootDirectoryPathHandlesEmptyDirectoryCorrectly(t *testing.T) {
	expectedErrMsg := "got an unexpected result for the git repo root directory."
	origRunCommand := runCommand
	runCommand = func(cmd string) (string, error) {
		return "", nil
	}
	defer func() { runCommand = origRunCommand }()

	gitDir, err := getRootDirectoryPath()
	assert.Error(t, err, expectedErrMsg)
	assert.Equal(t, "", gitDir)
}

func TestGetGitRepoRootDirectoryPathReturnsDirectoryCorrectly(t *testing.T) {
	expectedGitDir := "/usr/repos/captain-githook"
	origRunCommand := runCommand
	runCommand = func(cmd string) (string, error) {
		return expectedGitDir, nil
	}
	defer func() { runCommand = origRunCommand }()

	gitDir, err := getGitRepoRootDirectoryPath()
	assert.Nil(t, err)
	assert.Equal(t, expectedGitDir, gitDir)
}

func TestGetGitRepoRootDirectoryUsesCorrectCommand(t *testing.T) {
	var actualCmd string
	origRunCommand := runCommand
	runCommand = func(cmd string) (string, error) {
		actualCmd = cmd
		return "", nil
	}
	defer func() { runCommand = origRunCommand }()
	getGitRepoRootDirectoryPath()
	assert.Equal(t, expGitRootCommand, actualCmd, "Used incorrect command. Expected: %s, but got: %s", expGitRootCommand, actualCmd)
}

func TestGetHooksDirectoryReturnsErrorWhenCommandFails(t *testing.T) {
	errMsgDetails := "ouch"
	expectedErrMsg := "unexpected error encountered while trying to determine the git repo hooks directory. Error details: " + errMsgDetails
	origRunCommand := runCommand
	runCommand = func(cmd string) (string, error) {
		return "", errors.New(errMsgDetails)
	}
	defer func() { runCommand = origRunCommand }()

	_, err := getHooksDirectory()
	assert.Error(t, err, expectedErrMsg)
}

func TestGetHooksDirectoryUsesCorrectCommand(t *testing.T) {
	var actualCmd string
	origRunCommand := runCommand
	runCommand = func(cmd string) (string, error) {
		actualCmd = cmd
		return "", nil
	}
	defer func() { runCommand = origRunCommand }()
	getHooksDirectory()
	assert.Equal(t, expGitHooksCommand, actualCmd, "Used incorrect command. Expected: %s, but got: %s", expGitHooksCommand, actualCmd)
}
