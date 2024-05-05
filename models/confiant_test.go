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

func testConfiants(t *testing.T) {
	t.Parallel()

	query := Confiants()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testConfiantsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Confiant{}
	if err = randomize.Struct(seed, o, confiantDBTypes, true, confiantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Confiant struct: %s", err)
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

	count, err := Confiants().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testConfiantsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Confiant{}
	if err = randomize.Struct(seed, o, confiantDBTypes, true, confiantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Confiant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Confiants().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Confiants().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testConfiantsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Confiant{}
	if err = randomize.Struct(seed, o, confiantDBTypes, true, confiantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Confiant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ConfiantSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Confiants().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testConfiantsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Confiant{}
	if err = randomize.Struct(seed, o, confiantDBTypes, true, confiantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Confiant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ConfiantExists(ctx, tx, o.ConfiantKey)
	if err != nil {
		t.Errorf("Unable to check if Confiant exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ConfiantExists to return true, but got false.")
	}
}

func testConfiantsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Confiant{}
	if err = randomize.Struct(seed, o, confiantDBTypes, true, confiantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Confiant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	confiantFound, err := FindConfiant(ctx, tx, o.ConfiantKey)
	if err != nil {
		t.Error(err)
	}

	if confiantFound == nil {
		t.Error("want a record, got nil")
	}
}

func testConfiantsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Confiant{}
	if err = randomize.Struct(seed, o, confiantDBTypes, true, confiantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Confiant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Confiants().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testConfiantsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Confiant{}
	if err = randomize.Struct(seed, o, confiantDBTypes, true, confiantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Confiant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Confiants().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testConfiantsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	confiantOne := &Confiant{}
	confiantTwo := &Confiant{}
	if err = randomize.Struct(seed, confiantOne, confiantDBTypes, false, confiantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Confiant struct: %s", err)
	}
	if err = randomize.Struct(seed, confiantTwo, confiantDBTypes, false, confiantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Confiant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = confiantOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = confiantTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Confiants().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testConfiantsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	confiantOne := &Confiant{}
	confiantTwo := &Confiant{}
	if err = randomize.Struct(seed, confiantOne, confiantDBTypes, false, confiantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Confiant struct: %s", err)
	}
	if err = randomize.Struct(seed, confiantTwo, confiantDBTypes, false, confiantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Confiant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = confiantOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = confiantTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Confiants().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func confiantBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Confiant) error {
	*o = Confiant{}
	return nil
}

func confiantAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Confiant) error {
	*o = Confiant{}
	return nil
}

func confiantAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Confiant) error {
	*o = Confiant{}
	return nil
}

func confiantBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Confiant) error {
	*o = Confiant{}
	return nil
}

func confiantAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Confiant) error {
	*o = Confiant{}
	return nil
}

func confiantBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Confiant) error {
	*o = Confiant{}
	return nil
}

func confiantAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Confiant) error {
	*o = Confiant{}
	return nil
}

func confiantBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Confiant) error {
	*o = Confiant{}
	return nil
}

func confiantAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Confiant) error {
	*o = Confiant{}
	return nil
}

func testConfiantsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Confiant{}
	o := &Confiant{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, confiantDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Confiant object: %s", err)
	}

	AddConfiantHook(boil.BeforeInsertHook, confiantBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	confiantBeforeInsertHooks = []ConfiantHook{}

	AddConfiantHook(boil.AfterInsertHook, confiantAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	confiantAfterInsertHooks = []ConfiantHook{}

	AddConfiantHook(boil.AfterSelectHook, confiantAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	confiantAfterSelectHooks = []ConfiantHook{}

	AddConfiantHook(boil.BeforeUpdateHook, confiantBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	confiantBeforeUpdateHooks = []ConfiantHook{}

	AddConfiantHook(boil.AfterUpdateHook, confiantAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	confiantAfterUpdateHooks = []ConfiantHook{}

	AddConfiantHook(boil.BeforeDeleteHook, confiantBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	confiantBeforeDeleteHooks = []ConfiantHook{}

	AddConfiantHook(boil.AfterDeleteHook, confiantAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	confiantAfterDeleteHooks = []ConfiantHook{}

	AddConfiantHook(boil.BeforeUpsertHook, confiantBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	confiantBeforeUpsertHooks = []ConfiantHook{}

	AddConfiantHook(boil.AfterUpsertHook, confiantAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	confiantAfterUpsertHooks = []ConfiantHook{}
}

func testConfiantsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Confiant{}
	if err = randomize.Struct(seed, o, confiantDBTypes, true, confiantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Confiant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Confiants().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testConfiantsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Confiant{}
	if err = randomize.Struct(seed, o, confiantDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Confiant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(confiantColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Confiants().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testConfiantToOnePublisherUsingPublisher(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Confiant
	var foreign Publisher

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, confiantDBTypes, false, confiantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Confiant struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, publisherDBTypes, false, publisherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Publisher struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.PublisherID = foreign.PublisherID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Publisher().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.PublisherID != foreign.PublisherID {
		t.Errorf("want: %v, got %v", foreign.PublisherID, check.PublisherID)
	}

	ranAfterSelectHook := false
	AddPublisherHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Publisher) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := ConfiantSlice{&local}
	if err = local.L.LoadPublisher(ctx, tx, false, (*[]*Confiant)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Publisher == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Publisher = nil
	if err = local.L.LoadPublisher(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Publisher == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testConfiantToOneSetOpPublisherUsingPublisher(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Confiant
	var b, c Publisher

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, confiantDBTypes, false, strmangle.SetComplement(confiantPrimaryKeyColumns, confiantColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, publisherDBTypes, false, strmangle.SetComplement(publisherPrimaryKeyColumns, publisherColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, publisherDBTypes, false, strmangle.SetComplement(publisherPrimaryKeyColumns, publisherColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Publisher{&b, &c} {
		err = a.SetPublisher(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Publisher != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Confiants[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.PublisherID != x.PublisherID {
			t.Error("foreign key was wrong value", a.PublisherID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.PublisherID))
		reflect.Indirect(reflect.ValueOf(&a.PublisherID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PublisherID != x.PublisherID {
			t.Error("foreign key was wrong value", a.PublisherID, x.PublisherID)
		}
	}
}

func testConfiantsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Confiant{}
	if err = randomize.Struct(seed, o, confiantDBTypes, true, confiantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Confiant struct: %s", err)
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

func testConfiantsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Confiant{}
	if err = randomize.Struct(seed, o, confiantDBTypes, true, confiantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Confiant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ConfiantSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testConfiantsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Confiant{}
	if err = randomize.Struct(seed, o, confiantDBTypes, true, confiantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Confiant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Confiants().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	confiantDBTypes = map[string]string{`ConfiantKey`: `character varying`, `PublisherID`: `character varying`, `Domain`: `character varying`, `Rate`: `double precision`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`}
	_               = bytes.MinRead
)

func testConfiantsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(confiantPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(confiantAllColumns) == len(confiantPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Confiant{}
	if err = randomize.Struct(seed, o, confiantDBTypes, true, confiantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Confiant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Confiants().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, confiantDBTypes, true, confiantPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Confiant struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testConfiantsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(confiantAllColumns) == len(confiantPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Confiant{}
	if err = randomize.Struct(seed, o, confiantDBTypes, true, confiantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Confiant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Confiants().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, confiantDBTypes, true, confiantPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Confiant struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(confiantAllColumns, confiantPrimaryKeyColumns) {
		fields = confiantAllColumns
	} else {
		fields = strmangle.SetComplement(
			confiantAllColumns,
			confiantPrimaryKeyColumns,
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

	slice := ConfiantSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testConfiantsUpsert(t *testing.T) {
	t.Parallel()

	if len(confiantAllColumns) == len(confiantPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Confiant{}
	if err = randomize.Struct(seed, &o, confiantDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Confiant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Confiant: %s", err)
	}

	count, err := Confiants().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, confiantDBTypes, false, confiantPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Confiant struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Confiant: %s", err)
	}

	count, err = Confiants().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
