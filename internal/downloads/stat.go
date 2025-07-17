package downloads

import "ghdui/internal/widgets"

func Stat() widgets.Stat {
	return widgets.Stat{
		Title:       "Downloads",
		Value:       "31K",
		Description: "Jan 1st - Feb 1st",
	}
}
