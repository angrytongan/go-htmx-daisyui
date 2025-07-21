package people

import "slices"

type Person struct {
	Name     string
	Age      int
	Archived bool
}

var (
	people = []Person{
		{
			Name:     "Alice",
			Age:      32,
			Archived: false,
		},
		{
			Name:     "Bob",
			Age:      35,
			Archived: false,
		},
		{
			Name:     "Carol",
			Age:      37,
			Archived: false,
		},
		{
			Name:     "David",
			Age:      37,
			Archived: false,
		},
		{
			Name:     "Eve",
			Age:      37,
			Archived: false,
		},
		{
			Name:     "Frank",
			Age:      37,
			Archived: false,
		},
	}
)

func All() []Person {
	return people
}

// ArchiveToggle toggles the archive status of a Person, and returns the new
// status.
func ArchiveToggle(name string) bool {
	idx := slices.IndexFunc(people, func(p Person) bool {
		return p.Name == name
	})

	people[idx].Archived = !people[idx].Archived

	return people[idx].Archived
}
