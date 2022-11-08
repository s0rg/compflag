package compflag

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	exitCode     = 0
	envCompLine  = "COMP_LINE"
	envCompPoint = "COMP_POINT"
	flagUsage    = "completion-triggering flag"
)

var ErrUnknownShell = errors.New("shell unknown")

// Complete handles completion process started by shell, returns false if no completion was requested.
func Complete(opts ...Option) (ok bool) {
	line, ok := os.LookupEnv(envCompLine)
	if !ok || line == "" {
		return false
	}

	pos, err := strconv.Atoi(os.Getenv(envCompPoint))
	if err != nil {
		return false
	}

	var (
		token = extractToken(line, pos)
		sb    strings.Builder
		cfg   options
	)

	for _, o := range opts {
		o(&cfg)
	}

	cfg.validate()

	cfg.flags.VisitAll(func(f *flag.Flag) {
		if _, ok := cfg.hidden[f.Name]; ok {
			return
		}

		if token == "" || strings.HasPrefix(f.Name, token) {
			_, _ = sb.WriteString("-" + f.Name + "\n")
		}
	})

	_, _ = io.WriteString(cfg.writer, sb.String())

	return true
}

func Var(name string, opts ...Option) {
	var cfg options

	for _, o := range opts {
		o(&cfg)
	}

	cfg.validate()

	cfg.flags.Func(name, flagUsage, func(arg string) (err error) {
		switch arg {
		case "bash", "zsh":
			if Complete(append(opts, WithHidden(name))...) {
				cfg.exit(exitCode)
			}
		default:
			err = fmt.Errorf("%w: %s", ErrUnknownShell, arg)
		}

		return
	})
}

func extractToken(line string, pos int) (val string) {
	if pos > len(line) {
		pos = len(line)
	}

	if t := strings.LastIndexByte(line[:pos], ' '); t != -1 {
		line = strings.TrimSpace(line[t:pos])
	}

	return strings.TrimPrefix(line, "-")
}
