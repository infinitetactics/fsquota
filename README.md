# fsquota

[![license](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://github.com/infinitetactics/fsquota/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/infinitetactics/fsquota?status.svg)](https://godoc.org/github.com/infinitetactics/fsquota)
[![Build Status](https://travis-ci.org/infinitetactics/fsquota.svg?branch=master)](https://travis-ci.org/infinitetactics/fsquota)
[![Go Report Card](https://goreportcard.com/badge/github.com/infinitetactics/fsquota)](https://goreportcard.com/report/github.com/infinitetactics/fsquota)


fsquota is a native Go library for interacting with (Linux) filesystem quotas.
This library does **not** make use of cgo or invoke external commands, but rather interacts directly with the kernel interface by use of syscalls.
This library was originally written by the [Anexia](https://www.anexia-it.com/) R&D team. This repository has a few modifications for our use by InfiniteTactics.t

## Portability

fsquota has been developed with Linux in mind and as such only supports Linux for now.
Support for other platforms may be added in the future.

## fsqm

This repository also ships *fsqm*, a simple command line interface to filesystem quotas. *fsqm* provides the ability to retrieve user and group quota reports and management of user and group quotas.

## Status

The current release is **v0.1.8**.


Changes to fsquota are subject to [semantic versioning](http://semver.org/).

## License

fsquota is licensed under the terms of the [MIT license](https://github.com/infinitetactics/fsquota/blob/master/LICENSE).
