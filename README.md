# Chronos

A library and command line tool for the chronos time system.

The chronos time system splits the day into 16 * 16 * 16 * 16 moments instead of 24 * 60 * 60 seconds.
This is then represented using a four digit hexadecimal number.
For example, *0000* is the start of the day, *8000* is half way through the day and *FFFF* is the last moment of the day.

This was made because I was bored and I like the idea of a hexadecimal time system.
It looks nicer and is more succint.

## Installation

Run `go get -u github.com/patrickmcnamara/chronos/...`.

## Usage

Run `chronos` after installation, assuming that your `$GOPATH` is in your `$PATH`. \
Or you can use the chronos package in your own project.
