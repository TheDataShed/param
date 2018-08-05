package param

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func Delete(name string) {
	deleteParameter(name)
}

func deleteParameter(name string) {
	_, err := service.DeleteParameter(&ssm.DeleteParameterInput{
		Name: aws.String(name),
	})

	if err != nil {
		exitErrorf("Unable to delete parameter, %v", err)
	}
}
