// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/disism/saikan/ent/audiobooks"
	"github.com/disism/saikan/ent/predicate"
)

// AudiobooksDelete is the builder for deleting a Audiobooks entity.
type AudiobooksDelete struct {
	config
	hooks    []Hook
	mutation *AudiobooksMutation
}

// Where appends a list predicates to the AudiobooksDelete builder.
func (ad *AudiobooksDelete) Where(ps ...predicate.Audiobooks) *AudiobooksDelete {
	ad.mutation.Where(ps...)
	return ad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ad *AudiobooksDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ad.sqlExec, ad.mutation, ad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ad *AudiobooksDelete) ExecX(ctx context.Context) int {
	n, err := ad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ad *AudiobooksDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(audiobooks.Table, sqlgraph.NewFieldSpec(audiobooks.FieldID, field.TypeInt))
	if ps := ad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ad.mutation.done = true
	return affected, err
}

// AudiobooksDeleteOne is the builder for deleting a single Audiobooks entity.
type AudiobooksDeleteOne struct {
	ad *AudiobooksDelete
}

// Where appends a list predicates to the AudiobooksDelete builder.
func (ado *AudiobooksDeleteOne) Where(ps ...predicate.Audiobooks) *AudiobooksDeleteOne {
	ado.ad.mutation.Where(ps...)
	return ado
}

// Exec executes the deletion query.
func (ado *AudiobooksDeleteOne) Exec(ctx context.Context) error {
	n, err := ado.ad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{audiobooks.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ado *AudiobooksDeleteOne) ExecX(ctx context.Context) {
	if err := ado.Exec(ctx); err != nil {
		panic(err)
	}
}
