package git

import (
	"fmt"
	"os/exec"
)

// LogFormat defines the format for git log output, using null characters as separators for easy parsing
const LogFormat = "%H%x00%an%x00%ae%x00%aI%x00%s"

// Log executes the git log command with the specified format and returns the raw output as a string
func Log(repoPath string) (string, error) {
	cmd := exec.Command("git", "-C", repoPath, "log", "--pretty=format:"+LogFormat, "--numstat")
	output, err := cmd.Output()
	if err != nil { // If the command fails, return an error with context about the failure
		return "", fmt.Errorf("failed to run git log: %w", err)
	}

	// Return the raw output as a string for further processing by the parser
	return string(output), nil
}
