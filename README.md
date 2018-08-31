# Chronos

Chronos tells the current chronos time. It is also a package that provides
utilities for working with the chronos time system.

The chronos time system splits the day into 16 * 16 * 16 * 16 parts instead of
the usual 24 * 60 * 60. It displays this time using four hexadecimal digits. For
example, 0000 is the start of the day, 8000 is half way through the day and FFFF
is the last part of the day.

This was made because I was bored and I like the idea of a hexadecimal time
system. The logic of the system probably doesn't make any sense but it's cool
and, better yet, it's succint.

## Installation

Run `go get -u github.com/patrickmcnamara/chronos/cmd/chronos`.

Or you can download it from the GitHub release page if you like being weird.

## Usage

Run `chronos` after installation.

Or you can use the chronos package in your own project.
