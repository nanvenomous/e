package cmd

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	shells     = []string{"bash", "zsh", "fish", "powershell"}
	completion string
	cfgFile    string
)

func genCompletions(rtCmd *cobra.Command) error {
	switch completion {
	case shells[0]:
		rtCmd.Root().GenBashCompletion(os.Stdout)
	case shells[1]:
		rtCmd.Root().GenZshCompletion(os.Stdout)
	case shells[2]:
		rtCmd.Root().GenFishCompletion(os.Stdout, true)
	case shells[3]:
		rtCmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
	default:
		return errors.New("not a recognized shell")
	}
	return nil
}

func initRoot() {
	cobra.OnInitialize(initConfig)
	rootCmd.SilenceUsage = true
	rootCmd.SilenceErrors = true
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	completionFlag := "completion"
	rootCmd.PersistentFlags().StringVar(&completion, completionFlag, "", "generate shell completion")
	rootCmd.RegisterFlagCompletionFunc(completionFlag, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return shells, cobra.ShellCompDirectiveDefault
	})

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is $HOME/.config/%s.yaml)", pkgName))

}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		viper.AddConfigPath(path.Join(home, ".config"))
		viper.SetConfigName(pkgName)
	}

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err == nil {
		// fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
