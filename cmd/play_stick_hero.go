package cmd

import (
	"github.com/gdejong/adb-golang/pkg/adb"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"image/color"
	"image/draw"
	"log"
	"time"
)

// ScreenYPosition is the height at which a vertical line can be drawn to go through all black areas
// our Stick Hero will walk on.
const screenYPosition = 1800

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

			filename := "screenshots/" + time.Now().Format(time.RFC3339) + ".png"

			ignore := true
			isBlack := true
			transitions := make([]int, 3)
			transitionCount := 0

			for x := 0; x < img.Bounds().Max.X; x++ {
				pixelColor := img.At(x, screenYPosition)
				r, g, b, _ := pixelColor.RGBA()

				// Draw a red vertical line to see at what height we are looking.
				img.(draw.Image).Set(x, screenYPosition, color.RGBA{R: 255, G: 1, B: 1, A: 255})

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

			adb.StoreImage(img, filename)

			start := transitions[0]
			target1 := transitions[1]
			target2 := transitions[2]

			gap := target1 - start
			target := target2 - target1

			distance := int(float64(gap+target/2) * .98)

			logrus.WithFields(logrus.Fields{"start": start, "target1": target1, "target2": target2, "gap": gap, "target": target, "distance": distance}).Debugln("Calculations")

			adb.Swipe(distance)

			time.Sleep(time.Millisecond * 2500)
		}
	},
}
