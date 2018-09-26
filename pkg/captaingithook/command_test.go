package captaingithook

import (
	"os/exec"
	"testing"
)

type cmd struct {}

// func (c *cmd) CombinedOutput() ()

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

func TestNewCommandUsesDirectoryWhenSpecified(t *testing.T) {
	dir := "/usr/some/other/directory"
	mockCmd := &exec.Cmd{}
	osCommand = func(name string, arg ...string) *exec.Cmd {
		mockCmd.Path = name
		mockCmd.Args = append([]string{name}, arg...)
		return mockCmd
	}
	name, cmdSwitch := "sh", "-c"
	arg := []string{ cmdSwitch, "echo", "foobar" }
	cmd := newCommand(dir, name, arg...)
	defer func() { osCommand = exec.Command }()

	if mockCmd.Dir != dir {
		t.Errorf("Target directory for command was incorrect. Expected: %s, but got: %s.", dir, mockCmd.Dir)
	}

	if cmd == nil {
		t.Errorf("Got a nil exec.Command object.")
	}
}

func TestNewCommandUsesCallingProcDirectoryWhenNotSpecified(t *testing.T) {
	dir := ""
	mockCmd := &exec.Cmd{}
	osCommand = func(name string, arg ...string) *exec.Cmd {
		mockCmd.Path = name
		mockCmd.Args = append([]string{name}, arg...)
		return mockCmd
	}
	name, cmdSwitch := "sh", "-c"
	arg := []string{ cmdSwitch, "echo", "foobar" }
	cmd := newCommand(dir, name, arg...)
	defer func() { osCommand = exec.Command }()

	if mockCmd.Dir != dir {
		t.Errorf("Target directory for command was incorrect. Expected: %s, but got: %s.", dir, mockCmd.Dir)
	}

	if cmd == nil {
		t.Errorf("Got a nil exec.Command object.")
	}
}

// func TestRunSetsDirectoryWhenSpecified(t *testing.T) {
// 	mockCmd := &command{}
// 	createCommand = func(name string, arg ...string) *exec.Cmd {
// 		return nil
// 	}
// 	defer func() { createCommand = exec.Command }()
// }
