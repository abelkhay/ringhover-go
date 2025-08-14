package models

// Category table metadata.
type Category struct {
	ID   uint64 `db:"id"`
	Name string `db:"name"`
}
