// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/disism/saikan/ent/albums"
	"github.com/disism/saikan/ent/artists"
	"github.com/disism/saikan/ent/musics"
	"github.com/disism/saikan/ent/predicate"
)

// ArtistsUpdate is the builder for updating Artists entities.
type ArtistsUpdate struct {
	config
	hooks    []Hook
	mutation *ArtistsMutation
}

// Where appends a list predicates to the ArtistsUpdate builder.
func (au *ArtistsUpdate) Where(ps ...predicate.Artists) *ArtistsUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetName sets the "name" field.
func (au *ArtistsUpdate) SetName(s string) *ArtistsUpdate {
	au.mutation.SetName(s)
	return au
}

// SetNillableName sets the "name" field if the given value is not nil.
func (au *ArtistsUpdate) SetNillableName(s *string) *ArtistsUpdate {
	if s != nil {
		au.SetName(*s)
	}
	return au
}

// AddMusicIDs adds the "musics" edge to the Musics entity by IDs.
func (au *ArtistsUpdate) AddMusicIDs(ids ...uint64) *ArtistsUpdate {
	au.mutation.AddMusicIDs(ids...)
	return au
}

// AddMusics adds the "musics" edges to the Musics entity.
func (au *ArtistsUpdate) AddMusics(m ...*Musics) *ArtistsUpdate {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return au.AddMusicIDs(ids...)
}

// AddAlbumIDs adds the "albums" edge to the Albums entity by IDs.
func (au *ArtistsUpdate) AddAlbumIDs(ids ...uint64) *ArtistsUpdate {
	au.mutation.AddAlbumIDs(ids...)
	return au
}

// AddAlbums adds the "albums" edges to the Albums entity.
func (au *ArtistsUpdate) AddAlbums(a ...*Albums) *ArtistsUpdate {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.AddAlbumIDs(ids...)
}

// Mutation returns the ArtistsMutation object of the builder.
func (au *ArtistsUpdate) Mutation() *ArtistsMutation {
	return au.mutation
}

// ClearMusics clears all "musics" edges to the Musics entity.
func (au *ArtistsUpdate) ClearMusics() *ArtistsUpdate {
	au.mutation.ClearMusics()
	return au
}

// RemoveMusicIDs removes the "musics" edge to Musics entities by IDs.
func (au *ArtistsUpdate) RemoveMusicIDs(ids ...uint64) *ArtistsUpdate {
	au.mutation.RemoveMusicIDs(ids...)
	return au
}

// RemoveMusics removes "musics" edges to Musics entities.
func (au *ArtistsUpdate) RemoveMusics(m ...*Musics) *ArtistsUpdate {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return au.RemoveMusicIDs(ids...)
}

// ClearAlbums clears all "albums" edges to the Albums entity.
func (au *ArtistsUpdate) ClearAlbums() *ArtistsUpdate {
	au.mutation.ClearAlbums()
	return au
}

// RemoveAlbumIDs removes the "albums" edge to Albums entities by IDs.
func (au *ArtistsUpdate) RemoveAlbumIDs(ids ...uint64) *ArtistsUpdate {
	au.mutation.RemoveAlbumIDs(ids...)
	return au
}

// RemoveAlbums removes "albums" edges to Albums entities.
func (au *ArtistsUpdate) RemoveAlbums(a ...*Albums) *ArtistsUpdate {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.RemoveAlbumIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ArtistsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *ArtistsUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ArtistsUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ArtistsUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *ArtistsUpdate) check() error {
	if v, ok := au.mutation.Name(); ok {
		if err := artists.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Artists.name": %w`, err)}
		}
	}
	return nil
}

func (au *ArtistsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(artists.Table, artists.Columns, sqlgraph.NewFieldSpec(artists.FieldID, field.TypeUint64))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(artists.FieldName, field.TypeString, value)
	}
	if au.mutation.MusicsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artists.MusicsTable,
			Columns: artists.MusicsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(musics.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedMusicsIDs(); len(nodes) > 0 && !au.mutation.MusicsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artists.MusicsTable,
			Columns: artists.MusicsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(musics.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.MusicsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artists.MusicsTable,
			Columns: artists.MusicsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(musics.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.AlbumsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artists.AlbumsTable,
			Columns: artists.AlbumsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(albums.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedAlbumsIDs(); len(nodes) > 0 && !au.mutation.AlbumsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artists.AlbumsTable,
			Columns: artists.AlbumsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(albums.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.AlbumsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artists.AlbumsTable,
			Columns: artists.AlbumsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(albums.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{artists.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// ArtistsUpdateOne is the builder for updating a single Artists entity.
type ArtistsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ArtistsMutation
}

// SetName sets the "name" field.
func (auo *ArtistsUpdateOne) SetName(s string) *ArtistsUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (auo *ArtistsUpdateOne) SetNillableName(s *string) *ArtistsUpdateOne {
	if s != nil {
		auo.SetName(*s)
	}
	return auo
}

// AddMusicIDs adds the "musics" edge to the Musics entity by IDs.
func (auo *ArtistsUpdateOne) AddMusicIDs(ids ...uint64) *ArtistsUpdateOne {
	auo.mutation.AddMusicIDs(ids...)
	return auo
}

// AddMusics adds the "musics" edges to the Musics entity.
func (auo *ArtistsUpdateOne) AddMusics(m ...*Musics) *ArtistsUpdateOne {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return auo.AddMusicIDs(ids...)
}

// AddAlbumIDs adds the "albums" edge to the Albums entity by IDs.
func (auo *ArtistsUpdateOne) AddAlbumIDs(ids ...uint64) *ArtistsUpdateOne {
	auo.mutation.AddAlbumIDs(ids...)
	return auo
}

// AddAlbums adds the "albums" edges to the Albums entity.
func (auo *ArtistsUpdateOne) AddAlbums(a ...*Albums) *ArtistsUpdateOne {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.AddAlbumIDs(ids...)
}

// Mutation returns the ArtistsMutation object of the builder.
func (auo *ArtistsUpdateOne) Mutation() *ArtistsMutation {
	return auo.mutation
}

// ClearMusics clears all "musics" edges to the Musics entity.
func (auo *ArtistsUpdateOne) ClearMusics() *ArtistsUpdateOne {
	auo.mutation.ClearMusics()
	return auo
}

// RemoveMusicIDs removes the "musics" edge to Musics entities by IDs.
func (auo *ArtistsUpdateOne) RemoveMusicIDs(ids ...uint64) *ArtistsUpdateOne {
	auo.mutation.RemoveMusicIDs(ids...)
	return auo
}

// RemoveMusics removes "musics" edges to Musics entities.
func (auo *ArtistsUpdateOne) RemoveMusics(m ...*Musics) *ArtistsUpdateOne {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return auo.RemoveMusicIDs(ids...)
}

// ClearAlbums clears all "albums" edges to the Albums entity.
func (auo *ArtistsUpdateOne) ClearAlbums() *ArtistsUpdateOne {
	auo.mutation.ClearAlbums()
	return auo
}

// RemoveAlbumIDs removes the "albums" edge to Albums entities by IDs.
func (auo *ArtistsUpdateOne) RemoveAlbumIDs(ids ...uint64) *ArtistsUpdateOne {
	auo.mutation.RemoveAlbumIDs(ids...)
	return auo
}

// RemoveAlbums removes "albums" edges to Albums entities.
func (auo *ArtistsUpdateOne) RemoveAlbums(a ...*Albums) *ArtistsUpdateOne {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.RemoveAlbumIDs(ids...)
}

// Where appends a list predicates to the ArtistsUpdate builder.
func (auo *ArtistsUpdateOne) Where(ps ...predicate.Artists) *ArtistsUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *ArtistsUpdateOne) Select(field string, fields ...string) *ArtistsUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Artists entity.
func (auo *ArtistsUpdateOne) Save(ctx context.Context) (*Artists, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ArtistsUpdateOne) SaveX(ctx context.Context) *Artists {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ArtistsUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ArtistsUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *ArtistsUpdateOne) check() error {
	if v, ok := auo.mutation.Name(); ok {
		if err := artists.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Artists.name": %w`, err)}
		}
	}
	return nil
}

func (auo *ArtistsUpdateOne) sqlSave(ctx context.Context) (_node *Artists, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(artists.Table, artists.Columns, sqlgraph.NewFieldSpec(artists.FieldID, field.TypeUint64))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Artists.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, artists.FieldID)
		for _, f := range fields {
			if !artists.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != artists.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.SetField(artists.FieldName, field.TypeString, value)
	}
	if auo.mutation.MusicsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artists.MusicsTable,
			Columns: artists.MusicsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(musics.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedMusicsIDs(); len(nodes) > 0 && !auo.mutation.MusicsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artists.MusicsTable,
			Columns: artists.MusicsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(musics.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.MusicsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artists.MusicsTable,
			Columns: artists.MusicsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(musics.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.AlbumsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artists.AlbumsTable,
			Columns: artists.AlbumsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(albums.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedAlbumsIDs(); len(nodes) > 0 && !auo.mutation.AlbumsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artists.AlbumsTable,
			Columns: artists.AlbumsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(albums.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.AlbumsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artists.AlbumsTable,
			Columns: artists.AlbumsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(albums.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Artists{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{artists.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
