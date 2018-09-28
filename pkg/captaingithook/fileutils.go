package captaingithook

import (
	"io/ioutil"
	"os"
)

var writeFile = write
var readFile = read
var ioWrite = ioutil.WriteFile
var ioRead = ioutil.ReadFile

func write(filePath, contents string) error {
	return ioWrite(filePath, []byte(contents), os.ModePerm)
}

func read(filePath string) ([]byte, error) {
	return ioRead(filePath)
}
