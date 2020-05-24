package adb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsOnReturnsTrue(t *testing.T) {
	on := isOn(`mDockLayer=268435456 mStatusBarLayer=0
    mShowingDream=false mDreamingLockscreen=false mDreamingSleepToken=null
    mStatusBar=Window{895aa8f u0 StatusBar} isStatusBarKeyguard=false`)

	assert.True(t, on)
}

func TestIsOnReturnsFalse(t *testing.T) {
	on := isOn(`mDockLayer=268435456 mStatusBarLayer=0
    mShowingDream=false mDreamingLockscreen=true mDreamingSleepToken=null
    mStatusBar=Window{895aa8f u0 StatusBar} isStatusBarKeyguard=false`)

	assert.False(t, on)
}
