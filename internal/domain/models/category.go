package models

// Category table metadata.
// TODO
type Category struct {
	ID   uint64 `db:"id"`
	Name string `db:"name"`
}
