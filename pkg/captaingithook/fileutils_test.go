package captaingithook

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestWriteFileDoesNotPanicOnSuccess(t *testing.T) {
	var actualFileName string
	var actualData []byte
	var actualPerm os.FileMode
	expectedPerm := os.ModePerm
	expectedFileName := "/usr/foo/bar.txt"
	data := "foobaroo"
	expectedData := []byte(data)

	ioWrite = func(filename string, data []byte, perm os.FileMode) error {
		actualFileName = filename
		actualData = data
		actualPerm = perm
		return nil
	}
	defer func() { ioWrite = ioutil.WriteFile }()
	write(expectedFileName, data)

	if actualFileName != expectedFileName {
		t.Errorf("Attempted to write to wrong file. Expected: %s, but got: %s.", expectedFileName, actualFileName)
	}

	if bytes.Compare(actualData, expectedData) != 0 {
		t.Errorf("Attempted to write incorrect data file. Expected: %s, but got: %s.", expectedData, actualData)
	}

	if actualPerm != expectedPerm {
		t.Errorf("Attempted to use wrong perms on file. Expected: %d, but got: %d.", expectedPerm, actualPerm)
	}

}
