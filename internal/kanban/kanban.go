package kanban

import "sync"

type WorkflowID int
type ColumnID int
type NoteID int

type Workflow struct {
	Columns    []Column   // a workflow is a number of columns
	LastNoteID NoteID     // track the last note id, so we can create (unique) new ones
	mu         sync.Mutex // we're accessing shared state...
}

type Column struct {
	ID          ColumnID
	WIPLimit    int    // work-in-progress limit
	Description string // the name of the column
	Notes       []Note // our sticky notes
}

type Note struct {
	ID          NoteID
	Description string // words on this note
}
