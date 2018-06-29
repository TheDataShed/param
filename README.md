# WIP


## TODO

- bash completion for parameter names.
- Distribute as a release (On multiple platforms).
https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04
- Set parameters.
- Common SSM service instead of duplicated in both.
- Write tests!

## Usage

Make sure your terminal session has the correct AWS credentials.

Copy a parameter to your clipboard:

    $ param copy parameter_name
    Copied parameter_name to clipboard.

Get a sorted list of parameters in SSM with optional prefix(es):

    $ param list
    parameter_name1
    parameter_name2
    ...

    $ param list prefix1 prefix2 ...
    prefix1.dev.password
    prefix1.prod.password
    prefix2.key
