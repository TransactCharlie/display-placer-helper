package displayplacer

func ListDisplays() ([]Display, error) {
	rawLines, err := execDisplayPlacerList()
	if err != nil {
		return nil, err
	}

	displays, err := rawListOutputToDisplays(rawLines)
	if err != nil {
		return nil, err
	}
	return displays, nil
}

// ApplyDisplays asks displayplacer to set displays.
func ApplyDisplays(displays []Display) error {
	if len(displays) == 0 {
		return nil
	}
	return execDisplayPlacerSetDisplays(displays)
}
