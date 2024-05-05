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

func testNBDemandHourlies(t *testing.T) {
	t.Parallel()

	query := NBDemandHourlies()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testNBDemandHourliesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NBDemandHourly{}
	if err = randomize.Struct(seed, o, nbDemandHourlyDBTypes, true, nbDemandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NBDemandHourly struct: %s", err)
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

	count, err := NBDemandHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNBDemandHourliesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NBDemandHourly{}
	if err = randomize.Struct(seed, o, nbDemandHourlyDBTypes, true, nbDemandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NBDemandHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := NBDemandHourlies().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := NBDemandHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNBDemandHourliesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NBDemandHourly{}
	if err = randomize.Struct(seed, o, nbDemandHourlyDBTypes, true, nbDemandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NBDemandHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NBDemandHourlySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := NBDemandHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNBDemandHourliesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NBDemandHourly{}
	if err = randomize.Struct(seed, o, nbDemandHourlyDBTypes, true, nbDemandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NBDemandHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := NBDemandHourlyExists(ctx, tx, o.Time, o.DemandPartnerID, o.DemandPartnerPlacementID, o.PublisherID, o.Domain, o.Os, o.Country, o.DeviceType, o.PlacementType, o.Size, o.RequestType, o.PaymentType)
	if err != nil {
		t.Errorf("Unable to check if NBDemandHourly exists: %s", err)
	}
	if !e {
		t.Errorf("Expected NBDemandHourlyExists to return true, but got false.")
	}
}

func testNBDemandHourliesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NBDemandHourly{}
	if err = randomize.Struct(seed, o, nbDemandHourlyDBTypes, true, nbDemandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NBDemandHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	nbDemandHourlyFound, err := FindNBDemandHourly(ctx, tx, o.Time, o.DemandPartnerID, o.DemandPartnerPlacementID, o.PublisherID, o.Domain, o.Os, o.Country, o.DeviceType, o.PlacementType, o.Size, o.RequestType, o.PaymentType)
	if err != nil {
		t.Error(err)
	}

	if nbDemandHourlyFound == nil {
		t.Error("want a record, got nil")
	}
}

func testNBDemandHourliesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NBDemandHourly{}
	if err = randomize.Struct(seed, o, nbDemandHourlyDBTypes, true, nbDemandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NBDemandHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = NBDemandHourlies().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testNBDemandHourliesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NBDemandHourly{}
	if err = randomize.Struct(seed, o, nbDemandHourlyDBTypes, true, nbDemandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NBDemandHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := NBDemandHourlies().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testNBDemandHourliesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	nbDemandHourlyOne := &NBDemandHourly{}
	nbDemandHourlyTwo := &NBDemandHourly{}
	if err = randomize.Struct(seed, nbDemandHourlyOne, nbDemandHourlyDBTypes, false, nbDemandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NBDemandHourly struct: %s", err)
	}
	if err = randomize.Struct(seed, nbDemandHourlyTwo, nbDemandHourlyDBTypes, false, nbDemandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NBDemandHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = nbDemandHourlyOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = nbDemandHourlyTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := NBDemandHourlies().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testNBDemandHourliesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	nbDemandHourlyOne := &NBDemandHourly{}
	nbDemandHourlyTwo := &NBDemandHourly{}
	if err = randomize.Struct(seed, nbDemandHourlyOne, nbDemandHourlyDBTypes, false, nbDemandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NBDemandHourly struct: %s", err)
	}
	if err = randomize.Struct(seed, nbDemandHourlyTwo, nbDemandHourlyDBTypes, false, nbDemandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NBDemandHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = nbDemandHourlyOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = nbDemandHourlyTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NBDemandHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func nbDemandHourlyBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *NBDemandHourly) error {
	*o = NBDemandHourly{}
	return nil
}

func nbDemandHourlyAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *NBDemandHourly) error {
	*o = NBDemandHourly{}
	return nil
}

func nbDemandHourlyAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *NBDemandHourly) error {
	*o = NBDemandHourly{}
	return nil
}

func nbDemandHourlyBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *NBDemandHourly) error {
	*o = NBDemandHourly{}
	return nil
}

func nbDemandHourlyAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *NBDemandHourly) error {
	*o = NBDemandHourly{}
	return nil
}

func nbDemandHourlyBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *NBDemandHourly) error {
	*o = NBDemandHourly{}
	return nil
}

func nbDemandHourlyAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *NBDemandHourly) error {
	*o = NBDemandHourly{}
	return nil
}

func nbDemandHourlyBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *NBDemandHourly) error {
	*o = NBDemandHourly{}
	return nil
}

func nbDemandHourlyAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *NBDemandHourly) error {
	*o = NBDemandHourly{}
	return nil
}

func testNBDemandHourliesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &NBDemandHourly{}
	o := &NBDemandHourly{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, nbDemandHourlyDBTypes, false); err != nil {
		t.Errorf("Unable to randomize NBDemandHourly object: %s", err)
	}

	AddNBDemandHourlyHook(boil.BeforeInsertHook, nbDemandHourlyBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	nbDemandHourlyBeforeInsertHooks = []NBDemandHourlyHook{}

	AddNBDemandHourlyHook(boil.AfterInsertHook, nbDemandHourlyAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	nbDemandHourlyAfterInsertHooks = []NBDemandHourlyHook{}

	AddNBDemandHourlyHook(boil.AfterSelectHook, nbDemandHourlyAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	nbDemandHourlyAfterSelectHooks = []NBDemandHourlyHook{}

	AddNBDemandHourlyHook(boil.BeforeUpdateHook, nbDemandHourlyBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	nbDemandHourlyBeforeUpdateHooks = []NBDemandHourlyHook{}

	AddNBDemandHourlyHook(boil.AfterUpdateHook, nbDemandHourlyAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	nbDemandHourlyAfterUpdateHooks = []NBDemandHourlyHook{}

	AddNBDemandHourlyHook(boil.BeforeDeleteHook, nbDemandHourlyBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	nbDemandHourlyBeforeDeleteHooks = []NBDemandHourlyHook{}

	AddNBDemandHourlyHook(boil.AfterDeleteHook, nbDemandHourlyAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	nbDemandHourlyAfterDeleteHooks = []NBDemandHourlyHook{}

	AddNBDemandHourlyHook(boil.BeforeUpsertHook, nbDemandHourlyBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	nbDemandHourlyBeforeUpsertHooks = []NBDemandHourlyHook{}

	AddNBDemandHourlyHook(boil.AfterUpsertHook, nbDemandHourlyAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	nbDemandHourlyAfterUpsertHooks = []NBDemandHourlyHook{}
}

func testNBDemandHourliesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NBDemandHourly{}
	if err = randomize.Struct(seed, o, nbDemandHourlyDBTypes, true, nbDemandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NBDemandHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NBDemandHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNBDemandHourliesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NBDemandHourly{}
	if err = randomize.Struct(seed, o, nbDemandHourlyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize NBDemandHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(nbDemandHourlyColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := NBDemandHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNBDemandHourliesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NBDemandHourly{}
	if err = randomize.Struct(seed, o, nbDemandHourlyDBTypes, true, nbDemandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NBDemandHourly struct: %s", err)
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

func testNBDemandHourliesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NBDemandHourly{}
	if err = randomize.Struct(seed, o, nbDemandHourlyDBTypes, true, nbDemandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NBDemandHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NBDemandHourlySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testNBDemandHourliesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NBDemandHourly{}
	if err = randomize.Struct(seed, o, nbDemandHourlyDBTypes, true, nbDemandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NBDemandHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := NBDemandHourlies().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	nbDemandHourlyDBTypes = map[string]string{`Time`: `timestamp without time zone`, `DemandPartnerID`: `character varying`, `DemandPartnerPlacementID`: `character varying`, `PublisherID`: `character varying`, `Domain`: `character varying`, `Os`: `character varying`, `Country`: `character varying`, `DeviceType`: `character varying`, `PlacementType`: `character varying`, `Size`: `character varying`, `RequestType`: `character varying`, `PaymentType`: `character varying`, `BidRequests`: `bigint`, `BidResponses`: `bigint`, `AvgBidPrice`: `double precision`, `DPFee`: `double precision`, `AuctionWins`: `bigint`, `Auction`: `double precision`, `SoldImpressions`: `bigint`, `Revenue`: `double precision`, `DataImpressions`: `bigint`, `DataFee`: `double precision`}
	_                     = bytes.MinRead
)

func testNBDemandHourliesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(nbDemandHourlyPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(nbDemandHourlyAllColumns) == len(nbDemandHourlyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &NBDemandHourly{}
	if err = randomize.Struct(seed, o, nbDemandHourlyDBTypes, true, nbDemandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NBDemandHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NBDemandHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, nbDemandHourlyDBTypes, true, nbDemandHourlyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize NBDemandHourly struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testNBDemandHourliesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(nbDemandHourlyAllColumns) == len(nbDemandHourlyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &NBDemandHourly{}
	if err = randomize.Struct(seed, o, nbDemandHourlyDBTypes, true, nbDemandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NBDemandHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NBDemandHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, nbDemandHourlyDBTypes, true, nbDemandHourlyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize NBDemandHourly struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(nbDemandHourlyAllColumns, nbDemandHourlyPrimaryKeyColumns) {
		fields = nbDemandHourlyAllColumns
	} else {
		fields = strmangle.SetComplement(
			nbDemandHourlyAllColumns,
			nbDemandHourlyPrimaryKeyColumns,
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

	slice := NBDemandHourlySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testNBDemandHourliesUpsert(t *testing.T) {
	t.Parallel()

	if len(nbDemandHourlyAllColumns) == len(nbDemandHourlyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := NBDemandHourly{}
	if err = randomize.Struct(seed, &o, nbDemandHourlyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize NBDemandHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert NBDemandHourly: %s", err)
	}

	count, err := NBDemandHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, nbDemandHourlyDBTypes, false, nbDemandHourlyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize NBDemandHourly struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert NBDemandHourly: %s", err)
	}

	count, err = NBDemandHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
