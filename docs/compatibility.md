# Compatibility

## Supported versions

- Library: `github.com/firstaadi/goledger` (`v0.x` line)
- Go: `1.21.13`
- hledger: `1.51.x`

## Runtime dependency

`hledger` binary must be preinstalled and accessible from `PATH`, or provided using `WithBinaryPath`.

## Platforms

- Linux
- macOS
- Windows

## Notes

This library wraps CLI behavior. Minor output differences can happen across hledger patch versions or locale-specific environments.
