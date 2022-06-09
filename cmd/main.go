package main

import (
	"github.com/transactcharlie/display-placer-helper/displayplacer"
)

var myDisplays = []displayplacer.Display{
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
		Origin:             "(62,-1440)",
		Rotation:           "0",
	},
	{
		PersistentScreenID: "0BCE8D4D-DF1E-7F20-CBDD-064991103F22",
		Type:               "24 inch external screen",
		Resolution:         "2560x1440",
		Hertz:              "59",
		ColorDepth:         "8",
		Scaling:            "off",
		Origin:             "(-2498,-1440)",
		Rotation:           "0",
	},
	{
		PersistentScreenID: "6161706C-6950-6164-0000-053900000000",
		Type:               "19 inch external screen",
		Resolution:         "1112x834",
		Hertz:              "60",
		ColorDepth:         "4",
		Scaling:            "on",
		Origin:             "(-1112,0)",
		Rotation:           "0",
	},
}

func main() {

	displays, err := displayplacer.ListDisplays()
	if err != nil {
		panic(err)
	}

	displayMap := map[string]displayplacer.Display{}
	for _, d := range displays {
		displayMap[d.PersistentScreenID] = d
	}

	var attachedDisplays []displayplacer.Display
	requireChange := false

	for _, d := range myDisplays {
		compareDisplay, found := displayMap[d.PersistentScreenID]

		// If one of our displays isn't found that's ok -- it isn't plugged in
		if !found {
			continue
		}
		attachedDisplays = append(attachedDisplays, compareDisplay)

		if compareDisplay != d {
			requireChange = true
		}
	}

	if requireChange {
		err = displayplacer.ApplyDisplays(attachedDisplays)
		if err != nil {
			panic(err)
		}
	}
}
