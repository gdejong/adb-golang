package cmd

import (
	"github.com/gdejong/adb-golang/pkg/adb"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(testConnectionCommand)
}

var testConnectionCommand = &cobra.Command{
	Use:   "adb:test-connection",
	Short: "Test if a device is connected",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Looking for connected device...")
		serialNumber, err := adb.GetAdbSerialNumber()
		adb.HandleErr(err)

		log.Println("Detected ADB serial number: " + serialNumber)
	},
}
