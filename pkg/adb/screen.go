package adb

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type ScreenResolution struct {
	Width  int
	Height int
}

func (s ScreenResolution) String() string {
	return fmt.Sprintf("width: %d, height: %d", s.Width, s.Height)
}

func IsOn() bool {
	// https://stackoverflow.com/questions/35275828/is-there-a-way-to-check-if-android-device-screen-is-locked-via-adb/60037241#60037241
	adbText := runAdbCommand("shell", "dumpsys", "window")

	return isOn(adbText)
}

func isOn(input string) bool {
	r, err := regexp.Compile("mDreamingLockscreen=([^\\s]+)")
	HandleErr(err)

	if !r.MatchString(input) {
		HandleErr(errors.New("regex failed"))
	}

	m := r.FindStringSubmatch(input)

	// mDreamingLockscreen=false => screen is on
	// mDreamingLockscreen=true => screen is off
	if m[1] == "true" {
		return false
	}

	return true
}

func GetScreenResolution() ScreenResolution {
	adbText := runAdbCommand("shell", "dumpsys", "display")

	return getScreenResolution(adbText)
}

func getScreenResolution(input string) ScreenResolution {
	r, err := regexp.Compile("mStableDisplaySize=Point\\((\\d+),\\s+(\\d+)\\)")
	HandleErr(err)

	if !r.MatchString(input) {
		HandleErr(errors.New("regex failed"))
	}

	m := r.FindStringSubmatch(input)

	w, err := strconv.Atoi(m[1])
	HandleErr(err)
	h, err := strconv.Atoi(m[2])
	HandleErr(err)

	return ScreenResolution{
		Width:  w,
		Height: h,
	}
}
