// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testPublisherHourlies(t *testing.T) {
	t.Parallel()

	query := PublisherHourlies()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPublisherHourliesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PublisherHourly{}
	if err = randomize.Struct(seed, o, publisherHourlyDBTypes, true, publisherHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PublisherHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPublisherHourliesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PublisherHourly{}
	if err = randomize.Struct(seed, o, publisherHourlyDBTypes, true, publisherHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := PublisherHourlies().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PublisherHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPublisherHourliesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PublisherHourly{}
	if err = randomize.Struct(seed, o, publisherHourlyDBTypes, true, publisherHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PublisherHourlySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PublisherHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPublisherHourliesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PublisherHourly{}
	if err = randomize.Struct(seed, o, publisherHourlyDBTypes, true, publisherHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PublisherHourlyExists(ctx, tx, o.Time, o.PublisherID, o.Domain, o.Os, o.Country, o.DeviceType)
	if err != nil {
		t.Errorf("Unable to check if PublisherHourly exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PublisherHourlyExists to return true, but got false.")
	}
}

func testPublisherHourliesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PublisherHourly{}
	if err = randomize.Struct(seed, o, publisherHourlyDBTypes, true, publisherHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	publisherHourlyFound, err := FindPublisherHourly(ctx, tx, o.Time, o.PublisherID, o.Domain, o.Os, o.Country, o.DeviceType)
	if err != nil {
		t.Error(err)
	}

	if publisherHourlyFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPublisherHourliesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PublisherHourly{}
	if err = randomize.Struct(seed, o, publisherHourlyDBTypes, true, publisherHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = PublisherHourlies().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPublisherHourliesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PublisherHourly{}
	if err = randomize.Struct(seed, o, publisherHourlyDBTypes, true, publisherHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := PublisherHourlies().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPublisherHourliesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	publisherHourlyOne := &PublisherHourly{}
	publisherHourlyTwo := &PublisherHourly{}
	if err = randomize.Struct(seed, publisherHourlyOne, publisherHourlyDBTypes, false, publisherHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherHourly struct: %s", err)
	}
	if err = randomize.Struct(seed, publisherHourlyTwo, publisherHourlyDBTypes, false, publisherHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = publisherHourlyOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = publisherHourlyTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PublisherHourlies().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPublisherHourliesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	publisherHourlyOne := &PublisherHourly{}
	publisherHourlyTwo := &PublisherHourly{}
	if err = randomize.Struct(seed, publisherHourlyOne, publisherHourlyDBTypes, false, publisherHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherHourly struct: %s", err)
	}
	if err = randomize.Struct(seed, publisherHourlyTwo, publisherHourlyDBTypes, false, publisherHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = publisherHourlyOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = publisherHourlyTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PublisherHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func publisherHourlyBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *PublisherHourly) error {
	*o = PublisherHourly{}
	return nil
}

func publisherHourlyAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *PublisherHourly) error {
	*o = PublisherHourly{}
	return nil
}

func publisherHourlyAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *PublisherHourly) error {
	*o = PublisherHourly{}
	return nil
}

func publisherHourlyBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *PublisherHourly) error {
	*o = PublisherHourly{}
	return nil
}

func publisherHourlyAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *PublisherHourly) error {
	*o = PublisherHourly{}
	return nil
}

func publisherHourlyBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *PublisherHourly) error {
	*o = PublisherHourly{}
	return nil
}

func publisherHourlyAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *PublisherHourly) error {
	*o = PublisherHourly{}
	return nil
}

func publisherHourlyBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *PublisherHourly) error {
	*o = PublisherHourly{}
	return nil
}

func publisherHourlyAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *PublisherHourly) error {
	*o = PublisherHourly{}
	return nil
}

func testPublisherHourliesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &PublisherHourly{}
	o := &PublisherHourly{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, publisherHourlyDBTypes, false); err != nil {
		t.Errorf("Unable to randomize PublisherHourly object: %s", err)
	}

	AddPublisherHourlyHook(boil.BeforeInsertHook, publisherHourlyBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	publisherHourlyBeforeInsertHooks = []PublisherHourlyHook{}

	AddPublisherHourlyHook(boil.AfterInsertHook, publisherHourlyAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	publisherHourlyAfterInsertHooks = []PublisherHourlyHook{}

	AddPublisherHourlyHook(boil.AfterSelectHook, publisherHourlyAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	publisherHourlyAfterSelectHooks = []PublisherHourlyHook{}

	AddPublisherHourlyHook(boil.BeforeUpdateHook, publisherHourlyBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	publisherHourlyBeforeUpdateHooks = []PublisherHourlyHook{}

	AddPublisherHourlyHook(boil.AfterUpdateHook, publisherHourlyAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	publisherHourlyAfterUpdateHooks = []PublisherHourlyHook{}

	AddPublisherHourlyHook(boil.BeforeDeleteHook, publisherHourlyBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	publisherHourlyBeforeDeleteHooks = []PublisherHourlyHook{}

	AddPublisherHourlyHook(boil.AfterDeleteHook, publisherHourlyAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	publisherHourlyAfterDeleteHooks = []PublisherHourlyHook{}

	AddPublisherHourlyHook(boil.BeforeUpsertHook, publisherHourlyBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	publisherHourlyBeforeUpsertHooks = []PublisherHourlyHook{}

	AddPublisherHourlyHook(boil.AfterUpsertHook, publisherHourlyAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	publisherHourlyAfterUpsertHooks = []PublisherHourlyHook{}
}

func testPublisherHourliesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PublisherHourly{}
	if err = randomize.Struct(seed, o, publisherHourlyDBTypes, true, publisherHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PublisherHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPublisherHourliesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PublisherHourly{}
	if err = randomize.Struct(seed, o, publisherHourlyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PublisherHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(publisherHourlyColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := PublisherHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPublisherHourliesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PublisherHourly{}
	if err = randomize.Struct(seed, o, publisherHourlyDBTypes, true, publisherHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPublisherHourliesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PublisherHourly{}
	if err = randomize.Struct(seed, o, publisherHourlyDBTypes, true, publisherHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PublisherHourlySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPublisherHourliesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PublisherHourly{}
	if err = randomize.Struct(seed, o, publisherHourlyDBTypes, true, publisherHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PublisherHourlies().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	publisherHourlyDBTypes = map[string]string{`Time`: `timestamp without time zone`, `PublisherID`: `character varying`, `Domain`: `character varying`, `Os`: `character varying`, `Country`: `character varying`, `DeviceType`: `character varying`, `BidRequests`: `bigint`, `BidResponses`: `bigint`, `BidPriceCount`: `bigint`, `BidPriceTotal`: `double precision`, `PublisherImpressions`: `bigint`, `DemandImpressions`: `bigint`, `MissedOpportunities`: `bigint`, `SupplyTotal`: `double precision`, `DemandTotal`: `double precision`, `DemandPartnerFee`: `double precision`}
	_                      = bytes.MinRead
)

func testPublisherHourliesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(publisherHourlyPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(publisherHourlyAllColumns) == len(publisherHourlyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PublisherHourly{}
	if err = randomize.Struct(seed, o, publisherHourlyDBTypes, true, publisherHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PublisherHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, publisherHourlyDBTypes, true, publisherHourlyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PublisherHourly struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPublisherHourliesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(publisherHourlyAllColumns) == len(publisherHourlyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PublisherHourly{}
	if err = randomize.Struct(seed, o, publisherHourlyDBTypes, true, publisherHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PublisherHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PublisherHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, publisherHourlyDBTypes, true, publisherHourlyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PublisherHourly struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(publisherHourlyAllColumns, publisherHourlyPrimaryKeyColumns) {
		fields = publisherHourlyAllColumns
	} else {
		fields = strmangle.SetComplement(
			publisherHourlyAllColumns,
			publisherHourlyPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := PublisherHourlySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPublisherHourliesUpsert(t *testing.T) {
	t.Parallel()

	if len(publisherHourlyAllColumns) == len(publisherHourlyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := PublisherHourly{}
	if err = randomize.Struct(seed, &o, publisherHourlyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PublisherHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PublisherHourly: %s", err)
	}

	count, err := PublisherHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, publisherHourlyDBTypes, false, publisherHourlyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PublisherHourly struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PublisherHourly: %s", err)
	}

	count, err = PublisherHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
