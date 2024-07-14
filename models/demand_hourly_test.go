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

func testDemandHourlies(t *testing.T) {
	t.Parallel()

	query := DemandHourlies()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDemandHourliesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DemandHourly{}
	if err = randomize.Struct(seed, o, demandHourlyDBTypes, true, demandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandHourly struct: %s", err)
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

	count, err := DemandHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDemandHourliesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DemandHourly{}
	if err = randomize.Struct(seed, o, demandHourlyDBTypes, true, demandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DemandHourlies().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DemandHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDemandHourliesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DemandHourly{}
	if err = randomize.Struct(seed, o, demandHourlyDBTypes, true, demandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DemandHourlySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DemandHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDemandHourliesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DemandHourly{}
	if err = randomize.Struct(seed, o, demandHourlyDBTypes, true, demandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DemandHourlyExists(ctx, tx, o.Time, o.PublisherID, o.DemandPartnerID, o.Domain)
	if err != nil {
		t.Errorf("Unable to check if DemandHourly exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DemandHourlyExists to return true, but got false.")
	}
}

func testDemandHourliesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DemandHourly{}
	if err = randomize.Struct(seed, o, demandHourlyDBTypes, true, demandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	demandHourlyFound, err := FindDemandHourly(ctx, tx, o.Time, o.PublisherID, o.DemandPartnerID, o.Domain)
	if err != nil {
		t.Error(err)
	}

	if demandHourlyFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDemandHourliesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DemandHourly{}
	if err = randomize.Struct(seed, o, demandHourlyDBTypes, true, demandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DemandHourlies().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDemandHourliesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DemandHourly{}
	if err = randomize.Struct(seed, o, demandHourlyDBTypes, true, demandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DemandHourlies().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDemandHourliesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	demandHourlyOne := &DemandHourly{}
	demandHourlyTwo := &DemandHourly{}
	if err = randomize.Struct(seed, demandHourlyOne, demandHourlyDBTypes, false, demandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandHourly struct: %s", err)
	}
	if err = randomize.Struct(seed, demandHourlyTwo, demandHourlyDBTypes, false, demandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = demandHourlyOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = demandHourlyTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DemandHourlies().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDemandHourliesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	demandHourlyOne := &DemandHourly{}
	demandHourlyTwo := &DemandHourly{}
	if err = randomize.Struct(seed, demandHourlyOne, demandHourlyDBTypes, false, demandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandHourly struct: %s", err)
	}
	if err = randomize.Struct(seed, demandHourlyTwo, demandHourlyDBTypes, false, demandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = demandHourlyOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = demandHourlyTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DemandHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func demandHourlyBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *DemandHourly) error {
	*o = DemandHourly{}
	return nil
}

func demandHourlyAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *DemandHourly) error {
	*o = DemandHourly{}
	return nil
}

func demandHourlyAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *DemandHourly) error {
	*o = DemandHourly{}
	return nil
}

func demandHourlyBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DemandHourly) error {
	*o = DemandHourly{}
	return nil
}

func demandHourlyAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DemandHourly) error {
	*o = DemandHourly{}
	return nil
}

func demandHourlyBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DemandHourly) error {
	*o = DemandHourly{}
	return nil
}

func demandHourlyAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DemandHourly) error {
	*o = DemandHourly{}
	return nil
}

func demandHourlyBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DemandHourly) error {
	*o = DemandHourly{}
	return nil
}

func demandHourlyAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DemandHourly) error {
	*o = DemandHourly{}
	return nil
}

func testDemandHourliesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &DemandHourly{}
	o := &DemandHourly{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, demandHourlyDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DemandHourly object: %s", err)
	}

	AddDemandHourlyHook(boil.BeforeInsertHook, demandHourlyBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	demandHourlyBeforeInsertHooks = []DemandHourlyHook{}

	AddDemandHourlyHook(boil.AfterInsertHook, demandHourlyAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	demandHourlyAfterInsertHooks = []DemandHourlyHook{}

	AddDemandHourlyHook(boil.AfterSelectHook, demandHourlyAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	demandHourlyAfterSelectHooks = []DemandHourlyHook{}

	AddDemandHourlyHook(boil.BeforeUpdateHook, demandHourlyBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	demandHourlyBeforeUpdateHooks = []DemandHourlyHook{}

	AddDemandHourlyHook(boil.AfterUpdateHook, demandHourlyAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	demandHourlyAfterUpdateHooks = []DemandHourlyHook{}

	AddDemandHourlyHook(boil.BeforeDeleteHook, demandHourlyBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	demandHourlyBeforeDeleteHooks = []DemandHourlyHook{}

	AddDemandHourlyHook(boil.AfterDeleteHook, demandHourlyAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	demandHourlyAfterDeleteHooks = []DemandHourlyHook{}

	AddDemandHourlyHook(boil.BeforeUpsertHook, demandHourlyBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	demandHourlyBeforeUpsertHooks = []DemandHourlyHook{}

	AddDemandHourlyHook(boil.AfterUpsertHook, demandHourlyAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	demandHourlyAfterUpsertHooks = []DemandHourlyHook{}
}

func testDemandHourliesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DemandHourly{}
	if err = randomize.Struct(seed, o, demandHourlyDBTypes, true, demandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DemandHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDemandHourliesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DemandHourly{}
	if err = randomize.Struct(seed, o, demandHourlyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DemandHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(demandHourlyColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DemandHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDemandHourliesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DemandHourly{}
	if err = randomize.Struct(seed, o, demandHourlyDBTypes, true, demandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandHourly struct: %s", err)
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

func testDemandHourliesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DemandHourly{}
	if err = randomize.Struct(seed, o, demandHourlyDBTypes, true, demandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DemandHourlySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDemandHourliesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DemandHourly{}
	if err = randomize.Struct(seed, o, demandHourlyDBTypes, true, demandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DemandHourlies().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	demandHourlyDBTypes = map[string]string{`Time`: `timestamp without time zone`, `DemandPartnerID`: `character varying`, `PublisherID`: `character varying`, `Domain`: `character varying`, `BidRequest`: `bigint`, `BidResponse`: `bigint`, `BidPrice`: `double precision`, `Impression`: `bigint`, `Revenue`: `double precision`, `DemandPartnerFee`: `double precision`, `DataFee`: `double precision`}
	_                   = bytes.MinRead
)

func testDemandHourliesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(demandHourlyPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(demandHourlyAllColumns) == len(demandHourlyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DemandHourly{}
	if err = randomize.Struct(seed, o, demandHourlyDBTypes, true, demandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DemandHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, demandHourlyDBTypes, true, demandHourlyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DemandHourly struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDemandHourliesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(demandHourlyAllColumns) == len(demandHourlyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DemandHourly{}
	if err = randomize.Struct(seed, o, demandHourlyDBTypes, true, demandHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DemandHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, demandHourlyDBTypes, true, demandHourlyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DemandHourly struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(demandHourlyAllColumns, demandHourlyPrimaryKeyColumns) {
		fields = demandHourlyAllColumns
	} else {
		fields = strmangle.SetComplement(
			demandHourlyAllColumns,
			demandHourlyPrimaryKeyColumns,
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

	slice := DemandHourlySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDemandHourliesUpsert(t *testing.T) {
	t.Parallel()

	if len(demandHourlyAllColumns) == len(demandHourlyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DemandHourly{}
	if err = randomize.Struct(seed, &o, demandHourlyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DemandHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DemandHourly: %s", err)
	}

	count, err := DemandHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, demandHourlyDBTypes, false, demandHourlyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DemandHourly struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DemandHourly: %s", err)
	}

	count, err = DemandHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
