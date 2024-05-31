package tass64

import (
	"bytes"
	"errors"
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
		return stderr.String(), err
	}
	return out.String(), nil
}

func Does64TassExist() bool {
	cmd := exec.Command(TASS_EXECUTABLE)
	err := cmd.Run()
	if err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			return false
		}
	}
	return true
}
