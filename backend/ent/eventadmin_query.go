// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/GreenTeaProgrammers/MatsuriSNS/ent/event"
	"github.com/GreenTeaProgrammers/MatsuriSNS/ent/eventadmin"
	"github.com/GreenTeaProgrammers/MatsuriSNS/ent/predicate"
	"github.com/GreenTeaProgrammers/MatsuriSNS/ent/user"
)

// EventAdminQuery is the builder for querying EventAdmin entities.
type EventAdminQuery struct {
	config
	ctx        *QueryContext
	order      []eventadmin.OrderOption
	inters     []Interceptor
	predicates []predicate.EventAdmin
	withEvent  *EventQuery
	withUser   *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EventAdminQuery builder.
func (eaq *EventAdminQuery) Where(ps ...predicate.EventAdmin) *EventAdminQuery {
	eaq.predicates = append(eaq.predicates, ps...)
	return eaq
}

// Limit the number of records to be returned by this query.
func (eaq *EventAdminQuery) Limit(limit int) *EventAdminQuery {
	eaq.ctx.Limit = &limit
	return eaq
}

// Offset to start from.
func (eaq *EventAdminQuery) Offset(offset int) *EventAdminQuery {
	eaq.ctx.Offset = &offset
	return eaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (eaq *EventAdminQuery) Unique(unique bool) *EventAdminQuery {
	eaq.ctx.Unique = &unique
	return eaq
}

// Order specifies how the records should be ordered.
func (eaq *EventAdminQuery) Order(o ...eventadmin.OrderOption) *EventAdminQuery {
	eaq.order = append(eaq.order, o...)
	return eaq
}

// QueryEvent chains the current query on the "event" edge.
func (eaq *EventAdminQuery) QueryEvent() *EventQuery {
	query := (&EventClient{config: eaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(eventadmin.Table, eventadmin.FieldID, selector),
			sqlgraph.To(event.Table, event.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, eventadmin.EventTable, eventadmin.EventColumn),
		)
		fromU = sqlgraph.SetNeighbors(eaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (eaq *EventAdminQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: eaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(eventadmin.Table, eventadmin.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, eventadmin.UserTable, eventadmin.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(eaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first EventAdmin entity from the query.
// Returns a *NotFoundError when no EventAdmin was found.
func (eaq *EventAdminQuery) First(ctx context.Context) (*EventAdmin, error) {
	nodes, err := eaq.Limit(1).All(setContextOp(ctx, eaq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{eventadmin.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (eaq *EventAdminQuery) FirstX(ctx context.Context) *EventAdmin {
	node, err := eaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EventAdmin ID from the query.
// Returns a *NotFoundError when no EventAdmin ID was found.
func (eaq *EventAdminQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = eaq.Limit(1).IDs(setContextOp(ctx, eaq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{eventadmin.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (eaq *EventAdminQuery) FirstIDX(ctx context.Context) int {
	id, err := eaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EventAdmin entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EventAdmin entity is found.
// Returns a *NotFoundError when no EventAdmin entities are found.
func (eaq *EventAdminQuery) Only(ctx context.Context) (*EventAdmin, error) {
	nodes, err := eaq.Limit(2).All(setContextOp(ctx, eaq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{eventadmin.Label}
	default:
		return nil, &NotSingularError{eventadmin.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (eaq *EventAdminQuery) OnlyX(ctx context.Context) *EventAdmin {
	node, err := eaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EventAdmin ID in the query.
// Returns a *NotSingularError when more than one EventAdmin ID is found.
// Returns a *NotFoundError when no entities are found.
func (eaq *EventAdminQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = eaq.Limit(2).IDs(setContextOp(ctx, eaq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{eventadmin.Label}
	default:
		err = &NotSingularError{eventadmin.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (eaq *EventAdminQuery) OnlyIDX(ctx context.Context) int {
	id, err := eaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EventAdmins.
func (eaq *EventAdminQuery) All(ctx context.Context) ([]*EventAdmin, error) {
	ctx = setContextOp(ctx, eaq.ctx, ent.OpQueryAll)
	if err := eaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*EventAdmin, *EventAdminQuery]()
	return withInterceptors[[]*EventAdmin](ctx, eaq, qr, eaq.inters)
}

// AllX is like All, but panics if an error occurs.
func (eaq *EventAdminQuery) AllX(ctx context.Context) []*EventAdmin {
	nodes, err := eaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EventAdmin IDs.
func (eaq *EventAdminQuery) IDs(ctx context.Context) (ids []int, err error) {
	if eaq.ctx.Unique == nil && eaq.path != nil {
		eaq.Unique(true)
	}
	ctx = setContextOp(ctx, eaq.ctx, ent.OpQueryIDs)
	if err = eaq.Select(eventadmin.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (eaq *EventAdminQuery) IDsX(ctx context.Context) []int {
	ids, err := eaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (eaq *EventAdminQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, eaq.ctx, ent.OpQueryCount)
	if err := eaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, eaq, querierCount[*EventAdminQuery](), eaq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (eaq *EventAdminQuery) CountX(ctx context.Context) int {
	count, err := eaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (eaq *EventAdminQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, eaq.ctx, ent.OpQueryExist)
	switch _, err := eaq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (eaq *EventAdminQuery) ExistX(ctx context.Context) bool {
	exist, err := eaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EventAdminQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (eaq *EventAdminQuery) Clone() *EventAdminQuery {
	if eaq == nil {
		return nil
	}
	return &EventAdminQuery{
		config:     eaq.config,
		ctx:        eaq.ctx.Clone(),
		order:      append([]eventadmin.OrderOption{}, eaq.order...),
		inters:     append([]Interceptor{}, eaq.inters...),
		predicates: append([]predicate.EventAdmin{}, eaq.predicates...),
		withEvent:  eaq.withEvent.Clone(),
		withUser:   eaq.withUser.Clone(),
		// clone intermediate query.
		sql:  eaq.sql.Clone(),
		path: eaq.path,
	}
}

// WithEvent tells the query-builder to eager-load the nodes that are connected to
// the "event" edge. The optional arguments are used to configure the query builder of the edge.
func (eaq *EventAdminQuery) WithEvent(opts ...func(*EventQuery)) *EventAdminQuery {
	query := (&EventClient{config: eaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eaq.withEvent = query
	return eaq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (eaq *EventAdminQuery) WithUser(opts ...func(*UserQuery)) *EventAdminQuery {
	query := (&UserClient{config: eaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eaq.withUser = query
	return eaq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		EventID int `json:"event_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EventAdmin.Query().
//		GroupBy(eventadmin.FieldEventID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (eaq *EventAdminQuery) GroupBy(field string, fields ...string) *EventAdminGroupBy {
	eaq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EventAdminGroupBy{build: eaq}
	grbuild.flds = &eaq.ctx.Fields
	grbuild.label = eventadmin.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		EventID int `json:"event_id,omitempty"`
//	}
//
//	client.EventAdmin.Query().
//		Select(eventadmin.FieldEventID).
//		Scan(ctx, &v)
func (eaq *EventAdminQuery) Select(fields ...string) *EventAdminSelect {
	eaq.ctx.Fields = append(eaq.ctx.Fields, fields...)
	sbuild := &EventAdminSelect{EventAdminQuery: eaq}
	sbuild.label = eventadmin.Label
	sbuild.flds, sbuild.scan = &eaq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EventAdminSelect configured with the given aggregations.
func (eaq *EventAdminQuery) Aggregate(fns ...AggregateFunc) *EventAdminSelect {
	return eaq.Select().Aggregate(fns...)
}

func (eaq *EventAdminQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range eaq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, eaq); err != nil {
				return err
			}
		}
	}
	for _, f := range eaq.ctx.Fields {
		if !eventadmin.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if eaq.path != nil {
		prev, err := eaq.path(ctx)
		if err != nil {
			return err
		}
		eaq.sql = prev
	}
	return nil
}

func (eaq *EventAdminQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EventAdmin, error) {
	var (
		nodes       = []*EventAdmin{}
		_spec       = eaq.querySpec()
		loadedTypes = [2]bool{
			eaq.withEvent != nil,
			eaq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EventAdmin).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EventAdmin{config: eaq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, eaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := eaq.withEvent; query != nil {
		if err := eaq.loadEvent(ctx, query, nodes, nil,
			func(n *EventAdmin, e *Event) { n.Edges.Event = e }); err != nil {
			return nil, err
		}
	}
	if query := eaq.withUser; query != nil {
		if err := eaq.loadUser(ctx, query, nodes, nil,
			func(n *EventAdmin, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (eaq *EventAdminQuery) loadEvent(ctx context.Context, query *EventQuery, nodes []*EventAdmin, init func(*EventAdmin), assign func(*EventAdmin, *Event)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*EventAdmin)
	for i := range nodes {
		fk := nodes[i].EventID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(event.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "event_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (eaq *EventAdminQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*EventAdmin, init func(*EventAdmin), assign func(*EventAdmin, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*EventAdmin)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (eaq *EventAdminQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := eaq.querySpec()
	_spec.Node.Columns = eaq.ctx.Fields
	if len(eaq.ctx.Fields) > 0 {
		_spec.Unique = eaq.ctx.Unique != nil && *eaq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, eaq.driver, _spec)
}

func (eaq *EventAdminQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(eventadmin.Table, eventadmin.Columns, sqlgraph.NewFieldSpec(eventadmin.FieldID, field.TypeInt))
	_spec.From = eaq.sql
	if unique := eaq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if eaq.path != nil {
		_spec.Unique = true
	}
	if fields := eaq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, eventadmin.FieldID)
		for i := range fields {
			if fields[i] != eventadmin.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if eaq.withEvent != nil {
			_spec.Node.AddColumnOnce(eventadmin.FieldEventID)
		}
		if eaq.withUser != nil {
			_spec.Node.AddColumnOnce(eventadmin.FieldUserID)
		}
	}
	if ps := eaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := eaq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := eaq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := eaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (eaq *EventAdminQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(eaq.driver.Dialect())
	t1 := builder.Table(eventadmin.Table)
	columns := eaq.ctx.Fields
	if len(columns) == 0 {
		columns = eventadmin.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if eaq.sql != nil {
		selector = eaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if eaq.ctx.Unique != nil && *eaq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range eaq.predicates {
		p(selector)
	}
	for _, p := range eaq.order {
		p(selector)
	}
	if offset := eaq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := eaq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EventAdminGroupBy is the group-by builder for EventAdmin entities.
type EventAdminGroupBy struct {
	selector
	build *EventAdminQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (eagb *EventAdminGroupBy) Aggregate(fns ...AggregateFunc) *EventAdminGroupBy {
	eagb.fns = append(eagb.fns, fns...)
	return eagb
}

// Scan applies the selector query and scans the result into the given value.
func (eagb *EventAdminGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, eagb.build.ctx, ent.OpQueryGroupBy)
	if err := eagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EventAdminQuery, *EventAdminGroupBy](ctx, eagb.build, eagb, eagb.build.inters, v)
}

func (eagb *EventAdminGroupBy) sqlScan(ctx context.Context, root *EventAdminQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(eagb.fns))
	for _, fn := range eagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*eagb.flds)+len(eagb.fns))
		for _, f := range *eagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*eagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := eagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EventAdminSelect is the builder for selecting fields of EventAdmin entities.
type EventAdminSelect struct {
	*EventAdminQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (eas *EventAdminSelect) Aggregate(fns ...AggregateFunc) *EventAdminSelect {
	eas.fns = append(eas.fns, fns...)
	return eas
}

// Scan applies the selector query and scans the result into the given value.
func (eas *EventAdminSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, eas.ctx, ent.OpQuerySelect)
	if err := eas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EventAdminQuery, *EventAdminSelect](ctx, eas.EventAdminQuery, eas, eas.inters, v)
}

func (eas *EventAdminSelect) sqlScan(ctx context.Context, root *EventAdminQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(eas.fns))
	for _, fn := range eas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*eas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := eas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
