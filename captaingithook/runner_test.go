package captaingithook

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, "", hookScript)
	assert.Equal(t, errNilHooksConfig, err)
}

func TestGetConfiguredHookScriptReturnsCorrectErrorOnUnknownHook(t *testing.T) {
	hookName := "foo"
	hookScript, err := getConfiguredHookScript(runnerHooks, hookName)
	assert.Equal(t, "", hookScript)
	assert.Error(t, fmt.Errorf("unknown hook name: %s", hookName), err)
}

func TestGetConfiguredHookReturnsCorrectScriptForHook(t *testing.T) {
	for _, hst := range hookScriptTests {
		expHookScript := hst.expHookScript
		hook := hst.hookName
		actHookScript, err := getConfiguredHookScript(runnerHooks, hook)
		assert.Nil(t, err)
		assert.Equal(t, expHookScript, actHookScript, "Did not get correct script for hook: %s. Expected: %s, but got: %s", hook, expHookScript, actHookScript)
	}
}

func TestRunHookReturnsCorrectErrorWhenConfigIsNil(t *testing.T) {
	output, err := runHook(nil, "", "")
	assert.Equal(t, errNilConfig, err)
	assert.Equal(t, "", output)
}

func TestRunHookReturnsCorrectErrorWhenScriptCannotBeDetermined(t *testing.T) {
	output, err := runHook(&Config{}, "", "")
	assert.Equal(t, "", output)
	assert.Error(t, fmt.Errorf("unknown hook name: %s", ""), err)
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
	assert.Equal(t, expOutput, output)
	assert.Equal(t, expErr, err)
	assert.Equal(t, expDir, actDir)
	assert.Equal(t, preCommitScript, actCmd, "Did not use correct hook script. Expected: %s, but got: %s", preCommitScript, actCmd)
}
