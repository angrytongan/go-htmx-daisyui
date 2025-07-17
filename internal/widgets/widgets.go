package widgets

// Stat.
// - https://daisyui.com/components/stat/
type Stat struct {
	Title       string
	Value       string
	Description string

	ValueClass       string
	DescriptionClass string
}

// Table.
// - https://daisyui.com/components/table/
type Table struct {
	Head []string
	Body [][]string
	Foot []string
}
