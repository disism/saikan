// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/disism/saikan/ent/albums"
	"github.com/disism/saikan/ent/images"
)

// Albums is the model entity for the Albums schema.
type Albums struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Date holds the value of the "date" field.
	Date string `json:"date,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AlbumsQuery when eager-loading is set.
	Edges         AlbumsEdges `json:"edges"`
	images_albums *uint64
	selectValues  sql.SelectValues
}

// AlbumsEdges holds the relations/edges for other nodes in the graph.
type AlbumsEdges struct {
	// Image holds the value of the image edge.
	Image *Images `json:"image,omitempty"`
	// Musics holds the value of the musics edge.
	Musics []*Musics `json:"musics,omitempty"`
	// Users holds the value of the users edge.
	Users []*Users `json:"users,omitempty"`
	// Artists holds the value of the artists edge.
	Artists []*Artists `json:"artists,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// ImageOrErr returns the Image value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AlbumsEdges) ImageOrErr() (*Images, error) {
	if e.loadedTypes[0] {
		if e.Image == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: images.Label}
		}
		return e.Image, nil
	}
	return nil, &NotLoadedError{edge: "image"}
}

// MusicsOrErr returns the Musics value or an error if the edge
// was not loaded in eager-loading.
func (e AlbumsEdges) MusicsOrErr() ([]*Musics, error) {
	if e.loadedTypes[1] {
		return e.Musics, nil
	}
	return nil, &NotLoadedError{edge: "musics"}
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e AlbumsEdges) UsersOrErr() ([]*Users, error) {
	if e.loadedTypes[2] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// ArtistsOrErr returns the Artists value or an error if the edge
// was not loaded in eager-loading.
func (e AlbumsEdges) ArtistsOrErr() ([]*Artists, error) {
	if e.loadedTypes[3] {
		return e.Artists, nil
	}
	return nil, &NotLoadedError{edge: "artists"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Albums) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case albums.FieldID:
			values[i] = new(sql.NullInt64)
		case albums.FieldTitle, albums.FieldDate, albums.FieldDescription:
			values[i] = new(sql.NullString)
		case albums.FieldCreateTime, albums.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case albums.ForeignKeys[0]: // images_albums
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Albums fields.
func (a *Albums) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case albums.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = uint64(value.Int64)
		case albums.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				a.CreateTime = value.Time
			}
		case albums.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				a.UpdateTime = value.Time
			}
		case albums.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				a.Title = value.String
			}
		case albums.FieldDate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field date", values[i])
			} else if value.Valid {
				a.Date = value.String
			}
		case albums.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				a.Description = value.String
			}
		case albums.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field images_albums", value)
			} else if value.Valid {
				a.images_albums = new(uint64)
				*a.images_albums = uint64(value.Int64)
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Albums.
// This includes values selected through modifiers, order, etc.
func (a *Albums) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryImage queries the "image" edge of the Albums entity.
func (a *Albums) QueryImage() *ImagesQuery {
	return NewAlbumsClient(a.config).QueryImage(a)
}

// QueryMusics queries the "musics" edge of the Albums entity.
func (a *Albums) QueryMusics() *MusicsQuery {
	return NewAlbumsClient(a.config).QueryMusics(a)
}

// QueryUsers queries the "users" edge of the Albums entity.
func (a *Albums) QueryUsers() *UsersQuery {
	return NewAlbumsClient(a.config).QueryUsers(a)
}

// QueryArtists queries the "artists" edge of the Albums entity.
func (a *Albums) QueryArtists() *ArtistsQuery {
	return NewAlbumsClient(a.config).QueryArtists(a)
}

// Update returns a builder for updating this Albums.
// Note that you need to call Albums.Unwrap() before calling this method if this Albums
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Albums) Update() *AlbumsUpdateOne {
	return NewAlbumsClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Albums entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Albums) Unwrap() *Albums {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Albums is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Albums) String() string {
	var builder strings.Builder
	builder.WriteString("Albums(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("create_time=")
	builder.WriteString(a.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(a.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(a.Title)
	builder.WriteString(", ")
	builder.WriteString("date=")
	builder.WriteString(a.Date)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(a.Description)
	builder.WriteByte(')')
	return builder.String()
}

// AlbumsSlice is a parsable slice of Albums.
type AlbumsSlice []*Albums