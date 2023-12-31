// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/disism/saikan/ent/album"
	"github.com/disism/saikan/ent/file"
	"github.com/disism/saikan/ent/image"
	"github.com/disism/saikan/ent/playlist"
)

// ImageCreate is the builder for creating a Image entity.
type ImageCreate struct {
	config
	mutation *ImageMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (ic *ImageCreate) SetCreateTime(t time.Time) *ImageCreate {
	ic.mutation.SetCreateTime(t)
	return ic
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ic *ImageCreate) SetNillableCreateTime(t *time.Time) *ImageCreate {
	if t != nil {
		ic.SetCreateTime(*t)
	}
	return ic
}

// SetUpdateTime sets the "update_time" field.
func (ic *ImageCreate) SetUpdateTime(t time.Time) *ImageCreate {
	ic.mutation.SetUpdateTime(t)
	return ic
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (ic *ImageCreate) SetNillableUpdateTime(t *time.Time) *ImageCreate {
	if t != nil {
		ic.SetUpdateTime(*t)
	}
	return ic
}

// SetWidth sets the "width" field.
func (ic *ImageCreate) SetWidth(i int32) *ImageCreate {
	ic.mutation.SetWidth(i)
	return ic
}

// SetHeight sets the "height" field.
func (ic *ImageCreate) SetHeight(i int32) *ImageCreate {
	ic.mutation.SetHeight(i)
	return ic
}

// SetID sets the "id" field.
func (ic *ImageCreate) SetID(u uint64) *ImageCreate {
	ic.mutation.SetID(u)
	return ic
}

// SetFileID sets the "file" edge to the File entity by ID.
func (ic *ImageCreate) SetFileID(id uint64) *ImageCreate {
	ic.mutation.SetFileID(id)
	return ic
}

// SetFile sets the "file" edge to the File entity.
func (ic *ImageCreate) SetFile(f *File) *ImageCreate {
	return ic.SetFileID(f.ID)
}

// AddAlbumIDs adds the "albums" edge to the Album entity by IDs.
func (ic *ImageCreate) AddAlbumIDs(ids ...uint64) *ImageCreate {
	ic.mutation.AddAlbumIDs(ids...)
	return ic
}

// AddAlbums adds the "albums" edges to the Album entity.
func (ic *ImageCreate) AddAlbums(a ...*Album) *ImageCreate {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ic.AddAlbumIDs(ids...)
}

// AddPlaylistIDs adds the "playlists" edge to the Playlist entity by IDs.
func (ic *ImageCreate) AddPlaylistIDs(ids ...uint64) *ImageCreate {
	ic.mutation.AddPlaylistIDs(ids...)
	return ic
}

// AddPlaylists adds the "playlists" edges to the Playlist entity.
func (ic *ImageCreate) AddPlaylists(p ...*Playlist) *ImageCreate {
	ids := make([]uint64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ic.AddPlaylistIDs(ids...)
}

// Mutation returns the ImageMutation object of the builder.
func (ic *ImageCreate) Mutation() *ImageMutation {
	return ic.mutation
}

// Save creates the Image in the database.
func (ic *ImageCreate) Save(ctx context.Context) (*Image, error) {
	ic.defaults()
	return withHooks(ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *ImageCreate) SaveX(ctx context.Context) *Image {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *ImageCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *ImageCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *ImageCreate) defaults() {
	if _, ok := ic.mutation.CreateTime(); !ok {
		v := image.DefaultCreateTime()
		ic.mutation.SetCreateTime(v)
	}
	if _, ok := ic.mutation.UpdateTime(); !ok {
		v := image.DefaultUpdateTime()
		ic.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *ImageCreate) check() error {
	if _, ok := ic.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Image.create_time"`)}
	}
	if _, ok := ic.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Image.update_time"`)}
	}
	if _, ok := ic.mutation.Width(); !ok {
		return &ValidationError{Name: "width", err: errors.New(`ent: missing required field "Image.width"`)}
	}
	if _, ok := ic.mutation.Height(); !ok {
		return &ValidationError{Name: "height", err: errors.New(`ent: missing required field "Image.height"`)}
	}
	if _, ok := ic.mutation.FileID(); !ok {
		return &ValidationError{Name: "file", err: errors.New(`ent: missing required edge "Image.file"`)}
	}
	return nil
}

func (ic *ImageCreate) sqlSave(ctx context.Context) (*Image, error) {
	if err := ic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	ic.mutation.id = &_node.ID
	ic.mutation.done = true
	return _node, nil
}

func (ic *ImageCreate) createSpec() (*Image, *sqlgraph.CreateSpec) {
	var (
		_node = &Image{config: ic.config}
		_spec = sqlgraph.NewCreateSpec(image.Table, sqlgraph.NewFieldSpec(image.FieldID, field.TypeUint64))
	)
	if id, ok := ic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ic.mutation.CreateTime(); ok {
		_spec.SetField(image.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := ic.mutation.UpdateTime(); ok {
		_spec.SetField(image.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := ic.mutation.Width(); ok {
		_spec.SetField(image.FieldWidth, field.TypeInt32, value)
		_node.Width = value
	}
	if value, ok := ic.mutation.Height(); ok {
		_spec.SetField(image.FieldHeight, field.TypeInt32, value)
		_node.Height = value
	}
	if nodes := ic.mutation.FileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   image.FileTable,
			Columns: []string{image.FileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.file_images = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.AlbumsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   image.AlbumsTable,
			Columns: []string{image.AlbumsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(album.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.PlaylistsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   image.PlaylistsTable,
			Columns: []string{image.PlaylistsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playlist.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ImageCreateBulk is the builder for creating many Image entities in bulk.
type ImageCreateBulk struct {
	config
	err      error
	builders []*ImageCreate
}

// Save creates the Image entities in the database.
func (icb *ImageCreateBulk) Save(ctx context.Context) ([]*Image, error) {
	if icb.err != nil {
		return nil, icb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Image, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ImageMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *ImageCreateBulk) SaveX(ctx context.Context) []*Image {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *ImageCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *ImageCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}
