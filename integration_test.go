package goledger

import (
	"context"
	"errors"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestIntegrationAccounts(t *testing.T) {
	if _, err := exec.LookPath("hledger"); err != nil {
		t.Skip("hledger not installed in PATH")
	}

	c, err := New(WithDefaultLedgerFile(filepath.Join("testdata", "basic.journal")))
	if err != nil {
		t.Fatalf("new client: %v", err)
	}

	res, err := c.Accounts(context.Background(), CommandRequest{})
	if err != nil {
		var ce *CommandError
		if errors.As(err, &ce) {
			t.Fatalf("hledger failed: stderr=%s", ce.Stderr)
		}
		t.Fatalf("accounts: %v", err)
	}

	if !strings.Contains(res.Text, "assets:cash") {
		t.Fatalf("expected account in output, got: %s", res.Text)
	}
}
