/*
Copyright © 2022 nanvenomous mrgarelli@gmail.com

*/
package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/nanvenomous/e/system"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	pkgName = "e"
	review  bool
)

func stringContainsPartOfOneOfSlice(candidate string, slc []string) bool {
	for _, elm := range slc {
		if strings.Contains(candidate, elm) {
			return true
		}
	}
	return false
}

func openFilesForReview(args []string) error {
	var (
		err         error
		chdFilesStr string
		errStr      string
		nvimArgs    = []string{"-p"} // nvim -p opens the files in tabs
		toIgnore    = viper.GetStringSlice("ignore")
		diffIndex   = []string{"HEAD", "HEAD~1"}
		idx         int
	)
	if len(args) > 0 {
		idx, err = strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		if idx > 1 {
			diffIndex = []string{fmt.Sprintf("HEAD~%d", idx-1), fmt.Sprintf("HEAD~%d", idx)}
		}
	}
	chdFilesStr, errStr, err = system.Capture("git", append([]string{"diff", "--name-only"}, diffIndex...))
	if err != nil {
		fmt.Println(errStr)
		return err
	}
	for _, fl := range strings.Split(chdFilesStr, "\n") {
		if _, err := os.Stat(fl); err == nil { // file exists exists
			if !stringContainsPartOfOneOfSlice(fl, toIgnore) {
				nvimArgs = append(nvimArgs, fl)
			}
		}
	}
	return system.Nvim(nvimArgs)
}

var rootCmd = &cobra.Command{
	Use:   pkgName,
	Short: "nvim wrapper for sessions",
	Long:  `nvim wrapper for sessions`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			err error
		)

		if completion != "" {
			return genCompletions(cmd)

		}

		if review {
			return openFilesForReview(args)
		}

		if len(args) > 0 {
			err := system.Nvim(args)
			if err != nil {
				return err
			}
			os.Exit(0)
		}
		fileName := "Session.vim"
		path := "./"
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		ndirs := strings.Split(wd, "/")
		for range ndirs {
			fullFilePath := path + fileName
			if _, err := os.Stat(fullFilePath); !os.IsNotExist(err) {
				err := system.Nvim([]string{"-S", fullFilePath})
				if err != nil {
					return err
				}
				return nil
			}
			if path == "./" {
				path = "../"
			} else {
				path += "../"
			}
		}
		system.Nvim([]string{})
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func init() {
	initRoot()
	rootCmd.Flags().BoolVarP(&review, "review", "r", false, "open all the contents to review")
}
