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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// DemandPartnerDaily is an object representing the database table.
type DemandPartnerDaily struct {
	Time            time.Time `boil:"time" json:"time" toml:"time" yaml:"time"`
	DemandPartnerID string    `boil:"demand_partner_id" json:"demand_partner_id" toml:"demand_partner_id" yaml:"demand_partner_id"`
	Domain          string    `boil:"domain" json:"domain" toml:"domain" yaml:"domain"`
	Impression      int64     `boil:"impression" json:"impression" toml:"impression" yaml:"impression"`
	Revenue         float64   `boil:"revenue" json:"revenue" toml:"revenue" yaml:"revenue"`

	R *demandPartnerDailyR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L demandPartnerDailyL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DemandPartnerDailyColumns = struct {
	Time            string
	DemandPartnerID string
	Domain          string
	Impression      string
	Revenue         string
}{
	Time:            "time",
	DemandPartnerID: "demand_partner_id",
	Domain:          "domain",
	Impression:      "impression",
	Revenue:         "revenue",
}

var DemandPartnerDailyTableColumns = struct {
	Time            string
	DemandPartnerID string
	Domain          string
	Impression      string
	Revenue         string
}{
	Time:            "demand_partner_daily.time",
	DemandPartnerID: "demand_partner_daily.demand_partner_id",
	Domain:          "demand_partner_daily.domain",
	Impression:      "demand_partner_daily.impression",
	Revenue:         "demand_partner_daily.revenue",
}

// Generated where

var DemandPartnerDailyWhere = struct {
	Time            whereHelpertime_Time
	DemandPartnerID whereHelperstring
	Domain          whereHelperstring
	Impression      whereHelperint64
	Revenue         whereHelperfloat64
}{
	Time:            whereHelpertime_Time{field: "\"demand_partner_daily\".\"time\""},
	DemandPartnerID: whereHelperstring{field: "\"demand_partner_daily\".\"demand_partner_id\""},
	Domain:          whereHelperstring{field: "\"demand_partner_daily\".\"domain\""},
	Impression:      whereHelperint64{field: "\"demand_partner_daily\".\"impression\""},
	Revenue:         whereHelperfloat64{field: "\"demand_partner_daily\".\"revenue\""},
}

// DemandPartnerDailyRels is where relationship names are stored.
var DemandPartnerDailyRels = struct {
}{}

// demandPartnerDailyR is where relationships are stored.
type demandPartnerDailyR struct {
}

// NewStruct creates a new relationship struct
func (*demandPartnerDailyR) NewStruct() *demandPartnerDailyR {
	return &demandPartnerDailyR{}
}

// demandPartnerDailyL is where Load methods for each relationship are stored.
type demandPartnerDailyL struct{}

var (
	demandPartnerDailyAllColumns            = []string{"time", "demand_partner_id", "domain", "impression", "revenue"}
	demandPartnerDailyColumnsWithoutDefault = []string{"time", "demand_partner_id", "domain"}
	demandPartnerDailyColumnsWithDefault    = []string{"impression", "revenue"}
	demandPartnerDailyPrimaryKeyColumns     = []string{"time", "demand_partner_id", "domain"}
	demandPartnerDailyGeneratedColumns      = []string{}
)

type (
	// DemandPartnerDailySlice is an alias for a slice of pointers to DemandPartnerDaily.
	// This should almost always be used instead of []DemandPartnerDaily.
	DemandPartnerDailySlice []*DemandPartnerDaily
	// DemandPartnerDailyHook is the signature for custom DemandPartnerDaily hook methods
	DemandPartnerDailyHook func(context.Context, boil.ContextExecutor, *DemandPartnerDaily) error

	demandPartnerDailyQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	demandPartnerDailyType                 = reflect.TypeOf(&DemandPartnerDaily{})
	demandPartnerDailyMapping              = queries.MakeStructMapping(demandPartnerDailyType)
	demandPartnerDailyPrimaryKeyMapping, _ = queries.BindMapping(demandPartnerDailyType, demandPartnerDailyMapping, demandPartnerDailyPrimaryKeyColumns)
	demandPartnerDailyInsertCacheMut       sync.RWMutex
	demandPartnerDailyInsertCache          = make(map[string]insertCache)
	demandPartnerDailyUpdateCacheMut       sync.RWMutex
	demandPartnerDailyUpdateCache          = make(map[string]updateCache)
	demandPartnerDailyUpsertCacheMut       sync.RWMutex
	demandPartnerDailyUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var demandPartnerDailyAfterSelectMu sync.Mutex
var demandPartnerDailyAfterSelectHooks []DemandPartnerDailyHook

var demandPartnerDailyBeforeInsertMu sync.Mutex
var demandPartnerDailyBeforeInsertHooks []DemandPartnerDailyHook
var demandPartnerDailyAfterInsertMu sync.Mutex
var demandPartnerDailyAfterInsertHooks []DemandPartnerDailyHook

var demandPartnerDailyBeforeUpdateMu sync.Mutex
var demandPartnerDailyBeforeUpdateHooks []DemandPartnerDailyHook
var demandPartnerDailyAfterUpdateMu sync.Mutex
var demandPartnerDailyAfterUpdateHooks []DemandPartnerDailyHook

var demandPartnerDailyBeforeDeleteMu sync.Mutex
var demandPartnerDailyBeforeDeleteHooks []DemandPartnerDailyHook
var demandPartnerDailyAfterDeleteMu sync.Mutex
var demandPartnerDailyAfterDeleteHooks []DemandPartnerDailyHook

var demandPartnerDailyBeforeUpsertMu sync.Mutex
var demandPartnerDailyBeforeUpsertHooks []DemandPartnerDailyHook
var demandPartnerDailyAfterUpsertMu sync.Mutex
var demandPartnerDailyAfterUpsertHooks []DemandPartnerDailyHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *DemandPartnerDaily) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range demandPartnerDailyAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *DemandPartnerDaily) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range demandPartnerDailyBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *DemandPartnerDaily) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range demandPartnerDailyAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *DemandPartnerDaily) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range demandPartnerDailyBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *DemandPartnerDaily) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range demandPartnerDailyAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *DemandPartnerDaily) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range demandPartnerDailyBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *DemandPartnerDaily) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range demandPartnerDailyAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *DemandPartnerDaily) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range demandPartnerDailyBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *DemandPartnerDaily) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range demandPartnerDailyAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDemandPartnerDailyHook registers your hook function for all future operations.
func AddDemandPartnerDailyHook(hookPoint boil.HookPoint, demandPartnerDailyHook DemandPartnerDailyHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		demandPartnerDailyAfterSelectMu.Lock()
		demandPartnerDailyAfterSelectHooks = append(demandPartnerDailyAfterSelectHooks, demandPartnerDailyHook)
		demandPartnerDailyAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		demandPartnerDailyBeforeInsertMu.Lock()
		demandPartnerDailyBeforeInsertHooks = append(demandPartnerDailyBeforeInsertHooks, demandPartnerDailyHook)
		demandPartnerDailyBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		demandPartnerDailyAfterInsertMu.Lock()
		demandPartnerDailyAfterInsertHooks = append(demandPartnerDailyAfterInsertHooks, demandPartnerDailyHook)
		demandPartnerDailyAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		demandPartnerDailyBeforeUpdateMu.Lock()
		demandPartnerDailyBeforeUpdateHooks = append(demandPartnerDailyBeforeUpdateHooks, demandPartnerDailyHook)
		demandPartnerDailyBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		demandPartnerDailyAfterUpdateMu.Lock()
		demandPartnerDailyAfterUpdateHooks = append(demandPartnerDailyAfterUpdateHooks, demandPartnerDailyHook)
		demandPartnerDailyAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		demandPartnerDailyBeforeDeleteMu.Lock()
		demandPartnerDailyBeforeDeleteHooks = append(demandPartnerDailyBeforeDeleteHooks, demandPartnerDailyHook)
		demandPartnerDailyBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		demandPartnerDailyAfterDeleteMu.Lock()
		demandPartnerDailyAfterDeleteHooks = append(demandPartnerDailyAfterDeleteHooks, demandPartnerDailyHook)
		demandPartnerDailyAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		demandPartnerDailyBeforeUpsertMu.Lock()
		demandPartnerDailyBeforeUpsertHooks = append(demandPartnerDailyBeforeUpsertHooks, demandPartnerDailyHook)
		demandPartnerDailyBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		demandPartnerDailyAfterUpsertMu.Lock()
		demandPartnerDailyAfterUpsertHooks = append(demandPartnerDailyAfterUpsertHooks, demandPartnerDailyHook)
		demandPartnerDailyAfterUpsertMu.Unlock()
	}
}

// One returns a single demandPartnerDaily record from the query.
func (q demandPartnerDailyQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DemandPartnerDaily, error) {
	o := &DemandPartnerDaily{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for demand_partner_daily")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all DemandPartnerDaily records from the query.
func (q demandPartnerDailyQuery) All(ctx context.Context, exec boil.ContextExecutor) (DemandPartnerDailySlice, error) {
	var o []*DemandPartnerDaily

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DemandPartnerDaily slice")
	}

	if len(demandPartnerDailyAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all DemandPartnerDaily records in the query.
func (q demandPartnerDailyQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count demand_partner_daily rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q demandPartnerDailyQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if demand_partner_daily exists")
	}

	return count > 0, nil
}

// DemandPartnerDailies retrieves all the records using an executor.
func DemandPartnerDailies(mods ...qm.QueryMod) demandPartnerDailyQuery {
	mods = append(mods, qm.From("\"demand_partner_daily\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"demand_partner_daily\".*"})
	}

	return demandPartnerDailyQuery{q}
}

// FindDemandPartnerDaily retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDemandPartnerDaily(ctx context.Context, exec boil.ContextExecutor, time time.Time, demandPartnerID string, domain string, selectCols ...string) (*DemandPartnerDaily, error) {
	demandPartnerDailyObj := &DemandPartnerDaily{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"demand_partner_daily\" where \"time\"=$1 AND \"demand_partner_id\"=$2 AND \"domain\"=$3", sel,
	)

	q := queries.Raw(query, time, demandPartnerID, domain)

	err := q.Bind(ctx, exec, demandPartnerDailyObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from demand_partner_daily")
	}

	if err = demandPartnerDailyObj.doAfterSelectHooks(ctx, exec); err != nil {
		return demandPartnerDailyObj, err
	}

	return demandPartnerDailyObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DemandPartnerDaily) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no demand_partner_daily provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(demandPartnerDailyColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	demandPartnerDailyInsertCacheMut.RLock()
	cache, cached := demandPartnerDailyInsertCache[key]
	demandPartnerDailyInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			demandPartnerDailyAllColumns,
			demandPartnerDailyColumnsWithDefault,
			demandPartnerDailyColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(demandPartnerDailyType, demandPartnerDailyMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(demandPartnerDailyType, demandPartnerDailyMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"demand_partner_daily\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"demand_partner_daily\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into demand_partner_daily")
	}

	if !cached {
		demandPartnerDailyInsertCacheMut.Lock()
		demandPartnerDailyInsertCache[key] = cache
		demandPartnerDailyInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the DemandPartnerDaily.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DemandPartnerDaily) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	demandPartnerDailyUpdateCacheMut.RLock()
	cache, cached := demandPartnerDailyUpdateCache[key]
	demandPartnerDailyUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			demandPartnerDailyAllColumns,
			demandPartnerDailyPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update demand_partner_daily, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"demand_partner_daily\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, demandPartnerDailyPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(demandPartnerDailyType, demandPartnerDailyMapping, append(wl, demandPartnerDailyPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update demand_partner_daily row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for demand_partner_daily")
	}

	if !cached {
		demandPartnerDailyUpdateCacheMut.Lock()
		demandPartnerDailyUpdateCache[key] = cache
		demandPartnerDailyUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q demandPartnerDailyQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for demand_partner_daily")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for demand_partner_daily")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DemandPartnerDailySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), demandPartnerDailyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"demand_partner_daily\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, demandPartnerDailyPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in demandPartnerDaily slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all demandPartnerDaily")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DemandPartnerDaily) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no demand_partner_daily provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(demandPartnerDailyColumnsWithDefault, o)

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

	demandPartnerDailyUpsertCacheMut.RLock()
	cache, cached := demandPartnerDailyUpsertCache[key]
	demandPartnerDailyUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			demandPartnerDailyAllColumns,
			demandPartnerDailyColumnsWithDefault,
			demandPartnerDailyColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			demandPartnerDailyAllColumns,
			demandPartnerDailyPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert demand_partner_daily, could not build update column list")
		}

		ret := strmangle.SetComplement(demandPartnerDailyAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(demandPartnerDailyPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert demand_partner_daily, could not build conflict column list")
			}

			conflict = make([]string, len(demandPartnerDailyPrimaryKeyColumns))
			copy(conflict, demandPartnerDailyPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"demand_partner_daily\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(demandPartnerDailyType, demandPartnerDailyMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(demandPartnerDailyType, demandPartnerDailyMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert demand_partner_daily")
	}

	if !cached {
		demandPartnerDailyUpsertCacheMut.Lock()
		demandPartnerDailyUpsertCache[key] = cache
		demandPartnerDailyUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single DemandPartnerDaily record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DemandPartnerDaily) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no DemandPartnerDaily provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), demandPartnerDailyPrimaryKeyMapping)
	sql := "DELETE FROM \"demand_partner_daily\" WHERE \"time\"=$1 AND \"demand_partner_id\"=$2 AND \"domain\"=$3"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from demand_partner_daily")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for demand_partner_daily")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q demandPartnerDailyQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no demandPartnerDailyQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from demand_partner_daily")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for demand_partner_daily")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DemandPartnerDailySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(demandPartnerDailyBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), demandPartnerDailyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"demand_partner_daily\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, demandPartnerDailyPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from demandPartnerDaily slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for demand_partner_daily")
	}

	if len(demandPartnerDailyAfterDeleteHooks) != 0 {
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
func (o *DemandPartnerDaily) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDemandPartnerDaily(ctx, exec, o.Time, o.DemandPartnerID, o.Domain)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DemandPartnerDailySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DemandPartnerDailySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), demandPartnerDailyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"demand_partner_daily\".* FROM \"demand_partner_daily\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, demandPartnerDailyPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DemandPartnerDailySlice")
	}

	*o = slice

	return nil
}

// DemandPartnerDailyExists checks if the DemandPartnerDaily row exists.
func DemandPartnerDailyExists(ctx context.Context, exec boil.ContextExecutor, time time.Time, demandPartnerID string, domain string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"demand_partner_daily\" where \"time\"=$1 AND \"demand_partner_id\"=$2 AND \"domain\"=$3 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, time, demandPartnerID, domain)
	}
	row := exec.QueryRowContext(ctx, sql, time, demandPartnerID, domain)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if demand_partner_daily exists")
	}

	return exists, nil
}

// Exists checks if the DemandPartnerDaily row exists.
func (o *DemandPartnerDaily) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return DemandPartnerDailyExists(ctx, exec, o.Time, o.DemandPartnerID, o.Domain)
}
