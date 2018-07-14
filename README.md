# param

[![Build Status](https://travis-ci.com/WillJCJ/param.svg?branch=master)](https://travis-ci.com/WillJCJ/param)

- [Introduction](#introduction)
- [Install](#install)
- [Usage](#usage)
- [Shell Completion](#shell-completion)
- [Docs](#docs)
- [Todo](#todo)

## Introduction

`param` is a cli tool for talking to AWS Parameter Store. Copy parameters
directly to your clipboard.

`param` uses [Cobra](https://github.com/spf13/cobra), a library providing a
simple interface to create powerful CLI interfaces similar to git & go tools.

`param` works with Linux and MacOS.
It also supports `bash` and `zsh` completion.

## Install

You can install `param` by downloading the latest binary from the
[Releases](https://github.com/WillJCJ/param/releases) page or by compiling
from source with `go`.

```console
go get github.com/willjcj/param
```

This should build a binary at `$GOPATH/bin/param`

Add `$GOPATH/bin` to your `PATH`, or move param to somewhere
already on your `PATH`.

### Curl

Run the below command to download the 1.5.0 binary and add it to
`/usr/local/bin`.

#### Linux

```console
curl -LO https://github.com/WillJCJ/param/releases/download/1.5.0/param-linux-amd64 && \
chmod +x param-linux-amd64 && \
sudo mv param-linux-amd64 /usr/local/bin/param
```

#### MacOS

```console
curl -LO https://github.com/WillJCJ/param/releases/download/1.5.0/param-darwin-amd64 && \
chmod +x param-darwin-amd64 && \
sudo mv param-darwin-amd64 /usr/local/bin/param
```

## Usage

Make sure your terminal session has the correct AWS credentials.

Below is a brief overview for each command.
Full docs for each command can be found at [`/docs`](docs/param.md).


### Copy

Copy a parameter to your clipboard:

```console
$ param copy parameter_name
```

You can optionally show the parameter value in your console with the `-v` flag:

```console
$ param copy parameter_name -v
password123
```

#### Auto-completion

With shell completion enabled, you can press tab to auto-complete the parameter
names.

### List

Get a sorted list of parameters in SSM with optional prefix(es):

```console
$ param list
parameter.name.key
parameter2.name.password
...
```

```console
$ param list -p prefix1,prefix2
prefix1.dev.password
prefix1.prod.password
prefix2.key
```

### Set

Set a parameter with type `SecureString`:

```console
$ param set parameter_name password123
```

If the parameter already exists, you must specify the `-f` flag
to overwrite it:

```console
$ param set parameter_name password456 -f
```

#### Auto-completion

With shell completion enabled, you can press tab to auto-complete the parameter
names.

### Show

If you'd like to print out the decrypted parameter without copying it the
clipboard, you can use:

```console
$ param show parameter_name
password123
```

#### Auto-completion

With shell completion enabled, you can press tab to auto-complete the parameter
names.

## Shell Completion

`param completion (bash|zsh)` outputs shell completion code for the specified
shell (`bash` or `zsh`).
The shell code must be evaluated to provide interactive completion of
param commands.
This can be done by sourcing it from `~/.bashrc` or `~/.zshrc`

### Installation

#### bash

Source the bash completion script every time you start a shell
by adding a line to your `~/.bashrc` file:

```bash
source <(param completion bash)
```

You can also run this command to append it for you:

```console
printf "
# param shell completion
source <(param completion bash)
" >> $HOME/.bashrc
source $HOME/.bashrc
```

### Caching

Because the call to SSM can be quite slow, `param copy` caches the list of
parameter names by exporting an array named `PARAM_CACHE`.

You can clear the cache by unsetting the cache array:
```bash
unset PARAM_CACHE
```

You can stop caching entirely by setting the `PARAM_NO_CACHE` variable to `1`:

```bash
export PARAM_NO_CACHE=1
```

## Docs

Docs for each command can be found at [`/docs`](docs/param.md).

They are generated by Cobra by running `go run build/generate_docs.go`.

## CI

`param` uses Travis CI to test each commit. It also handles GitHub releases when
a new tag is pushed to GitHub.

## TODO

- New Commands
	- Delete command - Command to delete parameters
		- Maybe with a `--yes` flag to confim.
- Improvements
	- Add a flag to specify parameter type.
	Currently only works with `SecureString`s.
	- Better logging and a verbose option to see what calls are made to AWS.
- Shell Completion
	- `zsh` completion doesn't seem to work.
	- Update or delete the bash completion cache after creating/deleting
	parameters
	- Add bash completion for other subcommands and flags.
	- Command to reset the cache.
- Write Tests
- Improve Documentation
	- More Examples
	- Write about how to set up AWS variables/profile.
	(And what capabilities are required)
