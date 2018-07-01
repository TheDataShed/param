# WIP

## TODO

- Set command - Command to set parameters
- Common SSM service instead of duplicated in all commands.
- Write tests!
- Write about how to set up AWS variables/profile.
- Add global type flag String/SecureString

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

    $ param list -p prefix1,prefix2 ...
    prefix1.dev.password
    prefix1.prod.password
    prefix2.key


## Docs

Docs for each command can be found at `/docs`
