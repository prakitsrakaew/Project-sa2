package schema

import (
	"github.com/facebookincubator/ent"

	"github.com/facebookincubator/ent/schema/edge"

	"github.com/facebookincubator/ent/schema/field"
)

// Bill holds the schema definition for the Bill entity.
type Bill struct {
	ent.Schema
}

// Fields of the Bill.
func (Bill) Fields() []ent.Field {
	return []ent.Field{

		field.Int("price").Unique(),

		field.Int("time").Unique(),
	}
}

// Edges of the Bill.
func (Bill) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Repairinvoice", Repairinvoice.Type).Ref("repairinvoices").Unique().Required(),

		edge.From("Employee", Employee.Type).Ref("employees").Unique(),

		edge.From("Billingstatus", Billingstatus.Type).Ref("billingstatuss").Unique(),
	}
}
