package cmd

import (
	"github.com/spf13/cobra"
	"github.com/willjcj/param/pkg/param"
)

var showCmd = &cobra.Command{
	Use:   "show parameter_name",
	Short: "Show a decrypted parameter in the console.",
	Long:  "Show the specified decrypted SSM Parameter from Paramter Store in your console.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		param.Show(args[0])
	},
}

func init() {
	RootCmd.AddCommand(showCmd)
}
