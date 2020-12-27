# actions-docs

![ci-workflow](https://github.com/rescDev/actions-docs/workflows/ci/badge.svg)
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/rescDev/actions-docs.svg)](https://github.com/rescDev/actions-docs)

A tool to generate documentation about your Github Action.

```text
actions-docs md
actions-docs md --format=list
actions-docs md --create-file # Creates a new skeleton README with inputs and outputs as table and a simple top-level header and description
```

## Installation

The tool can be installed using `go get`:

```sh
go get github.com/rescDev/actions-docs@latest
```

This will put the binary in `$GOPATH/bin`. Be sure to add the directory to your `$PATH` to avoid getting `actions-docs: command not found` errors.

Binaries are also available on the [releases](https://github.com/rescDev/actions-docs/releases) page. Download one of the binaries (e.g. by using `curl`) and place it into your `$PATH`.

## Usage

Detailed usage examples can be found in the [`examples`](./examples) folder.

```text
NAME:
   actions-docs - Generate documentation for your GitHub Action

USAGE:
   actions-docs [global options] command [command options]

AUTHOR:
   Rene Schach - https://github.com/rescDev

COMMANDS:
   markdown, md  creates Markdown documentation for the GitHub Action
   update        perform a self update of the tool to its latest version

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```
