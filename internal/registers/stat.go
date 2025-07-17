package registers

import "ghdui/internal/widgets"

func Stat() widgets.Stat {
	return widgets.Stat{
		Title:       "Registers",
		Value:       "1,200",
		Description: "↘︎ 90 (14%)",
	}
}
