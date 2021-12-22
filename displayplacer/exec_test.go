package displayplacer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_execDisplayPlacer(t *testing.T) {
	output, err := execDisplayPlacerList()
	assert.NoError(t, err)
	t.Log(output)
	for _, l := range output {
		t.Log(l)
	}
}

// Test_execDisplayPlacerSetDisplays is set up for my own three monitors and is probably
// not portable in any way.
func Test_execDisplayPlacerSetDisplays(t *testing.T) {
	displays := []Display{
		{
			PersistentScreenID: "1C5968B4-BED8-7273-7BFA-4A8911ED44BC",
			Type:               "MacBook built in screen",
			Resolution:         "1680x1050",
			Hertz:              "N/A",
			ColorDepth:         "8",
			Scaling:            "on",
			Origin:             "(0,0)",
			Rotation:           "0",
		},
		{
			PersistentScreenID: "C3C2CD03-9FB9-45BE-DFBF-9F5A5CA823FF",
			Type:               "24 inch external screen",
			Resolution:         "2560x1440",
			Hertz:              "59",
			ColorDepth:         "8",
			Scaling:            "off",
			Origin:             "(415,-1440)",
			Rotation:           "0",
		},
		{
			PersistentScreenID: "0BCE8D4D-DF1E-7F20-CBDD-064991103F22",
			Type:               "24 inch external screen",
			Resolution:         "2560x1440",
			Hertz:              "59",
			ColorDepth:         "8",
			Scaling:            "off",
			Origin:             "(-2145,-1440)",
			Rotation:           "0",
		},
	}
	err := execDisplayPlacerSetDisplays(displays)
	assert.NoError(t, err)
}
