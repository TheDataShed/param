package param

import (
	"fmt"
)

func Show(name string) {
	value := getDecryptedParameter(name)

	fmt.Println(value)
}
