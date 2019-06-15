[![Travis CI](https://img.shields.io/travis/roaldnefs/gitignore.svg?style=for-the-badge)](https://travis-ci.org/roaldnefs/gitignore)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge)](https://godoc.org/github.com/roaldnefs/gitignore)
[![Github All Releases](https://img.shields.io/github/downloads/roaldnefs/gitignore/total.svg?style=for-the-badge)](https://github.com/roaldnefs/gitignore/releases)
[![GitHub](https://img.shields.io/github/license/roaldnefs/gitignore.svg?style=for-the-badge)](https://github.com/roaldnefs/gitignore/blob/master/LICENSE)

A tool for downloading `.gitignore` templates.

* [Installation](README.md#installation)
     * [Binaries](README.md#binaries)
     * [Via Go](README.md#via-go)
* [Usage](README.md#usage)

## Installation

### Binaries

For installation instructions from binaries please visit the [Releases Page](https://github.com/roaldnefs/gitignore/releases).

### Via Go

```console
$ go get github.com/roaldnefs/gitignore
```

## Usage

```console
$ gitignore --help
Gitignore will create a new .gitignore file in the current working
directory.

Example: gitignore Python -> resulting in a new .gitignore file for Python.

Usage:
  gitignore [language name] [flags]

Flags:
  -h, --help   help for gitignore

$ gitignore go
.gitignore created at /home/roald/go/src/github.com/roaldnefs/gitignore/.gitignore
```
