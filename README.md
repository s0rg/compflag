[![PkgGoDev](https://pkg.go.dev/badge/github.com/s0rg/compflag)](https://pkg.go.dev/github.com/s0rg/compflag)
[![License](https://img.shields.io/github/license/s0rg/compflag)](https://github.com/s0rg/compflag/blob/master/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/s0rg/compflag)](go.mod)
[![Tag](https://img.shields.io/github/v/tag/s0rg/compflag?sort=semver)](https://github.com/s0rg/compflag/tags)

[![CI](https://github.com/s0rg/compflag/workflows/ci/badge.svg)](https://github.com/s0rg/compflag/actions?query=workflow%3Aci)
[![Go Report Card](https://goreportcard.com/badge/github.com/s0rg/compflag)](https://goreportcard.com/report/github.com/s0rg/compflag)
[![Maintainability](https://api.codeclimate.com/v1/badges/b1ab20a6dd9536e9fbc8/maintainability)](https://codeclimate.com/github/s0rg/compflag/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/b1ab20a6dd9536e9fbc8/test_coverage)](https://codeclimate.com/github/s0rg/compflag/test_coverage)
![Issues](https://img.shields.io/github/issues/s0rg/compflag)

# compflag

Auto-completion for stdlib flag items

# usage

Call compflag.Complete() somewhere before actual app logic, best point is right at the start:

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

Please note, that you need to exit app if any completion happened.

Build your app, put binary somewhere in your "PATH", then run:

```bash
    complete -C %your-binary-name% %your-binary-name%
```

Now enter `%your-binary-name%`, and hit `TAB` twice )

# shell compatability

This will work with any shell compatible with `complete` (bash and zsh are both good with it), for zsh you may need
also use `bashcompinit` in addition to `compinit`.
