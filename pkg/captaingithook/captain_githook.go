package captaingithook

// Initialize sets up a repo
func Initialize(desiredConfigFileName string) error {
	configErr := initializeCaptainGithookConfigFile(desiredConfigFileName)

	if configErr != nil {
		return configErr
	}

	return nil
}
