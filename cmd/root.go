package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "param",
	Short: "Tools to improve Parameter Store on the command line.",
	Long:  "Param is a cli tool to improve interacting with AWS Parameter Store.",
	BashCompletionFunction: bash_completion_func,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.param.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".param" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".param")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

const (
	bash_completion_func = `__param_list()
{
    local words_no_flags=()
    for word in ${COMP_WORDS[@]}; do
        if [[ "${word}" != -* ]]; then
            words_no_flags+=("${word}")
        fi
    done

    # Check the word to complete is the 3rd one, not including flags
    if [ "${#words_no_flags[@]}" -gt "3" ]; then
        return
    fi

    # Cache parameter names as an env var.
    #
    # Ignore cache:
    # export PARAM_NO_CACHE=1
    #
    # Clear cache:
    # unset PARAM_CACHE
    local param_names=(${PARAM_CACHE[@]})
    if [ -z "${param_names[*]}" ] || [ "${PARAM_NO_CACHE}" = "1" ]; then
        param_names=($(param list 2>/dev/null | awk '{print $1}'))
    fi

    # If PARAM_NO_CACHE isn't set, then export the param_names as PARAM_CACHE
    if [ "${PARAM_NO_CACHE}" != "1" ]; then
        export PARAM_CACHE=(${param_names[@]})
    else
        unset PARAM_CACHE
    fi

    COMPREPLY=( $( compgen -W "${param_names[*]}" -- "${cur}" ) )
}

__custom_func() {
    case ${last_command} in
        param_copy | param_show)
            __param_list
            return
            ;;
        *)
            ;;
    esac
}
`
)
