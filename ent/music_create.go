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
	"github.com/disism/saikan/ent/artist"
	"github.com/disism/saikan/ent/file"
	"github.com/disism/saikan/ent/music"
	"github.com/disism/saikan/ent/playlist"
	"github.com/disism/saikan/ent/user"
)

// MusicCreate is the builder for creating a Music entity.
type MusicCreate struct {
	config
	mutation *MusicMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (mc *MusicCreate) SetCreateTime(t time.Time) *MusicCreate {
	mc.mutation.SetCreateTime(t)
	return mc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (mc *MusicCreate) SetNillableCreateTime(t *time.Time) *MusicCreate {
	if t != nil {
		mc.SetCreateTime(*t)
	}
	return mc
}

// SetUpdateTime sets the "update_time" field.
func (mc *MusicCreate) SetUpdateTime(t time.Time) *MusicCreate {
	mc.mutation.SetUpdateTime(t)
	return mc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (mc *MusicCreate) SetNillableUpdateTime(t *time.Time) *MusicCreate {
	if t != nil {
		mc.SetUpdateTime(*t)
	}
	return mc
}

// SetName sets the "name" field.
func (mc *MusicCreate) SetName(s string) *MusicCreate {
	mc.mutation.SetName(s)
	return mc
}

// SetDescription sets the "description" field.
func (mc *MusicCreate) SetDescription(s string) *MusicCreate {
	mc.mutation.SetDescription(s)
	return mc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (mc *MusicCreate) SetNillableDescription(s *string) *MusicCreate {
	if s != nil {
		mc.SetDescription(*s)
	}
	return mc
}

// SetID sets the "id" field.
func (mc *MusicCreate) SetID(u uint64) *MusicCreate {
	mc.mutation.SetID(u)
	return mc
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (mc *MusicCreate) AddUserIDs(ids ...uint64) *MusicCreate {
	mc.mutation.AddUserIDs(ids...)
	return mc
}

// AddUser adds the "user" edges to the User entity.
func (mc *MusicCreate) AddUser(u ...*User) *MusicCreate {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return mc.AddUserIDs(ids...)
}

// SetFileID sets the "file" edge to the File entity by ID.
func (mc *MusicCreate) SetFileID(id uint64) *MusicCreate {
	mc.mutation.SetFileID(id)
	return mc
}

// SetFile sets the "file" edge to the File entity.
func (mc *MusicCreate) SetFile(f *File) *MusicCreate {
	return mc.SetFileID(f.ID)
}

// AddArtistIDs adds the "artists" edge to the Artist entity by IDs.
func (mc *MusicCreate) AddArtistIDs(ids ...uint64) *MusicCreate {
	mc.mutation.AddArtistIDs(ids...)
	return mc
}

// AddArtists adds the "artists" edges to the Artist entity.
func (mc *MusicCreate) AddArtists(a ...*Artist) *MusicCreate {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return mc.AddArtistIDs(ids...)
}

// AddPlaylistIDs adds the "playlists" edge to the Playlist entity by IDs.
func (mc *MusicCreate) AddPlaylistIDs(ids ...uint64) *MusicCreate {
	mc.mutation.AddPlaylistIDs(ids...)
	return mc
}

// AddPlaylists adds the "playlists" edges to the Playlist entity.
func (mc *MusicCreate) AddPlaylists(p ...*Playlist) *MusicCreate {
	ids := make([]uint64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mc.AddPlaylistIDs(ids...)
}

// AddAlbumIDs adds the "albums" edge to the Album entity by IDs.
func (mc *MusicCreate) AddAlbumIDs(ids ...uint64) *MusicCreate {
	mc.mutation.AddAlbumIDs(ids...)
	return mc
}

// AddAlbums adds the "albums" edges to the Album entity.
func (mc *MusicCreate) AddAlbums(a ...*Album) *MusicCreate {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return mc.AddAlbumIDs(ids...)
}

// Mutation returns the MusicMutation object of the builder.
func (mc *MusicCreate) Mutation() *MusicMutation {
	return mc.mutation
}

// Save creates the Music in the database.
func (mc *MusicCreate) Save(ctx context.Context) (*Music, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MusicCreate) SaveX(ctx context.Context) *Music {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MusicCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MusicCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MusicCreate) defaults() {
	if _, ok := mc.mutation.CreateTime(); !ok {
		v := music.DefaultCreateTime()
		mc.mutation.SetCreateTime(v)
	}
	if _, ok := mc.mutation.UpdateTime(); !ok {
		v := music.DefaultUpdateTime()
		mc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MusicCreate) check() error {
	if _, ok := mc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Music.create_time"`)}
	}
	if _, ok := mc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Music.update_time"`)}
	}
	if _, ok := mc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Music.name"`)}
	}
	if v, ok := mc.mutation.Description(); ok {
		if err := music.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Music.description": %w`, err)}
		}
	}
	if len(mc.mutation.UserIDs()) == 0 {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Music.user"`)}
	}
	if _, ok := mc.mutation.FileID(); !ok {
		return &ValidationError{Name: "file", err: errors.New(`ent: missing required edge "Music.file"`)}
	}
	return nil
}

func (mc *MusicCreate) sqlSave(ctx context.Context) (*Music, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MusicCreate) createSpec() (*Music, *sqlgraph.CreateSpec) {
	var (
		_node = &Music{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(music.Table, sqlgraph.NewFieldSpec(music.FieldID, field.TypeUint64))
	)
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.CreateTime(); ok {
		_spec.SetField(music.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := mc.mutation.UpdateTime(); ok {
		_spec.SetField(music.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := mc.mutation.Name(); ok {
		_spec.SetField(music.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := mc.mutation.Description(); ok {
		_spec.SetField(music.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := mc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   music.UserTable,
			Columns: music.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.FileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   music.FileTable,
			Columns: []string{music.FileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.file_musics = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.ArtistsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   music.ArtistsTable,
			Columns: music.ArtistsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artist.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.PlaylistsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   music.PlaylistsTable,
			Columns: music.PlaylistsPrimaryKey,
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
	if nodes := mc.mutation.AlbumsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   music.AlbumsTable,
			Columns: music.AlbumsPrimaryKey,
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
	return _node, _spec
}

// MusicCreateBulk is the builder for creating many Music entities in bulk.
type MusicCreateBulk struct {
	config
	err      error
	builders []*MusicCreate
}

// Save creates the Music entities in the database.
func (mcb *MusicCreateBulk) Save(ctx context.Context) ([]*Music, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Music, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MusicMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MusicCreateBulk) SaveX(ctx context.Context) []*Music {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MusicCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MusicCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
