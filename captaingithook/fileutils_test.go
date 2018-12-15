package captaingithook

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestWriteFileUsesCorrectValues(t *testing.T) {
	var actualFileName string
	var actualData []byte
	var actualPerm os.FileMode
	expectedPerm := os.ModePerm
	expectedFileName := "/usr/foo/bar.txt"
	data := []byte("fooBarRoo")
	expectedData := []byte(data)

	ioWrite = func(filename string, data []byte, perm os.FileMode) error {
		actualFileName = filename
		actualData = data
		actualPerm = perm
		return nil
	}
	defer func() { ioWrite = ioutil.WriteFile }()
	actualErr := write(expectedFileName, data)
	assert.Nil(t, actualErr)
	assert.Equal(t, expectedFileName, actualFileName)
	assert.Equal(t, expectedData, actualData, "Attempted to write incorrect data file. Expected: %s, but got: %s.", expectedData, actualData)
	assert.Equal(t, expectedPerm, actualPerm, "Attempted to use wrong perms on file. Expected: %d, but got: %d.", expectedPerm, actualPerm)
}

func TestReadFileUsesCorrectValues(t *testing.T) {
	var actualFilePath string
	expectedFilePath := "/usr/foo/bar.txt"
	expectedData := []byte("fooBarRoo")

	ioRead = func(filepath string) ([]byte, error) {
		actualFilePath = filepath
		return expectedData, nil
	}
	defer func() { ioRead = ioutil.ReadFile }()
	data, err := read(expectedFilePath)
	assert.Nil(t, err)
	assert.Equal(t, expectedFilePath, actualFilePath, "Attempted to write to wrong file. Expected: %s, but got: %s.", expectedFilePath, actualFilePath)
	assert.Equal(t, expectedData, data, "Did not get correct file data. Expected: %s, but got: %s.", expectedData, data)
}

func TestFileExistsReturnsFalseWhenErrorIsOsNotExist(t *testing.T) {
	osStat = func(file string) (os.FileInfo, error) { return nil, nil }
	defer func() { osStat = os.Stat }()
	osIsNotExist = func(err error) bool { return true }
	defer func() { osIsNotExist = os.IsNotExist }()
	assert.False(t, exists("/usr/foo/repos/my-repo/captaingithook.json"))
}

func TestFileExistsReturnsTrueWhenErrorIsNotOsNotExist(t *testing.T) {
	osStat = func(file string) (os.FileInfo, error) { return nil, nil }
	defer func() { osStat = os.Stat }()
	osIsNotExist = func(err error) bool { return false }
	defer func() { osIsNotExist = os.IsNotExist }()
	assert.True(t, exists("/usr/foo/repos/my-repo/captaingithook.json"))
}
