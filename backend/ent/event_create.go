// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/GreenTeaProgrammers/MatsuriSNS/ent/event"
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

// SetCreatedByID sets the "created_by" edge to the User entity by ID.
func (ec *EventCreate) SetCreatedByID(id int) *EventCreate {
	ec.mutation.SetCreatedByID(id)
	return ec
}

// SetNillableCreatedByID sets the "created_by" edge to the User entity by ID if the given value is not nil.
func (ec *EventCreate) SetNillableCreatedByID(id *int) *EventCreate {
	if id != nil {
		ec = ec.SetCreatedByID(*id)
	}
	return ec
}

// SetCreatedBy sets the "created_by" edge to the User entity.
func (ec *EventCreate) SetCreatedBy(u *User) *EventCreate {
	return ec.SetCreatedByID(u.ID)
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
	if nodes := ec.mutation.CreatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   event.CreatedByTable,
			Columns: []string{event.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_events = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   event.PostsTable,
			Columns: event.PostsPrimaryKey,
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
