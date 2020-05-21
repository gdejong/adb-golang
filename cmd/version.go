package cmd

import (
	"fmt"
	"github.com/gdejong/adb-golang/pkg/adb"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "adb:version",
	Short: "Print ADB version info",
	Run: func(cmd *cobra.Command, args []string) {
		versionText := adb.GetVersion()

		fmt.Println(versionText)
	},
}
