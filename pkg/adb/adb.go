package adb

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"os/exec"
	"time"
)

func HandleErr(err error) {
	if err != nil {
		logrus.WithError(err).Panic("An error occurred!")
	}
}

func runAdbCommand(command ...string) string {
	commandBytes := runAdbCommandRawOutput(command...)

	return commandBytes.String()
}

func runAdbCommandRawOutput(command ...string) bytes.Buffer {
	cmd := exec.Command("adb", command...)
	var out bytes.Buffer
	cmd.Stdout = &out

	start := time.Now()
	err := cmd.Run()
	elapsed := time.Since(start)
	logrus.WithFields(logrus.Fields{"duration": elapsed, "adb command": command}).Debug("Ran ADB command")
	HandleErr(err)

	return out
}
