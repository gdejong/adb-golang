package cmd

import (
	"github.com/gdejong/adb-golang/pkg/adb"
	"github.com/spf13/cobra"
	"log"
	"time"
)

func init() {
	rootCmd.AddCommand(playStickHeroCommand)
}

var playStickHeroCommand = &cobra.Command{
	Use:   "adb:play-stick-hero",
	Short: "Play the Stick Hero game",
	Run: func(cmd *cobra.Command, args []string) {
		// This is basically just a Golang implementation of https://www.youtube.com/watch?v=Du__JfXqsAs
		// Check out his videos, he has a great YouTube channel!

		// Verify the connection to the device.
		_, err := adb.GetAdbSerialNumber()
		adb.HandleErr(err)

		on := adb.IsOn()
		if !on {
			log.Fatalln("Device screen is off, please unlock the device.")
		}

		for {
			img := adb.MakeScreenshot()

			ignore := true
			isBlack := true
			transitions := make([]int, 3)
			transitionCount := 0

			for x := 0; x < img.Bounds().Max.X; x++ {
				color := img.At(x, 1800)
				r, g, b, _ := color.RGBA()

				// Ignore pixels until we encounter the first black area.
				if ignore && (r+g+b != 0) {
					continue
				}

				ignore = false

				if isBlack && (r+g+b != 0) {
					isBlack = !isBlack

					transitions[transitionCount] = x
					transitionCount++

					continue
				}

				if !isBlack && (r+g+b == 0) {
					isBlack = !isBlack

					transitions[transitionCount] = x
					transitionCount++

					continue
				}
			}

			start := transitions[0]
			target1 := transitions[1]
			target2 := transitions[2]

			gap := target1 - start
			target := target2 - target1

			distance := int(float64(gap+target/2) * .98)

			adb.Swipe(distance)

			time.Sleep(time.Millisecond * 2500)
		}
	},
}
