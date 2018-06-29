package paramcopy

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

var service = createSSMService()

var usage = fmt.Sprintf("Usage: \n    %s paramter_name", path.Base(os.Args[0]))

func getName() string {
	if flag.NArg() < 1 {
		exitErrorf("Parameter name required.\n%s", usage)
	} else if flag.NArg() > 1 {
		exitErrorf("Too many parameters.\n%s", usage)
	}

	return flag.Arg(0)
}

func GetDecryptedParameter(name string) string {
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
