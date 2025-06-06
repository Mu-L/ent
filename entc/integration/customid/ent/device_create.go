// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/device"
	"entgo.io/ent/entc/integration/customid/ent/schema"
	"entgo.io/ent/entc/integration/customid/ent/session"
	"entgo.io/ent/schema/field"
)

// DeviceCreate is the builder for creating a Device entity.
type DeviceCreate struct {
	config
	mutation *DeviceMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetID sets the "id" field.
func (_c *DeviceCreate) SetID(v schema.ID) *DeviceCreate {
	_c.mutation.SetID(v)
	return _c
}

// SetNillableID sets the "id" field if the given value is not nil.
func (_c *DeviceCreate) SetNillableID(v *schema.ID) *DeviceCreate {
	if v != nil {
		_c.SetID(*v)
	}
	return _c
}

// SetActiveSessionID sets the "active_session" edge to the Session entity by ID.
func (_c *DeviceCreate) SetActiveSessionID(id schema.ID) *DeviceCreate {
	_c.mutation.SetActiveSessionID(id)
	return _c
}

// SetNillableActiveSessionID sets the "active_session" edge to the Session entity by ID if the given value is not nil.
func (_c *DeviceCreate) SetNillableActiveSessionID(id *schema.ID) *DeviceCreate {
	if id != nil {
		_c = _c.SetActiveSessionID(*id)
	}
	return _c
}

// SetActiveSession sets the "active_session" edge to the Session entity.
func (_c *DeviceCreate) SetActiveSession(v *Session) *DeviceCreate {
	return _c.SetActiveSessionID(v.ID)
}

// AddSessionIDs adds the "sessions" edge to the Session entity by IDs.
func (_c *DeviceCreate) AddSessionIDs(ids ...schema.ID) *DeviceCreate {
	_c.mutation.AddSessionIDs(ids...)
	return _c
}

// AddSessions adds the "sessions" edges to the Session entity.
func (_c *DeviceCreate) AddSessions(v ...*Session) *DeviceCreate {
	ids := make([]schema.ID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _c.AddSessionIDs(ids...)
}

// Mutation returns the DeviceMutation object of the builder.
func (_c *DeviceCreate) Mutation() *DeviceMutation {
	return _c.mutation
}

// Save creates the Device in the database.
func (_c *DeviceCreate) Save(ctx context.Context) (*Device, error) {
	_c.defaults()
	return withHooks(ctx, _c.sqlSave, _c.mutation, _c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (_c *DeviceCreate) SaveX(ctx context.Context) *Device {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *DeviceCreate) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *DeviceCreate) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (_c *DeviceCreate) defaults() {
	if _, ok := _c.mutation.ID(); !ok {
		v := device.DefaultID()
		_c.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *DeviceCreate) check() error {
	if v, ok := _c.mutation.ID(); ok {
		if err := device.IDValidator(v[:]); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Device.id": %w`, err)}
		}
	}
	return nil
}

func (_c *DeviceCreate) sqlSave(ctx context.Context) (*Device, error) {
	if err := _c.check(); err != nil {
		return nil, err
	}
	_node, _spec := _c.createSpec()
	if err := sqlgraph.CreateNode(ctx, _c.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*schema.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	_c.mutation.id = &_node.ID
	_c.mutation.done = true
	return _node, nil
}

func (_c *DeviceCreate) createSpec() (*Device, *sqlgraph.CreateSpec) {
	var (
		_node = &Device{config: _c.config}
		_spec = sqlgraph.NewCreateSpec(device.Table, sqlgraph.NewFieldSpec(device.FieldID, field.TypeBytes))
	)
	_spec.OnConflict = _c.conflict
	if id, ok := _c.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if nodes := _c.mutation.ActiveSessionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   device.ActiveSessionTable,
			Columns: []string{device.ActiveSessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeBytes),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.device_active_session = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := _c.mutation.SessionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   device.SessionsTable,
			Columns: []string{device.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeBytes),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Device.Create().
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (_c *DeviceCreate) OnConflict(opts ...sql.ConflictOption) *DeviceUpsertOne {
	_c.conflict = opts
	return &DeviceUpsertOne{
		create: _c,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Device.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (_c *DeviceCreate) OnConflictColumns(columns ...string) *DeviceUpsertOne {
	_c.conflict = append(_c.conflict, sql.ConflictColumns(columns...))
	return &DeviceUpsertOne{
		create: _c,
	}
}

type (
	// DeviceUpsertOne is the builder for "upsert"-ing
	//  one Device node.
	DeviceUpsertOne struct {
		create *DeviceCreate
	}

	// DeviceUpsert is the "OnConflict" setter.
	DeviceUpsert struct {
		*sql.UpdateSet
	}
)

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Device.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(device.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DeviceUpsertOne) UpdateNewValues() *DeviceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(device.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Device.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DeviceUpsertOne) Ignore() *DeviceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeviceUpsertOne) DoNothing() *DeviceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeviceCreate.OnConflict
// documentation for more info.
func (u *DeviceUpsertOne) Update(set func(*DeviceUpsert)) *DeviceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeviceUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *DeviceUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DeviceCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeviceUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DeviceUpsertOne) ID(ctx context.Context) (id schema.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: DeviceUpsertOne.ID is not supported by MySQL driver. Use DeviceUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DeviceUpsertOne) IDX(ctx context.Context) schema.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DeviceCreateBulk is the builder for creating many Device entities in bulk.
type DeviceCreateBulk struct {
	config
	err      error
	builders []*DeviceCreate
	conflict []sql.ConflictOption
}

// Save creates the Device entities in the database.
func (_c *DeviceCreateBulk) Save(ctx context.Context) ([]*Device, error) {
	if _c.err != nil {
		return nil, _c.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(_c.builders))
	nodes := make([]*Device, len(_c.builders))
	mutators := make([]Mutator, len(_c.builders))
	for i := range _c.builders {
		func(i int, root context.Context) {
			builder := _c.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeviceMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, _c.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = _c.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, _c.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, _c.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (_c *DeviceCreateBulk) SaveX(ctx context.Context) []*Device {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *DeviceCreateBulk) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *DeviceCreateBulk) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Device.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (_c *DeviceCreateBulk) OnConflict(opts ...sql.ConflictOption) *DeviceUpsertBulk {
	_c.conflict = opts
	return &DeviceUpsertBulk{
		create: _c,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Device.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (_c *DeviceCreateBulk) OnConflictColumns(columns ...string) *DeviceUpsertBulk {
	_c.conflict = append(_c.conflict, sql.ConflictColumns(columns...))
	return &DeviceUpsertBulk{
		create: _c,
	}
}

// DeviceUpsertBulk is the builder for "upsert"-ing
// a bulk of Device nodes.
type DeviceUpsertBulk struct {
	create *DeviceCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Device.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(device.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DeviceUpsertBulk) UpdateNewValues() *DeviceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(device.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Device.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DeviceUpsertBulk) Ignore() *DeviceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeviceUpsertBulk) DoNothing() *DeviceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeviceCreateBulk.OnConflict
// documentation for more info.
func (u *DeviceUpsertBulk) Update(set func(*DeviceUpsert)) *DeviceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeviceUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *DeviceUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DeviceCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DeviceCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeviceUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
