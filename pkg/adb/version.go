package adb

import (
	"strings"
)

func GetVersion() string {
	adbCommandOutput := runAdbCommand("version")

	return strings.TrimSpace(adbCommandOutput)
}
