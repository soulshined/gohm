package test_utils

import (
	"fmt"
	"gohm/cli"
	"strings"
	"testing"
)

// CreateTestCommand creates a cli.Command for testing with the specified flags and args.
// flags: single-value flags as map[name]value
// multiFlags: multi-value flags as map[name][]values (can be nil)
// args: command arguments
func CreateTestCommand(handler func(*cli.Command) string, flags map[string]string, multiFlags map[string][]string, args []string) *cli.Command {
	cmd := &cli.Command{
		Handler:    handler,
		Args:       args,
		ArgsLength: len(args),
	}

	for name, value := range flags {
		flag := &cli.Flag{
			Name:    name,
			Value:   value,
			Default: value,
			IsSet:   value != "",
		}
		cmd.Flags = append(cmd.Flags, flag)
	}

	for name, values := range multiFlags {
		flag := &cli.Flag{
			Name:    name,
			IsMulti: true,
			Values:  values,
			IsSet:   len(values) > 0,
		}
		if len(values) > 0 {
			flag.Value = values[0]
		}
		cmd.Flags = append(cmd.Flags, flag)
	}

	return cmd
}

// ExpectPanic runs fn and verifies it panics with the expected message.
// Returns true if the panic occurred with the expected message.
func ExpectPanic(t *testing.T, expected string, fn func()) {
	t.Helper()
	defer func() {
		r := recover()
		if r == nil {
			t.Error("expected panic but none occurred")
			return
		}
		msg := fmt.Sprintf("%v", r)
		if msg != expected {
			t.Errorf("expected panic message %q, got %q", expected, msg)
		}
	}()
	fn()
}

// ExpectPanicContains runs fn and verifies it panics with a message containing substr.
func ExpectPanicContains(t *testing.T, substr string, fn func()) {
	t.Helper()
	defer func() {
		r := recover()
		if r == nil {
			t.Error("expected panic but none occurred")
			return
		}
		msg := fmt.Sprintf("%v", r)
		if !strings.Contains(msg, substr) {
			t.Errorf("expected panic message to contain %q, got %q", substr, msg)
		}
	}()
	fn()
}

// ExpectNoPanic runs fn and fails if it panics.
func ExpectNoPanic(t *testing.T, fn func()) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("unexpected panic: %v", r)
		}
	}()
	fn()
}

// AssertContains checks that result contains all expected substrings.
func AssertContains(t *testing.T, result string, substrings ...string) {
	t.Helper()
	for _, substr := range substrings {
		if !strings.Contains(result, substr) {
			t.Errorf("expected result to contain %q, got %q", substr, result)
		}
	}
}

// AssertEquals checks that got equals expected.
func AssertEquals[T comparable](t *testing.T, got, expected T) {
	t.Helper()
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
