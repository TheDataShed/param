package param

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func Copy(service ssm.SSM, name string, verbose bool) {
	value := getDecryptedParameter(service, name)

	clipboard.WriteAll(value)

	if verbose {
		fmt.Println(value)
	}
}
