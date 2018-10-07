package captaingithook

import (
	"errors"
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

	if err == nil {
		t.Errorf("Expected error but got nil")
	}

	if actualErrMsg := err.Error(); actualErrMsg != expectedErrMsg {
		t.Errorf("Incorrect error message. Expected: %s, but got: %s", expectedErrMsg, actualErrMsg)
	}
}

func TestGetRootDirectoryPathHandlesEmptyDirectoryCorrectly(t *testing.T) {
	expectedErrMsg := "got an unexpected result for the git repo root directory."
	origRunCommand := runCommand
	runCommand = func(cmd string) (string, error) {
		return "", nil
	}
	defer func() { runCommand = origRunCommand }()

	gitDir, err := getRootDirectoryPath()

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
	runCommand = func(cmd string) (string, error) {
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
	var actualCmd string
	origRunCommand := runCommand
	runCommand = func(cmd string) (string, error) {
		actualCmd = cmd
		return "", nil
	}
	defer func() { runCommand = origRunCommand }()
	getGitRepoRootDirectoryPath()

	if actualCmd != expGitRootCommand {
		t.Errorf("Used incorrect command. Expected: %s, but got: %s", expGitRootCommand, actualCmd)
	}
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

	if err == nil {
		t.Errorf("Expected error but got nil")
	}

	if actualErrMsg := err.Error(); actualErrMsg != expectedErrMsg {
		t.Errorf("Incorrect error message. Expected: %s, but got: %s", expectedErrMsg, actualErrMsg)
	}
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

	if actualCmd != expGitHooksCommand {
		t.Errorf("Used incorrect command. Expected: %s, but got: %s", expGitHooksCommand, actualCmd)
	}
}
