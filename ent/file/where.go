// Code generated by entc, DO NOT EDIT.

package file

import (
	"interface_project/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Path applies equality check predicate on the "path" field. It's identical to PathEQ.
func Path(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPath), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Size applies equality check predicate on the "size" field. It's identical to SizeEQ.
func Size(v int16) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSize), v))
	})
}

// Deleted applies equality check predicate on the "deleted" field. It's identical to DeletedEQ.
func Deleted(v bool) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleted), v))
	})
}

// CreatedDate applies equality check predicate on the "created_date" field. It's identical to CreatedDateEQ.
func CreatedDate(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedDate), v))
	})
}

// PathEQ applies the EQ predicate on the "path" field.
func PathEQ(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPath), v))
	})
}

// PathNEQ applies the NEQ predicate on the "path" field.
func PathNEQ(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPath), v))
	})
}

// PathIn applies the In predicate on the "path" field.
func PathIn(vs ...string) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPath), v...))
	})
}

// PathNotIn applies the NotIn predicate on the "path" field.
func PathNotIn(vs ...string) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPath), v...))
	})
}

// PathGT applies the GT predicate on the "path" field.
func PathGT(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPath), v))
	})
}

// PathGTE applies the GTE predicate on the "path" field.
func PathGTE(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPath), v))
	})
}

// PathLT applies the LT predicate on the "path" field.
func PathLT(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPath), v))
	})
}

// PathLTE applies the LTE predicate on the "path" field.
func PathLTE(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPath), v))
	})
}

// PathContains applies the Contains predicate on the "path" field.
func PathContains(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPath), v))
	})
}

// PathHasPrefix applies the HasPrefix predicate on the "path" field.
func PathHasPrefix(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPath), v))
	})
}

// PathHasSuffix applies the HasSuffix predicate on the "path" field.
func PathHasSuffix(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPath), v))
	})
}

// PathEqualFold applies the EqualFold predicate on the "path" field.
func PathEqualFold(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPath), v))
	})
}

// PathContainsFold applies the ContainsFold predicate on the "path" field.
func PathContainsFold(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPath), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// SizeEQ applies the EQ predicate on the "size" field.
func SizeEQ(v int16) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSize), v))
	})
}

// SizeNEQ applies the NEQ predicate on the "size" field.
func SizeNEQ(v int16) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSize), v))
	})
}

// SizeIn applies the In predicate on the "size" field.
func SizeIn(vs ...int16) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSize), v...))
	})
}

// SizeNotIn applies the NotIn predicate on the "size" field.
func SizeNotIn(vs ...int16) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSize), v...))
	})
}

// SizeGT applies the GT predicate on the "size" field.
func SizeGT(v int16) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSize), v))
	})
}

// SizeGTE applies the GTE predicate on the "size" field.
func SizeGTE(v int16) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSize), v))
	})
}

// SizeLT applies the LT predicate on the "size" field.
func SizeLT(v int16) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSize), v))
	})
}

// SizeLTE applies the LTE predicate on the "size" field.
func SizeLTE(v int16) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSize), v))
	})
}

// DeletedEQ applies the EQ predicate on the "deleted" field.
func DeletedEQ(v bool) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleted), v))
	})
}

// DeletedNEQ applies the NEQ predicate on the "deleted" field.
func DeletedNEQ(v bool) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeleted), v))
	})
}

// CreatedDateEQ applies the EQ predicate on the "created_date" field.
func CreatedDateEQ(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedDate), v))
	})
}

// CreatedDateNEQ applies the NEQ predicate on the "created_date" field.
func CreatedDateNEQ(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedDate), v))
	})
}

// CreatedDateIn applies the In predicate on the "created_date" field.
func CreatedDateIn(vs ...time.Time) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedDate), v...))
	})
}

// CreatedDateNotIn applies the NotIn predicate on the "created_date" field.
func CreatedDateNotIn(vs ...time.Time) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedDate), v...))
	})
}

// CreatedDateGT applies the GT predicate on the "created_date" field.
func CreatedDateGT(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedDate), v))
	})
}

// CreatedDateGTE applies the GTE predicate on the "created_date" field.
func CreatedDateGTE(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedDate), v))
	})
}

// CreatedDateLT applies the LT predicate on the "created_date" field.
func CreatedDateLT(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedDate), v))
	})
}

// CreatedDateLTE applies the LTE predicate on the "created_date" field.
func CreatedDateLTE(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedDate), v))
	})
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.File {
	return predicate.File(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.File) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.File) predicate.File {
	return predicate.File(func(s *sql.Selector) {
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
func Not(p predicate.File) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		p(s.Not())
	})
}
