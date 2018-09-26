package captaingithook

import (
	"os/exec"
	"testing"
)

type CommandRunner interface {
	Run() string
}

func TestGetRunnerInfoReturnsCorrectValueOnWindows(t *testing.T) {
	const expectedRunner = "cmd.exe"
	const expectedRunnerArg = "/C"
	runner, runnerArg := getRunnerInfo("windows")

	if runner != expectedRunner {
		t.Errorf("Runner was incorrect. Expected: %s, but got: %s.", expectedRunner, runner)
	}

	if runnerArg != expectedRunnerArg {
		t.Errorf("Runner Arg was incorrect. Expected: %s, but got: %s.", expectedRunnerArg, runnerArg)
	}
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

		if runner != expectedRunner {
			t.Errorf("Runner was incorrect for OS: %s. Expected: %s, but got: %s.", os, expectedRunner, runner)
		}

		if runnerArg != expectedRunnerArg {
			t.Errorf("Runner Arg was incorrect for OS: %s. Expected: %s, but got: %s.", os, expectedRunnerArg, runnerArg)
		}
	}
}

func TestRunSetsDirectoryWhenSpecified(t *testing.T) {
	createCommand = func(name string, arg ...string) *exec.Cmd {
		return nil
	}
	defer func() { createCommand = exec.Command }()
}
