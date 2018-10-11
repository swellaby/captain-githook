package captaingithook

import (
	"errors"
	"fmt"
)

var runHookScript = runHook
var errNilHooksConfig = errors.New("hooks configuration was nil")
var errNilConfig = errors.New("configuration was nil")

func getConfiguredHookScript(hooksConfig *HooksConfig, hookName string) (hookScript string, err error) {
	if hooksConfig == nil {
		return "", errNilHooksConfig
	}

	switch hookName {
	case "pre-commit":
		hookScript = hooksConfig.PreCommit
	case "prepare-commit-msg":
		hookScript = hooksConfig.PrepareCommitMsg
	case "commit-msg":
		hookScript = hooksConfig.CommitMsg
	case "post-commit":
		hookScript = hooksConfig.PostCommit
	case "pre-push":
		hookScript = hooksConfig.PrePush
	case "pre-rebase":
		hookScript = hooksConfig.PreRebase
	case "post-checkout":
		hookScript = hooksConfig.PostCheckout
	case "post-merge":
		hookScript = hooksConfig.PostMerge
	case "applypatch-msg":
		hookScript = hooksConfig.ApplyPatchMsg
	case "pre-applypatch":
		hookScript = hooksConfig.PreApplyPatch
	case "post-applypatch":
		hookScript = hooksConfig.PostApplyPatch
	case "pre-receive":
		hookScript = hooksConfig.PreReceive
	case "update":
		hookScript = hooksConfig.Update
	case "post-receive":
		hookScript = hooksConfig.PostReceive
	case "post-update":
		hookScript = hooksConfig.PostUpdate
	case "push-to-checkout":
		hookScript = hooksConfig.PushToCheckout
	case "pre-auto-gc":
		hookScript = hooksConfig.PreAutoGc
	case "post-rewrite":
		hookScript = hooksConfig.PostRewrite
	case "sendemail-validate":
		hookScript = hooksConfig.SendEmailValidate
	default:
		return "", fmt.Errorf("unknown hook name: %s", hookName)
	}

	return hookScript, nil
}

func runHook(config *Config, hookName, rootDirectory string) (output string, err error) {
	if config == nil {
		return "", errNilConfig
	}

	hookScript, hookErr := getConfiguredHookScript(&config.Hooks, hookName)

	if hookErr != nil {
		return "", hookErr
	}

	return runCommandInDir(rootDirectory, hookScript)
}
