[![PkgGoDev](https://pkg.go.dev/badge/github.com/s0rg/compflag)](https://pkg.go.dev/github.com/s0rg/compflag)
[![License](https://img.shields.io/github/license/s0rg/compflag)](https://github.com/s0rg/compflag/blob/master/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/s0rg/compflag)](go.mod)
[![Tag](https://img.shields.io/github/v/tag/s0rg/compflag?sort=semver)](https://github.com/s0rg/compflag/tags)

[![CI](https://github.com/s0rg/compflag/workflows/ci/badge.svg)](https://github.com/s0rg/compflag/actions?query=workflow%3Aci)
[![Go Report Card](https://goreportcard.com/badge/github.com/s0rg/compflag)](https://goreportcard.com/report/github.com/s0rg/compflag)
[![Maintainability](https://qlty.sh/badges/40c7eb9f-11aa-4fbd-bac9-0a674b89270b/maintainability.svg)](https://qlty.sh/gh/s0rg/projects/compflag)
[![Code Coverage](https://qlty.sh/badges/40c7eb9f-11aa-4fbd-bac9-0a674b89270b/test_coverage.svg)](https://qlty.sh/gh/s0rg/projects/compflag)
![Issues](https://img.shields.io/github/issues/s0rg/compflag)

# compflag

Auto-completion for stdlib flag items

# usage

you got two options here:

- Call compflag.Complete() somewhere before actual app logic, best point is right at the start:
```go
    package main

    import (
        "os"
        "flag"

        "github.com/s0rg/compflag"
    )

    func main() {
        if compflag.Complete() {
            os.Exit(0)
        }

        flag.Parse()

        // other startup logic...
    }
```

- Define trigger flag for completion:
```go
    package main

    import (
        "os"
        "flag"

        "github.com/s0rg/compflag"
    )

    func main() {
        compflag.Var("complete")

        flag.Parse()

        // other startup logic...
    }

```

Please note, that you need to exit app if any completion happened.

Build your app, put binary somewhere in your "PATH", then run:

```bash
    complete -C %your-binary-name% %your-binary-name%
```

if you prefer flag-triggered version:
```bash
    complete -C "%your-binary-name% -%your-flag% bash" %your-binary-name%
```

Now enter `%your-binary-name%`, and hit `TAB` twice )

# shell compatability

This will work with any shell compatible with `complete` (bash and zsh are both good with it), for zsh you may need
also use `bashcompinit` in addition to `compinit`.
