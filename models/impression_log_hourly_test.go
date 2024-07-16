// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testImpressionLogHourlies(t *testing.T) {
	t.Parallel()

	query := ImpressionLogHourlies()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testImpressionLogHourliesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ImpressionLogHourly{}
	if err = randomize.Struct(seed, o, impressionLogHourlyDBTypes, true, impressionLogHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ImpressionLogHourly struct: %s", err)
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

	count, err := ImpressionLogHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testImpressionLogHourliesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ImpressionLogHourly{}
	if err = randomize.Struct(seed, o, impressionLogHourlyDBTypes, true, impressionLogHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ImpressionLogHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ImpressionLogHourlies().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ImpressionLogHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testImpressionLogHourliesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ImpressionLogHourly{}
	if err = randomize.Struct(seed, o, impressionLogHourlyDBTypes, true, impressionLogHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ImpressionLogHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ImpressionLogHourlySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ImpressionLogHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testImpressionLogHourliesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ImpressionLogHourly{}
	if err = randomize.Struct(seed, o, impressionLogHourlyDBTypes, true, impressionLogHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ImpressionLogHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ImpressionLogHourlyExists(ctx, tx, o.Time, o.PublisherID, o.DemandPartnerID, o.Domain, o.Os, o.Country, o.DeviceType, o.Size, o.IsFirst, o.HadFollowup)
	if err != nil {
		t.Errorf("Unable to check if ImpressionLogHourly exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ImpressionLogHourlyExists to return true, but got false.")
	}
}

func testImpressionLogHourliesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ImpressionLogHourly{}
	if err = randomize.Struct(seed, o, impressionLogHourlyDBTypes, true, impressionLogHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ImpressionLogHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	impressionLogHourlyFound, err := FindImpressionLogHourly(ctx, tx, o.Time, o.PublisherID, o.DemandPartnerID, o.Domain, o.Os, o.Country, o.DeviceType, o.Size, o.IsFirst, o.HadFollowup)
	if err != nil {
		t.Error(err)
	}

	if impressionLogHourlyFound == nil {
		t.Error("want a record, got nil")
	}
}

func testImpressionLogHourliesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ImpressionLogHourly{}
	if err = randomize.Struct(seed, o, impressionLogHourlyDBTypes, true, impressionLogHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ImpressionLogHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ImpressionLogHourlies().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testImpressionLogHourliesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ImpressionLogHourly{}
	if err = randomize.Struct(seed, o, impressionLogHourlyDBTypes, true, impressionLogHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ImpressionLogHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ImpressionLogHourlies().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testImpressionLogHourliesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	impressionLogHourlyOne := &ImpressionLogHourly{}
	impressionLogHourlyTwo := &ImpressionLogHourly{}
	if err = randomize.Struct(seed, impressionLogHourlyOne, impressionLogHourlyDBTypes, false, impressionLogHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ImpressionLogHourly struct: %s", err)
	}
	if err = randomize.Struct(seed, impressionLogHourlyTwo, impressionLogHourlyDBTypes, false, impressionLogHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ImpressionLogHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = impressionLogHourlyOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = impressionLogHourlyTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ImpressionLogHourlies().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testImpressionLogHourliesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	impressionLogHourlyOne := &ImpressionLogHourly{}
	impressionLogHourlyTwo := &ImpressionLogHourly{}
	if err = randomize.Struct(seed, impressionLogHourlyOne, impressionLogHourlyDBTypes, false, impressionLogHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ImpressionLogHourly struct: %s", err)
	}
	if err = randomize.Struct(seed, impressionLogHourlyTwo, impressionLogHourlyDBTypes, false, impressionLogHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ImpressionLogHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = impressionLogHourlyOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = impressionLogHourlyTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ImpressionLogHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func impressionLogHourlyBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ImpressionLogHourly) error {
	*o = ImpressionLogHourly{}
	return nil
}

func impressionLogHourlyAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ImpressionLogHourly) error {
	*o = ImpressionLogHourly{}
	return nil
}

func impressionLogHourlyAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ImpressionLogHourly) error {
	*o = ImpressionLogHourly{}
	return nil
}

func impressionLogHourlyBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ImpressionLogHourly) error {
	*o = ImpressionLogHourly{}
	return nil
}

func impressionLogHourlyAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ImpressionLogHourly) error {
	*o = ImpressionLogHourly{}
	return nil
}

func impressionLogHourlyBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ImpressionLogHourly) error {
	*o = ImpressionLogHourly{}
	return nil
}

func impressionLogHourlyAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ImpressionLogHourly) error {
	*o = ImpressionLogHourly{}
	return nil
}

func impressionLogHourlyBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ImpressionLogHourly) error {
	*o = ImpressionLogHourly{}
	return nil
}

func impressionLogHourlyAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ImpressionLogHourly) error {
	*o = ImpressionLogHourly{}
	return nil
}

func testImpressionLogHourliesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ImpressionLogHourly{}
	o := &ImpressionLogHourly{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, impressionLogHourlyDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ImpressionLogHourly object: %s", err)
	}

	AddImpressionLogHourlyHook(boil.BeforeInsertHook, impressionLogHourlyBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	impressionLogHourlyBeforeInsertHooks = []ImpressionLogHourlyHook{}

	AddImpressionLogHourlyHook(boil.AfterInsertHook, impressionLogHourlyAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	impressionLogHourlyAfterInsertHooks = []ImpressionLogHourlyHook{}

	AddImpressionLogHourlyHook(boil.AfterSelectHook, impressionLogHourlyAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	impressionLogHourlyAfterSelectHooks = []ImpressionLogHourlyHook{}

	AddImpressionLogHourlyHook(boil.BeforeUpdateHook, impressionLogHourlyBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	impressionLogHourlyBeforeUpdateHooks = []ImpressionLogHourlyHook{}

	AddImpressionLogHourlyHook(boil.AfterUpdateHook, impressionLogHourlyAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	impressionLogHourlyAfterUpdateHooks = []ImpressionLogHourlyHook{}

	AddImpressionLogHourlyHook(boil.BeforeDeleteHook, impressionLogHourlyBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	impressionLogHourlyBeforeDeleteHooks = []ImpressionLogHourlyHook{}

	AddImpressionLogHourlyHook(boil.AfterDeleteHook, impressionLogHourlyAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	impressionLogHourlyAfterDeleteHooks = []ImpressionLogHourlyHook{}

	AddImpressionLogHourlyHook(boil.BeforeUpsertHook, impressionLogHourlyBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	impressionLogHourlyBeforeUpsertHooks = []ImpressionLogHourlyHook{}

	AddImpressionLogHourlyHook(boil.AfterUpsertHook, impressionLogHourlyAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	impressionLogHourlyAfterUpsertHooks = []ImpressionLogHourlyHook{}
}

func testImpressionLogHourliesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ImpressionLogHourly{}
	if err = randomize.Struct(seed, o, impressionLogHourlyDBTypes, true, impressionLogHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ImpressionLogHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ImpressionLogHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testImpressionLogHourliesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ImpressionLogHourly{}
	if err = randomize.Struct(seed, o, impressionLogHourlyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ImpressionLogHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(impressionLogHourlyColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ImpressionLogHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testImpressionLogHourliesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ImpressionLogHourly{}
	if err = randomize.Struct(seed, o, impressionLogHourlyDBTypes, true, impressionLogHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ImpressionLogHourly struct: %s", err)
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

func testImpressionLogHourliesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ImpressionLogHourly{}
	if err = randomize.Struct(seed, o, impressionLogHourlyDBTypes, true, impressionLogHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ImpressionLogHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ImpressionLogHourlySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testImpressionLogHourliesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ImpressionLogHourly{}
	if err = randomize.Struct(seed, o, impressionLogHourlyDBTypes, true, impressionLogHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ImpressionLogHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ImpressionLogHourlies().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	impressionLogHourlyDBTypes = map[string]string{`Time`: `timestamp without time zone`, `PublisherID`: `character varying`, `DemandPartnerID`: `character varying`, `Domain`: `character varying`, `Os`: `character varying`, `Country`: `character varying`, `DeviceType`: `character varying`, `Size`: `character varying`, `IsFirst`: `boolean`, `HadFollowup`: `boolean`, `SoldImpressions`: `bigint`, `PubImpressions`: `bigint`, `Cost`: `double precision`, `Revenue`: `double precision`, `DemandPartnerFees`: `double precision`}
	_                          = bytes.MinRead
)

func testImpressionLogHourliesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(impressionLogHourlyPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(impressionLogHourlyAllColumns) == len(impressionLogHourlyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ImpressionLogHourly{}
	if err = randomize.Struct(seed, o, impressionLogHourlyDBTypes, true, impressionLogHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ImpressionLogHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ImpressionLogHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, impressionLogHourlyDBTypes, true, impressionLogHourlyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ImpressionLogHourly struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testImpressionLogHourliesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(impressionLogHourlyAllColumns) == len(impressionLogHourlyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ImpressionLogHourly{}
	if err = randomize.Struct(seed, o, impressionLogHourlyDBTypes, true, impressionLogHourlyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ImpressionLogHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ImpressionLogHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, impressionLogHourlyDBTypes, true, impressionLogHourlyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ImpressionLogHourly struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(impressionLogHourlyAllColumns, impressionLogHourlyPrimaryKeyColumns) {
		fields = impressionLogHourlyAllColumns
	} else {
		fields = strmangle.SetComplement(
			impressionLogHourlyAllColumns,
			impressionLogHourlyPrimaryKeyColumns,
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

	slice := ImpressionLogHourlySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testImpressionLogHourliesUpsert(t *testing.T) {
	t.Parallel()

	if len(impressionLogHourlyAllColumns) == len(impressionLogHourlyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ImpressionLogHourly{}
	if err = randomize.Struct(seed, &o, impressionLogHourlyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ImpressionLogHourly struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ImpressionLogHourly: %s", err)
	}

	count, err := ImpressionLogHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, impressionLogHourlyDBTypes, false, impressionLogHourlyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ImpressionLogHourly struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ImpressionLogHourly: %s", err)
	}

	count, err = ImpressionLogHourlies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
