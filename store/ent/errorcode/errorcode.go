// Code generated by ent, DO NOT EDIT.

package errorcode

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the errorcode type in the database.
	Label = "error_code"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldErrorCode holds the string denoting the error_code field in the database.
	FieldErrorCode = "error_code"
	// FieldGrpcStatus holds the string denoting the grpc_status field in the database.
	FieldGrpcStatus = "grpc_status"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldMessage holds the string denoting the message field in the database.
	FieldMessage = "message"
	// Table holds the table name of the errorcode in the database.
	Table = "error_codes"
)

// Columns holds all SQL columns for errorcode fields.
var Columns = []string{
	FieldID,
	FieldErrorCode,
	FieldGrpcStatus,
	FieldName,
	FieldMessage,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Order defines the ordering method for the ErrorCode queries.
type Order func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByErrorCode orders the results by the error_code field.
func ByErrorCode(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldErrorCode, opts...).ToFunc()
}

// ByGrpcStatus orders the results by the grpc_status field.
func ByGrpcStatus(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldGrpcStatus, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByMessage orders the results by the message field.
func ByMessage(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldMessage, opts...).ToFunc()
}
