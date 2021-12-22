package displayplacer

import "strings"

func rawListOutputToDisplaySetLines(rawLines []string) [][]string {

	displays := [][]string{}

	spaceIndex := []int{}

	// find the parts
	for ix, line := range rawLines {
		if line == "" {
			spaceIndex = append(spaceIndex, ix)
		}
	}

	// now nibble the parts starting at 0
	// We will not attempt to process the last chunk of the rawLines output as this never is a display.
	index := 0

	for _, nextSpaceIndex := range spaceIndex {
		// All the display sections we are interested in
		// start with Persistent screen id
		if strings.HasPrefix(rawLines[index], "Persistent screen id:") {
			displays = append(displays, rawLines[index:nextSpaceIndex])
		}

		// set index to start of next chunk
		index = nextSpaceIndex + 1
	}

	return displays
}

func rawListOutputToDisplays(rawLines []string) ([]Display, error) {
	r := []Display{}
	displayListLines := rawListOutputToDisplaySetLines(rawLines)
	for _, displayLines := range displayListLines {
		display, err := DisplayPlacerMonitorToMonitor(displayLines)
		if err != nil {
			return nil, err
		}
		r = append(r, display)
	}

	return r, nil
}
