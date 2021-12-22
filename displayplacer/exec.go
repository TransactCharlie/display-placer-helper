package displayplacer

import (
	"bufio"
	"fmt"
	"os/exec"
)

func execDisplayPlacerSetDisplays(displays []Display) error {
	args := []string{}
	for _, d := range displays {
		args = append(args, DisplayToDisplayPlacerString(d))
	}
	cmd := exec.Command("displayplacer", args...)

	if err := cmd.Start(); err != nil {
		return err
	}

	return cmd.Wait()
}

// execDisplayPlacer
func execDisplayPlacerList() ([]string, error) {
	cmd := exec.Command("displayplacer", "list")

	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	// Read output
	buff := bufio.NewScanner(stdOut)
	var allText []string

	// We return the scutil output as a []string for each line
	// This blocks until scutil exits / We've scanned all the output
	for buff.Scan() {
		allText = append(allText, buff.Text())
	}
	if buff.Err() != nil {
		return nil, fmt.Errorf("error while scanning buffer: %w", buff.Err())
	}

	return allText, cmd.Wait()
}
