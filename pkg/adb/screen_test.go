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

func TestGetScreenResolution(t *testing.T) {
	input := `DISPLAY MANAGER (dumpsys display)
  mOnlyCode=false
  mSafeMode=false
  mPendingTraversal=false
  mGlobalDisplayState=OFF
  mNextNonDefaultDisplayId=3
  mDefaultViewport=DisplayViewport{valid=true, displayId=0, uniqueId='null', orientation=0, logicalFrame=Rect(0, 0 - 1080, 2220), physicalFrame=Rect(0, 0 - 1440, 2960), deviceWidth=1440, deviceHeight=2960}
  mExternalTouchViewport=DisplayViewport{valid=false, displayId=0, uniqueId='null', orientation=0, logicalFrame=Rect(0, 0 - 0, 0), physicalFrame=Rect(0, 0 - 0, 0), deviceWidth=0, deviceHeight=0}
  mVirtualTouchViewports=[]
  mDefaultDisplayDefaultColorMode=0
  mSingleDisplayDemoMode=false
  mStableDisplaySize=Point(1440, 2960)

Display Adapters: size=5
`

	expected := ScreenResolution{
		Width:  1440,
		Height: 2960,
	}

	assert.Equal(t, expected, getScreenResolution(input))
}
