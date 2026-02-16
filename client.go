package goledger

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Logger is a minimal log sink for command execution events.
type Logger interface {
	Printf(format string, args ...any)
}

// Client executes hledger commands.
type Client struct {
	binaryPath        string
	env               map[string]string
	workingDir        string
	defaultLedgerFile string
	timeout           time.Duration
	logger            Logger
}

// Option mutates client settings.
type Option func(*Client)

// WithBinaryPath sets explicit hledger binary path.
func WithBinaryPath(path string) Option {
	return func(c *Client) { c.binaryPath = path }
}

// WithEnv adds or overrides command env vars.
func WithEnv(env map[string]string) Option {
	return func(c *Client) {
		if c.env == nil {
			c.env = make(map[string]string)
		}
		for k, v := range env {
			c.env[k] = v
		}
	}
}

// WithWorkingDir sets command working directory.
func WithWorkingDir(dir string) Option {
	return func(c *Client) { c.workingDir = dir }
}

// WithDefaultLedgerFile sets the default journal file if request does not provide one.
func WithDefaultLedgerFile(path string) Option {
	return func(c *Client) { c.defaultLedgerFile = path }
}

// WithTimeout sets default timeout for command execution.
func WithTimeout(d time.Duration) Option {
	return func(c *Client) { c.timeout = d }
}

// WithLogger sets optional logger.
func WithLogger(logger Logger) Option {
	return func(c *Client) { c.logger = logger }
}

// New creates a new hledger client.
func New(opts ...Option) (*Client, error) {
	c := &Client{}
	for _, opt := range opts {
		opt(c)
	}

	if c.binaryPath == "" {
		path, err := exec.LookPath("hledger")
		if err != nil {
			return nil, errors.New("hledger binary not found in PATH; install hledger or provide WithBinaryPath")
		}
		c.binaryPath = path
	}

	return c, nil
}

func (c *Client) buildExecConfig(ctx context.Context, timeout time.Duration) (context.Context, context.CancelFunc, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if timeout <= 0 {
		timeout = c.timeout
	}
	if timeout > 0 {
		cctx, cancel := context.WithTimeout(ctx, timeout)
		return cctx, cancel, nil
	}
	return ctx, func() {}, nil
}

func (c *Client) logf(format string, args ...any) {
	if c.logger != nil {
		c.logger.Printf(format, args...)
	}
}

func (c *Client) envForCommand() []string {
	if len(c.env) == 0 {
		return os.Environ()
	}
	base := os.Environ()
	for k, v := range c.env {
		base = append(base, fmt.Sprintf("%s=%s", k, v))
	}
	return base
}

func parseVersionOutput(out string) string {
	line := strings.TrimSpace(out)
	if i := strings.IndexByte(line, '\n'); i >= 0 {
		line = line[:i]
	}
	return strings.TrimSpace(line)
}
