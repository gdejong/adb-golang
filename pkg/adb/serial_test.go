package adb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindSerialNumber(t *testing.T) {
	serial, err := findSerialNumber(`List of devices attached
ce091829adc7dd2901      device

`)

	assert.Nil(t, err)
	assert.Equal(t, "ce091829adc7dd2901", serial)
}
