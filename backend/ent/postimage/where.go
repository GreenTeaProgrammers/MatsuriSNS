// Code generated by ent, DO NOT EDIT.

package postimage

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/GreenTeaProgrammers/MatsuriSNS/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.PostImage {
	return predicate.PostImage(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PostImage {
	return predicate.PostImage(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PostImage {
	return predicate.PostImage(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PostImage {
	return predicate.PostImage(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.PostImage {
	return predicate.PostImage(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.PostImage {
	return predicate.PostImage(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PostImage {
	return predicate.PostImage(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PostImage {
	return predicate.PostImage(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PostImage {
	return predicate.PostImage(sql.FieldLTE(FieldID, id))
}

// PostID applies equality check predicate on the "post_id" field. It's identical to PostIDEQ.
func PostID(v int) predicate.PostImage {
	return predicate.PostImage(sql.FieldEQ(FieldPostID, v))
}

// ImageURL applies equality check predicate on the "image_url" field. It's identical to ImageURLEQ.
func ImageURL(v string) predicate.PostImage {
	return predicate.PostImage(sql.FieldEQ(FieldImageURL, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.PostImage {
	return predicate.PostImage(sql.FieldEQ(FieldCreatedAt, v))
}

// PostIDEQ applies the EQ predicate on the "post_id" field.
func PostIDEQ(v int) predicate.PostImage {
	return predicate.PostImage(sql.FieldEQ(FieldPostID, v))
}

// PostIDNEQ applies the NEQ predicate on the "post_id" field.
func PostIDNEQ(v int) predicate.PostImage {
	return predicate.PostImage(sql.FieldNEQ(FieldPostID, v))
}

// PostIDIn applies the In predicate on the "post_id" field.
func PostIDIn(vs ...int) predicate.PostImage {
	return predicate.PostImage(sql.FieldIn(FieldPostID, vs...))
}

// PostIDNotIn applies the NotIn predicate on the "post_id" field.
func PostIDNotIn(vs ...int) predicate.PostImage {
	return predicate.PostImage(sql.FieldNotIn(FieldPostID, vs...))
}

// PostIDGT applies the GT predicate on the "post_id" field.
func PostIDGT(v int) predicate.PostImage {
	return predicate.PostImage(sql.FieldGT(FieldPostID, v))
}

// PostIDGTE applies the GTE predicate on the "post_id" field.
func PostIDGTE(v int) predicate.PostImage {
	return predicate.PostImage(sql.FieldGTE(FieldPostID, v))
}

// PostIDLT applies the LT predicate on the "post_id" field.
func PostIDLT(v int) predicate.PostImage {
	return predicate.PostImage(sql.FieldLT(FieldPostID, v))
}

// PostIDLTE applies the LTE predicate on the "post_id" field.
func PostIDLTE(v int) predicate.PostImage {
	return predicate.PostImage(sql.FieldLTE(FieldPostID, v))
}

// ImageURLEQ applies the EQ predicate on the "image_url" field.
func ImageURLEQ(v string) predicate.PostImage {
	return predicate.PostImage(sql.FieldEQ(FieldImageURL, v))
}

// ImageURLNEQ applies the NEQ predicate on the "image_url" field.
func ImageURLNEQ(v string) predicate.PostImage {
	return predicate.PostImage(sql.FieldNEQ(FieldImageURL, v))
}

// ImageURLIn applies the In predicate on the "image_url" field.
func ImageURLIn(vs ...string) predicate.PostImage {
	return predicate.PostImage(sql.FieldIn(FieldImageURL, vs...))
}

// ImageURLNotIn applies the NotIn predicate on the "image_url" field.
func ImageURLNotIn(vs ...string) predicate.PostImage {
	return predicate.PostImage(sql.FieldNotIn(FieldImageURL, vs...))
}

// ImageURLGT applies the GT predicate on the "image_url" field.
func ImageURLGT(v string) predicate.PostImage {
	return predicate.PostImage(sql.FieldGT(FieldImageURL, v))
}

// ImageURLGTE applies the GTE predicate on the "image_url" field.
func ImageURLGTE(v string) predicate.PostImage {
	return predicate.PostImage(sql.FieldGTE(FieldImageURL, v))
}

// ImageURLLT applies the LT predicate on the "image_url" field.
func ImageURLLT(v string) predicate.PostImage {
	return predicate.PostImage(sql.FieldLT(FieldImageURL, v))
}

// ImageURLLTE applies the LTE predicate on the "image_url" field.
func ImageURLLTE(v string) predicate.PostImage {
	return predicate.PostImage(sql.FieldLTE(FieldImageURL, v))
}

// ImageURLContains applies the Contains predicate on the "image_url" field.
func ImageURLContains(v string) predicate.PostImage {
	return predicate.PostImage(sql.FieldContains(FieldImageURL, v))
}

// ImageURLHasPrefix applies the HasPrefix predicate on the "image_url" field.
func ImageURLHasPrefix(v string) predicate.PostImage {
	return predicate.PostImage(sql.FieldHasPrefix(FieldImageURL, v))
}

// ImageURLHasSuffix applies the HasSuffix predicate on the "image_url" field.
func ImageURLHasSuffix(v string) predicate.PostImage {
	return predicate.PostImage(sql.FieldHasSuffix(FieldImageURL, v))
}

// ImageURLEqualFold applies the EqualFold predicate on the "image_url" field.
func ImageURLEqualFold(v string) predicate.PostImage {
	return predicate.PostImage(sql.FieldEqualFold(FieldImageURL, v))
}

// ImageURLContainsFold applies the ContainsFold predicate on the "image_url" field.
func ImageURLContainsFold(v string) predicate.PostImage {
	return predicate.PostImage(sql.FieldContainsFold(FieldImageURL, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.PostImage {
	return predicate.PostImage(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.PostImage {
	return predicate.PostImage(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.PostImage {
	return predicate.PostImage(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.PostImage {
	return predicate.PostImage(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.PostImage {
	return predicate.PostImage(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.PostImage {
	return predicate.PostImage(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.PostImage {
	return predicate.PostImage(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.PostImage {
	return predicate.PostImage(sql.FieldLTE(FieldCreatedAt, v))
}

// HasPost applies the HasEdge predicate on the "post" edge.
func HasPost() predicate.PostImage {
	return predicate.PostImage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, PostTable, PostColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPostWith applies the HasEdge predicate on the "post" edge with a given conditions (other predicates).
func HasPostWith(preds ...predicate.Post) predicate.PostImage {
	return predicate.PostImage(func(s *sql.Selector) {
		step := newPostStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PostImage) predicate.PostImage {
	return predicate.PostImage(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PostImage) predicate.PostImage {
	return predicate.PostImage(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PostImage) predicate.PostImage {
	return predicate.PostImage(sql.NotPredicates(p))
}
