package captaingithook

import ()

var configFileNames = [...]string{
	"captaingithook.json",
	".captaingithook.json",
	"captaingithookrc",
	".captaingithookrc",
	"captaingithookrc.json",
	".captaingithookrc.json",
	"captain-githook.json",
	".captain-githook.json",
	"captain-githookrc",
	".captain-githookrc",
	"captain-githookrc.json",
	".captain-githookrc.json",
}

// Config foo
type Config struct {
}

func isValidConfigFileName(fileName string) bool {
	if len(fileName) > 1 {
		for _, validConfigFileName := range configFileNames {
			if fileName == validConfigFileName {
				return true
			}
		}
	}

	return false
}

func createConfigFile(desiredFileName string) error {
	var configFileName string
	configFileName = configFileNames[0]

	if isValidConfigFileName(desiredFileName) {
		configFileName = desiredFileName
	}

	writeFile(configFileName)

	return nil
}

func getRepoConfig(repoRootDirectoryPath string) *Config {
	if len(repoRootDirectoryPath) < 1 {
		path, err := getGitRepoRootDirectoryPath()
		if err != nil {
			return nil
		}
		repoRootDirectoryPath = path
	}

	for _, configFileName := range configFileNames {
		configFilePath := repoRootDirectoryPath + "/" + configFileName
		readFile(configFilePath)
	}

	return nil
}
