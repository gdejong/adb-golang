package adb

import (
	"strconv"
)

func Swipe(msDuration int) {
	msDurationString := strconv.Itoa(msDuration)

	runAdbCommand("shell", "input", "touchscreen", "swipe", "500", "500", "500", "500", msDurationString)
}
