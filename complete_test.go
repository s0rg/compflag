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
		fs   = flag.NewFlagSet("test", flag.ExitOnError)
		buf  bytes.Buffer
		sbuf string
	)

	const (
		nfoo = "foo"
		nbar = "bar"
	)

	_ = fs.Bool(nfoo, false, "")
	_ = fs.String(nbar, "", "")

	t.Setenv(envCompLine, "app ")
	t.Setenv(envCompPoint, "5")

	if !Complete(
		WithWriter(&buf),
		WithFlagSet(fs),
	) {
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

	if !Complete(
		WithWriter(&buf),
		WithFlagSet(fs),
	) {
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

	if !Complete(
		WithWriter(&buf),
		WithFlagSet(fs),
	) {
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
	var (
		fs   = flag.NewFlagSet("test", flag.ExitOnError)
		buf  bytes.Buffer
		sbuf string
	)

	const (
		nfoo = "foo"
		nbar = "bar"
		nfud = "fud"
	)

	_ = fs.Bool(nfoo, false, "")
	_ = fs.String(nbar, "", "")
	_ = fs.Int(nfud, 0, "")

	t.Setenv(envCompLine, "app f")
	t.Setenv(envCompPoint, "6")

	if !Complete(
		WithWriter(&buf),
		WithFlagSet(fs),
		WithHidden(nfud),
	) {
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
