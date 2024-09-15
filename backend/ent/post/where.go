// Code generated by ent, DO NOT EDIT.

package post

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/GreenTeaProgrammers/MatsuriSNS/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldUserID, v))
}

// EventID applies equality check predicate on the "event_id" field. It's identical to EventIDEQ.
func EventID(v int) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldEventID, v))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldContent, v))
}

// LocationX applies equality check predicate on the "location_x" field. It's identical to LocationXEQ.
func LocationX(v float64) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldLocationX, v))
}

// LocationY applies equality check predicate on the "location_y" field. It's identical to LocationYEQ.
func LocationY(v float64) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldLocationY, v))
}

// VideoURL applies equality check predicate on the "video_url" field. It's identical to VideoURLEQ.
func VideoURL(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldVideoURL, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldUpdatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldUserID, vs...))
}

// EventIDEQ applies the EQ predicate on the "event_id" field.
func EventIDEQ(v int) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldEventID, v))
}

// EventIDNEQ applies the NEQ predicate on the "event_id" field.
func EventIDNEQ(v int) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldEventID, v))
}

// EventIDIn applies the In predicate on the "event_id" field.
func EventIDIn(vs ...int) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldEventID, vs...))
}

// EventIDNotIn applies the NotIn predicate on the "event_id" field.
func EventIDNotIn(vs ...int) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldEventID, vs...))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Post {
	return predicate.Post(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldContent, v))
}

// LocationXEQ applies the EQ predicate on the "location_x" field.
func LocationXEQ(v float64) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldLocationX, v))
}

// LocationXNEQ applies the NEQ predicate on the "location_x" field.
func LocationXNEQ(v float64) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldLocationX, v))
}

// LocationXIn applies the In predicate on the "location_x" field.
func LocationXIn(vs ...float64) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldLocationX, vs...))
}

// LocationXNotIn applies the NotIn predicate on the "location_x" field.
func LocationXNotIn(vs ...float64) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldLocationX, vs...))
}

// LocationXGT applies the GT predicate on the "location_x" field.
func LocationXGT(v float64) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldLocationX, v))
}

// LocationXGTE applies the GTE predicate on the "location_x" field.
func LocationXGTE(v float64) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldLocationX, v))
}

// LocationXLT applies the LT predicate on the "location_x" field.
func LocationXLT(v float64) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldLocationX, v))
}

// LocationXLTE applies the LTE predicate on the "location_x" field.
func LocationXLTE(v float64) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldLocationX, v))
}

// LocationYEQ applies the EQ predicate on the "location_y" field.
func LocationYEQ(v float64) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldLocationY, v))
}

// LocationYNEQ applies the NEQ predicate on the "location_y" field.
func LocationYNEQ(v float64) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldLocationY, v))
}

// LocationYIn applies the In predicate on the "location_y" field.
func LocationYIn(vs ...float64) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldLocationY, vs...))
}

// LocationYNotIn applies the NotIn predicate on the "location_y" field.
func LocationYNotIn(vs ...float64) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldLocationY, vs...))
}

// LocationYGT applies the GT predicate on the "location_y" field.
func LocationYGT(v float64) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldLocationY, v))
}

// LocationYGTE applies the GTE predicate on the "location_y" field.
func LocationYGTE(v float64) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldLocationY, v))
}

// LocationYLT applies the LT predicate on the "location_y" field.
func LocationYLT(v float64) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldLocationY, v))
}

// LocationYLTE applies the LTE predicate on the "location_y" field.
func LocationYLTE(v float64) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldLocationY, v))
}

// VideoURLEQ applies the EQ predicate on the "video_url" field.
func VideoURLEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldVideoURL, v))
}

// VideoURLNEQ applies the NEQ predicate on the "video_url" field.
func VideoURLNEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldVideoURL, v))
}

// VideoURLIn applies the In predicate on the "video_url" field.
func VideoURLIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldVideoURL, vs...))
}

// VideoURLNotIn applies the NotIn predicate on the "video_url" field.
func VideoURLNotIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldVideoURL, vs...))
}

// VideoURLGT applies the GT predicate on the "video_url" field.
func VideoURLGT(v string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldVideoURL, v))
}

// VideoURLGTE applies the GTE predicate on the "video_url" field.
func VideoURLGTE(v string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldVideoURL, v))
}

// VideoURLLT applies the LT predicate on the "video_url" field.
func VideoURLLT(v string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldVideoURL, v))
}

// VideoURLLTE applies the LTE predicate on the "video_url" field.
func VideoURLLTE(v string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldVideoURL, v))
}

// VideoURLContains applies the Contains predicate on the "video_url" field.
func VideoURLContains(v string) predicate.Post {
	return predicate.Post(sql.FieldContains(FieldVideoURL, v))
}

// VideoURLHasPrefix applies the HasPrefix predicate on the "video_url" field.
func VideoURLHasPrefix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasPrefix(FieldVideoURL, v))
}

// VideoURLHasSuffix applies the HasSuffix predicate on the "video_url" field.
func VideoURLHasSuffix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasSuffix(FieldVideoURL, v))
}

// VideoURLIsNil applies the IsNil predicate on the "video_url" field.
func VideoURLIsNil() predicate.Post {
	return predicate.Post(sql.FieldIsNull(FieldVideoURL))
}

// VideoURLNotNil applies the NotNil predicate on the "video_url" field.
func VideoURLNotNil() predicate.Post {
	return predicate.Post(sql.FieldNotNull(FieldVideoURL))
}

// VideoURLEqualFold applies the EqualFold predicate on the "video_url" field.
func VideoURLEqualFold(v string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldVideoURL, v))
}

// VideoURLContainsFold applies the ContainsFold predicate on the "video_url" field.
func VideoURLContainsFold(v string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldVideoURL, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEvent applies the HasEdge predicate on the "event" edge.
func HasEvent() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, EventTable, EventColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEventWith applies the HasEdge predicate on the "event" edge with a given conditions (other predicates).
func HasEventWith(preds ...predicate.Event) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := newEventStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasImages applies the HasEdge predicate on the "images" edge.
func HasImages() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, ImagesTable, ImagesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasImagesWith applies the HasEdge predicate on the "images" edge with a given conditions (other predicates).
func HasImagesWith(preds ...predicate.PostImage) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := newImagesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Post) predicate.Post {
	return predicate.Post(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Post) predicate.Post {
	return predicate.Post(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Post) predicate.Post {
	return predicate.Post(sql.NotPredicates(p))
}
