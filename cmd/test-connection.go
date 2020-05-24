package cmd

import (
	"github.com/gdejong/adb-golang/pkg/adb"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(testConnectionCommand)
}

var testConnectionCommand = &cobra.Command{
	Use:   "adb:test-connection",
	Short: "Test if a device is connected",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Debugln("Looking for connected device...")
		serialNumber, err := adb.GetAdbSerialNumber()
		adb.HandleErr(err)

		logrus.WithField("serial", serialNumber).Infoln("Detected serial number")
	},
}
