package captaingithook

import (
	"errors"
	"strings"
)

var getGitRepoRootDirectoryPath = getRootDirectoryPath
var getGitRepoHooksDirectoryPath = getHooksDirectoryPath

func getRootDirectoryPath() (gitDirPath string, err error) {
	result, err := runCommand("git", "rev-parse", "--show-toplevel")
	if err != nil {
		baseErr := err.Error()
		errMsg := "unexpected error encountered while trying to determine the git repo root directory. Error details: " + baseErr
		return result, errors.New(errMsg)
	} else if len(result) == 0 {
		errMsg := "got an unexpected result for the git repo root directory."
		return result, errors.New(errMsg)
	}
	gitDirPath = strings.TrimSuffix(result, "\n")
	return gitDirPath, err
}

func getHooksDirectoryPath() (string, error) {
	return "", nil
}
