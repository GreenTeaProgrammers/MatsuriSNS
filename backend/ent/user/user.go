// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldHashedPassword holds the string denoting the hashed_password field in the database.
	FieldHashedPassword = "hashed_password"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgePosts holds the string denoting the posts edge name in mutations.
	EdgePosts = "posts"
	// EdgeCreatedEvents holds the string denoting the created_events edge name in mutations.
	EdgeCreatedEvents = "created_events"
	// EdgeEventAdmins holds the string denoting the event_admins edge name in mutations.
	EdgeEventAdmins = "event_admins"
	// Table holds the table name of the user in the database.
	Table = "users"
	// PostsTable is the table that holds the posts relation/edge.
	PostsTable = "posts"
	// PostsInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	PostsInverseTable = "posts"
	// PostsColumn is the table column denoting the posts relation/edge.
	PostsColumn = "user_id"
	// CreatedEventsTable is the table that holds the created_events relation/edge.
	CreatedEventsTable = "events"
	// CreatedEventsInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	CreatedEventsInverseTable = "events"
	// CreatedEventsColumn is the table column denoting the created_events relation/edge.
	CreatedEventsColumn = "creator_id"
	// EventAdminsTable is the table that holds the event_admins relation/edge.
	EventAdminsTable = "event_admins"
	// EventAdminsInverseTable is the table name for the EventAdmin entity.
	// It exists in this package in order to avoid circular dependency with the "eventadmin" package.
	EventAdminsInverseTable = "event_admins"
	// EventAdminsColumn is the table column denoting the event_admins relation/edge.
	EventAdminsColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldHashedPassword,
	FieldCreatedAt,
	FieldUpdatedAt,
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

var (
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// HashedPasswordValidator is a validator for the "hashed_password" field. It is called by the builders before save.
	HashedPasswordValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByHashedPassword orders the results by the hashed_password field.
func ByHashedPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHashedPassword, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByPostsCount orders the results by posts count.
func ByPostsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPostsStep(), opts...)
	}
}

// ByPosts orders the results by posts terms.
func ByPosts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPostsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCreatedEventsCount orders the results by created_events count.
func ByCreatedEventsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCreatedEventsStep(), opts...)
	}
}

// ByCreatedEvents orders the results by created_events terms.
func ByCreatedEvents(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCreatedEventsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByEventAdminsCount orders the results by event_admins count.
func ByEventAdminsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newEventAdminsStep(), opts...)
	}
}

// ByEventAdmins orders the results by event_admins terms.
func ByEventAdmins(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEventAdminsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPostsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PostsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, PostsTable, PostsColumn),
	)
}
func newCreatedEventsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CreatedEventsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, CreatedEventsTable, CreatedEventsColumn),
	)
}
func newEventAdminsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EventAdminsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, EventAdminsTable, EventAdminsColumn),
	)
}
