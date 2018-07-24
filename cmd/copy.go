package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thedatashed/param/pkg/param"
)

var verbose bool

var copyCmd = &cobra.Command{
	Use:   "copy name",
	Short: "Copy a parameter to clipboard.",
	Long:  "Copy the specified SSM Parameter from Paramter Store to your clipboard.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		param.Copy(args[0], verbose)
	},
}

func init() {
	RootCmd.AddCommand(copyCmd)
	copyCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Also print parameter value to the stdout.")
}
