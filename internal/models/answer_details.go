// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// AnswerDetail is an object representing the database table.
type AnswerDetail struct {
	ID        string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	AnswerID  string      `boil:"answer_id" json:"answer_id" toml:"answer_id" yaml:"answer_id"`
	Port      null.Int64  `boil:"port" json:"port,omitempty" toml:"port" yaml:"port,omitempty"`
	Priority  null.Int64  `boil:"priority" json:"priority,omitempty" toml:"priority" yaml:"priority,omitempty"`
	Protocol  null.String `boil:"protocol" json:"protocol,omitempty" toml:"protocol" yaml:"protocol,omitempty"`
	Weight    null.Int64  `boil:"weight" json:"weight,omitempty" toml:"weight" yaml:"weight,omitempty"`
	CreatedAt time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *answerDetailR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L answerDetailL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AnswerDetailColumns = struct {
	ID        string
	AnswerID  string
	Port      string
	Priority  string
	Protocol  string
	Weight    string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	AnswerID:  "answer_id",
	Port:      "port",
	Priority:  "priority",
	Protocol:  "protocol",
	Weight:    "weight",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var AnswerDetailTableColumns = struct {
	ID        string
	AnswerID  string
	Port      string
	Priority  string
	Protocol  string
	Weight    string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "answer_details.id",
	AnswerID:  "answer_details.answer_id",
	Port:      "answer_details.port",
	Priority:  "answer_details.priority",
	Protocol:  "answer_details.protocol",
	Weight:    "answer_details.weight",
	CreatedAt: "answer_details.created_at",
	UpdatedAt: "answer_details.updated_at",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_Int64 struct{ field string }

func (w whereHelpernull_Int64) EQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int64) NEQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int64) LT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int64) LTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int64) GT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int64) GTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_Int64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_Int64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_Int64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_String) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_String) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var AnswerDetailWhere = struct {
	ID        whereHelperstring
	AnswerID  whereHelperstring
	Port      whereHelpernull_Int64
	Priority  whereHelpernull_Int64
	Protocol  whereHelpernull_String
	Weight    whereHelpernull_Int64
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperstring{field: "\"answer_details\".\"id\""},
	AnswerID:  whereHelperstring{field: "\"answer_details\".\"answer_id\""},
	Port:      whereHelpernull_Int64{field: "\"answer_details\".\"port\""},
	Priority:  whereHelpernull_Int64{field: "\"answer_details\".\"priority\""},
	Protocol:  whereHelpernull_String{field: "\"answer_details\".\"protocol\""},
	Weight:    whereHelpernull_Int64{field: "\"answer_details\".\"weight\""},
	CreatedAt: whereHelpertime_Time{field: "\"answer_details\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"answer_details\".\"updated_at\""},
}

// AnswerDetailRels is where relationship names are stored.
var AnswerDetailRels = struct {
	Answer string
}{
	Answer: "Answer",
}

// answerDetailR is where relationships are stored.
type answerDetailR struct {
	Answer *Answer `boil:"Answer" json:"Answer" toml:"Answer" yaml:"Answer"`
}

// NewStruct creates a new relationship struct
func (*answerDetailR) NewStruct() *answerDetailR {
	return &answerDetailR{}
}

func (r *answerDetailR) GetAnswer() *Answer {
	if r == nil {
		return nil
	}
	return r.Answer
}

// answerDetailL is where Load methods for each relationship are stored.
type answerDetailL struct{}

var (
	answerDetailAllColumns            = []string{"id", "answer_id", "port", "priority", "protocol", "weight", "created_at", "updated_at"}
	answerDetailColumnsWithoutDefault = []string{"answer_id", "created_at", "updated_at"}
	answerDetailColumnsWithDefault    = []string{"id", "port", "priority", "protocol", "weight"}
	answerDetailPrimaryKeyColumns     = []string{"id"}
	answerDetailGeneratedColumns      = []string{}
)

type (
	// AnswerDetailSlice is an alias for a slice of pointers to AnswerDetail.
	// This should almost always be used instead of []AnswerDetail.
	AnswerDetailSlice []*AnswerDetail
	// AnswerDetailHook is the signature for custom AnswerDetail hook methods
	AnswerDetailHook func(context.Context, boil.ContextExecutor, *AnswerDetail) error

	answerDetailQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	answerDetailType                 = reflect.TypeOf(&AnswerDetail{})
	answerDetailMapping              = queries.MakeStructMapping(answerDetailType)
	answerDetailPrimaryKeyMapping, _ = queries.BindMapping(answerDetailType, answerDetailMapping, answerDetailPrimaryKeyColumns)
	answerDetailInsertCacheMut       sync.RWMutex
	answerDetailInsertCache          = make(map[string]insertCache)
	answerDetailUpdateCacheMut       sync.RWMutex
	answerDetailUpdateCache          = make(map[string]updateCache)
	answerDetailUpsertCacheMut       sync.RWMutex
	answerDetailUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var answerDetailAfterSelectHooks []AnswerDetailHook

var answerDetailBeforeInsertHooks []AnswerDetailHook
var answerDetailAfterInsertHooks []AnswerDetailHook

var answerDetailBeforeUpdateHooks []AnswerDetailHook
var answerDetailAfterUpdateHooks []AnswerDetailHook

var answerDetailBeforeDeleteHooks []AnswerDetailHook
var answerDetailAfterDeleteHooks []AnswerDetailHook

var answerDetailBeforeUpsertHooks []AnswerDetailHook
var answerDetailAfterUpsertHooks []AnswerDetailHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AnswerDetail) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerDetailAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AnswerDetail) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerDetailBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AnswerDetail) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerDetailAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AnswerDetail) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerDetailBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AnswerDetail) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerDetailAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AnswerDetail) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerDetailBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AnswerDetail) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerDetailAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AnswerDetail) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerDetailBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AnswerDetail) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerDetailAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAnswerDetailHook registers your hook function for all future operations.
func AddAnswerDetailHook(hookPoint boil.HookPoint, answerDetailHook AnswerDetailHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		answerDetailAfterSelectHooks = append(answerDetailAfterSelectHooks, answerDetailHook)
	case boil.BeforeInsertHook:
		answerDetailBeforeInsertHooks = append(answerDetailBeforeInsertHooks, answerDetailHook)
	case boil.AfterInsertHook:
		answerDetailAfterInsertHooks = append(answerDetailAfterInsertHooks, answerDetailHook)
	case boil.BeforeUpdateHook:
		answerDetailBeforeUpdateHooks = append(answerDetailBeforeUpdateHooks, answerDetailHook)
	case boil.AfterUpdateHook:
		answerDetailAfterUpdateHooks = append(answerDetailAfterUpdateHooks, answerDetailHook)
	case boil.BeforeDeleteHook:
		answerDetailBeforeDeleteHooks = append(answerDetailBeforeDeleteHooks, answerDetailHook)
	case boil.AfterDeleteHook:
		answerDetailAfterDeleteHooks = append(answerDetailAfterDeleteHooks, answerDetailHook)
	case boil.BeforeUpsertHook:
		answerDetailBeforeUpsertHooks = append(answerDetailBeforeUpsertHooks, answerDetailHook)
	case boil.AfterUpsertHook:
		answerDetailAfterUpsertHooks = append(answerDetailAfterUpsertHooks, answerDetailHook)
	}
}

// One returns a single answerDetail record from the query.
func (q answerDetailQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AnswerDetail, error) {
	o := &AnswerDetail{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for answer_details")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all AnswerDetail records from the query.
func (q answerDetailQuery) All(ctx context.Context, exec boil.ContextExecutor) (AnswerDetailSlice, error) {
	var o []*AnswerDetail

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AnswerDetail slice")
	}

	if len(answerDetailAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all AnswerDetail records in the query.
func (q answerDetailQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count answer_details rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q answerDetailQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if answer_details exists")
	}

	return count > 0, nil
}

// Answer pointed to by the foreign key.
func (o *AnswerDetail) Answer(mods ...qm.QueryMod) answerQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.AnswerID),
	}

	queryMods = append(queryMods, mods...)

	return Answers(queryMods...)
}

// LoadAnswer allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (answerDetailL) LoadAnswer(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAnswerDetail interface{}, mods queries.Applicator) error {
	var slice []*AnswerDetail
	var object *AnswerDetail

	if singular {
		var ok bool
		object, ok = maybeAnswerDetail.(*AnswerDetail)
		if !ok {
			object = new(AnswerDetail)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeAnswerDetail)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeAnswerDetail))
			}
		}
	} else {
		s, ok := maybeAnswerDetail.(*[]*AnswerDetail)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeAnswerDetail)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeAnswerDetail))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &answerDetailR{}
		}
		args = append(args, object.AnswerID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &answerDetailR{}
			}

			for _, a := range args {
				if a == obj.AnswerID {
					continue Outer
				}
			}

			args = append(args, obj.AnswerID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`answers`),
		qm.WhereIn(`answers.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Answer")
	}

	var resultSlice []*Answer
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Answer")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for answers")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for answers")
	}

	if len(answerDetailAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Answer = foreign
		if foreign.R == nil {
			foreign.R = &answerR{}
		}
		foreign.R.AnswerDetail = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.AnswerID == foreign.ID {
				local.R.Answer = foreign
				if foreign.R == nil {
					foreign.R = &answerR{}
				}
				foreign.R.AnswerDetail = local
				break
			}
		}
	}

	return nil
}

// SetAnswer of the answerDetail to the related item.
// Sets o.R.Answer to related.
// Adds o to related.R.AnswerDetail.
func (o *AnswerDetail) SetAnswer(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Answer) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"answer_details\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"answer_id"}),
		strmangle.WhereClause("\"", "\"", 2, answerDetailPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.AnswerID = related.ID
	if o.R == nil {
		o.R = &answerDetailR{
			Answer: related,
		}
	} else {
		o.R.Answer = related
	}

	if related.R == nil {
		related.R = &answerR{
			AnswerDetail: o,
		}
	} else {
		related.R.AnswerDetail = o
	}

	return nil
}

// AnswerDetails retrieves all the records using an executor.
func AnswerDetails(mods ...qm.QueryMod) answerDetailQuery {
	mods = append(mods, qm.From("\"answer_details\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"answer_details\".*"})
	}

	return answerDetailQuery{q}
}

// FindAnswerDetail retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAnswerDetail(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*AnswerDetail, error) {
	answerDetailObj := &AnswerDetail{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"answer_details\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, answerDetailObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from answer_details")
	}

	if err = answerDetailObj.doAfterSelectHooks(ctx, exec); err != nil {
		return answerDetailObj, err
	}

	return answerDetailObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AnswerDetail) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no answer_details provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(answerDetailColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	answerDetailInsertCacheMut.RLock()
	cache, cached := answerDetailInsertCache[key]
	answerDetailInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			answerDetailAllColumns,
			answerDetailColumnsWithDefault,
			answerDetailColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(answerDetailType, answerDetailMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(answerDetailType, answerDetailMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"answer_details\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"answer_details\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into answer_details")
	}

	if !cached {
		answerDetailInsertCacheMut.Lock()
		answerDetailInsertCache[key] = cache
		answerDetailInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the AnswerDetail.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AnswerDetail) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	answerDetailUpdateCacheMut.RLock()
	cache, cached := answerDetailUpdateCache[key]
	answerDetailUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			answerDetailAllColumns,
			answerDetailPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update answer_details, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"answer_details\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, answerDetailPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(answerDetailType, answerDetailMapping, append(wl, answerDetailPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update answer_details row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for answer_details")
	}

	if !cached {
		answerDetailUpdateCacheMut.Lock()
		answerDetailUpdateCache[key] = cache
		answerDetailUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q answerDetailQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for answer_details")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for answer_details")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AnswerDetailSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), answerDetailPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"answer_details\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, answerDetailPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in answerDetail slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all answerDetail")
	}
	return rowsAff, nil
}

// Delete deletes a single AnswerDetail record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AnswerDetail) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AnswerDetail provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), answerDetailPrimaryKeyMapping)
	sql := "DELETE FROM \"answer_details\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from answer_details")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for answer_details")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q answerDetailQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no answerDetailQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from answer_details")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for answer_details")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AnswerDetailSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(answerDetailBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), answerDetailPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"answer_details\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, answerDetailPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from answerDetail slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for answer_details")
	}

	if len(answerDetailAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *AnswerDetail) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAnswerDetail(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AnswerDetailSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AnswerDetailSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), answerDetailPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"answer_details\".* FROM \"answer_details\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, answerDetailPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AnswerDetailSlice")
	}

	*o = slice

	return nil
}

// AnswerDetailExists checks if the AnswerDetail row exists.
func AnswerDetailExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"answer_details\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if answer_details exists")
	}

	return exists, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AnswerDetail) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no answer_details provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(answerDetailColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	answerDetailUpsertCacheMut.RLock()
	cache, cached := answerDetailUpsertCache[key]
	answerDetailUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			answerDetailAllColumns,
			answerDetailColumnsWithDefault,
			answerDetailColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			answerDetailAllColumns,
			answerDetailPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert answer_details, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(answerDetailPrimaryKeyColumns))
			copy(conflict, answerDetailPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryCockroachDB(dialect, "\"answer_details\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(answerDetailType, answerDetailMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(answerDetailType, answerDetailMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		_, _ = fmt.Fprintln(boil.DebugWriter, cache.query)
		_, _ = fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // CockcorachDB doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert answer_details")
	}

	if !cached {
		answerDetailUpsertCacheMut.Lock()
		answerDetailUpsertCache[key] = cache
		answerDetailUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}
