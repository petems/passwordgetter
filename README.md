# passwordgetter

A basic cli tool that uses `golang.org/x/crypto/ssh/terminal`'s method `terminal.ReadPassword` to get a password.

## Purpose

This was to help me understand stubbing the interface for stdin capture, as I've been writing tests for a lot of CLI apps recently.

The idea was from this StackOverflow question: https://stackoverflow.com/questions/38573176/mocking-crypto-ssh-terminal

And the code was based on https://github.com/isaacasensio/memo - which already had some stubbing in it!

## Installation

#### Via Go

```bash
$ go get github.com/petems/passwordgetter
```

## Usage

```console
$ passwordgetter
Enter password:

Your password was: hunter2
```
