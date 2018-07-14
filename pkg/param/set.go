package param

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func Set(name string, value string, force bool) {
	putSecureStringParameter(name, value, force)
}

func putSecureStringParameter(name string, value string, overwrite bool) {
	_, err := service.PutParameter(&ssm.PutParameterInput{
		Name:      aws.String(name),
		Type:      aws.String("SecureString"),
		Value:     aws.String(value),
		Overwrite: aws.Bool(overwrite),
	})

	if err != nil {
		exitErrorf("Unable to set parameter, %v", err)
	}
}
