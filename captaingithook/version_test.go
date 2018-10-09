package captaingithook

import "testing"

func TestVersionIsSet(t *testing.T) {
	if v := Version; v == "" {
		t.Errorf("Version value was not set.")
	}
}
