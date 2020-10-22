// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/prakitsrakaew/app/ent/bill"
	"github.com/prakitsrakaew/app/ent/predicate"
	"github.com/prakitsrakaew/app/ent/repairinvoice"
)

// RepairinvoiceUpdate is the builder for updating Repairinvoice entities.
type RepairinvoiceUpdate struct {
	config
	hooks      []Hook
	mutation   *RepairinvoiceMutation
	predicates []predicate.Repairinvoice
}

// Where adds a new predicate for the builder.
func (ru *RepairinvoiceUpdate) Where(ps ...predicate.Repairinvoice) *RepairinvoiceUpdate {
	ru.predicates = append(ru.predicates, ps...)
	return ru
}

// SetSymptomid sets the symptomid field.
func (ru *RepairinvoiceUpdate) SetSymptomid(i int) *RepairinvoiceUpdate {
	ru.mutation.ResetSymptomid()
	ru.mutation.SetSymptomid(i)
	return ru
}

// AddSymptomid adds i to symptomid.
func (ru *RepairinvoiceUpdate) AddSymptomid(i int) *RepairinvoiceUpdate {
	ru.mutation.AddSymptomid(i)
	return ru
}

// SetDeviceid sets the deviceid field.
func (ru *RepairinvoiceUpdate) SetDeviceid(i int) *RepairinvoiceUpdate {
	ru.mutation.ResetDeviceid()
	ru.mutation.SetDeviceid(i)
	return ru
}

// AddDeviceid adds i to deviceid.
func (ru *RepairinvoiceUpdate) AddDeviceid(i int) *RepairinvoiceUpdate {
	ru.mutation.AddDeviceid(i)
	return ru
}

// SetUserid sets the userid field.
func (ru *RepairinvoiceUpdate) SetUserid(i int) *RepairinvoiceUpdate {
	ru.mutation.ResetUserid()
	ru.mutation.SetUserid(i)
	return ru
}

// AddUserid adds i to userid.
func (ru *RepairinvoiceUpdate) AddUserid(i int) *RepairinvoiceUpdate {
	ru.mutation.AddUserid(i)
	return ru
}

// SetStatusrepairid sets the statusrepairid field.
func (ru *RepairinvoiceUpdate) SetStatusrepairid(i int) *RepairinvoiceUpdate {
	ru.mutation.ResetStatusrepairid()
	ru.mutation.SetStatusrepairid(i)
	return ru
}

// AddStatusrepairid adds i to statusrepairid.
func (ru *RepairinvoiceUpdate) AddStatusrepairid(i int) *RepairinvoiceUpdate {
	ru.mutation.AddStatusrepairid(i)
	return ru
}

// SetRepairinvoicesID sets the repairinvoices edge to Bill by id.
func (ru *RepairinvoiceUpdate) SetRepairinvoicesID(id int) *RepairinvoiceUpdate {
	ru.mutation.SetRepairinvoicesID(id)
	return ru
}

// SetNillableRepairinvoicesID sets the repairinvoices edge to Bill by id if the given value is not nil.
func (ru *RepairinvoiceUpdate) SetNillableRepairinvoicesID(id *int) *RepairinvoiceUpdate {
	if id != nil {
		ru = ru.SetRepairinvoicesID(*id)
	}
	return ru
}

// SetRepairinvoices sets the repairinvoices edge to Bill.
func (ru *RepairinvoiceUpdate) SetRepairinvoices(b *Bill) *RepairinvoiceUpdate {
	return ru.SetRepairinvoicesID(b.ID)
}

// Mutation returns the RepairinvoiceMutation object of the builder.
func (ru *RepairinvoiceUpdate) Mutation() *RepairinvoiceMutation {
	return ru.mutation
}

// ClearRepairinvoices clears the repairinvoices edge to Bill.
func (ru *RepairinvoiceUpdate) ClearRepairinvoices() *RepairinvoiceUpdate {
	ru.mutation.ClearRepairinvoices()
	return ru
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ru *RepairinvoiceUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RepairinvoiceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RepairinvoiceUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RepairinvoiceUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RepairinvoiceUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *RepairinvoiceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   repairinvoice.Table,
			Columns: repairinvoice.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: repairinvoice.FieldID,
			},
		},
	}
	if ps := ru.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Symptomid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repairinvoice.FieldSymptomid,
		})
	}
	if value, ok := ru.mutation.AddedSymptomid(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repairinvoice.FieldSymptomid,
		})
	}
	if value, ok := ru.mutation.Deviceid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repairinvoice.FieldDeviceid,
		})
	}
	if value, ok := ru.mutation.AddedDeviceid(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repairinvoice.FieldDeviceid,
		})
	}
	if value, ok := ru.mutation.Userid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repairinvoice.FieldUserid,
		})
	}
	if value, ok := ru.mutation.AddedUserid(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repairinvoice.FieldUserid,
		})
	}
	if value, ok := ru.mutation.Statusrepairid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repairinvoice.FieldStatusrepairid,
		})
	}
	if value, ok := ru.mutation.AddedStatusrepairid(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repairinvoice.FieldStatusrepairid,
		})
	}
	if ru.mutation.RepairinvoicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   repairinvoice.RepairinvoicesTable,
			Columns: []string{repairinvoice.RepairinvoicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bill.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RepairinvoicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   repairinvoice.RepairinvoicesTable,
			Columns: []string{repairinvoice.RepairinvoicesColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{repairinvoice.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// RepairinvoiceUpdateOne is the builder for updating a single Repairinvoice entity.
type RepairinvoiceUpdateOne struct {
	config
	hooks    []Hook
	mutation *RepairinvoiceMutation
}

// SetSymptomid sets the symptomid field.
func (ruo *RepairinvoiceUpdateOne) SetSymptomid(i int) *RepairinvoiceUpdateOne {
	ruo.mutation.ResetSymptomid()
	ruo.mutation.SetSymptomid(i)
	return ruo
}

// AddSymptomid adds i to symptomid.
func (ruo *RepairinvoiceUpdateOne) AddSymptomid(i int) *RepairinvoiceUpdateOne {
	ruo.mutation.AddSymptomid(i)
	return ruo
}

// SetDeviceid sets the deviceid field.
func (ruo *RepairinvoiceUpdateOne) SetDeviceid(i int) *RepairinvoiceUpdateOne {
	ruo.mutation.ResetDeviceid()
	ruo.mutation.SetDeviceid(i)
	return ruo
}

// AddDeviceid adds i to deviceid.
func (ruo *RepairinvoiceUpdateOne) AddDeviceid(i int) *RepairinvoiceUpdateOne {
	ruo.mutation.AddDeviceid(i)
	return ruo
}

// SetUserid sets the userid field.
func (ruo *RepairinvoiceUpdateOne) SetUserid(i int) *RepairinvoiceUpdateOne {
	ruo.mutation.ResetUserid()
	ruo.mutation.SetUserid(i)
	return ruo
}

// AddUserid adds i to userid.
func (ruo *RepairinvoiceUpdateOne) AddUserid(i int) *RepairinvoiceUpdateOne {
	ruo.mutation.AddUserid(i)
	return ruo
}

// SetStatusrepairid sets the statusrepairid field.
func (ruo *RepairinvoiceUpdateOne) SetStatusrepairid(i int) *RepairinvoiceUpdateOne {
	ruo.mutation.ResetStatusrepairid()
	ruo.mutation.SetStatusrepairid(i)
	return ruo
}

// AddStatusrepairid adds i to statusrepairid.
func (ruo *RepairinvoiceUpdateOne) AddStatusrepairid(i int) *RepairinvoiceUpdateOne {
	ruo.mutation.AddStatusrepairid(i)
	return ruo
}

// SetRepairinvoicesID sets the repairinvoices edge to Bill by id.
func (ruo *RepairinvoiceUpdateOne) SetRepairinvoicesID(id int) *RepairinvoiceUpdateOne {
	ruo.mutation.SetRepairinvoicesID(id)
	return ruo
}

// SetNillableRepairinvoicesID sets the repairinvoices edge to Bill by id if the given value is not nil.
func (ruo *RepairinvoiceUpdateOne) SetNillableRepairinvoicesID(id *int) *RepairinvoiceUpdateOne {
	if id != nil {
		ruo = ruo.SetRepairinvoicesID(*id)
	}
	return ruo
}

// SetRepairinvoices sets the repairinvoices edge to Bill.
func (ruo *RepairinvoiceUpdateOne) SetRepairinvoices(b *Bill) *RepairinvoiceUpdateOne {
	return ruo.SetRepairinvoicesID(b.ID)
}

// Mutation returns the RepairinvoiceMutation object of the builder.
func (ruo *RepairinvoiceUpdateOne) Mutation() *RepairinvoiceMutation {
	return ruo.mutation
}

// ClearRepairinvoices clears the repairinvoices edge to Bill.
func (ruo *RepairinvoiceUpdateOne) ClearRepairinvoices() *RepairinvoiceUpdateOne {
	ruo.mutation.ClearRepairinvoices()
	return ruo
}

// Save executes the query and returns the updated entity.
func (ruo *RepairinvoiceUpdateOne) Save(ctx context.Context) (*Repairinvoice, error) {

	var (
		err  error
		node *Repairinvoice
	)
	if len(ruo.hooks) == 0 {
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RepairinvoiceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RepairinvoiceUpdateOne) SaveX(ctx context.Context) *Repairinvoice {
	r, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return r
}

// Exec executes the query on the entity.
func (ruo *RepairinvoiceUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RepairinvoiceUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *RepairinvoiceUpdateOne) sqlSave(ctx context.Context) (r *Repairinvoice, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   repairinvoice.Table,
			Columns: repairinvoice.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: repairinvoice.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Repairinvoice.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := ruo.mutation.Symptomid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repairinvoice.FieldSymptomid,
		})
	}
	if value, ok := ruo.mutation.AddedSymptomid(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repairinvoice.FieldSymptomid,
		})
	}
	if value, ok := ruo.mutation.Deviceid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repairinvoice.FieldDeviceid,
		})
	}
	if value, ok := ruo.mutation.AddedDeviceid(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repairinvoice.FieldDeviceid,
		})
	}
	if value, ok := ruo.mutation.Userid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repairinvoice.FieldUserid,
		})
	}
	if value, ok := ruo.mutation.AddedUserid(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repairinvoice.FieldUserid,
		})
	}
	if value, ok := ruo.mutation.Statusrepairid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repairinvoice.FieldStatusrepairid,
		})
	}
	if value, ok := ruo.mutation.AddedStatusrepairid(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repairinvoice.FieldStatusrepairid,
		})
	}
	if ruo.mutation.RepairinvoicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   repairinvoice.RepairinvoicesTable,
			Columns: []string{repairinvoice.RepairinvoicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bill.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RepairinvoicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   repairinvoice.RepairinvoicesTable,
			Columns: []string{repairinvoice.RepairinvoicesColumn},
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
	r = &Repairinvoice{config: ruo.config}
	_spec.Assign = r.assignValues
	_spec.ScanValues = r.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{repairinvoice.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return r, nil
}
