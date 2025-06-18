package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.Text("description"),
		field.Float("price").
			Min(0),
		field.String("sku").
			Unique(),
		field.String("image_url"),
		field.Int("stock_count").
			Default(0).
			Min(0),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("category", Category.Type).
			Ref("products").
			Unique(),
		edge.To("reviews", Review.Type),
		edge.To("cart_items", CartItem.Type),
		edge.To("order_items", OrderItem.Type),
	}
}
