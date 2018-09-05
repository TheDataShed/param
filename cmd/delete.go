package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thedatashed/param/pkg/param"
)

var deleteCmd = &cobra.Command{
	Use:   "delete name value",
	Short: "Delete a parameter.",
	Long:  "Delete the given paramter from Parameter Store.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		param.Delete(*createSSMService(), args[0])
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
