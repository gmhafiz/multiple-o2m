// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/bug/ent/multiplemany"
	"entgo.io/bug/ent/user"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MultipleManyCreate is the builder for creating a MultipleMany entity.
type MultipleManyCreate struct {
	config
	mutation *MultipleManyMutation
	hooks    []Hook
}

// SetUser1 sets the "user_1" field.
func (mmc *MultipleManyCreate) SetUser1(i int) *MultipleManyCreate {
	mmc.mutation.SetUser1(i)
	return mmc
}

// SetUser2 sets the "user_2" field.
func (mmc *MultipleManyCreate) SetUser2(i int) *MultipleManyCreate {
	mmc.mutation.SetUser2(i)
	return mmc
}

// SetUser1ID sets the "user1" edge to the User entity by ID.
func (mmc *MultipleManyCreate) SetUser1ID(id int) *MultipleManyCreate {
	mmc.mutation.SetUser1ID(id)
	return mmc
}

// SetUser1 sets the "user1" edge to the User entity.
func (mmc *MultipleManyCreate) SetUser1(u *User) *MultipleManyCreate {
	return mmc.SetUser1ID(u.ID)
}

// SetUser2ID sets the "user2" edge to the User entity by ID.
func (mmc *MultipleManyCreate) SetUser2ID(id int) *MultipleManyCreate {
	mmc.mutation.SetUser2ID(id)
	return mmc
}

// SetUser2 sets the "user2" edge to the User entity.
func (mmc *MultipleManyCreate) SetUser2(u *User) *MultipleManyCreate {
	return mmc.SetUser2ID(u.ID)
}

// Mutation returns the MultipleManyMutation object of the builder.
func (mmc *MultipleManyCreate) Mutation() *MultipleManyMutation {
	return mmc.mutation
}

// Save creates the MultipleMany in the database.
func (mmc *MultipleManyCreate) Save(ctx context.Context) (*MultipleMany, error) {
	var (
		err  error
		node *MultipleMany
	)
	if len(mmc.hooks) == 0 {
		if err = mmc.check(); err != nil {
			return nil, err
		}
		node, err = mmc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MultipleManyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mmc.check(); err != nil {
				return nil, err
			}
			mmc.mutation = mutation
			if node, err = mmc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mmc.hooks) - 1; i >= 0; i-- {
			if mmc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mmc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mmc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mmc *MultipleManyCreate) SaveX(ctx context.Context) *MultipleMany {
	v, err := mmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mmc *MultipleManyCreate) Exec(ctx context.Context) error {
	_, err := mmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mmc *MultipleManyCreate) ExecX(ctx context.Context) {
	if err := mmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mmc *MultipleManyCreate) check() error {
	if _, ok := mmc.mutation.User1(); !ok {
		return &ValidationError{Name: "user_1", err: errors.New(`ent: missing required field "MultipleMany.user_1"`)}
	}
	if _, ok := mmc.mutation.User2(); !ok {
		return &ValidationError{Name: "user_2", err: errors.New(`ent: missing required field "MultipleMany.user_2"`)}
	}
	if _, ok := mmc.mutation.User1ID(); !ok {
		return &ValidationError{Name: "user1", err: errors.New(`ent: missing required edge "MultipleMany.user1"`)}
	}
	if _, ok := mmc.mutation.User2ID(); !ok {
		return &ValidationError{Name: "user2", err: errors.New(`ent: missing required edge "MultipleMany.user2"`)}
	}
	return nil
}

func (mmc *MultipleManyCreate) sqlSave(ctx context.Context) (*MultipleMany, error) {
	_node, _spec := mmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (mmc *MultipleManyCreate) createSpec() (*MultipleMany, *sqlgraph.CreateSpec) {
	var (
		_node = &MultipleMany{config: mmc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: multiplemany.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: multiplemany.FieldID,
			},
		}
	)
	if nodes := mmc.mutation.User1IDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   multiplemany.User1Table,
			Columns: []string{multiplemany.User1Column},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.User1 = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mmc.mutation.User2IDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   multiplemany.User2Table,
			Columns: []string{multiplemany.User2Column},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.User2 = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MultipleManyCreateBulk is the builder for creating many MultipleMany entities in bulk.
type MultipleManyCreateBulk struct {
	config
	builders []*MultipleManyCreate
}

// Save creates the MultipleMany entities in the database.
func (mmcb *MultipleManyCreateBulk) Save(ctx context.Context) ([]*MultipleMany, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mmcb.builders))
	nodes := make([]*MultipleMany, len(mmcb.builders))
	mutators := make([]Mutator, len(mmcb.builders))
	for i := range mmcb.builders {
		func(i int, root context.Context) {
			builder := mmcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MultipleManyMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mmcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mmcb *MultipleManyCreateBulk) SaveX(ctx context.Context) []*MultipleMany {
	v, err := mmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mmcb *MultipleManyCreateBulk) Exec(ctx context.Context) error {
	_, err := mmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mmcb *MultipleManyCreateBulk) ExecX(ctx context.Context) {
	if err := mmcb.Exec(ctx); err != nil {
		panic(err)
	}
}