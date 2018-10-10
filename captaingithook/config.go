package captaingithook

import (
	"encoding/json"
	"errors"
	"path/filepath"
)

var jsonMarshallIndent = json.MarshalIndent
var jsonUnmarshall = json.Unmarshal
var initializeCaptainGithookConfigFile = initConfigFile
var errConfigFileNotFound = errors.New("did not find a captain-githook config file, unable to run hook")
var errConfigFileParseFailed = errors.New("encountered a fatal error while attempting to parse the captain-githook config file")

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

// Config represents a captain-githook configuration
type Config struct {
	Hooks HooksConfig `json:"hooks,omitempty"`
}

// HooksConfig represents the git hooks configuration
type HooksConfig struct {
	PreCommit string `json:"pre-commit,omitempty"`
	PrePush   string `json:"pre-push,omitempty"`
	CommitMsg string `json:"commit-msg,omitempty"`
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

func getDefaultConfigJSONContent() ([]byte, error) {
	hooksConfig := &HooksConfig{
		PreCommit: "go test ./...",
	}
	config := &Config{Hooks: *hooksConfig}
	return jsonMarshallIndent(config, "", "  ")
}

func configFileExists(path string) bool {
	for _, configFileName := range configFileNames {
		configFilePath := filepath.Join(path, configFileName)
		if fileExists(configFilePath) {
			return true
		}
	}
	return false
}

func getConfigFileName(desiredFileName string) string {
	configFileName := ""

	if isValidConfigFileName(desiredFileName) {
		configFileName = desiredFileName
	} else {
		configFileName = configFileNames[0]
	}

	return configFileName
}

func initConfigFile(repoPath, desiredFileName string) error {
	if !configFileExists(repoPath) {
		configFileName := getConfigFileName(desiredFileName)
		data, _ := getDefaultConfigJSONContent()
		configFilePath := filepath.Join(repoPath, configFileName)
		err := writeFile(configFilePath, data)

		if err != nil {
			baseErr := err.Error()
			errMsg := "unexpected error encountered while trying to create the config file. Error details: " + baseErr
			return errors.New(errMsg)
		}
	}

	return nil
}

func getRepoConfig(repoRootDirectoryPath string) (config *Config, err error) {
	var configFileContents []byte
	foundConfigFile := false

	for _, configFileName := range configFileNames {
		configFilePath := filepath.Join(repoRootDirectoryPath, configFileName)

		if fileContents, err := readFile(configFilePath); err == nil {
			configFileContents = fileContents
			foundConfigFile = true
			break
		}
	}

	if !foundConfigFile {
		return nil, errConfigFileNotFound
	}

	parseErr := jsonUnmarshall(configFileContents, &config)

	if parseErr != nil {
		return nil, errConfigFileParseFailed
	}

	return config, nil
}
