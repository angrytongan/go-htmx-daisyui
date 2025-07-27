package kanban_test

import (
	"ghdui/internal/kanban"
	"slices"
	"testing"
)

// TestAddColumn tests adding a column to a workflow.
func TestAddColumn(t *testing.T) {
	k := kanban.NewMemory()
	description := "Foobar"
	WIPLimit := 0

	err := k.AddColumn(WIPLimit, description)
	if err != nil {
		t.Fatalf("k.AddColumn(%d, %s): %v", WIPLimit, description, err)
	}

	// Check length.
	gotNumColumns := len(k.Columns)
	wantNumColumns := 5
	if gotNumColumns != wantNumColumns {
		t.Fatalf("len(k.Columns): got %d, want %d", gotNumColumns, wantNumColumns)
	}

	// Check IDs.
	gotColumnIDs := []kanban.ColumnID{}
	for _, v := range k.Columns {
		gotColumnIDs = append(gotColumnIDs, v.ID)
	}
	wantColumnIDs := []kanban.ColumnID{1, 2, 3, 4, 5}

	if !slices.Equal(gotColumnIDs, wantColumnIDs) {
		t.Fatalf("slices.Equal(): got %v, want %v", gotColumnIDs, wantColumnIDs)
	}
}

// TestDelColumn tests deleting a column from a workflow.
func TestDelColumn(t *testing.T) {
	k := kanban.NewMemory()
	var columnToDelete kanban.ColumnID = 2

	err := k.DelColumn(columnToDelete)
	if err != nil {
		t.Fatalf("k.DelColumn(%d): %v", columnToDelete, err)
	}

	// Check length.
	gotNumColumns := len(k.Columns)
	wantNumColumns := 3
	if gotNumColumns != wantNumColumns {
		t.Fatalf("len(k.Columns): got %d, want %d", gotNumColumns, wantNumColumns)
	}

	// Check IDs.
	gotColumnIDs := []kanban.ColumnID{}
	for _, v := range k.Columns {
		gotColumnIDs = append(gotColumnIDs, v.ID)
	}
	wantColumnIDs := []kanban.ColumnID{1, 3, 4}

	if !slices.Equal(gotColumnIDs, wantColumnIDs) {
		t.Fatalf("slices.Equal(): got %v, want %v", gotColumnIDs, wantColumnIDs)
	}
}

// TestOrderColumn tests column reordering.
func TestOrderColumn(t *testing.T) {
	k := kanban.NewMemory()

	k.OrderColumn(4, 0)
	k.OrderColumn(3, 0)

	// Check IDs.
	gotColumnIDs := []kanban.ColumnID{}
	for _, v := range k.Columns {
		gotColumnIDs = append(gotColumnIDs, v.ID)
	}
	wantColumnIDs := []kanban.ColumnID{3, 4, 1, 2}

	if !slices.Equal(gotColumnIDs, wantColumnIDs) {
		t.Fatalf("slices.Equal(): got %v, want %v", gotColumnIDs, wantColumnIDs)
	}
}

// TestAddNote tests that notes added to columns are added correctly.
func TestAddNote(t *testing.T) {
	k := kanban.NewMemory()

	adds := []struct {
		ColumnID    kanban.ColumnID
		Description string
	}{
		{1, "column 1, item 1"},
		{2, "column 2, item 1"},
		{3, "column 3, item 1"},
		{1, "column 1, item 2"},
		{2, "column 2, item 2"},
		{3, "column 3, item 2"},
	}

	for _, v := range adds {
		k.AddNote(v.ColumnID, v.Description)
	}

	// Check number of notes added to each column.
	numNotes := []struct {
		ColumnID kanban.ColumnID
		NumItems int
	}{{1, 2}, {2, 2}, {3, 2}}

	for _, v := range numNotes {
		idx := slices.IndexFunc(k.Columns, func(c kanban.Column) bool {
			return v.ColumnID == c.ID
		})

		if idx == -1 {
			t.Fatalf("slices.IndexFunc(%d): missing", v.ColumnID)
		}

		column := k.Columns[idx]
		gotNumNotes := len(column.Notes)
		wantNumNotes := v.NumItems

		if gotNumNotes != wantNumNotes {
			t.Fatalf(
				"number of notes for column %d (%s): got %d, want %d",
				v.ColumnID,
				column.Description,
				gotNumNotes,
				wantNumNotes,
			)
		}
	}

	// Check content of notes in each column is correct.
	// TODO
}
