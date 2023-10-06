// Code generated by ent, DO NOT EDIT.

package tasklog

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/suyuan32/simple-admin-job/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldLTE(FieldID, id))
}

// StartedAt applies equality check predicate on the "started_at" field. It's identical to StartedAtEQ.
func StartedAt(v time.Time) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldEQ(FieldStartedAt, v))
}

// FinishedAt applies equality check predicate on the "finished_at" field. It's identical to FinishedAtEQ.
func FinishedAt(v time.Time) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldEQ(FieldFinishedAt, v))
}

// Result applies equality check predicate on the "result" field. It's identical to ResultEQ.
func Result(v uint8) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldEQ(FieldResult, v))
}

// StartedAtEQ applies the EQ predicate on the "started_at" field.
func StartedAtEQ(v time.Time) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldEQ(FieldStartedAt, v))
}

// StartedAtNEQ applies the NEQ predicate on the "started_at" field.
func StartedAtNEQ(v time.Time) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldNEQ(FieldStartedAt, v))
}

// StartedAtIn applies the In predicate on the "started_at" field.
func StartedAtIn(vs ...time.Time) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldIn(FieldStartedAt, vs...))
}

// StartedAtNotIn applies the NotIn predicate on the "started_at" field.
func StartedAtNotIn(vs ...time.Time) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldNotIn(FieldStartedAt, vs...))
}

// StartedAtGT applies the GT predicate on the "started_at" field.
func StartedAtGT(v time.Time) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldGT(FieldStartedAt, v))
}

// StartedAtGTE applies the GTE predicate on the "started_at" field.
func StartedAtGTE(v time.Time) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldGTE(FieldStartedAt, v))
}

// StartedAtLT applies the LT predicate on the "started_at" field.
func StartedAtLT(v time.Time) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldLT(FieldStartedAt, v))
}

// StartedAtLTE applies the LTE predicate on the "started_at" field.
func StartedAtLTE(v time.Time) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldLTE(FieldStartedAt, v))
}

// FinishedAtEQ applies the EQ predicate on the "finished_at" field.
func FinishedAtEQ(v time.Time) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldEQ(FieldFinishedAt, v))
}

// FinishedAtNEQ applies the NEQ predicate on the "finished_at" field.
func FinishedAtNEQ(v time.Time) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldNEQ(FieldFinishedAt, v))
}

// FinishedAtIn applies the In predicate on the "finished_at" field.
func FinishedAtIn(vs ...time.Time) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldIn(FieldFinishedAt, vs...))
}

// FinishedAtNotIn applies the NotIn predicate on the "finished_at" field.
func FinishedAtNotIn(vs ...time.Time) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldNotIn(FieldFinishedAt, vs...))
}

// FinishedAtGT applies the GT predicate on the "finished_at" field.
func FinishedAtGT(v time.Time) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldGT(FieldFinishedAt, v))
}

// FinishedAtGTE applies the GTE predicate on the "finished_at" field.
func FinishedAtGTE(v time.Time) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldGTE(FieldFinishedAt, v))
}

// FinishedAtLT applies the LT predicate on the "finished_at" field.
func FinishedAtLT(v time.Time) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldLT(FieldFinishedAt, v))
}

// FinishedAtLTE applies the LTE predicate on the "finished_at" field.
func FinishedAtLTE(v time.Time) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldLTE(FieldFinishedAt, v))
}

// ResultEQ applies the EQ predicate on the "result" field.
func ResultEQ(v uint8) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldEQ(FieldResult, v))
}

// ResultNEQ applies the NEQ predicate on the "result" field.
func ResultNEQ(v uint8) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldNEQ(FieldResult, v))
}

// ResultIn applies the In predicate on the "result" field.
func ResultIn(vs ...uint8) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldIn(FieldResult, vs...))
}

// ResultNotIn applies the NotIn predicate on the "result" field.
func ResultNotIn(vs ...uint8) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldNotIn(FieldResult, vs...))
}

// ResultGT applies the GT predicate on the "result" field.
func ResultGT(v uint8) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldGT(FieldResult, v))
}

// ResultGTE applies the GTE predicate on the "result" field.
func ResultGTE(v uint8) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldGTE(FieldResult, v))
}

// ResultLT applies the LT predicate on the "result" field.
func ResultLT(v uint8) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldLT(FieldResult, v))
}

// ResultLTE applies the LTE predicate on the "result" field.
func ResultLTE(v uint8) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldLTE(FieldResult, v))
}

// HasTasks applies the HasEdge predicate on the "tasks" edge.
func HasTasks() predicate.TaskLog {
	return predicate.TaskLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TasksTable, TasksColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTasksWith applies the HasEdge predicate on the "tasks" edge with a given conditions (other predicates).
func HasTasksWith(preds ...predicate.Task) predicate.TaskLog {
	return predicate.TaskLog(func(s *sql.Selector) {
		step := newTasksStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TaskLog) predicate.TaskLog {
	return predicate.TaskLog(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TaskLog) predicate.TaskLog {
	return predicate.TaskLog(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TaskLog) predicate.TaskLog {
	return predicate.TaskLog(sql.NotPredicates(p))
}
