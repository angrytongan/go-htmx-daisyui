package users

import "ghdui/internal/widgets"

func Stat() widgets.Stat {
	return widgets.Stat{
		Title:            "Users",
		Value:            "4,200",
		Description:      "↗︎ 40 (2%)",
		ValueClass:       "text-secondary",
		DescriptionClass: "text-secondary",
	}
}
