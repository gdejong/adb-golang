package cmd

import (
	"bufio"
	"github.com/gdejong/adb-golang/pkg/adb"
	"github.com/sirupsen/logrus"
	"strings"

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

		scanner := bufio.NewScanner(strings.NewReader(versionText))
		for scanner.Scan() {
			logrus.Info(scanner.Text())
		}
	},
}
