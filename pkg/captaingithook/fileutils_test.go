package captaingithook

import (
	"bytes"
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
	data := "foobaroo"
	expectedData := []byte(data)

	ioWrite = func(filename string, data []byte, perm os.FileMode) error {
		actualFileName = filename
		actualData = data
		actualPerm = perm
		return nil
	}
	defer func() { ioWrite = ioutil.WriteFile }()
	actualErr := write(expectedFileName, data)

	if actualErr != nil {
		t.Errorf("Got unexpected error %s.", actualErr)
	}

	if actualFileName != expectedFileName {
		t.Errorf("Attempted to write to wrong file. Expected: %s, but got: %s.", expectedFileName, actualFileName)
	}

	if !bytes.Equal(actualData, expectedData) {
		t.Errorf("Attempted to write incorrect data file. Expected: %s, but got: %s.", expectedData, actualData)
	}

	if actualPerm != expectedPerm {
		t.Errorf("Attempted to use wrong perms on file. Expected: %d, but got: %d.", expectedPerm, actualPerm)
	}

}
