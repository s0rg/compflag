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

// WithFlagSet set custom output for completion.
func WithFlagSet(fset *flag.FlagSet) Option {
	return func(o *options) {
		o.flags = fset
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
