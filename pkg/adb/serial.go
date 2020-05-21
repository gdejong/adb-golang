package adb

import (
	"bytes"
	"log"
	"os/exec"
	"regexp"
)

func GetAdbSerialNumber() string {
	cmd := exec.Command("adb", "devices")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	HandleErr(err)
	text := out.String()

	// If no phones are connected, exit.
	if text == "List of devices attached\n\n" {
		log.Fatalln("No phone connected, please make sure your phone is connected.")
	}

	r, err := regexp.Compile("attached\\s*([a-zA-Z0-9]+)\\s*device")
	HandleErr(err)

	if !r.MatchString(text) {
		log.Fatalln("A phone is connected but could not determine its ID, maybe the phone is locked?")
	}

	m := r.FindStringSubmatch(text)

	return m[1]
}
