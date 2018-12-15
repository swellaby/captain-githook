package cli

import (
	"github.com/stretchr/testify/assert"
	"github.com/swellaby/captain-githook/captaingithook"
	"testing"
)

func TestRootCommandConfiguresUseCorrectly(t *testing.T) {
	assert.Equal(t, "captain-githook", rootCmd.Use)
}

func TestRootCommandConfiguresVersionCorrectly(t *testing.T) {
	assert.Equal(t, captaingithook.Version, rootCmd.Version)
}

func TestGetRunnerReturnsRootCommand(t *testing.T) {
	assert.Equal(t, rootCmd, GetRunner())
}

func TestRootCommandHasCorrectSubcommands(t *testing.T) {
	actCommands := rootCmd.Commands()
	assert.Equal(t, 2, len(actCommands))
	assert.Equal(t, initCmd, actCommands[0])
	assert.Equal(t, runCmd, actCommands[1])
}
