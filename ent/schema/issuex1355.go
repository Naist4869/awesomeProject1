package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// IssueX1355 holds the schema definition for the IssueX1355 entity.
type IssueX1355 struct {
	ent.Schema
}

// Fields of the IssueX1355.
func (IssueX1355) Fields() []ent.Field {
	return []ent.Field{
		field.Time("time").Optional(),
		field.Int32("int32").Optional(),
		field.String("str").Default("abc"),
	}
}

// Edges of the IssueX1355.
func (IssueX1355) Edges() []ent.Edge {
	return nil
}
