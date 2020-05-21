package adb

import (
	"bytes"
	"log"
	"os/exec"
)

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
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

	err := cmd.Run()
	HandleErr(err)

	return out
}
