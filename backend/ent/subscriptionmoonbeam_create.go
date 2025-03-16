// Code generated by ent, DO NOT EDIT.

package ent

import (
	"arjunmal1311/only_fans_on_chain/backend/ent/subscriptionmoonbeam"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SubscriptionMoonbeamCreate is the builder for creating a SubscriptionMoonbeam entity.
type SubscriptionMoonbeamCreate struct {
	config
	mutation *SubscriptionMoonbeamMutation
	hooks    []Hook
}

// Mutation returns the SubscriptionMoonbeamMutation object of the builder.
func (smc *SubscriptionMoonbeamCreate) Mutation() *SubscriptionMoonbeamMutation {
	return smc.mutation
}

// Save creates the SubscriptionMoonbeam in the database.
func (smc *SubscriptionMoonbeamCreate) Save(ctx context.Context) (*SubscriptionMoonbeam, error) {
	return withHooks(ctx, smc.sqlSave, smc.mutation, smc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (smc *SubscriptionMoonbeamCreate) SaveX(ctx context.Context) *SubscriptionMoonbeam {
	v, err := smc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (smc *SubscriptionMoonbeamCreate) Exec(ctx context.Context) error {
	_, err := smc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smc *SubscriptionMoonbeamCreate) ExecX(ctx context.Context) {
	if err := smc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (smc *SubscriptionMoonbeamCreate) check() error {
	return nil
}

func (smc *SubscriptionMoonbeamCreate) sqlSave(ctx context.Context) (*SubscriptionMoonbeam, error) {
	if err := smc.check(); err != nil {
		return nil, err
	}
	_node, _spec := smc.createSpec()
	if err := sqlgraph.CreateNode(ctx, smc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	smc.mutation.id = &_node.ID
	smc.mutation.done = true
	return _node, nil
}

func (smc *SubscriptionMoonbeamCreate) createSpec() (*SubscriptionMoonbeam, *sqlgraph.CreateSpec) {
	var (
		_node = &SubscriptionMoonbeam{config: smc.config}
		_spec = sqlgraph.NewCreateSpec(subscriptionmoonbeam.Table, sqlgraph.NewFieldSpec(subscriptionmoonbeam.FieldID, field.TypeInt))
	)
	return _node, _spec
}

// SubscriptionMoonbeamCreateBulk is the builder for creating many SubscriptionMoonbeam entities in bulk.
type SubscriptionMoonbeamCreateBulk struct {
	config
	err      error
	builders []*SubscriptionMoonbeamCreate
}

// Save creates the SubscriptionMoonbeam entities in the database.
func (smcb *SubscriptionMoonbeamCreateBulk) Save(ctx context.Context) ([]*SubscriptionMoonbeam, error) {
	if smcb.err != nil {
		return nil, smcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(smcb.builders))
	nodes := make([]*SubscriptionMoonbeam, len(smcb.builders))
	mutators := make([]Mutator, len(smcb.builders))
	for i := range smcb.builders {
		func(i int, root context.Context) {
			builder := smcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SubscriptionMoonbeamMutation)
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
					_, err = mutators[i+1].Mutate(root, smcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, smcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, smcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (smcb *SubscriptionMoonbeamCreateBulk) SaveX(ctx context.Context) []*SubscriptionMoonbeam {
	v, err := smcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (smcb *SubscriptionMoonbeamCreateBulk) Exec(ctx context.Context) error {
	_, err := smcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smcb *SubscriptionMoonbeamCreateBulk) ExecX(ctx context.Context) {
	if err := smcb.Exec(ctx); err != nil {
		panic(err)
	}
}
