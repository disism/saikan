// Code generated by ent, DO NOT EDIT.

package files

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the files type in the database.
	Label = "files"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHash holds the string denoting the hash field in the database.
	FieldHash = "hash"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSize holds the string denoting the size field in the database.
	FieldSize = "size"
	// EdgeMusics holds the string denoting the musics edge name in mutations.
	EdgeMusics = "musics"
	// EdgeImages holds the string denoting the images edge name in mutations.
	EdgeImages = "images"
	// Table holds the table name of the files in the database.
	Table = "files"
	// MusicsTable is the table that holds the musics relation/edge.
	MusicsTable = "musics"
	// MusicsInverseTable is the table name for the Musics entity.
	// It exists in this package in order to avoid circular dependency with the "musics" package.
	MusicsInverseTable = "musics"
	// MusicsColumn is the table column denoting the musics relation/edge.
	MusicsColumn = "files_musics"
	// ImagesTable is the table that holds the images relation/edge.
	ImagesTable = "images"
	// ImagesInverseTable is the table name for the Images entity.
	// It exists in this package in order to avoid circular dependency with the "images" package.
	ImagesInverseTable = "images"
	// ImagesColumn is the table column denoting the images relation/edge.
	ImagesColumn = "files_images"
)

// Columns holds all SQL columns for files fields.
var Columns = []string{
	FieldID,
	FieldHash,
	FieldName,
	FieldSize,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// OrderOption defines the ordering options for the Files queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByHash orders the results by the hash field.
func ByHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHash, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// BySize orders the results by the size field.
func BySize(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSize, opts...).ToFunc()
}

// ByMusicsCount orders the results by musics count.
func ByMusicsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMusicsStep(), opts...)
	}
}

// ByMusics orders the results by musics terms.
func ByMusics(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMusicsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByImagesCount orders the results by images count.
func ByImagesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newImagesStep(), opts...)
	}
}

// ByImages orders the results by images terms.
func ByImages(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newImagesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newMusicsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MusicsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MusicsTable, MusicsColumn),
	)
}
func newImagesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ImagesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ImagesTable, ImagesColumn),
	)
}