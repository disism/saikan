// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/disism/saikan/ent/albums"
	"github.com/disism/saikan/ent/files"
	"github.com/disism/saikan/ent/images"
	"github.com/disism/saikan/ent/playlists"
	"github.com/disism/saikan/ent/predicate"
)

// ImagesUpdate is the builder for updating Images entities.
type ImagesUpdate struct {
	config
	hooks    []Hook
	mutation *ImagesMutation
}

// Where appends a list predicates to the ImagesUpdate builder.
func (iu *ImagesUpdate) Where(ps ...predicate.Images) *ImagesUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetUpdateTime sets the "update_time" field.
func (iu *ImagesUpdate) SetUpdateTime(t time.Time) *ImagesUpdate {
	iu.mutation.SetUpdateTime(t)
	return iu
}

// SetWidth sets the "width" field.
func (iu *ImagesUpdate) SetWidth(i int32) *ImagesUpdate {
	iu.mutation.ResetWidth()
	iu.mutation.SetWidth(i)
	return iu
}

// SetNillableWidth sets the "width" field if the given value is not nil.
func (iu *ImagesUpdate) SetNillableWidth(i *int32) *ImagesUpdate {
	if i != nil {
		iu.SetWidth(*i)
	}
	return iu
}

// AddWidth adds i to the "width" field.
func (iu *ImagesUpdate) AddWidth(i int32) *ImagesUpdate {
	iu.mutation.AddWidth(i)
	return iu
}

// SetHeight sets the "height" field.
func (iu *ImagesUpdate) SetHeight(i int32) *ImagesUpdate {
	iu.mutation.ResetHeight()
	iu.mutation.SetHeight(i)
	return iu
}

// SetNillableHeight sets the "height" field if the given value is not nil.
func (iu *ImagesUpdate) SetNillableHeight(i *int32) *ImagesUpdate {
	if i != nil {
		iu.SetHeight(*i)
	}
	return iu
}

// AddHeight adds i to the "height" field.
func (iu *ImagesUpdate) AddHeight(i int32) *ImagesUpdate {
	iu.mutation.AddHeight(i)
	return iu
}

// SetFileID sets the "file" edge to the Files entity by ID.
func (iu *ImagesUpdate) SetFileID(id uint64) *ImagesUpdate {
	iu.mutation.SetFileID(id)
	return iu
}

// SetFile sets the "file" edge to the Files entity.
func (iu *ImagesUpdate) SetFile(f *Files) *ImagesUpdate {
	return iu.SetFileID(f.ID)
}

// AddAlbumIDs adds the "albums" edge to the Albums entity by IDs.
func (iu *ImagesUpdate) AddAlbumIDs(ids ...uint64) *ImagesUpdate {
	iu.mutation.AddAlbumIDs(ids...)
	return iu
}

// AddAlbums adds the "albums" edges to the Albums entity.
func (iu *ImagesUpdate) AddAlbums(a ...*Albums) *ImagesUpdate {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return iu.AddAlbumIDs(ids...)
}

// AddPlaylistIDs adds the "playlists" edge to the Playlists entity by IDs.
func (iu *ImagesUpdate) AddPlaylistIDs(ids ...uint64) *ImagesUpdate {
	iu.mutation.AddPlaylistIDs(ids...)
	return iu
}

// AddPlaylists adds the "playlists" edges to the Playlists entity.
func (iu *ImagesUpdate) AddPlaylists(p ...*Playlists) *ImagesUpdate {
	ids := make([]uint64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return iu.AddPlaylistIDs(ids...)
}

// Mutation returns the ImagesMutation object of the builder.
func (iu *ImagesUpdate) Mutation() *ImagesMutation {
	return iu.mutation
}

// ClearFile clears the "file" edge to the Files entity.
func (iu *ImagesUpdate) ClearFile() *ImagesUpdate {
	iu.mutation.ClearFile()
	return iu
}

// ClearAlbums clears all "albums" edges to the Albums entity.
func (iu *ImagesUpdate) ClearAlbums() *ImagesUpdate {
	iu.mutation.ClearAlbums()
	return iu
}

// RemoveAlbumIDs removes the "albums" edge to Albums entities by IDs.
func (iu *ImagesUpdate) RemoveAlbumIDs(ids ...uint64) *ImagesUpdate {
	iu.mutation.RemoveAlbumIDs(ids...)
	return iu
}

// RemoveAlbums removes "albums" edges to Albums entities.
func (iu *ImagesUpdate) RemoveAlbums(a ...*Albums) *ImagesUpdate {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return iu.RemoveAlbumIDs(ids...)
}

// ClearPlaylists clears all "playlists" edges to the Playlists entity.
func (iu *ImagesUpdate) ClearPlaylists() *ImagesUpdate {
	iu.mutation.ClearPlaylists()
	return iu
}

// RemovePlaylistIDs removes the "playlists" edge to Playlists entities by IDs.
func (iu *ImagesUpdate) RemovePlaylistIDs(ids ...uint64) *ImagesUpdate {
	iu.mutation.RemovePlaylistIDs(ids...)
	return iu
}

// RemovePlaylists removes "playlists" edges to Playlists entities.
func (iu *ImagesUpdate) RemovePlaylists(p ...*Playlists) *ImagesUpdate {
	ids := make([]uint64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return iu.RemovePlaylistIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *ImagesUpdate) Save(ctx context.Context) (int, error) {
	iu.defaults()
	return withHooks(ctx, iu.sqlSave, iu.mutation, iu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iu *ImagesUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *ImagesUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *ImagesUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iu *ImagesUpdate) defaults() {
	if _, ok := iu.mutation.UpdateTime(); !ok {
		v := images.UpdateDefaultUpdateTime()
		iu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *ImagesUpdate) check() error {
	if _, ok := iu.mutation.FileID(); iu.mutation.FileCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Images.file"`)
	}
	return nil
}

func (iu *ImagesUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := iu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(images.Table, images.Columns, sqlgraph.NewFieldSpec(images.FieldID, field.TypeUint64))
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.UpdateTime(); ok {
		_spec.SetField(images.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := iu.mutation.Width(); ok {
		_spec.SetField(images.FieldWidth, field.TypeInt32, value)
	}
	if value, ok := iu.mutation.AddedWidth(); ok {
		_spec.AddField(images.FieldWidth, field.TypeInt32, value)
	}
	if value, ok := iu.mutation.Height(); ok {
		_spec.SetField(images.FieldHeight, field.TypeInt32, value)
	}
	if value, ok := iu.mutation.AddedHeight(); ok {
		_spec.AddField(images.FieldHeight, field.TypeInt32, value)
	}
	if iu.mutation.FileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   images.FileTable,
			Columns: []string{images.FileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(files.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.FileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   images.FileTable,
			Columns: []string{images.FileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(files.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.AlbumsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   images.AlbumsTable,
			Columns: []string{images.AlbumsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(albums.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedAlbumsIDs(); len(nodes) > 0 && !iu.mutation.AlbumsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   images.AlbumsTable,
			Columns: []string{images.AlbumsColumn},
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
	if nodes := iu.mutation.AlbumsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   images.AlbumsTable,
			Columns: []string{images.AlbumsColumn},
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
	if iu.mutation.PlaylistsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   images.PlaylistsTable,
			Columns: []string{images.PlaylistsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playlists.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedPlaylistsIDs(); len(nodes) > 0 && !iu.mutation.PlaylistsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   images.PlaylistsTable,
			Columns: []string{images.PlaylistsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playlists.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.PlaylistsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   images.PlaylistsTable,
			Columns: []string{images.PlaylistsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playlists.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{images.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iu.mutation.done = true
	return n, nil
}

// ImagesUpdateOne is the builder for updating a single Images entity.
type ImagesUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ImagesMutation
}

// SetUpdateTime sets the "update_time" field.
func (iuo *ImagesUpdateOne) SetUpdateTime(t time.Time) *ImagesUpdateOne {
	iuo.mutation.SetUpdateTime(t)
	return iuo
}

// SetWidth sets the "width" field.
func (iuo *ImagesUpdateOne) SetWidth(i int32) *ImagesUpdateOne {
	iuo.mutation.ResetWidth()
	iuo.mutation.SetWidth(i)
	return iuo
}

// SetNillableWidth sets the "width" field if the given value is not nil.
func (iuo *ImagesUpdateOne) SetNillableWidth(i *int32) *ImagesUpdateOne {
	if i != nil {
		iuo.SetWidth(*i)
	}
	return iuo
}

// AddWidth adds i to the "width" field.
func (iuo *ImagesUpdateOne) AddWidth(i int32) *ImagesUpdateOne {
	iuo.mutation.AddWidth(i)
	return iuo
}

// SetHeight sets the "height" field.
func (iuo *ImagesUpdateOne) SetHeight(i int32) *ImagesUpdateOne {
	iuo.mutation.ResetHeight()
	iuo.mutation.SetHeight(i)
	return iuo
}

// SetNillableHeight sets the "height" field if the given value is not nil.
func (iuo *ImagesUpdateOne) SetNillableHeight(i *int32) *ImagesUpdateOne {
	if i != nil {
		iuo.SetHeight(*i)
	}
	return iuo
}

// AddHeight adds i to the "height" field.
func (iuo *ImagesUpdateOne) AddHeight(i int32) *ImagesUpdateOne {
	iuo.mutation.AddHeight(i)
	return iuo
}

// SetFileID sets the "file" edge to the Files entity by ID.
func (iuo *ImagesUpdateOne) SetFileID(id uint64) *ImagesUpdateOne {
	iuo.mutation.SetFileID(id)
	return iuo
}

// SetFile sets the "file" edge to the Files entity.
func (iuo *ImagesUpdateOne) SetFile(f *Files) *ImagesUpdateOne {
	return iuo.SetFileID(f.ID)
}

// AddAlbumIDs adds the "albums" edge to the Albums entity by IDs.
func (iuo *ImagesUpdateOne) AddAlbumIDs(ids ...uint64) *ImagesUpdateOne {
	iuo.mutation.AddAlbumIDs(ids...)
	return iuo
}

// AddAlbums adds the "albums" edges to the Albums entity.
func (iuo *ImagesUpdateOne) AddAlbums(a ...*Albums) *ImagesUpdateOne {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return iuo.AddAlbumIDs(ids...)
}

// AddPlaylistIDs adds the "playlists" edge to the Playlists entity by IDs.
func (iuo *ImagesUpdateOne) AddPlaylistIDs(ids ...uint64) *ImagesUpdateOne {
	iuo.mutation.AddPlaylistIDs(ids...)
	return iuo
}

// AddPlaylists adds the "playlists" edges to the Playlists entity.
func (iuo *ImagesUpdateOne) AddPlaylists(p ...*Playlists) *ImagesUpdateOne {
	ids := make([]uint64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return iuo.AddPlaylistIDs(ids...)
}

// Mutation returns the ImagesMutation object of the builder.
func (iuo *ImagesUpdateOne) Mutation() *ImagesMutation {
	return iuo.mutation
}

// ClearFile clears the "file" edge to the Files entity.
func (iuo *ImagesUpdateOne) ClearFile() *ImagesUpdateOne {
	iuo.mutation.ClearFile()
	return iuo
}

// ClearAlbums clears all "albums" edges to the Albums entity.
func (iuo *ImagesUpdateOne) ClearAlbums() *ImagesUpdateOne {
	iuo.mutation.ClearAlbums()
	return iuo
}

// RemoveAlbumIDs removes the "albums" edge to Albums entities by IDs.
func (iuo *ImagesUpdateOne) RemoveAlbumIDs(ids ...uint64) *ImagesUpdateOne {
	iuo.mutation.RemoveAlbumIDs(ids...)
	return iuo
}

// RemoveAlbums removes "albums" edges to Albums entities.
func (iuo *ImagesUpdateOne) RemoveAlbums(a ...*Albums) *ImagesUpdateOne {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return iuo.RemoveAlbumIDs(ids...)
}

// ClearPlaylists clears all "playlists" edges to the Playlists entity.
func (iuo *ImagesUpdateOne) ClearPlaylists() *ImagesUpdateOne {
	iuo.mutation.ClearPlaylists()
	return iuo
}

// RemovePlaylistIDs removes the "playlists" edge to Playlists entities by IDs.
func (iuo *ImagesUpdateOne) RemovePlaylistIDs(ids ...uint64) *ImagesUpdateOne {
	iuo.mutation.RemovePlaylistIDs(ids...)
	return iuo
}

// RemovePlaylists removes "playlists" edges to Playlists entities.
func (iuo *ImagesUpdateOne) RemovePlaylists(p ...*Playlists) *ImagesUpdateOne {
	ids := make([]uint64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return iuo.RemovePlaylistIDs(ids...)
}

// Where appends a list predicates to the ImagesUpdate builder.
func (iuo *ImagesUpdateOne) Where(ps ...predicate.Images) *ImagesUpdateOne {
	iuo.mutation.Where(ps...)
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *ImagesUpdateOne) Select(field string, fields ...string) *ImagesUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Images entity.
func (iuo *ImagesUpdateOne) Save(ctx context.Context) (*Images, error) {
	iuo.defaults()
	return withHooks(ctx, iuo.sqlSave, iuo.mutation, iuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *ImagesUpdateOne) SaveX(ctx context.Context) *Images {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *ImagesUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *ImagesUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iuo *ImagesUpdateOne) defaults() {
	if _, ok := iuo.mutation.UpdateTime(); !ok {
		v := images.UpdateDefaultUpdateTime()
		iuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *ImagesUpdateOne) check() error {
	if _, ok := iuo.mutation.FileID(); iuo.mutation.FileCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Images.file"`)
	}
	return nil
}

func (iuo *ImagesUpdateOne) sqlSave(ctx context.Context) (_node *Images, err error) {
	if err := iuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(images.Table, images.Columns, sqlgraph.NewFieldSpec(images.FieldID, field.TypeUint64))
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Images.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, images.FieldID)
		for _, f := range fields {
			if !images.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != images.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.UpdateTime(); ok {
		_spec.SetField(images.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := iuo.mutation.Width(); ok {
		_spec.SetField(images.FieldWidth, field.TypeInt32, value)
	}
	if value, ok := iuo.mutation.AddedWidth(); ok {
		_spec.AddField(images.FieldWidth, field.TypeInt32, value)
	}
	if value, ok := iuo.mutation.Height(); ok {
		_spec.SetField(images.FieldHeight, field.TypeInt32, value)
	}
	if value, ok := iuo.mutation.AddedHeight(); ok {
		_spec.AddField(images.FieldHeight, field.TypeInt32, value)
	}
	if iuo.mutation.FileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   images.FileTable,
			Columns: []string{images.FileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(files.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.FileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   images.FileTable,
			Columns: []string{images.FileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(files.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.AlbumsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   images.AlbumsTable,
			Columns: []string{images.AlbumsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(albums.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedAlbumsIDs(); len(nodes) > 0 && !iuo.mutation.AlbumsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   images.AlbumsTable,
			Columns: []string{images.AlbumsColumn},
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
	if nodes := iuo.mutation.AlbumsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   images.AlbumsTable,
			Columns: []string{images.AlbumsColumn},
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
	if iuo.mutation.PlaylistsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   images.PlaylistsTable,
			Columns: []string{images.PlaylistsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playlists.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedPlaylistsIDs(); len(nodes) > 0 && !iuo.mutation.PlaylistsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   images.PlaylistsTable,
			Columns: []string{images.PlaylistsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playlists.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.PlaylistsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   images.PlaylistsTable,
			Columns: []string{images.PlaylistsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playlists.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Images{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{images.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iuo.mutation.done = true
	return _node, nil
}