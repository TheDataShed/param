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

var completionFlag = flag.Bool("completion-bash", false, "Print bash completion and exit.")

func main() {
	flag.Parse()
	printBashCompletion()

	param := getName()

	clipboard.WriteAll(getDecryptedParameter(param))
	fmt.Printf("Copied %s to clipboard\n", param)
}

func printBashCompletion() {
	if *completionFlag {
		fmt.Println(bashCompletionFile)
		os.Exit(0)
	}
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

var bashCompletionFile string = `#!/bin/bash
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
`
