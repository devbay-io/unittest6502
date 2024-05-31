package tass64

import (
	"bytes"
	"fmt"
	"os/exec"
)

var (
	TASS_EXECUTABLE = "64tass"
)

// Run64tass executes the 64tass command with the provided arguments.
func Run64tass(args ...string) (string, error) {
	cmd := exec.Command(TASS_EXECUTABLE, args...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return stderr.String(), fmt.Errorf("failed to run 64tass with args %v: %w", args, err)
	}
	return out.String(), nil
}

func Does64TassExist() bool {
	_, err := exec.LookPath(TASS_EXECUTABLE)
	return err == nil
}
