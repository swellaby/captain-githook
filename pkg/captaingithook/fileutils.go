package captaingithook

import (
	"io/ioutil"
	"os"
)

var writeFile = write
var readFile = read
var ioWrite = ioutil.WriteFile
var ioRead = ioutil.ReadFile
var fileExists = exists
var osStat = os.Stat
var osIsNotExist = os.IsNotExist

func write(filePath string, contents []byte) error {
	return ioWrite(filePath, contents, os.ModePerm)
}

func read(filePath string) ([]byte, error) {
	return ioRead(filePath)
}

func exists(filePath string) bool {
	if _, err := osStat(filePath); osIsNotExist(err) {
		return false
	}

	return true
}
