package goledger

import (
	"context"
	"errors"
	"testing"
)

type fakeExitErr struct{ code int }

func (e fakeExitErr) Error() string { return "exit" }
func (e fakeExitErr) ExitCode() int { return e.code }

func TestCommonOptsArgv(t *testing.T) {
	o := CommonOpts{
		File:         "a.ledger",
		Begin:        "2026-01-01",
		End:          "2026-12-31",
		Period:       "monthly",
		Depth:        2,
		OutputFormat: "json",
		Value:        "then,USD",
		RulesFile:    "rules.csv",
		Alias:        []string{"expenses=ex"},
		Args:         []string{"--no-total"},
	}

	got := o.argv("")
	if len(got) == 0 {
		t.Fatal("expected args")
	}
}

func TestDecodeJSONLoose(t *testing.T) {
	v := decodeJSONLoose([]byte(`{"ok":true}`))
	m, ok := v.(map[string]any)
	if !ok {
		t.Fatalf("expected map, got %T", v)
	}
	if m["ok"] != true {
		t.Fatalf("unexpected value: %#v", m)
	}
}

func TestToCommandError(t *testing.T) {
	err := toCommandError([]string{"hledger", "balance"}, []byte("a"), []byte("b"), fakeExitErr{code: 2})
	var ce *CommandError
	if !errors.As(err, &ce) {
		t.Fatal("expected CommandError")
	}
	if ce.ExitCode != 2 {
		t.Fatalf("expected exit code 2, got %d", ce.ExitCode)
	}
}

func TestClientNewFailsWithoutBinary(t *testing.T) {
	c, err := New(WithBinaryPath("/definitely/not-found/hledger"))
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	_, err = c.Raw(context.Background(), RawRequest{Command: "--version"})
	if err == nil {
		t.Fatal("expected execution error")
	}
}
