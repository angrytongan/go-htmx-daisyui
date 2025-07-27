package kanban

type Kanban interface {
	NewMemory() *Workflow
	AddColumn(wipLimit int, description string) error
	AddNote(idColumn ColumnID, description string) error
	DelColumn(id ColumnID) error
	DelNote(idColumn ColumnID, idNote NoteID) error
	OrderColumn(id ColumnID, position int) error
}
