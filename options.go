package compflag

import (
	"flag"
	"io"
	"os"
)

type (
	options struct {
		flags  *flag.FlagSet
		writer io.Writer
		exit   func(int)
		hidden map[string]struct{}
	}

	// Option is a functional option type.
	Option func(*options)
)

func (o *options) validate() {
	if o.flags == nil {
		o.flags = flag.CommandLine
	}

	if o.writer == nil {
		o.writer = os.Stdout
	}

	if o.exit == nil {
		o.exit = os.Exit
	}

	if o.hidden == nil {
		o.hidden = make(map[string]struct{})
	}
}

// WithWriter set custom output for completion.
func WithWriter(w io.Writer) Option {
	return func(o *options) {
		o.writer = w
	}
}

// WithFlagSet set custom [flag.FlagSet] as completion source.
func WithFlagSet(fset *flag.FlagSet) Option {
	return func(o *options) {
		o.flags = fset
	}
}

// WithExitFunc set custom exit handler on successful complete.
func WithExitFunc(fn func(int)) Option {
	return func(o *options) {
		o.exit = fn
	}
}

// WithHidden hides given options from completion.
func WithHidden(names ...string) Option {
	return func(o *options) {
		if o.hidden == nil {
			o.hidden = make(map[string]struct{})
		}

		for _, k := range names {
			o.hidden[k] = struct{}{}
		}
	}
}
