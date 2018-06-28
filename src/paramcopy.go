package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/atotto/clipboard"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

var service = createSSMService()

var usage = fmt.Sprintf("Usage: \n    %s paramter_name", path.Base(os.Args[0]))

var completionPrintFlag = flag.Bool("completion-bash", false, "Print bash completion and exit.")
var completionClearCacheFlag = flag.Bool("clear", false, "Clear bash completion cache.")

func main() {
	flag.Parse()

	printBashCompletion()
	clearCompletionCache()

	param := getName()

	clipboard.WriteAll(getDecryptedParameter(param))
	fmt.Printf("Copied %s to clipboard\n", param)
}
func getName() string {
	if flag.NArg() < 1 {
		exitErrorf("Parameter name required.\n%s", usage)
	} else if flag.NArg() > 1 {
		exitErrorf("Too many parameters.\n%s", usage)
	}

	return flag.Arg(0)
}

func getDecryptedParameter(name string) string {
	output, err := service.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(name),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		exitErrorf("Unable to describe parameters, %v", err)
	}
	return *output.Parameter.Value
}

func createSSMService() *ssm.SSM {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1")},
	)

	if err != nil {
		exitErrorf("Unable to describe parameters, %v", err)
	}

	// Create SSM service client
	return ssm.New(sess)
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

// Print bash completion string and exit.
func printBashCompletion() {
	if *completionPrintFlag {
		fmt.Println(bashCompletionFile)
		os.Exit(0)
	}
}

// Clear bash completion cache and exit.
func clearCompletionCache() {
	if *completionClearCacheFlag {
		os.Setenv("PARAM_COPY_CACHE", "")
		fmt.Println("PARAM_COPY_CACHE removed")
		os.Exit(0)
	}
}

var bashCompletionFile string = `#!/bin/bash
_paramcopy_completions()
{
    # Don't complete if there's already an argument
    if [ "${#COMP_WORDS[@]}" != "2" ]; then
        return
    fi

    # Complete flags
    if [[ ${COMP_WORDS[1]:0:1} == "-" ]] ; then
        SUGGEST="-completion-bash -clear"
    # If not a flag, get a list of possible parameters
    else
        # Cache returned parameters as an env var.
        if [ -z "${PARAM_COPY_CACHE}" ]; then
            export PARAM_COPY_CACHE=$(./paramlist)
        fi
        SUGGEST=${PARAM_COPY_CACHE}
    fi

    # Make sure to add -- so it doesn't try to parse the last
    # argument as a flag if it starts with a -
    local suggestions=$(compgen -W "${SUGGEST}" -- "${COMP_WORDS[1]}")
    for param in ${suggestions}; do
        COMPREPLY+=("${param}")
    done
}
complete -F _paramcopy_completions ./paramcopy
`
