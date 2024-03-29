// Code generated by ent, DO NOT EDIT.

package images

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the images type in the database.
	Label = "images"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldWidth holds the string denoting the width field in the database.
	FieldWidth = "width"
	// FieldHeight holds the string denoting the height field in the database.
	FieldHeight = "height"
	// EdgeFile holds the string denoting the file edge name in mutations.
	EdgeFile = "file"
	// EdgeAlbums holds the string denoting the albums edge name in mutations.
	EdgeAlbums = "albums"
	// EdgePlaylists holds the string denoting the playlists edge name in mutations.
	EdgePlaylists = "playlists"
	// Table holds the table name of the images in the database.
	Table = "images"
	// FileTable is the table that holds the file relation/edge.
	FileTable = "images"
	// FileInverseTable is the table name for the Files entity.
	// It exists in this package in order to avoid circular dependency with the "files" package.
	FileInverseTable = "files"
	// FileColumn is the table column denoting the file relation/edge.
	FileColumn = "files_images"
	// AlbumsTable is the table that holds the albums relation/edge.
	AlbumsTable = "albums"
	// AlbumsInverseTable is the table name for the Albums entity.
	// It exists in this package in order to avoid circular dependency with the "albums" package.
	AlbumsInverseTable = "albums"
	// AlbumsColumn is the table column denoting the albums relation/edge.
	AlbumsColumn = "images_albums"
	// PlaylistsTable is the table that holds the playlists relation/edge.
	PlaylistsTable = "playlists"
	// PlaylistsInverseTable is the table name for the Playlists entity.
	// It exists in this package in order to avoid circular dependency with the "playlists" package.
	PlaylistsInverseTable = "playlists"
	// PlaylistsColumn is the table column denoting the playlists relation/edge.
	PlaylistsColumn = "images_playlists"
)

// Columns holds all SQL columns for images fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldWidth,
	FieldHeight,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "images"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"files_images",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
)

// OrderOption defines the ordering options for the Images queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByWidth orders the results by the width field.
func ByWidth(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWidth, opts...).ToFunc()
}

// ByHeight orders the results by the height field.
func ByHeight(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHeight, opts...).ToFunc()
}

// ByFileField orders the results by file field.
func ByFileField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFileStep(), sql.OrderByField(field, opts...))
	}
}

// ByAlbumsCount orders the results by albums count.
func ByAlbumsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAlbumsStep(), opts...)
	}
}

// ByAlbums orders the results by albums terms.
func ByAlbums(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAlbumsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPlaylistsCount orders the results by playlists count.
func ByPlaylistsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPlaylistsStep(), opts...)
	}
}

// ByPlaylists orders the results by playlists terms.
func ByPlaylists(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPlaylistsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newFileStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FileInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, FileTable, FileColumn),
	)
}
func newAlbumsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AlbumsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AlbumsTable, AlbumsColumn),
	)
}
func newPlaylistsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PlaylistsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PlaylistsTable, PlaylistsColumn),
	)
}
