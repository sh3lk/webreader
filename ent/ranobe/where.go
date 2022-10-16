// Code generated by ent, DO NOT EDIT.

package ranobe

import (
	"time"
	"webreader/ent/predicate"
	"webreader/ent/schema/ulid"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id ulid.ID) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id ulid.ID) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id ulid.ID) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...ulid.ID) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...ulid.ID) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id ulid.ID) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id ulid.ID) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id ulid.ID) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id ulid.ID) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// EngName applies equality check predicate on the "engName" field. It's identical to EngNameEQ.
func EngName(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEngName), v))
	})
}

// RusName applies equality check predicate on the "rusName" field. It's identical to RusNameEQ.
func RusName(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRusName), v))
	})
}

// Summary applies equality check predicate on the "summary" field. It's identical to SummaryEQ.
func Summary(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSummary), v))
	})
}

// ReleaseDate applies equality check predicate on the "releaseDate" field. It's identical to ReleaseDateEQ.
func ReleaseDate(v int) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReleaseDate), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Ranobe {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Ranobe {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Ranobe {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Ranobe {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Ranobe {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Ranobe {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// EngNameEQ applies the EQ predicate on the "engName" field.
func EngNameEQ(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEngName), v))
	})
}

// EngNameNEQ applies the NEQ predicate on the "engName" field.
func EngNameNEQ(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEngName), v))
	})
}

// EngNameIn applies the In predicate on the "engName" field.
func EngNameIn(vs ...string) predicate.Ranobe {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEngName), v...))
	})
}

// EngNameNotIn applies the NotIn predicate on the "engName" field.
func EngNameNotIn(vs ...string) predicate.Ranobe {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEngName), v...))
	})
}

// EngNameGT applies the GT predicate on the "engName" field.
func EngNameGT(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEngName), v))
	})
}

// EngNameGTE applies the GTE predicate on the "engName" field.
func EngNameGTE(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEngName), v))
	})
}

// EngNameLT applies the LT predicate on the "engName" field.
func EngNameLT(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEngName), v))
	})
}

// EngNameLTE applies the LTE predicate on the "engName" field.
func EngNameLTE(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEngName), v))
	})
}

// EngNameContains applies the Contains predicate on the "engName" field.
func EngNameContains(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEngName), v))
	})
}

// EngNameHasPrefix applies the HasPrefix predicate on the "engName" field.
func EngNameHasPrefix(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEngName), v))
	})
}

// EngNameHasSuffix applies the HasSuffix predicate on the "engName" field.
func EngNameHasSuffix(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEngName), v))
	})
}

// EngNameEqualFold applies the EqualFold predicate on the "engName" field.
func EngNameEqualFold(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEngName), v))
	})
}

// EngNameContainsFold applies the ContainsFold predicate on the "engName" field.
func EngNameContainsFold(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEngName), v))
	})
}

// RusNameEQ applies the EQ predicate on the "rusName" field.
func RusNameEQ(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRusName), v))
	})
}

// RusNameNEQ applies the NEQ predicate on the "rusName" field.
func RusNameNEQ(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRusName), v))
	})
}

// RusNameIn applies the In predicate on the "rusName" field.
func RusNameIn(vs ...string) predicate.Ranobe {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRusName), v...))
	})
}

// RusNameNotIn applies the NotIn predicate on the "rusName" field.
func RusNameNotIn(vs ...string) predicate.Ranobe {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRusName), v...))
	})
}

// RusNameGT applies the GT predicate on the "rusName" field.
func RusNameGT(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRusName), v))
	})
}

// RusNameGTE applies the GTE predicate on the "rusName" field.
func RusNameGTE(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRusName), v))
	})
}

// RusNameLT applies the LT predicate on the "rusName" field.
func RusNameLT(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRusName), v))
	})
}

// RusNameLTE applies the LTE predicate on the "rusName" field.
func RusNameLTE(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRusName), v))
	})
}

// RusNameContains applies the Contains predicate on the "rusName" field.
func RusNameContains(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRusName), v))
	})
}

// RusNameHasPrefix applies the HasPrefix predicate on the "rusName" field.
func RusNameHasPrefix(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRusName), v))
	})
}

// RusNameHasSuffix applies the HasSuffix predicate on the "rusName" field.
func RusNameHasSuffix(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRusName), v))
	})
}

// RusNameEqualFold applies the EqualFold predicate on the "rusName" field.
func RusNameEqualFold(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRusName), v))
	})
}

// RusNameContainsFold applies the ContainsFold predicate on the "rusName" field.
func RusNameContainsFold(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRusName), v))
	})
}

// SummaryEQ applies the EQ predicate on the "summary" field.
func SummaryEQ(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSummary), v))
	})
}

// SummaryNEQ applies the NEQ predicate on the "summary" field.
func SummaryNEQ(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSummary), v))
	})
}

// SummaryIn applies the In predicate on the "summary" field.
func SummaryIn(vs ...string) predicate.Ranobe {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSummary), v...))
	})
}

// SummaryNotIn applies the NotIn predicate on the "summary" field.
func SummaryNotIn(vs ...string) predicate.Ranobe {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSummary), v...))
	})
}

// SummaryGT applies the GT predicate on the "summary" field.
func SummaryGT(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSummary), v))
	})
}

// SummaryGTE applies the GTE predicate on the "summary" field.
func SummaryGTE(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSummary), v))
	})
}

// SummaryLT applies the LT predicate on the "summary" field.
func SummaryLT(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSummary), v))
	})
}

// SummaryLTE applies the LTE predicate on the "summary" field.
func SummaryLTE(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSummary), v))
	})
}

// SummaryContains applies the Contains predicate on the "summary" field.
func SummaryContains(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSummary), v))
	})
}

// SummaryHasPrefix applies the HasPrefix predicate on the "summary" field.
func SummaryHasPrefix(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSummary), v))
	})
}

// SummaryHasSuffix applies the HasSuffix predicate on the "summary" field.
func SummaryHasSuffix(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSummary), v))
	})
}

// SummaryEqualFold applies the EqualFold predicate on the "summary" field.
func SummaryEqualFold(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSummary), v))
	})
}

// SummaryContainsFold applies the ContainsFold predicate on the "summary" field.
func SummaryContainsFold(v string) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSummary), v))
	})
}

// ReleaseDateEQ applies the EQ predicate on the "releaseDate" field.
func ReleaseDateEQ(v int) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReleaseDate), v))
	})
}

// ReleaseDateNEQ applies the NEQ predicate on the "releaseDate" field.
func ReleaseDateNEQ(v int) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReleaseDate), v))
	})
}

// ReleaseDateIn applies the In predicate on the "releaseDate" field.
func ReleaseDateIn(vs ...int) predicate.Ranobe {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldReleaseDate), v...))
	})
}

// ReleaseDateNotIn applies the NotIn predicate on the "releaseDate" field.
func ReleaseDateNotIn(vs ...int) predicate.Ranobe {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldReleaseDate), v...))
	})
}

// ReleaseDateGT applies the GT predicate on the "releaseDate" field.
func ReleaseDateGT(v int) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldReleaseDate), v))
	})
}

// ReleaseDateGTE applies the GTE predicate on the "releaseDate" field.
func ReleaseDateGTE(v int) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldReleaseDate), v))
	})
}

// ReleaseDateLT applies the LT predicate on the "releaseDate" field.
func ReleaseDateLT(v int) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldReleaseDate), v))
	})
}

// ReleaseDateLTE applies the LTE predicate on the "releaseDate" field.
func ReleaseDateLTE(v int) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldReleaseDate), v))
	})
}

// HasCategories applies the HasEdge predicate on the "categories" edge.
func HasCategories() predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CategoriesTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, CategoriesTable, CategoriesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCategoriesWith applies the HasEdge predicate on the "categories" edge with a given conditions (other predicates).
func HasCategoriesWith(preds ...predicate.Category) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CategoriesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, CategoriesTable, CategoriesPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Ranobe) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Ranobe) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Ranobe) predicate.Ranobe {
	return predicate.Ranobe(func(s *sql.Selector) {
		p(s.Not())
	})
}
