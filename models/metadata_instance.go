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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// MetadataInstance is an object representing the database table.
type MetadataInstance struct {
	InstanceID string    `boil:"instance_id" json:"instance_id" toml:"instance_id" yaml:"instance_id"`
	Bitwise    int64     `boil:"bitwise" json:"bitwise" toml:"bitwise" yaml:"bitwise"`
	Type       string    `boil:"type" json:"type" toml:"type" yaml:"type"`
	Config     null.JSON `boil:"config" json:"config,omitempty" toml:"config" yaml:"config,omitempty"`

	R *metadataInstanceR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L metadataInstanceL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MetadataInstanceColumns = struct {
	InstanceID string
	Bitwise    string
	Type       string
	Config     string
}{
	InstanceID: "instance_id",
	Bitwise:    "bitwise",
	Type:       "type",
	Config:     "config",
}

var MetadataInstanceTableColumns = struct {
	InstanceID string
	Bitwise    string
	Type       string
	Config     string
}{
	InstanceID: "metadata_instance.instance_id",
	Bitwise:    "metadata_instance.bitwise",
	Type:       "metadata_instance.type",
	Config:     "metadata_instance.config",
}

// Generated where

type whereHelpernull_JSON struct{ field string }

func (w whereHelpernull_JSON) EQ(x null.JSON) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_JSON) NEQ(x null.JSON) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_JSON) LT(x null.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_JSON) LTE(x null.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_JSON) GT(x null.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_JSON) GTE(x null.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_JSON) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_JSON) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var MetadataInstanceWhere = struct {
	InstanceID whereHelperstring
	Bitwise    whereHelperint64
	Type       whereHelperstring
	Config     whereHelpernull_JSON
}{
	InstanceID: whereHelperstring{field: "\"metadata_instance\".\"instance_id\""},
	Bitwise:    whereHelperint64{field: "\"metadata_instance\".\"bitwise\""},
	Type:       whereHelperstring{field: "\"metadata_instance\".\"type\""},
	Config:     whereHelpernull_JSON{field: "\"metadata_instance\".\"config\""},
}

// MetadataInstanceRels is where relationship names are stored.
var MetadataInstanceRels = struct {
}{}

// metadataInstanceR is where relationships are stored.
type metadataInstanceR struct {
}

// NewStruct creates a new relationship struct
func (*metadataInstanceR) NewStruct() *metadataInstanceR {
	return &metadataInstanceR{}
}

// metadataInstanceL is where Load methods for each relationship are stored.
type metadataInstanceL struct{}

var (
	metadataInstanceAllColumns            = []string{"instance_id", "bitwise", "type", "config"}
	metadataInstanceColumnsWithoutDefault = []string{"instance_id", "bitwise", "type"}
	metadataInstanceColumnsWithDefault    = []string{"config"}
	metadataInstancePrimaryKeyColumns     = []string{"instance_id"}
	metadataInstanceGeneratedColumns      = []string{}
)

type (
	// MetadataInstanceSlice is an alias for a slice of pointers to MetadataInstance.
	// This should almost always be used instead of []MetadataInstance.
	MetadataInstanceSlice []*MetadataInstance
	// MetadataInstanceHook is the signature for custom MetadataInstance hook methods
	MetadataInstanceHook func(context.Context, boil.ContextExecutor, *MetadataInstance) error

	metadataInstanceQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	metadataInstanceType                 = reflect.TypeOf(&MetadataInstance{})
	metadataInstanceMapping              = queries.MakeStructMapping(metadataInstanceType)
	metadataInstancePrimaryKeyMapping, _ = queries.BindMapping(metadataInstanceType, metadataInstanceMapping, metadataInstancePrimaryKeyColumns)
	metadataInstanceInsertCacheMut       sync.RWMutex
	metadataInstanceInsertCache          = make(map[string]insertCache)
	metadataInstanceUpdateCacheMut       sync.RWMutex
	metadataInstanceUpdateCache          = make(map[string]updateCache)
	metadataInstanceUpsertCacheMut       sync.RWMutex
	metadataInstanceUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var metadataInstanceAfterSelectHooks []MetadataInstanceHook

var metadataInstanceBeforeInsertHooks []MetadataInstanceHook
var metadataInstanceAfterInsertHooks []MetadataInstanceHook

var metadataInstanceBeforeUpdateHooks []MetadataInstanceHook
var metadataInstanceAfterUpdateHooks []MetadataInstanceHook

var metadataInstanceBeforeDeleteHooks []MetadataInstanceHook
var metadataInstanceAfterDeleteHooks []MetadataInstanceHook

var metadataInstanceBeforeUpsertHooks []MetadataInstanceHook
var metadataInstanceAfterUpsertHooks []MetadataInstanceHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *MetadataInstance) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range metadataInstanceAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *MetadataInstance) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range metadataInstanceBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *MetadataInstance) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range metadataInstanceAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *MetadataInstance) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range metadataInstanceBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *MetadataInstance) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range metadataInstanceAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *MetadataInstance) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range metadataInstanceBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *MetadataInstance) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range metadataInstanceAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *MetadataInstance) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range metadataInstanceBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *MetadataInstance) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range metadataInstanceAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMetadataInstanceHook registers your hook function for all future operations.
func AddMetadataInstanceHook(hookPoint boil.HookPoint, metadataInstanceHook MetadataInstanceHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		metadataInstanceAfterSelectHooks = append(metadataInstanceAfterSelectHooks, metadataInstanceHook)
	case boil.BeforeInsertHook:
		metadataInstanceBeforeInsertHooks = append(metadataInstanceBeforeInsertHooks, metadataInstanceHook)
	case boil.AfterInsertHook:
		metadataInstanceAfterInsertHooks = append(metadataInstanceAfterInsertHooks, metadataInstanceHook)
	case boil.BeforeUpdateHook:
		metadataInstanceBeforeUpdateHooks = append(metadataInstanceBeforeUpdateHooks, metadataInstanceHook)
	case boil.AfterUpdateHook:
		metadataInstanceAfterUpdateHooks = append(metadataInstanceAfterUpdateHooks, metadataInstanceHook)
	case boil.BeforeDeleteHook:
		metadataInstanceBeforeDeleteHooks = append(metadataInstanceBeforeDeleteHooks, metadataInstanceHook)
	case boil.AfterDeleteHook:
		metadataInstanceAfterDeleteHooks = append(metadataInstanceAfterDeleteHooks, metadataInstanceHook)
	case boil.BeforeUpsertHook:
		metadataInstanceBeforeUpsertHooks = append(metadataInstanceBeforeUpsertHooks, metadataInstanceHook)
	case boil.AfterUpsertHook:
		metadataInstanceAfterUpsertHooks = append(metadataInstanceAfterUpsertHooks, metadataInstanceHook)
	}
}

// One returns a single metadataInstance record from the query.
func (q metadataInstanceQuery) One(ctx context.Context, exec boil.ContextExecutor) (*MetadataInstance, error) {
	o := &MetadataInstance{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for metadata_instance")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all MetadataInstance records from the query.
func (q metadataInstanceQuery) All(ctx context.Context, exec boil.ContextExecutor) (MetadataInstanceSlice, error) {
	var o []*MetadataInstance

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to MetadataInstance slice")
	}

	if len(metadataInstanceAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all MetadataInstance records in the query.
func (q metadataInstanceQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count metadata_instance rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q metadataInstanceQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if metadata_instance exists")
	}

	return count > 0, nil
}

// MetadataInstances retrieves all the records using an executor.
func MetadataInstances(mods ...qm.QueryMod) metadataInstanceQuery {
	mods = append(mods, qm.From("\"metadata_instance\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"metadata_instance\".*"})
	}

	return metadataInstanceQuery{q}
}

// FindMetadataInstance retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMetadataInstance(ctx context.Context, exec boil.ContextExecutor, instanceID string, selectCols ...string) (*MetadataInstance, error) {
	metadataInstanceObj := &MetadataInstance{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"metadata_instance\" where \"instance_id\"=$1", sel,
	)

	q := queries.Raw(query, instanceID)

	err := q.Bind(ctx, exec, metadataInstanceObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from metadata_instance")
	}

	if err = metadataInstanceObj.doAfterSelectHooks(ctx, exec); err != nil {
		return metadataInstanceObj, err
	}

	return metadataInstanceObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *MetadataInstance) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no metadata_instance provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(metadataInstanceColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	metadataInstanceInsertCacheMut.RLock()
	cache, cached := metadataInstanceInsertCache[key]
	metadataInstanceInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			metadataInstanceAllColumns,
			metadataInstanceColumnsWithDefault,
			metadataInstanceColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(metadataInstanceType, metadataInstanceMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(metadataInstanceType, metadataInstanceMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"metadata_instance\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"metadata_instance\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into metadata_instance")
	}

	if !cached {
		metadataInstanceInsertCacheMut.Lock()
		metadataInstanceInsertCache[key] = cache
		metadataInstanceInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the MetadataInstance.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *MetadataInstance) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	metadataInstanceUpdateCacheMut.RLock()
	cache, cached := metadataInstanceUpdateCache[key]
	metadataInstanceUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			metadataInstanceAllColumns,
			metadataInstancePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update metadata_instance, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"metadata_instance\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, metadataInstancePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(metadataInstanceType, metadataInstanceMapping, append(wl, metadataInstancePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update metadata_instance row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for metadata_instance")
	}

	if !cached {
		metadataInstanceUpdateCacheMut.Lock()
		metadataInstanceUpdateCache[key] = cache
		metadataInstanceUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q metadataInstanceQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for metadata_instance")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for metadata_instance")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MetadataInstanceSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), metadataInstancePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"metadata_instance\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, metadataInstancePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in metadataInstance slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all metadataInstance")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *MetadataInstance) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no metadata_instance provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(metadataInstanceColumnsWithDefault, o)

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

	metadataInstanceUpsertCacheMut.RLock()
	cache, cached := metadataInstanceUpsertCache[key]
	metadataInstanceUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			metadataInstanceAllColumns,
			metadataInstanceColumnsWithDefault,
			metadataInstanceColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			metadataInstanceAllColumns,
			metadataInstancePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert metadata_instance, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(metadataInstancePrimaryKeyColumns))
			copy(conflict, metadataInstancePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"metadata_instance\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(metadataInstanceType, metadataInstanceMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(metadataInstanceType, metadataInstanceMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert metadata_instance")
	}

	if !cached {
		metadataInstanceUpsertCacheMut.Lock()
		metadataInstanceUpsertCache[key] = cache
		metadataInstanceUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single MetadataInstance record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *MetadataInstance) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no MetadataInstance provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), metadataInstancePrimaryKeyMapping)
	sql := "DELETE FROM \"metadata_instance\" WHERE \"instance_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from metadata_instance")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for metadata_instance")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q metadataInstanceQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no metadataInstanceQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from metadata_instance")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for metadata_instance")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MetadataInstanceSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(metadataInstanceBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), metadataInstancePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"metadata_instance\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, metadataInstancePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from metadataInstance slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for metadata_instance")
	}

	if len(metadataInstanceAfterDeleteHooks) != 0 {
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
func (o *MetadataInstance) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMetadataInstance(ctx, exec, o.InstanceID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MetadataInstanceSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MetadataInstanceSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), metadataInstancePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"metadata_instance\".* FROM \"metadata_instance\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, metadataInstancePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MetadataInstanceSlice")
	}

	*o = slice

	return nil
}

// MetadataInstanceExists checks if the MetadataInstance row exists.
func MetadataInstanceExists(ctx context.Context, exec boil.ContextExecutor, instanceID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"metadata_instance\" where \"instance_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, instanceID)
	}
	row := exec.QueryRowContext(ctx, sql, instanceID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if metadata_instance exists")
	}

	return exists, nil
}

// Exists checks if the MetadataInstance row exists.
func (o *MetadataInstance) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return MetadataInstanceExists(ctx, exec, o.InstanceID)
}
