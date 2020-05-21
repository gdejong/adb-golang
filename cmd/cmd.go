package cmd

import (
	"github.com/gdejong/adb-golang/pkg/adb"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "adb"}

func Execute() {
	err := rootCmd.Execute()
	adb.HandleErr(err)
}
