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
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// Factor is an object representing the database table.
type Factor struct {
	Publisher string  `boil:"publisher" json:"publisher" toml:"publisher" yaml:"publisher"`
	Domain    string  `boil:"domain" json:"domain" toml:"domain" yaml:"domain"`
	Device    string  `boil:"device" json:"device,omitempty" toml:"device" yaml:"device,omitempty"`
	Factor    float64 `boil:"factor" json:"factor,omitempty" toml:"factor" yaml:"factor,omitempty"`
	Country   string  `boil:"country" json:"country,omitempty" toml:"country" yaml:"country,omitempty"`

	R *factorR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L factorL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var FactorColumns = struct {
	Publisher string
	Domain    string
	Device    string
	Factor    string
	Country   string
}{
	Publisher: "publisher",
	Domain:    "domain",
	Device:    "device",
	Factor:    "factor",
	Country:   "country",
}

var FactorTableColumns = struct {
	Publisher string
	Domain    string
	Device    string
	Factor    string
	Country   string
}{
	Publisher: "factor.publisher",
	Domain:    "factor.domain",
	Device:    "factor.device",
	Factor:    "factor.factor",
	Country:   "factor.country",
}

// Generated where

type whereHelpertypes_NullDecimal struct{ field string }

func (w whereHelpertypes_NullDecimal) EQ(x types.NullDecimal) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpertypes_NullDecimal) NEQ(x types.NullDecimal) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpertypes_NullDecimal) LT(x types.NullDecimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_NullDecimal) LTE(x types.NullDecimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_NullDecimal) GT(x types.NullDecimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_NullDecimal) GTE(x types.NullDecimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpertypes_NullDecimal) IsNull() qm.QueryMod { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpertypes_NullDecimal) IsNotNull() qm.QueryMod {
	return qmhelper.WhereIsNotNull(w.field)
}

var FactorWhere = struct {
	Publisher whereHelperstring
	Domain    whereHelperstring
	Device    whereHelpernull_String
	Factor    whereHelpertypes_NullDecimal
	Country   whereHelpernull_String
}{
	Publisher: whereHelperstring{field: "\"factor\".\"publisher\""},
	Domain:    whereHelperstring{field: "\"factor\".\"domain\""},
	Device:    whereHelpernull_String{field: "\"factor\".\"device\""},
	Factor:    whereHelpertypes_NullDecimal{field: "\"factor\".\"factor\""},
	Country:   whereHelpernull_String{field: "\"factor\".\"country\""},
}

// FactorRels is where relationship names are stored.
var FactorRels = struct {
}{}

// factorR is where relationships are stored.
type factorR struct {
}

// NewStruct creates a new relationship struct
func (*factorR) NewStruct() *factorR {
	return &factorR{}
}

// factorL is where Load methods for each relationship are stored.
type factorL struct{}

var (
	factorAllColumns            = []string{"publisher", "domain", "device", "factor", "country"}
	factorColumnsWithoutDefault = []string{"publisher", "domain"}
	factorColumnsWithDefault    = []string{"device", "factor", "country"}
	factorPrimaryKeyColumns     = []string{"publisher", "domain"}
	factorGeneratedColumns      = []string{}
)

type (
	// FactorSlice is an alias for a slice of pointers to Factor.
	// This should almost always be used instead of []Factor.
	FactorSlice []*Factor
	// FactorHook is the signature for custom Factor hook methods
	FactorHook func(context.Context, boil.ContextExecutor, *Factor) error

	factorQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	factorType                 = reflect.TypeOf(&Factor{})
	factorMapping              = queries.MakeStructMapping(factorType)
	factorPrimaryKeyMapping, _ = queries.BindMapping(factorType, factorMapping, factorPrimaryKeyColumns)
	factorInsertCacheMut       sync.RWMutex
	factorInsertCache          = make(map[string]insertCache)
	factorUpdateCacheMut       sync.RWMutex
	factorUpdateCache          = make(map[string]updateCache)
	factorUpsertCacheMut       sync.RWMutex
	factorUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var factorAfterSelectMu sync.Mutex
var factorAfterSelectHooks []FactorHook

var factorBeforeInsertMu sync.Mutex
var factorBeforeInsertHooks []FactorHook
var factorAfterInsertMu sync.Mutex
var factorAfterInsertHooks []FactorHook

var factorBeforeUpdateMu sync.Mutex
var factorBeforeUpdateHooks []FactorHook
var factorAfterUpdateMu sync.Mutex
var factorAfterUpdateHooks []FactorHook

var factorBeforeDeleteMu sync.Mutex
var factorBeforeDeleteHooks []FactorHook
var factorAfterDeleteMu sync.Mutex
var factorAfterDeleteHooks []FactorHook

var factorBeforeUpsertMu sync.Mutex
var factorBeforeUpsertHooks []FactorHook
var factorAfterUpsertMu sync.Mutex
var factorAfterUpsertHooks []FactorHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Factor) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range factorAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Factor) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range factorBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Factor) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range factorAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Factor) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range factorBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Factor) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range factorAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Factor) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range factorBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Factor) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range factorAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Factor) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range factorBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Factor) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range factorAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFactorHook registers your hook function for all future operations.
func AddFactorHook(hookPoint boil.HookPoint, factorHook FactorHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		factorAfterSelectMu.Lock()
		factorAfterSelectHooks = append(factorAfterSelectHooks, factorHook)
		factorAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		factorBeforeInsertMu.Lock()
		factorBeforeInsertHooks = append(factorBeforeInsertHooks, factorHook)
		factorBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		factorAfterInsertMu.Lock()
		factorAfterInsertHooks = append(factorAfterInsertHooks, factorHook)
		factorAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		factorBeforeUpdateMu.Lock()
		factorBeforeUpdateHooks = append(factorBeforeUpdateHooks, factorHook)
		factorBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		factorAfterUpdateMu.Lock()
		factorAfterUpdateHooks = append(factorAfterUpdateHooks, factorHook)
		factorAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		factorBeforeDeleteMu.Lock()
		factorBeforeDeleteHooks = append(factorBeforeDeleteHooks, factorHook)
		factorBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		factorAfterDeleteMu.Lock()
		factorAfterDeleteHooks = append(factorAfterDeleteHooks, factorHook)
		factorAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		factorBeforeUpsertMu.Lock()
		factorBeforeUpsertHooks = append(factorBeforeUpsertHooks, factorHook)
		factorBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		factorAfterUpsertMu.Lock()
		factorAfterUpsertHooks = append(factorAfterUpsertHooks, factorHook)
		factorAfterUpsertMu.Unlock()
	}
}

// One returns a single factor record from the query.
func (q factorQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Factor, error) {
	o := &Factor{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for factor")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Factor records from the query.
func (q factorQuery) All(ctx context.Context, exec boil.ContextExecutor) (FactorSlice, error) {
	var o []*Factor

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Factor slice")
	}

	if len(factorAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Factor records in the query.
func (q factorQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count factor rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q factorQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if factor exists")
	}

	return count > 0, nil
}

// Factors retrieves all the records using an executor.
func Factors(mods ...qm.QueryMod) factorQuery {
	mods = append(mods, qm.From("\"factor\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"factor\".*"})
	}

	return factorQuery{q}
}

// FindFactor retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFactor(ctx context.Context, exec boil.ContextExecutor, publisher string, domain string, selectCols ...string) (*Factor, error) {
	factorObj := &Factor{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"factor\" where \"publisher\"=$1 AND \"domain\"=$2", sel,
	)

	q := queries.Raw(query, publisher, domain)

	err := q.Bind(ctx, exec, factorObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from factor")
	}

	if err = factorObj.doAfterSelectHooks(ctx, exec); err != nil {
		return factorObj, err
	}

	return factorObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Factor) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no factor provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(factorColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	factorInsertCacheMut.RLock()
	cache, cached := factorInsertCache[key]
	factorInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			factorAllColumns,
			factorColumnsWithDefault,
			factorColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(factorType, factorMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(factorType, factorMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"factor\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"factor\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into factor")
	}

	if !cached {
		factorInsertCacheMut.Lock()
		factorInsertCache[key] = cache
		factorInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Factor.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Factor) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	factorUpdateCacheMut.RLock()
	cache, cached := factorUpdateCache[key]
	factorUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			factorAllColumns,
			factorPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update factor, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"factor\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, factorPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(factorType, factorMapping, append(wl, factorPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update factor row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for factor")
	}

	if !cached {
		factorUpdateCacheMut.Lock()
		factorUpdateCache[key] = cache
		factorUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q factorQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for factor")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for factor")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FactorSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), factorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"factor\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, factorPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in factor slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all factor")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Factor) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no factor provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(factorColumnsWithDefault, o)

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

	factorUpsertCacheMut.RLock()
	cache, cached := factorUpsertCache[key]
	factorUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			factorAllColumns,
			factorColumnsWithDefault,
			factorColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			factorAllColumns,
			factorPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert factor, could not build update column list")
		}

		ret := strmangle.SetComplement(factorAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(factorPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert factor, could not build conflict column list")
			}

			conflict = make([]string, len(factorPrimaryKeyColumns))
			copy(conflict, factorPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"factor\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(factorType, factorMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(factorType, factorMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert factor")
	}

	if !cached {
		factorUpsertCacheMut.Lock()
		factorUpsertCache[key] = cache
		factorUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Factor record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Factor) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Factor provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), factorPrimaryKeyMapping)
	sql := "DELETE FROM \"factor\" WHERE \"publisher\"=$1 AND \"domain\"=$2"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from factor")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for factor")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q factorQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no factorQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from factor")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for factor")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FactorSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(factorBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), factorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"factor\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, factorPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from factor slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for factor")
	}

	if len(factorAfterDeleteHooks) != 0 {
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
func (o *Factor) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindFactor(ctx, exec, o.Publisher, o.Domain)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FactorSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := FactorSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), factorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"factor\".* FROM \"factor\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, factorPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FactorSlice")
	}

	*o = slice

	return nil
}

// FactorExists checks if the Factor row exists.
func FactorExists(ctx context.Context, exec boil.ContextExecutor, publisher string, domain string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"factor\" where \"publisher\"=$1 AND \"domain\"=$2 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, publisher, domain)
	}
	row := exec.QueryRowContext(ctx, sql, publisher, domain)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if factor exists")
	}

	return exists, nil
}

// Exists checks if the Factor row exists.
func (o *Factor) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return FactorExists(ctx, exec, o.Publisher, o.Domain)
}
