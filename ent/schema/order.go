package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.String("order_number").
			Unique(),
		field.Enum("status").
			Values("pending", "processing", "shipped", "delivered", "cancelled").
			Default("pending"),
		field.Float("shipping_cost").
			Default(0),
		field.Float("tax").
			Default(0),
		field.Float("total").
			Default(0),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("orders").
			Unique(),
		edge.To("items", OrderItem.Type),
		edge.To("payments", Payment.Type),
		edge.From("shipping_address", Address.Type).
			Ref("shipping_orders").
			Unique(),
		edge.From("billing_address", Address.Type).
			Ref("billing_orders").
			Unique(),
	}
}
