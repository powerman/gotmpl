# gotmpl [![GoDoc](https://godoc.org/github.com/powerman/gotmpl?status.svg)](http://godoc.org/github.com/powerman/gotmpl) [![Go Report Card](https://goreportcard.com/badge/github.com/powerman/gotmpl)](https://goreportcard.com/report/github.com/powerman/gotmpl) [![CircleCI](https://circleci.com/gh/powerman/gotmpl.svg?style=svg)](https://circleci.com/gh/powerman/gotmpl) [![Coverage Status](https://coveralls.io/repos/github/powerman/gotmpl/badge.svg?branch=master)](https://coveralls.io/github/powerman/gotmpl?branch=master)

Command line tool for processing template file using Go text/template syntax.

## Installation

```sh
go get github.com/powerman/gotmpl
```

Or use static binary:

```sh
curl -s -L -o /usr/local/bin/gotmpl \
    https://github.com/powerman/gotmpl/releases/download/v1.1.1/gotmpl-`uname -s`-`uname -m`
chmod +x /usr/local/bin/gotmpl
```

## Usage

```sh
gotmpl <config.yaml.tmpl >config.yaml
```
