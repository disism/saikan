// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/disism/saikan/ent/files"
	"github.com/disism/saikan/ent/images"
	"github.com/disism/saikan/ent/musics"
)

// FilesCreate is the builder for creating a Files entity.
type FilesCreate struct {
	config
	mutation *FilesMutation
	hooks    []Hook
}

// SetHash sets the "hash" field.
func (fc *FilesCreate) SetHash(s string) *FilesCreate {
	fc.mutation.SetHash(s)
	return fc
}

// SetName sets the "name" field.
func (fc *FilesCreate) SetName(s string) *FilesCreate {
	fc.mutation.SetName(s)
	return fc
}

// SetSize sets the "size" field.
func (fc *FilesCreate) SetSize(u uint64) *FilesCreate {
	fc.mutation.SetSize(u)
	return fc
}

// SetID sets the "id" field.
func (fc *FilesCreate) SetID(u uint64) *FilesCreate {
	fc.mutation.SetID(u)
	return fc
}

// AddMusicIDs adds the "musics" edge to the Musics entity by IDs.
func (fc *FilesCreate) AddMusicIDs(ids ...uint64) *FilesCreate {
	fc.mutation.AddMusicIDs(ids...)
	return fc
}

// AddMusics adds the "musics" edges to the Musics entity.
func (fc *FilesCreate) AddMusics(m ...*Musics) *FilesCreate {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return fc.AddMusicIDs(ids...)
}

// AddImageIDs adds the "images" edge to the Images entity by IDs.
func (fc *FilesCreate) AddImageIDs(ids ...uint64) *FilesCreate {
	fc.mutation.AddImageIDs(ids...)
	return fc
}

// AddImages adds the "images" edges to the Images entity.
func (fc *FilesCreate) AddImages(i ...*Images) *FilesCreate {
	ids := make([]uint64, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return fc.AddImageIDs(ids...)
}

// Mutation returns the FilesMutation object of the builder.
func (fc *FilesCreate) Mutation() *FilesMutation {
	return fc.mutation
}

// Save creates the Files in the database.
func (fc *FilesCreate) Save(ctx context.Context) (*Files, error) {
	return withHooks(ctx, fc.sqlSave, fc.mutation, fc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FilesCreate) SaveX(ctx context.Context) *Files {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FilesCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FilesCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FilesCreate) check() error {
	if _, ok := fc.mutation.Hash(); !ok {
		return &ValidationError{Name: "hash", err: errors.New(`ent: missing required field "Files.hash"`)}
	}
	if _, ok := fc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Files.name"`)}
	}
	if v, ok := fc.mutation.Name(); ok {
		if err := files.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Files.name": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New(`ent: missing required field "Files.size"`)}
	}
	return nil
}

func (fc *FilesCreate) sqlSave(ctx context.Context) (*Files, error) {
	if err := fc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	fc.mutation.id = &_node.ID
	fc.mutation.done = true
	return _node, nil
}

func (fc *FilesCreate) createSpec() (*Files, *sqlgraph.CreateSpec) {
	var (
		_node = &Files{config: fc.config}
		_spec = sqlgraph.NewCreateSpec(files.Table, sqlgraph.NewFieldSpec(files.FieldID, field.TypeUint64))
	)
	if id, ok := fc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := fc.mutation.Hash(); ok {
		_spec.SetField(files.FieldHash, field.TypeString, value)
		_node.Hash = value
	}
	if value, ok := fc.mutation.Name(); ok {
		_spec.SetField(files.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := fc.mutation.Size(); ok {
		_spec.SetField(files.FieldSize, field.TypeUint64, value)
		_node.Size = value
	}
	if nodes := fc.mutation.MusicsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   files.MusicsTable,
			Columns: []string{files.MusicsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(musics.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.ImagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   files.ImagesTable,
			Columns: []string{files.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(images.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FilesCreateBulk is the builder for creating many Files entities in bulk.
type FilesCreateBulk struct {
	config
	err      error
	builders []*FilesCreate
}

// Save creates the Files entities in the database.
func (fcb *FilesCreateBulk) Save(ctx context.Context) ([]*Files, error) {
	if fcb.err != nil {
		return nil, fcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Files, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FilesMutation)
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
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FilesCreateBulk) SaveX(ctx context.Context) []*Files {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FilesCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FilesCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}
