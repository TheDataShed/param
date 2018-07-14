package param

import (
	"fmt"

	"github.com/atotto/clipboard"
)

func Copy(name string, verbose bool) {
	value := getDecryptedParameter(name)

	clipboard.WriteAll(value)

	if verbose {
		fmt.Println(value)
	}
}
