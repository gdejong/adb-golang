package adb

import (
	"errors"
	"regexp"
)

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
