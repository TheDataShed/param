package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thedatashed/param/pkg/param"
)

var force bool

var setCmd = &cobra.Command{
	Use:   "set name value",
	Short: "Set a paramter in Parameter Store.",
	Long:  "Add a SecureString paramter to Parameter Store.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		param.Set(args[0], args[1], force)
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
	setCmd.Flags().BoolVarP(&force, "force", "f", false, "Overwrite the parameter if it exists.")

}
