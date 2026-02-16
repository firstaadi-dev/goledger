package goledger

import (
	"errors"
	"fmt"
)

// CommandError captures command exit failure and stderr/stdout context.
type CommandError struct {
	Command  []string
	ExitCode int
	Stdout   string
	Stderr   string
	Cause    error
}

func (e *CommandError) Error() string {
	if e == nil {
		return "<nil>"
	}
	return fmt.Sprintf("hledger command failed (exit=%d): %v", e.ExitCode, e.Command)
}

func (e *CommandError) Unwrap() error { return e.Cause }

func toCommandError(cmd []string, stdout, stderr []byte, err error) error {
	if err == nil {
		return nil
	}

	ce := &CommandError{
		Command:  cmd,
		Stdout:   string(stdout),
		Stderr:   string(stderr),
		Cause:    err,
		ExitCode: -1,
	}

	var ee interface{ ExitCode() int }
	if errors.As(err, &ee) {
		ce.ExitCode = ee.ExitCode()
	}

	return ce
}
