package goledger

import "time"

// CommonOpts contains frequently used hledger flags.
type CommonOpts struct {
	File         string
	Begin        string
	End          string
	Period       string
	Depth        int
	OutputFormat string
	Value        string
	RulesFile    string
	Alias        []string
	Args         []string
}

func (o CommonOpts) argv(defaultLedgerFile string) []string {
	args := make([]string, 0, 16)
	file := o.File
	if file == "" {
		file = defaultLedgerFile
	}
	if file != "" {
		args = append(args, "--file", file)
	}
	if o.Begin != "" {
		args = append(args, "--begin", o.Begin)
	}
	if o.End != "" {
		args = append(args, "--end", o.End)
	}
	if o.Period != "" {
		args = append(args, "--period", o.Period)
	}
	if o.Depth > 0 {
		args = append(args, "--depth", itoa(o.Depth))
	}
	if o.OutputFormat != "" {
		args = append(args, "--output-format", o.OutputFormat)
	}
	if o.Value != "" {
		args = append(args, "--value", o.Value)
	}
	if o.RulesFile != "" {
		args = append(args, "--rules-file", o.RulesFile)
	}
	for _, a := range o.Alias {
		if a != "" {
			args = append(args, "--alias", a)
		}
	}
	args = append(args, o.Args...)
	return args
}

// CommandRequest is the generic typed request used by all command wrappers.
type CommandRequest struct {
	Common  CommonOpts
	Query   string
	Args    []string
	Stdin   []byte
	Timeout time.Duration
}

// CommandResult is the standard response contract for wrapper commands.
type CommandResult struct {
	Structured any
	Text       string
	Raw        RawResult
}

// VersionInfo contains hledger version details.
type VersionInfo struct {
	Raw string
}

func itoa(v int) string {
	if v == 0 {
		return "0"
	}
	buf := [20]byte{}
	i := len(buf)
	for v > 0 {
		i--
		buf[i] = byte('0' + v%10)
		v /= 10
	}
	return string(buf[i:])
}
