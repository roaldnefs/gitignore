<a href="https://github.com/roaldnefs/gitignore" style="color: black;">
    <h1 align="center">gitignore</h1>
</a>
<p align="center">
    <a href="https://github.com/roaldnefs/gitignore/releases">
        <img src="https://img.shields.io/github/v/release/roaldnefs/gitignore?style=for-the-badge&color=blue"
            alt="Latest release version">
    </a>
    <a href="https://github.com/roaldnefs/gitignore/blob/master/LICENSE">
        <img src="https://img.shields.io/github/license/roaldnefs/gitignore.svg?style=for-the-badge&color=blue"
            alt="GitHub - License">
    </a>
    <a href="https://github.com/roaldnefs/gitignore/actions">
        <img src="https://img.shields.io/github/workflow/status/roaldnefs/gitignore/build?style=for-the-badge&color=blue"
            alt="GitHub Workflow Status">
    </a>
    <a href="https://github.com/roaldnefs/gitignore/graphs/contributors">
        <img src="https://img.shields.io/github/contributors/roaldnefs/gitignore?style=for-the-badge&color=blue"
            alt="GitHub contributors">
    </a>
    </br>
    <b>gitignore</b> is a tool for downloading `.gitignore` templates..
    <br />
    <a href="https://github.com/roaldnefs/gitignore/releases"><strong>Download Latest Release »</strong></a>
    <br />
    <a href="https://github.com/roaldnefs/gitignore/issues/new?title=Bug%3A">Report Bug</a>
    ·
    <a href="https://github.com/roaldnefs/gitignore/issues/new?title=Feature+Request%3A">Request Feature</a>
</p>

## Introduction

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
  -g, --global   Search globally useful gitignores
  -h, --help     help for gitignore

$ gitignore go
.gitignore created at /home/roald/go/src/github.com/roaldnefs/gitignore/.gitignore
```
