// Code generated by ent, DO NOT EDIT.

package errorcode

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Yostardev/errors/store/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldEQ(FieldUpdatedAt, v))
}

// ErrorCode applies equality check predicate on the "error_code" field. It's identical to ErrorCodeEQ.
func ErrorCode(v int) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldEQ(FieldErrorCode, v))
}

// GrpcStatus applies equality check predicate on the "grpc_status" field. It's identical to GrpcStatusEQ.
func GrpcStatus(v uint32) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldEQ(FieldGrpcStatus, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldEQ(FieldName, v))
}

// Message applies equality check predicate on the "message" field. It's identical to MessageEQ.
func Message(v string) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldEQ(FieldMessage, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldLTE(FieldUpdatedAt, v))
}

// ErrorCodeEQ applies the EQ predicate on the "error_code" field.
func ErrorCodeEQ(v int) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldEQ(FieldErrorCode, v))
}

// ErrorCodeNEQ applies the NEQ predicate on the "error_code" field.
func ErrorCodeNEQ(v int) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldNEQ(FieldErrorCode, v))
}

// ErrorCodeIn applies the In predicate on the "error_code" field.
func ErrorCodeIn(vs ...int) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldIn(FieldErrorCode, vs...))
}

// ErrorCodeNotIn applies the NotIn predicate on the "error_code" field.
func ErrorCodeNotIn(vs ...int) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldNotIn(FieldErrorCode, vs...))
}

// ErrorCodeGT applies the GT predicate on the "error_code" field.
func ErrorCodeGT(v int) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldGT(FieldErrorCode, v))
}

// ErrorCodeGTE applies the GTE predicate on the "error_code" field.
func ErrorCodeGTE(v int) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldGTE(FieldErrorCode, v))
}

// ErrorCodeLT applies the LT predicate on the "error_code" field.
func ErrorCodeLT(v int) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldLT(FieldErrorCode, v))
}

// ErrorCodeLTE applies the LTE predicate on the "error_code" field.
func ErrorCodeLTE(v int) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldLTE(FieldErrorCode, v))
}

// GrpcStatusEQ applies the EQ predicate on the "grpc_status" field.
func GrpcStatusEQ(v uint32) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldEQ(FieldGrpcStatus, v))
}

// GrpcStatusNEQ applies the NEQ predicate on the "grpc_status" field.
func GrpcStatusNEQ(v uint32) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldNEQ(FieldGrpcStatus, v))
}

// GrpcStatusIn applies the In predicate on the "grpc_status" field.
func GrpcStatusIn(vs ...uint32) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldIn(FieldGrpcStatus, vs...))
}

// GrpcStatusNotIn applies the NotIn predicate on the "grpc_status" field.
func GrpcStatusNotIn(vs ...uint32) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldNotIn(FieldGrpcStatus, vs...))
}

// GrpcStatusGT applies the GT predicate on the "grpc_status" field.
func GrpcStatusGT(v uint32) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldGT(FieldGrpcStatus, v))
}

// GrpcStatusGTE applies the GTE predicate on the "grpc_status" field.
func GrpcStatusGTE(v uint32) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldGTE(FieldGrpcStatus, v))
}

// GrpcStatusLT applies the LT predicate on the "grpc_status" field.
func GrpcStatusLT(v uint32) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldLT(FieldGrpcStatus, v))
}

// GrpcStatusLTE applies the LTE predicate on the "grpc_status" field.
func GrpcStatusLTE(v uint32) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldLTE(FieldGrpcStatus, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldContainsFold(FieldName, v))
}

// MessageEQ applies the EQ predicate on the "message" field.
func MessageEQ(v string) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldEQ(FieldMessage, v))
}

// MessageNEQ applies the NEQ predicate on the "message" field.
func MessageNEQ(v string) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldNEQ(FieldMessage, v))
}

// MessageIn applies the In predicate on the "message" field.
func MessageIn(vs ...string) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldIn(FieldMessage, vs...))
}

// MessageNotIn applies the NotIn predicate on the "message" field.
func MessageNotIn(vs ...string) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldNotIn(FieldMessage, vs...))
}

// MessageGT applies the GT predicate on the "message" field.
func MessageGT(v string) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldGT(FieldMessage, v))
}

// MessageGTE applies the GTE predicate on the "message" field.
func MessageGTE(v string) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldGTE(FieldMessage, v))
}

// MessageLT applies the LT predicate on the "message" field.
func MessageLT(v string) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldLT(FieldMessage, v))
}

// MessageLTE applies the LTE predicate on the "message" field.
func MessageLTE(v string) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldLTE(FieldMessage, v))
}

// MessageContains applies the Contains predicate on the "message" field.
func MessageContains(v string) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldContains(FieldMessage, v))
}

// MessageHasPrefix applies the HasPrefix predicate on the "message" field.
func MessageHasPrefix(v string) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldHasPrefix(FieldMessage, v))
}

// MessageHasSuffix applies the HasSuffix predicate on the "message" field.
func MessageHasSuffix(v string) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldHasSuffix(FieldMessage, v))
}

// MessageEqualFold applies the EqualFold predicate on the "message" field.
func MessageEqualFold(v string) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldEqualFold(FieldMessage, v))
}

// MessageContainsFold applies the ContainsFold predicate on the "message" field.
func MessageContainsFold(v string) predicate.ErrorCode {
	return predicate.ErrorCode(sql.FieldContainsFold(FieldMessage, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ErrorCode) predicate.ErrorCode {
	return predicate.ErrorCode(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ErrorCode) predicate.ErrorCode {
	return predicate.ErrorCode(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ErrorCode) predicate.ErrorCode {
	return predicate.ErrorCode(func(s *sql.Selector) {
		p(s.Not())
	})
}