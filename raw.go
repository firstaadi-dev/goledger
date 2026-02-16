package goledger

import (
	"bytes"
	"context"
	"os/exec"
	"strings"
	"time"
)

// RawRequest executes arbitrary hledger command arguments.
type RawRequest struct {
	Command string
	Args    []string
	Stdin   []byte
}

// RawResult captures process output and exit code.
type RawResult struct {
	Stdout   []byte
	Stderr   []byte
	ExitCode int
}

// Raw executes hledger with a raw command + args.
func (c *Client) Raw(ctx context.Context, req RawRequest) (RawResult, error) {
	args := make([]string, 0, len(req.Args)+1)
	if req.Command != "" {
		args = append(args, req.Command)
	}
	args = append(args, req.Args...)
	return c.exec(ctx, args, req.Stdin, 0)
}

func (c *Client) exec(ctx context.Context, args []string, stdin []byte, timeout time.Duration) (RawResult, error) {
	ctx, cancel, err := c.buildExecConfig(ctx, timeout)
	if err != nil {
		return RawResult{}, err
	}
	defer cancel()

	cmd := exec.CommandContext(ctx, c.binaryPath, args...)
	if c.workingDir != "" {
		cmd.Dir = c.workingDir
	}
	cmd.Env = c.envForCommand()
	if len(stdin) > 0 {
		cmd.Stdin = bytes.NewReader(stdin)
	}

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	c.logf("running: %s %s", c.binaryPath, strings.Join(args, " "))
	err = cmd.Run()
	res := RawResult{Stdout: stdout.Bytes(), Stderr: stderr.Bytes(), ExitCode: 0}
	if err == nil {
		return res, nil
	}

	if ee, ok := err.(interface{ ExitCode() int }); ok {
		res.ExitCode = ee.ExitCode()
	} else {
		res.ExitCode = -1
	}
	return res, toCommandError(append([]string{c.binaryPath}, args...), res.Stdout, res.Stderr, err)
}

// Version returns hledger version line from `hledger --version`.
func (c *Client) Version(ctx context.Context) (VersionInfo, error) {
	res, err := c.exec(ctx, []string{"--version"}, nil, 0)
	if err != nil {
		return VersionInfo{}, err
	}
	return VersionInfo{Raw: parseVersionOutput(string(res.Stdout))}, nil
}
