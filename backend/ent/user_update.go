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
	"github.com/GreenTeaProgrammers/MatsuriSNS/ent/event"
	"github.com/GreenTeaProgrammers/MatsuriSNS/ent/eventadmin"
	"github.com/GreenTeaProgrammers/MatsuriSNS/ent/post"
	"github.com/GreenTeaProgrammers/MatsuriSNS/ent/predicate"
	"github.com/GreenTeaProgrammers/MatsuriSNS/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUsername sets the "username" field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUsername(s *string) *UserUpdate {
	if s != nil {
		uu.SetUsername(*s)
	}
	return uu
}

// SetHashedPassword sets the "hashed_password" field.
func (uu *UserUpdate) SetHashedPassword(s string) *UserUpdate {
	uu.mutation.SetHashedPassword(s)
	return uu
}

// SetNillableHashedPassword sets the "hashed_password" field if the given value is not nil.
func (uu *UserUpdate) SetNillableHashedPassword(s *string) *UserUpdate {
	if s != nil {
		uu.SetHashedPassword(*s)
	}
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (uu *UserUpdate) AddPostIDs(ids ...int) *UserUpdate {
	uu.mutation.AddPostIDs(ids...)
	return uu
}

// AddPosts adds the "posts" edges to the Post entity.
func (uu *UserUpdate) AddPosts(p ...*Post) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.AddPostIDs(ids...)
}

// AddCreatedEventIDs adds the "created_events" edge to the Event entity by IDs.
func (uu *UserUpdate) AddCreatedEventIDs(ids ...int) *UserUpdate {
	uu.mutation.AddCreatedEventIDs(ids...)
	return uu
}

// AddCreatedEvents adds the "created_events" edges to the Event entity.
func (uu *UserUpdate) AddCreatedEvents(e ...*Event) *UserUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uu.AddCreatedEventIDs(ids...)
}

// AddEventAdminIDs adds the "event_admins" edge to the EventAdmin entity by IDs.
func (uu *UserUpdate) AddEventAdminIDs(ids ...int) *UserUpdate {
	uu.mutation.AddEventAdminIDs(ids...)
	return uu
}

// AddEventAdmins adds the "event_admins" edges to the EventAdmin entity.
func (uu *UserUpdate) AddEventAdmins(e ...*EventAdmin) *UserUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uu.AddEventAdminIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearPosts clears all "posts" edges to the Post entity.
func (uu *UserUpdate) ClearPosts() *UserUpdate {
	uu.mutation.ClearPosts()
	return uu
}

// RemovePostIDs removes the "posts" edge to Post entities by IDs.
func (uu *UserUpdate) RemovePostIDs(ids ...int) *UserUpdate {
	uu.mutation.RemovePostIDs(ids...)
	return uu
}

// RemovePosts removes "posts" edges to Post entities.
func (uu *UserUpdate) RemovePosts(p ...*Post) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.RemovePostIDs(ids...)
}

// ClearCreatedEvents clears all "created_events" edges to the Event entity.
func (uu *UserUpdate) ClearCreatedEvents() *UserUpdate {
	uu.mutation.ClearCreatedEvents()
	return uu
}

// RemoveCreatedEventIDs removes the "created_events" edge to Event entities by IDs.
func (uu *UserUpdate) RemoveCreatedEventIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveCreatedEventIDs(ids...)
	return uu
}

// RemoveCreatedEvents removes "created_events" edges to Event entities.
func (uu *UserUpdate) RemoveCreatedEvents(e ...*Event) *UserUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uu.RemoveCreatedEventIDs(ids...)
}

// ClearEventAdmins clears all "event_admins" edges to the EventAdmin entity.
func (uu *UserUpdate) ClearEventAdmins() *UserUpdate {
	uu.mutation.ClearEventAdmins()
	return uu
}

// RemoveEventAdminIDs removes the "event_admins" edge to EventAdmin entities by IDs.
func (uu *UserUpdate) RemoveEventAdminIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveEventAdminIDs(ids...)
	return uu
}

// RemoveEventAdmins removes "event_admins" edges to EventAdmin entities.
func (uu *UserUpdate) RemoveEventAdmins(e ...*EventAdmin) *UserUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uu.RemoveEventAdminIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	uu.defaults()
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Username(); ok {
		if err := user.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "User.username": %w`, err)}
		}
	}
	if v, ok := uu.mutation.HashedPassword(); ok {
		if err := user.HashedPasswordValidator(v); err != nil {
			return &ValidationError{Name: "hashed_password", err: fmt.Errorf(`ent: validator failed for field "User.hashed_password": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uu.mutation.HashedPassword(); ok {
		_spec.SetField(user.FieldHashedPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if uu.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.PostsTable,
			Columns: []string{user.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedPostsIDs(); len(nodes) > 0 && !uu.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.PostsTable,
			Columns: []string{user.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.PostsTable,
			Columns: []string{user.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.CreatedEventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.CreatedEventsTable,
			Columns: []string{user.CreatedEventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedCreatedEventsIDs(); len(nodes) > 0 && !uu.mutation.CreatedEventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.CreatedEventsTable,
			Columns: []string{user.CreatedEventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.CreatedEventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.CreatedEventsTable,
			Columns: []string{user.CreatedEventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.EventAdminsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.EventAdminsTable,
			Columns: []string{user.EventAdminsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(eventadmin.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedEventAdminsIDs(); len(nodes) > 0 && !uu.mutation.EventAdminsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.EventAdminsTable,
			Columns: []string{user.EventAdminsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(eventadmin.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.EventAdminsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.EventAdminsTable,
			Columns: []string{user.EventAdminsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(eventadmin.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUsername sets the "username" field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUsername(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetUsername(*s)
	}
	return uuo
}

// SetHashedPassword sets the "hashed_password" field.
func (uuo *UserUpdateOne) SetHashedPassword(s string) *UserUpdateOne {
	uuo.mutation.SetHashedPassword(s)
	return uuo
}

// SetNillableHashedPassword sets the "hashed_password" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableHashedPassword(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetHashedPassword(*s)
	}
	return uuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (uuo *UserUpdateOne) AddPostIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddPostIDs(ids...)
	return uuo
}

// AddPosts adds the "posts" edges to the Post entity.
func (uuo *UserUpdateOne) AddPosts(p ...*Post) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.AddPostIDs(ids...)
}

// AddCreatedEventIDs adds the "created_events" edge to the Event entity by IDs.
func (uuo *UserUpdateOne) AddCreatedEventIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddCreatedEventIDs(ids...)
	return uuo
}

// AddCreatedEvents adds the "created_events" edges to the Event entity.
func (uuo *UserUpdateOne) AddCreatedEvents(e ...*Event) *UserUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uuo.AddCreatedEventIDs(ids...)
}

// AddEventAdminIDs adds the "event_admins" edge to the EventAdmin entity by IDs.
func (uuo *UserUpdateOne) AddEventAdminIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddEventAdminIDs(ids...)
	return uuo
}

// AddEventAdmins adds the "event_admins" edges to the EventAdmin entity.
func (uuo *UserUpdateOne) AddEventAdmins(e ...*EventAdmin) *UserUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uuo.AddEventAdminIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearPosts clears all "posts" edges to the Post entity.
func (uuo *UserUpdateOne) ClearPosts() *UserUpdateOne {
	uuo.mutation.ClearPosts()
	return uuo
}

// RemovePostIDs removes the "posts" edge to Post entities by IDs.
func (uuo *UserUpdateOne) RemovePostIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemovePostIDs(ids...)
	return uuo
}

// RemovePosts removes "posts" edges to Post entities.
func (uuo *UserUpdateOne) RemovePosts(p ...*Post) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.RemovePostIDs(ids...)
}

// ClearCreatedEvents clears all "created_events" edges to the Event entity.
func (uuo *UserUpdateOne) ClearCreatedEvents() *UserUpdateOne {
	uuo.mutation.ClearCreatedEvents()
	return uuo
}

// RemoveCreatedEventIDs removes the "created_events" edge to Event entities by IDs.
func (uuo *UserUpdateOne) RemoveCreatedEventIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveCreatedEventIDs(ids...)
	return uuo
}

// RemoveCreatedEvents removes "created_events" edges to Event entities.
func (uuo *UserUpdateOne) RemoveCreatedEvents(e ...*Event) *UserUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uuo.RemoveCreatedEventIDs(ids...)
}

// ClearEventAdmins clears all "event_admins" edges to the EventAdmin entity.
func (uuo *UserUpdateOne) ClearEventAdmins() *UserUpdateOne {
	uuo.mutation.ClearEventAdmins()
	return uuo
}

// RemoveEventAdminIDs removes the "event_admins" edge to EventAdmin entities by IDs.
func (uuo *UserUpdateOne) RemoveEventAdminIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveEventAdminIDs(ids...)
	return uuo
}

// RemoveEventAdmins removes "event_admins" edges to EventAdmin entities.
func (uuo *UserUpdateOne) RemoveEventAdmins(e ...*EventAdmin) *UserUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uuo.RemoveEventAdminIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	uuo.defaults()
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Username(); ok {
		if err := user.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "User.username": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.HashedPassword(); ok {
		if err := user.HashedPasswordValidator(v); err != nil {
			return &ValidationError{Name: "hashed_password", err: fmt.Errorf(`ent: validator failed for field "User.hashed_password": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uuo.mutation.HashedPassword(); ok {
		_spec.SetField(user.FieldHashedPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if uuo.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.PostsTable,
			Columns: []string{user.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedPostsIDs(); len(nodes) > 0 && !uuo.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.PostsTable,
			Columns: []string{user.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.PostsTable,
			Columns: []string{user.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.CreatedEventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.CreatedEventsTable,
			Columns: []string{user.CreatedEventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedCreatedEventsIDs(); len(nodes) > 0 && !uuo.mutation.CreatedEventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.CreatedEventsTable,
			Columns: []string{user.CreatedEventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.CreatedEventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.CreatedEventsTable,
			Columns: []string{user.CreatedEventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.EventAdminsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.EventAdminsTable,
			Columns: []string{user.EventAdminsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(eventadmin.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedEventAdminsIDs(); len(nodes) > 0 && !uuo.mutation.EventAdminsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.EventAdminsTable,
			Columns: []string{user.EventAdminsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(eventadmin.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.EventAdminsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.EventAdminsTable,
			Columns: []string{user.EventAdminsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(eventadmin.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
