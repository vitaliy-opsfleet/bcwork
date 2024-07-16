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

// PublisherDomain is an object representing the database table.
type PublisherDomain struct {
	Name        string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	PublisherID string    `boil:"publisher_id" json:"publisher_id" toml:"publisher_id" yaml:"publisher_id"`
	CreatedAt   time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *publisherDomainR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L publisherDomainL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PublisherDomainColumns = struct {
	Name        string
	PublisherID string
	CreatedAt   string
}{
	Name:        "name",
	PublisherID: "publisher_id",
	CreatedAt:   "created_at",
}

var PublisherDomainTableColumns = struct {
	Name        string
	PublisherID string
	CreatedAt   string
}{
	Name:        "publisher_domain.name",
	PublisherID: "publisher_domain.publisher_id",
	CreatedAt:   "publisher_domain.created_at",
}

// Generated where

var PublisherDomainWhere = struct {
	Name        whereHelperstring
	PublisherID whereHelperstring
	CreatedAt   whereHelpertime_Time
}{
	Name:        whereHelperstring{field: "\"publisher_domain\".\"name\""},
	PublisherID: whereHelperstring{field: "\"publisher_domain\".\"publisher_id\""},
	CreatedAt:   whereHelpertime_Time{field: "\"publisher_domain\".\"created_at\""},
}

// PublisherDomainRels is where relationship names are stored.
var PublisherDomainRels = struct {
	Publisher string
}{
	Publisher: "Publisher",
}

// publisherDomainR is where relationships are stored.
type publisherDomainR struct {
	Publisher *Publisher `boil:"Publisher" json:"Publisher" toml:"Publisher" yaml:"Publisher"`
}

// NewStruct creates a new relationship struct
func (*publisherDomainR) NewStruct() *publisherDomainR {
	return &publisherDomainR{}
}

func (r *publisherDomainR) GetPublisher() *Publisher {
	if r == nil {
		return nil
	}
	return r.Publisher
}

// publisherDomainL is where Load methods for each relationship are stored.
type publisherDomainL struct{}

var (
	publisherDomainAllColumns            = []string{"name", "publisher_id", "created_at"}
	publisherDomainColumnsWithoutDefault = []string{"name", "publisher_id", "created_at"}
	publisherDomainColumnsWithDefault    = []string{}
	publisherDomainPrimaryKeyColumns     = []string{"name", "publisher_id"}
	publisherDomainGeneratedColumns      = []string{}
)

type (
	// PublisherDomainSlice is an alias for a slice of pointers to PublisherDomain.
	// This should almost always be used instead of []PublisherDomain.
	PublisherDomainSlice []*PublisherDomain
	// PublisherDomainHook is the signature for custom PublisherDomain hook methods
	PublisherDomainHook func(context.Context, boil.ContextExecutor, *PublisherDomain) error

	publisherDomainQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	publisherDomainType                 = reflect.TypeOf(&PublisherDomain{})
	publisherDomainMapping              = queries.MakeStructMapping(publisherDomainType)
	publisherDomainPrimaryKeyMapping, _ = queries.BindMapping(publisherDomainType, publisherDomainMapping, publisherDomainPrimaryKeyColumns)
	publisherDomainInsertCacheMut       sync.RWMutex
	publisherDomainInsertCache          = make(map[string]insertCache)
	publisherDomainUpdateCacheMut       sync.RWMutex
	publisherDomainUpdateCache          = make(map[string]updateCache)
	publisherDomainUpsertCacheMut       sync.RWMutex
	publisherDomainUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var publisherDomainAfterSelectMu sync.Mutex
var publisherDomainAfterSelectHooks []PublisherDomainHook

var publisherDomainBeforeInsertMu sync.Mutex
var publisherDomainBeforeInsertHooks []PublisherDomainHook
var publisherDomainAfterInsertMu sync.Mutex
var publisherDomainAfterInsertHooks []PublisherDomainHook

var publisherDomainBeforeUpdateMu sync.Mutex
var publisherDomainBeforeUpdateHooks []PublisherDomainHook
var publisherDomainAfterUpdateMu sync.Mutex
var publisherDomainAfterUpdateHooks []PublisherDomainHook

var publisherDomainBeforeDeleteMu sync.Mutex
var publisherDomainBeforeDeleteHooks []PublisherDomainHook
var publisherDomainAfterDeleteMu sync.Mutex
var publisherDomainAfterDeleteHooks []PublisherDomainHook

var publisherDomainBeforeUpsertMu sync.Mutex
var publisherDomainBeforeUpsertHooks []PublisherDomainHook
var publisherDomainAfterUpsertMu sync.Mutex
var publisherDomainAfterUpsertHooks []PublisherDomainHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *PublisherDomain) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range publisherDomainAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *PublisherDomain) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range publisherDomainBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *PublisherDomain) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range publisherDomainAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *PublisherDomain) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range publisherDomainBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *PublisherDomain) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range publisherDomainAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *PublisherDomain) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range publisherDomainBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *PublisherDomain) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range publisherDomainAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *PublisherDomain) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range publisherDomainBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *PublisherDomain) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range publisherDomainAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPublisherDomainHook registers your hook function for all future operations.
func AddPublisherDomainHook(hookPoint boil.HookPoint, publisherDomainHook PublisherDomainHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		publisherDomainAfterSelectMu.Lock()
		publisherDomainAfterSelectHooks = append(publisherDomainAfterSelectHooks, publisherDomainHook)
		publisherDomainAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		publisherDomainBeforeInsertMu.Lock()
		publisherDomainBeforeInsertHooks = append(publisherDomainBeforeInsertHooks, publisherDomainHook)
		publisherDomainBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		publisherDomainAfterInsertMu.Lock()
		publisherDomainAfterInsertHooks = append(publisherDomainAfterInsertHooks, publisherDomainHook)
		publisherDomainAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		publisherDomainBeforeUpdateMu.Lock()
		publisherDomainBeforeUpdateHooks = append(publisherDomainBeforeUpdateHooks, publisherDomainHook)
		publisherDomainBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		publisherDomainAfterUpdateMu.Lock()
		publisherDomainAfterUpdateHooks = append(publisherDomainAfterUpdateHooks, publisherDomainHook)
		publisherDomainAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		publisherDomainBeforeDeleteMu.Lock()
		publisherDomainBeforeDeleteHooks = append(publisherDomainBeforeDeleteHooks, publisherDomainHook)
		publisherDomainBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		publisherDomainAfterDeleteMu.Lock()
		publisherDomainAfterDeleteHooks = append(publisherDomainAfterDeleteHooks, publisherDomainHook)
		publisherDomainAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		publisherDomainBeforeUpsertMu.Lock()
		publisherDomainBeforeUpsertHooks = append(publisherDomainBeforeUpsertHooks, publisherDomainHook)
		publisherDomainBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		publisherDomainAfterUpsertMu.Lock()
		publisherDomainAfterUpsertHooks = append(publisherDomainAfterUpsertHooks, publisherDomainHook)
		publisherDomainAfterUpsertMu.Unlock()
	}
}

// One returns a single publisherDomain record from the query.
func (q publisherDomainQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PublisherDomain, error) {
	o := &PublisherDomain{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for publisher_domain")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all PublisherDomain records from the query.
func (q publisherDomainQuery) All(ctx context.Context, exec boil.ContextExecutor) (PublisherDomainSlice, error) {
	var o []*PublisherDomain

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to PublisherDomain slice")
	}

	if len(publisherDomainAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all PublisherDomain records in the query.
func (q publisherDomainQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count publisher_domain rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q publisherDomainQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if publisher_domain exists")
	}

	return count > 0, nil
}

// Publisher pointed to by the foreign key.
func (o *PublisherDomain) Publisher(mods ...qm.QueryMod) publisherQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"publisher_id\" = ?", o.PublisherID),
	}

	queryMods = append(queryMods, mods...)

	return Publishers(queryMods...)
}

// LoadPublisher allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (publisherDomainL) LoadPublisher(ctx context.Context, e boil.ContextExecutor, singular bool, maybePublisherDomain interface{}, mods queries.Applicator) error {
	var slice []*PublisherDomain
	var object *PublisherDomain

	if singular {
		var ok bool
		object, ok = maybePublisherDomain.(*PublisherDomain)
		if !ok {
			object = new(PublisherDomain)
			ok = queries.SetFromEmbeddedStruct(&object, &maybePublisherDomain)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybePublisherDomain))
			}
		}
	} else {
		s, ok := maybePublisherDomain.(*[]*PublisherDomain)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybePublisherDomain)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybePublisherDomain))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &publisherDomainR{}
		}
		args[object.PublisherID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &publisherDomainR{}
			}

			args[obj.PublisherID] = struct{}{}

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
		qm.From(`publisher`),
		qm.WhereIn(`publisher.publisher_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Publisher")
	}

	var resultSlice []*Publisher
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Publisher")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for publisher")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for publisher")
	}

	if len(publisherAfterSelectHooks) != 0 {
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
		object.R.Publisher = foreign
		if foreign.R == nil {
			foreign.R = &publisherR{}
		}
		foreign.R.PublisherDomains = append(foreign.R.PublisherDomains, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.PublisherID == foreign.PublisherID {
				local.R.Publisher = foreign
				if foreign.R == nil {
					foreign.R = &publisherR{}
				}
				foreign.R.PublisherDomains = append(foreign.R.PublisherDomains, local)
				break
			}
		}
	}

	return nil
}

// SetPublisher of the publisherDomain to the related item.
// Sets o.R.Publisher to related.
// Adds o to related.R.PublisherDomains.
func (o *PublisherDomain) SetPublisher(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Publisher) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"publisher_domain\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"publisher_id"}),
		strmangle.WhereClause("\"", "\"", 2, publisherDomainPrimaryKeyColumns),
	)
	values := []interface{}{related.PublisherID, o.Name, o.PublisherID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PublisherID = related.PublisherID
	if o.R == nil {
		o.R = &publisherDomainR{
			Publisher: related,
		}
	} else {
		o.R.Publisher = related
	}

	if related.R == nil {
		related.R = &publisherR{
			PublisherDomains: PublisherDomainSlice{o},
		}
	} else {
		related.R.PublisherDomains = append(related.R.PublisherDomains, o)
	}

	return nil
}

// PublisherDomains retrieves all the records using an executor.
func PublisherDomains(mods ...qm.QueryMod) publisherDomainQuery {
	mods = append(mods, qm.From("\"publisher_domain\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"publisher_domain\".*"})
	}

	return publisherDomainQuery{q}
}

// FindPublisherDomain retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPublisherDomain(ctx context.Context, exec boil.ContextExecutor, name string, publisherID string, selectCols ...string) (*PublisherDomain, error) {
	publisherDomainObj := &PublisherDomain{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"publisher_domain\" where \"name\"=$1 AND \"publisher_id\"=$2", sel,
	)

	q := queries.Raw(query, name, publisherID)

	err := q.Bind(ctx, exec, publisherDomainObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from publisher_domain")
	}

	if err = publisherDomainObj.doAfterSelectHooks(ctx, exec); err != nil {
		return publisherDomainObj, err
	}

	return publisherDomainObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PublisherDomain) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no publisher_domain provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(publisherDomainColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	publisherDomainInsertCacheMut.RLock()
	cache, cached := publisherDomainInsertCache[key]
	publisherDomainInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			publisherDomainAllColumns,
			publisherDomainColumnsWithDefault,
			publisherDomainColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(publisherDomainType, publisherDomainMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(publisherDomainType, publisherDomainMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"publisher_domain\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"publisher_domain\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into publisher_domain")
	}

	if !cached {
		publisherDomainInsertCacheMut.Lock()
		publisherDomainInsertCache[key] = cache
		publisherDomainInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the PublisherDomain.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PublisherDomain) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	publisherDomainUpdateCacheMut.RLock()
	cache, cached := publisherDomainUpdateCache[key]
	publisherDomainUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			publisherDomainAllColumns,
			publisherDomainPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update publisher_domain, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"publisher_domain\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, publisherDomainPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(publisherDomainType, publisherDomainMapping, append(wl, publisherDomainPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update publisher_domain row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for publisher_domain")
	}

	if !cached {
		publisherDomainUpdateCacheMut.Lock()
		publisherDomainUpdateCache[key] = cache
		publisherDomainUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q publisherDomainQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for publisher_domain")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for publisher_domain")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PublisherDomainSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), publisherDomainPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"publisher_domain\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, publisherDomainPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in publisherDomain slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all publisherDomain")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PublisherDomain) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no publisher_domain provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(publisherDomainColumnsWithDefault, o)

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

	publisherDomainUpsertCacheMut.RLock()
	cache, cached := publisherDomainUpsertCache[key]
	publisherDomainUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			publisherDomainAllColumns,
			publisherDomainColumnsWithDefault,
			publisherDomainColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			publisherDomainAllColumns,
			publisherDomainPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert publisher_domain, could not build update column list")
		}

		ret := strmangle.SetComplement(publisherDomainAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(publisherDomainPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert publisher_domain, could not build conflict column list")
			}

			conflict = make([]string, len(publisherDomainPrimaryKeyColumns))
			copy(conflict, publisherDomainPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"publisher_domain\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(publisherDomainType, publisherDomainMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(publisherDomainType, publisherDomainMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert publisher_domain")
	}

	if !cached {
		publisherDomainUpsertCacheMut.Lock()
		publisherDomainUpsertCache[key] = cache
		publisherDomainUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single PublisherDomain record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PublisherDomain) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no PublisherDomain provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), publisherDomainPrimaryKeyMapping)
	sql := "DELETE FROM \"publisher_domain\" WHERE \"name\"=$1 AND \"publisher_id\"=$2"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from publisher_domain")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for publisher_domain")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q publisherDomainQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no publisherDomainQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from publisher_domain")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for publisher_domain")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PublisherDomainSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(publisherDomainBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), publisherDomainPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"publisher_domain\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, publisherDomainPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from publisherDomain slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for publisher_domain")
	}

	if len(publisherDomainAfterDeleteHooks) != 0 {
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
func (o *PublisherDomain) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPublisherDomain(ctx, exec, o.Name, o.PublisherID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PublisherDomainSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PublisherDomainSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), publisherDomainPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"publisher_domain\".* FROM \"publisher_domain\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, publisherDomainPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PublisherDomainSlice")
	}

	*o = slice

	return nil
}

// PublisherDomainExists checks if the PublisherDomain row exists.
func PublisherDomainExists(ctx context.Context, exec boil.ContextExecutor, name string, publisherID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"publisher_domain\" where \"name\"=$1 AND \"publisher_id\"=$2 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, name, publisherID)
	}
	row := exec.QueryRowContext(ctx, sql, name, publisherID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if publisher_domain exists")
	}

	return exists, nil
}

// Exists checks if the PublisherDomain row exists.
func (o *PublisherDomain) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return PublisherDomainExists(ctx, exec, o.Name, o.PublisherID)
}
