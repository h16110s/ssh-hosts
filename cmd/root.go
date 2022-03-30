package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ssh-hosts",
	Short: "SSH hosts manager",
	Long:  "Manage ~/.ssh/config",
	Run:   ls, // 何も指定しなければlsと同じ挙動
}

var CONFIG_PATH =  "/.ssh/config"

func Execute() error {
	return rootCmd.Execute()
}
