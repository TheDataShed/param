package param

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/ssm"
)

func Show(service ssm.SSM, name string) {
	value := getDecryptedParameter(service, name)

	fmt.Println(value)
}
