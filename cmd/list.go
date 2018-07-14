package cmd

import (
	"fmt"
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
		listParameters()
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

	listCmd.Flags().StringVarP(&prefixes, "prefix", "p", "", "Prefixes to fileter by")
}

func listParameters() {
	prefixSlice := strings.Split(prefixes, ",")
	for _, param := range param.DescribeParameters(prefixSlice) {
		fmt.Println(param)
	}
}
