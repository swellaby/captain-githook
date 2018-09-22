package captaingithook

import (
	"io/ioutil"
	"os"
)

var write = ioutil.WriteFile
var read = ioutil.ReadFile

var writeFile = func(filePath, contents string) {
	err := write(filePath, []byte(contents), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

var readFile = func(filePath string) []byte {
	data, err := read(filePath)
	if err != nil {
		panic(err)
	}

	return data
}
