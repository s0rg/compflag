package compflag

import (
	"bytes"
	"flag"
	"strings"
	"testing"
)

func TestCompleteErrors(t *testing.T) {
	t.Setenv(envCompLine, "")
	t.Setenv(envCompPoint, "")

	if Complete() {
		t.Error("step 1: fail")
	}

	t.Setenv(envCompLine, "foo ")

	if Complete() {
		t.Error("step 2: fail")
	}

	t.Setenv(envCompPoint, "bar")

	if Complete() {
		t.Error("step 3: fail")
	}

	t.Setenv(envCompLine, "app ")
	t.Setenv(envCompPoint, "5")

	if !Complete() {
		t.Error("step 5: fail")
	}
}

func TestComplete(t *testing.T) {
	var (
		buf  bytes.Buffer
		sbuf string
		fs   = flag.NewFlagSet("test", flag.ExitOnError)
		opts = []Option{
			WithFlagSet(fs),
			WithWriter(&buf),
		}
	)

	const (
		nfoo = "foo"
		nbar = "bar"
	)

	_ = fs.Bool(nfoo, false, "")
	_ = fs.String(nbar, "", "")

	t.Setenv(envCompLine, "app ")
	t.Setenv(envCompPoint, "5")

	if !Complete(opts...) {
		t.Error("step 1: no complete")
	}

	sbuf = buf.String()

	if strings.Count(sbuf, "\n") != 2 {
		t.Error("step 1: unexpected lines")
	}

	if !strings.Contains(sbuf, nfoo) {
		t.Error("step 1: no foo")
	}

	if !strings.Contains(sbuf, nbar) {
		t.Error("step 1: no bar")
	}

	buf.Reset()

	t.Setenv(envCompLine, "app f")
	t.Setenv(envCompPoint, "6")

	if !Complete(opts...) {
		t.Error("step 2: no complete")
	}

	sbuf = buf.String()

	if strings.Count(sbuf, "\n") != 1 {
		t.Error("step 2: unexpected lines")
	}

	if !strings.Contains(sbuf, nfoo) {
		t.Error("step 2: no foo")
	}

	if strings.Contains(sbuf, nbar) {
		t.Error("step 2: has bar")
	}

	buf.Reset()

	t.Setenv(envCompLine, "app b")
	t.Setenv(envCompPoint, "6")

	if !Complete(opts...) {
		t.Error("step 3: no complete")
	}

	sbuf = buf.String()

	if strings.Count(sbuf, "\n") != 1 {
		t.Error("step 3: unexpected lines")
	}

	if strings.Contains(sbuf, nfoo) {
		t.Error("step 3: has foo")
	}

	if !strings.Contains(sbuf, nbar) {
		t.Error("step 3: no bar")
	}
}

func TestHidden(t *testing.T) {
	const (
		nfoo = "foo"
		nbar = "bar"
		nfud = "fud"
	)

	var (
		fs   = flag.NewFlagSet("test", flag.ExitOnError)
		buf  bytes.Buffer
		sbuf string
		opts = []Option{
			WithFlagSet(fs),
			WithWriter(&buf),
			WithHidden(nfud),
		}
	)

	_ = fs.Bool(nfoo, false, "")
	_ = fs.String(nbar, "", "")
	_ = fs.Int(nfud, 0, "")

	t.Setenv(envCompLine, "app f")
	t.Setenv(envCompPoint, "6")

	if !Complete(opts...) {
		t.Error("step 1: no complete")
	}

	sbuf = buf.String()

	if strings.Count(sbuf, "\n") != 1 {
		t.Error("step 1: unexpected lines")
	}

	if !strings.Contains(sbuf, nfoo) {
		t.Error("step 1: no foo")
	}

	if strings.Contains(sbuf, nbar) {
		t.Error("step 1: has bar")
	}

	if strings.Contains(sbuf, nfud) {
		t.Error("step 1: has fud")
	}
}

func TestCompleteVar(t *testing.T) {
	var (
		fs   = flag.NewFlagSet("test", flag.ContinueOnError)
		buf  bytes.Buffer
		sbuf string
	)

	const (
		ncoo = "coo"
		nbar = "bar"
		ncud = "cud"
		narg = "complete"
	)

	_ = fs.Bool(ncoo, false, "")
	_ = fs.String(nbar, "", "")
	_ = fs.Int(ncud, 0, "")

	t.Setenv(envCompLine, "app -c")
	t.Setenv(envCompPoint, "6")

	Var(narg,
		WithFlagSet(fs),
		WithWriter(&buf),
		WithHidden(ncud),
		WithExitFunc(func(int) {}),
	)

	if err := fs.Parse([]string{"-complete", "bash"}); err != nil {
		t.Fatalf("parse: %v", err)
	}

	sbuf = buf.String()

	if !strings.Contains(sbuf, ncoo) {
		t.Error("no coo")
	}

	if strings.Contains(sbuf, nbar) {
		t.Error("has bar")
	}

	if strings.Contains(sbuf, ncud) {
		t.Error("has cud")
	}

	if strings.Contains(sbuf, narg) {
		t.Error("has complete")
	}
}

func TestCompleteVarError(t *testing.T) {
	var (
		fs  = flag.NewFlagSet("test", flag.ContinueOnError)
		err error
		buf bytes.Buffer
	)

	const narg = "complete"

	t.Setenv(envCompLine, "app -c")
	t.Setenv(envCompPoint, "6")

	Var(narg,
		WithFlagSet(fs),
		WithWriter(&buf),
		WithExitFunc(func(int) {}),
	)

	if err = fs.Parse([]string{"-complete", "tch"}); err != nil {
		if !strings.Contains(err.Error(), ErrUnknownShell.Error()) {
			t.Errorf("unexpected error: %v", err)
		}
	} else {
		t.Error("no error")
	}
}
