// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// IiqDaily is an object representing the database table.
type IiqDaily struct {
	Time       time.Time `boil:"time" json:"time" toml:"time" yaml:"time"`
	Dpid       string    `boil:"dpid" json:"dpid" toml:"dpid" yaml:"dpid"`
	Datacenter string    `boil:"datacenter" json:"datacenter" toml:"datacenter" yaml:"datacenter"`
	Request    int64     `boil:"request" json:"request" toml:"request" yaml:"request"`
	Response   int64     `boil:"response" json:"response" toml:"response" yaml:"response"`
	Impression int64     `boil:"impression" json:"impression" toml:"impression" yaml:"impression"`
	Revenue    float64   `boil:"revenue" json:"revenue" toml:"revenue" yaml:"revenue"`

	R *iiqDailyR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L iiqDailyL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var IiqDailyColumns = struct {
	Time       string
	Dpid       string
	Datacenter string
	Request    string
	Response   string
	Impression string
	Revenue    string
}{
	Time:       "time",
	Dpid:       "dpid",
	Datacenter: "datacenter",
	Request:    "request",
	Response:   "response",
	Impression: "impression",
	Revenue:    "revenue",
}

var IiqDailyTableColumns = struct {
	Time       string
	Dpid       string
	Datacenter string
	Request    string
	Response   string
	Impression string
	Revenue    string
}{
	Time:       "iiq_daily.time",
	Dpid:       "iiq_daily.dpid",
	Datacenter: "iiq_daily.datacenter",
	Request:    "iiq_daily.request",
	Response:   "iiq_daily.response",
	Impression: "iiq_daily.impression",
	Revenue:    "iiq_daily.revenue",
}

// Generated where

var IiqDailyWhere = struct {
	Time       whereHelpertime_Time
	Dpid       whereHelperstring
	Datacenter whereHelperstring
	Request    whereHelperint64
	Response   whereHelperint64
	Impression whereHelperint64
	Revenue    whereHelperfloat64
}{
	Time:       whereHelpertime_Time{field: "\"iiq_daily\".\"time\""},
	Dpid:       whereHelperstring{field: "\"iiq_daily\".\"dpid\""},
	Datacenter: whereHelperstring{field: "\"iiq_daily\".\"datacenter\""},
	Request:    whereHelperint64{field: "\"iiq_daily\".\"request\""},
	Response:   whereHelperint64{field: "\"iiq_daily\".\"response\""},
	Impression: whereHelperint64{field: "\"iiq_daily\".\"impression\""},
	Revenue:    whereHelperfloat64{field: "\"iiq_daily\".\"revenue\""},
}

// IiqDailyRels is where relationship names are stored.
var IiqDailyRels = struct {
}{}

// iiqDailyR is where relationships are stored.
type iiqDailyR struct {
}

// NewStruct creates a new relationship struct
func (*iiqDailyR) NewStruct() *iiqDailyR {
	return &iiqDailyR{}
}

// iiqDailyL is where Load methods for each relationship are stored.
type iiqDailyL struct{}

var (
	iiqDailyAllColumns            = []string{"time", "dpid", "datacenter", "request", "response", "impression", "revenue"}
	iiqDailyColumnsWithoutDefault = []string{"time", "dpid"}
	iiqDailyColumnsWithDefault    = []string{"datacenter", "request", "response", "impression", "revenue"}
	iiqDailyPrimaryKeyColumns     = []string{"time", "dpid", "datacenter"}
	iiqDailyGeneratedColumns      = []string{}
)

type (
	// IiqDailySlice is an alias for a slice of pointers to IiqDaily.
	// This should almost always be used instead of []IiqDaily.
	IiqDailySlice []*IiqDaily
	// IiqDailyHook is the signature for custom IiqDaily hook methods
	IiqDailyHook func(context.Context, boil.ContextExecutor, *IiqDaily) error

	iiqDailyQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	iiqDailyType                 = reflect.TypeOf(&IiqDaily{})
	iiqDailyMapping              = queries.MakeStructMapping(iiqDailyType)
	iiqDailyPrimaryKeyMapping, _ = queries.BindMapping(iiqDailyType, iiqDailyMapping, iiqDailyPrimaryKeyColumns)
	iiqDailyInsertCacheMut       sync.RWMutex
	iiqDailyInsertCache          = make(map[string]insertCache)
	iiqDailyUpdateCacheMut       sync.RWMutex
	iiqDailyUpdateCache          = make(map[string]updateCache)
	iiqDailyUpsertCacheMut       sync.RWMutex
	iiqDailyUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var iiqDailyAfterSelectHooks []IiqDailyHook

var iiqDailyBeforeInsertHooks []IiqDailyHook
var iiqDailyAfterInsertHooks []IiqDailyHook

var iiqDailyBeforeUpdateHooks []IiqDailyHook
var iiqDailyAfterUpdateHooks []IiqDailyHook

var iiqDailyBeforeDeleteHooks []IiqDailyHook
var iiqDailyAfterDeleteHooks []IiqDailyHook

var iiqDailyBeforeUpsertHooks []IiqDailyHook
var iiqDailyAfterUpsertHooks []IiqDailyHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *IiqDaily) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range iiqDailyAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *IiqDaily) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range iiqDailyBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *IiqDaily) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range iiqDailyAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *IiqDaily) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range iiqDailyBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *IiqDaily) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range iiqDailyAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *IiqDaily) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range iiqDailyBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *IiqDaily) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range iiqDailyAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *IiqDaily) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range iiqDailyBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *IiqDaily) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range iiqDailyAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddIiqDailyHook registers your hook function for all future operations.
func AddIiqDailyHook(hookPoint boil.HookPoint, iiqDailyHook IiqDailyHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		iiqDailyAfterSelectHooks = append(iiqDailyAfterSelectHooks, iiqDailyHook)
	case boil.BeforeInsertHook:
		iiqDailyBeforeInsertHooks = append(iiqDailyBeforeInsertHooks, iiqDailyHook)
	case boil.AfterInsertHook:
		iiqDailyAfterInsertHooks = append(iiqDailyAfterInsertHooks, iiqDailyHook)
	case boil.BeforeUpdateHook:
		iiqDailyBeforeUpdateHooks = append(iiqDailyBeforeUpdateHooks, iiqDailyHook)
	case boil.AfterUpdateHook:
		iiqDailyAfterUpdateHooks = append(iiqDailyAfterUpdateHooks, iiqDailyHook)
	case boil.BeforeDeleteHook:
		iiqDailyBeforeDeleteHooks = append(iiqDailyBeforeDeleteHooks, iiqDailyHook)
	case boil.AfterDeleteHook:
		iiqDailyAfterDeleteHooks = append(iiqDailyAfterDeleteHooks, iiqDailyHook)
	case boil.BeforeUpsertHook:
		iiqDailyBeforeUpsertHooks = append(iiqDailyBeforeUpsertHooks, iiqDailyHook)
	case boil.AfterUpsertHook:
		iiqDailyAfterUpsertHooks = append(iiqDailyAfterUpsertHooks, iiqDailyHook)
	}
}

// One returns a single iiqDaily record from the query.
func (q iiqDailyQuery) One(ctx context.Context, exec boil.ContextExecutor) (*IiqDaily, error) {
	o := &IiqDaily{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for iiq_daily")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all IiqDaily records from the query.
func (q iiqDailyQuery) All(ctx context.Context, exec boil.ContextExecutor) (IiqDailySlice, error) {
	var o []*IiqDaily

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to IiqDaily slice")
	}

	if len(iiqDailyAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all IiqDaily records in the query.
func (q iiqDailyQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count iiq_daily rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q iiqDailyQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if iiq_daily exists")
	}

	return count > 0, nil
}

// IiqDailies retrieves all the records using an executor.
func IiqDailies(mods ...qm.QueryMod) iiqDailyQuery {
	mods = append(mods, qm.From("\"iiq_daily\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"iiq_daily\".*"})
	}

	return iiqDailyQuery{q}
}

// FindIiqDaily retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindIiqDaily(ctx context.Context, exec boil.ContextExecutor, time time.Time, dpid string, datacenter string, selectCols ...string) (*IiqDaily, error) {
	iiqDailyObj := &IiqDaily{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"iiq_daily\" where \"time\"=$1 AND \"dpid\"=$2 AND \"datacenter\"=$3", sel,
	)

	q := queries.Raw(query, time, dpid, datacenter)

	err := q.Bind(ctx, exec, iiqDailyObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from iiq_daily")
	}

	if err = iiqDailyObj.doAfterSelectHooks(ctx, exec); err != nil {
		return iiqDailyObj, err
	}

	return iiqDailyObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *IiqDaily) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no iiq_daily provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(iiqDailyColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	iiqDailyInsertCacheMut.RLock()
	cache, cached := iiqDailyInsertCache[key]
	iiqDailyInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			iiqDailyAllColumns,
			iiqDailyColumnsWithDefault,
			iiqDailyColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(iiqDailyType, iiqDailyMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(iiqDailyType, iiqDailyMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"iiq_daily\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"iiq_daily\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into iiq_daily")
	}

	if !cached {
		iiqDailyInsertCacheMut.Lock()
		iiqDailyInsertCache[key] = cache
		iiqDailyInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the IiqDaily.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *IiqDaily) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	iiqDailyUpdateCacheMut.RLock()
	cache, cached := iiqDailyUpdateCache[key]
	iiqDailyUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			iiqDailyAllColumns,
			iiqDailyPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update iiq_daily, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"iiq_daily\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, iiqDailyPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(iiqDailyType, iiqDailyMapping, append(wl, iiqDailyPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update iiq_daily row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for iiq_daily")
	}

	if !cached {
		iiqDailyUpdateCacheMut.Lock()
		iiqDailyUpdateCache[key] = cache
		iiqDailyUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q iiqDailyQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for iiq_daily")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for iiq_daily")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o IiqDailySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), iiqDailyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"iiq_daily\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, iiqDailyPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in iiqDaily slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all iiqDaily")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *IiqDaily) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no iiq_daily provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(iiqDailyColumnsWithDefault, o)

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

	iiqDailyUpsertCacheMut.RLock()
	cache, cached := iiqDailyUpsertCache[key]
	iiqDailyUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			iiqDailyAllColumns,
			iiqDailyColumnsWithDefault,
			iiqDailyColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			iiqDailyAllColumns,
			iiqDailyPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert iiq_daily, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(iiqDailyPrimaryKeyColumns))
			copy(conflict, iiqDailyPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"iiq_daily\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(iiqDailyType, iiqDailyMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(iiqDailyType, iiqDailyMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert iiq_daily")
	}

	if !cached {
		iiqDailyUpsertCacheMut.Lock()
		iiqDailyUpsertCache[key] = cache
		iiqDailyUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single IiqDaily record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *IiqDaily) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no IiqDaily provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), iiqDailyPrimaryKeyMapping)
	sql := "DELETE FROM \"iiq_daily\" WHERE \"time\"=$1 AND \"dpid\"=$2 AND \"datacenter\"=$3"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from iiq_daily")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for iiq_daily")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q iiqDailyQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no iiqDailyQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from iiq_daily")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for iiq_daily")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o IiqDailySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(iiqDailyBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), iiqDailyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"iiq_daily\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, iiqDailyPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from iiqDaily slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for iiq_daily")
	}

	if len(iiqDailyAfterDeleteHooks) != 0 {
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
func (o *IiqDaily) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindIiqDaily(ctx, exec, o.Time, o.Dpid, o.Datacenter)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *IiqDailySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := IiqDailySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), iiqDailyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"iiq_daily\".* FROM \"iiq_daily\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, iiqDailyPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in IiqDailySlice")
	}

	*o = slice

	return nil
}

// IiqDailyExists checks if the IiqDaily row exists.
func IiqDailyExists(ctx context.Context, exec boil.ContextExecutor, time time.Time, dpid string, datacenter string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"iiq_daily\" where \"time\"=$1 AND \"dpid\"=$2 AND \"datacenter\"=$3 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, time, dpid, datacenter)
	}
	row := exec.QueryRowContext(ctx, sql, time, dpid, datacenter)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if iiq_daily exists")
	}

	return exists, nil
}

// Exists checks if the IiqDaily row exists.
func (o *IiqDaily) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return IiqDailyExists(ctx, exec, o.Time, o.Dpid, o.Datacenter)
}
