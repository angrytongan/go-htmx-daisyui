package employees

import "ghdui/internal/widgets"

func Table() widgets.Table {
	return widgets.Table{
		Head: []string{
			"ID", "Name", "Job", "Company", "Location", "Last Login", "Favourite Colour",
		},
		Body: [][]string{
			{
				"1",
				"Cy Ganderton",
				"Quality Control Specialist",
				"Littel, Schaden and Vandervort",
				"Canada",
				"12/16/2020",
				"Blue",
			},
			{
				"2",
				"Hart Hagerty",
				"Desktop Support Technician",
				"Zemlak, Daniel and Leannon",
				"United States",
				"12/5/2020",
				"Purple",
			},
		},
		Foot: []string{},
	}
}
