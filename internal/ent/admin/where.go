// Code generated by ent, DO NOT EDIT.

package admin

import (
	"entgo.io/ent/dialect/sql"
	"github.com/chetan0402/veripass/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Admin {
	return predicate.Admin(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Admin {
	return predicate.Admin(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Admin {
	return predicate.Admin(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Admin {
	return predicate.Admin(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Admin {
	return predicate.Admin(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Admin {
	return predicate.Admin(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Admin {
	return predicate.Admin(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Admin {
	return predicate.Admin(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Admin {
	return predicate.Admin(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Admin {
	return predicate.Admin(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Admin {
	return predicate.Admin(sql.FieldContainsFold(FieldID, id))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Admin {
	return predicate.Admin(sql.FieldEQ(FieldEmail, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Admin {
	return predicate.Admin(sql.FieldEQ(FieldName, v))
}

// Hostel applies equality check predicate on the "hostel" field. It's identical to HostelEQ.
func Hostel(v string) predicate.Admin {
	return predicate.Admin(sql.FieldEQ(FieldHostel, v))
}

// CanAddPass applies equality check predicate on the "can_add_pass" field. It's identical to CanAddPassEQ.
func CanAddPass(v bool) predicate.Admin {
	return predicate.Admin(sql.FieldEQ(FieldCanAddPass, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Admin {
	return predicate.Admin(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Admin {
	return predicate.Admin(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Admin {
	return predicate.Admin(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Admin {
	return predicate.Admin(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Admin {
	return predicate.Admin(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Admin {
	return predicate.Admin(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Admin {
	return predicate.Admin(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Admin {
	return predicate.Admin(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Admin {
	return predicate.Admin(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Admin {
	return predicate.Admin(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Admin {
	return predicate.Admin(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Admin {
	return predicate.Admin(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Admin {
	return predicate.Admin(sql.FieldContainsFold(FieldEmail, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Admin {
	return predicate.Admin(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Admin {
	return predicate.Admin(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Admin {
	return predicate.Admin(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Admin {
	return predicate.Admin(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Admin {
	return predicate.Admin(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Admin {
	return predicate.Admin(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Admin {
	return predicate.Admin(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Admin {
	return predicate.Admin(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Admin {
	return predicate.Admin(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Admin {
	return predicate.Admin(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Admin {
	return predicate.Admin(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Admin {
	return predicate.Admin(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Admin {
	return predicate.Admin(sql.FieldContainsFold(FieldName, v))
}

// HostelEQ applies the EQ predicate on the "hostel" field.
func HostelEQ(v string) predicate.Admin {
	return predicate.Admin(sql.FieldEQ(FieldHostel, v))
}

// HostelNEQ applies the NEQ predicate on the "hostel" field.
func HostelNEQ(v string) predicate.Admin {
	return predicate.Admin(sql.FieldNEQ(FieldHostel, v))
}

// HostelIn applies the In predicate on the "hostel" field.
func HostelIn(vs ...string) predicate.Admin {
	return predicate.Admin(sql.FieldIn(FieldHostel, vs...))
}

// HostelNotIn applies the NotIn predicate on the "hostel" field.
func HostelNotIn(vs ...string) predicate.Admin {
	return predicate.Admin(sql.FieldNotIn(FieldHostel, vs...))
}

// HostelGT applies the GT predicate on the "hostel" field.
func HostelGT(v string) predicate.Admin {
	return predicate.Admin(sql.FieldGT(FieldHostel, v))
}

// HostelGTE applies the GTE predicate on the "hostel" field.
func HostelGTE(v string) predicate.Admin {
	return predicate.Admin(sql.FieldGTE(FieldHostel, v))
}

// HostelLT applies the LT predicate on the "hostel" field.
func HostelLT(v string) predicate.Admin {
	return predicate.Admin(sql.FieldLT(FieldHostel, v))
}

// HostelLTE applies the LTE predicate on the "hostel" field.
func HostelLTE(v string) predicate.Admin {
	return predicate.Admin(sql.FieldLTE(FieldHostel, v))
}

// HostelContains applies the Contains predicate on the "hostel" field.
func HostelContains(v string) predicate.Admin {
	return predicate.Admin(sql.FieldContains(FieldHostel, v))
}

// HostelHasPrefix applies the HasPrefix predicate on the "hostel" field.
func HostelHasPrefix(v string) predicate.Admin {
	return predicate.Admin(sql.FieldHasPrefix(FieldHostel, v))
}

// HostelHasSuffix applies the HasSuffix predicate on the "hostel" field.
func HostelHasSuffix(v string) predicate.Admin {
	return predicate.Admin(sql.FieldHasSuffix(FieldHostel, v))
}

// HostelEqualFold applies the EqualFold predicate on the "hostel" field.
func HostelEqualFold(v string) predicate.Admin {
	return predicate.Admin(sql.FieldEqualFold(FieldHostel, v))
}

// HostelContainsFold applies the ContainsFold predicate on the "hostel" field.
func HostelContainsFold(v string) predicate.Admin {
	return predicate.Admin(sql.FieldContainsFold(FieldHostel, v))
}

// CanAddPassEQ applies the EQ predicate on the "can_add_pass" field.
func CanAddPassEQ(v bool) predicate.Admin {
	return predicate.Admin(sql.FieldEQ(FieldCanAddPass, v))
}

// CanAddPassNEQ applies the NEQ predicate on the "can_add_pass" field.
func CanAddPassNEQ(v bool) predicate.Admin {
	return predicate.Admin(sql.FieldNEQ(FieldCanAddPass, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Admin) predicate.Admin {
	return predicate.Admin(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Admin) predicate.Admin {
	return predicate.Admin(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Admin) predicate.Admin {
	return predicate.Admin(sql.NotPredicates(p))
}
