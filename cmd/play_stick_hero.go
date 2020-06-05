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
const screenYPosition = 1600

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

		screenSize := adb.GetScreenResolution()
		logrus.WithField("size", screenSize).Debugln("Screen size")

		colorRed := color.RGBA{R: 255, G: 1, B: 1, A: 255}
		colorGreen := color.RGBA{R: 0, G: 255, B: 1, A: 255}
		colorBlue := color.RGBA{R: 0, G: 0, B: 255, A: 255}
		colorYellow := color.RGBA{R: 255, G: 255, B: 0, A: 255}

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
				img.(draw.Image).Set(x, screenYPosition, colorRed)

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

			currentPlatformEnd := transitions[0]
			nextPlatformStart := transitions[1]
			nextPlatformEnd := transitions[2]

			gap := nextPlatformStart - currentPlatformEnd
			nextPlatformWidth := nextPlatformEnd - nextPlatformStart
			distance := int(float64(gap + nextPlatformWidth/2))

			// Draw the gap
			for i := currentPlatformEnd; i < nextPlatformStart; i++ {
				img.(draw.Image).Set(i, screenYPosition+2, colorGreen)
			}

			// Draw the second platform
			for i := nextPlatformStart; i < nextPlatformEnd; i++ {
				img.(draw.Image).Set(i, screenYPosition+4, colorBlue)
			}

			// Draw the distance
			for i := currentPlatformEnd; i < currentPlatformEnd+distance; i++ {
				img.(draw.Image).Set(i, screenYPosition+6, colorYellow)
			}

			adb.StoreImage(img, filename)

			logrus.WithFields(logrus.Fields{
				"currentPlatformEnd": currentPlatformEnd,
				"nextPlatformStart":  nextPlatformStart,
				"nextPlatformEnd":    nextPlatformEnd,
				"gap":                gap,
				"nextPlatformWidth":  nextPlatformWidth,
				"distance":           distance,
				"filename":           filename,
			}).Debugln("Calculations")

			if distance == 0 {
				logrus.Fatalln("distance is zero, this can't be right")
			}

			msToSwipe := int(float64(distance) * 0.99)

			adb.Swipe(msToSwipe)

			time.Sleep(time.Millisecond * 3000)
		}
	},
}
