// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/GreenTeaProgrammers/MatsuriSNS/ent/event"
	"github.com/GreenTeaProgrammers/MatsuriSNS/ent/eventadmin"
	"github.com/GreenTeaProgrammers/MatsuriSNS/ent/post"
	"github.com/GreenTeaProgrammers/MatsuriSNS/ent/user"
)

// EventCreate is the builder for creating a Event entity.
type EventCreate struct {
	config
	mutation *EventMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (ec *EventCreate) SetTitle(s string) *EventCreate {
	ec.mutation.SetTitle(s)
	return ec
}

// SetDescription sets the "description" field.
func (ec *EventCreate) SetDescription(s string) *EventCreate {
	ec.mutation.SetDescription(s)
	return ec
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ec *EventCreate) SetNillableDescription(s *string) *EventCreate {
	if s != nil {
		ec.SetDescription(*s)
	}
	return ec
}

// SetMapURL sets the "map_url" field.
func (ec *EventCreate) SetMapURL(s string) *EventCreate {
	ec.mutation.SetMapURL(s)
	return ec
}

// SetQrCodeURL sets the "qr_code_url" field.
func (ec *EventCreate) SetQrCodeURL(s string) *EventCreate {
	ec.mutation.SetQrCodeURL(s)
	return ec
}

// SetNillableQrCodeURL sets the "qr_code_url" field if the given value is not nil.
func (ec *EventCreate) SetNillableQrCodeURL(s *string) *EventCreate {
	if s != nil {
		ec.SetQrCodeURL(*s)
	}
	return ec
}

// SetStartTime sets the "start_time" field.
func (ec *EventCreate) SetStartTime(t time.Time) *EventCreate {
	ec.mutation.SetStartTime(t)
	return ec
}

// SetEndTime sets the "end_time" field.
func (ec *EventCreate) SetEndTime(t time.Time) *EventCreate {
	ec.mutation.SetEndTime(t)
	return ec
}

// SetCreatedAt sets the "created_at" field.
func (ec *EventCreate) SetCreatedAt(t time.Time) *EventCreate {
	ec.mutation.SetCreatedAt(t)
	return ec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ec *EventCreate) SetNillableCreatedAt(t *time.Time) *EventCreate {
	if t != nil {
		ec.SetCreatedAt(*t)
	}
	return ec
}

// SetUpdatedAt sets the "updated_at" field.
func (ec *EventCreate) SetUpdatedAt(t time.Time) *EventCreate {
	ec.mutation.SetUpdatedAt(t)
	return ec
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ec *EventCreate) SetNillableUpdatedAt(t *time.Time) *EventCreate {
	if t != nil {
		ec.SetUpdatedAt(*t)
	}
	return ec
}

// SetCreatorID sets the "creator_id" field.
func (ec *EventCreate) SetCreatorID(i int) *EventCreate {
	ec.mutation.SetCreatorID(i)
	return ec
}

// SetCreator sets the "creator" edge to the User entity.
func (ec *EventCreate) SetCreator(u *User) *EventCreate {
	return ec.SetCreatorID(u.ID)
}

// AddEventAdminIDs adds the "event_admins" edge to the EventAdmin entity by IDs.
func (ec *EventCreate) AddEventAdminIDs(ids ...int) *EventCreate {
	ec.mutation.AddEventAdminIDs(ids...)
	return ec
}

// AddEventAdmins adds the "event_admins" edges to the EventAdmin entity.
func (ec *EventCreate) AddEventAdmins(e ...*EventAdmin) *EventCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ec.AddEventAdminIDs(ids...)
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (ec *EventCreate) AddPostIDs(ids ...int) *EventCreate {
	ec.mutation.AddPostIDs(ids...)
	return ec
}

// AddPosts adds the "posts" edges to the Post entity.
func (ec *EventCreate) AddPosts(p ...*Post) *EventCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ec.AddPostIDs(ids...)
}

// Mutation returns the EventMutation object of the builder.
func (ec *EventCreate) Mutation() *EventMutation {
	return ec.mutation
}

// Save creates the Event in the database.
func (ec *EventCreate) Save(ctx context.Context) (*Event, error) {
	ec.defaults()
	return withHooks(ctx, ec.sqlSave, ec.mutation, ec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EventCreate) SaveX(ctx context.Context) *Event {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EventCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EventCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *EventCreate) defaults() {
	if _, ok := ec.mutation.CreatedAt(); !ok {
		v := event.DefaultCreatedAt()
		ec.mutation.SetCreatedAt(v)
	}
	if _, ok := ec.mutation.UpdatedAt(); !ok {
		v := event.DefaultUpdatedAt()
		ec.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EventCreate) check() error {
	if _, ok := ec.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Event.title"`)}
	}
	if v, ok := ec.mutation.Title(); ok {
		if err := event.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Event.title": %w`, err)}
		}
	}
	if _, ok := ec.mutation.MapURL(); !ok {
		return &ValidationError{Name: "map_url", err: errors.New(`ent: missing required field "Event.map_url"`)}
	}
	if v, ok := ec.mutation.MapURL(); ok {
		if err := event.MapURLValidator(v); err != nil {
			return &ValidationError{Name: "map_url", err: fmt.Errorf(`ent: validator failed for field "Event.map_url": %w`, err)}
		}
	}
	if _, ok := ec.mutation.StartTime(); !ok {
		return &ValidationError{Name: "start_time", err: errors.New(`ent: missing required field "Event.start_time"`)}
	}
	if _, ok := ec.mutation.EndTime(); !ok {
		return &ValidationError{Name: "end_time", err: errors.New(`ent: missing required field "Event.end_time"`)}
	}
	if _, ok := ec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Event.created_at"`)}
	}
	if _, ok := ec.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Event.updated_at"`)}
	}
	if _, ok := ec.mutation.CreatorID(); !ok {
		return &ValidationError{Name: "creator_id", err: errors.New(`ent: missing required field "Event.creator_id"`)}
	}
	if len(ec.mutation.CreatorIDs()) == 0 {
		return &ValidationError{Name: "creator", err: errors.New(`ent: missing required edge "Event.creator"`)}
	}
	return nil
}

func (ec *EventCreate) sqlSave(ctx context.Context) (*Event, error) {
	if err := ec.check(); err != nil {
		return nil, err
	}
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ec.mutation.id = &_node.ID
	ec.mutation.done = true
	return _node, nil
}

func (ec *EventCreate) createSpec() (*Event, *sqlgraph.CreateSpec) {
	var (
		_node = &Event{config: ec.config}
		_spec = sqlgraph.NewCreateSpec(event.Table, sqlgraph.NewFieldSpec(event.FieldID, field.TypeInt))
	)
	if value, ok := ec.mutation.Title(); ok {
		_spec.SetField(event.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := ec.mutation.Description(); ok {
		_spec.SetField(event.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ec.mutation.MapURL(); ok {
		_spec.SetField(event.FieldMapURL, field.TypeString, value)
		_node.MapURL = value
	}
	if value, ok := ec.mutation.QrCodeURL(); ok {
		_spec.SetField(event.FieldQrCodeURL, field.TypeString, value)
		_node.QrCodeURL = value
	}
	if value, ok := ec.mutation.StartTime(); ok {
		_spec.SetField(event.FieldStartTime, field.TypeTime, value)
		_node.StartTime = value
	}
	if value, ok := ec.mutation.EndTime(); ok {
		_spec.SetField(event.FieldEndTime, field.TypeTime, value)
		_node.EndTime = value
	}
	if value, ok := ec.mutation.CreatedAt(); ok {
		_spec.SetField(event.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ec.mutation.UpdatedAt(); ok {
		_spec.SetField(event.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := ec.mutation.CreatorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   event.CreatorTable,
			Columns: []string{event.CreatorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CreatorID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.EventAdminsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   event.EventAdminsTable,
			Columns: []string{event.EventAdminsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(eventadmin.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   event.PostsTable,
			Columns: []string{event.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EventCreateBulk is the builder for creating many Event entities in bulk.
type EventCreateBulk struct {
	config
	err      error
	builders []*EventCreate
}

// Save creates the Event entities in the database.
func (ecb *EventCreateBulk) Save(ctx context.Context) ([]*Event, error) {
	if ecb.err != nil {
		return nil, ecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Event, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EventMutation)
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
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EventCreateBulk) SaveX(ctx context.Context) []*Event {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EventCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EventCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}
