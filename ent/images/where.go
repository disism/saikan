// Code generated by ent, DO NOT EDIT.

package images

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/disism/saikan/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Images {
	return predicate.Images(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Images {
	return predicate.Images(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Images {
	return predicate.Images(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Images {
	return predicate.Images(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Images {
	return predicate.Images(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Images {
	return predicate.Images(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Images {
	return predicate.Images(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Images {
	return predicate.Images(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Images {
	return predicate.Images(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Images {
	return predicate.Images(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Images {
	return predicate.Images(sql.FieldEQ(FieldUpdateTime, v))
}

// Width applies equality check predicate on the "width" field. It's identical to WidthEQ.
func Width(v int32) predicate.Images {
	return predicate.Images(sql.FieldEQ(FieldWidth, v))
}

// Height applies equality check predicate on the "height" field. It's identical to HeightEQ.
func Height(v int32) predicate.Images {
	return predicate.Images(sql.FieldEQ(FieldHeight, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Images {
	return predicate.Images(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Images {
	return predicate.Images(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Images {
	return predicate.Images(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Images {
	return predicate.Images(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Images {
	return predicate.Images(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Images {
	return predicate.Images(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Images {
	return predicate.Images(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Images {
	return predicate.Images(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Images {
	return predicate.Images(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Images {
	return predicate.Images(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Images {
	return predicate.Images(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Images {
	return predicate.Images(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Images {
	return predicate.Images(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Images {
	return predicate.Images(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Images {
	return predicate.Images(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Images {
	return predicate.Images(sql.FieldLTE(FieldUpdateTime, v))
}

// WidthEQ applies the EQ predicate on the "width" field.
func WidthEQ(v int32) predicate.Images {
	return predicate.Images(sql.FieldEQ(FieldWidth, v))
}

// WidthNEQ applies the NEQ predicate on the "width" field.
func WidthNEQ(v int32) predicate.Images {
	return predicate.Images(sql.FieldNEQ(FieldWidth, v))
}

// WidthIn applies the In predicate on the "width" field.
func WidthIn(vs ...int32) predicate.Images {
	return predicate.Images(sql.FieldIn(FieldWidth, vs...))
}

// WidthNotIn applies the NotIn predicate on the "width" field.
func WidthNotIn(vs ...int32) predicate.Images {
	return predicate.Images(sql.FieldNotIn(FieldWidth, vs...))
}

// WidthGT applies the GT predicate on the "width" field.
func WidthGT(v int32) predicate.Images {
	return predicate.Images(sql.FieldGT(FieldWidth, v))
}

// WidthGTE applies the GTE predicate on the "width" field.
func WidthGTE(v int32) predicate.Images {
	return predicate.Images(sql.FieldGTE(FieldWidth, v))
}

// WidthLT applies the LT predicate on the "width" field.
func WidthLT(v int32) predicate.Images {
	return predicate.Images(sql.FieldLT(FieldWidth, v))
}

// WidthLTE applies the LTE predicate on the "width" field.
func WidthLTE(v int32) predicate.Images {
	return predicate.Images(sql.FieldLTE(FieldWidth, v))
}

// HeightEQ applies the EQ predicate on the "height" field.
func HeightEQ(v int32) predicate.Images {
	return predicate.Images(sql.FieldEQ(FieldHeight, v))
}

// HeightNEQ applies the NEQ predicate on the "height" field.
func HeightNEQ(v int32) predicate.Images {
	return predicate.Images(sql.FieldNEQ(FieldHeight, v))
}

// HeightIn applies the In predicate on the "height" field.
func HeightIn(vs ...int32) predicate.Images {
	return predicate.Images(sql.FieldIn(FieldHeight, vs...))
}

// HeightNotIn applies the NotIn predicate on the "height" field.
func HeightNotIn(vs ...int32) predicate.Images {
	return predicate.Images(sql.FieldNotIn(FieldHeight, vs...))
}

// HeightGT applies the GT predicate on the "height" field.
func HeightGT(v int32) predicate.Images {
	return predicate.Images(sql.FieldGT(FieldHeight, v))
}

// HeightGTE applies the GTE predicate on the "height" field.
func HeightGTE(v int32) predicate.Images {
	return predicate.Images(sql.FieldGTE(FieldHeight, v))
}

// HeightLT applies the LT predicate on the "height" field.
func HeightLT(v int32) predicate.Images {
	return predicate.Images(sql.FieldLT(FieldHeight, v))
}

// HeightLTE applies the LTE predicate on the "height" field.
func HeightLTE(v int32) predicate.Images {
	return predicate.Images(sql.FieldLTE(FieldHeight, v))
}

// HasFile applies the HasEdge predicate on the "file" edge.
func HasFile() predicate.Images {
	return predicate.Images(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FileTable, FileColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFileWith applies the HasEdge predicate on the "file" edge with a given conditions (other predicates).
func HasFileWith(preds ...predicate.Files) predicate.Images {
	return predicate.Images(func(s *sql.Selector) {
		step := newFileStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAlbums applies the HasEdge predicate on the "albums" edge.
func HasAlbums() predicate.Images {
	return predicate.Images(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AlbumsTable, AlbumsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAlbumsWith applies the HasEdge predicate on the "albums" edge with a given conditions (other predicates).
func HasAlbumsWith(preds ...predicate.Albums) predicate.Images {
	return predicate.Images(func(s *sql.Selector) {
		step := newAlbumsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPlaylists applies the HasEdge predicate on the "playlists" edge.
func HasPlaylists() predicate.Images {
	return predicate.Images(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PlaylistsTable, PlaylistsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlaylistsWith applies the HasEdge predicate on the "playlists" edge with a given conditions (other predicates).
func HasPlaylistsWith(preds ...predicate.Playlists) predicate.Images {
	return predicate.Images(func(s *sql.Selector) {
		step := newPlaylistsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Images) predicate.Images {
	return predicate.Images(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Images) predicate.Images {
	return predicate.Images(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Images) predicate.Images {
	return predicate.Images(sql.NotPredicates(p))
}
