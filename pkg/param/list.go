package param

import (
	"fmt"
	"sort"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func List(prefixes []string) {
	for _, param := range describeParameters(prefixes) {
		fmt.Println(param)
	}
}

func describeParameters(prefixes []string) []string {
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
