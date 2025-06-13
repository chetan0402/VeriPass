// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chetan0402/veripass/internal/ent/pass"
	"github.com/chetan0402/veripass/internal/ent/predicate"
)

// PassUpdate is the builder for updating Pass entities.
type PassUpdate struct {
	config
	hooks    []Hook
	mutation *PassMutation
}

// Where appends a list predicates to the PassUpdate builder.
func (pu *PassUpdate) Where(ps ...predicate.Pass) *PassUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUserID sets the "user_id" field.
func (pu *PassUpdate) SetUserID(s string) *PassUpdate {
	pu.mutation.SetUserID(s)
	return pu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (pu *PassUpdate) SetNillableUserID(s *string) *PassUpdate {
	if s != nil {
		pu.SetUserID(*s)
	}
	return pu
}

// SetType sets the "type" field.
func (pu *PassUpdate) SetType(pa pass.Type) *PassUpdate {
	pu.mutation.SetType(pa)
	return pu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (pu *PassUpdate) SetNillableType(pa *pass.Type) *PassUpdate {
	if pa != nil {
		pu.SetType(*pa)
	}
	return pu
}

// SetStartTime sets the "start_time" field.
func (pu *PassUpdate) SetStartTime(t time.Time) *PassUpdate {
	pu.mutation.SetStartTime(t)
	return pu
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (pu *PassUpdate) SetNillableStartTime(t *time.Time) *PassUpdate {
	if t != nil {
		pu.SetStartTime(*t)
	}
	return pu
}

// SetEndTime sets the "end_time" field.
func (pu *PassUpdate) SetEndTime(t time.Time) *PassUpdate {
	pu.mutation.SetEndTime(t)
	return pu
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (pu *PassUpdate) SetNillableEndTime(t *time.Time) *PassUpdate {
	if t != nil {
		pu.SetEndTime(*t)
	}
	return pu
}

// ClearEndTime clears the value of the "end_time" field.
func (pu *PassUpdate) ClearEndTime() *PassUpdate {
	pu.mutation.ClearEndTime()
	return pu
}

// Mutation returns the PassMutation object of the builder.
func (pu *PassUpdate) Mutation() *PassMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PassUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PassUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PassUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PassUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PassUpdate) check() error {
	if v, ok := pu.mutation.GetType(); ok {
		if err := pass.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Pass.type": %w`, err)}
		}
	}
	return nil
}

func (pu *PassUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(pass.Table, pass.Columns, sqlgraph.NewFieldSpec(pass.FieldID, field.TypeUUID))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UserID(); ok {
		_spec.SetField(pass.FieldUserID, field.TypeString, value)
	}
	if value, ok := pu.mutation.GetType(); ok {
		_spec.SetField(pass.FieldType, field.TypeEnum, value)
	}
	if value, ok := pu.mutation.StartTime(); ok {
		_spec.SetField(pass.FieldStartTime, field.TypeTime, value)
	}
	if value, ok := pu.mutation.EndTime(); ok {
		_spec.SetField(pass.FieldEndTime, field.TypeTime, value)
	}
	if pu.mutation.EndTimeCleared() {
		_spec.ClearField(pass.FieldEndTime, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pass.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PassUpdateOne is the builder for updating a single Pass entity.
type PassUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PassMutation
}

// SetUserID sets the "user_id" field.
func (puo *PassUpdateOne) SetUserID(s string) *PassUpdateOne {
	puo.mutation.SetUserID(s)
	return puo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (puo *PassUpdateOne) SetNillableUserID(s *string) *PassUpdateOne {
	if s != nil {
		puo.SetUserID(*s)
	}
	return puo
}

// SetType sets the "type" field.
func (puo *PassUpdateOne) SetType(pa pass.Type) *PassUpdateOne {
	puo.mutation.SetType(pa)
	return puo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (puo *PassUpdateOne) SetNillableType(pa *pass.Type) *PassUpdateOne {
	if pa != nil {
		puo.SetType(*pa)
	}
	return puo
}

// SetStartTime sets the "start_time" field.
func (puo *PassUpdateOne) SetStartTime(t time.Time) *PassUpdateOne {
	puo.mutation.SetStartTime(t)
	return puo
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (puo *PassUpdateOne) SetNillableStartTime(t *time.Time) *PassUpdateOne {
	if t != nil {
		puo.SetStartTime(*t)
	}
	return puo
}

// SetEndTime sets the "end_time" field.
func (puo *PassUpdateOne) SetEndTime(t time.Time) *PassUpdateOne {
	puo.mutation.SetEndTime(t)
	return puo
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (puo *PassUpdateOne) SetNillableEndTime(t *time.Time) *PassUpdateOne {
	if t != nil {
		puo.SetEndTime(*t)
	}
	return puo
}

// ClearEndTime clears the value of the "end_time" field.
func (puo *PassUpdateOne) ClearEndTime() *PassUpdateOne {
	puo.mutation.ClearEndTime()
	return puo
}

// Mutation returns the PassMutation object of the builder.
func (puo *PassUpdateOne) Mutation() *PassMutation {
	return puo.mutation
}

// Where appends a list predicates to the PassUpdate builder.
func (puo *PassUpdateOne) Where(ps ...predicate.Pass) *PassUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PassUpdateOne) Select(field string, fields ...string) *PassUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Pass entity.
func (puo *PassUpdateOne) Save(ctx context.Context) (*Pass, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PassUpdateOne) SaveX(ctx context.Context) *Pass {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PassUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PassUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PassUpdateOne) check() error {
	if v, ok := puo.mutation.GetType(); ok {
		if err := pass.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Pass.type": %w`, err)}
		}
	}
	return nil
}

func (puo *PassUpdateOne) sqlSave(ctx context.Context) (_node *Pass, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(pass.Table, pass.Columns, sqlgraph.NewFieldSpec(pass.FieldID, field.TypeUUID))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Pass.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pass.FieldID)
		for _, f := range fields {
			if !pass.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != pass.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UserID(); ok {
		_spec.SetField(pass.FieldUserID, field.TypeString, value)
	}
	if value, ok := puo.mutation.GetType(); ok {
		_spec.SetField(pass.FieldType, field.TypeEnum, value)
	}
	if value, ok := puo.mutation.StartTime(); ok {
		_spec.SetField(pass.FieldStartTime, field.TypeTime, value)
	}
	if value, ok := puo.mutation.EndTime(); ok {
		_spec.SetField(pass.FieldEndTime, field.TypeTime, value)
	}
	if puo.mutation.EndTimeCleared() {
		_spec.ClearField(pass.FieldEndTime, field.TypeTime)
	}
	_node = &Pass{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pass.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
