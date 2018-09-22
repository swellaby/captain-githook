package captaingithook

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestWriteFileDoesNotPanicOnSuccess(t *testing.T) {
	write = func(filename string, data []byte, perm os.FileMode) error {
		return nil
	}
	defer func() { write = ioutil.WriteFile }()
}
