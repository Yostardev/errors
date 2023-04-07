package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ErrorCode holds the schema definition for the ErrorCode entity.
type ErrorCode struct {
	ent.Schema
}

// Fields of the ErrorCode.
func (ErrorCode) Fields() []ent.Field {
	return []ent.Field{
		field.Int("error_code").Unique().Comment("错误码"),
		field.Uint32("grpc_status").Comment("Grpc状态"),
		field.String("name").Unique().Comment("错误名（唯一标识）"),
		field.String("message").Comment("错误信息"),
	}
}

// Edges of the ErrorCode.
func (ErrorCode) Edges() []ent.Edge {
	return nil
}

// Mixin of the ErrorCode.
func (ErrorCode) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Meta{},
	}
}
