package adb

import (
	"bytes"
	"os/exec"
	"strings"
)

func GetVersion() string {
	cmd := exec.Command("adb", "version")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	HandleErr(err)

	text := out.String()

	return strings.TrimSpace(text)
}
