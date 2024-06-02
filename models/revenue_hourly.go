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

// RevenueHourly is an object representing the database table.
type RevenueHourly struct {
	Time                 time.Time `boil:"time" json:"time" toml:"time" yaml:"time"`
	PublisherImpressions int64     `boil:"publisher_impressions" json:"publisher_impressions" toml:"publisher_impressions" yaml:"publisher_impressions"`
	SoldImpressions      int64     `boil:"sold_impressions" json:"sold_impressions" toml:"sold_impressions" yaml:"sold_impressions"`
	Cost                 float64   `boil:"cost" json:"cost" toml:"cost" yaml:"cost"`
	Revenue              float64   `boil:"revenue" json:"revenue" toml:"revenue" yaml:"revenue"`
	DemandPartnerFees    float64   `boil:"demand_partner_fees" json:"demand_partner_fees" toml:"demand_partner_fees" yaml:"demand_partner_fees"`
	MissedOpportunities  int64     `boil:"missed_opportunities" json:"missed_opportunities" toml:"missed_opportunities" yaml:"missed_opportunities"`
	DataFee              float64   `boil:"data_fee" json:"data_fee" toml:"data_fee" yaml:"data_fee"`
	DPBidRequests        int64     `boil:"dp_bid_requests" json:"dp_bid_requests" toml:"dp_bid_requests" yaml:"dp_bid_requests"`

	R *revenueHourlyR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L revenueHourlyL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RevenueHourlyColumns = struct {
	Time                 string
	PublisherImpressions string
	SoldImpressions      string
	Cost                 string
	Revenue              string
	DemandPartnerFees    string
	MissedOpportunities  string
	DataFee              string
	DPBidRequests        string
}{
	Time:                 "time",
	PublisherImpressions: "publisher_impressions",
	SoldImpressions:      "sold_impressions",
	Cost:                 "cost",
	Revenue:              "revenue",
	DemandPartnerFees:    "demand_partner_fees",
	MissedOpportunities:  "missed_opportunities",
	DataFee:              "data_fee",
	DPBidRequests:        "dp_bid_requests",
}

var RevenueHourlyTableColumns = struct {
	Time                 string
	PublisherImpressions string
	SoldImpressions      string
	Cost                 string
	Revenue              string
	DemandPartnerFees    string
	MissedOpportunities  string
	DataFee              string
	DPBidRequests        string
}{
	Time:                 "revenue_hourly.time",
	PublisherImpressions: "revenue_hourly.publisher_impressions",
	SoldImpressions:      "revenue_hourly.sold_impressions",
	Cost:                 "revenue_hourly.cost",
	Revenue:              "revenue_hourly.revenue",
	DemandPartnerFees:    "revenue_hourly.demand_partner_fees",
	MissedOpportunities:  "revenue_hourly.missed_opportunities",
	DataFee:              "revenue_hourly.data_fee",
	DPBidRequests:        "revenue_hourly.dp_bid_requests",
}

// Generated where

var RevenueHourlyWhere = struct {
	Time                 whereHelpertime_Time
	PublisherImpressions whereHelperint64
	SoldImpressions      whereHelperint64
	Cost                 whereHelperfloat64
	Revenue              whereHelperfloat64
	DemandPartnerFees    whereHelperfloat64
	MissedOpportunities  whereHelperint64
	DataFee              whereHelperfloat64
	DPBidRequests        whereHelperint64
}{
	Time:                 whereHelpertime_Time{field: "\"revenue_hourly\".\"time\""},
	PublisherImpressions: whereHelperint64{field: "\"revenue_hourly\".\"publisher_impressions\""},
	SoldImpressions:      whereHelperint64{field: "\"revenue_hourly\".\"sold_impressions\""},
	Cost:                 whereHelperfloat64{field: "\"revenue_hourly\".\"cost\""},
	Revenue:              whereHelperfloat64{field: "\"revenue_hourly\".\"revenue\""},
	DemandPartnerFees:    whereHelperfloat64{field: "\"revenue_hourly\".\"demand_partner_fees\""},
	MissedOpportunities:  whereHelperint64{field: "\"revenue_hourly\".\"missed_opportunities\""},
	DataFee:              whereHelperfloat64{field: "\"revenue_hourly\".\"data_fee\""},
	DPBidRequests:        whereHelperint64{field: "\"revenue_hourly\".\"dp_bid_requests\""},
}

// RevenueHourlyRels is where relationship names are stored.
var RevenueHourlyRels = struct {
}{}

// revenueHourlyR is where relationships are stored.
type revenueHourlyR struct {
}

// NewStruct creates a new relationship struct
func (*revenueHourlyR) NewStruct() *revenueHourlyR {
	return &revenueHourlyR{}
}

// revenueHourlyL is where Load methods for each relationship are stored.
type revenueHourlyL struct{}

var (
	revenueHourlyAllColumns            = []string{"time", "publisher_impressions", "sold_impressions", "cost", "revenue", "demand_partner_fees", "missed_opportunities", "data_fee", "dp_bid_requests"}
	revenueHourlyColumnsWithoutDefault = []string{"time"}
	revenueHourlyColumnsWithDefault    = []string{"publisher_impressions", "sold_impressions", "cost", "revenue", "demand_partner_fees", "missed_opportunities", "data_fee", "dp_bid_requests"}
	revenueHourlyPrimaryKeyColumns     = []string{"time"}
	revenueHourlyGeneratedColumns      = []string{}
)

type (
	// RevenueHourlySlice is an alias for a slice of pointers to RevenueHourly.
	// This should almost always be used instead of []RevenueHourly.
	RevenueHourlySlice []*RevenueHourly
	// RevenueHourlyHook is the signature for custom RevenueHourly hook methods
	RevenueHourlyHook func(context.Context, boil.ContextExecutor, *RevenueHourly) error

	revenueHourlyQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	revenueHourlyType                 = reflect.TypeOf(&RevenueHourly{})
	revenueHourlyMapping              = queries.MakeStructMapping(revenueHourlyType)
	revenueHourlyPrimaryKeyMapping, _ = queries.BindMapping(revenueHourlyType, revenueHourlyMapping, revenueHourlyPrimaryKeyColumns)
	revenueHourlyInsertCacheMut       sync.RWMutex
	revenueHourlyInsertCache          = make(map[string]insertCache)
	revenueHourlyUpdateCacheMut       sync.RWMutex
	revenueHourlyUpdateCache          = make(map[string]updateCache)
	revenueHourlyUpsertCacheMut       sync.RWMutex
	revenueHourlyUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var revenueHourlyAfterSelectMu sync.Mutex
var revenueHourlyAfterSelectHooks []RevenueHourlyHook

var revenueHourlyBeforeInsertMu sync.Mutex
var revenueHourlyBeforeInsertHooks []RevenueHourlyHook
var revenueHourlyAfterInsertMu sync.Mutex
var revenueHourlyAfterInsertHooks []RevenueHourlyHook

var revenueHourlyBeforeUpdateMu sync.Mutex
var revenueHourlyBeforeUpdateHooks []RevenueHourlyHook
var revenueHourlyAfterUpdateMu sync.Mutex
var revenueHourlyAfterUpdateHooks []RevenueHourlyHook

var revenueHourlyBeforeDeleteMu sync.Mutex
var revenueHourlyBeforeDeleteHooks []RevenueHourlyHook
var revenueHourlyAfterDeleteMu sync.Mutex
var revenueHourlyAfterDeleteHooks []RevenueHourlyHook

var revenueHourlyBeforeUpsertMu sync.Mutex
var revenueHourlyBeforeUpsertHooks []RevenueHourlyHook
var revenueHourlyAfterUpsertMu sync.Mutex
var revenueHourlyAfterUpsertHooks []RevenueHourlyHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *RevenueHourly) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range revenueHourlyAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *RevenueHourly) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range revenueHourlyBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *RevenueHourly) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range revenueHourlyAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *RevenueHourly) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range revenueHourlyBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *RevenueHourly) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range revenueHourlyAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *RevenueHourly) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range revenueHourlyBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *RevenueHourly) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range revenueHourlyAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *RevenueHourly) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range revenueHourlyBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *RevenueHourly) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range revenueHourlyAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRevenueHourlyHook registers your hook function for all future operations.
func AddRevenueHourlyHook(hookPoint boil.HookPoint, revenueHourlyHook RevenueHourlyHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		revenueHourlyAfterSelectMu.Lock()
		revenueHourlyAfterSelectHooks = append(revenueHourlyAfterSelectHooks, revenueHourlyHook)
		revenueHourlyAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		revenueHourlyBeforeInsertMu.Lock()
		revenueHourlyBeforeInsertHooks = append(revenueHourlyBeforeInsertHooks, revenueHourlyHook)
		revenueHourlyBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		revenueHourlyAfterInsertMu.Lock()
		revenueHourlyAfterInsertHooks = append(revenueHourlyAfterInsertHooks, revenueHourlyHook)
		revenueHourlyAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		revenueHourlyBeforeUpdateMu.Lock()
		revenueHourlyBeforeUpdateHooks = append(revenueHourlyBeforeUpdateHooks, revenueHourlyHook)
		revenueHourlyBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		revenueHourlyAfterUpdateMu.Lock()
		revenueHourlyAfterUpdateHooks = append(revenueHourlyAfterUpdateHooks, revenueHourlyHook)
		revenueHourlyAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		revenueHourlyBeforeDeleteMu.Lock()
		revenueHourlyBeforeDeleteHooks = append(revenueHourlyBeforeDeleteHooks, revenueHourlyHook)
		revenueHourlyBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		revenueHourlyAfterDeleteMu.Lock()
		revenueHourlyAfterDeleteHooks = append(revenueHourlyAfterDeleteHooks, revenueHourlyHook)
		revenueHourlyAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		revenueHourlyBeforeUpsertMu.Lock()
		revenueHourlyBeforeUpsertHooks = append(revenueHourlyBeforeUpsertHooks, revenueHourlyHook)
		revenueHourlyBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		revenueHourlyAfterUpsertMu.Lock()
		revenueHourlyAfterUpsertHooks = append(revenueHourlyAfterUpsertHooks, revenueHourlyHook)
		revenueHourlyAfterUpsertMu.Unlock()
	}
}

// One returns a single revenueHourly record from the query.
func (q revenueHourlyQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RevenueHourly, error) {
	o := &RevenueHourly{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for revenue_hourly")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all RevenueHourly records from the query.
func (q revenueHourlyQuery) All(ctx context.Context, exec boil.ContextExecutor) (RevenueHourlySlice, error) {
	var o []*RevenueHourly

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RevenueHourly slice")
	}

	if len(revenueHourlyAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all RevenueHourly records in the query.
func (q revenueHourlyQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count revenue_hourly rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q revenueHourlyQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if revenue_hourly exists")
	}

	return count > 0, nil
}

// RevenueHourlies retrieves all the records using an executor.
func RevenueHourlies(mods ...qm.QueryMod) revenueHourlyQuery {
	mods = append(mods, qm.From("\"revenue_hourly\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"revenue_hourly\".*"})
	}

	return revenueHourlyQuery{q}
}

// FindRevenueHourly retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRevenueHourly(ctx context.Context, exec boil.ContextExecutor, time time.Time, selectCols ...string) (*RevenueHourly, error) {
	revenueHourlyObj := &RevenueHourly{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"revenue_hourly\" where \"time\"=$1", sel,
	)

	q := queries.Raw(query, time)

	err := q.Bind(ctx, exec, revenueHourlyObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from revenue_hourly")
	}

	if err = revenueHourlyObj.doAfterSelectHooks(ctx, exec); err != nil {
		return revenueHourlyObj, err
	}

	return revenueHourlyObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RevenueHourly) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no revenue_hourly provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(revenueHourlyColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	revenueHourlyInsertCacheMut.RLock()
	cache, cached := revenueHourlyInsertCache[key]
	revenueHourlyInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			revenueHourlyAllColumns,
			revenueHourlyColumnsWithDefault,
			revenueHourlyColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(revenueHourlyType, revenueHourlyMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(revenueHourlyType, revenueHourlyMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"revenue_hourly\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"revenue_hourly\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into revenue_hourly")
	}

	if !cached {
		revenueHourlyInsertCacheMut.Lock()
		revenueHourlyInsertCache[key] = cache
		revenueHourlyInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the RevenueHourly.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RevenueHourly) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	revenueHourlyUpdateCacheMut.RLock()
	cache, cached := revenueHourlyUpdateCache[key]
	revenueHourlyUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			revenueHourlyAllColumns,
			revenueHourlyPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update revenue_hourly, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"revenue_hourly\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, revenueHourlyPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(revenueHourlyType, revenueHourlyMapping, append(wl, revenueHourlyPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update revenue_hourly row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for revenue_hourly")
	}

	if !cached {
		revenueHourlyUpdateCacheMut.Lock()
		revenueHourlyUpdateCache[key] = cache
		revenueHourlyUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q revenueHourlyQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for revenue_hourly")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for revenue_hourly")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RevenueHourlySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), revenueHourlyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"revenue_hourly\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, revenueHourlyPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in revenueHourly slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all revenueHourly")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RevenueHourly) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no revenue_hourly provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(revenueHourlyColumnsWithDefault, o)

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

	revenueHourlyUpsertCacheMut.RLock()
	cache, cached := revenueHourlyUpsertCache[key]
	revenueHourlyUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			revenueHourlyAllColumns,
			revenueHourlyColumnsWithDefault,
			revenueHourlyColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			revenueHourlyAllColumns,
			revenueHourlyPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert revenue_hourly, could not build update column list")
		}

		ret := strmangle.SetComplement(revenueHourlyAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(revenueHourlyPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert revenue_hourly, could not build conflict column list")
			}

			conflict = make([]string, len(revenueHourlyPrimaryKeyColumns))
			copy(conflict, revenueHourlyPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"revenue_hourly\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(revenueHourlyType, revenueHourlyMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(revenueHourlyType, revenueHourlyMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert revenue_hourly")
	}

	if !cached {
		revenueHourlyUpsertCacheMut.Lock()
		revenueHourlyUpsertCache[key] = cache
		revenueHourlyUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single RevenueHourly record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RevenueHourly) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RevenueHourly provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), revenueHourlyPrimaryKeyMapping)
	sql := "DELETE FROM \"revenue_hourly\" WHERE \"time\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from revenue_hourly")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for revenue_hourly")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q revenueHourlyQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no revenueHourlyQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from revenue_hourly")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for revenue_hourly")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RevenueHourlySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(revenueHourlyBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), revenueHourlyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"revenue_hourly\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, revenueHourlyPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from revenueHourly slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for revenue_hourly")
	}

	if len(revenueHourlyAfterDeleteHooks) != 0 {
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
func (o *RevenueHourly) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRevenueHourly(ctx, exec, o.Time)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RevenueHourlySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RevenueHourlySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), revenueHourlyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"revenue_hourly\".* FROM \"revenue_hourly\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, revenueHourlyPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RevenueHourlySlice")
	}

	*o = slice

	return nil
}

// RevenueHourlyExists checks if the RevenueHourly row exists.
func RevenueHourlyExists(ctx context.Context, exec boil.ContextExecutor, time time.Time) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"revenue_hourly\" where \"time\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, time)
	}
	row := exec.QueryRowContext(ctx, sql, time)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if revenue_hourly exists")
	}

	return exists, nil
}

// Exists checks if the RevenueHourly row exists.
func (o *RevenueHourly) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return RevenueHourlyExists(ctx, exec, o.Time)
}
