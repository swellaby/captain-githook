package captaingithook

import (
	"io/ioutil"
	"os"
)

var writeFile = write
var readFile = read
var ioWrite = ioutil.WriteFile
var ioRead = ioutil.ReadFile

func write(filePath, contents string) {
	err := ioWrite(filePath, []byte(contents), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func read(filePath string) []byte {
	data, err := ioRead(filePath)
	if err != nil {
		panic(err)
	}

	return data
}
