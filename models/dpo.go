// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// Dpo is an object representing the database table.
type Dpo struct {
	DemandPartnerID string    `boil:"demand_partner_id" json:"demand_partner_id" toml:"demand_partner_id" yaml:"demand_partner_id"`
	IsInclude       bool      `boil:"is_include" json:"is_include" toml:"is_include" yaml:"is_include"`
	CreatedAt       time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt       null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *dpoR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dpoL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DpoColumns = struct {
	DemandPartnerID string
	IsInclude       string
	CreatedAt       string
	UpdatedAt       string
}{
	DemandPartnerID: "demand_partner_id",
	IsInclude:       "is_include",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

var DpoTableColumns = struct {
	DemandPartnerID string
	IsInclude       string
	CreatedAt       string
	UpdatedAt       string
}{
	DemandPartnerID: "dpo.demand_partner_id",
	IsInclude:       "dpo.is_include",
	CreatedAt:       "dpo.created_at",
	UpdatedAt:       "dpo.updated_at",
}

// Generated where

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var DpoWhere = struct {
	DemandPartnerID whereHelperstring
	IsInclude       whereHelperbool
	CreatedAt       whereHelpertime_Time
	UpdatedAt       whereHelpernull_Time
}{
	DemandPartnerID: whereHelperstring{field: "\"dpo\".\"demand_partner_id\""},
	IsInclude:       whereHelperbool{field: "\"dpo\".\"is_include\""},
	CreatedAt:       whereHelpertime_Time{field: "\"dpo\".\"created_at\""},
	UpdatedAt:       whereHelpernull_Time{field: "\"dpo\".\"updated_at\""},
}

// DpoRels is where relationship names are stored.
var DpoRels = struct {
	DemandPartnerDpoRules string
}{
	DemandPartnerDpoRules: "DemandPartnerDpoRules",
}

// dpoR is where relationships are stored.
type dpoR struct {
	DemandPartnerDpoRules DpoRuleSlice `boil:"DemandPartnerDpoRules" json:"DemandPartnerDpoRules" toml:"DemandPartnerDpoRules" yaml:"DemandPartnerDpoRules"`
}

// NewStruct creates a new relationship struct
func (*dpoR) NewStruct() *dpoR {
	return &dpoR{}
}

func (r *dpoR) GetDemandPartnerDpoRules() DpoRuleSlice {
	if r == nil {
		return nil
	}
	return r.DemandPartnerDpoRules
}

// dpoL is where Load methods for each relationship are stored.
type dpoL struct{}

var (
	dpoAllColumns            = []string{"demand_partner_id", "is_include", "created_at", "updated_at"}
	dpoColumnsWithoutDefault = []string{"demand_partner_id", "created_at"}
	dpoColumnsWithDefault    = []string{"is_include", "updated_at"}
	dpoPrimaryKeyColumns     = []string{"demand_partner_id"}
	dpoGeneratedColumns      = []string{}
)

type (
	// DpoSlice is an alias for a slice of pointers to Dpo.
	// This should almost always be used instead of []Dpo.
	DpoSlice []*Dpo
	// DpoHook is the signature for custom Dpo hook methods
	DpoHook func(context.Context, boil.ContextExecutor, *Dpo) error

	dpoQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dpoType                 = reflect.TypeOf(&Dpo{})
	dpoMapping              = queries.MakeStructMapping(dpoType)
	dpoPrimaryKeyMapping, _ = queries.BindMapping(dpoType, dpoMapping, dpoPrimaryKeyColumns)
	dpoInsertCacheMut       sync.RWMutex
	dpoInsertCache          = make(map[string]insertCache)
	dpoUpdateCacheMut       sync.RWMutex
	dpoUpdateCache          = make(map[string]updateCache)
	dpoUpsertCacheMut       sync.RWMutex
	dpoUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var dpoAfterSelectMu sync.Mutex
var dpoAfterSelectHooks []DpoHook

var dpoBeforeInsertMu sync.Mutex
var dpoBeforeInsertHooks []DpoHook
var dpoAfterInsertMu sync.Mutex
var dpoAfterInsertHooks []DpoHook

var dpoBeforeUpdateMu sync.Mutex
var dpoBeforeUpdateHooks []DpoHook
var dpoAfterUpdateMu sync.Mutex
var dpoAfterUpdateHooks []DpoHook

var dpoBeforeDeleteMu sync.Mutex
var dpoBeforeDeleteHooks []DpoHook
var dpoAfterDeleteMu sync.Mutex
var dpoAfterDeleteHooks []DpoHook

var dpoBeforeUpsertMu sync.Mutex
var dpoBeforeUpsertHooks []DpoHook
var dpoAfterUpsertMu sync.Mutex
var dpoAfterUpsertHooks []DpoHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Dpo) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dpoAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Dpo) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dpoBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Dpo) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dpoAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Dpo) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dpoBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Dpo) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dpoAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Dpo) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dpoBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Dpo) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dpoAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Dpo) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dpoBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Dpo) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dpoAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDpoHook registers your hook function for all future operations.
func AddDpoHook(hookPoint boil.HookPoint, dpoHook DpoHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		dpoAfterSelectMu.Lock()
		dpoAfterSelectHooks = append(dpoAfterSelectHooks, dpoHook)
		dpoAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		dpoBeforeInsertMu.Lock()
		dpoBeforeInsertHooks = append(dpoBeforeInsertHooks, dpoHook)
		dpoBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		dpoAfterInsertMu.Lock()
		dpoAfterInsertHooks = append(dpoAfterInsertHooks, dpoHook)
		dpoAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		dpoBeforeUpdateMu.Lock()
		dpoBeforeUpdateHooks = append(dpoBeforeUpdateHooks, dpoHook)
		dpoBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		dpoAfterUpdateMu.Lock()
		dpoAfterUpdateHooks = append(dpoAfterUpdateHooks, dpoHook)
		dpoAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		dpoBeforeDeleteMu.Lock()
		dpoBeforeDeleteHooks = append(dpoBeforeDeleteHooks, dpoHook)
		dpoBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		dpoAfterDeleteMu.Lock()
		dpoAfterDeleteHooks = append(dpoAfterDeleteHooks, dpoHook)
		dpoAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		dpoBeforeUpsertMu.Lock()
		dpoBeforeUpsertHooks = append(dpoBeforeUpsertHooks, dpoHook)
		dpoBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		dpoAfterUpsertMu.Lock()
		dpoAfterUpsertHooks = append(dpoAfterUpsertHooks, dpoHook)
		dpoAfterUpsertMu.Unlock()
	}
}

// One returns a single dpo record from the query.
func (q dpoQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Dpo, error) {
	o := &Dpo{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for dpo")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Dpo records from the query.
func (q dpoQuery) All(ctx context.Context, exec boil.ContextExecutor) (DpoSlice, error) {
	var o []*Dpo

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Dpo slice")
	}

	if len(dpoAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Dpo records in the query.
func (q dpoQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count dpo rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q dpoQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if dpo exists")
	}

	return count > 0, nil
}

// DemandPartnerDpoRules retrieves all the dpo_rule's DpoRules with an executor via demand_partner_id column.
func (o *Dpo) DemandPartnerDpoRules(mods ...qm.QueryMod) dpoRuleQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"dpo_rule\".\"demand_partner_id\"=?", o.DemandPartnerID),
	)

	return DpoRules(queryMods...)
}

// LoadDemandPartnerDpoRules allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (dpoL) LoadDemandPartnerDpoRules(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDpo interface{}, mods queries.Applicator) error {
	var slice []*Dpo
	var object *Dpo

	if singular {
		var ok bool
		object, ok = maybeDpo.(*Dpo)
		if !ok {
			object = new(Dpo)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeDpo)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeDpo))
			}
		}
	} else {
		s, ok := maybeDpo.(*[]*Dpo)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeDpo)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeDpo))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &dpoR{}
		}
		args[object.DemandPartnerID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &dpoR{}
			}
			args[obj.DemandPartnerID] = struct{}{}
		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`dpo_rule`),
		qm.WhereIn(`dpo_rule.demand_partner_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load dpo_rule")
	}

	var resultSlice []*DpoRule
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice dpo_rule")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on dpo_rule")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for dpo_rule")
	}

	if len(dpoRuleAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.DemandPartnerDpoRules = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &dpoRuleR{}
			}
			foreign.R.DemandPartner = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.DemandPartnerID == foreign.DemandPartnerID {
				local.R.DemandPartnerDpoRules = append(local.R.DemandPartnerDpoRules, foreign)
				if foreign.R == nil {
					foreign.R = &dpoRuleR{}
				}
				foreign.R.DemandPartner = local
				break
			}
		}
	}

	return nil
}

// AddDemandPartnerDpoRules adds the given related objects to the existing relationships
// of the dpo, optionally inserting them as new records.
// Appends related to o.R.DemandPartnerDpoRules.
// Sets related.R.DemandPartner appropriately.
func (o *Dpo) AddDemandPartnerDpoRules(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*DpoRule) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.DemandPartnerID = o.DemandPartnerID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"dpo_rule\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"demand_partner_id"}),
				strmangle.WhereClause("\"", "\"", 2, dpoRulePrimaryKeyColumns),
			)
			values := []interface{}{o.DemandPartnerID, rel.RuleID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.DemandPartnerID = o.DemandPartnerID
		}
	}

	if o.R == nil {
		o.R = &dpoR{
			DemandPartnerDpoRules: related,
		}
	} else {
		o.R.DemandPartnerDpoRules = append(o.R.DemandPartnerDpoRules, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &dpoRuleR{
				DemandPartner: o,
			}
		} else {
			rel.R.DemandPartner = o
		}
	}
	return nil
}

// Dpos retrieves all the records using an executor.
func Dpos(mods ...qm.QueryMod) dpoQuery {
	mods = append(mods, qm.From("\"dpo\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"dpo\".*"})
	}

	return dpoQuery{q}
}

// FindDpo retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDpo(ctx context.Context, exec boil.ContextExecutor, demandPartnerID string, selectCols ...string) (*Dpo, error) {
	dpoObj := &Dpo{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"dpo\" where \"demand_partner_id\"=$1", sel,
	)

	q := queries.Raw(query, demandPartnerID)

	err := q.Bind(ctx, exec, dpoObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from dpo")
	}

	if err = dpoObj.doAfterSelectHooks(ctx, exec); err != nil {
		return dpoObj, err
	}

	return dpoObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Dpo) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no dpo provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dpoColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	dpoInsertCacheMut.RLock()
	cache, cached := dpoInsertCache[key]
	dpoInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			dpoAllColumns,
			dpoColumnsWithDefault,
			dpoColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(dpoType, dpoMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dpoType, dpoMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"dpo\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"dpo\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into dpo")
	}

	if !cached {
		dpoInsertCacheMut.Lock()
		dpoInsertCache[key] = cache
		dpoInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Dpo.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Dpo) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	dpoUpdateCacheMut.RLock()
	cache, cached := dpoUpdateCache[key]
	dpoUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			dpoAllColumns,
			dpoPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update dpo, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"dpo\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, dpoPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dpoType, dpoMapping, append(wl, dpoPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update dpo row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for dpo")
	}

	if !cached {
		dpoUpdateCacheMut.Lock()
		dpoUpdateCache[key] = cache
		dpoUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q dpoQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for dpo")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for dpo")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DpoSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dpoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"dpo\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, dpoPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in dpo slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all dpo")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Dpo) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no dpo provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dpoColumnsWithDefault, o)

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

	dpoUpsertCacheMut.RLock()
	cache, cached := dpoUpsertCache[key]
	dpoUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			dpoAllColumns,
			dpoColumnsWithDefault,
			dpoColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			dpoAllColumns,
			dpoPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert dpo, could not build update column list")
		}

		ret := strmangle.SetComplement(dpoAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(dpoPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert dpo, could not build conflict column list")
			}

			conflict = make([]string, len(dpoPrimaryKeyColumns))
			copy(conflict, dpoPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"dpo\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(dpoType, dpoMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dpoType, dpoMapping, ret)
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

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert dpo")
	}

	if !cached {
		dpoUpsertCacheMut.Lock()
		dpoUpsertCache[key] = cache
		dpoUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Dpo record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Dpo) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Dpo provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dpoPrimaryKeyMapping)
	sql := "DELETE FROM \"dpo\" WHERE \"demand_partner_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from dpo")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for dpo")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q dpoQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no dpoQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from dpo")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for dpo")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DpoSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(dpoBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dpoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"dpo\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, dpoPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from dpo slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for dpo")
	}

	if len(dpoAfterDeleteHooks) != 0 {
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
func (o *Dpo) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDpo(ctx, exec, o.DemandPartnerID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DpoSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DpoSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dpoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"dpo\".* FROM \"dpo\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, dpoPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DpoSlice")
	}

	*o = slice

	return nil
}

// DpoExists checks if the Dpo row exists.
func DpoExists(ctx context.Context, exec boil.ContextExecutor, demandPartnerID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"dpo\" where \"demand_partner_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, demandPartnerID)
	}
	row := exec.QueryRowContext(ctx, sql, demandPartnerID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if dpo exists")
	}

	return exists, nil
}

// Exists checks if the Dpo row exists.
func (o *Dpo) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return DpoExists(ctx, exec, o.DemandPartnerID)
}
