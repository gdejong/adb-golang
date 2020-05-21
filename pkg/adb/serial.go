package adb

import (
	"errors"
	"regexp"
)

func GetAdbSerialNumber() string {
	adbCommandOutput := runAdbCommand("devices")

	serial, err := findSerialNumber(adbCommandOutput)
	HandleErr(err)

	return serial
}

func findSerialNumber(input string) (string, error) {
	// If no phones are connected, exit.
	if input == "List of devices attached\n\n" {
		return "", errors.New("no device connected, please make sure your phone is connected")
	}

	r, err := regexp.Compile("attached\\s*([a-zA-Z0-9]+)\\s*device")
	HandleErr(err)

	if !r.MatchString(input) {
		return "", errors.New("a phone is connected but could not determine its ID, maybe the phone is locked")
	}

	m := r.FindStringSubmatch(input)

	return m[1], nil
}
