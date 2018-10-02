package captaingithook

import (
	"errors"
	"path/filepath"
)

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

var errConfigFileSearch = errors.New("encountered a fatal error while checking for existing config files")


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

func configFileExists() bool {
	path, err := getGitRepoRootDirectoryPath()
	if err != nil {
		panic(err)
	}

	for _, configFileName := range configFileNames {
		configFilePath := filepath.Join(path, configFileName)
		if fileExists(configFilePath) {
			return true
		}
	}
	return false
}

func createConfigFile(desiredFileName string) (err error) {
	configFileName := ""

	if isValidConfigFileName(desiredFileName) {
		configFileName = desiredFileName
	} else {
		configFileName = configFileNames[0]
	}

	defer func() {
		if r := recover(); r != nil {
			err = errConfigFileSearch
		}
	}()

	foundFile := configFileExists()

	if !foundFile {
		err = writeFile(configFileName, "")
	}

	return
}

func getRepoConfig() (config *Config, err error) {
	path, err := getGitRepoRootDirectoryPath()
	if err != nil {
		return nil, err
	}

	for _, configFileName := range configFileNames {
		configFilePath := filepath.Join(path, configFileName)
		readFile(configFilePath)
	}

	return nil, nil
}
