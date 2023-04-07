// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Yostardev/errors/store/ent/errorcode"
)

// ErrorCodeCreate is the builder for creating a ErrorCode entity.
type ErrorCodeCreate struct {
	config
	mutation *ErrorCodeMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ecc *ErrorCodeCreate) SetCreatedAt(t time.Time) *ErrorCodeCreate {
	ecc.mutation.SetCreatedAt(t)
	return ecc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ecc *ErrorCodeCreate) SetNillableCreatedAt(t *time.Time) *ErrorCodeCreate {
	if t != nil {
		ecc.SetCreatedAt(*t)
	}
	return ecc
}

// SetUpdatedAt sets the "updated_at" field.
func (ecc *ErrorCodeCreate) SetUpdatedAt(t time.Time) *ErrorCodeCreate {
	ecc.mutation.SetUpdatedAt(t)
	return ecc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ecc *ErrorCodeCreate) SetNillableUpdatedAt(t *time.Time) *ErrorCodeCreate {
	if t != nil {
		ecc.SetUpdatedAt(*t)
	}
	return ecc
}

// SetErrorCode sets the "error_code" field.
func (ecc *ErrorCodeCreate) SetErrorCode(i int) *ErrorCodeCreate {
	ecc.mutation.SetErrorCode(i)
	return ecc
}

// SetGrpcStatus sets the "grpc_status" field.
func (ecc *ErrorCodeCreate) SetGrpcStatus(u uint32) *ErrorCodeCreate {
	ecc.mutation.SetGrpcStatus(u)
	return ecc
}

// SetName sets the "name" field.
func (ecc *ErrorCodeCreate) SetName(s string) *ErrorCodeCreate {
	ecc.mutation.SetName(s)
	return ecc
}

// SetMessage sets the "message" field.
func (ecc *ErrorCodeCreate) SetMessage(s string) *ErrorCodeCreate {
	ecc.mutation.SetMessage(s)
	return ecc
}

// SetID sets the "id" field.
func (ecc *ErrorCodeCreate) SetID(i int64) *ErrorCodeCreate {
	ecc.mutation.SetID(i)
	return ecc
}

// Mutation returns the ErrorCodeMutation object of the builder.
func (ecc *ErrorCodeCreate) Mutation() *ErrorCodeMutation {
	return ecc.mutation
}

// Save creates the ErrorCode in the database.
func (ecc *ErrorCodeCreate) Save(ctx context.Context) (*ErrorCode, error) {
	ecc.defaults()
	return withHooks[*ErrorCode, ErrorCodeMutation](ctx, ecc.sqlSave, ecc.mutation, ecc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ecc *ErrorCodeCreate) SaveX(ctx context.Context) *ErrorCode {
	v, err := ecc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecc *ErrorCodeCreate) Exec(ctx context.Context) error {
	_, err := ecc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecc *ErrorCodeCreate) ExecX(ctx context.Context) {
	if err := ecc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ecc *ErrorCodeCreate) defaults() {
	if _, ok := ecc.mutation.CreatedAt(); !ok {
		v := errorcode.DefaultCreatedAt()
		ecc.mutation.SetCreatedAt(v)
	}
	if _, ok := ecc.mutation.UpdatedAt(); !ok {
		v := errorcode.DefaultUpdatedAt()
		ecc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ecc *ErrorCodeCreate) check() error {
	if _, ok := ecc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ErrorCode.created_at"`)}
	}
	if _, ok := ecc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ErrorCode.updated_at"`)}
	}
	if _, ok := ecc.mutation.ErrorCode(); !ok {
		return &ValidationError{Name: "error_code", err: errors.New(`ent: missing required field "ErrorCode.error_code"`)}
	}
	if _, ok := ecc.mutation.GrpcStatus(); !ok {
		return &ValidationError{Name: "grpc_status", err: errors.New(`ent: missing required field "ErrorCode.grpc_status"`)}
	}
	if _, ok := ecc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "ErrorCode.name"`)}
	}
	if _, ok := ecc.mutation.Message(); !ok {
		return &ValidationError{Name: "message", err: errors.New(`ent: missing required field "ErrorCode.message"`)}
	}
	return nil
}

func (ecc *ErrorCodeCreate) sqlSave(ctx context.Context) (*ErrorCode, error) {
	if err := ecc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ecc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ecc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	ecc.mutation.id = &_node.ID
	ecc.mutation.done = true
	return _node, nil
}

func (ecc *ErrorCodeCreate) createSpec() (*ErrorCode, *sqlgraph.CreateSpec) {
	var (
		_node = &ErrorCode{config: ecc.config}
		_spec = sqlgraph.NewCreateSpec(errorcode.Table, sqlgraph.NewFieldSpec(errorcode.FieldID, field.TypeInt64))
	)
	if id, ok := ecc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ecc.mutation.CreatedAt(); ok {
		_spec.SetField(errorcode.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ecc.mutation.UpdatedAt(); ok {
		_spec.SetField(errorcode.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ecc.mutation.ErrorCode(); ok {
		_spec.SetField(errorcode.FieldErrorCode, field.TypeInt, value)
		_node.ErrorCode = value
	}
	if value, ok := ecc.mutation.GrpcStatus(); ok {
		_spec.SetField(errorcode.FieldGrpcStatus, field.TypeUint32, value)
		_node.GrpcStatus = value
	}
	if value, ok := ecc.mutation.Name(); ok {
		_spec.SetField(errorcode.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ecc.mutation.Message(); ok {
		_spec.SetField(errorcode.FieldMessage, field.TypeString, value)
		_node.Message = value
	}
	return _node, _spec
}

// ErrorCodeCreateBulk is the builder for creating many ErrorCode entities in bulk.
type ErrorCodeCreateBulk struct {
	config
	builders []*ErrorCodeCreate
}

// Save creates the ErrorCode entities in the database.
func (eccb *ErrorCodeCreateBulk) Save(ctx context.Context) ([]*ErrorCode, error) {
	specs := make([]*sqlgraph.CreateSpec, len(eccb.builders))
	nodes := make([]*ErrorCode, len(eccb.builders))
	mutators := make([]Mutator, len(eccb.builders))
	for i := range eccb.builders {
		func(i int, root context.Context) {
			builder := eccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ErrorCodeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, eccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, eccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, eccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (eccb *ErrorCodeCreateBulk) SaveX(ctx context.Context) []*ErrorCode {
	v, err := eccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eccb *ErrorCodeCreateBulk) Exec(ctx context.Context) error {
	_, err := eccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eccb *ErrorCodeCreateBulk) ExecX(ctx context.Context) {
	if err := eccb.Exec(ctx); err != nil {
		panic(err)
	}
}
