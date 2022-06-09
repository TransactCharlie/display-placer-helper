# display-placer-helper

A small utility that uses [displayplacer](https://github.com/jakehilborn/displayplacer) to maintain the position of your monitors on osx.

OSX is awful at maintining screen orientation of monitors. After a reboot monitors can be swapped about. Maybe this only happens if you have 2 external monitors of the exact same type but it's really annoying.

Displayplacer doesn't check to see if monitors are already in the same situation as they were when you ask it to set monitors so it always takes a few seconds and blanks all your screens. This makes it a bit awkward to call from a shell init for example.

## Usage

well... Not really general purpose usable but if you have a look at `cmd/main.go` we define the state of displays we wish to have and this util will check that the displays are in that state or not.

If they are not it will call displayplacer to make them be in that state.

```@golang

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
		Origin:             "(190,-1440)",
		Rotation:           "0",
	},
	{
		PersistentScreenID: "0BCE8D4D-DF1E-7F20-CBDD-064991103F22",
		Type:               "24 inch external screen",
		Resolution:         "2560x1440",
		Hertz:              "59",
		ColorDepth:         "8",
		Scaling:            "off",
		Origin:             "(-2370,-1440)",
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
```