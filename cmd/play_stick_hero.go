package cmd

import (
	"github.com/gdejong/adb-golang/pkg/adb"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(playStickHeroCommand)
}

var playStickHeroCommand = &cobra.Command{
	Use:   "adb:play-stick-hero",
	Short: "Play the Stick Hero game",
	Run: func(cmd *cobra.Command, args []string) {
		// Verify the connection to the device.
		_, err := adb.GetAdbSerialNumber()
		adb.HandleErr(err)

		on := adb.IsOn()
		if !on {
			log.Fatalln("Device screen is off, please unlock the device.")
		}

		img := adb.MakeScreenshot()

		adb.StoreImage(img, "test.png")
		adb.Swipe(100)
	},
}
