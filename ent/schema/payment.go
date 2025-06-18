package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Payment holds the schema definition for the Payment entity.
type Payment struct {
	ent.Schema
}

// Fields of the Payment.
func (Payment) Fields() []ent.Field {
	return []ent.Field{
		field.String("transaction_id").
			Unique(),
		field.Enum("method").
			Values("credit_card", "paypal", "bank_transfer", "other"),
		field.Float("amount").
			Min(0),
		field.Enum("status").
			Values("pending", "completed", "failed", "refunded").
			Default("pending"),
		field.Time("processed_at").
			Optional(),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Payment.
func (Payment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("payments").
			Unique(),
		edge.From("order", Order.Type).
			Ref("payments").
			Unique(),
	}
}
