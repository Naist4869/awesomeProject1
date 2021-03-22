// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Naist4869/awesomeProject1/ent/issuex1355"
	"github.com/Naist4869/awesomeProject1/ent/predicate"
)

// IssueX1355Delete is the builder for deleting a IssueX1355 entity.
type IssueX1355Delete struct {
	config
	hooks    []Hook
	mutation *IssueX1355Mutation
}

// Where adds a new predicate to the IssueX1355Delete builder.
func (ix *IssueX1355Delete) Where(ps ...predicate.IssueX1355) *IssueX1355Delete {
	ix.mutation.predicates = append(ix.mutation.predicates, ps...)
	return ix
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ix *IssueX1355Delete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ix.hooks) == 0 {
		affected, err = ix.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IssueX1355Mutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ix.mutation = mutation
			affected, err = ix.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ix.hooks) - 1; i >= 0; i-- {
			mut = ix.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ix.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ix *IssueX1355Delete) ExecX(ctx context.Context) int {
	n, err := ix.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ix *IssueX1355Delete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: issuex1355.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: issuex1355.FieldID,
			},
		},
	}
	if ps := ix.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ix.driver, _spec)
}

// IssueX1355DeleteOne is the builder for deleting a single IssueX1355 entity.
type IssueX1355DeleteOne struct {
	ix *IssueX1355Delete
}

// Exec executes the deletion query.
func (ixo *IssueX1355DeleteOne) Exec(ctx context.Context) error {
	n, err := ixo.ix.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{issuex1355.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ixo *IssueX1355DeleteOne) ExecX(ctx context.Context) {
	ixo.ix.ExecX(ctx)
}