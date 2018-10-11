package captaingithook

import (
	"errors"
	"fmt"
	"testing"
)

const applyPatchMsgScript = "echo applypatch-msg"
const preApplyPatchScript = "echo pre-applypatch"
const postApplyPatchScript = "echo post-applypatch"
const preCommitScript = "echo pre-commit"
const prepareCommitMsgScript = "echo prepare-commit-msg"
const commitMsgScript = "echo commit-msg"
const postCommitMsgScript = "echo post-commit"
const preRebaseScript = "echo pre-rebase"
const postCheckoutScript = "echo post-checkout"
const postMergeScript = "echo post-merge"
const prePushScript = "echo pre-push"
const preReceiveScript = "echo pre-receive"
const updateScript = "echo update"
const postReceiveScript = "echo post-receive"
const postUpdateScript = "echo post-update"
const pushToCheckoutScript = "echo push-to-checkout"
const preAutoGcScript = "echo pre-auto-gc"
const postRewriteScript = "echo post-rewrite"
const sendEmailValidateScript = "echo sendemail-validate"

var runnerHooks = &HooksConfig{
	ApplyPatchMsg:     applyPatchMsgScript,
	PreApplyPatch:     preApplyPatchScript,
	PostApplyPatch:    postApplyPatchScript,
	PreCommit:         preCommitScript,
	PrepareCommitMsg:  prepareCommitMsgScript,
	CommitMsg:         commitMsgScript,
	PostCommit:        postCommitMsgScript,
	PreRebase:         preRebaseScript,
	PostCheckout:      postCheckoutScript,
	PostMerge:         postMergeScript,
	PrePush:           prePushScript,
	PreReceive:        preReceiveScript,
	Update:            updateScript,
	PostReceive:       postReceiveScript,
	PostUpdate:        postUpdateScript,
	PushToCheckout:    pushToCheckoutScript,
	PreAutoGc:         preAutoGcScript,
	PostRewrite:       postRewriteScript,
	SendEmailValidate: sendEmailValidateScript,
}
var runnerConfig = &Config{Hooks: *runnerHooks}

var hookScriptTests = []struct {
	hookName      string
	expHookScript string
}{
	{"applypatch-msg", applyPatchMsgScript},
	{"pre-applypatch", preApplyPatchScript},
	{"post-applypatch", postApplyPatchScript},
	{"pre-commit", preCommitScript},
	{"prepare-commit-msg", prepareCommitMsgScript},
	{"commit-msg", commitMsgScript},
	{"post-commit", postCommitMsgScript},
	{"pre-rebase", preRebaseScript},
	{"post-checkout", postCheckoutScript},
	{"post-merge", postMergeScript},
	{"pre-push", prePushScript},
	{"pre-receive", preReceiveScript},
	{"update", updateScript},
	{"post-receive", postReceiveScript},
	{"post-update", postUpdateScript},
	{"push-to-checkout", pushToCheckoutScript},
	{"pre-auto-gc", preAutoGcScript},
	{"post-rewrite", postRewriteScript},
	{"sendemail-validate", sendEmailValidateScript},
}

func TestGetConfiguredHooksScriptReturnsCorrectErrorOnNilHooksConfig(t *testing.T) {
	hookScript, err := getConfiguredHookScript(nil, "foobar")

	if hookScript != "" {
		t.Errorf("Did not get correct hook script value. Expected empty string, but got: %s", hookScript)
	}

	if err != errNilHooksConfig {
		t.Errorf("Did not get correct error message. Expected: %s, but got: %s", errNilHooksConfig, err)
	}
}

func TestGetConfiguredHookScriptReturnsCorrectErrorOnUnknownHook(t *testing.T) {
	hookName := "foo"
	hookScript, err := getConfiguredHookScript(runnerHooks, hookName)
	expErrMsg := fmt.Errorf("unknown hook name: %s", hookName).Error()

	if hookScript != "" {
		t.Errorf("Did not get correct hook script value. Expected empty string, but got: %s", hookScript)
	}

	if errMsg := err.Error(); errMsg != expErrMsg {
		t.Errorf("Did not get correct error message. Expected: %s, but got: %s", expErrMsg, errMsg)
	}
}

func TestGetConfiguredHookReturnsCorrectScriptForHook(t *testing.T) {
	for _, hst := range hookScriptTests {
		expHookScript := hst.expHookScript
		hook := hst.hookName
		actHookScript, err := getConfiguredHookScript(runnerHooks, hook)

		if err != nil {
			t.Errorf("Error was not nil for hook: %s. Error: %s", hook, err)
		}

		if actHookScript != expHookScript {
			t.Errorf("Did not get correct script for hook: %s. Expected: %s, but got: %s", hook, expHookScript, actHookScript)
		}
	}
}

func TestRunHookReturnsCorrectErrorWhenConfigIsNil(t *testing.T) {
	output, err := runHook(nil, "", "")

	if err != errNilConfig {
		t.Errorf("Did not get correct error. Expected: %s, but got: %s", errNilConfig, err)
	}

	if output != "" {
		t.Errorf("Output was not an empty string. Output: %s", output)
	}
}

func TestRunHookReturnsCorrectErrorWhenScriptCannotBeDetermined(t *testing.T) {
	expErrMsg := fmt.Errorf("unknown hook name: %s", "").Error()
	output, err := runHook(&Config{}, "", "")

	if errMsg := err.Error(); errMsg != expErrMsg {
		t.Errorf("Did not get correct error. Expected: %s, but got: %s", expErrMsg, errMsg)
	}

	if output != "" {
		t.Errorf("Output was not an empty string. Output: %s", output)
	}
}

func TestRunHookInvokesScriptCorrectly(t *testing.T) {
	expDir := "/usr/repos/foo"
	actDir := ""
	actCmd := ""
	origFunc := runCommandInDir
	expOutput := "hello world"
	expErr := errors.New("foobar")
	defer func() { runCommandInDir = origFunc }()
	runCommandInDir = func(directory, command string) (string, error) {
		actDir = directory
		actCmd = command
		return expOutput, expErr
	}

	output, err := runHook(runnerConfig, "pre-commit", expDir)

	if output != expOutput {
		t.Errorf("Did not get correct output. Expected: %s, but got: %s", expOutput, output)
	}

	if err != expErr {
		t.Errorf("Did not get correct error. Expected: %s, but got: %s", expErr, err)
	}

	if actDir != expDir {
		t.Errorf("Did not use correct root directory. Expected: %s, but got: %s", expDir, actDir)
	}

	if actCmd != preCommitScript {
		t.Errorf("Did not use correct hook script. Expected: %s, but got: %s", preCommitScript, actCmd)
	}
}
