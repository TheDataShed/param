# WIP

## TODO

- Move both utilities into one binary.
- Distribute as a release.
- Write tests!

## paramcopy

Copies the specified SSM parameter to your clipboard.

## Usage

Make sure valid AWS credentials are accessible in your terminal.

    go build pkg/paramcopy.go

    ./paramcopy parameter

Copies the specified parameter to your clipboard.

### Bash Completion

For now, you need the paramlist binary in the same directory for bash completion.

    source <(./paramcopy -completion-bash)
    go build pkg/paramlist.go

## paramlist

There's also a untility to list SSM parameters.

It was creted for use in the bash completion.

You can specify at least one prefix to filter the parameter names by.
