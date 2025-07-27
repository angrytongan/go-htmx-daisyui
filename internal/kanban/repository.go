package kanban

type Kanban interface {
	NewMemory() *Workflow
	AddColumn(wipLimit int, description string) error
	DelColumn(id ColumnID) error
	OrderColumn(id ColumnID, position int) error
}
