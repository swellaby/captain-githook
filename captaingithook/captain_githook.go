package captaingithook

// Initialize sets up a repo
func Initialize() error {
	return InitializeWithFileName("")
}

// InitializeWithFileName sets up a repo using the specified config file name
func InitializeWithFileName(desiredConfigFileName string) error {
	path, err := getGitRepoRootDirectoryPath()
	if err != nil {
		return err
	}
	configErr := initializeCaptainGithookConfigFile(path, desiredConfigFileName)

	if configErr != nil {
		return configErr
	}

	hooksErr := initializeGitHookFiles()
	if hooksErr != nil {
		return hooksErr
	}

	return nil
}

// RunHook runs the specified hook
func RunHook(hookName string) (output string, err error) {
	dirPath, err := getGitRepoRootDirectoryPath()
	if err != nil {
		return "", err
	}

	config, configErr := getCaptainGithookConfig(dirPath)

	if configErr != nil {
		return "", configErr
	}

	return runHookScript(config, hookName, dirPath)
}
