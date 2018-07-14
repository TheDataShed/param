package param

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
)

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
