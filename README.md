# WIP

## TODO

- bash completion for parameter names
- Distribute as a release.
- Common SSM service instead of duplicated in both.
- Write tests!

## Usage

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
