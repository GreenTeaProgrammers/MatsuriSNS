// Code generated by ent, DO NOT EDIT.

package event

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/GreenTeaProgrammers/MatsuriSNS/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldTitle, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldDescription, v))
}

// MapURL applies equality check predicate on the "map_url" field. It's identical to MapURLEQ.
func MapURL(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldMapURL, v))
}

// QrCodeURL applies equality check predicate on the "qr_code_url" field. It's identical to QrCodeURLEQ.
func QrCodeURL(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldQrCodeURL, v))
}

// StartTime applies equality check predicate on the "start_time" field. It's identical to StartTimeEQ.
func StartTime(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldStartTime, v))
}

// EndTime applies equality check predicate on the "end_time" field. It's identical to EndTimeEQ.
func EndTime(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldEndTime, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatorID applies equality check predicate on the "creator_id" field. It's identical to CreatorIDEQ.
func CreatorID(v int) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldCreatorID, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldTitle, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Event {
	return predicate.Event(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Event {
	return predicate.Event(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldDescription, v))
}

// MapURLEQ applies the EQ predicate on the "map_url" field.
func MapURLEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldMapURL, v))
}

// MapURLNEQ applies the NEQ predicate on the "map_url" field.
func MapURLNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldMapURL, v))
}

// MapURLIn applies the In predicate on the "map_url" field.
func MapURLIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldMapURL, vs...))
}

// MapURLNotIn applies the NotIn predicate on the "map_url" field.
func MapURLNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldMapURL, vs...))
}

// MapURLGT applies the GT predicate on the "map_url" field.
func MapURLGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldMapURL, v))
}

// MapURLGTE applies the GTE predicate on the "map_url" field.
func MapURLGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldMapURL, v))
}

// MapURLLT applies the LT predicate on the "map_url" field.
func MapURLLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldMapURL, v))
}

// MapURLLTE applies the LTE predicate on the "map_url" field.
func MapURLLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldMapURL, v))
}

// MapURLContains applies the Contains predicate on the "map_url" field.
func MapURLContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldMapURL, v))
}

// MapURLHasPrefix applies the HasPrefix predicate on the "map_url" field.
func MapURLHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldMapURL, v))
}

// MapURLHasSuffix applies the HasSuffix predicate on the "map_url" field.
func MapURLHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldMapURL, v))
}

// MapURLEqualFold applies the EqualFold predicate on the "map_url" field.
func MapURLEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldMapURL, v))
}

// MapURLContainsFold applies the ContainsFold predicate on the "map_url" field.
func MapURLContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldMapURL, v))
}

// QrCodeURLEQ applies the EQ predicate on the "qr_code_url" field.
func QrCodeURLEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldQrCodeURL, v))
}

// QrCodeURLNEQ applies the NEQ predicate on the "qr_code_url" field.
func QrCodeURLNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldQrCodeURL, v))
}

// QrCodeURLIn applies the In predicate on the "qr_code_url" field.
func QrCodeURLIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldQrCodeURL, vs...))
}

// QrCodeURLNotIn applies the NotIn predicate on the "qr_code_url" field.
func QrCodeURLNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldQrCodeURL, vs...))
}

// QrCodeURLGT applies the GT predicate on the "qr_code_url" field.
func QrCodeURLGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldQrCodeURL, v))
}

// QrCodeURLGTE applies the GTE predicate on the "qr_code_url" field.
func QrCodeURLGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldQrCodeURL, v))
}

// QrCodeURLLT applies the LT predicate on the "qr_code_url" field.
func QrCodeURLLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldQrCodeURL, v))
}

// QrCodeURLLTE applies the LTE predicate on the "qr_code_url" field.
func QrCodeURLLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldQrCodeURL, v))
}

// QrCodeURLContains applies the Contains predicate on the "qr_code_url" field.
func QrCodeURLContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldQrCodeURL, v))
}

// QrCodeURLHasPrefix applies the HasPrefix predicate on the "qr_code_url" field.
func QrCodeURLHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldQrCodeURL, v))
}

// QrCodeURLHasSuffix applies the HasSuffix predicate on the "qr_code_url" field.
func QrCodeURLHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldQrCodeURL, v))
}

// QrCodeURLIsNil applies the IsNil predicate on the "qr_code_url" field.
func QrCodeURLIsNil() predicate.Event {
	return predicate.Event(sql.FieldIsNull(FieldQrCodeURL))
}

// QrCodeURLNotNil applies the NotNil predicate on the "qr_code_url" field.
func QrCodeURLNotNil() predicate.Event {
	return predicate.Event(sql.FieldNotNull(FieldQrCodeURL))
}

// QrCodeURLEqualFold applies the EqualFold predicate on the "qr_code_url" field.
func QrCodeURLEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldQrCodeURL, v))
}

// QrCodeURLContainsFold applies the ContainsFold predicate on the "qr_code_url" field.
func QrCodeURLContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldQrCodeURL, v))
}

// StartTimeEQ applies the EQ predicate on the "start_time" field.
func StartTimeEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldStartTime, v))
}

// StartTimeNEQ applies the NEQ predicate on the "start_time" field.
func StartTimeNEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldStartTime, v))
}

// StartTimeIn applies the In predicate on the "start_time" field.
func StartTimeIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldStartTime, vs...))
}

// StartTimeNotIn applies the NotIn predicate on the "start_time" field.
func StartTimeNotIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldStartTime, vs...))
}

// StartTimeGT applies the GT predicate on the "start_time" field.
func StartTimeGT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldStartTime, v))
}

// StartTimeGTE applies the GTE predicate on the "start_time" field.
func StartTimeGTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldStartTime, v))
}

// StartTimeLT applies the LT predicate on the "start_time" field.
func StartTimeLT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldStartTime, v))
}

// StartTimeLTE applies the LTE predicate on the "start_time" field.
func StartTimeLTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldStartTime, v))
}

// EndTimeEQ applies the EQ predicate on the "end_time" field.
func EndTimeEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldEndTime, v))
}

// EndTimeNEQ applies the NEQ predicate on the "end_time" field.
func EndTimeNEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldEndTime, v))
}

// EndTimeIn applies the In predicate on the "end_time" field.
func EndTimeIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldEndTime, vs...))
}

// EndTimeNotIn applies the NotIn predicate on the "end_time" field.
func EndTimeNotIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldEndTime, vs...))
}

// EndTimeGT applies the GT predicate on the "end_time" field.
func EndTimeGT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldEndTime, v))
}

// EndTimeGTE applies the GTE predicate on the "end_time" field.
func EndTimeGTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldEndTime, v))
}

// EndTimeLT applies the LT predicate on the "end_time" field.
func EndTimeLT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldEndTime, v))
}

// EndTimeLTE applies the LTE predicate on the "end_time" field.
func EndTimeLTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldEndTime, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatorIDEQ applies the EQ predicate on the "creator_id" field.
func CreatorIDEQ(v int) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldCreatorID, v))
}

// CreatorIDNEQ applies the NEQ predicate on the "creator_id" field.
func CreatorIDNEQ(v int) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldCreatorID, v))
}

// CreatorIDIn applies the In predicate on the "creator_id" field.
func CreatorIDIn(vs ...int) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldCreatorID, vs...))
}

// CreatorIDNotIn applies the NotIn predicate on the "creator_id" field.
func CreatorIDNotIn(vs ...int) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldCreatorID, vs...))
}

// HasCreator applies the HasEdge predicate on the "creator" edge.
func HasCreator() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CreatorTable, CreatorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCreatorWith applies the HasEdge predicate on the "creator" edge with a given conditions (other predicates).
func HasCreatorWith(preds ...predicate.User) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := newCreatorStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEventAdmins applies the HasEdge predicate on the "event_admins" edge.
func HasEventAdmins() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, EventAdminsTable, EventAdminsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEventAdminsWith applies the HasEdge predicate on the "event_admins" edge with a given conditions (other predicates).
func HasEventAdminsWith(preds ...predicate.EventAdmin) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := newEventAdminsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPosts applies the HasEdge predicate on the "posts" edge.
func HasPosts() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, PostsTable, PostsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPostsWith applies the HasEdge predicate on the "posts" edge with a given conditions (other predicates).
func HasPostsWith(preds ...predicate.Post) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := newPostsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Event) predicate.Event {
	return predicate.Event(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Event) predicate.Event {
	return predicate.Event(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Event) predicate.Event {
	return predicate.Event(sql.NotPredicates(p))
}
