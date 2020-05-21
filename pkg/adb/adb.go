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

func runAdbCommand(command string) string {
	cmd := exec.Command("adb", command)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	HandleErr(err)

	return out.String()
}
