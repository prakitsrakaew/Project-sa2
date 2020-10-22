// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/prakitsrakaew/app/ent/bill"
	"github.com/prakitsrakaew/app/ent/billingstatus"
	"github.com/prakitsrakaew/app/ent/predicate"
)

// BillingstatusUpdate is the builder for updating Billingstatus entities.
type BillingstatusUpdate struct {
	config
	hooks      []Hook
	mutation   *BillingstatusMutation
	predicates []predicate.Billingstatus
}

// Where adds a new predicate for the builder.
func (bu *BillingstatusUpdate) Where(ps ...predicate.Billingstatus) *BillingstatusUpdate {
	bu.predicates = append(bu.predicates, ps...)
	return bu
}

// SetBillingstatusname sets the billingstatusname field.
func (bu *BillingstatusUpdate) SetBillingstatusname(s string) *BillingstatusUpdate {
	bu.mutation.SetBillingstatusname(s)
	return bu
}

// AddBillingstatusIDs adds the billingstatuss edge to Bill by ids.
func (bu *BillingstatusUpdate) AddBillingstatusIDs(ids ...int) *BillingstatusUpdate {
	bu.mutation.AddBillingstatusIDs(ids...)
	return bu
}

// AddBillingstatuss adds the billingstatuss edges to Bill.
func (bu *BillingstatusUpdate) AddBillingstatuss(b ...*Bill) *BillingstatusUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bu.AddBillingstatusIDs(ids...)
}

// Mutation returns the BillingstatusMutation object of the builder.
func (bu *BillingstatusUpdate) Mutation() *BillingstatusMutation {
	return bu.mutation
}

// RemoveBillingstatusIDs removes the billingstatuss edge to Bill by ids.
func (bu *BillingstatusUpdate) RemoveBillingstatusIDs(ids ...int) *BillingstatusUpdate {
	bu.mutation.RemoveBillingstatusIDs(ids...)
	return bu
}

// RemoveBillingstatuss removes billingstatuss edges to Bill.
func (bu *BillingstatusUpdate) RemoveBillingstatuss(b ...*Bill) *BillingstatusUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bu.RemoveBillingstatusIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (bu *BillingstatusUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(bu.hooks) == 0 {
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BillingstatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bu.mutation = mutation
			affected, err = bu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bu.hooks) - 1; i >= 0; i-- {
			mut = bu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BillingstatusUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BillingstatusUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BillingstatusUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BillingstatusUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   billingstatus.Table,
			Columns: billingstatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: billingstatus.FieldID,
			},
		},
	}
	if ps := bu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Billingstatusname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: billingstatus.FieldBillingstatusname,
		})
	}
	if nodes := bu.mutation.RemovedBillingstatussIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   billingstatus.BillingstatussTable,
			Columns: []string{billingstatus.BillingstatussColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.BillingstatussIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   billingstatus.BillingstatussTable,
			Columns: []string{billingstatus.BillingstatussColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{billingstatus.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// BillingstatusUpdateOne is the builder for updating a single Billingstatus entity.
type BillingstatusUpdateOne struct {
	config
	hooks    []Hook
	mutation *BillingstatusMutation
}

// SetBillingstatusname sets the billingstatusname field.
func (buo *BillingstatusUpdateOne) SetBillingstatusname(s string) *BillingstatusUpdateOne {
	buo.mutation.SetBillingstatusname(s)
	return buo
}

// AddBillingstatusIDs adds the billingstatuss edge to Bill by ids.
func (buo *BillingstatusUpdateOne) AddBillingstatusIDs(ids ...int) *BillingstatusUpdateOne {
	buo.mutation.AddBillingstatusIDs(ids...)
	return buo
}

// AddBillingstatuss adds the billingstatuss edges to Bill.
func (buo *BillingstatusUpdateOne) AddBillingstatuss(b ...*Bill) *BillingstatusUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return buo.AddBillingstatusIDs(ids...)
}

// Mutation returns the BillingstatusMutation object of the builder.
func (buo *BillingstatusUpdateOne) Mutation() *BillingstatusMutation {
	return buo.mutation
}

// RemoveBillingstatusIDs removes the billingstatuss edge to Bill by ids.
func (buo *BillingstatusUpdateOne) RemoveBillingstatusIDs(ids ...int) *BillingstatusUpdateOne {
	buo.mutation.RemoveBillingstatusIDs(ids...)
	return buo
}

// RemoveBillingstatuss removes billingstatuss edges to Bill.
func (buo *BillingstatusUpdateOne) RemoveBillingstatuss(b ...*Bill) *BillingstatusUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return buo.RemoveBillingstatusIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (buo *BillingstatusUpdateOne) Save(ctx context.Context) (*Billingstatus, error) {

	var (
		err  error
		node *Billingstatus
	)
	if len(buo.hooks) == 0 {
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BillingstatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			buo.mutation = mutation
			node, err = buo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(buo.hooks) - 1; i >= 0; i-- {
			mut = buo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, buo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BillingstatusUpdateOne) SaveX(ctx context.Context) *Billingstatus {
	b, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return b
}

// Exec executes the query on the entity.
func (buo *BillingstatusUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BillingstatusUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BillingstatusUpdateOne) sqlSave(ctx context.Context) (b *Billingstatus, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   billingstatus.Table,
			Columns: billingstatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: billingstatus.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Billingstatus.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := buo.mutation.Billingstatusname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: billingstatus.FieldBillingstatusname,
		})
	}
	if nodes := buo.mutation.RemovedBillingstatussIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   billingstatus.BillingstatussTable,
			Columns: []string{billingstatus.BillingstatussColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.BillingstatussIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   billingstatus.BillingstatussTable,
			Columns: []string{billingstatus.BillingstatussColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	b = &Billingstatus{config: buo.config}
	_spec.Assign = b.assignValues
	_spec.ScanValues = b.scanValues()
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{billingstatus.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return b, nil
}