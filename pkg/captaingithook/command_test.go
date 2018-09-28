package captaingithook

import (
	"os/exec"
	"testing"
)

type MockCommand struct {
	CombinedOutputFunc func()([]byte, error)
}

func (m MockCommand) CombinedOutput()([]byte, error) {
	if m.CombinedOutputFunc != nil {
		return m.CombinedOutputFunc()
	}
	return nil, nil
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

func TestNewCommandUsesDirectoryWhenSpecified(t *testing.T) {
	dir := "/usr/some/other/directory"
	mockCmd := &exec.Cmd{}
	osCommand = func(name string, arg ...string) *exec.Cmd {
		mockCmd.Path = name
		mockCmd.Args = append([]string{name}, arg...)
		return mockCmd
	}
	defer func() { osCommand = exec.Command }()
	name, cmdSwitch := "sh", "-c"
	arg := []string{ cmdSwitch, "echo", "foobar" }

	if newCommand(dir, name, arg...) == nil {
		t.Errorf("Got a nil exec.Command object.")
	}

	if mockCmd.Dir != dir {
		t.Errorf("Target directory for command was incorrect. Expected: %s, but got: %s.", dir, mockCmd.Dir)
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
	defer func() { osCommand = exec.Command }()

	name, cmdSwitch := "sh", "-c"
	arg := []string{ cmdSwitch, "echo", "foobar" }

	if newCommand(dir, name, arg...) == nil {
		t.Errorf("Got a nil exec.Command object.")
	}

	if mockCmd.Dir != dir {
		t.Errorf("Target directory for command was incorrect. Expected: %s, but got: %s.", dir, mockCmd.Dir)
	}
}

func TestRunReturnsCorrectResults(t *testing.T) {
	mockBytes := []byte("foobar")
	createCommand = func(directory, name string, args ...string) command {
		return &MockCommand{ CombinedOutputFunc: func() ([]byte, error) {
			return mockBytes, nil
		}}
	}
	defer func() { createCommand = newCommand }()

	result, err := run("", "")

	if err != nil {
		t.Errorf("Err was not nil. Got: %s.", err)
	}

	if result != string(mockBytes) {
		t.Errorf("Result from run was incorrect. Expected: %s, but got: %s.", result, string(mockBytes))
	}
}
