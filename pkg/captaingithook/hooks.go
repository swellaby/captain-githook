package captaingithook

import (
	"path/filepath"
)

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

func writeAllHookFiles(gitHooksDirectoryPath string) error {
	firstHook := gitHooks[0]
	hookPath := filepath.Join(gitHooksDirectoryPath, firstHook)
	writeFile(hookPath, hookFileContents)
	return nil
}

func removeAllHookFiles() {

}
