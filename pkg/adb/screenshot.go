package adb

import (
	"bytes"
	"image"
	"image/png"
	"os"
)

func MakeScreenshot(filename string) {
	b := runAdbCommandRawOutput("exec-out", "screencap", "-p")

	storePng(b.Bytes(), filename)
}

func storePng(b []byte, filename string) {
	out, err := os.Create(filename)
	HandleErr(err)

	r := bytes.NewReader(b)
	img, _, err := image.Decode(r)
	HandleErr(err)

	err = png.Encode(out, img)
	HandleErr(err)
}
