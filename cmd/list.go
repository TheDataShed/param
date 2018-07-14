package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/willjcj/param/pkg/param"
)

var prefixes string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List parameters in Parameter Store.",
	Long: `List all parameters from parameter store with an optional prefix.
    Results are sorted in alphabetical order.`,
	Run: func(cmd *cobra.Command, args []string) {
		prefixSlice := strings.Split(prefixes, ",")
		param.List(prefixSlice)
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

	listCmd.Flags().StringVarP(&prefixes, "prefix", "p", "", "Prefixes to fileter by")
}
