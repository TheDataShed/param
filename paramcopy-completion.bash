#!/bin/bash
_paramcopy_completions()
{
    # Don't complete if there's already an argument
    if [ "${#COMP_WORDS[@]}" != "2" ]; then
        return
    fi

    # Cache returned parameters as an env var.
    # export PARAM_COPY_NO_CACHE=1 to ignore the cache.
    if [ -z "${PARAM_COPY_CACHE}" ] || [ "${PARAM_COPY_NO_CACHE}" = "1" ]; then
        export PARAM_COPY_CACHE=$(./paramlist "${COMP_WORDS[1]}")
    fi
    local suggestions=$(compgen -W "${PARAM_COPY_CACHE}" "${COMP_WORDS[1]}")
    # Use paramlist to complete the paramcopy command
    for param in ${suggestions}; do
        COMPREPLY+=("${param}")
    done
}
complete -F _paramcopy_completions ./paramcopy
