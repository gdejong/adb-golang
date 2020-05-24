package cmd

import (
	"github.com/gdejong/adb-golang/pkg/adb"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var verbose bool

var rootCmd = &cobra.Command{
	Use: "adb",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if verbose {
			logrus.SetLevel(logrus.DebugLevel)
		} else {
			logrus.SetLevel(logrus.InfoLevel)
		}
	},
}

func Execute() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "make logging verbose")

	err := rootCmd.Execute()
	adb.HandleErr(err)
}
