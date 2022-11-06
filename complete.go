package compflag

import (
	"flag"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	envCompLine  = "COMP_LINE"
	envCompPoint = "COMP_POINT"
)

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

func extractToken(line string, pos int) (val string) {
	if pos > len(line) {
		pos = len(line)
	}

	if t := strings.LastIndexByte(line[:pos], ' '); t != -1 {
		line = strings.TrimSpace(line[t:pos])
	}

	return strings.TrimPrefix(line, "-")
}
