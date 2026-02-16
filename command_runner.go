package goledger

import "context"

func (c *Client) runCommand(ctx context.Context, command string, req CommandRequest) (CommandResult, error) {
	args := make([]string, 0, 20)
	args = append(args, command)
	args = append(args, req.Common.argv(c.defaultLedgerFile)...)
	if req.Query != "" {
		args = append(args, req.Query)
	}
	args = append(args, req.Args...)

	raw, err := c.exec(ctx, args, req.Stdin, req.Timeout)
	if err != nil {
		return CommandResult{}, err
	}

	result := CommandResult{
		Text: string(raw.Stdout),
		Raw:  raw,
	}
	if req.Common.OutputFormat == "json" {
		result.Structured = decodeJSONLoose(raw.Stdout)
	}
	return result, nil
}
