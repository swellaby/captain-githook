package captaingithook

import (
	"errors"
	"testing"
)

func TestGetGitRepoRootDirectoryPathHandlesErrorCorrectly(t *testing.T) {
	errMsgDetails := "not a git repo"
	expectedErrMsg := "unexpected error encountered while trying to determine the git repo root directory. Error details: " + errMsgDetails
	origRunCommand := runCommand
	runCommand = func(cmd string, cmdArgs ...string) (string, error) {
		return "", errors.New(errMsgDetails)
	}
	defer func() { runCommand = origRunCommand }()

	_, err := getGitRepoRootDirectoryPath()

	if err == nil {
		t.Errorf("Expected error but got nil")
	}

	if actualErrMsg := err.Error(); actualErrMsg != expectedErrMsg {
		t.Errorf("Incorrect error message. Expected: %s, but got: %s", expectedErrMsg, actualErrMsg)
	}
}

func TestGetGitRepoRootDirectoryPathHandlesEmptyDirectoryCorrectly(t *testing.T) {
	expectedErrMsg := "got an unexpected result for the git repo root directory."
	origRunCommand := runCommand
	runCommand = func(cmd string, cmdArgs ...string) (string, error) {
		return "", nil
	}
	defer func() { runCommand = origRunCommand }()

	gitDir, err := getGitRepoRootDirectoryPath()

	if err == nil {
		t.Errorf("Expected error but got nil")
	}

	if gitDir != "" {
		t.Errorf("Expected empty string for directory but got: %s", gitDir)
	}

	if actualErrMsg := err.Error(); actualErrMsg != expectedErrMsg {
		t.Errorf("Incorrect error message. Expected: %s, but got: %s", expectedErrMsg, actualErrMsg)
	}
}

func TestGetGitRepoRootDirectoryPathReturnsDirectoryCorrectly(t *testing.T) {
	expectedGitDir := "/usr/repos/captain-githook"
	origRunCommand := runCommand
	runCommand = func(cmd string, cmdArgs ...string) (string, error) {
		return expectedGitDir, nil
	}
	defer func() { runCommand = origRunCommand }()

	gitDir, err := getGitRepoRootDirectoryPath()

	if err != nil {
		t.Errorf("Expected nil for error but got: %s", err)
	}

	if gitDir != expectedGitDir {
		t.Errorf("Incorrect git directory. Expected: %s, but got: %s", expectedGitDir, gitDir)
	}
}

func TestGetGitRepoRootDirectoryUsesCorrectCommand(t *testing.T) {
	expectedCmd := "git"
	expectedCmdArgs := []string{ "rev-parse", "--show-toplevel" }
	var actualCmd string
	var actualCmdArgs []string
	origRunCommand := runCommand
	runCommand = func(cmd string, cmdArgs ...string) (string, error) {
		actualCmd = cmd
		actualCmdArgs = cmdArgs
		return "", nil
	}
	defer func() { runCommand = origRunCommand }()
	getGitRepoRootDirectoryPath()

	if actualCmd != expectedCmd {
		t.Errorf("Used incorrect command. Expected: %s, but got: %s", expectedCmd, actualCmd)
	}

	aCount := len(actualCmdArgs)
	eCount := len(expectedCmdArgs)

	if aCount != eCount {
		t.Errorf("Used incorrect number of commandArgs. Expected: %d, but got: %d", eCount, aCount)
	}

	if actualCmdArgs[0] != expectedCmdArgs[0] {
		t.Errorf("Used incorrect command switch. Expected: %s, but got: %s", expectedCmdArgs[0], actualCmdArgs[0])
	}

	if actualCmdArgs[1] != expectedCmdArgs[1] {
		t.Errorf("Used incorrect command switch. Expected: %s, but got: %s", expectedCmdArgs[1], actualCmdArgs[1])
	}
}
