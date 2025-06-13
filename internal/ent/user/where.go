// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"github.com/chetan0402/veripass/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// Room applies equality check predicate on the "room" field. It's identical to RoomEQ.
func Room(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldRoom, v))
}

// Hostel applies equality check predicate on the "hostel" field. It's identical to HostelEQ.
func Hostel(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldHostel, v))
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPhone, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldName, v))
}

// RoomEQ applies the EQ predicate on the "room" field.
func RoomEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldRoom, v))
}

// RoomNEQ applies the NEQ predicate on the "room" field.
func RoomNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldRoom, v))
}

// RoomIn applies the In predicate on the "room" field.
func RoomIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldRoom, vs...))
}

// RoomNotIn applies the NotIn predicate on the "room" field.
func RoomNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldRoom, vs...))
}

// RoomGT applies the GT predicate on the "room" field.
func RoomGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldRoom, v))
}

// RoomGTE applies the GTE predicate on the "room" field.
func RoomGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldRoom, v))
}

// RoomLT applies the LT predicate on the "room" field.
func RoomLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldRoom, v))
}

// RoomLTE applies the LTE predicate on the "room" field.
func RoomLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldRoom, v))
}

// RoomContains applies the Contains predicate on the "room" field.
func RoomContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldRoom, v))
}

// RoomHasPrefix applies the HasPrefix predicate on the "room" field.
func RoomHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldRoom, v))
}

// RoomHasSuffix applies the HasSuffix predicate on the "room" field.
func RoomHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldRoom, v))
}

// RoomEqualFold applies the EqualFold predicate on the "room" field.
func RoomEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldRoom, v))
}

// RoomContainsFold applies the ContainsFold predicate on the "room" field.
func RoomContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldRoom, v))
}

// HostelEQ applies the EQ predicate on the "hostel" field.
func HostelEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldHostel, v))
}

// HostelNEQ applies the NEQ predicate on the "hostel" field.
func HostelNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldHostel, v))
}

// HostelIn applies the In predicate on the "hostel" field.
func HostelIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldHostel, vs...))
}

// HostelNotIn applies the NotIn predicate on the "hostel" field.
func HostelNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldHostel, vs...))
}

// HostelGT applies the GT predicate on the "hostel" field.
func HostelGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldHostel, v))
}

// HostelGTE applies the GTE predicate on the "hostel" field.
func HostelGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldHostel, v))
}

// HostelLT applies the LT predicate on the "hostel" field.
func HostelLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldHostel, v))
}

// HostelLTE applies the LTE predicate on the "hostel" field.
func HostelLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldHostel, v))
}

// HostelContains applies the Contains predicate on the "hostel" field.
func HostelContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldHostel, v))
}

// HostelHasPrefix applies the HasPrefix predicate on the "hostel" field.
func HostelHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldHostel, v))
}

// HostelHasSuffix applies the HasSuffix predicate on the "hostel" field.
func HostelHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldHostel, v))
}

// HostelEqualFold applies the EqualFold predicate on the "hostel" field.
func HostelEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldHostel, v))
}

// HostelContainsFold applies the ContainsFold predicate on the "hostel" field.
func HostelContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldHostel, v))
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPhone, v))
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPhone, v))
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPhone, vs...))
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPhone, vs...))
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPhone, v))
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPhone, v))
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPhone, v))
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPhone, v))
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPhone, v))
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPhone, v))
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPhone, v))
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPhone, v))
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPhone, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
