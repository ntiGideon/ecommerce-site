package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Address holds the schema definition for the Address entity.
type Address struct {
	ent.Schema
}

// Fields of the Address.
func (Address) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name").
			NotEmpty(),
		field.String("last_name").
			NotEmpty(),
		field.String("company"),
		field.String("address1").
			NotEmpty(),
		field.String("address2"),
		field.String("city").
			NotEmpty(),
		field.String("state").
			NotEmpty(),
		field.String("zip").
			NotEmpty(),
		field.String("country").
			NotEmpty(),
		field.String("phone"),
		field.Bool("default").
			Default(false),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Address.
func (Address) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("addresses").
			Unique(),
		edge.To("shipping_orders", Order.Type),
		edge.To("billing_orders", Order.Type),
	}
}
