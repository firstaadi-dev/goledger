# Architecture

## Layers

1. `Client` configuration and lifecycle (`client.go`)
2. Command execution (`raw.go`) using `os/exec` without shell interpolation
3. Typed wrappers for each supported hledger command (`commands.go`)
4. Shared request/response contracts (`types.go`)

## Request flow

1. Wrapper method receives `CommandRequest`
2. `CommonOpts` gets translated to CLI flags
3. Command is executed by `Client.exec`
4. Result returns raw bytes and text, and optional loose JSON decode when `--output-format json` is used

## Error flow

Execution failures are normalized to `CommandError`, containing command, exit code, stdout, and stderr.
