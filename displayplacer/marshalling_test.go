package displayplacer

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DisplayToDisplayPlacerString(t *testing.T) {
	payload := Display{
		PersistentScreenID: "C3C2CD03-9FB9-45BE-DFBF-9F5A5CA823FF",
		Type:               "24 inch external screen",
		Resolution:         "2560x1440",
		Hertz:              "59",
		ColorDepth:         "8",
		Scaling:            "off",
		Origin:             "(0,-1440)",
		Rotation:           "0",
	}
	actual := DisplayToDisplayPlacerString(payload)
	expected := "id:C3C2CD03-9FB9-45BE-DFBF-9F5A5CA823FF res:2560x1440 hz:59 color_depth:8 scaling:off origin:(0,-1440) degree:0"
	assert.Equal(t, expected, actual)
}

func Test_DisplayPlacerDisplayToDisplay(t *testing.T) {
	payload := `
Persistent screen id: C3C2CD03-9FB9-45BE-DFBF-9F5A5CA823FF
Contextual screen id: 731409289
Type: 24 inch external screen
Resolution: 2560x1440
Hertz: 59
Color Depth: 8
Scaling:off
Origin: (0,-1440)
Rotation: 0
Resolutions for rotation 0:
  mode 0: res:2560x1440 hz:59 color_depth:4
  mode 1: res:2560x1440 hz:59 color_depth:8 <-- current mode
  mode 2: res:1024x768 hz:60 color_depth:4
  mode 3: res:1024x768 hz:60 color_depth:8
  mode 4: res:800x600 hz:60 color_depth:4
  mode 5: res:800x600 hz:60 color_depth:8
  mode 6: res:640x480 hz:59 color_depth:4
  mode 7: res:640x480 hz:59 color_depth:8
  mode 8: res:1280x1024 hz:60 color_depth:4
  mode 9: res:1280x1024 hz:60 color_depth:8
  mode 10: res:1920x1080 hz:60 color_depth:4
  mode 11: res:1920x1080 hz:60 color_depth:8
  mode 12: res:640x480 hz:59 color_depth:4
  mode 13: res:640x480 hz:59 color_depth:8
  mode 14: res:720x480 hz:59 color_depth:4
  mode 15: res:720x480 hz:59 color_depth:8
  mode 16: res:1280x720 hz:60 color_depth:4
  mode 17: res:1280x720 hz:60 color_depth:8
  mode 18: res:1920x1080 hz:60 color_depth:4
  mode 19: res:1920x1080 hz:60 color_depth:8
  mode 20: res:1920x1080 hz:50 color_depth:4
  mode 21: res:1920x1080 hz:50 color_depth:8
  mode 22: res:1280x720 hz:50 color_depth:4
  mode 23: res:1280x720 hz:50 color_depth:8
  mode 24: res:1920x1080 hz:60 color_depth:4
  mode 25: res:1920x1080 hz:60 color_depth:8
  mode 26: res:1920x1080 hz:50 color_depth:4
  mode 27: res:1920x1080 hz:50 color_depth:8
  mode 28: res:1280x720 hz:59 color_depth:4
  mode 29: res:1280x720 hz:59 color_depth:8
  mode 30: res:800x600 hz:59 color_depth:4
  mode 31: res:800x600 hz:59 color_depth:8
  mode 32: res:1024x768 hz:59 color_depth:4
  mode 33: res:1024x768 hz:59 color_depth:8
  mode 34: res:1024x576 hz:59 color_depth:4
  mode 35: res:1024x576 hz:59 color_depth:8
  mode 36: res:1280x960 hz:59 color_depth:4
  mode 37: res:1280x960 hz:59 color_depth:8
  mode 38: res:1344x1008 hz:59 color_depth:4
  mode 39: res:1344x1008 hz:59 color_depth:8
  mode 40: res:1344x756 hz:59 color_depth:4
  mode 41: res:1344x756 hz:59 color_depth:8
  mode 42: res:1440x900 hz:59 color_depth:4
  mode 43: res:1440x900 hz:59 color_depth:8
  mode 44: res:1680x1050 hz:59 color_depth:4
  mode 45: res:1680x1050 hz:59 color_depth:8
  mode 46: res:1600x1200 hz:59 color_depth:4
  mode 47: res:1600x1200 hz:59 color_depth:8
  mode 48: res:1600x900 hz:59 color_depth:4
  mode 49: res:1600x900 hz:59 color_depth:8
  mode 50: res:1920x1080 hz:59 color_depth:4
  mode 51: res:1920x1080 hz:59 color_depth:8
  mode 52: res:1920x1200 hz:59 color_depth:4
  mode 53: res:1920x1200 hz:59 color_depth:8
  mode 54: res:2048x1152 hz:59 color_depth:4
  mode 55: res:2048x1152 hz:59 color_depth:8
  mode 56: res:2560x1600 hz:59 color_depth:4
  mode 57: res:2560x1600 hz:59 color_depth:8
  mode 58: res:3840x2400 hz:59 color_depth:4
  mode 59: res:3840x2400 hz:59 color_depth:8
  mode 60: res:3360x2100 hz:59 color_depth:4
  mode 61: res:3360x2100 hz:59 color_depth:8
  mode 62: res:2880x1800 hz:59 color_depth:4
  mode 63: res:2880x1800 hz:59 color_depth:8
  mode 64: res:2048x1280 hz:59 color_depth:4
  mode 65: res:2048x1280 hz:59 color_depth:8
  mode 66: res:1280x720 hz:59 color_depth:4 scaling:on
  mode 67: res:1280x720 hz:59 color_depth:8 scaling:on
  mode 68: res:960x540 hz:60 color_depth:4 scaling:on
  mode 69: res:960x540 hz:60 color_depth:8 scaling:on
  mode 70: res:960x540 hz:60 color_depth:4 scaling:on
  mode 71: res:960x540 hz:60 color_depth:8 scaling:on
  mode 72: res:960x540 hz:50 color_depth:4 scaling:on
  mode 73: res:960x540 hz:50 color_depth:8 scaling:on
  mode 74: res:960x540 hz:60 color_depth:4 scaling:on
  mode 75: res:960x540 hz:60 color_depth:8 scaling:on
  mode 76: res:960x540 hz:50 color_depth:4 scaling:on
  mode 77: res:960x540 hz:50 color_depth:8 scaling:on
  mode 78: res:720x450 hz:59 color_depth:4 scaling:on
  mode 79: res:720x450 hz:59 color_depth:8 scaling:on
  mode 80: res:840x525 hz:59 color_depth:4 scaling:on
  mode 81: res:840x525 hz:59 color_depth:8 scaling:on
  mode 82: res:800x600 hz:59 color_depth:4 scaling:on
  mode 83: res:800x600 hz:59 color_depth:8 scaling:on
  mode 84: res:800x450 hz:59 color_depth:4 scaling:on
  mode 85: res:800x450 hz:59 color_depth:8 scaling:on
  mode 86: res:960x540 hz:59 color_depth:4 scaling:on
  mode 87: res:960x540 hz:59 color_depth:8 scaling:on
  mode 88: res:960x600 hz:59 color_depth:4 scaling:on
  mode 89: res:960x600 hz:59 color_depth:8 scaling:on
  mode 90: res:1024x576 hz:59 color_depth:4 scaling:on
  mode 91: res:1024x576 hz:59 color_depth:8 scaling:on
  mode 92: res:1280x800 hz:59 color_depth:4 scaling:on
  mode 93: res:1280x800 hz:59 color_depth:8 scaling:on
  mode 94: res:1920x1200 hz:59 color_depth:4 scaling:on
  mode 95: res:1920x1200 hz:59 color_depth:8 scaling:on
  mode 96: res:1680x1050 hz:59 color_depth:4 scaling:on
  mode 97: res:1680x1050 hz:59 color_depth:8 scaling:on
  mode 98: res:1440x900 hz:59 color_depth:4 scaling:on
  mode 99: res:1440x900 hz:59 color_depth:8 scaling:on
  mode 100: res:1024x640 hz:59 color_depth:4 scaling:on
  mode 101: res:1024x640 hz:59 color_depth:8 scaling:on
`
	lines := strings.Split(payload, "\n")
	actual, err := DisplayPlacerMonitorToMonitor(lines)
	assert.NoError(t, err)
	assert.NotNil(t, actual)

	expected := Display{
		PersistentScreenID: "C3C2CD03-9FB9-45BE-DFBF-9F5A5CA823FF",
		Type:               "24 inch external screen",
		Resolution:         "2560x1440",
		Hertz:              "59",
		ColorDepth:         "8",
		Origin:             "(0,-1440)",
		Rotation:           "0",
		Scaling:            "off",
	}
	assert.EqualValues(t, expected, actual)
}
