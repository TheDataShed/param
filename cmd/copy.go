package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"github.com/willjcj/param/pkg/paramget"
)

var verbose bool

var copyCmd = &cobra.Command{
	Use:   "copy parameter_name",
	Short: "Copy a parameter to clipboard.",
	Long:  "Copy the specified SSM Parameter from Paramter Store to your clipboard.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		copyToClipboard(args[0])
	},
}

func init() {
	RootCmd.AddCommand(copyCmd)
	copyCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Also print parameter value to the stdout.")
}

func copyToClipboard(name string) {
	value := paramget.GetDecryptedParameter(name)

	clipboard.WriteAll(value)

	if verbose {
		fmt.Println(value)
	}
}
