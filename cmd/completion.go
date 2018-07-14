package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion (bash|zsh)",
	Short: "Generates shell completion scripts",
	Long: `Output shell completion code for the specified shell (bash or zsh).
The shell code must be evaluated to provide interactive completion of param commands.
This can be done by sourcing it from ~/.bashrc

Examples:
    # Installing bash completion
    printf "
    # param shell completion
    source <(param completion bash)
    " >> $HOME/.bashrc
    source $HOME/.bashrc`,
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"bash", "zsh"},
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "bash" {
			RootCmd.GenBashCompletion(os.Stdout)
		} else if args[0] == "zsh" {
			RootCmd.GenZshCompletion(os.Stdout)
		}
	},
}

func init() {
	RootCmd.AddCommand(completionCmd)
}
