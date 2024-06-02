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

// CompassPublisherTag is an object representing the database table.
type CompassPublisherTag struct {
	ID          int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	PublisherID string `boil:"publisher_id" json:"publisher_id" toml:"publisher_id" yaml:"publisher_id"`
	DeviceType  string `boil:"device_type" json:"device_type" toml:"device_type" yaml:"device_type"`
	Domain      string `boil:"domain" json:"domain" toml:"domain" yaml:"domain"`

	R *compassPublisherTagR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L compassPublisherTagL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CompassPublisherTagColumns = struct {
	ID          string
	PublisherID string
	DeviceType  string
	Domain      string
}{
	ID:          "id",
	PublisherID: "publisher_id",
	DeviceType:  "device_type",
	Domain:      "domain",
}

var CompassPublisherTagTableColumns = struct {
	ID          string
	PublisherID string
	DeviceType  string
	Domain      string
}{
	ID:          "compass_publisher_tag.id",
	PublisherID: "compass_publisher_tag.publisher_id",
	DeviceType:  "compass_publisher_tag.device_type",
	Domain:      "compass_publisher_tag.domain",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var CompassPublisherTagWhere = struct {
	ID          whereHelperint
	PublisherID whereHelperstring
	DeviceType  whereHelperstring
	Domain      whereHelperstring
}{
	ID:          whereHelperint{field: "\"compass_publisher_tag\".\"id\""},
	PublisherID: whereHelperstring{field: "\"compass_publisher_tag\".\"publisher_id\""},
	DeviceType:  whereHelperstring{field: "\"compass_publisher_tag\".\"device_type\""},
	Domain:      whereHelperstring{field: "\"compass_publisher_tag\".\"domain\""},
}

// CompassPublisherTagRels is where relationship names are stored.
var CompassPublisherTagRels = struct {
}{}

// compassPublisherTagR is where relationships are stored.
type compassPublisherTagR struct {
}

// NewStruct creates a new relationship struct
func (*compassPublisherTagR) NewStruct() *compassPublisherTagR {
	return &compassPublisherTagR{}
}

// compassPublisherTagL is where Load methods for each relationship are stored.
type compassPublisherTagL struct{}

var (
	compassPublisherTagAllColumns            = []string{"id", "publisher_id", "device_type", "domain"}
	compassPublisherTagColumnsWithoutDefault = []string{"publisher_id", "device_type", "domain"}
	compassPublisherTagColumnsWithDefault    = []string{"id"}
	compassPublisherTagPrimaryKeyColumns     = []string{"id"}
	compassPublisherTagGeneratedColumns      = []string{}
)

type (
	// CompassPublisherTagSlice is an alias for a slice of pointers to CompassPublisherTag.
	// This should almost always be used instead of []CompassPublisherTag.
	CompassPublisherTagSlice []*CompassPublisherTag
	// CompassPublisherTagHook is the signature for custom CompassPublisherTag hook methods
	CompassPublisherTagHook func(context.Context, boil.ContextExecutor, *CompassPublisherTag) error

	compassPublisherTagQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	compassPublisherTagType                 = reflect.TypeOf(&CompassPublisherTag{})
	compassPublisherTagMapping              = queries.MakeStructMapping(compassPublisherTagType)
	compassPublisherTagPrimaryKeyMapping, _ = queries.BindMapping(compassPublisherTagType, compassPublisherTagMapping, compassPublisherTagPrimaryKeyColumns)
	compassPublisherTagInsertCacheMut       sync.RWMutex
	compassPublisherTagInsertCache          = make(map[string]insertCache)
	compassPublisherTagUpdateCacheMut       sync.RWMutex
	compassPublisherTagUpdateCache          = make(map[string]updateCache)
	compassPublisherTagUpsertCacheMut       sync.RWMutex
	compassPublisherTagUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var compassPublisherTagAfterSelectMu sync.Mutex
var compassPublisherTagAfterSelectHooks []CompassPublisherTagHook

var compassPublisherTagBeforeInsertMu sync.Mutex
var compassPublisherTagBeforeInsertHooks []CompassPublisherTagHook
var compassPublisherTagAfterInsertMu sync.Mutex
var compassPublisherTagAfterInsertHooks []CompassPublisherTagHook

var compassPublisherTagBeforeUpdateMu sync.Mutex
var compassPublisherTagBeforeUpdateHooks []CompassPublisherTagHook
var compassPublisherTagAfterUpdateMu sync.Mutex
var compassPublisherTagAfterUpdateHooks []CompassPublisherTagHook

var compassPublisherTagBeforeDeleteMu sync.Mutex
var compassPublisherTagBeforeDeleteHooks []CompassPublisherTagHook
var compassPublisherTagAfterDeleteMu sync.Mutex
var compassPublisherTagAfterDeleteHooks []CompassPublisherTagHook

var compassPublisherTagBeforeUpsertMu sync.Mutex
var compassPublisherTagBeforeUpsertHooks []CompassPublisherTagHook
var compassPublisherTagAfterUpsertMu sync.Mutex
var compassPublisherTagAfterUpsertHooks []CompassPublisherTagHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CompassPublisherTag) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range compassPublisherTagAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CompassPublisherTag) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range compassPublisherTagBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CompassPublisherTag) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range compassPublisherTagAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CompassPublisherTag) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range compassPublisherTagBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CompassPublisherTag) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range compassPublisherTagAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CompassPublisherTag) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range compassPublisherTagBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CompassPublisherTag) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range compassPublisherTagAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CompassPublisherTag) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range compassPublisherTagBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CompassPublisherTag) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range compassPublisherTagAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCompassPublisherTagHook registers your hook function for all future operations.
func AddCompassPublisherTagHook(hookPoint boil.HookPoint, compassPublisherTagHook CompassPublisherTagHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		compassPublisherTagAfterSelectMu.Lock()
		compassPublisherTagAfterSelectHooks = append(compassPublisherTagAfterSelectHooks, compassPublisherTagHook)
		compassPublisherTagAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		compassPublisherTagBeforeInsertMu.Lock()
		compassPublisherTagBeforeInsertHooks = append(compassPublisherTagBeforeInsertHooks, compassPublisherTagHook)
		compassPublisherTagBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		compassPublisherTagAfterInsertMu.Lock()
		compassPublisherTagAfterInsertHooks = append(compassPublisherTagAfterInsertHooks, compassPublisherTagHook)
		compassPublisherTagAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		compassPublisherTagBeforeUpdateMu.Lock()
		compassPublisherTagBeforeUpdateHooks = append(compassPublisherTagBeforeUpdateHooks, compassPublisherTagHook)
		compassPublisherTagBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		compassPublisherTagAfterUpdateMu.Lock()
		compassPublisherTagAfterUpdateHooks = append(compassPublisherTagAfterUpdateHooks, compassPublisherTagHook)
		compassPublisherTagAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		compassPublisherTagBeforeDeleteMu.Lock()
		compassPublisherTagBeforeDeleteHooks = append(compassPublisherTagBeforeDeleteHooks, compassPublisherTagHook)
		compassPublisherTagBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		compassPublisherTagAfterDeleteMu.Lock()
		compassPublisherTagAfterDeleteHooks = append(compassPublisherTagAfterDeleteHooks, compassPublisherTagHook)
		compassPublisherTagAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		compassPublisherTagBeforeUpsertMu.Lock()
		compassPublisherTagBeforeUpsertHooks = append(compassPublisherTagBeforeUpsertHooks, compassPublisherTagHook)
		compassPublisherTagBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		compassPublisherTagAfterUpsertMu.Lock()
		compassPublisherTagAfterUpsertHooks = append(compassPublisherTagAfterUpsertHooks, compassPublisherTagHook)
		compassPublisherTagAfterUpsertMu.Unlock()
	}
}

// One returns a single compassPublisherTag record from the query.
func (q compassPublisherTagQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CompassPublisherTag, error) {
	o := &CompassPublisherTag{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for compass_publisher_tag")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CompassPublisherTag records from the query.
func (q compassPublisherTagQuery) All(ctx context.Context, exec boil.ContextExecutor) (CompassPublisherTagSlice, error) {
	var o []*CompassPublisherTag

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CompassPublisherTag slice")
	}

	if len(compassPublisherTagAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CompassPublisherTag records in the query.
func (q compassPublisherTagQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count compass_publisher_tag rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q compassPublisherTagQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if compass_publisher_tag exists")
	}

	return count > 0, nil
}

// CompassPublisherTags retrieves all the records using an executor.
func CompassPublisherTags(mods ...qm.QueryMod) compassPublisherTagQuery {
	mods = append(mods, qm.From("\"compass_publisher_tag\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"compass_publisher_tag\".*"})
	}

	return compassPublisherTagQuery{q}
}

// FindCompassPublisherTag retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCompassPublisherTag(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*CompassPublisherTag, error) {
	compassPublisherTagObj := &CompassPublisherTag{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"compass_publisher_tag\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, compassPublisherTagObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from compass_publisher_tag")
	}

	if err = compassPublisherTagObj.doAfterSelectHooks(ctx, exec); err != nil {
		return compassPublisherTagObj, err
	}

	return compassPublisherTagObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CompassPublisherTag) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no compass_publisher_tag provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(compassPublisherTagColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	compassPublisherTagInsertCacheMut.RLock()
	cache, cached := compassPublisherTagInsertCache[key]
	compassPublisherTagInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			compassPublisherTagAllColumns,
			compassPublisherTagColumnsWithDefault,
			compassPublisherTagColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(compassPublisherTagType, compassPublisherTagMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(compassPublisherTagType, compassPublisherTagMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"compass_publisher_tag\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"compass_publisher_tag\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into compass_publisher_tag")
	}

	if !cached {
		compassPublisherTagInsertCacheMut.Lock()
		compassPublisherTagInsertCache[key] = cache
		compassPublisherTagInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CompassPublisherTag.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CompassPublisherTag) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	compassPublisherTagUpdateCacheMut.RLock()
	cache, cached := compassPublisherTagUpdateCache[key]
	compassPublisherTagUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			compassPublisherTagAllColumns,
			compassPublisherTagPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update compass_publisher_tag, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"compass_publisher_tag\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, compassPublisherTagPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(compassPublisherTagType, compassPublisherTagMapping, append(wl, compassPublisherTagPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update compass_publisher_tag row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for compass_publisher_tag")
	}

	if !cached {
		compassPublisherTagUpdateCacheMut.Lock()
		compassPublisherTagUpdateCache[key] = cache
		compassPublisherTagUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q compassPublisherTagQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for compass_publisher_tag")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for compass_publisher_tag")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CompassPublisherTagSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), compassPublisherTagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"compass_publisher_tag\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, compassPublisherTagPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in compassPublisherTag slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all compassPublisherTag")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CompassPublisherTag) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no compass_publisher_tag provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(compassPublisherTagColumnsWithDefault, o)

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

	compassPublisherTagUpsertCacheMut.RLock()
	cache, cached := compassPublisherTagUpsertCache[key]
	compassPublisherTagUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			compassPublisherTagAllColumns,
			compassPublisherTagColumnsWithDefault,
			compassPublisherTagColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			compassPublisherTagAllColumns,
			compassPublisherTagPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert compass_publisher_tag, could not build update column list")
		}

		ret := strmangle.SetComplement(compassPublisherTagAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(compassPublisherTagPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert compass_publisher_tag, could not build conflict column list")
			}

			conflict = make([]string, len(compassPublisherTagPrimaryKeyColumns))
			copy(conflict, compassPublisherTagPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"compass_publisher_tag\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(compassPublisherTagType, compassPublisherTagMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(compassPublisherTagType, compassPublisherTagMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert compass_publisher_tag")
	}

	if !cached {
		compassPublisherTagUpsertCacheMut.Lock()
		compassPublisherTagUpsertCache[key] = cache
		compassPublisherTagUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CompassPublisherTag record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CompassPublisherTag) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CompassPublisherTag provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), compassPublisherTagPrimaryKeyMapping)
	sql := "DELETE FROM \"compass_publisher_tag\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from compass_publisher_tag")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for compass_publisher_tag")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q compassPublisherTagQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no compassPublisherTagQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from compass_publisher_tag")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for compass_publisher_tag")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CompassPublisherTagSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(compassPublisherTagBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), compassPublisherTagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"compass_publisher_tag\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, compassPublisherTagPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from compassPublisherTag slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for compass_publisher_tag")
	}

	if len(compassPublisherTagAfterDeleteHooks) != 0 {
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
func (o *CompassPublisherTag) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCompassPublisherTag(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CompassPublisherTagSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CompassPublisherTagSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), compassPublisherTagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"compass_publisher_tag\".* FROM \"compass_publisher_tag\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, compassPublisherTagPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CompassPublisherTagSlice")
	}

	*o = slice

	return nil
}

// CompassPublisherTagExists checks if the CompassPublisherTag row exists.
func CompassPublisherTagExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"compass_publisher_tag\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if compass_publisher_tag exists")
	}

	return exists, nil
}

// Exists checks if the CompassPublisherTag row exists.
func (o *CompassPublisherTag) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return CompassPublisherTagExists(ctx, exec, o.ID)
}
