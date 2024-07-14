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

// ID5Testing is an object representing the database table.
type ID5Testing struct {
	Time              time.Time `boil:"time" json:"time" toml:"time" yaml:"time"`
	DemandPartnerID   string    `boil:"demand_partner_id" json:"demand_partner_id" toml:"demand_partner_id" yaml:"demand_partner_id"`
	ID5Requests       int64     `boil:"id5_requests" json:"id5_requests" toml:"id5_requests" yaml:"id5_requests"`
	NonID5Requests    int64     `boil:"non_id5_requests" json:"non_id5_requests" toml:"non_id5_requests" yaml:"non_id5_requests"`
	ID5Impressions    int64     `boil:"id5_impressions" json:"id5_impressions" toml:"id5_impressions" yaml:"id5_impressions"`
	NonID5Impressions int64     `boil:"non_id5_impressions" json:"non_id5_impressions" toml:"non_id5_impressions" yaml:"non_id5_impressions"`

	R *id5TestingR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L id5TestingL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ID5TestingColumns = struct {
	Time              string
	DemandPartnerID   string
	ID5Requests       string
	NonID5Requests    string
	ID5Impressions    string
	NonID5Impressions string
}{
	Time:              "time",
	DemandPartnerID:   "demand_partner_id",
	ID5Requests:       "id5_requests",
	NonID5Requests:    "non_id5_requests",
	ID5Impressions:    "id5_impressions",
	NonID5Impressions: "non_id5_impressions",
}

var ID5TestingTableColumns = struct {
	Time              string
	DemandPartnerID   string
	ID5Requests       string
	NonID5Requests    string
	ID5Impressions    string
	NonID5Impressions string
}{
	Time:              "id5_testing.time",
	DemandPartnerID:   "id5_testing.demand_partner_id",
	ID5Requests:       "id5_testing.id5_requests",
	NonID5Requests:    "id5_testing.non_id5_requests",
	ID5Impressions:    "id5_testing.id5_impressions",
	NonID5Impressions: "id5_testing.non_id5_impressions",
}

// Generated where

var ID5TestingWhere = struct {
	Time              whereHelpertime_Time
	DemandPartnerID   whereHelperstring
	ID5Requests       whereHelperint64
	NonID5Requests    whereHelperint64
	ID5Impressions    whereHelperint64
	NonID5Impressions whereHelperint64
}{
	Time:              whereHelpertime_Time{field: "\"id5_testing\".\"time\""},
	DemandPartnerID:   whereHelperstring{field: "\"id5_testing\".\"demand_partner_id\""},
	ID5Requests:       whereHelperint64{field: "\"id5_testing\".\"id5_requests\""},
	NonID5Requests:    whereHelperint64{field: "\"id5_testing\".\"non_id5_requests\""},
	ID5Impressions:    whereHelperint64{field: "\"id5_testing\".\"id5_impressions\""},
	NonID5Impressions: whereHelperint64{field: "\"id5_testing\".\"non_id5_impressions\""},
}

// ID5TestingRels is where relationship names are stored.
var ID5TestingRels = struct {
}{}

// id5TestingR is where relationships are stored.
type id5TestingR struct {
}

// NewStruct creates a new relationship struct
func (*id5TestingR) NewStruct() *id5TestingR {
	return &id5TestingR{}
}

// id5TestingL is where Load methods for each relationship are stored.
type id5TestingL struct{}

var (
	id5TestingAllColumns            = []string{"time", "demand_partner_id", "id5_requests", "non_id5_requests", "id5_impressions", "non_id5_impressions"}
	id5TestingColumnsWithoutDefault = []string{"time", "demand_partner_id", "id5_requests", "non_id5_requests", "id5_impressions", "non_id5_impressions"}
	id5TestingColumnsWithDefault    = []string{}
	id5TestingPrimaryKeyColumns     = []string{"time", "demand_partner_id"}
	id5TestingGeneratedColumns      = []string{}
)

type (
	// ID5TestingSlice is an alias for a slice of pointers to ID5Testing.
	// This should almost always be used instead of []ID5Testing.
	ID5TestingSlice []*ID5Testing
	// ID5TestingHook is the signature for custom ID5Testing hook methods
	ID5TestingHook func(context.Context, boil.ContextExecutor, *ID5Testing) error

	id5TestingQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	id5TestingType                 = reflect.TypeOf(&ID5Testing{})
	id5TestingMapping              = queries.MakeStructMapping(id5TestingType)
	id5TestingPrimaryKeyMapping, _ = queries.BindMapping(id5TestingType, id5TestingMapping, id5TestingPrimaryKeyColumns)
	id5TestingInsertCacheMut       sync.RWMutex
	id5TestingInsertCache          = make(map[string]insertCache)
	id5TestingUpdateCacheMut       sync.RWMutex
	id5TestingUpdateCache          = make(map[string]updateCache)
	id5TestingUpsertCacheMut       sync.RWMutex
	id5TestingUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var id5TestingAfterSelectHooks []ID5TestingHook

var id5TestingBeforeInsertHooks []ID5TestingHook
var id5TestingAfterInsertHooks []ID5TestingHook

var id5TestingBeforeUpdateHooks []ID5TestingHook
var id5TestingAfterUpdateHooks []ID5TestingHook

var id5TestingBeforeDeleteHooks []ID5TestingHook
var id5TestingAfterDeleteHooks []ID5TestingHook

var id5TestingBeforeUpsertHooks []ID5TestingHook
var id5TestingAfterUpsertHooks []ID5TestingHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ID5Testing) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range id5TestingAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ID5Testing) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range id5TestingBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ID5Testing) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range id5TestingAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ID5Testing) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range id5TestingBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ID5Testing) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range id5TestingAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ID5Testing) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range id5TestingBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ID5Testing) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range id5TestingAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ID5Testing) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range id5TestingBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ID5Testing) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range id5TestingAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddID5TestingHook registers your hook function for all future operations.
func AddID5TestingHook(hookPoint boil.HookPoint, id5TestingHook ID5TestingHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		id5TestingAfterSelectHooks = append(id5TestingAfterSelectHooks, id5TestingHook)
	case boil.BeforeInsertHook:
		id5TestingBeforeInsertHooks = append(id5TestingBeforeInsertHooks, id5TestingHook)
	case boil.AfterInsertHook:
		id5TestingAfterInsertHooks = append(id5TestingAfterInsertHooks, id5TestingHook)
	case boil.BeforeUpdateHook:
		id5TestingBeforeUpdateHooks = append(id5TestingBeforeUpdateHooks, id5TestingHook)
	case boil.AfterUpdateHook:
		id5TestingAfterUpdateHooks = append(id5TestingAfterUpdateHooks, id5TestingHook)
	case boil.BeforeDeleteHook:
		id5TestingBeforeDeleteHooks = append(id5TestingBeforeDeleteHooks, id5TestingHook)
	case boil.AfterDeleteHook:
		id5TestingAfterDeleteHooks = append(id5TestingAfterDeleteHooks, id5TestingHook)
	case boil.BeforeUpsertHook:
		id5TestingBeforeUpsertHooks = append(id5TestingBeforeUpsertHooks, id5TestingHook)
	case boil.AfterUpsertHook:
		id5TestingAfterUpsertHooks = append(id5TestingAfterUpsertHooks, id5TestingHook)
	}
}

// One returns a single id5Testing record from the query.
func (q id5TestingQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ID5Testing, error) {
	o := &ID5Testing{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for id5_testing")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ID5Testing records from the query.
func (q id5TestingQuery) All(ctx context.Context, exec boil.ContextExecutor) (ID5TestingSlice, error) {
	var o []*ID5Testing

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ID5Testing slice")
	}

	if len(id5TestingAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ID5Testing records in the query.
func (q id5TestingQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count id5_testing rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q id5TestingQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if id5_testing exists")
	}

	return count > 0, nil
}

// ID5Testings retrieves all the records using an executor.
func ID5Testings(mods ...qm.QueryMod) id5TestingQuery {
	mods = append(mods, qm.From("\"id5_testing\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"id5_testing\".*"})
	}

	return id5TestingQuery{q}
}

// FindID5Testing retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindID5Testing(ctx context.Context, exec boil.ContextExecutor, time time.Time, demandPartnerID string, selectCols ...string) (*ID5Testing, error) {
	id5TestingObj := &ID5Testing{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"id5_testing\" where \"time\"=$1 AND \"demand_partner_id\"=$2", sel,
	)

	q := queries.Raw(query, time, demandPartnerID)

	err := q.Bind(ctx, exec, id5TestingObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from id5_testing")
	}

	if err = id5TestingObj.doAfterSelectHooks(ctx, exec); err != nil {
		return id5TestingObj, err
	}

	return id5TestingObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ID5Testing) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no id5_testing provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(id5TestingColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	id5TestingInsertCacheMut.RLock()
	cache, cached := id5TestingInsertCache[key]
	id5TestingInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			id5TestingAllColumns,
			id5TestingColumnsWithDefault,
			id5TestingColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(id5TestingType, id5TestingMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(id5TestingType, id5TestingMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"id5_testing\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"id5_testing\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into id5_testing")
	}

	if !cached {
		id5TestingInsertCacheMut.Lock()
		id5TestingInsertCache[key] = cache
		id5TestingInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ID5Testing.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ID5Testing) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	id5TestingUpdateCacheMut.RLock()
	cache, cached := id5TestingUpdateCache[key]
	id5TestingUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			id5TestingAllColumns,
			id5TestingPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update id5_testing, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"id5_testing\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, id5TestingPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(id5TestingType, id5TestingMapping, append(wl, id5TestingPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update id5_testing row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for id5_testing")
	}

	if !cached {
		id5TestingUpdateCacheMut.Lock()
		id5TestingUpdateCache[key] = cache
		id5TestingUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q id5TestingQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for id5_testing")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for id5_testing")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ID5TestingSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), id5TestingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"id5_testing\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, id5TestingPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in id5Testing slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all id5Testing")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ID5Testing) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no id5_testing provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(id5TestingColumnsWithDefault, o)

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

	id5TestingUpsertCacheMut.RLock()
	cache, cached := id5TestingUpsertCache[key]
	id5TestingUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			id5TestingAllColumns,
			id5TestingColumnsWithDefault,
			id5TestingColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			id5TestingAllColumns,
			id5TestingPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert id5_testing, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(id5TestingPrimaryKeyColumns))
			copy(conflict, id5TestingPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"id5_testing\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(id5TestingType, id5TestingMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(id5TestingType, id5TestingMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert id5_testing")
	}

	if !cached {
		id5TestingUpsertCacheMut.Lock()
		id5TestingUpsertCache[key] = cache
		id5TestingUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ID5Testing record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ID5Testing) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ID5Testing provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), id5TestingPrimaryKeyMapping)
	sql := "DELETE FROM \"id5_testing\" WHERE \"time\"=$1 AND \"demand_partner_id\"=$2"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from id5_testing")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for id5_testing")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q id5TestingQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no id5TestingQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from id5_testing")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for id5_testing")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ID5TestingSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(id5TestingBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), id5TestingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"id5_testing\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, id5TestingPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from id5Testing slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for id5_testing")
	}

	if len(id5TestingAfterDeleteHooks) != 0 {
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
func (o *ID5Testing) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindID5Testing(ctx, exec, o.Time, o.DemandPartnerID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ID5TestingSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ID5TestingSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), id5TestingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"id5_testing\".* FROM \"id5_testing\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, id5TestingPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ID5TestingSlice")
	}

	*o = slice

	return nil
}

// ID5TestingExists checks if the ID5Testing row exists.
func ID5TestingExists(ctx context.Context, exec boil.ContextExecutor, time time.Time, demandPartnerID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"id5_testing\" where \"time\"=$1 AND \"demand_partner_id\"=$2 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, time, demandPartnerID)
	}
	row := exec.QueryRowContext(ctx, sql, time, demandPartnerID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if id5_testing exists")
	}

	return exists, nil
}

// Exists checks if the ID5Testing row exists.
func (o *ID5Testing) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ID5TestingExists(ctx, exec, o.Time, o.DemandPartnerID)
}
