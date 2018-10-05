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

	return nil
}
