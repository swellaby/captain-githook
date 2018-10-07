package captaingithook

import (
	"errors"
	"fmt"
	"path/filepath"
)

var initializeGitHookFiles = createAllHookFiles

var gitHooks = [...]string{
	"applypatch-msg",
	"pre-applypatch",
	"post-applypatch",
	"pre-commit",
	"prepare-commit-msg",
	"commit-msg",
	"post-commit",
	"pre-rebase",
	"post-checkout",
	"post-merge",
	"pre-push",
	"pre-receive",
	"update",
	"post-receive",
	"post-update",
	"push-to-checkout",
	"pre-auto-gc",
	"post-rewrite",
	"sendemail-validate",
}

const hookFileScript = `#!/bin/sh
# captain-githook
# version v0.0.1

hookName=` + "`basename \"$0\"`" + `
gitParams="$*"

if command -v captain-githook >/dev/null 2>&1; then
  captain-githook $hookName "$gitParams"
else
  echo "Can't find captain-githook, skipping $hookName hook"
  echo "You can reinstall it using 'go get -u github.com/swellaby/captain-githook' or delete this hook"
fi`

var hookFileContents = []byte(hookFileScript)
var errInvalidGitHooksDirectoryPath = errors.New("invalid git hooks directory path")

func createAllHookFiles() error {
	hooksDir, hooksErr := getGitRepoHooksDirectory()
	if hooksErr != nil {
		return errInvalidGitHooksDirectoryPath
	}

	var notCreatedHooks []string

	for _, hook := range gitHooks {
		hookPath := filepath.Join(hooksDir, hook)
		err := writeFile(hookPath, hookFileContents)
		if err != nil {
			notCreatedHooks = append(notCreatedHooks, hook)
		}
	}

	if len(notCreatedHooks) > 0 {
		return fmt.Errorf("encountered an error while attempting to create one or more hook files. did not create hooks: %v", notCreatedHooks)
	}

	return nil
}

func removeAllHookFiles() error {
	return nil
}
