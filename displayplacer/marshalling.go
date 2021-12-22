package displayplacer

import (
	"fmt"
	"strings"
)

func DisplayPlacerMonitorToMonitor(dpLines []string) (Display, error) {
	d := Display{}

	for _, l := range dpLines {
		parts := strings.Split(l, ":")
		key := strings.TrimSpace(parts[0])
		value := ""
		if len(parts) > 1 {
			value = strings.TrimSpace(parts[1])
		}
		switch key {
		case "Persistent screen id":
			d.PersistentScreenID = value
		case "Type":
			d.Type = value
		case "Origin":
			// Remove any note such as: `Origin: (0,0) - main display`
			originParts := strings.Split(value, " - ")
			d.Origin = strings.TrimSpace(originParts[0])
		case "Rotation":
			// Remove any note such as: Rotation: 0 - rotate internal screen example ...
			rotationParts := strings.Split(value, " - ")
			d.Rotation = strings.TrimSpace(rotationParts[0])
		case "Resolution":
			d.Resolution = value
		case "Hertz":
			d.Hertz = value
		case "Color Depth":
			d.ColorDepth = value
		case "Scaling":
			d.Scaling = value
		}
	}

	return d, nil
}

func DisplayToDisplayPlacerString(d Display) string {
	switch d.Hertz {
	case "N/A":
		return fmt.Sprintf("id:%s res:%s color_depth:%s scaling:%s origin:%s degree:%s",
			d.PersistentScreenID,
			d.Resolution,
			d.ColorDepth,
			d.Scaling,
			d.Origin,
			d.Rotation,
		)
	default:
		return fmt.Sprintf("id:%s res:%s hz:%s color_depth:%s scaling:%s origin:%s degree:%s",
			d.PersistentScreenID,
			d.Resolution,
			d.Hertz,
			d.ColorDepth,
			d.Scaling,
			d.Origin,
			d.Rotation,
		)
	}
}
