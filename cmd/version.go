package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version",
	Long:  `show version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ssh-host-manager version " + "0.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
