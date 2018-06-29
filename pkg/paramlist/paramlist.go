package paramlist

import (
	"fmt"
	"os"
	"sort"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

var service = createSSMService()

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

func DescribeParameters(prefixes []string) []string {
	paramNames := []string{}
	if len(prefixes) <= 0 {
		paramNames = getAllParamNames()
	} else {
		for _, prefix := range prefixes {
			for _, name := range getParamNames(prefix) {
				paramNames = appendIfMissing(paramNames, name)
			}
		}
	}
	sort.Strings(paramNames)
	return paramNames
}

func getParamNames(prefix string) []string {
	// Fix passing empty string to SSM API
	if len(prefix) < 1 {
		return getAllParamNames()
	}

	filters := []*ssm.ParametersFilter{&ssm.ParametersFilter{
		Key:    aws.String("Name"),
		Values: []*string{aws.String(prefix)},
	}}
	paramNames := []string{}
	err := service.DescribeParametersPages(&ssm.DescribeParametersInput{
		Filters: filters},
		func(page *ssm.DescribeParametersOutput, lastPage bool) bool {
			for _, parameter := range page.Parameters {
				paramNames = append(paramNames, aws.StringValue(parameter.Name))
			}
			return true
		})

	if err != nil {
		exitErrorf("Unable to describe parameters, %v", err)
	}

	return paramNames
}

func getAllParamNames() []string {
	return getParamNames(" ")
}

func appendIfMissing(slice []string, s string) []string {
	for _, ele := range slice {
		if ele == s {
			return slice
		}
	}
	return append(slice, s)
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
