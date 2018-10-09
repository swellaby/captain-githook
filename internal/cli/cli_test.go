package cli

import (
	"github.com/swellaby/captain-githook/captaingithook"
	"testing"
)

func TestRootCommandConfiguresUseCorrectly(t *testing.T) {
	expUse := "captain-githook"
	use := rootCmd.Use
	if use != expUse {
		t.Errorf("Did not set correct use value for root command. Expected: %s, but got: %s", expUse, use)
	}
}

func TestRootCommandConfiguresVersionCorrectly(t *testing.T) {
	expVersion := captaingithook.Version
	version := rootCmd.Version
	if version != expVersion {
		t.Errorf("Did not set correct version. Expected: %s, but got: %s", expVersion, version)
	}
}

func TestGetRunnerReturnsRootCommand(t *testing.T) {
	actRunner := GetRunner()
	if actRunner != rootCmd {
		t.Errorf("Did not get correct runner. Expected: %v, but got: %v", rootCmd, actRunner)
	}
}

func TestRootCommandHasCorrectSubcommands(t *testing.T) {
	expNumCommands := 1
	actCommands := rootCmd.Commands()
	actNumCommands := len(actCommands)

	if actNumCommands != expNumCommands {
		t.Errorf("Did not get correct number of subcommands. Expected: %d, but got: %d", expNumCommands, actNumCommands)
	}

	if cmd := actCommands[0]; cmd != initCmd {
		t.Errorf("Did not get correct init subcommand. Expected: %s, but got: %s", initCmd.Use, cmd.Use)
	}
}
