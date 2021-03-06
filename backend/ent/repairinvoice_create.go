// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/prakitsrakaew/app/ent/bill"
	"github.com/prakitsrakaew/app/ent/repairinvoice"
)

// RepairinvoiceCreate is the builder for creating a Repairinvoice entity.
type RepairinvoiceCreate struct {
	config
	mutation *RepairinvoiceMutation
	hooks    []Hook
}

// SetSymptomid sets the symptomid field.
func (rc *RepairinvoiceCreate) SetSymptomid(i int) *RepairinvoiceCreate {
	rc.mutation.SetSymptomid(i)
	return rc
}

// SetDeviceid sets the deviceid field.
func (rc *RepairinvoiceCreate) SetDeviceid(i int) *RepairinvoiceCreate {
	rc.mutation.SetDeviceid(i)
	return rc
}

// SetUserid sets the userid field.
func (rc *RepairinvoiceCreate) SetUserid(i int) *RepairinvoiceCreate {
	rc.mutation.SetUserid(i)
	return rc
}

// SetStatusrepairid sets the statusrepairid field.
func (rc *RepairinvoiceCreate) SetStatusrepairid(i int) *RepairinvoiceCreate {
	rc.mutation.SetStatusrepairid(i)
	return rc
}

// SetRepairinvoicesID sets the repairinvoices edge to Bill by id.
func (rc *RepairinvoiceCreate) SetRepairinvoicesID(id int) *RepairinvoiceCreate {
	rc.mutation.SetRepairinvoicesID(id)
	return rc
}

// SetNillableRepairinvoicesID sets the repairinvoices edge to Bill by id if the given value is not nil.
func (rc *RepairinvoiceCreate) SetNillableRepairinvoicesID(id *int) *RepairinvoiceCreate {
	if id != nil {
		rc = rc.SetRepairinvoicesID(*id)
	}
	return rc
}

// SetRepairinvoices sets the repairinvoices edge to Bill.
func (rc *RepairinvoiceCreate) SetRepairinvoices(b *Bill) *RepairinvoiceCreate {
	return rc.SetRepairinvoicesID(b.ID)
}

// Mutation returns the RepairinvoiceMutation object of the builder.
func (rc *RepairinvoiceCreate) Mutation() *RepairinvoiceMutation {
	return rc.mutation
}

// Save creates the Repairinvoice in the database.
func (rc *RepairinvoiceCreate) Save(ctx context.Context) (*Repairinvoice, error) {
	if _, ok := rc.mutation.Symptomid(); !ok {
		return nil, &ValidationError{Name: "symptomid", err: errors.New("ent: missing required field \"symptomid\"")}
	}
	if _, ok := rc.mutation.Deviceid(); !ok {
		return nil, &ValidationError{Name: "deviceid", err: errors.New("ent: missing required field \"deviceid\"")}
	}
	if _, ok := rc.mutation.Userid(); !ok {
		return nil, &ValidationError{Name: "userid", err: errors.New("ent: missing required field \"userid\"")}
	}
	if _, ok := rc.mutation.Statusrepairid(); !ok {
		return nil, &ValidationError{Name: "statusrepairid", err: errors.New("ent: missing required field \"statusrepairid\"")}
	}
	var (
		err  error
		node *Repairinvoice
	)
	if len(rc.hooks) == 0 {
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RepairinvoiceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rc.mutation = mutation
			node, err = rc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RepairinvoiceCreate) SaveX(ctx context.Context) *Repairinvoice {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rc *RepairinvoiceCreate) sqlSave(ctx context.Context) (*Repairinvoice, error) {
	r, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	r.ID = int(id)
	return r, nil
}

func (rc *RepairinvoiceCreate) createSpec() (*Repairinvoice, *sqlgraph.CreateSpec) {
	var (
		r     = &Repairinvoice{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: repairinvoice.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: repairinvoice.FieldID,
			},
		}
	)
	if value, ok := rc.mutation.Symptomid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repairinvoice.FieldSymptomid,
		})
		r.Symptomid = value
	}
	if value, ok := rc.mutation.Deviceid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repairinvoice.FieldDeviceid,
		})
		r.Deviceid = value
	}
	if value, ok := rc.mutation.Userid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repairinvoice.FieldUserid,
		})
		r.Userid = value
	}
	if value, ok := rc.mutation.Statusrepairid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repairinvoice.FieldStatusrepairid,
		})
		r.Statusrepairid = value
	}
	if nodes := rc.mutation.RepairinvoicesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return r, _spec
}
