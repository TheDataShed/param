#!/bin/bash
_paramcopy_completions()
{
    for param in $(./paramlist "${COMP_WORDS[1]}"); do
        COMPREPLY+=("${param}")
    done
}
complete -F _paramcopy_completions ./paramcopy
