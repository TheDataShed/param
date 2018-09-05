package param

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func Delete(service ssm.SSM, name string) {
	deleteParameter(service, name)
}

func deleteParameter(service ssm.SSM, name string) {
	_, err := service.DeleteParameter(&ssm.DeleteParameterInput{
		Name: aws.String(name),
	})

	if err != nil {
		exitErrorf("Unable to delete parameter, %v", err)
	}
}
