# goledger

`goledger` is an MIT-licensed Go wrapper around [`hledger`](https://hledger.org/doc.html).

## Requirements

- Go `1.21.13`
- `hledger` binary installed and available in `PATH` (or configure via `WithBinaryPath`)

## Install

```bash
go get github.com/firstaadi/goledger
```

## Quick start

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/firstaadi/goledger"
)

func main() {
	c, err := goledger.New(goledger.WithDefaultLedgerFile("./journal.ledger"))
	if err != nil {
		log.Fatal(err)
	}

	res, err := c.Balance(context.Background(), goledger.CommandRequest{
		Common: goledger.CommonOpts{OutputFormat: "json"},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.Text)
}
```

## Supported commands

`accounts`, `activity`, `add`, `aregister`, `balance`, `balancesheet`, `bs`, `cashflow`, `cf`, `check`, `close`, `codes`, `commodities`, `descriptions`, `diff`, `files`, `help`, `import`, `incomestatement`, `is`, `notes`, `payees`, `prices`, `print`, `register`, `rewrite`, `roi`, `stats`, `tags`, `test`.

All commands are also reachable through `Client.Raw(...)`.
