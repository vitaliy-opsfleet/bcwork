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

func testFactors(t *testing.T) {
	t.Parallel()

	query := Factors()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testFactorsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Factor{}
	if err = randomize.Struct(seed, o, factorDBTypes, true, factorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Factor struct: %s", err)
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

	count, err := Factors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFactorsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Factor{}
	if err = randomize.Struct(seed, o, factorDBTypes, true, factorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Factor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Factors().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Factors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFactorsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Factor{}
	if err = randomize.Struct(seed, o, factorDBTypes, true, factorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Factor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FactorSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Factors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFactorsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Factor{}
	if err = randomize.Struct(seed, o, factorDBTypes, true, factorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Factor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := FactorExists(ctx, tx, o.Publisher, o.Domain, o.Device, o.Country)
	if err != nil {
		t.Errorf("Unable to check if Factor exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FactorExists to return true, but got false.")
	}
}

func testFactorsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Factor{}
	if err = randomize.Struct(seed, o, factorDBTypes, true, factorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Factor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	factorFound, err := FindFactor(ctx, tx, o.Publisher, o.Domain, o.Device, o.Country)
	if err != nil {
		t.Error(err)
	}

	if factorFound == nil {
		t.Error("want a record, got nil")
	}
}

func testFactorsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Factor{}
	if err = randomize.Struct(seed, o, factorDBTypes, true, factorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Factor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Factors().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testFactorsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Factor{}
	if err = randomize.Struct(seed, o, factorDBTypes, true, factorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Factor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Factors().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFactorsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	factorOne := &Factor{}
	factorTwo := &Factor{}
	if err = randomize.Struct(seed, factorOne, factorDBTypes, false, factorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Factor struct: %s", err)
	}
	if err = randomize.Struct(seed, factorTwo, factorDBTypes, false, factorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Factor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = factorOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = factorTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Factors().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFactorsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	factorOne := &Factor{}
	factorTwo := &Factor{}
	if err = randomize.Struct(seed, factorOne, factorDBTypes, false, factorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Factor struct: %s", err)
	}
	if err = randomize.Struct(seed, factorTwo, factorDBTypes, false, factorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Factor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = factorOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = factorTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Factors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func factorBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Factor) error {
	*o = Factor{}
	return nil
}

func factorAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Factor) error {
	*o = Factor{}
	return nil
}

func factorAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Factor) error {
	*o = Factor{}
	return nil
}

func factorBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Factor) error {
	*o = Factor{}
	return nil
}

func factorAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Factor) error {
	*o = Factor{}
	return nil
}

func factorBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Factor) error {
	*o = Factor{}
	return nil
}

func factorAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Factor) error {
	*o = Factor{}
	return nil
}

func factorBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Factor) error {
	*o = Factor{}
	return nil
}

func factorAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Factor) error {
	*o = Factor{}
	return nil
}

func testFactorsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Factor{}
	o := &Factor{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, factorDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Factor object: %s", err)
	}

	AddFactorHook(boil.BeforeInsertHook, factorBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	factorBeforeInsertHooks = []FactorHook{}

	AddFactorHook(boil.AfterInsertHook, factorAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	factorAfterInsertHooks = []FactorHook{}

	AddFactorHook(boil.AfterSelectHook, factorAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	factorAfterSelectHooks = []FactorHook{}

	AddFactorHook(boil.BeforeUpdateHook, factorBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	factorBeforeUpdateHooks = []FactorHook{}

	AddFactorHook(boil.AfterUpdateHook, factorAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	factorAfterUpdateHooks = []FactorHook{}

	AddFactorHook(boil.BeforeDeleteHook, factorBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	factorBeforeDeleteHooks = []FactorHook{}

	AddFactorHook(boil.AfterDeleteHook, factorAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	factorAfterDeleteHooks = []FactorHook{}

	AddFactorHook(boil.BeforeUpsertHook, factorBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	factorBeforeUpsertHooks = []FactorHook{}

	AddFactorHook(boil.AfterUpsertHook, factorAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	factorAfterUpsertHooks = []FactorHook{}
}

func testFactorsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Factor{}
	if err = randomize.Struct(seed, o, factorDBTypes, true, factorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Factor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Factors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFactorsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Factor{}
	if err = randomize.Struct(seed, o, factorDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Factor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(factorColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Factors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFactorsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Factor{}
	if err = randomize.Struct(seed, o, factorDBTypes, true, factorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Factor struct: %s", err)
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

func testFactorsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Factor{}
	if err = randomize.Struct(seed, o, factorDBTypes, true, factorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Factor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FactorSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testFactorsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Factor{}
	if err = randomize.Struct(seed, o, factorDBTypes, true, factorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Factor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Factors().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	factorDBTypes = map[string]string{`Publisher`: `character varying`, `Domain`: `character varying`, `Country`: `character varying`, `Device`: `character varying`, `Factor`: `double precision`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `RuleID`: `character varying`, `DemandPartnerID`: `character varying`, `Browser`: `character varying`, `Os`: `character varying`, `PlacementType`: `character varying`}
	_             = bytes.MinRead
)

func testFactorsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(factorPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(factorAllColumns) == len(factorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Factor{}
	if err = randomize.Struct(seed, o, factorDBTypes, true, factorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Factor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Factors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, factorDBTypes, true, factorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Factor struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testFactorsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(factorAllColumns) == len(factorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Factor{}
	if err = randomize.Struct(seed, o, factorDBTypes, true, factorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Factor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Factors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, factorDBTypes, true, factorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Factor struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(factorAllColumns, factorPrimaryKeyColumns) {
		fields = factorAllColumns
	} else {
		fields = strmangle.SetComplement(
			factorAllColumns,
			factorPrimaryKeyColumns,
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

	slice := FactorSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testFactorsUpsert(t *testing.T) {
	t.Parallel()

	if len(factorAllColumns) == len(factorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Factor{}
	if err = randomize.Struct(seed, &o, factorDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Factor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Factor: %s", err)
	}

	count, err := Factors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, factorDBTypes, false, factorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Factor struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Factor: %s", err)
	}

	count, err = Factors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}