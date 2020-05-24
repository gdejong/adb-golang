package adb

import (
	"errors"
	"regexp"
)

func GetAdbSerialNumber() (string, error) {
	adbCommandOutput := runAdbCommand("devices")

	return findSerialNumber(adbCommandOutput)
}

func findSerialNumber(input string) (string, error) {
	// If no devices are connected, exit.
	if input == "List of devices attached\n\n" {
		return "", errors.New("no device connected, please make sure your device is connected")
	}

	r, err := regexp.Compile("attached\\s*([a-zA-Z0-9]+)\\s*device")
	HandleErr(err)

	if !r.MatchString(input) {
		return "", errors.New("a device is connected but could not determine its ID, maybe the device is locked")
	}

	m := r.FindStringSubmatch(input)

	return m[1], nil
}
