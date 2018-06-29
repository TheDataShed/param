package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/atotto/clipboard"
	"github.com/posener/complete"
	"github.com/willjcj/param/pkg/paramcopy"
	"github.com/willjcj/param/pkg/paramlist"
)

func main() {

	// create the complete command
	cmp := complete.New(
		"param",
		complete.Command{
			Flags: complete.Flags{
				"-complete":   complete.PredictNothing,
				"-uncomplete": complete.PredictNothing,
				"-h":          complete.PredictNothing,
				"-help":       complete.PredictNothing,
				"-y":          complete.PredictNothing,
			},
			Sub: complete.Commands{
				"copy": complete.Command{
					// TODO fill with list output
					Args: complete.PredictNothing,
				},
				"list": complete.Command{
					// TODO fill with list output
					Args: complete.PredictNothing,
				},
			},
		},
	)

	cmp.InstallName = "complete"
	cmp.UninstallName = "uncomplete"
	cmp.AddFlags(nil)

	// parse the flags - both the program's flags and the completion flags
	flag.Parse()

	// run the completion, in case that the completion was invoked
	// and ran as a completion script or handled a flag that passed
	// as argument, the Run method will return true,
	// in that case, our program have nothing to do and should return.
	if cmp.Complete() {
		return
	}

	// command := os.Args[0]
	subCommand := os.Args[1]

	if subCommand == "copy" {
		paramName := os.Args[2]

		decryptedParameter := paramcopy.GetDecryptedParameter(paramName)

		clipboard.WriteAll(decryptedParameter)
		fmt.Printf("Copied %s to clipboard.\n", paramName)
	} else if subCommand == "list" {
		prefixes := os.Args[2:]

		params := paramlist.DescribeParameters(prefixes)

		for _, param := range params {
			fmt.Println(param)
		}
	}
}
