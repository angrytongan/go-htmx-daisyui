package kanban

import "slices"

func NewMemory() *Workflow {
	columns := []Column{
		{
			ID:          1,
			WIPLimit:    0, // no limit on this column
			Description: "Backlog",
			Notes:       []Note{},
		},
		{
			ID:          2,
			WIPLimit:    5,
			Description: "Planning",
			Notes:       []Note{},
		},
		{
			ID:          3,
			WIPLimit:    5,
			Description: "In progress",
			Notes:       []Note{},
		},
		{
			ID:          4,
			WIPLimit:    0, // no limit on this column
			Description: "Done",
			Notes:       []Note{},
		},
	}

	return &Workflow{
		Columns: columns,
	}
}

// AddColumn adds a new column to the board, in the last position.
func (wf *Workflow) AddColumn(WIPLimit int, description string) error {
	max := slices.MaxFunc(wf.Columns, func(a, b Column) int {
		if a.ID > b.ID {
			return int(a.ID)
		}
		return int(b.ID)
	})

	wf.Columns = append(wf.Columns, Column{
		ID:          max.ID + 1,
		WIPLimit:    WIPLimit,
		Description: description,
	})

	return nil
}

func (wf *Workflow) DelColumn(id ColumnID) error {
	idx := slices.IndexFunc(wf.Columns, func(c Column) bool {
		return c.ID == id
	})

	if idx != -1 {
		wf.Columns = slices.Delete(wf.Columns, idx, idx+1)
	}

	return nil
}

func (wf *Workflow) OrderColumn(id ColumnID, position int) error {
	idx := slices.IndexFunc(wf.Columns, func(c Column) bool {
		return c.ID == id
	})

	if idx != -1 {
		c := wf.Columns[idx]
		wf.Columns = slices.Delete(wf.Columns, idx, idx+1)
		wf.Columns = slices.Insert(wf.Columns, position, c)
	}

	return nil
}
