package adb

import (
	"bytes"
	"image"
	"image/png"
	"os"
)

func MakeScreenshot() image.Image {
	b := runAdbCommandRawOutput("exec-out", "screencap", "-p")

	r := bytes.NewReader(b.Bytes())
	img, _, err := image.Decode(r)
	HandleErr(err)

	return img
}

func StoreImage(img image.Image, filename string) {
	out, err := os.Create(filename)
	HandleErr(err)

	err = png.Encode(out, img)
	HandleErr(err)
}
