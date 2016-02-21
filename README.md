# Simple GIT (SG)

## Getting Started

### Installing

To start using Simple GIT, install Go and run `go get`:

```sh
$ go get github.com/nazwa/sg
```

This will install the `sg` command line utility into
your `$GOBIN` path.

### Usage

```sh
$ sg
```

This will commit with a default message `Updates...`, pull and push from remote repository.
You can also specity a custom message, like so:

```sh
$ sg My super new feature
```
