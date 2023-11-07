// Code generated by ent, DO NOT EDIT.

package oidc

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/disism/saikan/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Oidc {
	return predicate.Oidc(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Oidc {
	return predicate.Oidc(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Oidc {
	return predicate.Oidc(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Oidc {
	return predicate.Oidc(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Oidc {
	return predicate.Oidc(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Oidc {
	return predicate.Oidc(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Oidc {
	return predicate.Oidc(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Oidc {
	return predicate.Oidc(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Oidc {
	return predicate.Oidc(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Oidc {
	return predicate.Oidc(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Oidc {
	return predicate.Oidc(sql.FieldEQ(FieldUpdateTime, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Oidc {
	return predicate.Oidc(sql.FieldEQ(FieldName, v))
}

// ConfigurationEndpoint applies equality check predicate on the "configuration_endpoint" field. It's identical to ConfigurationEndpointEQ.
func ConfigurationEndpoint(v string) predicate.Oidc {
	return predicate.Oidc(sql.FieldEQ(FieldConfigurationEndpoint, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Oidc {
	return predicate.Oidc(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Oidc {
	return predicate.Oidc(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Oidc {
	return predicate.Oidc(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Oidc {
	return predicate.Oidc(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Oidc {
	return predicate.Oidc(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Oidc {
	return predicate.Oidc(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Oidc {
	return predicate.Oidc(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Oidc {
	return predicate.Oidc(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Oidc {
	return predicate.Oidc(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Oidc {
	return predicate.Oidc(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Oidc {
	return predicate.Oidc(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Oidc {
	return predicate.Oidc(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Oidc {
	return predicate.Oidc(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Oidc {
	return predicate.Oidc(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Oidc {
	return predicate.Oidc(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Oidc {
	return predicate.Oidc(sql.FieldLTE(FieldUpdateTime, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Oidc {
	return predicate.Oidc(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Oidc {
	return predicate.Oidc(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Oidc {
	return predicate.Oidc(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Oidc {
	return predicate.Oidc(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Oidc {
	return predicate.Oidc(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Oidc {
	return predicate.Oidc(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Oidc {
	return predicate.Oidc(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Oidc {
	return predicate.Oidc(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Oidc {
	return predicate.Oidc(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Oidc {
	return predicate.Oidc(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Oidc {
	return predicate.Oidc(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Oidc {
	return predicate.Oidc(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Oidc {
	return predicate.Oidc(sql.FieldContainsFold(FieldName, v))
}

// ConfigurationEndpointEQ applies the EQ predicate on the "configuration_endpoint" field.
func ConfigurationEndpointEQ(v string) predicate.Oidc {
	return predicate.Oidc(sql.FieldEQ(FieldConfigurationEndpoint, v))
}

// ConfigurationEndpointNEQ applies the NEQ predicate on the "configuration_endpoint" field.
func ConfigurationEndpointNEQ(v string) predicate.Oidc {
	return predicate.Oidc(sql.FieldNEQ(FieldConfigurationEndpoint, v))
}

// ConfigurationEndpointIn applies the In predicate on the "configuration_endpoint" field.
func ConfigurationEndpointIn(vs ...string) predicate.Oidc {
	return predicate.Oidc(sql.FieldIn(FieldConfigurationEndpoint, vs...))
}

// ConfigurationEndpointNotIn applies the NotIn predicate on the "configuration_endpoint" field.
func ConfigurationEndpointNotIn(vs ...string) predicate.Oidc {
	return predicate.Oidc(sql.FieldNotIn(FieldConfigurationEndpoint, vs...))
}

// ConfigurationEndpointGT applies the GT predicate on the "configuration_endpoint" field.
func ConfigurationEndpointGT(v string) predicate.Oidc {
	return predicate.Oidc(sql.FieldGT(FieldConfigurationEndpoint, v))
}

// ConfigurationEndpointGTE applies the GTE predicate on the "configuration_endpoint" field.
func ConfigurationEndpointGTE(v string) predicate.Oidc {
	return predicate.Oidc(sql.FieldGTE(FieldConfigurationEndpoint, v))
}

// ConfigurationEndpointLT applies the LT predicate on the "configuration_endpoint" field.
func ConfigurationEndpointLT(v string) predicate.Oidc {
	return predicate.Oidc(sql.FieldLT(FieldConfigurationEndpoint, v))
}

// ConfigurationEndpointLTE applies the LTE predicate on the "configuration_endpoint" field.
func ConfigurationEndpointLTE(v string) predicate.Oidc {
	return predicate.Oidc(sql.FieldLTE(FieldConfigurationEndpoint, v))
}

// ConfigurationEndpointContains applies the Contains predicate on the "configuration_endpoint" field.
func ConfigurationEndpointContains(v string) predicate.Oidc {
	return predicate.Oidc(sql.FieldContains(FieldConfigurationEndpoint, v))
}

// ConfigurationEndpointHasPrefix applies the HasPrefix predicate on the "configuration_endpoint" field.
func ConfigurationEndpointHasPrefix(v string) predicate.Oidc {
	return predicate.Oidc(sql.FieldHasPrefix(FieldConfigurationEndpoint, v))
}

// ConfigurationEndpointHasSuffix applies the HasSuffix predicate on the "configuration_endpoint" field.
func ConfigurationEndpointHasSuffix(v string) predicate.Oidc {
	return predicate.Oidc(sql.FieldHasSuffix(FieldConfigurationEndpoint, v))
}

// ConfigurationEndpointEqualFold applies the EqualFold predicate on the "configuration_endpoint" field.
func ConfigurationEndpointEqualFold(v string) predicate.Oidc {
	return predicate.Oidc(sql.FieldEqualFold(FieldConfigurationEndpoint, v))
}

// ConfigurationEndpointContainsFold applies the ContainsFold predicate on the "configuration_endpoint" field.
func ConfigurationEndpointContainsFold(v string) predicate.Oidc {
	return predicate.Oidc(sql.FieldContainsFold(FieldConfigurationEndpoint, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Oidc) predicate.Oidc {
	return predicate.Oidc(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Oidc) predicate.Oidc {
	return predicate.Oidc(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Oidc) predicate.Oidc {
	return predicate.Oidc(sql.NotPredicates(p))
}
