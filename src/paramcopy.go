package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/atotto/clipboard"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

var service = createSSMService()

var usage = "Usage: \n    paramcopy paramter_name"

func main() {
	param := getName()
	output, err := service.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(param),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		exitErrorf("Unable to describe parameters, %v", err)
	}
	clipboard.WriteAll(*output.Parameter.Value)
	fmt.Printf("Copied %s to clipboard\n", param)
}

func getName() string {
	flag.Parse()

	if flag.NArg() < 1 {
		exitErrorf("Parameter name required.\n%s", usage)
	} else if flag.NArg() > 1 {
		exitErrorf("Too many parameters.\n%s", usage)
	}

	return flag.Arg(0)
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

var version string = "😂👌💯🔥🔥😂"
