package pageviews

import "ghdui/internal/widgets"

func Stat() widgets.Stat {
	return widgets.Stat{
		Title:       "Total Page Views",
		Value:       "89,400",
		Description: "21% more than last month",
	}
}
