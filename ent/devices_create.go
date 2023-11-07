// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/disism/saikan/ent/devices"
	"github.com/disism/saikan/ent/users"
)

// DevicesCreate is the builder for creating a Devices entity.
type DevicesCreate struct {
	config
	mutation *DevicesMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (dc *DevicesCreate) SetCreateTime(t time.Time) *DevicesCreate {
	dc.mutation.SetCreateTime(t)
	return dc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (dc *DevicesCreate) SetNillableCreateTime(t *time.Time) *DevicesCreate {
	if t != nil {
		dc.SetCreateTime(*t)
	}
	return dc
}

// SetUpdateTime sets the "update_time" field.
func (dc *DevicesCreate) SetUpdateTime(t time.Time) *DevicesCreate {
	dc.mutation.SetUpdateTime(t)
	return dc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (dc *DevicesCreate) SetNillableUpdateTime(t *time.Time) *DevicesCreate {
	if t != nil {
		dc.SetUpdateTime(*t)
	}
	return dc
}

// SetIP sets the "ip" field.
func (dc *DevicesCreate) SetIP(s string) *DevicesCreate {
	dc.mutation.SetIP(s)
	return dc
}

// SetDevice sets the "device" field.
func (dc *DevicesCreate) SetDevice(s string) *DevicesCreate {
	dc.mutation.SetDevice(s)
	return dc
}

// SetID sets the "id" field.
func (dc *DevicesCreate) SetID(u uint64) *DevicesCreate {
	dc.mutation.SetID(u)
	return dc
}

// SetUserID sets the "user" edge to the Users entity by ID.
func (dc *DevicesCreate) SetUserID(id uint64) *DevicesCreate {
	dc.mutation.SetUserID(id)
	return dc
}

// SetNillableUserID sets the "user" edge to the Users entity by ID if the given value is not nil.
func (dc *DevicesCreate) SetNillableUserID(id *uint64) *DevicesCreate {
	if id != nil {
		dc = dc.SetUserID(*id)
	}
	return dc
}

// SetUser sets the "user" edge to the Users entity.
func (dc *DevicesCreate) SetUser(u *Users) *DevicesCreate {
	return dc.SetUserID(u.ID)
}

// Mutation returns the DevicesMutation object of the builder.
func (dc *DevicesCreate) Mutation() *DevicesMutation {
	return dc.mutation
}

// Save creates the Devices in the database.
func (dc *DevicesCreate) Save(ctx context.Context) (*Devices, error) {
	dc.defaults()
	return withHooks(ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DevicesCreate) SaveX(ctx context.Context) *Devices {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DevicesCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DevicesCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DevicesCreate) defaults() {
	if _, ok := dc.mutation.CreateTime(); !ok {
		v := devices.DefaultCreateTime()
		dc.mutation.SetCreateTime(v)
	}
	if _, ok := dc.mutation.UpdateTime(); !ok {
		v := devices.DefaultUpdateTime()
		dc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DevicesCreate) check() error {
	if _, ok := dc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Devices.create_time"`)}
	}
	if _, ok := dc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Devices.update_time"`)}
	}
	if _, ok := dc.mutation.IP(); !ok {
		return &ValidationError{Name: "ip", err: errors.New(`ent: missing required field "Devices.ip"`)}
	}
	if _, ok := dc.mutation.Device(); !ok {
		return &ValidationError{Name: "device", err: errors.New(`ent: missing required field "Devices.device"`)}
	}
	return nil
}

func (dc *DevicesCreate) sqlSave(ctx context.Context) (*Devices, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DevicesCreate) createSpec() (*Devices, *sqlgraph.CreateSpec) {
	var (
		_node = &Devices{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(devices.Table, sqlgraph.NewFieldSpec(devices.FieldID, field.TypeUint64))
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dc.mutation.CreateTime(); ok {
		_spec.SetField(devices.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := dc.mutation.UpdateTime(); ok {
		_spec.SetField(devices.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := dc.mutation.IP(); ok {
		_spec.SetField(devices.FieldIP, field.TypeString, value)
		_node.IP = value
	}
	if value, ok := dc.mutation.Device(); ok {
		_spec.SetField(devices.FieldDevice, field.TypeString, value)
		_node.Device = value
	}
	if nodes := dc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   devices.UserTable,
			Columns: []string{devices.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(users.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.users_devices = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DevicesCreateBulk is the builder for creating many Devices entities in bulk.
type DevicesCreateBulk struct {
	config
	err      error
	builders []*DevicesCreate
}

// Save creates the Devices entities in the database.
func (dcb *DevicesCreateBulk) Save(ctx context.Context) ([]*Devices, error) {
	if dcb.err != nil {
		return nil, dcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Devices, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DevicesMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DevicesCreateBulk) SaveX(ctx context.Context) []*Devices {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DevicesCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DevicesCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
