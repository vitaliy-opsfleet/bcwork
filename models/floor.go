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

// Floor is an object representing the database table.
type Floor struct {
	Publisher       string      `boil:"publisher" json:"publisher" toml:"publisher" yaml:"publisher"`
	Domain          string      `boil:"domain" json:"domain" toml:"domain" yaml:"domain"`
	Country         string      `boil:"country" json:"country" toml:"country" yaml:"country"`
	Device          string      `boil:"device" json:"device" toml:"device" yaml:"device"`
	Floor           float64     `boil:"floor" json:"floor" toml:"floor" yaml:"floor"`
	CreatedAt       time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt       null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	RuleID          string      `boil:"rule_id" json:"rule_id" toml:"rule_id" yaml:"rule_id"`
	DemandPartnerID string      `boil:"demand_partner_id" json:"demand_partner_id" toml:"demand_partner_id" yaml:"demand_partner_id"`
	Browser         null.String `boil:"browser" json:"browser,omitempty" toml:"browser" yaml:"browser,omitempty"`
	Os              null.String `boil:"os" json:"os,omitempty" toml:"os" yaml:"os,omitempty"`
	PlacementType   null.String `boil:"placement_type" json:"placement_type,omitempty" toml:"placement_type" yaml:"placement_type,omitempty"`

	R *floorR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L floorL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var FloorColumns = struct {
	Publisher       string
	Domain          string
	Country         string
	Device          string
	Floor           string
	CreatedAt       string
	UpdatedAt       string
	RuleID          string
	DemandPartnerID string
	Browser         string
	Os              string
	PlacementType   string
}{
	Publisher:       "publisher",
	Domain:          "domain",
	Country:         "country",
	Device:          "device",
	Floor:           "floor",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	RuleID:          "rule_id",
	DemandPartnerID: "demand_partner_id",
	Browser:         "browser",
	Os:              "os",
	PlacementType:   "placement_type",
}

var FloorTableColumns = struct {
	Publisher       string
	Domain          string
	Country         string
	Device          string
	Floor           string
	CreatedAt       string
	UpdatedAt       string
	RuleID          string
	DemandPartnerID string
	Browser         string
	Os              string
	PlacementType   string
}{
	Publisher:       "floor.publisher",
	Domain:          "floor.domain",
	Country:         "floor.country",
	Device:          "floor.device",
	Floor:           "floor.floor",
	CreatedAt:       "floor.created_at",
	UpdatedAt:       "floor.updated_at",
	RuleID:          "floor.rule_id",
	DemandPartnerID: "floor.demand_partner_id",
	Browser:         "floor.browser",
	Os:              "floor.os",
	PlacementType:   "floor.placement_type",
}

// Generated where

var FloorWhere = struct {
	Publisher       whereHelperstring
	Domain          whereHelperstring
	Country         whereHelperstring
	Device          whereHelperstring
	Floor           whereHelperfloat64
	CreatedAt       whereHelpertime_Time
	UpdatedAt       whereHelpernull_Time
	RuleID          whereHelperstring
	DemandPartnerID whereHelperstring
	Browser         whereHelpernull_String
	Os              whereHelpernull_String
	PlacementType   whereHelpernull_String
}{
	Publisher:       whereHelperstring{field: "\"floor\".\"publisher\""},
	Domain:          whereHelperstring{field: "\"floor\".\"domain\""},
	Country:         whereHelperstring{field: "\"floor\".\"country\""},
	Device:          whereHelperstring{field: "\"floor\".\"device\""},
	Floor:           whereHelperfloat64{field: "\"floor\".\"floor\""},
	CreatedAt:       whereHelpertime_Time{field: "\"floor\".\"created_at\""},
	UpdatedAt:       whereHelpernull_Time{field: "\"floor\".\"updated_at\""},
	RuleID:          whereHelperstring{field: "\"floor\".\"rule_id\""},
	DemandPartnerID: whereHelperstring{field: "\"floor\".\"demand_partner_id\""},
	Browser:         whereHelpernull_String{field: "\"floor\".\"browser\""},
	Os:              whereHelpernull_String{field: "\"floor\".\"os\""},
	PlacementType:   whereHelpernull_String{field: "\"floor\".\"placement_type\""},
}

// FloorRels is where relationship names are stored.
var FloorRels = struct {
}{}

// floorR is where relationships are stored.
type floorR struct {
}

// NewStruct creates a new relationship struct
func (*floorR) NewStruct() *floorR {
	return &floorR{}
}

// floorL is where Load methods for each relationship are stored.
type floorL struct{}

var (
	floorAllColumns            = []string{"publisher", "domain", "country", "device", "floor", "created_at", "updated_at", "rule_id", "demand_partner_id", "browser", "os", "placement_type"}
	floorColumnsWithoutDefault = []string{"publisher", "domain", "country", "device", "created_at"}
	floorColumnsWithDefault    = []string{"floor", "updated_at", "rule_id", "demand_partner_id", "browser", "os", "placement_type"}
	floorPrimaryKeyColumns     = []string{"publisher", "domain", "device", "country"}
	floorGeneratedColumns      = []string{}
)

type (
	// FloorSlice is an alias for a slice of pointers to Floor.
	// This should almost always be used instead of []Floor.
	FloorSlice []*Floor
	// FloorHook is the signature for custom Floor hook methods
	FloorHook func(context.Context, boil.ContextExecutor, *Floor) error

	floorQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	floorType                 = reflect.TypeOf(&Floor{})
	floorMapping              = queries.MakeStructMapping(floorType)
	floorPrimaryKeyMapping, _ = queries.BindMapping(floorType, floorMapping, floorPrimaryKeyColumns)
	floorInsertCacheMut       sync.RWMutex
	floorInsertCache          = make(map[string]insertCache)
	floorUpdateCacheMut       sync.RWMutex
	floorUpdateCache          = make(map[string]updateCache)
	floorUpsertCacheMut       sync.RWMutex
	floorUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var floorAfterSelectMu sync.Mutex
var floorAfterSelectHooks []FloorHook

var floorBeforeInsertMu sync.Mutex
var floorBeforeInsertHooks []FloorHook
var floorAfterInsertMu sync.Mutex
var floorAfterInsertHooks []FloorHook

var floorBeforeUpdateMu sync.Mutex
var floorBeforeUpdateHooks []FloorHook
var floorAfterUpdateMu sync.Mutex
var floorAfterUpdateHooks []FloorHook

var floorBeforeDeleteMu sync.Mutex
var floorBeforeDeleteHooks []FloorHook
var floorAfterDeleteMu sync.Mutex
var floorAfterDeleteHooks []FloorHook

var floorBeforeUpsertMu sync.Mutex
var floorBeforeUpsertHooks []FloorHook
var floorAfterUpsertMu sync.Mutex
var floorAfterUpsertHooks []FloorHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Floor) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range floorAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Floor) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range floorBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Floor) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range floorAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Floor) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range floorBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Floor) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range floorAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Floor) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range floorBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Floor) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range floorAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Floor) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range floorBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Floor) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range floorAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFloorHook registers your hook function for all future operations.
func AddFloorHook(hookPoint boil.HookPoint, floorHook FloorHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		floorAfterSelectMu.Lock()
		floorAfterSelectHooks = append(floorAfterSelectHooks, floorHook)
		floorAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		floorBeforeInsertMu.Lock()
		floorBeforeInsertHooks = append(floorBeforeInsertHooks, floorHook)
		floorBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		floorAfterInsertMu.Lock()
		floorAfterInsertHooks = append(floorAfterInsertHooks, floorHook)
		floorAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		floorBeforeUpdateMu.Lock()
		floorBeforeUpdateHooks = append(floorBeforeUpdateHooks, floorHook)
		floorBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		floorAfterUpdateMu.Lock()
		floorAfterUpdateHooks = append(floorAfterUpdateHooks, floorHook)
		floorAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		floorBeforeDeleteMu.Lock()
		floorBeforeDeleteHooks = append(floorBeforeDeleteHooks, floorHook)
		floorBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		floorAfterDeleteMu.Lock()
		floorAfterDeleteHooks = append(floorAfterDeleteHooks, floorHook)
		floorAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		floorBeforeUpsertMu.Lock()
		floorBeforeUpsertHooks = append(floorBeforeUpsertHooks, floorHook)
		floorBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		floorAfterUpsertMu.Lock()
		floorAfterUpsertHooks = append(floorAfterUpsertHooks, floorHook)
		floorAfterUpsertMu.Unlock()
	}
}

// One returns a single floor record from the query.
func (q floorQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Floor, error) {
	o := &Floor{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for floor")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Floor records from the query.
func (q floorQuery) All(ctx context.Context, exec boil.ContextExecutor) (FloorSlice, error) {
	var o []*Floor

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Floor slice")
	}

	if len(floorAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Floor records in the query.
func (q floorQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count floor rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q floorQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if floor exists")
	}

	return count > 0, nil
}

// Floors retrieves all the records using an executor.
func Floors(mods ...qm.QueryMod) floorQuery {
	mods = append(mods, qm.From("\"floor\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"floor\".*"})
	}

	return floorQuery{q}
}

// FindFloor retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFloor(ctx context.Context, exec boil.ContextExecutor, publisher string, domain string, device string, country string, selectCols ...string) (*Floor, error) {
	floorObj := &Floor{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"floor\" where \"publisher\"=$1 AND \"domain\"=$2 AND \"device\"=$3 AND \"country\"=$4", sel,
	)

	q := queries.Raw(query, publisher, domain, device, country)

	err := q.Bind(ctx, exec, floorObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from floor")
	}

	if err = floorObj.doAfterSelectHooks(ctx, exec); err != nil {
		return floorObj, err
	}

	return floorObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Floor) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no floor provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(floorColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	floorInsertCacheMut.RLock()
	cache, cached := floorInsertCache[key]
	floorInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			floorAllColumns,
			floorColumnsWithDefault,
			floorColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(floorType, floorMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(floorType, floorMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"floor\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"floor\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into floor")
	}

	if !cached {
		floorInsertCacheMut.Lock()
		floorInsertCache[key] = cache
		floorInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Floor.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Floor) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	floorUpdateCacheMut.RLock()
	cache, cached := floorUpdateCache[key]
	floorUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			floorAllColumns,
			floorPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update floor, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"floor\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, floorPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(floorType, floorMapping, append(wl, floorPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update floor row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for floor")
	}

	if !cached {
		floorUpdateCacheMut.Lock()
		floorUpdateCache[key] = cache
		floorUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q floorQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for floor")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for floor")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FloorSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), floorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"floor\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, floorPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in floor slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all floor")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Floor) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no floor provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(floorColumnsWithDefault, o)

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

	floorUpsertCacheMut.RLock()
	cache, cached := floorUpsertCache[key]
	floorUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			floorAllColumns,
			floorColumnsWithDefault,
			floorColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			floorAllColumns,
			floorPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert floor, could not build update column list")
		}

		ret := strmangle.SetComplement(floorAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(floorPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert floor, could not build conflict column list")
			}

			conflict = make([]string, len(floorPrimaryKeyColumns))
			copy(conflict, floorPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"floor\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(floorType, floorMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(floorType, floorMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert floor")
	}

	if !cached {
		floorUpsertCacheMut.Lock()
		floorUpsertCache[key] = cache
		floorUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Floor record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Floor) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Floor provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), floorPrimaryKeyMapping)
	sql := "DELETE FROM \"floor\" WHERE \"publisher\"=$1 AND \"domain\"=$2 AND \"device\"=$3 AND \"country\"=$4"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from floor")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for floor")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q floorQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no floorQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from floor")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for floor")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FloorSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(floorBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), floorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"floor\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, floorPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from floor slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for floor")
	}

	if len(floorAfterDeleteHooks) != 0 {
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
func (o *Floor) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindFloor(ctx, exec, o.Publisher, o.Domain, o.Device, o.Country)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FloorSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := FloorSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), floorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"floor\".* FROM \"floor\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, floorPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FloorSlice")
	}

	*o = slice

	return nil
}

// FloorExists checks if the Floor row exists.
func FloorExists(ctx context.Context, exec boil.ContextExecutor, publisher string, domain string, device string, country string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"floor\" where \"publisher\"=$1 AND \"domain\"=$2 AND \"device\"=$3 AND \"country\"=$4 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, publisher, domain, device, country)
	}
	row := exec.QueryRowContext(ctx, sql, publisher, domain, device, country)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if floor exists")
	}

	return exists, nil
}

// Exists checks if the Floor row exists.
func (o *Floor) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return FloorExists(ctx, exec, o.Publisher, o.Domain, o.Device, o.Country)
}