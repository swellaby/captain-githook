package captaingithook

import (
	"testing"
)

var expGitHooks = [...]string{
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

func TestWriteAllHookFiles(t *testing.T) {
	// var actHookPaths []string
	originalWriteFile := writeFile
	defer func() { writeFile = originalWriteFile }()
	writeFile = func(filePath string, contents []byte) error {
		// actHookPaths.
		return nil
	}
	writeAllHookFiles("")
}
